// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	externalContainerDatabaseManagementRepresentation = map[string]interface{}{
		"external_container_database_id": Representation{RepType: Required, Create: `${oci_database_external_container_database.test_external_container_database.id}`},
		"external_database_connector_id": Representation{RepType: Required, Create: `${oci_database_external_database_connector.test_external_database_connector.id}`},
		"license_model":                  Representation{RepType: Required, Create: `BRING_YOUR_OWN_LICENSE`},
		"enable_management":              Representation{RepType: Required, Create: `true`, Update: `false`},
	}
	externalContainerDatabaseConnectorRepresentation = map[string]interface{}{
		"connection_credentials": RepresentationGroup{Required, externalDatabaseConnectorConnectionCredentialsRepresentation},
		"connection_string":      RepresentationGroup{Required, externalDatabaseConnectorConnectionStringRepresentation},
		"connector_agent_id":     Representation{RepType: Required, Create: `ocid1.managementagent.oc1.phx.amaaaaaajobtc3iaes4ijczgekzqigoji25xocsny7yundummydummydummy`},
		"display_name":           Representation{RepType: Required, Create: `myTestConn`},
		"external_database_id":   Representation{RepType: Required, Create: `${oci_database_external_container_database.test_external_container_database.id}`},
		"connector_type":         Representation{RepType: Optional, Create: `MACS`},
	}

	ExternalContainerDatabaseManagementResourceDependencies = GenerateResourceFromRepresentationMap("oci_database_external_container_database", "test_external_container_database", Required, Create, externalContainerDatabaseRepresentation) +
		GenerateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", Required, Create, externalContainerDatabaseConnectorRepresentation)
)

// issue-routing-tag: database/default
func TestDatabaseExternalContainerDatabaseManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExternalContainerDatabaseManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_external_container_database_management.test_external_container_database_management"

	resourceCdb := "oci_database_external_container_database.test_external_container_database"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+ExternalContainerDatabaseManagementResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", Required, Create, externalContainerDatabaseManagementRepresentation), "database", "externalContainerDatabaseManagement", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create (Enable Database Management)
		{
			Config: config + compartmentIdVariableStr + ExternalContainerDatabaseManagementResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", Required, Create, externalContainerDatabaseManagementRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
			),
		},
		// Verify Enablement
		{
			Config: config + compartmentIdVariableStr + ExternalContainerDatabaseManagementResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", Required, Create, externalContainerDatabaseManagementRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceCdb, "database_management_config.0.database_management_status", "ENABLED"),
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ExternalContainerDatabaseManagementResourceDependencies,
		},
		// verify Update (Enable Database Management)
		{
			Config: config + compartmentIdVariableStr + ExternalContainerDatabaseManagementResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", Optional, Create, externalContainerDatabaseManagementRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
			),
		},
		// verify Update (Disable Database Management)
		{
			Config: config + compartmentIdVariableStr + ExternalContainerDatabaseManagementResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", Optional, Update, externalContainerDatabaseManagementRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
			),
		},
		// Verify Disablement
		{
			Config: config + compartmentIdVariableStr + ExternalContainerDatabaseManagementResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", Optional, Update, externalContainerDatabaseManagementRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceCdb, "database_management_config.0.database_management_status", "NOT_ENABLED"),
			),
		},
	})
}
