// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"sync/atomic"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/globalvar"
)

func TestDatabaseDatabaseResource_importedReadOnlyDestroyPreservesState(t *testing.T) {
	const databaseID = "database-test-id"

	var deleteAttempts int32
	var deleted atomic.Bool
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.Header().Set("opc-request-id", "test-request")

		switch r.Method {
		case http.MethodGet:
			if deleted.Load() {
				w.WriteHeader(http.StatusNotFound)
				_, _ = w.Write([]byte(`{"code":"NotAuthorizedOrNotFound","message":"not found"}`))
				return
			}
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{
				"id":"` + databaseID + `",
				"compartmentId":"compartment-test-id",
				"dbName":"testdb",
				"dbUniqueName":"testdb_unique",
				"lifecycleState":"AVAILABLE"
			}`))
		case http.MethodDelete:
			if atomic.AddInt32(&deleteAttempts, 1) == 1 {
				w.WriteHeader(http.StatusNotFound)
				_, _ = w.Write([]byte(`{"code":"NotAuthorizedOrNotFound","message":"not found"}`))
				return
			}
			deleted.Store(true)
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Fatalf("unexpected request %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	configureDatabaseReadOnlyDestroyTest(t, server.URL)

	config := acctest.ProviderTestConfig()
	resourceName := "oci_database_database.test_database"
	var checkedRetainedState atomic.Bool

	acctest.ResourceTest(t, checkDatabaseReadOnlyDestroyRetainedState(t, resourceName, databaseID, &checkedRetainedState), []resource.TestStep{
		{
			Config: config + `
resource "oci_database_database" "test_database" {
  db_home_id = "db-home-test-id"
  source     = "NONE"

  database {}
}
`,
			ImportState:        true,
			ImportStateId:      databaseID,
			ImportStatePersist: true,
			ResourceName:       resourceName,
		},
		{
			Config:      config,
			ExpectError: regexp.MustCompile("(?i)(NotAuthorizedOrNotFound|not found)"),
		},
	})

	if !checkedRetainedState.Load() {
		t.Fatalf("expected retained state id check to run for %s", resourceName)
	}

	if atomic.LoadInt32(&deleteAttempts) < 2 {
		t.Fatalf("expected final cleanup to retry deleting %s from retained state after failed destroy, got %d delete attempts", resourceName, deleteAttempts)
	}
}

func TestDatabaseDbHomeResource_importedReadOnlyDestroyPreservesState(t *testing.T) {
	const dbHomeID = "db-home-test-id"

	var deleteAttempts int32
	var deleted atomic.Bool
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.Header().Set("opc-request-id", "test-request")

		switch r.Method {
		case http.MethodGet:
			if deleted.Load() {
				w.WriteHeader(http.StatusNotFound)
				_, _ = w.Write([]byte(`{"code":"NotAuthorizedOrNotFound","message":"not found"}`))
				return
			}
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{
				"id":"` + dbHomeID + `",
				"compartmentId":"compartment-test-id",
				"displayName":"test-db-home",
				"lifecycleState":"AVAILABLE",
				"dbVersion":"19.0.0.0",
				"dbHomeLocation":"/u01/app/oracle/product/19.0.0/dbhome_1"
			}`))
		case http.MethodDelete:
			if atomic.AddInt32(&deleteAttempts, 1) == 1 {
				w.WriteHeader(http.StatusNotFound)
				_, _ = w.Write([]byte(`{"code":"NotAuthorizedOrNotFound","message":"not found"}`))
				return
			}
			deleted.Store(true)
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Fatalf("unexpected request %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	configureDatabaseReadOnlyDestroyTest(t, server.URL)

	config := acctest.ProviderTestConfig()
	resourceName := "oci_database_db_home.test_db_home"
	var checkedRetainedState atomic.Bool

	acctest.ResourceTest(t, checkDatabaseReadOnlyDestroyRetainedState(t, resourceName, dbHomeID, &checkedRetainedState), []resource.TestStep{
		{
			Config: config + `
resource "oci_database_db_home" "test_db_home" {
}
`,
			ImportState:        true,
			ImportStateId:      dbHomeID,
			ImportStatePersist: true,
			ResourceName:       resourceName,
		},
		{
			Config:      config,
			ExpectError: regexp.MustCompile("(?i)(NotAuthorizedOrNotFound|not found)"),
		},
	})

	if !checkedRetainedState.Load() {
		t.Fatalf("expected retained state id check to run for %s", resourceName)
	}

	if atomic.LoadInt32(&deleteAttempts) < 2 {
		t.Fatalf("expected final cleanup to retry deleting %s from retained state after failed destroy, got %d delete attempts", resourceName, deleteAttempts)
	}
}

