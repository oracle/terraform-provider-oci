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
	externalNonContainerDatabaseOperationsInsightsManagementRepresentation = map[string]interface{}{
		"external_database_connector_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_database_external_database_connector.test_external_database_connector.id}`},
		"external_non_container_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_external_non_container_database.test_external_non_container_database.id}`},
		"enable_operations_insights":         acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}

	ExternalNonContainerDatabaseOperationsInsightsManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", acctest.Required, acctest.Create, externalDatabaseConnectorRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_non_container_database", "test_external_non_container_database", acctest.Required, acctest.Create, externalNonContainerDatabaseRepresentation)
)

// issue-routing-tag: database/default
func TestDatabaseExternalNonContainerDatabaseOperationsInsightsManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExternalNonContainerDatabaseOperationsInsightsManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_external_non_container_database_operations_insights_management.test_external_non_container_database_operations_insights_management"
	resourceNonCDB := "oci_database_external_non_container_database.test_external_non_container_database"
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ExternalNonContainerDatabaseOperationsInsightsManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_non_container_database_operations_insights_management", "test_external_non_container_database_operations_insights_management", acctest.Required, acctest.Create, externalNonContainerDatabaseOperationsInsightsManagementRepresentation), "database", "externalNonContainerDatabaseOperationsInsightsManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ExternalNonContainerDatabaseOperationsInsightsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_non_container_database_operations_insights_management", "test_external_non_container_database_operations_insights_management", acctest.Required, acctest.Create, externalNonContainerDatabaseOperationsInsightsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_non_container_database_id"),
			),
		},

		// verify enabled
		{
			Config: config + compartmentIdVariableStr + ExternalNonContainerDatabaseOperationsInsightsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_non_container_database_operations_insights_management", "test_external_non_container_database_operations_insights_management", acctest.Required, acctest.Create, externalNonContainerDatabaseOperationsInsightsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceNonCDB, "operations_insights_config.0.operations_insights_status", "ENABLED"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ExternalNonContainerDatabaseOperationsInsightsManagementResourceDependencies,
		},
		// verify Update (Enable Operations Insights)
		{
			Config: config + compartmentIdVariableStr + ExternalNonContainerDatabaseOperationsInsightsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_non_container_database_operations_insights_management", "test_external_non_container_database_operations_insights_management", acctest.Optional, acctest.Create, externalNonContainerDatabaseOperationsInsightsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_non_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
			),
		},
		// verify Update (Disable Operations Insights)
		{
			Config: config + compartmentIdVariableStr + ExternalNonContainerDatabaseOperationsInsightsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_non_container_database_operations_insights_management", "test_external_non_container_database_operations_insights_management", acctest.Optional, acctest.Update, externalNonContainerDatabaseOperationsInsightsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_non_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
			),
		},
		// Verify Disablement
		{
			Config: config + compartmentIdVariableStr + ExternalNonContainerDatabaseOperationsInsightsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_non_container_database_operations_insights_management", "test_external_non_container_database_operations_insights_management", acctest.Optional, acctest.Update, externalNonContainerDatabaseOperationsInsightsManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceNonCDB, "operations_insights_config.0.operations_insights_status", "NOT_ENABLED"),
			),
		},
	})
}
