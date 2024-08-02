// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	BdsAutoScalingConfigurationRequiredOnlyResource = BdsAutoScalingConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Required, acctest.Create, BdsAutoScalingConfigurationRepresentation)

	BdsAutoScalingConfigurationResourceConfig = BdsAutoScalingConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Optional, acctest.Update, BdsAutoScalingConfigurationRepresentation)

	BdsBdsAutoScalingConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"auto_scaling_configuration_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_auto_scaling_configuration.test_auto_scaling_configuration.id}`},
		"bds_instance_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
	}

	BdsBdsAutoScalingConfigurationDataSourceRepresentation = map[string]interface{}{
		"bds_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":    acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":           acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":          acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsAutoScalingConfigurationDataSourceFilterRepresentation}}
	BdsAutoScalingConfigurationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_bds_auto_scaling_configuration.test_auto_scaling_configuration.id}`}},
	}

	BdsAutoScalingConfigurationRepresentation = map[string]interface{}{
		"bds_instance_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
		"cluster_admin_password": acctest.Representation{RepType: acctest.Required, Create: `T3JhY2xlVGVhbVVTQSExMjM=`},
		"is_enabled":             acctest.Representation{RepType: acctest.Required, Create: `true`},
		"node_type":              acctest.Representation{RepType: acctest.Required, Create: `WORKER`},
		"display_name":           acctest.Representation{RepType: acctest.Required, Create: `mem-schedule-vertical`, Update: `mem-schedule-vertical`},
		"policy_details":         acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsAutoScalingConfigurationPolicyDetailsRepresentationScheduleVertical},
	}

	BdsAutoScalingConfigurationPolicyDetailsRepresentationMetricHorizontal = map[string]interface{}{
		"policy_type":      acctest.Representation{RepType: acctest.Required, Create: `METRIC_BASED_HORIZONTAL_SCALING_POLICY`},
		"scale_in_config":  acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsAutoScalingConfigurationPolicyDetailsScaleInConfigRepresentation},
		"scale_out_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsAutoScalingConfigurationPolicyDetailsScaleOutConfigRepresentation},
	}

	BdsAutoScalingConfigurationPolicyDetailsRepresentationMetricVertical = map[string]interface{}{
		"policy_type":       acctest.Representation{RepType: acctest.Required, Create: `METRIC_BASED_VERTICAL_SCALING_POLICY`},
		"scale_down_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsAutoScalingConfigurationPolicyDetailsScaleDownConfigRepresentation},
		"scale_up_config":   acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsAutoScalingConfigurationPolicyDetailsScaleUpConfigRepresentation},
	}

	BdsAutoScalingConfigurationPolicyDetailsRepresentationScheduleHorizontal = map[string]interface{}{
		"policy_type":      acctest.Representation{RepType: acctest.Required, Create: `SCHEDULE_BASED_HORIZONTAL_SCALING_POLICY`},
		"timezone":         acctest.Representation{RepType: acctest.Required, Create: `Asia/Kolkata`, Update: `Asia/Kolkata`},
		"schedule_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsAutoScalingConfigurationPolicyDetailsScheduleHorizontalDetailsRepresentation},
	}

	BdsAutoScalingConfigurationPolicyDetailsRepresentationScheduleVertical = map[string]interface{}{
		"policy_type":      acctest.Representation{RepType: acctest.Required, Create: `SCHEDULE_BASED_VERTICAL_SCALING_POLICY`},
		"timezone":         acctest.Representation{RepType: acctest.Required, Create: `Asia/Kolkata`, Update: `Asia/Kolkata`},
		"schedule_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsAutoScalingConfigurationPolicyDetailsScheduleVerticalDetailsRepresentation},
	}

	BdsAutoScalingConfigurationPolicyDetailsScaleInConfigRepresentation = map[string]interface{}{
		"metric":         acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsAutoScalingConfigurationPolicyDetailsScaleInConfigMetricRepresentation},
		"min_node_count": acctest.Representation{RepType: acctest.Required, Create: `3`, Update: `3`},
		"step_size":      acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `1`},
	}

	BdsAutoScalingConfigurationPolicyDetailsScaleOutConfigRepresentation = map[string]interface{}{
		"max_node_count": acctest.Representation{RepType: acctest.Required, Create: `4`, Update: `4`},
		"metric":         acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsAutoScalingConfigurationPolicyDetailsScaleOutConfigMetricRepresentation},
		"step_size":      acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `1`},
	}

	BdsAutoScalingConfigurationPolicyDetailsScaleDownConfigRepresentation = map[string]interface{}{
		"memory_step_size":    acctest.Representation{RepType: acctest.Required, Create: `2`, Update: `2`},
		"metric":              acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsAutoScalingConfigurationPolicyDetailsScaleDownConfigMetricRepresentation},
		"min_memory_per_node": acctest.Representation{RepType: acctest.Required, Create: `16`, Update: `16`},
		"min_ocpus_per_node":  acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `1`},
		"ocpu_step_size":      acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `1`},
	}

	BdsAutoScalingConfigurationPolicyDetailsScaleUpConfigRepresentation = map[string]interface{}{
		"max_memory_per_node": acctest.Representation{RepType: acctest.Required, Create: `20`, Update: `20`},
		"max_ocpus_per_node":  acctest.Representation{RepType: acctest.Required, Create: `3`, Update: `3`},
		"memory_step_size":    acctest.Representation{RepType: acctest.Required, Create: `2`, Update: `2`},
		"metric":              acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsAutoScalingConfigurationPolicyDetailsScaleUpConfigMetricRepresentation},
		"ocpu_step_size":      acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `1`},
	}

	BdsAutoScalingConfigurationPolicyDetailsScheduleHorizontalDetailsRepresentation = map[string]interface{}{
		"schedule_type":                      acctest.Representation{RepType: acctest.Required, Create: `DAY_BASED`},
		"time_and_horizontal_scaling_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsAutoScalingConfigurationPolicyDetailsScheduleDetailsTimeAndHorizontalScalingConfigRepresentation},
	}

	BdsAutoScalingConfigurationPolicyDetailsScheduleVerticalDetailsRepresentation = map[string]interface{}{
		"schedule_type":                    acctest.Representation{RepType: acctest.Required, Create: `DAY_BASED`},
		"time_and_vertical_scaling_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsAutoScalingConfigurationPolicyDetailsScheduleDetailsTimeAndVerticalScalingConfigRepresentation},
	}

	BdsAutoScalingConfigurationPolicyDetailsScheduleDetailsTimeAndHorizontalScalingConfigRepresentation = map[string]interface{}{
		"time_recurrence":   acctest.Representation{RepType: acctest.Required, Create: `FREQ=WEEKLY;BYDAY=TH;BYHOUR=13;BYMINUTE=30`, Update: `BYDAY=SU;BYHOUR=8,BYMINUTE=30`},
		"target_node_count": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `10`},
	}

	BdsAutoScalingConfigurationPolicyDetailsScheduleDetailsTimeAndVerticalScalingConfigRepresentation = map[string]interface{}{
		"time_recurrence":        acctest.Representation{RepType: acctest.Required, Create: `FREQ=WEEKLY;BYDAY=SU;BYHOUR=8;BYMINUTE=30`, Update: `FREQ=WEEKLY;BYDAY=SU;BYHOUR=8;BYMINUTE=30`},
		"target_shape":           acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E4.Flex`, Update: `VM.Standard.E4.Flex`},
		"target_memory_per_node": acctest.Representation{RepType: acctest.Required, Create: `20`, Update: `20`},
		"target_ocpus_per_node":  acctest.Representation{RepType: acctest.Required, Create: `3`, Update: `3`},
	}

	BdsAutoScalingConfigurationPolicyDetailsScaleDownConfigMetricRepresentation = map[string]interface{}{
		"metric_type": acctest.Representation{RepType: acctest.Required, Create: `MEMORY_UTILIZATION`, Update: `MEMORY_UTILIZATION`},
		"threshold":   acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsAutoScalingConfigurationPolicyDetailsScaleDownConfigMetricThresholdRepresentation},
	}

	BdsAutoScalingConfigurationPolicyDetailsScaleDownConfigMetricThresholdRepresentation = map[string]interface{}{
		"duration_in_minutes": acctest.Representation{RepType: acctest.Required, Create: `25`, Update: `50`},
		"operator":            acctest.Representation{RepType: acctest.Required, Create: `LT`, Update: `LT`},
		"value":               acctest.Representation{RepType: acctest.Required, Create: `35`, Update: `35`},
	}

	BdsAutoScalingConfigurationPolicyDetailsScaleInConfigMetricRepresentation = map[string]interface{}{
		"metric_type": acctest.Representation{RepType: acctest.Required, Create: `MEMORY_UTILIZATION`, Update: `MEMORY_UTILIZATION`},
		"threshold":   acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsAutoScalingConfigurationPolicyDetailsScaleInConfigMetricThresholdRepresentation},
	}

	BdsAutoScalingConfigurationPolicyDetailsScaleInConfigMetricThresholdRepresentation = map[string]interface{}{
		"duration_in_minutes": acctest.Representation{RepType: acctest.Required, Create: `5`, Update: `5`},
		"operator":            acctest.Representation{RepType: acctest.Required, Create: `LT`, Update: `LT`},
		"value":               acctest.Representation{RepType: acctest.Required, Create: `29`, Update: `29`},
	}

	BdsAutoScalingConfigurationPolicyDetailsScaleOutConfigMetricRepresentation = map[string]interface{}{
		"metric_type": acctest.Representation{RepType: acctest.Required, Create: `MEMORY_UTILIZATION`, Update: `MEMORY_UTILIZATION`},
		"threshold":   acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsAutoScalingConfigurationPolicyDetailsScaleOutConfigMetricThresholdRepresentation},
	}

	BdsAutoScalingConfigurationPolicyDetailsScaleOutConfigMetricThresholdRepresentation = map[string]interface{}{
		"duration_in_minutes": acctest.Representation{RepType: acctest.Required, Create: `5`, Update: `3`},
		"operator":            acctest.Representation{RepType: acctest.Required, Create: `GT`, Update: `GT`},
		"value":               acctest.Representation{RepType: acctest.Required, Create: `60`, Update: `60`},
	}

	BdsAutoScalingConfigurationPolicyDetailsScaleUpConfigMetricRepresentation = map[string]interface{}{
		"metric_type": acctest.Representation{RepType: acctest.Required, Create: `MEMORY_UTILIZATION`, Update: `MEMORY_UTILIZATION`},
		"threshold":   acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsAutoScalingConfigurationPolicyDetailsScaleUpConfigMetricThresholdRepresentation},
	}

	BdsAutoScalingConfigurationPolicyDetailsScaleUpConfigMetricThresholdRepresentation = map[string]interface{}{
		"duration_in_minutes": acctest.Representation{RepType: acctest.Required, Create: `3`, Update: `3`},
		"operator":            acctest.Representation{RepType: acctest.Required, Create: `GT`, Update: `GT`},
		"value":               acctest.Representation{RepType: acctest.Required, Create: `80`, Update: `90`},
	}

	BdsAutoScalingConfigurationResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Required, acctest.Create, bdsInstanceOdhRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation)
)

