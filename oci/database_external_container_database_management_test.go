// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
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
	externalContainerDatabaseManagementRepresentation = map[string]interface{}{
		"external_container_database_id": Representation{repType: Required, create: `${oci_database_external_container_database.test_external_container_database.id}`},
		"external_database_connector_id": Representation{repType: Required, create: `${oci_database_external_database_connector.test_external_database_connector.id}`},
		"license_model":                  Representation{repType: Required, create: `BRING_YOUR_OWN_LICENSE`},
		"enable_management":              Representation{repType: Required, create: `true`, update: `false`},
	}
	externalContainerDatabaseConnectorRepresentation = map[string]interface{}{
		"connection_credentials": RepresentationGroup{Required, externalDatabaseConnectorConnectionCredentialsRepresentation},
		"connection_string":      RepresentationGroup{Required, externalDatabaseConnectorConnectionStringRepresentation},
		"connector_agent_id":     Representation{repType: Required, create: `ocid1.managementagent.oc1.phx.amaaaaaajobtc3iaes4ijczgekzqigoji25xocsny7yundummydummydummy`},
		"display_name":           Representation{repType: Required, create: `myTestConn`},
		"external_database_id":   Representation{repType: Required, create: `${oci_database_external_container_database.test_external_container_database.id}`},
		"connector_type":         Representation{repType: Optional, create: `MACS`},
	}

	ExternalContainerDatabaseManagementResourceDependencies = generateResourceFromRepresentationMap("oci_database_external_container_database", "test_external_container_database", Required, Create, externalContainerDatabaseRepresentation) +
		generateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", Required, Create, externalContainerDatabaseConnectorRepresentation)
)

func TestDatabaseExternalContainerDatabaseManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExternalContainerDatabaseManagementResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_external_container_database_management.test_external_container_database_management"

	resourceCdb := "oci_database_external_container_database.test_external_container_database"

	// Save TF content to create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	saveConfigContent(config+compartmentIdVariableStr+ExternalContainerDatabaseManagementResourceDependencies+
		generateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", Required, Create, externalContainerDatabaseManagementRepresentation), "database", "externalContainerDatabaseManagement", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create (Enable Database Management)
			{
				Config: config + compartmentIdVariableStr + ExternalContainerDatabaseManagementResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", Required, Create, externalContainerDatabaseManagementRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "external_container_database_id"),
					resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
				),
			},
			// Verify Enablement
			{
				Config: config + compartmentIdVariableStr + ExternalContainerDatabaseManagementResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", Required, Create, externalContainerDatabaseManagementRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceCdb, "database_management_config.0.database_management_status", "ENABLED"),
				),
			},
			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ExternalContainerDatabaseManagementResourceDependencies,
			},
			// verify update (Disable Database Management)
			{
				Config: config + compartmentIdVariableStr + ExternalContainerDatabaseManagementResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", Optional, Update, externalContainerDatabaseManagementRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "external_container_database_id"),
					resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
				),
			},
			// Verify Disablement
			{
				Config: config + compartmentIdVariableStr + ExternalContainerDatabaseManagementResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", Optional, Update, externalContainerDatabaseManagementRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceCdb, "database_management_config.0.database_management_status", "NOT_ENABLED"),
				),
			},
		},
	})
}
