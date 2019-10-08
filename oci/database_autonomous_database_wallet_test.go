// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	autonomousDatabaseWalletSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_database_id": Representation{repType: Required, create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
		"password":               Representation{repType: Required, create: `BEstrO0ng_#11`},
		"generate_type":          Representation{repType: Optional, create: `ALL`},
		"base64_encode_content":  Representation{repType: Optional, create: `true`},
	}

	AutonomousDatabaseWalletResourceConfig = generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create, autonomousDatabaseRepresentation)
)

func TestDatabaseAutonomousDatabaseWalletResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousDatabaseWalletResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_database_autonomous_database_wallet.test_autonomous_database_wallet"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_database_wallet", "test_autonomous_database_wallet", Required, Create, autonomousDatabaseWalletSingularDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousDatabaseWalletResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "generate_type", "SINGLE"),
					resource.TestCheckResourceAttr(singularDatasourceName, "password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(singularDatasourceName, "base64_encode_content", "false"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "content"),
				),
			},
			// verify datasource with base64 encoded content
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_database_wallet", "test_autonomous_database_wallet", Optional, Create, autonomousDatabaseWalletSingularDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousDatabaseWalletResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "generate_type", "ALL"),
					resource.TestCheckResourceAttr(singularDatasourceName, "password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(singularDatasourceName, "base64_encode_content", "true"),
					testCheckAttributeBase64Encoded(singularDatasourceName, "content", true),
				),
			},
		},
	})
}
