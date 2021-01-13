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
	BdsAutoScalingConfigurationRequiredOnlyResource = BdsAutoScalingConfigurationResourceDependencies +
		generateResourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration", Required, Create, bdsAutoScalingConfigurationRepresentation)

	BdsAutoScalingConfigurationResourceConfig = BdsAutoScalingConfigurationResourceDependencies +
		generateResourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration", Optional, Update, bdsAutoScalingConfigurationRepresentation)

	bdsAutoScalingConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"auto_scaling_configuration_id": Representation{repType: Required, create: `${oci_bds_auto_scaling_configuration.test_auto_scaling_configuration.id}`},
		"bds_instance_id":               Representation{repType: Required, create: `${oci_bds_bds_instance.test_bds_instance.id}`},
	}

	bdsAutoScalingConfigurationDataSourceRepresentation = map[string]interface{}{
		"bds_instance_id": Representation{repType: Required, create: `${oci_bds_bds_instance.test_bds_instance.id}`},
		"compartment_id":  Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":    Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"state":           Representation{repType: Optional, create: `ACTIVE`},
		"filter":          RepresentationGroup{Required, bdsAutoScalingConfigurationDataSourceFilterRepresentation}}
	bdsAutoScalingConfigurationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_bds_auto_scaling_configuration.test_auto_scaling_configuration.id}`}},
	}

	bdsAutoScalingConfigurationRepresentation = map[string]interface{}{
		"bds_instance_id":        Representation{repType: Required, create: `${oci_bds_bds_instance.test_bds_instance.id}`},
		"cluster_admin_password": Representation{repType: Required, create: `V2VsY29tZTE=`},
		"is_enabled":             Representation{repType: Required, create: `true`},
		"node_type":              Representation{repType: Required, create: `WORKER`},
		"policy":                 RepresentationGroup{Required, autoScalingConfigurationPolicyRepresentation},
		"display_name":           Representation{repType: Optional, create: `displayName`, update: `displayName2`},
	}
	autoScalingConfigurationPolicyRepresentation = map[string]interface{}{
		"policy_type": Representation{repType: Required, create: `THRESHOLD_BASED`, update: `THRESHOLD_BASED`},
		"rules":       []RepresentationGroup{{Required, autoScalingConfigurationPolicyScaleUpRulesRepresentation}, {Required, autoScalingConfigurationPolicyScaleDownRulesRepresentation}},
	}
	autoScalingConfigurationPolicyScaleUpRulesRepresentation = map[string]interface{}{
		"action": Representation{repType: Required, create: `CHANGE_SHAPE_SCALE_UP`, update: `CHANGE_SHAPE_SCALE_UP`},
		"metric": RepresentationGroup{Required, autoScalingConfigurationPolicyScaleUpRulesMetricRepresentation},
	}
	autoScalingConfigurationPolicyScaleUpRulesMetricRepresentation = map[string]interface{}{
		"metric_type": Representation{repType: Required, create: `CPU_UTILIZATION`, update: `CPU_UTILIZATION`},
		"threshold":   RepresentationGroup{Required, autoScalingConfigurationPolicyScaleUpRulesMetricThresholdRepresentation},
	}
	autoScalingConfigurationPolicyScaleUpRulesMetricThresholdRepresentation = map[string]interface{}{
		"duration_in_minutes": Representation{repType: Required, create: `25`, update: `50`},
		"operator":            Representation{repType: Required, create: `GT`, update: `GT`},
		"value":               Representation{repType: Required, create: `80`, update: `90`},
	}
	autoScalingConfigurationPolicyScaleDownRulesRepresentation = map[string]interface{}{
		"action": Representation{repType: Required, create: `CHANGE_SHAPE_SCALE_DOWN`, update: `CHANGE_SHAPE_SCALE_DOWN`},
		"metric": RepresentationGroup{Required, autoScalingConfigurationPolicyScaleDownRulesMetricRepresentation},
	}
	autoScalingConfigurationPolicyScaleDownRulesMetricRepresentation = map[string]interface{}{
		"metric_type": Representation{repType: Required, create: `CPU_UTILIZATION`, update: `CPU_UTILIZATION`},
		"threshold":   RepresentationGroup{Required, autoScalingConfigurationPolicyScaleDownRulesMetricThresholdRepresentation},
	}
	autoScalingConfigurationPolicyScaleDownRulesMetricThresholdRepresentation = map[string]interface{}{
		"duration_in_minutes": Representation{repType: Required, create: `25`, update: `50`},
		"operator":            Representation{repType: Required, create: `LT`, update: `LT`},
		"value":               Representation{repType: Required, create: `15`, update: `20`},
	}

	BdsAutoScalingConfigurationResourceDependencies = generateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", Required, Create, bdsInstanceRepresentation) +
		generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation)
)

func TestBdsAutoScalingConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBdsAutoScalingConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_bds_auto_scaling_configuration.test_auto_scaling_configuration"
	datasourceName := "data.oci_bds_auto_scaling_configurations.test_auto_scaling_configuration"
	singularDatasourceName := "data.oci_bds_auto_scaling_configuration.test_auto_scaling_configuration"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + BdsAutoScalingConfigurationResourceDependencies +
					generateResourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration", Required, Create, bdsAutoScalingConfigurationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
					resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "V2VsY29tZTE="),
					resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "node_type", "WORKER"),
					resource.TestCheckResourceAttr(resourceName, "policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "policy.0.policy_type", "THRESHOLD_BASED"),
					resource.TestCheckResourceAttr(resourceName, "policy.0.rules.#", "2"),
					CheckResourceSetContainsElementWithProperties(resourceName, "policy.0.rules", map[string]string{
						"action":               "CHANGE_SHAPE_SCALE_UP",
						"metric.#":             "1",
						"metric.0.metric_type": "CPU_UTILIZATION",
						"metric.0.threshold.#": "1",
						"metric.0.threshold.0.duration_in_minutes": "25",
						"metric.0.threshold.0.operator":            "GT",
						"metric.0.threshold.0.value":               "80",
					},
						[]string{}),
					CheckResourceSetContainsElementWithProperties(resourceName, "policy.0.rules", map[string]string{
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
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + BdsAutoScalingConfigurationResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + BdsAutoScalingConfigurationResourceDependencies +
					generateResourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration", Optional, Create, bdsAutoScalingConfigurationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
					resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "V2VsY29tZTE="),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "node_type", "WORKER"),
					resource.TestCheckResourceAttr(resourceName, "policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "policy.0.policy_type", "THRESHOLD_BASED"),
					resource.TestCheckResourceAttr(resourceName, "policy.0.rules.#", "2"),
					CheckResourceSetContainsElementWithProperties(resourceName, "policy.0.rules", map[string]string{
						"action":               "CHANGE_SHAPE_SCALE_UP",
						"metric.#":             "1",
						"metric.0.metric_type": "CPU_UTILIZATION",
						"metric.0.threshold.#": "1",
						"metric.0.threshold.0.duration_in_minutes": "25",
						"metric.0.threshold.0.operator":            "GT",
						"metric.0.threshold.0.value":               "80",
					},
						[]string{}),
					CheckResourceSetContainsElementWithProperties(resourceName, "policy.0.rules", map[string]string{
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
				Config: config + compartmentIdVariableStr + BdsAutoScalingConfigurationResourceDependencies +
					generateResourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration", Optional, Update, bdsAutoScalingConfigurationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
					resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "V2VsY29tZTE="),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "node_type", "WORKER"),
					resource.TestCheckResourceAttr(resourceName, "policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "policy.0.policy_type", "THRESHOLD_BASED"),
					resource.TestCheckResourceAttr(resourceName, "policy.0.rules.#", "2"),
					CheckResourceSetContainsElementWithProperties(resourceName, "policy.0.rules", map[string]string{
						"action":               "CHANGE_SHAPE_SCALE_UP",
						"metric.#":             "1",
						"metric.0.metric_type": "CPU_UTILIZATION",
						"metric.0.threshold.#": "1",
						"metric.0.threshold.0.duration_in_minutes": "50",
						"metric.0.threshold.0.operator":            "GT",
						"metric.0.threshold.0.value":               "90",
					},
						[]string{}),
					CheckResourceSetContainsElementWithProperties(resourceName, "policy.0.rules", map[string]string{
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_bds_auto_scaling_configurations", "test_auto_scaling_configuration", Optional, Update, bdsAutoScalingConfigurationDataSourceRepresentation) +
					compartmentIdVariableStr + BdsAutoScalingConfigurationResourceDependencies +
					generateResourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration", Optional, Update, bdsAutoScalingConfigurationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration", Required, Create, bdsAutoScalingConfigurationSingularDataSourceRepresentation) +
					compartmentIdVariableStr + BdsAutoScalingConfigurationResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "auto_scaling_configuration_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "bds_instance_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "node_type", "WORKER"),
					resource.TestCheckResourceAttr(singularDatasourceName, "policy.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "policy.0.policy_type", "THRESHOLD_BASED"),
					resource.TestCheckResourceAttr(singularDatasourceName, "policy.0.rules.#", "2"),
					CheckResourceSetContainsElementWithProperties(resourceName, "policy.0.rules", map[string]string{
						"action":               "CHANGE_SHAPE_SCALE_UP",
						"metric.#":             "1",
						"metric.0.metric_type": "CPU_UTILIZATION",
						"metric.0.threshold.#": "1",
						"metric.0.threshold.0.duration_in_minutes": "50",
						"metric.0.threshold.0.operator":            "GT",
						"metric.0.threshold.0.value":               "90",
					},
						[]string{}),
					CheckResourceSetContainsElementWithProperties(resourceName, "policy.0.rules", map[string]string{
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
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + BdsAutoScalingConfigurationResourceConfig,
			},
			// verify resource import
			{
				Config:            config,
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
