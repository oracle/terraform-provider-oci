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
	DatabaseExternalNonContainerDatabaseManagementRepresentation = map[string]interface{}{
		"external_database_connector_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_database_external_database_connector.test_external_database_connector.id}`},
		"external_non_container_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_external_non_container_database.test_external_non_container_database.id}`},
		"license_model":                      acctest.Representation{RepType: acctest.Required, Create: `BRING_YOUR_OWN_LICENSE`},
		"enable_management":                  acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}

	DatabaseExternalNonContainerDatabaseManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_external_non_container_database", "test_external_non_container_database", acctest.Required, acctest.Create, DatabaseExternalNonContainerDatabaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", acctest.Required, acctest.Create, DatabaseExternalDatabaseConnectorRepresentation)
)

// issue-routing-tag: database/default
func TestDatabaseExternalNonContainerDatabaseManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExternalNonContainerDatabaseManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_external_non_container_database_management.test_external_non_container_database_management"
	resourceNonCDB := "oci_database_external_non_container_database.test_external_non_container_database"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseExternalNonContainerDatabaseManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_non_container_database_management", "test_external_non_container_database_management", acctest.Required, acctest.Create, DatabaseExternalNonContainerDatabaseManagementRepresentation), "database", "externalNonContainerDatabaseManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create (Enable Database Management)
		{
			Config: config + compartmentIdVariableStr + DatabaseExternalNonContainerDatabaseManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_non_container_database_management", "test_external_non_container_database_management", acctest.Required, acctest.Create, DatabaseExternalNonContainerDatabaseManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_non_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
			),
		},
		// Verify Enablement
		{
			Config: config + compartmentIdVariableStr + DatabaseExternalNonContainerDatabaseManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_non_container_database_management", "test_external_non_container_database_management", acctest.Required, acctest.Create, DatabaseExternalNonContainerDatabaseManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceNonCDB, "database_management_config.0.database_management_status", "ENABLED"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseExternalNonContainerDatabaseManagementResourceDependencies,
		},
		// verify Update (Enable Database Management)
		{
			Config: config + compartmentIdVariableStr + DatabaseExternalNonContainerDatabaseManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_non_container_database_management", "test_external_non_container_database_management", acctest.Optional, acctest.Create, DatabaseExternalNonContainerDatabaseManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_non_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
			),
		},
		// verify Update (Disable Database Management)
		{
			Config: config + compartmentIdVariableStr + DatabaseExternalNonContainerDatabaseManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_non_container_database_management", "test_external_non_container_database_management", acctest.Optional, acctest.Update, DatabaseExternalNonContainerDatabaseManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_non_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
			),
		},
		// Verify Disablement
		{
			Config: config + compartmentIdVariableStr + DatabaseExternalNonContainerDatabaseManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_non_container_database_management", "test_external_non_container_database_management", acctest.Optional, acctest.Update, DatabaseExternalNonContainerDatabaseManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceNonCDB, "database_management_config.0.database_management_status", "NOT_ENABLED"),
			),
		},
	})
}
