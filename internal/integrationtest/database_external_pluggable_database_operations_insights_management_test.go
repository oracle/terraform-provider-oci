// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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

	DatabaseExternalPluggableDatabaseOperationsInsightsManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database", "test_external_container_database", acctest.Required, acctest.Create, DatabaseExternalContainerDatabaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database", "test_external_pluggable_database", acctest.Required, acctest.Create, DatabaseExternalPluggable1DatabaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_pluggable_database_connector", acctest.Required, acctest.Create, DatabaseExternalPluggableDatabaseConnectorRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", acctest.Required, acctest.Create, DatabaseExternalContainerDatabaseConnectorRepresentation)
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
