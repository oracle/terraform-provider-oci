// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	//"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseDbNodeSnapshotManagementRepresentation = map[string]interface{}{
		"exadb_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_exadb_vm_cluster.test_exadb_vm_cluster.id}`},
		"name":                acctest.Representation{RepType: acctest.Required, Create: `snapshot1`},
		"source_dbnode_ids":   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_exadb_vm_cluster.test_exadb_vm_cluster.node_resource[*].node_id}`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"example-tag-namespace-all.example-tag": "value"}},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseDbNodeSnapshotManagementIgnoreTagsRepresentation},
	}

	DatabaseDbNodeSnapshotManagementIgnoreTagsRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `freeform_tags`}},
	}

	DatabaseDbNodeSnapshotManagementResourceDependencies = `
        variable exadb_vm_cluster_id {}

        data "oci_database_exadb_vm_cluster" "test_exadb_vm_cluster" {
            exadb_vm_cluster_id = var.exadb_vm_cluster_id
        }
    `
)

// issue-routing-tag: database/default
func TestDatabaseDbNodeSnapshotManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDbNodeSnapshotManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_db_node_snapshot_management.test_db_node_snapshot_management"

	//var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseDbNodeSnapshotManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_node_snapshot_management", "test_db_node_snapshot_management", acctest.Optional, acctest.Create, DatabaseDbNodeSnapshotManagementRepresentation), "database", "dbNodeSnapshotManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseDbNodeSnapshotManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_node_snapshot_management", "test_db_node_snapshot_management", acctest.Required, acctest.Create, DatabaseDbNodeSnapshotManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "exadb_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "snapshot1"),
				resource.TestCheckResourceAttr(resourceName, "source_dbnode_ids.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "snapshots.#", "3"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseDbNodeSnapshotManagementResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseDbNodeSnapshotManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_node_snapshot_management", "test_db_node_snapshot_management", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseDbNodeSnapshotManagementRepresentation, map[string]interface{}{
						"name": acctest.Representation{RepType: acctest.Required, Create: `snapshot2`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "exadb_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "snapshot2"),
				resource.TestCheckResourceAttr(resourceName, "source_dbnode_ids.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "snapshots.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.Department", "Finance"),
				// commenting out this code due to make test-compile failure
				//func(s *terraform.State) (err error) {
				//	resId, err = acctest.FromInstanceState(s, resourceName, "id")
				//	if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
				//		if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
				//			return errExport
				//		}
				//	}
				//	return err
				//},
			),
		},
	})
}
