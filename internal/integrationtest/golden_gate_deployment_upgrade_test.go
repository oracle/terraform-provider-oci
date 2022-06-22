// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"terraform-provider-oci/httpreplay"
)

// issue-routing-tag: golden_gate/default
func TestGoldenGateDeploymentUpgradeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGoldenGateDeploymentUpgradeResource_basic")
	defer httpreplay.SaveScenario()

	const (
		COMPARTMENT_ID     = "compartment_id"
		TEST_DEPLOYMENT_ID = "test_deployment_id"
	)

	var (
		datasourceName = "data.oci_golden_gate_deployment_upgrades.test_deployment_upgrades"
	)
	var (
		GoldenGateDeploymentUpgradeResourceDependencies = ""

		goldenGatedeploymentUpgradeDataSourceRepresentation = map[string]interface{}{
			"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
			"deployment_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.test_deployment_id}`},
		}
	)

	config := acctest.ProviderTestConfig() +
		makeVariableStr(COMPARTMENT_ID, t) +
		makeVariableStr(TEST_DEPLOYMENT_ID, t) +
		GoldenGateDeploymentUpgradeResourceDependencies

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, testAccCheckGoldenGateDeploymentUpgradeDestroy, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_deployment_upgrades", "test_deployment_upgrades", acctest.Required, acctest.Create, goldenGatedeploymentUpgradeDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "deployment_upgrade_collection.0.items.#", "0"),
			),
		},

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_deployment_upgrades", "test_deployment_upgrades", acctest.Required, acctest.Create, goldenGatedeploymentUpgradeDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "deployment_upgrade_collection.0.items.#", "0"),
			),
		},

		//	// verify singular datasource
		//	{
		//		Config: config +
		//			generateDataSourceFromRepresentationMap("oci_golden_gate_deployment_upgrade", "test_deployment_upgrade", acctest.Required, acctest.Create, goldenGateDeploymentUpgradeSingularDataSourceRepresentation) +
		//			compartmentIdVariableStr + DeploymentUpgradeResourceConfig,
		//		Check: acctest.ComposeAggregateTestCheckFuncWrapper(
		//			resource.TestCheckResourceAttrSet(singularDatasourceName, "deployment_upgrade_id"),
		//
		//			resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
		//			resource.TestCheckResourceAttrSet(singularDatasourceName, "deployment_upgrade_type"),
		//			resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
		//			resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
		//			resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
		//			resource.TestCheckResourceAttrSet(singularDatasourceName, "lifecycle_sub_state"),
		//			resource.TestCheckResourceAttrSet(singularDatasourceName, "ogg_version"),
		//			resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
		//			resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
		//			resource.TestCheckResourceAttrSet(singularDatasourceName, "time_finished"),
		//			resource.TestCheckResourceAttrSet(singularDatasourceName, "time_started"),
		//			resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
		//		),
		//	},
	})
}

func testAccCheckGoldenGateDeploymentUpgradeDestroy(s *terraform.State) error {
	return nil
}
