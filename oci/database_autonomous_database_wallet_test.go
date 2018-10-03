// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	AutonomousDatabaseWalletResourceConfig = AutonomousDatabaseWalletResourceDependencies + `

`
	AutonomousDatabaseWalletPropertyVariables = `
variable "autonomous_database_wallet_password" { default = "BEstrO0ng_#11" }

`
	AutonomousDatabaseWalletResourceDependencies = AutonomousDatabasePropertyVariables + AutonomousDatabaseResourceConfig
)

func TestDatabaseAutonomousDatabaseWalletResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_database_autonomous_database_wallet.test_autonomous_database_wallet"

	testResourceName := GenerateTestResourceName("adb1", 14)
	setEnvSetting("TF_VAR_autonomous_database_db_name", testResourceName)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config + `
variable "autonomous_database_wallet_password" { default = "BEstrO0ng_#11" }

data "oci_database_autonomous_database_wallet" "test_autonomous_database_wallet" {
	#Required
	autonomous_database_id = "${oci_database_autonomous_database.test_autonomous_database.id}"
	password = "${var.autonomous_database_wallet_password}"
}
                ` + compartmentIdVariableStr + AutonomousDatabaseWalletResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "content"),
				),
			},
		},
	})
}
