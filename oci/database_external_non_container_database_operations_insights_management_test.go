// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	externalNonContainerDatabaseOperationsInsightsManagementRepresentation = map[string]interface{}{
		"external_database_connector_id":     Representation{repType: Required, create: `${oci_database_external_database_connector.test_external_database_connector.id}`},
		"external_non_container_database_id": Representation{repType: Required, create: `${oci_database_external_non_container_database.test_external_non_container_database.id}`},
		"enable_operations_insights":         Representation{repType: Required, create: `true`, update: `false`},
	}

	ExternalNonContainerDatabaseOperationsInsightsManagementResourceDependencies = generateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", Required, Create, externalDatabaseConnectorRepresentation) +
		generateResourceFromRepresentationMap("oci_database_external_non_container_database", "test_external_non_container_database", Required, Create, externalNonContainerDatabaseRepresentation)
)

// issue-routing-tag: database/default
func TestDatabaseExternalNonContainerDatabaseOperationsInsightsManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExternalNonContainerDatabaseOperationsInsightsManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_external_non_container_database_operations_insights_management.test_external_non_container_database_operations_insights_management"
	resourceNonCDB := "oci_database_external_non_container_database.test_external_non_container_database"
	// Save TF content to create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	saveConfigContent(config+compartmentIdVariableStr+ExternalNonContainerDatabaseOperationsInsightsManagementResourceDependencies+
		generateResourceFromRepresentationMap("oci_database_external_non_container_database_operations_insights_management", "test_external_non_container_database_operations_insights_management", Required, Create, externalNonContainerDatabaseOperationsInsightsManagementRepresentation), "database", "externalNonContainerDatabaseOperationsInsightsManagement", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + ExternalNonContainerDatabaseOperationsInsightsManagementResourceDependencies +
				generateResourceFromRepresentationMap("oci_database_external_non_container_database_operations_insights_management", "test_external_non_container_database_operations_insights_management", Required, Create, externalNonContainerDatabaseOperationsInsightsManagementRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_non_container_database_id"),
			),
		},

		// verify enabled
		{
			Config: config + compartmentIdVariableStr + ExternalNonContainerDatabaseOperationsInsightsManagementResourceDependencies +
				generateResourceFromRepresentationMap("oci_database_external_non_container_database_operations_insights_management", "test_external_non_container_database_operations_insights_management", Required, Create, externalNonContainerDatabaseOperationsInsightsManagementRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceNonCDB, "operations_insights_config.0.operations_insights_status", "ENABLED"),
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr + ExternalNonContainerDatabaseOperationsInsightsManagementResourceDependencies,
		},
		// verify update (Enable Operations Insights)
		{
			Config: config + compartmentIdVariableStr + ExternalNonContainerDatabaseOperationsInsightsManagementResourceDependencies +
				generateResourceFromRepresentationMap("oci_database_external_non_container_database_operations_insights_management", "test_external_non_container_database_operations_insights_management", Optional, Create, externalNonContainerDatabaseOperationsInsightsManagementRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_non_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
			),
		},
		// verify update (Disable Operations Insights)
		{
			Config: config + compartmentIdVariableStr + ExternalNonContainerDatabaseOperationsInsightsManagementResourceDependencies +
				generateResourceFromRepresentationMap("oci_database_external_non_container_database_operations_insights_management", "test_external_non_container_database_operations_insights_management", Optional, Update, externalNonContainerDatabaseOperationsInsightsManagementRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_non_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
			),
		},
		// Verify Disablement
		{
			Config: config + compartmentIdVariableStr + ExternalNonContainerDatabaseOperationsInsightsManagementResourceDependencies +
				generateResourceFromRepresentationMap("oci_database_external_non_container_database_operations_insights_management", "test_external_non_container_database_operations_insights_management", Optional, Update, externalNonContainerDatabaseOperationsInsightsManagementRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceNonCDB, "operations_insights_config.0.operations_insights_status", "NOT_ENABLED"),
			),
		},
	})
}
