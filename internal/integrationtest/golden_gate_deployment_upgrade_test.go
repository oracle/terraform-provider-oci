// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

/*
 Note:
  this test requires an existing deployment, with at least 1 upgrade record.
  Set the following environmentVariables in order to make it work:
   TF_VAR_compartment_id=compartment_id_of_the_deployment
TF_VAR_test_deployment_id=ocid_of_deployment
*/
// issue-routing-tag: golden_gate/default
func TestGoldenGateDeploymentUpgradeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGoldenGateDeploymentUpgradeResource_basic")
	defer httpreplay.SaveScenario()

	const (
		COMPARTMENT_ID     = "compartment_id"
		TEST_DEPLOYMENT_ID = "test_deployment_id"
	)

	var (
		datasourceName         = "data.oci_golden_gate_deployment_upgrades.test_deployment_upgrades"
		singularDatasourceName = "data.oci_golden_gate_deployment_upgrade.test_deployment_upgrade"
	)
	var (
		GoldenGateDeploymentUpgradeResourceDependencies = ""

		goldenGatedeploymentUpgradeDataSourceRepresentation = map[string]interface{}{
			"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
			"deployment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.test_deployment_id}`},
		}

		goldenGateDeploymentUpgradeSingularDataSourceRepresentation = map[string]interface{}{
			"deployment_upgrade_id": acctest.Representation{RepType: acctest.Required,
				Create: `${data.oci_golden_gate_deployment_upgrades.test_deployment_upgrades.deployment_upgrade_collection[0].items[0].id}`},
		}
	)

	config := acctest.ProviderTestConfig() +
		makeVariableStr(COMPARTMENT_ID, t) +
		makeVariableStr(TEST_DEPLOYMENT_ID, t) +
		GoldenGateDeploymentUpgradeResourceDependencies

	var (
		compartmentId = utils.GetEnvSettingWithBlankDefault(COMPARTMENT_ID)
		deploymentId  = utils.GetEnvSettingWithBlankDefault(TEST_DEPLOYMENT_ID)
	)
	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, testAccCheckGoldenGateDeploymentUpgradeDestroy, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_deployment_upgrades", "test_deployment_upgrades", acctest.Required, acctest.Create, goldenGatedeploymentUpgradeDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "deployment_upgrade_collection.0.items.#", "1"),
			),
		},

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_deployment_upgrades", "test_deployment_upgrades", acctest.Required, acctest.Create, goldenGatedeploymentUpgradeDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "deployment_upgrade_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "deployment_upgrade_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "deployment_id", deploymentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "deployment_upgrade_collection.0.items.0.deployment_upgrade_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "deployment_upgrade_collection.0.items.0.is_rollback_allowed"),
				resource.TestCheckResourceAttrSet(datasourceName, "deployment_upgrade_collection.0.items.0.is_security_fix"),
				resource.TestCheckResourceAttrSet(datasourceName, "deployment_upgrade_collection.0.items.0.is_snoozed"),
				resource.TestCheckResourceAttrSet(datasourceName, "deployment_upgrade_collection.0.items.0.ogg_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "deployment_upgrade_collection.0.items.0.previous_ogg_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "deployment_upgrade_collection.0.items.0.release_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "deployment_upgrade_collection.0.items.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "deployment_upgrade_collection.0.items.0.time_finished"),
				resource.TestCheckResourceAttrSet(datasourceName, "deployment_upgrade_collection.0.items.0.time_released"),
				resource.TestCheckResourceAttrSet(datasourceName, "deployment_upgrade_collection.0.items.0.time_started"),
				resource.TestCheckResourceAttrSet(datasourceName, "deployment_upgrade_collection.0.items.0.time_updated"),
			),
		},

		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_deployment_upgrades", "test_deployment_upgrades", acctest.Required, acctest.Create, goldenGatedeploymentUpgradeDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_deployment_upgrade", "test_deployment_upgrade", acctest.Required, acctest.Create, goldenGateDeploymentUpgradeSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_cancel_allowed"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_reschedule_allowed"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deployment_id", deploymentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deployment_upgrade_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_rollback_allowed"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_security_fix"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_snoozed"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ogg_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "previous_ogg_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "release_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_finished"),
				resource.TestCheckNoResourceAttr(singularDatasourceName, "time_ogg_version_supported_until"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_released"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_schedule"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_schedule_max"),
				resource.TestCheckNoResourceAttr(singularDatasourceName, "time_snoozed_until"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_started"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}

func testAccCheckGoldenGateDeploymentUpgradeDestroy(s *terraform.State) error {
	return nil
}
