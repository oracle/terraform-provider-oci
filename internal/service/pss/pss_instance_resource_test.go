// Copyright (c) 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package pss

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_pss "github.com/oracle/oci-go-sdk/v65/pss"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func TestPssInstanceDefinedTagsPlan(t *testing.T) {
	previous := tfresource.DefinedTagsToSuppress
	t.Cleanup(func() { tfresource.DefinedTagsToSuppress = previous })

	for _, tc := range []struct {
		name          string
		ignoredTags   []string
		configuredTag string
		wantChange    bool
	}{
		{
			name:          "auto tags reproduce the nonempty plan",
			configuredTag: "value",
			wantChange:    true,
		},
		{
			name:          "ignoring only Oracle auto tags produces an empty plan",
			ignoredTags:   []string{"Oracle-Tags.CreatedBy", "Oracle-Tags.CreatedOn"},
			configuredTag: "value",
		},
		{
			name:          "custom tag changes remain planned",
			ignoredTags:   []string{"Oracle-Tags.CreatedBy", "Oracle-Tags.CreatedOn"},
			configuredTag: "updatedValue",
			wantChange:    true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			tfresource.DefinedTagsToSuppress = tc.ignoredTags
			resource := PssInstanceResource()
			// The acceptance response includes Oracle-added tags alongside the requested tag.
			// This fixture preserves that shape with sanitized identifiers and values.
			stateData := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{
				"compartment_id":                "test-compartment",
				"admin_identity_domain_user_id": "test-admin",
				"identity_domain_url":           "https://identity.example.com",
				"display_name":                  "test-instance",
				"state":                         "ACTIVE",
				"freeform_tags":                 map[string]interface{}{"Department": "Finance"},
				"defined_tags": map[string]interface{}{
					"test_namespace.test_tag": "value",
					"Oracle-Tags.CreatedBy":   "test-user",
					"Oracle-Tags.CreatedOn":   "2026-01-01T00:00:00Z",
				},
			})
			stateData.SetId("test-instance-id")
			// Computed fields are populated by refresh rather than configuration.
			if err := stateData.Set("system_tags", map[string]interface{}{}); err != nil {
				t.Fatal(err)
			}
			config := terraform.NewResourceConfigRaw(map[string]interface{}{
				"compartment_id":                "test-compartment",
				"admin_identity_domain_user_id": "test-admin",
				"identity_domain_url":           "https://identity.example.com",
				"display_name":                  "test-instance",
				"freeform_tags":                 map[string]interface{}{"Department": "Finance"},
				"defined_tags": map[string]interface{}{
					"test_namespace.test_tag": tc.configuredTag,
				},
			})
			diff, err := resource.Diff(context.Background(), stateData.State(), config, nil)
			if err != nil {
				t.Fatal(err)
			}
			gotChange := diff != nil && !diff.Empty()
			if gotChange != tc.wantChange {
				t.Fatalf("plan has changes = %t, want %t; diff: %#v", gotChange, tc.wantChange, diff)
			}
			if diff != nil && diff.RequiresNew() {
				t.Fatal("tag changes must not recreate the instance")
			}
			if tc.configuredTag == "updatedValue" {
				change := diff.Attributes["defined_tags.test_namespace.test_tag"]
				if change == nil || change.Old != "value" || change.New != "updatedValue" {
					t.Fatalf("custom tag update was lost: %#v", change)
				}
			}
		})
	}
}

func TestPssInstancesDataSourceStateIdentity(t *testing.T) {
	instanceID := "test-instance-id"
	data := schema.TestResourceDataRaw(t, PssInstancesDataSource().Schema, map[string]interface{}{
		"compartment_id": "test-compartment",
		"id":             instanceID,
	})
	sync := PssInstancesDataSourceCrud{
		D: data,
		Res: &oci_pss.ListInstancesResponse{
			InstanceCollection: oci_pss.InstanceCollection{
				Items: []oci_pss.InstanceSummary{{
					Id:             oci_common.String(instanceID),
					CompartmentId:  oci_common.String("test-compartment"),
					LifecycleState: oci_pss.InstanceLifecycleStateActive,
				}},
			},
		},
	}
	if err := sync.SetData(); err != nil {
		t.Fatal(err)
	}
	state := data.State()
	if state == nil || state.ID == "" || state.ID == instanceID {
		t.Fatalf("expected an independent data-source state identity: %#v", state)
	}
	if got := state.Attributes["id"]; got != state.ID {
		t.Fatalf("data-source id attribute = %q, want its state ID %q", got, state.ID)
	}
	if got := state.Attributes["instance_collection.0.items.0.id"]; got != instanceID {
		t.Fatalf("collection instance ID = %q, want %q", got, instanceID)
	}
}
