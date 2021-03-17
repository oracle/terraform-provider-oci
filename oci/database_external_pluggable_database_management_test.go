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
	externalPluggableDatabaseManagementRepresentation = map[string]interface{}{
		"external_pluggable_database_id": Representation{repType: Required, create: `${oci_database_external_pluggable_database.test_external_pluggable_database.id}`},
		"external_database_connector_id": Representation{repType: Required, create: `${oci_database_external_database_connector.test_external_pluggable_database_connector.id}`},
		"enable_management":              Representation{repType: Required, create: `true`, update: `false`},
	}
	externalPluggableDatabaseConnectorRepresentation = map[string]interface{}{
		"connection_credentials": RepresentationGroup{Required, externalDatabaseConnectorConnectionCredentialsRepresentation},
		"connection_string":      RepresentationGroup{Required, externalDatabaseConnectorConnectionStringRepresentation},
		"connector_agent_id":     Representation{repType: Required, create: `ocid1.managementagent.oc1.phx.amaaaaaajobtc3iaes4ijczgekzqigoji25xocsny7yundummydummydummy`},
		"display_name":           Representation{repType: Required, create: `myTestConn`},
		"external_database_id":   Representation{repType: Required, create: `${oci_database_external_pluggable_database.test_external_pluggable_database.id}`},
		"connector_type":         Representation{repType: Optional, create: `MACS`},
	}

	externalPluggable1DatabaseRepresentation = map[string]interface{}{
		"compartment_id":                 Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":                   Representation{repType: Required, create: `myTestExternalPdb`},
		"external_container_database_id": Representation{repType: Required, create: `${oci_database_external_container_database.test_external_container_database.id}`},
		"defined_tags":                   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	ExternalPluggableDatabaseManagementResourceDependencies = generateResourceFromRepresentationMap("oci_database_external_container_database", "test_external_container_database", Required, Create, externalContainerDatabaseRepresentation) +
		generateResourceFromRepresentationMap("oci_database_external_pluggable_database", "test_external_pluggable_database", Required, Create, externalPluggable1DatabaseRepresentation) +
		generateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_pluggable_database_connector", Required, Create, externalPluggableDatabaseConnectorRepresentation) +
		generateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", Required, Create, externalContainerDatabaseConnectorRepresentation)
)

func TestDatabaseExternalPluggableDatabaseManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExternalPluggableDatabaseManagementResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_external_pluggable_database_management.test_external_pluggable_database_management"
	resourcePDB := "oci_database_external_pluggable_database.test_external_pluggable_database"

	// Save TF content to create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	saveConfigContent(config+compartmentIdVariableStr+ExternalPluggableDatabaseManagementResourceDependencies+
		generateResourceFromRepresentationMap("oci_database_external_pluggable_database_management", "test_external_pluggable_database_management", Required, Create, externalPluggableDatabaseManagementRepresentation), "database", "externalPluggableDatabaseManagement", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// Enablement of parent CDB
			{
				Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseManagementResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", Required, Create, externalContainerDatabaseManagementRepresentation),
			},
			// Enablement of PDB
			{
				Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseManagementResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", Required, Create, externalContainerDatabaseManagementRepresentation) +
					generateResourceFromRepresentationMap("oci_database_external_pluggable_database_management", "test_external_pluggable_database_management", Required, Create, externalPluggableDatabaseManagementRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "external_pluggable_database_id"),
					resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
				),
			},
			// Verify Enablement of PDB
			{
				Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseManagementResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", Required, Create, externalContainerDatabaseManagementRepresentation) +
					generateResourceFromRepresentationMap("oci_database_external_pluggable_database_management", "test_external_pluggable_database_management", Required, Create, externalPluggableDatabaseManagementRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourcePDB, "database_management_config.0.database_management_status", "ENABLED"),
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseManagementResourceDependencies,
			},
			// Disablement of parent CDB
			{
				Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseManagementResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", Optional, Update, externalContainerDatabaseManagementRepresentation),
			},
			// Disablement of PDB
			{
				Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseManagementResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", Optional, Update, externalContainerDatabaseManagementRepresentation) +
					generateResourceFromRepresentationMap("oci_database_external_pluggable_database_management", "test_external_pluggable_database_management", Optional, Update, externalPluggableDatabaseManagementRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "external_pluggable_database_id"),
					resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
				),
			},
			// Verify Disablement of PDB
			{
				Config: config + compartmentIdVariableStr + ExternalPluggableDatabaseManagementResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", Optional, Update, externalContainerDatabaseManagementRepresentation) +
					generateResourceFromRepresentationMap("oci_database_external_pluggable_database_management", "test_external_pluggable_database_management", Optional, Update, externalPluggableDatabaseManagementRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourcePDB, "database_management_config.0.database_management_status", "NOT_ENABLED"),
				),
			},
		},
	})
}
