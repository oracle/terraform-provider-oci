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

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	OpsiOperationsInsightsWarehouseDownloadWarehouseWalletRepresentation = map[string]interface{}{
		"operations_insights_warehouse_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id}`},
		"operations_insights_warehouse_wallet_password": acctest.Representation{RepType: acctest.Required, Create: `password123`},
	}

	OpsiOperationsInsightsWarehouseDownloadWarehouseWalletResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse", "test_operations_insights_warehouse", acctest.Required, acctest.Create, OpsiOperationsInsightsWarehouseRepresentation)
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiOperationsInsightsWarehouseDownloadWarehouseWalletResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpsiOperationsInsightsWarehouseDownloadWarehouseWalletResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_opsi_operations_insights_warehouse_download_warehouse_wallet.test_operations_insights_warehouse_download_warehouse_wallet"

	var resId string
	// Save TF content to create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OpsiOperationsInsightsWarehouseDownloadWarehouseWalletResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse_download_warehouse_wallet", "test_operations_insights_warehouse_download_warehouse_wallet", acctest.Required, acctest.Create, OpsiOperationsInsightsWarehouseDownloadWarehouseWalletRepresentation), "operationsinsights", "operationsInsightsWarehouseDownloadWarehouseWallet", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + OpsiOperationsInsightsWarehouseDownloadWarehouseWalletResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse_download_warehouse_wallet", "test_operations_insights_warehouse_download_warehouse_wallet", acctest.Required, acctest.Create, OpsiOperationsInsightsWarehouseDownloadWarehouseWalletRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "operations_insights_warehouse_id"),
				//resource.TestCheckResourceAttr(resourceName, "operations_insights_warehouse_wallet_password", "operationsInsightsWarehouseWalletPassword"),

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
	})
}
