// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	AutonomousDatabaseRegionalWalletManagementResourceConfig = AutonomousDatabaseRegionalWalletManagementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_regional_wallet_management", "test_autonomous_database_regional_wallet_management", acctest.Optional, acctest.Update, autonomousDatabaseRegionalWalletManagementRepresentation)

	autonomousDatabaseRegionalWalletManagementSingularDataSourceRepresentation = map[string]interface{}{}

	autonomousDatabaseRegionalWalletManagementRepresentation = map[string]interface{}{
		"should_rotate": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	AutonomousDatabaseRegionalWalletManagementResourceDependencies = ""
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
	acctest.SaveConfigContent(config+compartmentIdVariableStr+AutonomousDatabaseRegionalWalletManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_regional_wallet_management", "test_autonomous_database_regional_wallet_management", acctest.Optional, acctest.Create, autonomousDatabaseRegionalWalletManagementRepresentation), "database", "autonomousDatabaseRegionalWalletManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseRegionalWalletManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_regional_wallet_management", "test_autonomous_database_regional_wallet_management", acctest.Required, acctest.Create, autonomousDatabaseRegionalWalletManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + AutonomousDatabaseRegionalWalletManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_regional_wallet_management", "test_autonomous_database_regional_wallet_management", acctest.Optional, acctest.Update, autonomousDatabaseRegionalWalletManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_database_regional_wallet_management", "test_autonomous_database_regional_wallet_management", acctest.Required, acctest.Create, autonomousDatabaseRegionalWalletManagementSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousDatabaseRegionalWalletManagementResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_rotated"),
			),
		},
	})
}
