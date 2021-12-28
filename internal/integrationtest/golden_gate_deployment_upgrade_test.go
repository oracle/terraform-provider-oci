// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	//goldenGateDeploymentUpgradeSingularDataSourceRepresentation = map[string]interface{}{
	//	"deployment_upgrade_id": acctest.Representation{RepType: acctest.Required, create: `${oci_golden_gate_deployment_upgrade.test_deployment_upgrade.id}`},
	//}

	goldenGatedeploymentUpgradeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"deployment_id":  acctest.Representation{RepType: acctest.Optional, Create: `${oci_golden_gate_deployment.test_ggsdeployment.id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
	}

	DeploymentUpgradeResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "test_ggsdeployment", acctest.Required, acctest.Create, goldenGateDeploymentRepresentation)
)

// issue-routing-tag: golden_gate/default
func TestGoldenGateDeploymentUpgradeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGoldenGateDeploymentUpgradeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_golden_gate_deployment_upgrades.test_deployment_upgrades"
	//	singularDatasourceName := "data.oci_golden_gate_deployment_upgrade.test_deployment_upgrade"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, testAccCheckGoldenGateDeploymentUpgradeDestroy, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_deployment_upgrades", "test_deployment_upgrades", acctest.Required, acctest.Create, goldenGatedeploymentUpgradeDataSourceRepresentation) +
				compartmentIdVariableStr + DeploymentUpgradeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "deployment_upgrade_collection.0.items.#", "0"),
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_deployment_upgrades", "test_deployment_upgrades", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(goldenGatedeploymentUpgradeDataSourceRepresentation, map[string]interface{}{
						"deployment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_golden_gate_deployment.test_ggsdeployment.id}`},
					})) +
				compartmentIdVariableStr + DeploymentUpgradeResourceConfig,
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
