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
	dbManagementPrivateEndpointAssociatedDatabaseSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                    Representation{repType: Required, create: `${var.compartment_id}`},
		"db_management_private_endpoint_id": Representation{repType: Required, create: `${oci_database_management_db_management_private_endpoint.test_db_management_private_endpoint.id}`},
	}

	dbManagementPrivateEndpointAssociatedDatabaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                    Representation{repType: Required, create: `${var.compartment_id}`},
		"db_management_private_endpoint_id": Representation{repType: Required, create: `${oci_database_management_db_management_private_endpoint.test_db_management_private_endpoint.id}`},
	}

	DbManagementPrivateEndpointAssociatedDatabaseResourceConfig = generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		generateResourceFromRepresentationMap("oci_database_management_db_management_private_endpoint", "test_db_management_private_endpoint", Required, Create, dbManagementPrivateEndpointRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementDbManagementPrivateEndpointAssociatedDatabaseResource_basic(t *testing.T) {
	t.Skip("Skip this test till Database Management service provides a better way of testing this. It requires a live managed database instance")
	httpreplay.SetScenario("TestDatabaseManagementDbManagementPrivateEndpointAssociatedDatabaseResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_management_db_management_private_endpoint_associated_databases.test_db_management_private_endpoint_associated_databases"
	singularDatasourceName := "data.oci_database_management_db_management_private_endpoint_associated_database.test_db_management_private_endpoint_associated_database"

	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_management_db_management_private_endpoint_associated_databases", "test_db_management_private_endpoint_associated_databases", Required, Create, dbManagementPrivateEndpointAssociatedDatabaseDataSourceRepresentation) +
					compartmentIdVariableStr + DbManagementPrivateEndpointAssociatedDatabaseResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "db_management_private_endpoint_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "associated_database_collection.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "associated_database_collection.0.items.#"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_management_db_management_private_endpoint_associated_database", "test_db_management_private_endpoint_associated_database", Required, Create, dbManagementPrivateEndpointAssociatedDatabaseSingularDataSourceRepresentation) +
					compartmentIdVariableStr + DbManagementPrivateEndpointAssociatedDatabaseResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "db_management_private_endpoint_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "items.#"),
				),
			},
		},
	})
}
