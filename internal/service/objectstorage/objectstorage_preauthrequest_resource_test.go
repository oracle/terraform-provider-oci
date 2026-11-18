// Copyright (c) 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package objectstorage

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_object_storage "github.com/oracle/oci-go-sdk/v65/objectstorage"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

// issue-routing-tag: object_storage/default
func TestObjectStoragePreauthenticatedRequestRefresh(t *testing.T) {
	for _, tt := range []struct {
		name          string
		listingAction string
		objectField   string
	}{
		{name: "ListObjects", listingAction: "ListObjects"},
		{name: "Deny", listingAction: "Deny"},
		{name: "object_name", objectField: "object_name"},
		{name: "legacy_object", objectField: "object"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			config := map[string]interface{}{
				"access_type": "AnyObjectRead",
				"bucket":      "test-bucket", "namespace": "test-namespace",
				"name": "test-par", "time_expires": "2038-01-01T12:00:00Z",
			}
			response := map[string]interface{}{
				"id": "test-par-id", "name": "test-par", "accessType": "AnyObjectRead",
				"timeCreated": "2026-01-01T00:00:00Z", "timeExpires": config["time_expires"],
			}
			if tt.listingAction != "" {
				config["bucket_listing_action"] = tt.listingAction
				response["bucketListingAction"] = tt.listingAction
			}
			if tt.objectField != "" {
				config["access_type"] = "ObjectRead"
				config[tt.objectField] = "test-object"
				response["accessType"] = "ObjectRead"
				response["objectName"] = "test-object"
			}

			creates, reads := 0, 0
			client := &oci_object_storage.ObjectStorageClient{
				BaseClient: oci_common.DefaultBaseClientWithSigner(preauthRequestTestSigner{}),
			}
			client.Host = "https://objectstorage.invalid"
			// Handle every SDK request in memory; no credentials or OCI endpoints are used.
			client.HTTPClient = preauthRequestTestTransport(func(r *http.Request) (*http.Response, error) {
				switch {
				case r.Method == http.MethodPost && r.URL.Path == "/n/test-namespace/b/test-bucket/p":
					creates++
					var request oci_object_storage.CreatePreauthenticatedRequestDetails
					if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
						return nil, err
					}
					if string(request.BucketListingAction) != tt.listingAction {
						t.Errorf("create listing action = %q, want %q", request.BucketListingAction, tt.listingAction)
					}
					response["accessUri"] = "/p/test-token/n/test-namespace/b/test-bucket/o/"
					response["fullPath"] = "https://objectstorage.invalid/p/test-token/n/test-namespace/b/test-bucket/o/"
				case r.Method == http.MethodGet && r.URL.Path == "/n/test-namespace/b/test-bucket/p/test-par-id":
					reads++
					// GET returns a summary without the create-only URL fields.
					delete(response, "accessUri")
					delete(response, "fullPath")
				default:
					return nil, fmt.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
				}
				body, err := json.Marshal(response)
				if err != nil {
					return nil, err
				}
				return &http.Response{
					StatusCode: http.StatusOK, Header: http.Header{"Content-Type": []string{"application/json"}},
					Body: io.NopCloser(strings.NewReader(string(body))), Request: r,
				}, nil
			})

			resource := ObjectStoragePreauthenticatedRequestResource()
			data := schema.TestResourceDataRaw(t, resource.Schema, config)
			crud := &ObjectStoragePreauthenticatedRequestResourceCrud{
				BaseCrud: tfresource.BaseCrud{D: data}, Client: client,
			}
			if err := tfresource.CreateResource(data, crud); err != nil {
				t.Fatal(err)
			}
			id := data.Id()
			for _, phase := range []string{"create", "refresh", "second_refresh", "import"} {
				t.Run(phase, func(t *testing.T) {
					if phase == "import" {
						data = schema.TestResourceDataRaw(t, resource.Schema, nil)
						data.SetId(id)
						imported, err := resource.Importer.State(data, nil)
						if err != nil || len(imported) != 1 {
							t.Fatalf("import returned %d resources, error: %v", len(imported), err)
						}
						data = imported[0]
						crud.D = data
					}
					if phase != "create" {
						if err := tfresource.ReadResource(crud); err != nil {
							t.Fatal(err)
						}
					}
					if got := data.Get("bucket_listing_action"); got != tt.listingAction {
						t.Errorf("listing action = %q, want %q", got, tt.listingAction)
					}
					if data.Id() != id {
						t.Errorf("resource ID = %q, want %q", data.Id(), id)
					}
					if tt.objectField != "" {
						for _, field := range []string{"object", "object_name"} {
							if got := data.Get(field); got != "test-object" {
								t.Errorf("%s = %q, want test-object", field, got)
							}
						}
					}
					for _, field := range []string{"access_uri", "full_path"} {
						if got := data.Get(field).(string); (got != "") != (phase != "import") {
							t.Errorf("unexpected %s after %s: %q", field, phase, got)
						}
					}
					// Exercise Terraform's schema diff against the original configuration.
					diff, err := resource.Diff(context.Background(), data.State(), terraform.NewResourceConfigRaw(config), nil)
					if err != nil {
						t.Fatal(err)
					}
					if diff != nil && !diff.Empty() {
						t.Errorf("expected no changes after %s, got: %#v", phase, diff)
					}
				})
			}
			if creates != 1 || reads != 3 {
				t.Errorf("requests: %d creates, %d reads; want 1 create, 3 reads", creates, reads)
			}
		})
	}
}

type preauthRequestTestTransport func(*http.Request) (*http.Response, error)

func (f preauthRequestTestTransport) Do(r *http.Request) (*http.Response, error) { return f(r) }

type preauthRequestTestSigner struct{}

func (preauthRequestTestSigner) Sign(*http.Request) error { return nil }
