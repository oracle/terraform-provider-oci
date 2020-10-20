// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	AutonomousDatabaseRegionalWalletManagementResourceConfig = AutonomousDatabaseRegionalWalletManagementResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_autonomous_database_regional_wallet_management", "test_autonomous_database_regional_wallet_management", Optional, Update, autonomousDatabaseRegionalWalletManagementRepresentation)

	autonomousDatabaseRegionalWalletManagementSingularDataSourceRepresentation = map[string]interface{}{}

	autonomousDatabaseRegionalWalletManagementRepresentation = map[string]interface{}{
		"should_rotate": Representation{repType: Optional, create: `false`, update: `true`},
	}

	AutonomousDatabaseRegionalWalletManagementResourceDependencies = ""
)

func TestDatabaseAutonomousDatabaseRegionalWalletManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousDatabaseRegionalWalletManagementResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database_regional_wallet_management.test_autonomous_database_regional_wallet_management"

	singularDatasourceName := "data.oci_database_autonomous_database_regional_wallet_management.test_autonomous_database_regional_wallet_management"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + AutonomousDatabaseRegionalWalletManagementResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_autonomous_database_regional_wallet_management", "test_autonomous_database_regional_wallet_management", Required, Create, autonomousDatabaseRegionalWalletManagementRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
					generateResourceFromRepresentationMap("oci_database_autonomous_database_regional_wallet_management", "test_autonomous_database_regional_wallet_management", Optional, Update, autonomousDatabaseRegionalWalletManagementRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttrSet(resourceName, "time_rotated"),
					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_database_autonomous_database_regional_wallet_management", "test_autonomous_database_regional_wallet_management", Required, Create, autonomousDatabaseRegionalWalletManagementSingularDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousDatabaseRegionalWalletManagementResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_rotated"),
				),
			},
		},
	})
}
