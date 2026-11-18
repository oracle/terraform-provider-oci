// Copyright (c) 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

// Explicit live inputs are prerequisites: missing inputs must never become a
// skipped test or successful acceptance evidence.
func databaseAckRequiredInput(t *testing.T, name string) string {
	t.Helper()
	value := os.Getenv("TF_VAR_" + name)
	if strings.TrimSpace(value) == "" {
		t.Fatalf("TF_VAR_%s must be selected from live preflight before this acceptance test", name)
	}
	return value
}

func databaseAckCloudConfig(t *testing.T, needsPassword bool) string {
	t.Helper()
	clusterID := databaseAckRequiredInput(t, "ack_test_cloud_vm_cluster_id")
	if !strings.HasPrefix(clusterID, "ocid1.cloudvmcluster.") {
		t.Fatal("TF_VAR_ack_test_cloud_vm_cluster_id must identify the disposable, test-owned regular ExaCS cluster")
	}
	databaseAckRequiredInput(t, "ack_test_db_home_version")
	config := acctest.ProviderTestConfig() + `
variable "ack_test_cloud_vm_cluster_id" { type = string }
variable "ack_test_db_home_version" { type = string }
data "oci_database_cloud_vm_cluster" "ack_parent_cluster" {
 cloud_vm_cluster_id = var.ack_test_cloud_vm_cluster_id
}
`
	if needsPassword {
		databaseAckRequiredInput(t, "ack_test_admin_password")
		config += `
variable "ack_test_admin_password" {
 type = string
 sensitive = true
}
`
	}
	return config
}

func databaseAckRequireReplacedID(resourceName string, previousID *string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		currentID, err := acctest.FromInstanceState(state, resourceName, "id")
		if err != nil {
			return err
		}
		if currentID == "" {
			return fmt.Errorf("%s has no ID", resourceName)
		}
		if *previousID != "" && currentID == *previousID {
			return fmt.Errorf("%s was not replaced after its create-only acknowledgment changed", resourceName)
		}
		*previousID = currentID
		return nil
	}
}

// PRECHECK action resources use the database ID as their Terraform ID. Their
// creation is proven by a new successful history entry, not a different ID.
func databaseAckRequireNewPrecheckHistory(dataSource, target string, previousIDs map[string]bool) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		root := state.RootModule()
		entry, ok := root.Resources[dataSource]
		if !ok || entry.Primary == nil {
			return fmt.Errorf("missing PRECHECK history datasource %s", dataSource)
		}
		attrs := entry.Primary.Attributes
		count, err := strconv.Atoi(attrs["database_upgrade_history_entries.#"])
		if err != nil || count == 0 {
			return fmt.Errorf("no successful PRECHECK history entry was returned")
		}
		foundNew := false
		currentIDs := map[string]bool{}
		for i := 0; i < count; i++ {
			prefix := fmt.Sprintf("database_upgrade_history_entries.%d.", i)
			id := attrs[prefix+"id"]
			if id == "" {
				return fmt.Errorf("PRECHECK history entry %d has no ID", i)
			}
			currentIDs[id] = true
			if attrs[prefix+"action"] == "PRECHECK" && attrs[prefix+"state"] == "SUCCEEDED" && attrs[prefix+"target_db_version"] == target && !previousIDs[id] {
				foundNew = true
			}
		}
		if !foundNew {
			return fmt.Errorf("acknowledgment input did not produce a new successful PRECHECK history entry for the selected target")
		}
		for id := range currentIDs {
			previousIDs[id] = true
		}
		return nil
	}
}

func databaseAckBaseDBDestroy(state *terraform.State) error {
	client := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	found := false
	for _, item := range state.RootModule().Resources {
		if item.Type != "oci_database_db_system" {
			continue
		}
		found = true
		id := item.Primary.ID
		response, err := client.GetDbSystem(context.Background(), oci_database.GetDbSystemRequest{DbSystemId: &id})
		if err != nil {
			if serviceError, ok := common.IsServiceError(err); ok && serviceError.GetHTTPStatusCode() == 404 {
				continue
			}
			return err
		}
		if response.LifecycleState != oci_database.DbSystemLifecycleStateTerminated {
			return fmt.Errorf("owned BaseDB system %s remains %s after normal cleanup", id, response.LifecycleState)
		}
	}
	if !found {
		return fmt.Errorf("owned BaseDB source system was absent from the test state")
	}
	return nil
}
