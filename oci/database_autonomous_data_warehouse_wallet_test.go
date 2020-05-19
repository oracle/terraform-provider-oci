// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	autonomousDataWarehouseWalletSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_data_warehouse_id": Representation{repType: Required, create: `${oci_database_autonomous_data_warehouse.test_autonomous_data_warehouse.id}`},
		"password":                     Representation{repType: Required, create: `BEstrO0ng_#11`},
		"base64_encode_content":        Representation{repType: Optional, create: `true`},
	}

	AutonomousDataWarehouseWalletResourceConfig = generateResourceFromRepresentationMap("oci_database_autonomous_data_warehouse", "test_autonomous_data_warehouse", Required, Create, autonomousDataWarehouseRepresentation)
)

func TestDatabaseAutonomousDataWarehouseWalletResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousDataWarehouseWalletResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_database_autonomous_data_warehouse_wallet.test_autonomous_data_warehouse_wallet"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_data_warehouse_wallet", "test_autonomous_data_warehouse_wallet", Required, Create, autonomousDataWarehouseWalletSingularDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousDataWarehouseWalletResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_data_warehouse_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(singularDatasourceName, "base64_encode_content", "false"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "content"),
				),
			},

			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_data_warehouse_wallet", "test_autonomous_data_warehouse_wallet", Optional, Create, autonomousDataWarehouseWalletSingularDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousDataWarehouseWalletResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_data_warehouse_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(singularDatasourceName, "base64_encode_content", "true"),
					testCheckAttributeBase64Encoded(singularDatasourceName, "content", true),
				),
			},
		},
	})
}
