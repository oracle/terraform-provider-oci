// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	AutonomousDatabaseWalletRequiredOnlyResource = AutonomousDatabaseWalletResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_autonomous_database_wallet", "test_autonomous_database_wallet", Required, Create, autonomousDatabaseWalletRepresentation)

	AutonomousDatabaseWalletResourceConfig = AutonomousDatabaseWalletResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_autonomous_database_wallet", "test_autonomous_database_wallet", Optional, Update, autonomousDatabaseWalletRepresentation)

	adbWalletDbName = RandomString(1, charsetWithoutDigits) + RandomString(13, charset)

	autonomousDatabaseWalletSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_database_id": Representation{RepType: Required, Create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
		"password":               Representation{RepType: Required, Create: `BEstrO0ng_#11`},
		"generate_type":          Representation{RepType: Optional, Create: `ALL`},
		"base64_encode_content":  Representation{RepType: Optional, Create: `true`},
	}

	autonomousDatabaseWalletRepresentation = map[string]interface{}{
		"autonomous_database_id": Representation{RepType: Required, Create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
		"password":               Representation{RepType: Required, Create: `BEstrO0ng_#11`},
		"base64_encode_content":  Representation{RepType: Optional, Create: `true`},
		"generate_type":          Representation{RepType: Optional, Create: `ALL`},
	}

	AutonomousDatabaseWalletResourceDependencies = GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create,
		GetUpdatedRepresentationCopy("db_name", Representation{RepType: Required, Create: adbWalletDbName}, autonomousDatabaseRepresentation))
)

// issue-routing-tag: database/dbaas-adb
func TestDatabaseAutonomousDatabaseWalletResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousDatabaseWalletResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database_wallet.test_autonomous_database_wallet"

	singularDatasourceName := "data.oci_database_autonomous_database_wallet.test_autonomous_database_wallet"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+AutonomousDatabaseWalletResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_database_autonomous_database_wallet", "test_autonomous_database_wallet", Optional, Create, autonomousDatabaseWalletRepresentation), "database", "autonomousDatabaseWallet", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseWalletResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database_wallet", "test_autonomous_database_wallet", Required, Create, autonomousDatabaseWalletRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_database_id"),
				resource.TestCheckResourceAttr(resourceName, "password", "BEstrO0ng_#11"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseWalletResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseWalletResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_autonomous_database_wallet", "test_autonomous_database_wallet", Optional, Create, autonomousDatabaseWalletRepresentation),
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
				GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_wallet", "test_autonomous_database_wallet", Required, Create, autonomousDatabaseWalletSingularDataSourceRepresentation) +
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
				GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_wallet", "test_autonomous_database_wallet", Optional, Create, autonomousDatabaseWalletSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousDatabaseWalletResourceDependencies,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "generate_type", "ALL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "base64_encode_content", "true"),
				TestCheckAttributeBase64Encoded(singularDatasourceName, "content", true),
			),
		},
	})
}
