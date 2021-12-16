// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
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
	externalPluggableDatabaseManagementRepresentation = map[string]interface{}{
		"external_pluggable_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_external_pluggable_database.test_external_pluggable_database.id}`},
		"external_database_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_external_database_connector.test_external_pluggable_database_connector.id}`},
		"enable_management":              acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}
	externalPluggableDatabaseConnectorRepresentation = map[string]interface{}{
		"connection_credentials": acctest.RepresentationGroup{RepType: acctest.Required, Group: externalDatabaseConnectorConnectionCredentialsRepresentation},
		"connection_string":      acctest.RepresentationGroup{RepType: acctest.Required, Group: externalDatabaseConnectorConnectionStringRepresentation},
		"connector_agent_id":     acctest.Representation{RepType: acctest.Required, Create: `ocid1.managementagent.oc1.phx.amaaaaaajobtc3iaes4ijczgekzqigoji25xocsny7yundummydummydummy`},
		"display_name":           acctest.Representation{RepType: acctest.Required, Create: `myTestConn`},
		"external_database_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_database_external_pluggable_database.test_external_pluggable_database.id}`},
		"connector_type":         acctest.Representation{RepType: acctest.Optional, Create: `MACS`},
	}

	externalPluggable1DatabaseRepresentation = map[string]interface{}{
		"compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                   acctest.Representation{RepType: acctest.Required, Create: `myTestExternalPdb`},
		"external_container_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_external_container_database.test_external_container_database.id}`},
		"defined_tags":                   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	ExternalPluggableDatabaseManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database", "test_external_container_database", acctest.Required, acctest.Create, externalContainerDatabaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database", "test_external_pluggable_database", acctest.Required, acctest.Create, externalPluggable1DatabaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_pluggable_database_connector", acctest.Required, acctest.Create, externalPluggableDatabaseConnectorRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", acctest.Required, acctest.Create, externalContainerDatabaseConnectorRepresentation)
)

// issue-routing-tag: database/default
func TestDatabaseExternalPluggableDatabaseManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExternalPluggableDatabaseManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_external_pluggable_database_management.test_external_pluggable_database_management"
	resourcePDB := "oci_database_external_pluggable_database.test_external_pluggable_database"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ExternalPluggableDatabaseManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database_management", "test_external_pluggable_database_management", acctest.Required, acctest.Create, externalPluggableDatabaseManagementRepresentation), "database", "externalPluggableDatabaseManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Enablement of parent CDB
		{
			Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", acctest.Required, acctest.Create, externalContainerDatabaseManagementRepresentation),
		},
		// Enablement of PDB
		{
			Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", acctest.Required, acctest.Create, externalContainerDatabaseManagementRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database_management", "test_external_pluggable_database_management", acctest.Required, acctest.Create, externalPluggableDatabaseManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_pluggable_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
			),
		},
		// Verify Enablement of PDB
		{
			Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", acctest.Required, acctest.Create, externalContainerDatabaseManagementRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database_management", "test_external_pluggable_database_management", acctest.Required, acctest.Create, externalPluggableDatabaseManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourcePDB, "database_management_config.0.database_management_status", "ENABLED"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseManagementResourceDependencies,
		},
		// Enablement of parent CDB
		{
			Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", acctest.Optional, acctest.Create, externalContainerDatabaseManagementRepresentation),
		},
		// Enablement of PDB
		{
			Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", acctest.Optional, acctest.Create, externalContainerDatabaseManagementRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database_management", "test_external_pluggable_database_management", acctest.Optional, acctest.Create, externalPluggableDatabaseManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_pluggable_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
			),
		},
		// Verify Enablement of PDB
		{
			Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", acctest.Optional, acctest.Create, externalContainerDatabaseManagementRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database_management", "test_external_pluggable_database_management", acctest.Optional, acctest.Create, externalPluggableDatabaseManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourcePDB, "database_management_config.0.database_management_status", "ENABLED"),
			),
		},
		// Disablement of parent CDB
		{
			Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", acctest.Optional, acctest.Update, externalContainerDatabaseManagementRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database_management", "test_external_pluggable_database_management", acctest.Optional, acctest.Update, externalPluggableDatabaseManagementRepresentation),
		},
		// Disablement of PDB
		{
			Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", acctest.Optional, acctest.Update, externalContainerDatabaseManagementRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database_management", "test_external_pluggable_database_management", acctest.Optional, acctest.Update, externalPluggableDatabaseManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_pluggable_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
			),
		},
		// Verify Disablement of PDB
		{
			Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", acctest.Optional, acctest.Update, externalContainerDatabaseManagementRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_pluggable_database_management", "test_external_pluggable_database_management", acctest.Optional, acctest.Update, externalPluggableDatabaseManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourcePDB, "database_management_config.0.database_management_status", "NOT_ENABLED"),
			),
		},
	})
}