func TestDatabaseDatabaseResource_importedMissingDestroyRemovesState(t *testing.T) {
	const databaseID = "database-missing-test-id"

	var deletePhase atomic.Bool
	var deletePhaseReads int32
	var destroyReadNotFound atomic.Bool
	var deleteAttempts int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.Header().Set("opc-request-id", "test-request")

		switch r.Method {
		case http.MethodGet:
			if deletePhase.Load() && atomic.AddInt32(&deletePhaseReads, 1) > 1 {
				destroyReadNotFound.Store(true)
				w.WriteHeader(http.StatusNotFound)
				_, _ = w.Write([]byte(`{"code":"NotAuthorizedOrNotFound","message":"not found"}`))
				return
			}
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{
				"id":"` + databaseID + `",
				"compartmentId":"compartment-test-id",
				"dbName":"testdb",
				"dbUniqueName":"testdb_unique",
				"lifecycleState":"AVAILABLE"
			}`))
		case http.MethodDelete:
			atomic.AddInt32(&deleteAttempts, 1)
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(`{"code":"NotAuthorizedOrNotFound","message":"not found"}`))
		default:
			t.Fatalf("unexpected request %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	configureDatabaseReadOnlyDestroyTest(t, server.URL)

	config := acctest.ProviderTestConfig()
	resourceName := "oci_database_database.test_database"
	var checkedRemovedState atomic.Bool

	acctest.ResourceTest(t, nil, []resource.TestStep{
		{
			Config: config + `
resource "oci_database_database" "test_database" {
  db_home_id = "db-home-test-id"
  source     = "NONE"

  database {}
}
`,
			ImportState:        true,
			ImportStateId:      databaseID,
			ImportStatePersist: true,
			ResourceName:       resourceName,
		},
		{
			PreConfig: func() {
				deletePhase.Store(true)
			},
			Config: config,
			Check:  checkDatabaseReadOnlyDestroyRemovedState(resourceName, &checkedRemovedState),
		},
	})

	if !checkedRemovedState.Load() {
		t.Fatalf("expected removed state check to run for %s", resourceName)
	}

	if atomic.LoadInt32(&deletePhaseReads) < 2 {
		t.Fatalf("expected config-removal refresh read to succeed and destroy-time read to return 404 for %s, got %d reads", resourceName, deletePhaseReads)
	}

	if !destroyReadNotFound.Load() {
		t.Fatalf("expected destroy-time read to return 404 for %s", resourceName)
	}

	if atomic.LoadInt32(&deleteAttempts) != 0 {
		t.Fatalf("expected no delete attempt for %s after pre-delete read returned 404, got %d", resourceName, deleteAttempts)
	}
}

func TestDatabaseDbHomeResource_importedMissingDestroyRemovesState(t *testing.T) {
	const dbHomeID = "db-home-missing-test-id"

	var deletePhase atomic.Bool
	var deletePhaseReads int32
	var destroyReadNotFound atomic.Bool
	var deleteAttempts int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.Header().Set("opc-request-id", "test-request")

		switch r.Method {
		case http.MethodGet:
			if deletePhase.Load() && atomic.AddInt32(&deletePhaseReads, 1) > 1 {
				destroyReadNotFound.Store(true)
				w.WriteHeader(http.StatusNotFound)
				_, _ = w.Write([]byte(`{"code":"NotAuthorizedOrNotFound","message":"not found"}`))
				return
			}
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{
				"id":"` + dbHomeID + `",
				"compartmentId":"compartment-test-id",
				"displayName":"test-db-home",
				"lifecycleState":"AVAILABLE",
				"dbVersion":"19.0.0.0",
				"dbHomeLocation":"/u01/app/oracle/product/19.0.0/dbhome_1"
			}`))
		case http.MethodDelete:
			atomic.AddInt32(&deleteAttempts, 1)
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(`{"code":"NotAuthorizedOrNotFound","message":"not found"}`))
		default:
			t.Fatalf("unexpected request %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	configureDatabaseReadOnlyDestroyTest(t, server.URL)

	config := acctest.ProviderTestConfig()
	resourceName := "oci_database_db_home.test_db_home"
	var checkedRemovedState atomic.Bool

	acctest.ResourceTest(t, nil, []resource.TestStep{
		{
			Config: config + `
resource "oci_database_db_home" "test_db_home" {
}
`,
			ImportState:        true,
			ImportStateId:      dbHomeID,
			ImportStatePersist: true,
			ResourceName:       resourceName,
		},
		{
			PreConfig: func() {
				deletePhase.Store(true)
			},
			Config: config,
			Check:  checkDatabaseReadOnlyDestroyRemovedState(resourceName, &checkedRemovedState),
		},
	})

	if !checkedRemovedState.Load() {
		t.Fatalf("expected removed state check to run for %s", resourceName)
	}

	if atomic.LoadInt32(&deletePhaseReads) < 2 {
		t.Fatalf("expected config-removal refresh read to succeed and destroy-time read to return 404 for %s, got %d reads", resourceName, deletePhaseReads)
	}

	if !destroyReadNotFound.Load() {
		t.Fatalf("expected destroy-time read to return 404 for %s", resourceName)
	}

	if atomic.LoadInt32(&deleteAttempts) != 0 {
		t.Fatalf("expected no delete attempt for %s after pre-delete read returned 404, got %d", resourceName, deleteAttempts)
	}
}

