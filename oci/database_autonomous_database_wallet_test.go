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
	AutonomousDatabaseWalletRequiredOnlyResource = AutonomousDatabaseWalletResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_autonomous_database_wallet", "test_autonomous_database_wallet", Required, Create, autonomousDatabaseWalletRepresentation)

	AutonomousDatabaseWalletResourceConfig = AutonomousDatabaseWalletResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_autonomous_database_wallet", "test_autonomous_database_wallet", Optional, Update, autonomousDatabaseWalletRepresentation)

	adbWalletDbName = randomString(1, charsetWithoutDigits) + randomString(13, charset)

	autonomousDatabaseWalletSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_database_id": Representation{repType: Required, create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
		"password":               Representation{repType: Required, create: `BEstrO0ng_#11`},
		"generate_type":          Representation{repType: Optional, create: `ALL`},
		"base64_encode_content":  Representation{repType: Optional, create: `true`},
	}

	autonomousDatabaseWalletRepresentation = map[string]interface{}{
		"autonomous_database_id": Representation{repType: Required, create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
		"password":               Representation{repType: Required, create: `BEstrO0ng_#11`},
		"base64_encode_content":  Representation{repType: Optional, create: `true`},
		"generate_type":          Representation{repType: Optional, create: `ALL`},
	}

	AutonomousDatabaseWalletResourceDependencies = generateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create,
		getUpdatedRepresentationCopy("db_name", Representation{repType: Required, create: adbWalletDbName}, autonomousDatabaseRepresentation))
)

// issue-routing-tag: database/dbaas-adb
func TestDatabaseAutonomousDatabaseWalletResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousDatabaseWalletResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database_wallet.test_autonomous_database_wallet"

	singularDatasourceName := "data.oci_database_autonomous_database_wallet.test_autonomous_database_wallet"

	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+AutonomousDatabaseWalletResourceDependencies+
		generateResourceFromRepresentationMap("oci_database_autonomous_database_wallet", "test_autonomous_database_wallet", Optional, Create, autonomousDatabaseWalletRepresentation), "database", "autonomousDatabaseWallet", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseWalletResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database_wallet", "test_autonomous_database_wallet", Required, Create, autonomousDatabaseWalletRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "autonomous_database_id"),
					resource.TestCheckResourceAttr(resourceName, "password", "BEstrO0ng_#11"),
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseWalletResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseWalletResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database_wallet", "test_autonomous_database_wallet", Optional, Create, autonomousDatabaseWalletRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "autonomous_database_id"),
					resource.TestCheckResourceAttr(resourceName, "base64_encode_content", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "content"),
					resource.TestCheckResourceAttr(resourceName, "generate_type", "ALL"),
					resource.TestCheckResourceAttr(resourceName, "password", "BEstrO0ng_#11"),
				),
			},

			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_database_wallet", "test_autonomous_database_wallet", Required, Create, autonomousDatabaseWalletSingularDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousDatabaseWalletResourceDependencies,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "generate_type", "SINGLE"),
					resource.TestCheckResourceAttr(singularDatasourceName, "password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttr(singularDatasourceName, "base64_encode_content", "false"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "content"),
				),
			},

			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_database_wallet", "test_autonomous_database_wallet", Optional, Create, autonomousDatabaseWalletSingularDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousDatabaseWalletResourceDependencies,
				Check: ComposeAggregateTestCheckFuncWrapper(
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
