// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
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
	operationsInsightsWarehouseDownloadWarehouseWalletRepresentation = map[string]interface{}{
		"operations_insights_warehouse_id":              Representation{RepType: Required, Create: `${oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id}`},
		"operations_insights_warehouse_wallet_password": Representation{RepType: Required, Create: `password123`},
	}

	OperationsInsightsWarehouseDownloadWarehouseWalletResourceDependencies = GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse", "test_operations_insights_warehouse", Required, Create, operationsInsightsWarehouseRepresentation)
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiOperationsInsightsWarehouseDownloadWarehouseWalletResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpsiOperationsInsightsWarehouseDownloadWarehouseWalletResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_opsi_operations_insights_warehouse_download_warehouse_wallet.test_operations_insights_warehouse_download_warehouse_wallet"

	var resId string
	// Save TF content to create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+OperationsInsightsWarehouseDownloadWarehouseWalletResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse_download_warehouse_wallet", "test_operations_insights_warehouse_download_warehouse_wallet", Required, Create, operationsInsightsWarehouseDownloadWarehouseWalletRepresentation), "operationsinsights", "operationsInsightsWarehouseDownloadWarehouseWallet", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + OperationsInsightsWarehouseDownloadWarehouseWalletResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse_download_warehouse_wallet", "test_operations_insights_warehouse_download_warehouse_wallet", Required, Create, operationsInsightsWarehouseDownloadWarehouseWalletRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "operations_insights_warehouse_id"),
				//resource.TestCheckResourceAttr(resourceName, "operations_insights_warehouse_wallet_password", "operationsInsightsWarehouseWalletPassword"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
	})
}
