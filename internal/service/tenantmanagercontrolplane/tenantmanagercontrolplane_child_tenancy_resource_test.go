// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package tenantmanagercontrolplane

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_tenantmanagercontrolplane "github.com/oracle/oci-go-sdk/v65/tenantmanagercontrolplane"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

// These fixtures exercise SDK request/response handling entirely in memory.
// No credentials, sockets, OCI endpoints, or real tenancy creation are used.
type childTenancyFixtureTransport func(*http.Request) (*http.Response, error)

func (f childTenancyFixtureTransport) Do(r *http.Request) (*http.Response, error) {
	return f(r)
}

type childTenancyFixtureSigner struct{}

func (childTenancyFixtureSigner) Sign(*http.Request) error { return nil }

func TestUnitChildTenancyUnchangedConfigurationHasEmptyPlan(t *testing.T) {
	for _, governance := range []string{"", "OPTED_IN"} {
		t.Run("governance="+governance, func(t *testing.T) {
			config := map[string]interface{}{
				"admin_email": "test@example.invalid", "compartment_id": "parent-id",
				"home_region": "us-seattle-1", "tenancy_name": "offlinechild",
			}
			// Match the post-create state: omitted request-only fields have no
			// response value because Create returns only a work request and ID.
			state := &terraform.InstanceState{ID: "child-id", Attributes: map[string]string{"id": "child-id"}}
			for name, value := range config {
				state.Attributes[name] = value.(string)
			}
			if governance != "" {
				config["governance_status"] = governance
				state.Attributes["governance_status"] = governance
			}
			diff, err := TenantmanagercontrolplaneChildTenancyResource().Diff(context.Background(), state, terraform.NewResourceConfigRaw(config), nil)
			if err != nil {
				t.Fatal(err)
			}
			if diff != nil && !diff.Empty() {
				t.Fatalf("unchanged creation inputs must not plan an update or replacement: %#v", diff.Attributes)
			}
		})
	}
}

func TestUnitChildTenancyCreateWorkRequest(t *testing.T) {
	for _, tc := range []struct {
		name      string
		status    string
		resources string
		wantError bool
		optional  bool
	}{
		{"success", "SUCCEEDED", `[{"entityType":"childTenancy","actionType":"CREATED","identifier":"child-id"}]`, false, false},
		{"optional inputs", "SUCCEEDED", `[{"entityType":"childTenancy","actionType":"CREATED","identifier":"child-id"}]`, false, true},
		{"failed", "FAILED", `[{"entityType":"childTenancy","actionType":"RELATED","identifier":"child-id"}]`, true, false},
		{"canceled", "CANCELED", `[]`, true, false},
		{"missing identifier", "SUCCEEDED", `[{"entityType":"childTenancy","actionType":"CREATED"}]`, true, false},
		{"wrong action", "SUCCEEDED", `[{"entityType":"childTenancy","actionType":"UPDATED","identifier":"child-id"}]`, true, false},
		{"missing entity type", "SUCCEEDED", `[{"actionType":"CREATED","identifier":"child-id"}]`, true, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			values := map[string]interface{}{
				"admin_email": "test@example.invalid", "compartment_id": "parent-id",
				"home_region": "us-seattle-1", "tenancy_name": "offlinechild",
			}
			if tc.optional {
				values["features"] = []interface{}{"CORE"}
				values["governance_status"] = "OPTED_IN"
				values["policy_name"] = "child_admin"
			}
			d := schema.TestResourceDataRaw(t, TenantmanagercontrolplaneChildTenancyResource().Schema, values)
			posts, gets := 0, 0
			base := oci_common.DefaultBaseClientWithSigner(childTenancyFixtureSigner{})
			base.Host, base.BasePath = "https://fixture.invalid", "20230401"
			base.HTTPClient = childTenancyFixtureTransport(func(r *http.Request) (*http.Response, error) {
				if r.URL.Host != "fixture.invalid" {
					return nil, fmt.Errorf("unexpected host %s", r.URL.Host)
				}
				status, body := http.StatusOK, ""
				headers := http.Header{"Content-Type": []string{"application/json"}}
				switch {
				case r.Method == http.MethodPost && r.URL.Path == "/20230401/childTenancies":
					posts++
					var payload map[string]interface{}
					if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
						t.Fatal(err)
					}
					for key, expected := range map[string]string{"adminEmail": "test@example.invalid", "compartmentId": "parent-id", "homeRegion": "us-seattle-1", "tenancyName": "offlinechild"} {
						if payload[key] != expected {
							t.Errorf("%s = %v, want %s", key, payload[key], expected)
						}
					}
					if tc.optional && (payload["policyName"] != "child_admin" || payload["governanceStatus"] != "OPTED_IN" || fmt.Sprint(payload["features"]) != "[CORE]") {
						t.Errorf("optional fields not mapped: %v", payload)
					}
					if r.Header.Get("opc-retry-token") == "" {
						t.Error("missing SDK retry token")
					}
					status = http.StatusAccepted
					headers.Set("opc-work-request-id", "work-id")
				case r.Method == http.MethodGet && r.URL.Path == "/20230401/workRequests/work-id":
					gets++
					body = fmt.Sprintf(`{"id":"work-id","status":%q,"timeFinished":%q,"resources":%s}`, tc.status, time.Now().UTC().Format(time.RFC3339), tc.resources)
				case r.Method == http.MethodGet && r.URL.Path == "/20230401/workRequests/work-id/errors":
					body = `{"items":[{"message":"fixture work request error"}]}`
				default:
					return nil, fmt.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
				}
				return &http.Response{StatusCode: status, Header: headers, Body: io.NopCloser(strings.NewReader(body)), Request: r}, nil
			})
			crud := &TenantmanagercontrolplaneChildTenancyResourceCrud{
				BaseCrud:          tfresource.BaseCrud{D: d},
				Client:            &oci_tenantmanagercontrolplane.OrganizationClient{BaseClient: base},
				WorkRequestClient: &oci_tenantmanagercontrolplane.WorkRequestClient{BaseClient: base},
			}
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()
			err := tfresource.CreateResourceWithContext(ctx, d, crud)
			if (err != nil) != tc.wantError {
				t.Fatalf("CreateResourceWithContext() error = %v, wantError %v", err, tc.wantError)
			}
			if !tc.wantError && (d.Id() != "child-id" || crud.ID() != "child-id") {
				t.Errorf("created ID lost: state=%q ID()=%q", d.Id(), crud.ID())
			}
			if posts != 1 || gets < 2 {
				t.Errorf("expected exactly one create and work-request polling, got POST=%d GET=%d", posts, gets)
			}
		})
	}
}
