// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

type databaseNoopSigner struct{}

func (databaseNoopSigner) Sign(*http.Request) error {
	return nil
}

func TestDatabaseDbSystemGetDbHomeInfoUsesDatabaseCurrentDbHome(t *testing.T) {
	const (
		databaseID    = "D1"
		oldDbHomeID   = "H1"
		currentHomeID = "H2"
	)

	var requestMutex sync.Mutex
	var requestPaths []string
	server := httptest.NewServer(http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		requestMutex.Lock()
		requestPaths = append(requestPaths, request.URL.Path)
		requestMutex.Unlock()

		responseWriter.Header().Set("Content-Type", "application/json")
		switch request.URL.Path {
		case "/20160918/databases/" + databaseID:
			_ = json.NewEncoder(responseWriter).Encode(map[string]interface{}{
				"id":             databaseID,
				"compartmentId":  "C1",
				"dbName":         "PROD",
				"dbUniqueName":   "PROD_UNIQUE",
				"lifecycleState": "AVAILABLE",
				"dbHomeId":       currentHomeID,
				"dbSystemId":     "S1",
			})
		case "/20160918/dbHomes/" + currentHomeID:
			_ = json.NewEncoder(responseWriter).Encode(map[string]interface{}{
				"id":             currentHomeID,
				"compartmentId":  "C1",
				"displayName":    "current-home",
				"lifecycleState": "AVAILABLE",
				"dbVersion":      "19.24.0.0",
				"dbHomeLocation": "/u01/app/oracle/product/19.0.0/dbhome_2",
				"dbSystemId":     "S1",
			})
		default:
			http.Error(responseWriter, "unexpected request", http.StatusNotFound)
		}
	}))
	defer server.Close()

	client := oci_database.DatabaseClient{BaseClient: oci_common.BaseClient{
		HTTPClient: server.Client(),
		Signer:     databaseNoopSigner{},
		Host:       server.URL,
		BasePath:   "20160918",
		UserAgent:  "terraform-provider-oci-test",
	}}

	resourceData := schema.TestResourceDataRaw(t, DatabaseDbSystemResource().Schema, map[string]interface{}{
		"availability_domain": "AD-1",
		"compartment_id":      "C1",
		"db_home": []interface{}{map[string]interface{}{
			"database": []interface{}{map[string]interface{}{
				"admin_password": "NotARealPassword1!",
				"db_name":        "PROD",
			}},
		}},
		"hostname":        "dbsystem",
		"shape":           "VM.Standard2.1",
		"ssh_public_keys": []interface{}{"ssh-rsa test-key"},
		"subnet_id":       "SUBNET1",
	})
	if err := resourceData.Set("db_home", []interface{}{map[string]interface{}{
		"id":         oldDbHomeID,
		"db_version": "19.18.0.0",
		"database": []interface{}{map[string]interface{}{
			"admin_password": "NotARealPassword1!",
			"db_name":        "PROD",
			"id":             databaseID,
		}},
	}}); err != nil {
		t.Fatalf("failed to seed old DB home state: %v", err)
	}

	crud := &DatabaseDbSystemResourceCrud{
		BaseCrud: tfresource.BaseCrud{D: resourceData},
		Client:   &client,
		Res: &oci_database.DbSystem{
			Id:            oci_common.String("S1"),
			CompartmentId: oci_common.String("C1"),
		},
	}

	if err := crud.getDbHomeInfo(context.Background()); err != nil {
		t.Fatalf("getDbHomeInfo returned an error: %v", err)
	}
	if err := crud.SetData(); err != nil {
		t.Fatalf("SetData returned an error: %v", err)
	}

	requestMutex.Lock()
	gotRequestPaths := append([]string(nil), requestPaths...)
	requestMutex.Unlock()
	wantRequestPaths := []string{
		"/20160918/databases/" + databaseID,
		"/20160918/dbHomes/" + currentHomeID,
	}
	if !reflect.DeepEqual(gotRequestPaths, wantRequestPaths) {
		t.Fatalf("unexpected refresh requests: got %v, want %v", gotRequestPaths, wantRequestPaths)
	}

	if got := resourceData.Get("db_home.0.id"); got != currentHomeID {
		t.Errorf("DB home ID in state = %q, want %q", got, currentHomeID)
	}
	if got := resourceData.Get("db_home.0.db_version"); got != "19.24.0.0" {
		t.Errorf("DB version in state = %q, want %q", got, "19.24.0.0")
	}
	if got := resourceData.Get("db_home.0.database.0.id"); got != databaseID {
		t.Errorf("database ID in state = %q, want unchanged ID %q", got, databaseID)
	}
}
