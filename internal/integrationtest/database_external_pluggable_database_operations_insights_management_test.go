// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	externalPluggableDatabaseOperationsInsightsManagementRepresentation = map[string]interface{}{
		"external_database_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_external_database_connector.test_external_database_connector.id}`},
		"external_pluggable_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_external_pluggable_database.test_external_pluggable_database.id}`},
		"enable_operations_insights":     acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}

	ExternalPluggableDatabaseOperationsInsightsManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database", "test_external_container_database", acctest.Required, acctest.Create, externalContainerDatabaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database", "test_external_pluggable_database", acctest.Required, acctest.Create, externalPluggable1DatabaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_pluggable_database_connector", acctest.Required, acctest.Create, externalPluggableDatabaseConnectorRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", acctest.Required, acctest.Create, externalContainerDatabaseConnectorRepresentation)
)

// issue-routing-tag: database/default
func TestDatabaseExternalPluggableDatabaseOperationsInsightsManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExternalPluggableDatabaseOperationsInsightsManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_external_pluggable_database_operations_insights_management.test_external_pluggable_database_operations_insights_management"
	resourcePDB := "oci_database_external_pluggable_database.test_external_pluggable_database"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ExternalPluggableDatabaseOperationsInsightsManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database_operations_insights_management", "test_external_pluggable_database_operations_insights_management", acctest.Required, acctest.Create, externalPluggableDatabaseOperationsInsightsManagementRepresentation), "database", "externalPluggableDatabaseOperationsInsightsManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Enablement of PDB
		{
			Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseOperationsInsightsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database_operations_insights_management", "test_external_pluggable_database_operations_insights_management", acctest.Required, acctest.Create, externalPluggableDatabaseOperationsInsightsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_pluggable_database_id"),
			),
		},

		// Verify Enablement of PDB
		{
			Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseOperationsInsightsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database_operations_insights_management", "test_external_pluggable_database_operations_insights_management", acctest.Required, acctest.Create, externalPluggableDatabaseOperationsInsightsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourcePDB, "operations_insights_config.0.operations_insights_status", "ENABLED"),
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseOperationsInsightsManagementResourceDependencies,
		},
		// Enablement of PDB
		{
			Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseOperationsInsightsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database_operations_insights_management", "test_external_pluggable_database_operations_insights_management", acctest.Required, acctest.Create, externalPluggableDatabaseOperationsInsightsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_pluggable_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
			),
		},
		// Verify Enablement of PDB
		{
			Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseOperationsInsightsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database_operations_insights_management", "test_external_pluggable_database_operations_insights_management", acctest.Required, acctest.Create, externalPluggableDatabaseOperationsInsightsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourcePDB, "operations_insights_config.0.operations_insights_status", "ENABLED"),
			),
		},

		// Disablement of PDB
		{
			Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseOperationsInsightsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database_operations_insights_management", "test_external_pluggable_database_operations_insights_management", acctest.Optional, acctest.Update, externalPluggableDatabaseOperationsInsightsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_pluggable_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
			),
		},
		// Verify Disablement of PDB
		{
			Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseOperationsInsightsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database_operations_insights_management", "test_external_pluggable_database_operations_insights_management", acctest.Optional, acctest.Update, externalPluggableDatabaseOperationsInsightsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourcePDB, "operations_insights_config.0.operations_insights_status", "NOT_ENABLED"),
			),
		},
	})
}
