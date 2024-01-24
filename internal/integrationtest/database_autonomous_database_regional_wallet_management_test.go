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
	DatabaseAutonomousDatabaseRegionalWalletManagementResourceConfig = DatabaseAutonomousDatabaseRegionalWalletManagementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_regional_wallet_management", "test_autonomous_database_regional_wallet_management", acctest.Optional, acctest.Update, DatabaseAutonomousDatabaseRegionalWalletManagementRepresentation)

	DatabaseDatabaseAutonomousDatabaseRegionalWalletManagementSingularDataSourceRepresentation = map[string]interface{}{}

	DatabaseAutonomousDatabaseRegionalWalletManagementRepresentation = map[string]interface{}{
		"grace_period":  acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"should_rotate": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	DatabaseAutonomousDatabaseRegionalWalletManagementResourceDependencies = ""
)

// issue-routing-tag: database/dbaas-adb
func TestDatabaseAutonomousDatabaseRegionalWalletManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousDatabaseRegionalWalletManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database_regional_wallet_management.test_autonomous_database_regional_wallet_management"

	singularDatasourceName := "data.oci_database_autonomous_database_regional_wallet_management.test_autonomous_database_regional_wallet_management"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseAutonomousDatabaseRegionalWalletManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_regional_wallet_management", "test_autonomous_database_regional_wallet_management", acctest.Optional, acctest.Create, DatabaseAutonomousDatabaseRegionalWalletManagementRepresentation), "database", "autonomousDatabaseRegionalWalletManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		//0. verify create
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseRegionalWalletManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_regional_wallet_management", "test_autonomous_database_regional_wallet_management", acctest.Required, acctest.Create, DatabaseAutonomousDatabaseRegionalWalletManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		//1. delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseRegionalWalletManagementResourceDependencies,
		},
		//2. verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseRegionalWalletManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_regional_wallet_management", "test_autonomous_database_regional_wallet_management", acctest.Optional, acctest.Create, DatabaseAutonomousDatabaseRegionalWalletManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
		//3. verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseRegionalWalletManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_regional_wallet_management", "test_autonomous_database_regional_wallet_management", acctest.Optional, acctest.Update, DatabaseAutonomousDatabaseRegionalWalletManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
		//4. verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_regional_wallet_management", "test_autonomous_database_regional_wallet_management", acctest.Required, acctest.Create, DatabaseDatabaseAutonomousDatabaseRegionalWalletManagementSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousDatabaseRegionalWalletManagementResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_rotated"),
			),
		},
	})
}
