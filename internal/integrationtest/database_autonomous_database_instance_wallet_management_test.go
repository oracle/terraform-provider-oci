// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	adbWalletMgmtDbName = utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)

	DatabaseAutonomousDatabaseInstanceWalletManagementResourceConfig = DatabaseAutonomousDatabaseInstanceWalletManagementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_instance_wallet_management", "test_autonomous_database_instance_wallet_management", acctest.Optional, acctest.Update, DatabaseAutonomousDatabaseInstanceWalletManagementRepresentation)

	DatabaseDatabaseAutonomousDatabaseInstanceWalletManagementSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
	}

	DatabaseAutonomousDatabaseInstanceWalletManagementRepresentation = map[string]interface{}{
		"autonomous_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
		"grace_period":           acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"should_rotate":          acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	DatabaseAutonomousDatabaseInstanceWalletManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create,
		acctest.GetUpdatedRepresentationCopy("db_name", acctest.Representation{RepType: acctest.Required, Create: adbWalletMgmtDbName}, DatabaseAutonomousDatabaseRepresentation))
)

// issue-routing-tag: database/dbaas-adb
func TestDatabaseAutonomousDatabaseInstanceWalletManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousDatabaseInstanceWalletManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database_instance_wallet_management.test_autonomous_database_instance_wallet_management"

	singularDatasourceName := "data.oci_database_autonomous_database_instance_wallet_management.test_autonomous_database_instance_wallet_management"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseAutonomousDatabaseInstanceWalletManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_instance_wallet_management", "test_autonomous_database_instance_wallet_management", acctest.Optional, acctest.Create, DatabaseAutonomousDatabaseInstanceWalletManagementRepresentation), "database", "autonomousDatabaseInstanceWalletManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		//0. verify create
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseInstanceWalletManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_instance_wallet_management", "test_autonomous_database_instance_wallet_management", acctest.Required, acctest.Create, DatabaseAutonomousDatabaseInstanceWalletManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_database_id"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseInstanceWalletManagementResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseInstanceWalletManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_instance_wallet_management", "test_autonomous_database_instance_wallet_management", acctest.Optional, acctest.Create, DatabaseAutonomousDatabaseInstanceWalletManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_database_id"),
				resource.TestCheckResourceAttr(resourceName, "grace_period", "10"),
				resource.TestCheckResourceAttr(resourceName, "should_rotate", "false"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		//1. verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseInstanceWalletManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_instance_wallet_management", "test_autonomous_database_instance_wallet_management", acctest.Optional, acctest.Update, DatabaseAutonomousDatabaseInstanceWalletManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_database_id"),
				resource.TestCheckResourceAttr(resourceName, "grace_period", "11"),
				resource.TestCheckResourceAttr(resourceName, "should_rotate", "true"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_rotated"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		//2. verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_instance_wallet_management", "test_autonomous_database_instance_wallet_management", acctest.Required, acctest.Create, DatabaseDatabaseAutonomousDatabaseInstanceWalletManagementSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousDatabaseInstanceWalletManagementResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_rotated"),
			),
		},
	})
}