func checkDatabaseReadOnlyDestroyRetainedState(t *testing.T, resourceName, expectedID string, checked *atomic.Bool) resource.TestCheckFunc {
	t.Helper()

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("expected %s to remain in state after failed read-only destroy", resourceName)
		}
		if rs.Primary == nil {
			return fmt.Errorf("expected %s to have primary state after failed read-only destroy", resourceName)
		}
		if got := rs.Primary.ID; got != expectedID {
			return fmt.Errorf("expected %s primary id to remain %q after failed read-only destroy, got %q", resourceName, expectedID, got)
		}
		if got := rs.Primary.Attributes["id"]; got != expectedID {
			return fmt.Errorf("expected %s state attribute id to remain %q after failed read-only destroy, got %q", resourceName, expectedID, got)
		}

		checked.Store(true)
		return nil
	}
}

func checkDatabaseReadOnlyDestroyRemovedState(resourceName string, checked *atomic.Bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if _, ok := s.RootModule().Resources[resourceName]; ok {
			return fmt.Errorf("expected %s to be removed from state when pre-delete read also returned 404", resourceName)
		}

		checked.Store(true)
		return nil
	}
}

func configureDatabaseReadOnlyDestroyTest(t *testing.T, databaseHost string) {
	t.Helper()

	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("failed to generate test key: %v", err)
	}
	privateKey := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	})

	t.Setenv("tenancy_ocid", "tenancy-test-id")
	t.Setenv("user_ocid", "user-test-id")
	t.Setenv("fingerprint", "test-fingerprint")
	t.Setenv("private_key", string(privateKey))
	t.Setenv("region", "us-phoenix-1")
	t.Setenv("compartment_ocid", "compartment-test-id")
	t.Setenv("compartment_id_for_create", "compartment-test-id")
	t.Setenv("compartment_id_for_update", "compartment-test-id")
	t.Setenv("tags_import_if_exists", "false")
	t.Setenv(globalvar.ClientHostOverridesEnv, "oci_database.DatabaseClient="+databaseHost)
}
