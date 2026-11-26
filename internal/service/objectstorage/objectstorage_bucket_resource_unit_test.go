// Copyright (c) 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package objectstorage

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_object_storage "github.com/oracle/oci-go-sdk/v65/objectstorage"
)

type bucketResourceTestSigner struct{}

func (bucketResourceTestSigner) Sign(*http.Request) error { return nil }

func TestObjectStorageBucketDelete(t *testing.T) {
	tests := []struct {
		name       string
		force      bool
		wantError  bool
		wantEvents []string
	}{
		{
			name:       "non-forced delete leaves objects and returns bucket error",
			wantError:  true,
			wantEvents: []string{"delete bucket"},
		},
		{
			name:       "forced delete removes versions before bucket",
			force:      true,
			wantEvents: []string{"list versions", "delete object version", "delete bucket"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var events []string
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				switch {
				case r.Method == http.MethodGet && r.URL.Path == "/n/namespace/b/bucket/objectversions":
					events = append(events, "list versions")
					w.Header().Set("Content-Type", "application/json")
					_, _ = w.Write([]byte(`{"items":[{"name":"object.txt","versionId":"version-1","isDeleteMarker":false}]}`))
				case r.Method == http.MethodDelete && r.URL.Path == "/n/namespace/b/bucket/o/object.txt":
					events = append(events, "delete object version")
					if r.URL.Query().Get("versionId") != "version-1" {
						http.Error(w, "version ID missing", http.StatusBadRequest)
						return
					}
					w.WriteHeader(http.StatusNoContent)
				case r.Method == http.MethodDelete && r.URL.Path == "/n/namespace/b/bucket":
					events = append(events, "delete bucket")
					if test.force {
						w.WriteHeader(http.StatusNoContent)
						return
					}
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusConflict)
					_, _ = w.Write([]byte(`{"code":"BucketNotEmpty","message":"bucket is not empty"}`))
				default:
					http.Error(w, fmt.Sprintf("unexpected request %s %s", r.Method, r.URL), http.StatusNotFound)
				}
			}))
			defer server.Close()

			client := &oci_object_storage.ObjectStorageClient{BaseClient: oci_common.DefaultBaseClientWithSigner(bucketResourceTestSigner{})}
			client.Host = server.URL
			resourceData := schema.TestResourceDataRaw(t, ObjectStorageBucketResource().Schema, map[string]interface{}{
				"name":         "bucket",
				"namespace":    "namespace",
				"force_delete": test.force,
			})
			crud := &ObjectStorageBucketResourceCrud{Client: client}
			crud.D = resourceData

			err := crud.Delete()
			if (err != nil) != test.wantError {
				t.Fatalf("Delete() error = %v, wantError %t", err, test.wantError)
			}
			if got := fmt.Sprint(events); got != fmt.Sprint(test.wantEvents) {
				t.Fatalf("request order = %v, want %v", events, test.wantEvents)
			}
		})
	}
}
