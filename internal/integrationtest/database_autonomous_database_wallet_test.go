// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	AutonomousDatabaseWalletRequiredOnlyResource = AutonomousDatabaseWalletResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_wallet", "test_autonomous_database_wallet", acctest.Required, acctest.Create, autonomousDatabaseWalletRepresentation)

	AutonomousDatabaseWalletResourceConfig = AutonomousDatabaseWalletResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_wallet", "test_autonomous_database_wallet", acctest.Optional, acctest.Update, autonomousDatabaseWalletRepresentation)

	adbWalletDbName = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)

	autonomousDatabaseWalletSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
		"password":               acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"generate_type":          acctest.Representation{RepType: acctest.Optional, Create: `ALL`},
		"base64_encode_content":  acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}

	autonomousDatabaseWalletRepresentation = map[string]interface{}{
		"autonomous_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
		"password":               acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"base64_encode_content":  acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"generate_type":          acctest.Representation{RepType: acctest.Optional, Create: `ALL`},
	}

	AutonomousDatabaseWalletResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create,
		acctest.GetUpdatedRepresentationCopy("db_name", acctest.Representation{RepType: acctest.Required, Create: adbWalletDbName}, autonomousDatabaseRepresentation))
)

// issue-routing-tag: database/dbaas-adb
func TestDatabaseAutonomousDatabaseWalletResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousDatabaseWalletResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database_wallet.test_autonomous_database_wallet"

	singularDatasourceName := "data.oci_database_autonomous_database_wallet.test_autonomous_database_wallet"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+AutonomousDatabaseWalletResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_wallet", "test_autonomous_database_wallet", acctest.Optional, acctest.Create, autonomousDatabaseWalletRepresentation), "database", "autonomousDatabaseWallet", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseWalletResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_wallet", "test_autonomous_database_wallet", acctest.Required, acctest.Create, autonomousDatabaseWalletRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_wallet", "test_autonomous_database_wallet", acctest.Optional, acctest.Create, autonomousDatabaseWalletRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_wallet", "test_autonomous_database_wallet", acctest.Required, acctest.Create, autonomousDatabaseWalletSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousDatabaseWalletResourceDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "generate_type", "SINGLE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "base64_encode_content", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "content"),
			),
		},

		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_wallet", "test_autonomous_database_wallet", acctest.Optional, acctest.Create, autonomousDatabaseWalletSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousDatabaseWalletResourceDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "generate_type", "ALL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "base64_encode_content", "true"),
				acctest.TestCheckAttributeBase64Encoded(singularDatasourceName, "content", true),
			),
		},
	})
}
