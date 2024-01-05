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
	DatabaseExternalPluggableDatabaseManagementRepresentation = map[string]interface{}{
		"external_pluggable_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_external_pluggable_database.test_external_pluggable_database.id}`},
		"external_database_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_external_database_connector.test_external_pluggable_database_connector.id}`},
		"enable_management":              acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}
	DatabaseExternalPluggableDatabaseConnectorRepresentation = map[string]interface{}{
		"connection_credentials": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseExternalDatabaseConnectorConnectionCredentialsRepresentation},
		"connection_string":      acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseExternalDatabaseConnectorConnectionStringRepresentation},
		"connector_agent_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.agent_id}`},
		"display_name":           acctest.Representation{RepType: acctest.Required, Create: `myTestConn`},
		"external_database_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_database_external_pluggable_database.test_external_pluggable_database.id}`},
		"connector_type":         acctest.Representation{RepType: acctest.Optional, Create: `MACS`},
	}

	DatabaseExternalPluggable1DatabaseRepresentation = map[string]interface{}{
		"compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                   acctest.Representation{RepType: acctest.Required, Create: `myTestExternalPdb`},
		"external_container_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_external_container_database.test_external_container_database.id}`},
		"defined_tags":                   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	DatabaseExternalPluggableDatabaseManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database", "test_external_container_database", acctest.Required, acctest.Create, DatabaseExternalContainerDatabaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database", "test_external_pluggable_database", acctest.Required, acctest.Create, DatabaseExternalPluggable1DatabaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_pluggable_database_connector", acctest.Required, acctest.Create, DatabaseExternalPluggableDatabaseConnectorRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", acctest.Required, acctest.Create, DatabaseExternalContainerDatabaseConnectorRepresentation)
)

// issue-routing-tag: database/default
func TestDatabaseExternalPluggableDatabaseManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExternalPluggableDatabaseManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	agentId := utils.GetEnvSettingWithBlankDefault("connector_agent_id")
	agentIdVariableStr := fmt.Sprintf("variable \"agent_id\" { default = \"%s\" }\n", agentId)

	resourceName := "oci_database_external_pluggable_database_management.test_external_pluggable_database_management"
	resourcePDB := "oci_database_external_pluggable_database.test_external_pluggable_database"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+agentIdVariableStr+DatabaseExternalPluggableDatabaseManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database_management", "test_external_pluggable_database_management", acctest.Required, acctest.Create, DatabaseExternalPluggableDatabaseManagementRepresentation), "database", "externalPluggableDatabaseManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Enablement of parent CDB
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + DatabaseExternalPluggableDatabaseManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", acctest.Required, acctest.Create, DatabaseExternalContainerDatabaseManagementRepresentation),
		},
		// Enablement of PDB
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + DatabaseExternalPluggableDatabaseManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", acctest.Required, acctest.Create, DatabaseExternalContainerDatabaseManagementRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database_management", "test_external_pluggable_database_management", acctest.Required, acctest.Create, DatabaseExternalPluggableDatabaseManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_pluggable_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
			),
		},
		// Verify Enablement of PDB
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + DatabaseExternalPluggableDatabaseManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", acctest.Required, acctest.Create, DatabaseExternalContainerDatabaseManagementRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database_management", "test_external_pluggable_database_management", acctest.Required, acctest.Create, DatabaseExternalPluggableDatabaseManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourcePDB, "database_management_config.0.database_management_status", "ENABLED"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + DatabaseExternalPluggableDatabaseManagementResourceDependencies,
		},
		// Enablement of parent CDB
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + DatabaseExternalPluggableDatabaseManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", acctest.Optional, acctest.Create, DatabaseExternalContainerDatabaseManagementRepresentation),
		},
		// Enablement of PDB
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + DatabaseExternalPluggableDatabaseManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", acctest.Optional, acctest.Create, DatabaseExternalContainerDatabaseManagementRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database_management", "test_external_pluggable_database_management", acctest.Optional, acctest.Create, DatabaseExternalPluggableDatabaseManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_pluggable_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
			),
		},
		// Verify Enablement of PDB
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + DatabaseExternalPluggableDatabaseManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", acctest.Optional, acctest.Create, DatabaseExternalContainerDatabaseManagementRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database_management", "test_external_pluggable_database_management", acctest.Optional, acctest.Create, DatabaseExternalPluggableDatabaseManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourcePDB, "database_management_config.0.database_management_status", "ENABLED"),
			),
		},
		// Disablement of parent CDB
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + DatabaseExternalPluggableDatabaseManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", acctest.Optional, acctest.Update, DatabaseExternalContainerDatabaseManagementRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database_management", "test_external_pluggable_database_management", acctest.Optional, acctest.Update, DatabaseExternalPluggableDatabaseManagementRepresentation),
		},
		// Disablement of PDB
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + DatabaseExternalPluggableDatabaseManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", acctest.Optional, acctest.Update, DatabaseExternalContainerDatabaseManagementRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database_management", "test_external_pluggable_database_management", acctest.Optional, acctest.Update, DatabaseExternalPluggableDatabaseManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_pluggable_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
			),
		},
		// Verify Disablement of PDB
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + DatabaseExternalPluggableDatabaseManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", acctest.Optional, acctest.Update, DatabaseExternalContainerDatabaseManagementRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database_management", "test_external_pluggable_database_management", acctest.Optional, acctest.Update, DatabaseExternalPluggableDatabaseManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourcePDB, "database_management_config.0.database_management_status", "NOT_ENABLED"),
			),
		},
	})
}
