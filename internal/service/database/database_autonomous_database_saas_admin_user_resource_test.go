// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func TestDatabaseAutonomousDatabaseSaasAdminUserReadDisabledClearsState(t *testing.T) {
	const autonomousDatabaseID = "ADB1"

	server := httptest.NewServer(http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			t.Errorf("request method = %q, want %q", request.Method, http.MethodPost)
		}
		wantPath := "/20160918/autonomousDatabases/" + autonomousDatabaseID + "/actions/getSaasAdminUserStatus"
		if request.URL.Path != wantPath {
			t.Errorf("request path = %q, want %q", request.URL.Path, wantPath)
		}

		responseWriter.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(responseWriter).Encode(map[string]bool{"isEnabled": false})
	}))
	defer server.Close()

	client := oci_database.DatabaseClient{BaseClient: oci_common.BaseClient{
		HTTPClient: server.Client(),
		Signer:     databaseNoopSigner{},
		Host:       server.URL,
		BasePath:   "20160918",
		UserAgent:  "terraform-provider-oci-test",
	}}

	resourceData := schema.TestResourceDataRaw(t, DatabaseAutonomousDatabaseSaasAdminUserResource().Schema, map[string]interface{}{
		"autonomous_database_id": autonomousDatabaseID,
		"access_type":            "READ_WRITE",
		"duration":               1,
		"password":               "NotARealPassword1!",
	})
	resourceData.SetId("saas-admin-user-admin-user-" + autonomousDatabaseID)

	crud := &DatabaseAutonomousDatabaseSaasAdminUserResourceCrud{
		BaseCrud: tfresource.BaseCrud{D: resourceData},
		Client:   &client,
	}

	if err := tfresource.ReadResourceWithContext(context.Background(), crud); err != nil {
		t.Fatalf("ReadResourceWithContext returned an error: %v", err)
	}
	if resourceData.Id() != "" {
		t.Fatalf("resource ID = %q, want empty", resourceData.Id())
	}
}