// issue-routing-tag: bds/default
func TestBdsOdhAutoScalingConfigurationResource(t *testing.T) {
	httpreplay.SetScenario("TestBdsOdhAutoScalingConfigurationResource")

	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_ocid")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	//bdsinstanceId := utils.GetEnvSettingWithBlankDefault("bdsinstance_ocid")
	//bdsinstanceIdVariableStr := fmt.Sprintf("variable \"bdsinstance_id\" { default = \"%s\" }\n", bdsinstanceId)

	resourceName := "oci_bds_auto_scaling_configuration.test_auto_scaling_configuration"
	datasourceName := "data.oci_bds_auto_scaling_configurations.test_auto_scaling_configuration"
	singularDatasourceName := "data.oci_bds_auto_scaling_configuration.test_auto_scaling_configuration"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+subnetIdVariableStr+BdsAutoScalingConfigurationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Optional, acctest.Create, BdsAutoScalingConfigurationRepresentation), "bds", "autoScalingConfiguration", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + BdsAutoScalingConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Required, acctest.Create, BdsAutoScalingConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "T3JhY2xlVGVhbVVTQSExMjM="),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "node_type", "WORKER"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.policy_type", "METRIC_BASED_VERTICAL_SCALING_POLICY"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.trigger_type", "METRIC_BASED"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.action_type", "VERTICAL_SCALING"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "policy_details.0.scale_down_config", map[string]string{
					"metric.#":                                 "1",
					"metric.0.metric_type":                     "MEMORY_UTILIZATION",
					"metric.0.threshold.#":                     "1",
					"metric.0.threshold.0.duration_in_minutes": "25",
					"metric.0.threshold.0.operator":            "LT",
					"metric.0.threshold.0.value":               "15",
				},
					[]string{}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "policy_details.0.scale_up_config", map[string]string{
					"metric.#":                                 "1",
					"metric.0.metric_type":                     "MEMORY_UTILIZATION",
					"metric.0.threshold.#":                     "1",
					"metric.0.threshold.0.duration_in_minutes": "25",
					"metric.0.threshold.0.operator":            "GT",
					"metric.0.threshold.0.value":               "80",
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
				acctest.GenerateResourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Optional, acctest.Create, BdsAutoScalingConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "V2VsY29tZTE="),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "node_type", "WORKER"),
				resource.TestCheckResourceAttr(resourceName, "policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.policy_type", "METRIC_BASED_VERTICAL_SCALING_POLICY"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "policy_details.0.scale_down_config", map[string]string{
					"metric.#":                                 "1",
					"metric.0.metric_type":                     "MEMORY_UTILIZATION",
					"metric.0.threshold.#":                     "1",
					"metric.0.threshold.0.duration_in_minutes": "25",
					"metric.0.threshold.0.operator":            "LT",
					"metric.0.threshold.0.value":               "15",
				},
					[]string{}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "policy_details.0.scale_up_config", map[string]string{
					"metric.#":                                 "1",
					"metric.0.metric_type":                     "MEMORY_UTILIZATION",
					"metric.0.threshold.#":                     "1",
					"metric.0.threshold.0.duration_in_minutes": "25",
					"metric.0.threshold.0.operator":            "GT",
					"metric.0.threshold.0.value":               "80",
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
				acctest.GenerateResourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Optional, acctest.Update, BdsAutoScalingConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "V2VsY29tZTE="),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "node_type", "WORKER"),

				resource.TestCheckResourceAttr(resourceName, "policy_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "policy_details.0.action_type"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.policy_type", "METRIC_BASED_HORIZONTAL_SCALING_POLICY"),

				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_down_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_down_config.0.memory_step_size", "11"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_down_config.0.metric.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_down_config.0.metric.0.metric_type", "MEMORY_UTILIZATION"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_down_config.0.metric.0.threshold.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_down_config.0.metric.0.threshold.0.duration_in_minutes", "11"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_down_config.0.metric.0.threshold.0.operator", "LT"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_down_config.0.metric.0.threshold.0.value", "11"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_down_config.0.min_memory_per_node", "11"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_down_config.0.min_ocpus_per_node", "11"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_down_config.0.ocpu_step_size", "11"),

				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_in_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_in_config.0.metric.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_in_config.0.metric.0.metric_type", "MEMORY_UTILIZATION"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_in_config.0.metric.0.threshold.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_in_config.0.metric.0.threshold.0.duration_in_minutes", "11"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_in_config.0.metric.0.threshold.0.operator", "LT"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_in_config.0.metric.0.threshold.0.value", "11"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_in_config.0.min_node_count", "11"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_in_config.0.step_size", "11"),

				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_out_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_out_config.0.max_node_count", "11"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_out_config.0.metric.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_out_config.0.metric.0.metric_type", "MEMORY_UTILIZATION"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_out_config.0.metric.0.threshold.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_out_config.0.metric.0.threshold.0.duration_in_minutes", "11"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_out_config.0.metric.0.threshold.0.operator", "LT"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_out_config.0.metric.0.threshold.0.value", "11"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_out_config.0.step_size", "11"),

				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_up_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_up_config.0.max_memory_per_node", "11"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_up_config.0.max_ocpus_per_node", "11"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_up_config.0.memory_step_size", "11"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_up_config.0.metric.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_up_config.0.metric.0.metric_type", "MEMORY_UTILIZATION"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_up_config.0.metric.0.threshold.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_up_config.0.metric.0.threshold.0.duration_in_minutes", "11"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_up_config.0.metric.0.threshold.0.operator", "LT"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_up_config.0.metric.0.threshold.0.value", "11"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.scale_up_config.0.ocpu_step_size", "11"),

				resource.TestCheckResourceAttr(resourceName, "policy_details.0.schedule_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.schedule_details.0.schedule_type", "DAY_BASED"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.schedule_details.0.time_and_horizontal_scaling_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.schedule_details.0.time_and_horizontal_scaling_config.0.target_node_count", "11"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.schedule_details.0.time_and_horizontal_scaling_config.0.time_recurrence", "timeRecurrence2"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.schedule_details.0.time_and_vertical_scaling_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.schedule_details.0.time_and_vertical_scaling_config.0.target_memory_per_node", "11"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.schedule_details.0.time_and_vertical_scaling_config.0.target_ocpus_per_node", "11"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.schedule_details.0.time_and_vertical_scaling_config.0.target_shape", "targetShape2"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.schedule_details.0.time_and_vertical_scaling_config.0.time_recurrence", "timeRecurrence2"),
				resource.TestCheckResourceAttr(resourceName, "policy_details.0.timezone", "timezone2"),

				resource.TestCheckResourceAttrSet(resourceName, "policy_details.0.trigger_type"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_bds_auto_scaling_configurations", "test_auto_scaling_configuration", acctest.Optional, acctest.Update, BdsBdsAutoScalingConfigurationDataSourceRepresentation) +
				compartmentIdVariableStr + BdsAutoScalingConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Optional, acctest.Update, BdsAutoScalingConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "auto_scaling_configurations.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.node_type", "WORKER"),

				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "auto_scaling_configurations.0.policy_details.0.action_type"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.policy_type", "METRIC_BASED_HORIZONTAL_SCALING_POLICY"),

				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_down_config.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_down_config.0.memory_step_size", "11"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_down_config.0.metric.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_down_config.0.metric.0.metric_type", "MEMORY_UTILIZATION"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_down_config.0.metric.0.threshold.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_down_config.0.metric.0.threshold.0.duration_in_minutes", "11"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_down_config.0.metric.0.threshold.0.operator", "LT"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_down_config.0.metric.0.threshold.0.value", "11"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_down_config.0.min_memory_per_node", "11"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_down_config.0.min_ocpus_per_node", "11"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_down_config.0.ocpu_step_size", "11"),

				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_in_config.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_in_config.0.metric.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_in_config.0.metric.0.metric_type", "MEMORY_UTILIZATION"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_in_config.0.metric.0.threshold.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_in_config.0.metric.0.threshold.0.duration_in_minutes", "11"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_in_config.0.metric.0.threshold.0.operator", "LT"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_in_config.0.metric.0.threshold.0.value", "11"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_in_config.0.min_node_count", "11"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_in_config.0.step_size", "11"),

				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_out_config.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_out_config.0.max_node_count", "11"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_out_config.0.metric.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_out_config.0.metric.0.metric_type", "MEMORY_UTILIZATION"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_out_config.0.metric.0.threshold.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_out_config.0.metric.0.threshold.0.duration_in_minutes", "11"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_out_config.0.metric.0.threshold.0.operator", "LT"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_out_config.0.metric.0.threshold.0.value", "11"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_out_config.0.step_size", "11"),

				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_up_config.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_up_config.0.max_memory_per_node", "11"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_up_config.0.max_ocpus_per_node", "11"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_up_config.0.memory_step_size", "11"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_up_config.0.metric.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_up_config.0.metric.0.metric_type", "MEMORY_UTILIZATION"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_up_config.0.metric.0.threshold.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_up_config.0.metric.0.threshold.0.duration_in_minutes", "11"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_up_config.0.metric.0.threshold.0.operator", "LT"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_up_config.0.metric.0.threshold.0.value", "11"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.scale_up_config.0.ocpu_step_size", "11"),

				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.schedule_details.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.schedule_details.0.schedule_type", "DAY_BASED"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.schedule_details.0.time_and_horizontal_scaling_config.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.schedule_details.0.time_and_horizontal_scaling_config.0.target_node_count", "11"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.schedule_details.0.time_and_horizontal_scaling_config.0.time_recurrence", "timeRecurrence2"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.schedule_details.0.time_and_vertical_scaling_config.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.schedule_details.0.time_and_vertical_scaling_config.0.target_memory_per_node", "11"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.schedule_details.0.time_and_vertical_scaling_config.0.target_ocpus_per_node", "11"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.schedule_details.0.time_and_vertical_scaling_config.0.target_shape", "targetShape2"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.schedule_details.0.time_and_vertical_scaling_config.0.time_recurrence", "timeRecurrence2"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.policy_details.0.timezone", "timezone2"),

				resource.TestCheckResourceAttrSet(datasourceName, "auto_scaling_configurations.0.policy_details.0.trigger_type"),
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
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "policy_details.0.action_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.policy_type", "METRIC_BASED_HORIZONTAL_SCALING_POLICY"),

				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_down_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_down_config.0.memory_step_size", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_down_config.0.metric.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_down_config.0.metric.0.metric_type", "MEMORY_UTILIZATION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_down_config.0.metric.0.threshold.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_down_config.0.metric.0.threshold.0.duration_in_minutes", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_down_config.0.metric.0.threshold.0.operator", "LT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_down_config.0.metric.0.threshold.0.value", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_down_config.0.min_memory_per_node", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_down_config.0.min_ocpus_per_node", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_down_config.0.ocpu_step_size", "11"),

				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_in_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_in_config.0.metric.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_in_config.0.metric.0.metric_type", "MEMORY_UTILIZATION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_in_config.0.metric.0.threshold.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_in_config.0.metric.0.threshold.0.duration_in_minutes", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_in_config.0.metric.0.threshold.0.operator", "LT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_in_config.0.metric.0.threshold.0.value", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_in_config.0.min_node_count", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_in_config.0.step_size", "11"),

				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_out_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_out_config.0.max_node_count", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_out_config.0.metric.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_out_config.0.metric.0.metric_type", "MEMORY_UTILIZATION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_out_config.0.metric.0.threshold.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_out_config.0.metric.0.threshold.0.duration_in_minutes", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_out_config.0.metric.0.threshold.0.operator", "LT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_out_config.0.metric.0.threshold.0.value", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_out_config.0.step_size", "11"),

				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_up_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_up_config.0.max_memory_per_node", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_up_config.0.max_ocpus_per_node", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_up_config.0.memory_step_size", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_up_config.0.metric.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_up_config.0.metric.0.metric_type", "MEMORY_UTILIZATION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_up_config.0.metric.0.threshold.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_up_config.0.metric.0.threshold.0.duration_in_minutes", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_up_config.0.metric.0.threshold.0.operator", "LT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_up_config.0.metric.0.threshold.0.value", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_up_config.0.ocpu_step_size", "11"),

				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.schedule_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.schedule_details.0.schedule_type", "DAY_BASED"),

				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.schedule_details.0.time_and_horizontal_scaling_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.schedule_details.0.time_and_horizontal_scaling_config.0.target_node_count", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.schedule_details.0.time_and_horizontal_scaling_config.0.time_recurrence", "timeRecurrence2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.schedule_details.0.time_and_vertical_scaling_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.schedule_details.0.time_and_vertical_scaling_config.0.target_memory_per_node", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.schedule_details.0.time_and_vertical_scaling_config.0.target_ocpus_per_node", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.schedule_details.0.time_and_vertical_scaling_config.0.target_shape", "targetShape2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.schedule_details.0.time_and_vertical_scaling_config.0.time_recurrence", "timeRecurrence2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.timezone", "timezone2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "policy_details.0.trigger_type"),

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
