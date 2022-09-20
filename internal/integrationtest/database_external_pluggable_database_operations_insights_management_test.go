// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseExternalPluggableDatabaseOperationsInsightsManagementRepresentation = map[string]interface{}{
		"external_database_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_external_database_connector.test_external_database_connector.id}`},
		"external_pluggable_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_external_pluggable_database.test_external_pluggable_database.id}`},
		"enable_operations_insights":     acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}

<<<<<<< ours
	DatabaseExternalPluggableDatabaseOperationsInsightsManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_agent", "test_agent", acctest.Required, acctest.Create, agentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_environment", "test_environment", acctest.Required, acctest.Create, environmentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_backup", "test_backup", acctest.Required, acctest.Create, backupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Required, acctest.Create, databaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home", acctest.Required, acctest.Create, dbHomeRepresentation) +
		BackupResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", Optional, Create, backupDestinationNFSRepresentation) +
		GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Optional, Update, RepresentationCopyWithNewProperties(exadataInfrastructureActivateRepresentation, map[string]interface{}{"activation_file": Representation{RepType: Optional, Update: activationFilePath}})) +
		GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", Optional, Update, vmClusterNetworkValidateRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Required, acctest.Create, exadataInfrastructureRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database", "test_external_container_database", acctest.Required, acctest.Create, externalContainerDatabaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", acctest.Required, acctest.Create, externalDatabaseConnectorRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database", "test_external_pluggable_database", acctest.Required, acctest.Create, externalPluggableDatabaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", acctest.Required, acctest.Create, vmClusterNetworkRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Required, acctest.Create, vmClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Required, acctest.Create, vaultRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_vault_secret", "test_secret", acctest.Required, acctest.Create, secretRepresentation)
=======
	DatabaseExternalPluggableDatabaseOperationsInsightsManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database", "test_external_container_database", acctest.Required, acctest.Create, DatabaseExternalContainerDatabaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database", "test_external_pluggable_database", acctest.Required, acctest.Create, DatabaseExternalPluggable1DatabaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_pluggable_database_connector", acctest.Required, acctest.Create, DatabaseExternalPluggableDatabaseConnectorRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", acctest.Required, acctest.Create, DatabaseExternalContainerDatabaseConnectorRepresentation)
>>>>>>> theirs
)

// issue-routing-tag: database/default
func TestDatabaseExternalPluggableDatabaseOperationsInsightsManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExternalPluggableDatabaseOperationsInsightsManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	agentId := utils.GetEnvSettingWithBlankDefault("connector_agent_id")
	agentIdVariableStr := fmt.Sprintf("variable \"agent_id\" { default = \"%s\" }\n", agentId)

	resourceName := "oci_database_external_pluggable_database_operations_insights_management.test_external_pluggable_database_operations_insights_management"
	resourcePDB := "oci_database_external_pluggable_database.test_external_pluggable_database"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+agentIdVariableStr+DatabaseExternalPluggableDatabaseOperationsInsightsManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database_operations_insights_management", "test_external_pluggable_database_operations_insights_management", acctest.Required, acctest.Create, DatabaseExternalPluggableDatabaseOperationsInsightsManagementRepresentation), "database", "externalPluggableDatabaseOperationsInsightsManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Enablement of PDB
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + DatabaseExternalPluggableDatabaseOperationsInsightsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database_operations_insights_management", "test_external_pluggable_database_operations_insights_management", acctest.Required, acctest.Create, DatabaseExternalPluggableDatabaseOperationsInsightsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_pluggable_database_id"),
			),
		},

		// Verify Enablement of PDB
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + DatabaseExternalPluggableDatabaseOperationsInsightsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database_operations_insights_management", "test_external_pluggable_database_operations_insights_management", acctest.Required, acctest.Create, DatabaseExternalPluggableDatabaseOperationsInsightsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourcePDB, "operations_insights_config.0.operations_insights_status", "ENABLED"),
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + DatabaseExternalPluggableDatabaseOperationsInsightsManagementResourceDependencies,
		},
		// Enablement of PDB
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + DatabaseExternalPluggableDatabaseOperationsInsightsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database_operations_insights_management", "test_external_pluggable_database_operations_insights_management", acctest.Required, acctest.Create, DatabaseExternalPluggableDatabaseOperationsInsightsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_pluggable_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
			),
		},
		// Verify Enablement of PDB
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + DatabaseExternalPluggableDatabaseOperationsInsightsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database_operations_insights_management", "test_external_pluggable_database_operations_insights_management", acctest.Required, acctest.Create, DatabaseExternalPluggableDatabaseOperationsInsightsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourcePDB, "operations_insights_config.0.operations_insights_status", "ENABLED"),
			),
		},

		// Disablement of PDB
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + DatabaseExternalPluggableDatabaseOperationsInsightsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database_operations_insights_management", "test_external_pluggable_database_operations_insights_management", acctest.Optional, acctest.Update, DatabaseExternalPluggableDatabaseOperationsInsightsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_pluggable_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
			),
		},
		// Verify Disablement of PDB
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + DatabaseExternalPluggableDatabaseOperationsInsightsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database_operations_insights_management", "test_external_pluggable_database_operations_insights_management", acctest.Optional, acctest.Update, DatabaseExternalPluggableDatabaseOperationsInsightsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourcePDB, "operations_insights_config.0.operations_insights_status", "NOT_ENABLED"),
			),
		},
	})
}
