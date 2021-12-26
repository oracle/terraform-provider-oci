// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	dbManagementPrivateEndpointAssociatedDatabaseSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_management_private_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_db_management_private_endpoint.test_db_management_private_endpoint.id}`},
	}

	dbManagementPrivateEndpointAssociatedDatabaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_management_private_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_db_management_private_endpoint.test_db_management_private_endpoint.id}`},
	}

	DbManagementPrivateEndpointAssociatedDatabaseResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_db_management_private_endpoint", "test_db_management_private_endpoint", acctest.Required, acctest.Create, dbManagementPrivateEndpointRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementDbManagementPrivateEndpointAssociatedDatabaseResource_basic(t *testing.T) {
	t.Skip("Skip this test till Database Management service provides a better way of testing this. It requires a live managed database instance")
	httpreplay.SetScenario("TestDatabaseManagementDbManagementPrivateEndpointAssociatedDatabaseResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_management_db_management_private_endpoint_associated_databases.test_db_management_private_endpoint_associated_databases"
	singularDatasourceName := "data.oci_database_management_db_management_private_endpoint_associated_database.test_db_management_private_endpoint_associated_database"

	acctest.SaveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_db_management_private_endpoint_associated_databases", "test_db_management_private_endpoint_associated_databases", acctest.Required, acctest.Create, dbManagementPrivateEndpointAssociatedDatabaseDataSourceRepresentation) +
					compartmentIdVariableStr + DbManagementPrivateEndpointAssociatedDatabaseResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "db_management_private_endpoint_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "associated_database_collection.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "associated_database_collection.0.items.#"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_db_management_private_endpoint_associated_database", "test_db_management_private_endpoint_associated_database", acctest.Required, acctest.Create, dbManagementPrivateEndpointAssociatedDatabaseSingularDataSourceRepresentation) +
					compartmentIdVariableStr + DbManagementPrivateEndpointAssociatedDatabaseResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "db_management_private_endpoint_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "items.#"),
				),
			},
		},
	})
}
