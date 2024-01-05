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
	BdsAutoScalingConfigurationRequiredOnlyResource = BdsAutoScalingConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Required, acctest.Create, BdsautoScalingConfigurationRepresentation)

	BdsAutoScalingConfigurationResourceConfig = BdsAutoScalingConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Optional, acctest.Update, BdsautoScalingConfigurationRepresentation)

	BdsBdsautoScalingConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"auto_scaling_configuration_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_auto_scaling_configuration.test_auto_scaling_configuration.id}`},
		"bds_instance_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
	}

	BdsBdsautoScalingConfigurationDataSourceRepresentation = map[string]interface{}{
		"bds_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":    acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":           acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":          acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsautoScalingConfigurationDataSourceFilterRepresentation}}
	BdsautoScalingConfigurationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_bds_auto_scaling_configuration.test_auto_scaling_configuration.id}`}},
	}

	BdsautoScalingConfigurationRepresentation = map[string]interface{}{
		"bds_instance_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
		"cluster_admin_password": acctest.Representation{RepType: acctest.Required, Create: `V2VsY29tZTE=`},
		"is_enabled":             acctest.Representation{RepType: acctest.Required, Create: `true`},
		"node_type":              acctest.Representation{RepType: acctest.Required, Create: `WORKER`},
		"policy":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsautoScalingConfigurationPolicyRepresentation},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
	}
	BdsautoScalingConfigurationPolicyRepresentation = map[string]interface{}{
		"policy_type": acctest.Representation{RepType: acctest.Required, Create: `THRESHOLD_BASED`, Update: `THRESHOLD_BASED`},
		"rules":       []acctest.RepresentationGroup{{RepType: acctest.Required, Group: BdsautoScalingConfigurationPolicyScaleUpRulesRepresentation}, {RepType: acctest.Required, Group: BdsautoScalingConfigurationPolicyScaleDownRulesRepresentation}},
	}
	BdsautoScalingConfigurationPolicyScaleUpRulesRepresentation = map[string]interface{}{
		"action": acctest.Representation{RepType: acctest.Required, Create: `CHANGE_SHAPE_SCALE_UP`, Update: `CHANGE_SHAPE_SCALE_UP`},
		"metric": acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsautoScalingConfigurationPolicyScaleUpRulesMetricRepresentation},
	}
	BdsautoScalingConfigurationPolicyScaleUpRulesMetricRepresentation = map[string]interface{}{
		"metric_type": acctest.Representation{RepType: acctest.Required, Create: `CPU_UTILIZATION`, Update: `CPU_UTILIZATION`},
		"threshold":   acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsautoScalingConfigurationPolicyScaleUpRulesMetricThresholdRepresentation},
	}
	BdsautoScalingConfigurationPolicyScaleUpRulesMetricThresholdRepresentation = map[string]interface{}{
		"duration_in_minutes": acctest.Representation{RepType: acctest.Required, Create: `25`, Update: `50`},
		"operator":            acctest.Representation{RepType: acctest.Required, Create: `GT`, Update: `GT`},
		"value":               acctest.Representation{RepType: acctest.Required, Create: `80`, Update: `90`},
	}
	BdsautoScalingConfigurationPolicyScaleDownRulesRepresentation = map[string]interface{}{
		"action": acctest.Representation{RepType: acctest.Required, Create: `CHANGE_SHAPE_SCALE_DOWN`, Update: `CHANGE_SHAPE_SCALE_DOWN`},
		"metric": acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsautoScalingConfigurationPolicyScaleDownRulesMetricRepresentation},
	}
	BdsautoScalingConfigurationPolicyScaleDownRulesMetricRepresentation = map[string]interface{}{
		"metric_type": acctest.Representation{RepType: acctest.Required, Create: `CPU_UTILIZATION`, Update: `CPU_UTILIZATION`},
		"threshold":   acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsautoScalingConfigurationPolicyScaleDownRulesMetricThresholdRepresentation},
	}
	BdsautoScalingConfigurationPolicyScaleDownRulesMetricThresholdRepresentation = map[string]interface{}{
		"duration_in_minutes": acctest.Representation{RepType: acctest.Required, Create: `25`, Update: `50`},
		"operator":            acctest.Representation{RepType: acctest.Required, Create: `LT`, Update: `LT`},
		"value":               acctest.Representation{RepType: acctest.Required, Create: `15`, Update: `20`},
	}

	BdsAutoScalingConfigurationResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Required, acctest.Create, BdsbdsInstanceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation)
)

// issue-routing-tag: bds/default
func TestBdsAutoScalingConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBdsAutoScalingConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_bds_auto_scaling_configuration.test_auto_scaling_configuration"
	datasourceName := "data.oci_bds_auto_scaling_configurations.test_auto_scaling_configuration"
	singularDatasourceName := "data.oci_bds_auto_scaling_configuration.test_auto_scaling_configuration"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+BdsAutoScalingConfigurationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Optional, acctest.Create, BdsautoScalingConfigurationRepresentation), "bds", "autoScalingConfiguration", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + BdsAutoScalingConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Required, acctest.Create, BdsautoScalingConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "V2VsY29tZTE="),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "node_type", "WORKER"),
				resource.TestCheckResourceAttr(resourceName, "policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy.0.policy_type", "THRESHOLD_BASED"),
				resource.TestCheckResourceAttr(resourceName, "policy.0.rules.#", "2"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "policy.0.rules", map[string]string{
					"action":               "CHANGE_SHAPE_SCALE_UP",
					"metric.#":             "1",
					"metric.0.metric_type": "CPU_UTILIZATION",
					"metric.0.threshold.#": "1",
					"metric.0.threshold.0.duration_in_minutes": "25",
					"metric.0.threshold.0.operator":            "GT",
					"metric.0.threshold.0.value":               "80",
				},
					[]string{}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "policy.0.rules", map[string]string{
					"action":               "CHANGE_SHAPE_SCALE_DOWN",
					"metric.#":             "1",
					"metric.0.metric_type": "CPU_UTILIZATION",
					"metric.0.threshold.#": "1",
					"metric.0.threshold.0.duration_in_minutes": "25",
					"metric.0.threshold.0.operator":            "LT",
					"metric.0.threshold.0.value":               "15",
				},
					[]string{}),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + BdsAutoScalingConfigurationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + BdsAutoScalingConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Optional, acctest.Create, BdsautoScalingConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "V2VsY29tZTE="),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "node_type", "WORKER"),
				resource.TestCheckResourceAttr(resourceName, "policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy.0.policy_type", "THRESHOLD_BASED"),
				resource.TestCheckResourceAttr(resourceName, "policy.0.rules.#", "2"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "policy.0.rules", map[string]string{
					"action":               "CHANGE_SHAPE_SCALE_UP",
					"metric.#":             "1",
					"metric.0.metric_type": "CPU_UTILIZATION",
					"metric.0.threshold.#": "1",
					"metric.0.threshold.0.duration_in_minutes": "25",
					"metric.0.threshold.0.operator":            "GT",
					"metric.0.threshold.0.value":               "80",
				},
					[]string{}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "policy.0.rules", map[string]string{
					"action":               "CHANGE_SHAPE_SCALE_DOWN",
					"metric.#":             "1",
					"metric.0.metric_type": "CPU_UTILIZATION",
					"metric.0.threshold.#": "1",
					"metric.0.threshold.0.duration_in_minutes": "25",
					"metric.0.threshold.0.operator":            "LT",
					"metric.0.threshold.0.value":               "15",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
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
			Config: config + compartmentIdVariableStr + BdsAutoScalingConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Optional, acctest.Update, BdsautoScalingConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "V2VsY29tZTE="),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "node_type", "WORKER"),
				resource.TestCheckResourceAttr(resourceName, "policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy.0.policy_type", "THRESHOLD_BASED"),
				resource.TestCheckResourceAttr(resourceName, "policy.0.rules.#", "2"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "policy.0.rules", map[string]string{
					"action":               "CHANGE_SHAPE_SCALE_UP",
					"metric.#":             "1",
					"metric.0.metric_type": "CPU_UTILIZATION",
					"metric.0.threshold.#": "1",
					"metric.0.threshold.0.duration_in_minutes": "50",
					"metric.0.threshold.0.operator":            "GT",
					"metric.0.threshold.0.value":               "90",
				},
					[]string{}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "policy.0.rules", map[string]string{
					"action":               "CHANGE_SHAPE_SCALE_DOWN",
					"metric.#":             "1",
					"metric.0.metric_type": "CPU_UTILIZATION",
					"metric.0.threshold.#": "1",
					"metric.0.threshold.0.duration_in_minutes": "50",
					"metric.0.threshold.0.operator":            "LT",
					"metric.0.threshold.0.value":               "20",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_bds_auto_scaling_configurations", "test_auto_scaling_configuration", acctest.Optional, acctest.Update, BdsBdsautoScalingConfigurationDataSourceRepresentation) +
				compartmentIdVariableStr + BdsAutoScalingConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Optional, acctest.Update, BdsautoScalingConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "auto_scaling_configurations.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.node_type", "WORKER"),
				resource.TestCheckResourceAttrSet(datasourceName, "auto_scaling_configurations.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "auto_scaling_configurations.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "auto_scaling_configurations.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Required, acctest.Create, BdsBdsautoScalingConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + BdsAutoScalingConfigurationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "auto_scaling_configuration_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bds_instance_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_type", "WORKER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy.0.policy_type", "THRESHOLD_BASED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy.0.rules.#", "2"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "policy.0.rules", map[string]string{
					"action":               "CHANGE_SHAPE_SCALE_UP",
					"metric.#":             "1",
					"metric.0.metric_type": "CPU_UTILIZATION",
					"metric.0.threshold.#": "1",
					"metric.0.threshold.0.duration_in_minutes": "50",
					"metric.0.threshold.0.operator":            "GT",
					"metric.0.threshold.0.value":               "90",
				},
					[]string{}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "policy.0.rules", map[string]string{
					"action":               "CHANGE_SHAPE_SCALE_DOWN",
					"metric.#":             "1",
					"metric.0.metric_type": "CPU_UTILIZATION",
					"metric.0.threshold.#": "1",
					"metric.0.threshold.0.duration_in_minutes": "50",
					"metric.0.threshold.0.operator":            "LT",
					"metric.0.threshold.0.value":               "20",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + BdsAutoScalingConfigurationRequiredOnlyResource,
			ImportState:       true,
			ImportStateIdFunc: getBdsAutoScalingConfigurationCompositeId(resourceName),
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"cluster_admin_password",
				"is_enabled",
				"time_updated",
			},
			ResourceName: resourceName,
		},
	})
}

func getBdsAutoScalingConfigurationCompositeId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}

		return fmt.Sprintf("bdsInstances/%s/autoScalingConfiguration/%s", rs.Primary.Attributes["bds_instance_id"], rs.Primary.Attributes["id"]), nil
	}
}
