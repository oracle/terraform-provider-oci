// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	externalPluggableDatabaseOperationsInsightsManagementRepresentation = map[string]interface{}{
		"external_database_connector_id": Representation{repType: Required, create: `${oci_database_external_database_connector.test_external_database_connector.id}`},
		"external_pluggable_database_id": Representation{repType: Required, create: `${oci_database_external_pluggable_database.test_external_pluggable_database.id}`},
		"enable_operations_insights":     Representation{repType: Required, create: `true`, update: `false`},
	}

	ExternalPluggableDatabaseOperationsInsightsManagementResourceDependencies = generateResourceFromRepresentationMap("oci_database_external_container_database", "test_external_container_database", Required, Create, externalContainerDatabaseRepresentation) +
		generateResourceFromRepresentationMap("oci_database_external_pluggable_database", "test_external_pluggable_database", Required, Create, externalPluggable1DatabaseRepresentation) +
		generateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_pluggable_database_connector", Required, Create, externalPluggableDatabaseConnectorRepresentation) +
		generateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", Required, Create, externalContainerDatabaseConnectorRepresentation)
)

// issue-routing-tag: database/default
func TestDatabaseExternalPluggableDatabaseOperationsInsightsManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExternalPluggableDatabaseOperationsInsightsManagementResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_external_pluggable_database_operations_insights_management.test_external_pluggable_database_operations_insights_management"
	resourcePDB := "oci_database_external_pluggable_database.test_external_pluggable_database"

	// Save TF content to create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	saveConfigContent(config+compartmentIdVariableStr+ExternalPluggableDatabaseOperationsInsightsManagementResourceDependencies+
		generateResourceFromRepresentationMap("oci_database_external_pluggable_database_operations_insights_management", "test_external_pluggable_database_operations_insights_management", Required, Create, externalPluggableDatabaseOperationsInsightsManagementRepresentation), "database", "externalPluggableDatabaseOperationsInsightsManagement", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// Enablement of PDB
			{
				Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseOperationsInsightsManagementResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_pluggable_database_operations_insights_management", "test_external_pluggable_database_operations_insights_management", Required, Create, externalPluggableDatabaseOperationsInsightsManagementRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
					resource.TestCheckResourceAttrSet(resourceName, "external_pluggable_database_id"),
				),
			},

			// Verify Enablement of PDB
			{
				Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseOperationsInsightsManagementResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_pluggable_database_operations_insights_management", "test_external_pluggable_database_operations_insights_management", Required, Create, externalPluggableDatabaseOperationsInsightsManagementRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourcePDB, "operations_insights_config.0.operations_insights_status", "ENABLED"),
				),
			},
			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseOperationsInsightsManagementResourceDependencies,
			},
			// Enablement of PDB
			{
				Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseOperationsInsightsManagementResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_pluggable_database_operations_insights_management", "test_external_pluggable_database_operations_insights_management", Required, Create, externalPluggableDatabaseOperationsInsightsManagementRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "external_pluggable_database_id"),
					resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
				),
			},
			// Verify Enablement of PDB
			{
				Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseOperationsInsightsManagementResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_pluggable_database_operations_insights_management", "test_external_pluggable_database_operations_insights_management", Required, Create, externalPluggableDatabaseOperationsInsightsManagementRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourcePDB, "operations_insights_config.0.operations_insights_status", "ENABLED"),
				),
			},

			// Disablement of PDB
			{
				Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseOperationsInsightsManagementResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_pluggable_database_operations_insights_management", "test_external_pluggable_database_operations_insights_management", Optional, Update, externalPluggableDatabaseOperationsInsightsManagementRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "external_pluggable_database_id"),
					resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
				),
			},
			// Verify Disablement of PDB
			{
				Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseOperationsInsightsManagementResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_pluggable_database_operations_insights_management", "test_external_pluggable_database_operations_insights_management", Optional, Update, externalPluggableDatabaseOperationsInsightsManagementRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourcePDB, "operations_insights_config.0.operations_insights_status", "NOT_ENABLED"),
				),
			},
		},
	})
}
