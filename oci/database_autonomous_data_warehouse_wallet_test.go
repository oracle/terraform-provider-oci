// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var (
	AutonomousDataWarehouseWalletResourceConfig = AutonomousDataWarehouseWalletResourceDependencies + `

`
	AutonomousDataWarehouseWalletPropertyVariables = `
variable "autonomous_data_warehouse_wallet_password" { default = "BEstrO0ng_#11" }

`
	AutonomousDataWarehouseWalletResourceDependencies = AutonomousDataWarehouseResourceConfig
)

func TestDatabaseAutonomousDataWarehouseWalletResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_database_autonomous_data_warehouse_wallet.test_autonomous_data_warehouse_wallet"

	testResourceName := GenerateTestResourceName("adwdb1", 14)
	setEnvSetting("TF_VAR_autonomous_data_warehouse_db_name", testResourceName)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config + `
variable "autonomous_data_warehouse_wallet_password" { default = "BEstrO0ng_#11" }

data "oci_database_autonomous_data_warehouse_wallet" "test_autonomous_data_warehouse_wallet" {
	#Required
	autonomous_data_warehouse_id = "${oci_database_autonomous_data_warehouse.test_autonomous_data_warehouse.id}"
	password = "${var.autonomous_data_warehouse_wallet_password}"
}
                ` + compartmentIdVariableStr + AutonomousDataWarehouseWalletResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_data_warehouse_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "content"),
				),
			},
		},
	})
}
