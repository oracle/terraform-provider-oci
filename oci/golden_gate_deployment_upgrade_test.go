// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	//goldenGateDeploymentUpgradeSingularDataSourceRepresentation = map[string]interface{}{
	//	"deployment_upgrade_id": Representation{repType: Required, create: `${oci_golden_gate_deployment_upgrade.test_deployment_upgrade.id}`},
	//}

	goldenGatedeploymentUpgradeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"deployment_id":  Representation{RepType: Optional, Create: `${oci_golden_gate_deployment.test_ggsdeployment.id}`},
		"display_name":   Representation{RepType: Optional, Create: `displayName`},
		"state":          Representation{RepType: Optional, Create: `AVAILABLE`},
	}

	DeploymentUpgradeResourceConfig = GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		GenerateResourceFromRepresentationMap("oci_golden_gate_deployment", "test_ggsdeployment", Required, Create, goldenGateDeploymentRepresentation)
)

// issue-routing-tag: golden_gate/default
func TestGoldenGateDeploymentUpgradeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGoldenGateDeploymentUpgradeResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_golden_gate_deployment_upgrades.test_deployment_upgrades"
	//	singularDatasourceName := "data.oci_golden_gate_deployment_upgrade.test_deployment_upgrade"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, testAccCheckGoldenGateDeploymentUpgradeDestroy, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_golden_gate_deployment_upgrades", "test_deployment_upgrades", Required, Create, goldenGatedeploymentUpgradeDataSourceRepresentation) +
				compartmentIdVariableStr + DeploymentUpgradeResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "deployment_upgrade_collection.0.items.#", "0"),
			),
		},
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_golden_gate_deployment_upgrades", "test_deployment_upgrades", Required, Create,
					RepresentationCopyWithNewProperties(goldenGatedeploymentUpgradeDataSourceRepresentation, map[string]interface{}{
						"deployment_id": Representation{RepType: Required, Create: `${oci_golden_gate_deployment.test_ggsdeployment.id}`},
					})) +
				compartmentIdVariableStr + DeploymentUpgradeResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "deployment_upgrade_collection.0.items.#", "0"),
			),
		},
		//	// verify singular datasource
		//	{
		//		Config: config +
		//			generateDataSourceFromRepresentationMap("oci_golden_gate_deployment_upgrade", "test_deployment_upgrade", Required, Create, goldenGateDeploymentUpgradeSingularDataSourceRepresentation) +
		//			compartmentIdVariableStr + DeploymentUpgradeResourceConfig,
		//		Check: ComposeAggregateTestCheckFuncWrapper(
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
