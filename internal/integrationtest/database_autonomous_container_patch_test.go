// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatabaseDatabaseAutonomousContainerPatchDataSourceRepresentation = map[string]interface{}{
		"autonomous_container_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
		"compartment_id":                   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"autonomous_patch_type":            acctest.Representation{RepType: acctest.Required, Create: `QUARTERLY`},
	}

	DatabaseAutonomousContainerPatchResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Required, acctest.Create, DatabasePatchAutonomousContainerDatabaseRepresentation) +
		DatabaseCloudAutonomousVmClusterResourceConfig

	ExaccDatabaseAutonomousContainerPatchResourceConfig = ACDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Update, ExaccMRACDatabaseRepresentation)

	DatabasePatchAutonomousContainerDatabaseRepresentation = map[string]interface{}{
		"version_preference":             acctest.Representation{RepType: acctest.Optional, Create: `LATEST_RELEASE_UPDATE`, Update: `NEXT_RELEASE_UPDATE`},
		"display_name":                   acctest.Representation{RepType: acctest.Required, Create: `containerDatabase2`, Update: `displayName2`},
		"patch_model":                    acctest.Representation{RepType: acctest.Required, Create: `RELEASE_UPDATES`, Update: `RELEASE_UPDATE_REVISIONS`},
		"db_version":                     acctest.Representation{RepType: acctest.Required, Create: `19.19.0.1.0`},
		"cloud_autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id}`},
		"backup_config":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: ACDatabaseBackupConfigRepresentation},
		"compartment_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"defined_tags":                   acctest.Representation{RepType: acctest.Optional, Create: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"})}`, Update: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedValue"})}`},
		"freeform_tags":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_automatic_failover_enabled":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"kms_key_id":                     acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"maintenance_window_details":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsRepresentation},
		"service_level_agreement_type":   acctest.Representation{RepType: acctest.Optional, Create: `STANDARD`},
		"vault_id":                       acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_kms_vault.test_vault.id}`},
		"db_name":                        acctest.Representation{RepType: acctest.Optional, Create: `DBNAME`},
		"is_dst_file_update_enabled":     acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
)

// issue-routing-tag: database/default
func TestDatabaseAutonomousContainerPatchResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousContainerPatchResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_container_patches.test_autonomous_container_patches"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_patches", "test_autonomous_container_patches", acctest.Required, acctest.Create, DatabaseDatabaseAutonomousContainerPatchDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousContainerPatchResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_patch_type", "QUARTERLY"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.0.autonomous_patch_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.0.description"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.0.patch_model"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.0.quarter"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.0.type"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.0.version"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.0.year"),
			),
		},
	})
}

// issue-routing-tag: database/default
func TestExaccDatabaseAutonomousContainerPatchResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousContainerPatchResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_container_patches.test_autonomous_container_patches"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_patches", "test_autonomous_container_patches", acctest.Required, acctest.Create, DatabaseDatabaseAutonomousContainerPatchDataSourceRepresentation) +
				compartmentIdVariableStr + ExaccDatabaseAutonomousContainerPatchResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_patch_type", "QUARTERLY"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.0.autonomous_patch_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.0.description"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.0.patch_model"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.0.quarter"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.0.type"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.0.version"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_patches.0.year"),
			),
		},
	})
}
