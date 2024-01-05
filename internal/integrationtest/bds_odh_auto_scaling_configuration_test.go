// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	BdsOdhAutoScalingConfigurationRequiredOnlyResource = BdsOdhAutoScalingConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration_compute_worker", acctest.Optional, acctest.Create, bdsOdhAutoScalingConfigurationRepresentationMetricHorizontalComputeWorker)

	bdsOdhAutoScalingConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"auto_scaling_configuration_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_auto_scaling_configuration.test_auto_scaling_configuration_compute_worker.id}`},
		"bds_instance_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
	}

	bdsOdhAutoScalingConfigurationDataSourceRepresentation = map[string]interface{}{
		"bds_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":    acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"state":           acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":          acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsOdhAutoScalingConfigurationDataSourceFilterRepresentation}}
	bdsOdhAutoScalingConfigurationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_bds_auto_scaling_configuration.test_auto_scaling_configuration_compute_worker.id}`}},
	}

	bdsOdhAutoScalingConfigurationRepresentationMetricBasedVertical = map[string]interface{}{
		"bds_instance_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
		"cluster_admin_password": acctest.Representation{RepType: acctest.Required, Create: `V2VsY29tZTE=`},
		"is_enabled":             acctest.Representation{RepType: acctest.Required, Create: `true`},
		"node_type":              acctest.Representation{RepType: acctest.Required, Create: `WORKER`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"policy_details":         acctest.RepresentationGroup{RepType: acctest.Required, Group: autoScalingOdhConfigurationPolicyDetailsRepresentationMetricVertical},
	}

	autoScalingOdhConfigurationRepresentationMetricVerticalComputeWorker = acctest.RepresentationCopyWithNewProperties(bdsOdhAutoScalingConfigurationRepresentationMetricBasedVertical,
		map[string]interface{}{
			"node_type": acctest.Representation{RepType: acctest.Required, Create: `COMPUTE_ONLY_WORKER`},
		})

	bdsOdhAutoScalingConfigurationRepresentationMetricHorizontalComputeWorker = acctest.RepresentationCopyWithNewProperties(autoScalingOdhConfigurationRepresentationMetricVerticalComputeWorker,
		map[string]interface{}{
			"policy_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: autoScalingOdhConfigurationPolicyDetailsRepresentationMetricHorizontal},
		})

	autoScalingOdhConfigurationPolicyDetailsRepresentationMetricVertical = map[string]interface{}{
		"policy_type":       acctest.Representation{RepType: acctest.Required, Create: `METRIC_BASED_VERTICAL_SCALING_POLICY`},
		"scale_down_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: autoScalingOdhConfigurationPolicyDetailsScaleDownConfigRepresentation},
		"scale_up_config":   acctest.RepresentationGroup{RepType: acctest.Required, Group: autoScalingOdhConfigurationPolicyDetailsScaleUpConfigRepresentation},
	}

	autoScalingOdhConfigurationPolicyDetailsRepresentationMetricHorizontal = map[string]interface{}{
		"policy_type":      acctest.Representation{RepType: acctest.Required, Create: `METRIC_BASED_HORIZONTAL_SCALING_POLICY`},
		"scale_in_config":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: autoScalingOdhConfigurationPolicyDetailsScaleInConfigRepresentation},
		"scale_out_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: autoScalingOdhConfigurationPolicyDetailsScaleOutConfigRepresentation},
	}
	autoScalingOdhConfigurationPolicyDetailsScaleDownConfigRepresentation = map[string]interface{}{
		"memory_step_size":    acctest.Representation{RepType: acctest.Optional, Create: `2`, Update: `4`},
		"metric":              acctest.RepresentationGroup{RepType: acctest.Required, Group: autoScalingOdhConfigurationPolicyDetailsScaleDownConfigMetricRepresentation},
		"min_memory_per_node": acctest.Representation{RepType: acctest.Optional, Create: `16`, Update: `24`},
		"min_ocpus_per_node":  acctest.Representation{RepType: acctest.Optional, Create: `3`, Update: `5`},
		"ocpu_step_size":      acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
	}
	autoScalingOdhConfigurationPolicyDetailsScaleInConfigRepresentation = map[string]interface{}{
		"metric":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: autoScalingOdhConfigurationPolicyDetailsScaleInConfigMetricRepresentation},
		"min_node_count": acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `3`},
		"step_size":      acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
	}
	autoScalingOdhConfigurationPolicyDetailsScaleOutConfigRepresentation = map[string]interface{}{
		"max_node_count": acctest.Representation{RepType: acctest.Optional, Create: `4`, Update: `6`},
		"metric":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: autoScalingOdhConfigurationPolicyDetailsScaleOutConfigMetricRepresentation},
		"step_size":      acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
	}
	autoScalingOdhConfigurationPolicyDetailsScaleUpConfigRepresentation = map[string]interface{}{
		"max_memory_per_node": acctest.Representation{RepType: acctest.Optional, Create: `32`, Update: `64`},
		"max_ocpus_per_node":  acctest.Representation{RepType: acctest.Optional, Create: `6`, Update: `8`},
		"memory_step_size":    acctest.Representation{RepType: acctest.Optional, Create: `2`, Update: `4`},
		"metric":              acctest.RepresentationGroup{RepType: acctest.Required, Group: autoScalingOdhConfigurationPolicyDetailsScaleUpConfigMetricRepresentation},
		"ocpu_step_size":      acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
	}
	autoScalingOdhConfigurationPolicyDetailsScaleDownConfigMetricRepresentation = map[string]interface{}{
		"metric_type": acctest.Representation{RepType: acctest.Required, Create: `CPU_UTILIZATION`},
		"threshold":   acctest.RepresentationGroup{RepType: acctest.Required, Group: autoScalingOdhConfigurationPolicyDetailsScaleDownConfigMetricThresholdRepresentation},
	}
	autoScalingOdhConfigurationPolicyDetailsScaleInConfigMetricRepresentation = map[string]interface{}{
		"metric_type": acctest.Representation{RepType: acctest.Optional, Create: `CPU_UTILIZATION`},
		"threshold":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: autoScalingOdhConfigurationPolicyDetailsScaleInConfigMetricThresholdRepresentation},
	}
	autoScalingOdhConfigurationPolicyDetailsScaleOutConfigMetricRepresentation = map[string]interface{}{
		"metric_type": acctest.Representation{RepType: acctest.Optional, Create: `CPU_UTILIZATION`},
		"threshold":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: autoScalingOdhConfigurationPolicyDetailsScaleOutConfigMetricThresholdRepresentation},
	}
	autoScalingOdhConfigurationPolicyDetailsScaleUpConfigMetricRepresentation = map[string]interface{}{
		"metric_type": acctest.Representation{RepType: acctest.Required, Create: `CPU_UTILIZATION`},
		"threshold":   acctest.RepresentationGroup{RepType: acctest.Required, Group: autoScalingOdhConfigurationPolicyDetailsScaleUpConfigMetricThresholdRepresentation},
	}
	autoScalingOdhConfigurationPolicyDetailsScaleDownConfigMetricThresholdRepresentation = map[string]interface{}{
		"duration_in_minutes": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"operator":            acctest.Representation{RepType: acctest.Required, Create: `LT`, Update: `LT`},
		"value":               acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
	}
	autoScalingOdhConfigurationPolicyDetailsScaleInConfigMetricThresholdRepresentation = map[string]interface{}{
		"duration_in_minutes": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"operator":            acctest.Representation{RepType: acctest.Optional, Create: `LT`, Update: `LT`},
		"value":               acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}
	autoScalingOdhConfigurationPolicyDetailsScaleOutConfigMetricThresholdRepresentation = map[string]interface{}{
		"duration_in_minutes": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"operator":            acctest.Representation{RepType: acctest.Optional, Create: `GT`, Update: `GT`},
		"value":               acctest.Representation{RepType: acctest.Optional, Create: `80`, Update: `90`},
	}
	autoScalingOdhConfigurationPolicyDetailsScaleUpConfigMetricThresholdRepresentation = map[string]interface{}{
		"duration_in_minutes": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"operator":            acctest.Representation{RepType: acctest.Required, Create: `GT`, Update: `GT`},
		"value":               acctest.Representation{RepType: acctest.Required, Create: `80`, Update: `90`},
	}

	bdsInstanceOdhRepresentationComputeWorkerOneOcpu = acctest.RepresentationCopyWithNewProperties(bdsInstanceOdhRepresentation,
		map[string]interface{}{
			"compute_only_worker_node": acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceNodeFlexShapeOneOcpuRepresentation},
		},
	)

	bdsInstanceNodeFlexShapeOneOcpuRepresentation = acctest.RepresentationCopyWithNewProperties(bdsInstanceNodeFlexShapeRepresentation,
		map[string]interface{}{
			"shape_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: bdsInstanceNodesShapeConfigOneOcpuRepresentation},
		})

	bdsInstanceNodesShapeConfigOneOcpuRepresentation = map[string]interface{}{
		"memory_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `16`},
		"ocpus":         acctest.Representation{RepType: acctest.Required, Create: `1`},
	}

	BdsOdhAutoScalingConfigurationResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Required, acctest.Create, bdsInstanceOdhRepresentationComputeWorkerOneOcpu) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation)
)

// issue-routing-tag: bds/default
func TestBdsOdhAutoScalingConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBdsOdhAutoScalingConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceNameWorker := "oci_bds_auto_scaling_configuration.test_auto_scaling_configuration_worker"
	resourceNameComputeWorker := "oci_bds_auto_scaling_configuration.test_auto_scaling_configuration_compute_worker"
	datasourceName := "data.oci_bds_auto_scaling_configurations.test_auto_scaling_configurations_compute_worker"
	singularDatasourceName := "data.oci_bds_auto_scaling_configuration.test_auto_scaling_configuration_compute_worker"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+BdsOdhAutoScalingConfigurationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Optional, acctest.Create, bdsOdhAutoScalingConfigurationRepresentationMetricBasedVertical), "bds", "autoScalingConfiguration", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify create worker auto-scaling config
		{
			Config: config + compartmentIdVariableStr + BdsOdhAutoScalingConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration_worker", acctest.Required, acctest.Create, bdsOdhAutoScalingConfigurationRepresentationMetricBasedVertical),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceNameWorker, "bds_instance_id"),
				resource.TestCheckResourceAttr(resourceNameWorker, "cluster_admin_password", "V2VsY29tZTE="),
				resource.TestCheckResourceAttr(resourceNameWorker, "is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceNameWorker, "node_type", "WORKER"),
				resource.TestCheckResourceAttr(resourceNameWorker, "policy_details.#", "1"),
				resource.TestCheckResourceAttr(resourceNameWorker, "policy_details.0.policy_type", "METRIC_BASED_VERTICAL_SCALING_POLICY"),
				resource.TestCheckResourceAttr(resourceNameWorker, "policy_details.0.trigger_type", "METRIC_BASED"),
				resource.TestCheckResourceAttr(resourceNameWorker, "policy_details.0.action_type", "VERTICAL_SCALING"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceNameWorker, "policy_details.0.scale_up_config", map[string]string{
					"metric.#":                                 "1",
					"metric.0.metric_type":                     "CPU_UTILIZATION",
					"metric.0.threshold.#":                     "1",
					"metric.0.threshold.0.duration_in_minutes": "10",
					"metric.0.threshold.0.operator":            "GT",
					"metric.0.threshold.0.value":               "80",
				},
					[]string{}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceNameWorker, "policy_details.0.scale_down_config", map[string]string{
					"metric.#":                                 "1",
					"metric.0.metric_type":                     "CPU_UTILIZATION",
					"metric.0.threshold.#":                     "1",
					"metric.0.threshold.0.duration_in_minutes": "10",
					"metric.0.threshold.0.operator":            "LT",
					"metric.0.threshold.0.value":               "10",
				},
					[]string{}),
			),
		},

		// verify create compute worker auto-scaling with optionals, metric based vertical scaling for flex compute worker
		{
			Config: config + compartmentIdVariableStr + BdsOdhAutoScalingConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration_worker", acctest.Required, acctest.Create, bdsOdhAutoScalingConfigurationRepresentationMetricBasedVertical) +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration_compute_worker", acctest.Optional, acctest.Create, autoScalingOdhConfigurationRepresentationMetricVerticalComputeWorker),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceNameComputeWorker, "bds_instance_id"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "cluster_admin_password", "V2VsY29tZTE="),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceNameComputeWorker, "id"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "node_type", "COMPUTE_ONLY_WORKER"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceNameComputeWorker, "policy_details.0.action_type"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.policy_type", "METRIC_BASED_VERTICAL_SCALING_POLICY"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_down_config.#", "1"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_down_config.0.memory_step_size", "2"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_down_config.0.metric.#", "1"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_down_config.0.metric.0.metric_type", "CPU_UTILIZATION"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_down_config.0.metric.0.threshold.#", "1"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_down_config.0.metric.0.threshold.0.duration_in_minutes", "10"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_down_config.0.metric.0.threshold.0.operator", "LT"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_down_config.0.metric.0.threshold.0.value", "10"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_down_config.0.min_memory_per_node", "16"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_down_config.0.min_ocpus_per_node", "3"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_down_config.0.ocpu_step_size", "1"),

				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_up_config.#", "1"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_up_config.0.max_memory_per_node", "32"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_up_config.0.max_ocpus_per_node", "6"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_up_config.0.memory_step_size", "2"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_up_config.0.metric.#", "1"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_up_config.0.metric.0.metric_type", "CPU_UTILIZATION"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_up_config.0.metric.0.threshold.#", "1"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_up_config.0.metric.0.threshold.0.duration_in_minutes", "10"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_up_config.0.metric.0.threshold.0.operator", "GT"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_up_config.0.metric.0.threshold.0.value", "80"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_up_config.0.ocpu_step_size", "1"),
				resource.TestCheckResourceAttrSet(resourceNameComputeWorker, "policy_details.0.trigger_type"),
				resource.TestCheckResourceAttrSet(resourceNameComputeWorker, "state"),
				resource.TestCheckResourceAttrSet(resourceNameComputeWorker, "time_created"),
				resource.TestCheckResourceAttrSet(resourceNameComputeWorker, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceNameComputeWorker, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceNameComputeWorker); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters for metric based vertical scaling for flex compute worker
		{
			Config: config + compartmentIdVariableStr + BdsOdhAutoScalingConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration_worker", acctest.Required, acctest.Create, bdsOdhAutoScalingConfigurationRepresentationMetricBasedVertical) +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration_compute_worker", acctest.Optional, acctest.Update, autoScalingOdhConfigurationRepresentationMetricVerticalComputeWorker),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceNameComputeWorker, "bds_instance_id"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "cluster_admin_password", "V2VsY29tZTE="),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceNameComputeWorker, "id"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "node_type", "COMPUTE_ONLY_WORKER"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceNameComputeWorker, "policy_details.0.action_type"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.policy_type", "METRIC_BASED_VERTICAL_SCALING_POLICY"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_down_config.#", "1"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_down_config.0.memory_step_size", "4"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_down_config.0.metric.#", "1"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_down_config.0.metric.0.metric_type", "CPU_UTILIZATION"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_down_config.0.metric.0.threshold.#", "1"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_down_config.0.metric.0.threshold.0.duration_in_minutes", "11"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_down_config.0.metric.0.threshold.0.operator", "LT"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_down_config.0.metric.0.threshold.0.value", "11"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_down_config.0.min_memory_per_node", "24"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_down_config.0.min_ocpus_per_node", "5"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_down_config.0.ocpu_step_size", "2"),

				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_up_config.#", "1"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_up_config.0.max_memory_per_node", "64"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_up_config.0.max_ocpus_per_node", "8"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_up_config.0.memory_step_size", "4"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_up_config.0.metric.#", "1"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_up_config.0.metric.0.metric_type", "CPU_UTILIZATION"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_up_config.0.metric.0.threshold.#", "1"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_up_config.0.metric.0.threshold.0.duration_in_minutes", "11"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_up_config.0.metric.0.threshold.0.operator", "GT"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_up_config.0.metric.0.threshold.0.value", "90"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_up_config.0.ocpu_step_size", "2"),
				resource.TestCheckResourceAttrSet(resourceNameComputeWorker, "policy_details.0.trigger_type"),
				resource.TestCheckResourceAttrSet(resourceNameComputeWorker, "state"),
				resource.TestCheckResourceAttrSet(resourceNameComputeWorker, "time_created"),
				resource.TestCheckResourceAttrSet(resourceNameComputeWorker, "time_updated"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceNameComputeWorker, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// Since policy type can't be updated, delete compute worker scaling config before next create with policy type as "METRIC_BASED_HORIZONTAL_SCALING_POLICY"
		{
			Config: config + compartmentIdVariableStr + BdsOdhAutoScalingConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration_worker", acctest.Required, acctest.Create, bdsOdhAutoScalingConfigurationRepresentationMetricBasedVertical),
		},

		// Add a new metric based horizontal scaling config for compute worker
		{
			Config: config + compartmentIdVariableStr + BdsOdhAutoScalingConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration_worker", acctest.Required, acctest.Create, bdsOdhAutoScalingConfigurationRepresentationMetricBasedVertical) +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration_compute_worker", acctest.Optional, acctest.Create, bdsOdhAutoScalingConfigurationRepresentationMetricHorizontalComputeWorker),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceNameComputeWorker, "bds_instance_id"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "cluster_admin_password", "V2VsY29tZTE="),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceNameComputeWorker, "id"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "node_type", "COMPUTE_ONLY_WORKER"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceNameComputeWorker, "policy_details.0.action_type"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.policy_type", "METRIC_BASED_HORIZONTAL_SCALING_POLICY"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_in_config.#", "1"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_in_config.0.metric.#", "1"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_in_config.0.metric.0.metric_type", "CPU_UTILIZATION"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_in_config.0.metric.0.threshold.#", "1"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_in_config.0.metric.0.threshold.0.duration_in_minutes", "10"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_in_config.0.metric.0.threshold.0.operator", "LT"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_in_config.0.metric.0.threshold.0.value", "10"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_in_config.0.min_node_count", "1"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_in_config.0.step_size", "1"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_out_config.#", "1"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_out_config.0.max_node_count", "4"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_out_config.0.metric.#", "1"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_out_config.0.metric.0.metric_type", "CPU_UTILIZATION"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_out_config.0.metric.0.threshold.#", "1"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_out_config.0.metric.0.threshold.0.duration_in_minutes", "10"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_out_config.0.metric.0.threshold.0.operator", "GT"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_out_config.0.metric.0.threshold.0.value", "80"),
				resource.TestCheckResourceAttr(resourceNameComputeWorker, "policy_details.0.scale_out_config.0.step_size", "1"),
				resource.TestCheckResourceAttrSet(resourceNameComputeWorker, "policy_details.0.trigger_type"),
				resource.TestCheckResourceAttrSet(resourceNameComputeWorker, "state"),
				resource.TestCheckResourceAttrSet(resourceNameComputeWorker, "time_created"),
				resource.TestCheckResourceAttrSet(resourceNameComputeWorker, "time_updated"),
			),
		},

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_bds_auto_scaling_configurations", "test_auto_scaling_configurations_compute_worker", acctest.Optional, acctest.Create, bdsOdhAutoScalingConfigurationDataSourceRepresentation) +
				compartmentIdVariableStr + BdsOdhAutoScalingConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration_compute_worker", acctest.Optional, acctest.Create, bdsOdhAutoScalingConfigurationRepresentationMetricHorizontalComputeWorker),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.display_name", "displayName"),
				resource.TestCheckResourceAttrSet(datasourceName, "auto_scaling_configurations.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.node_type", "COMPUTE_ONLY_WORKER"),
				resource.TestCheckResourceAttrSet(datasourceName, "auto_scaling_configurations.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "auto_scaling_configurations.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "auto_scaling_configurations.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration_compute_worker", acctest.Required, acctest.Create, bdsOdhAutoScalingConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + BdsOdhAutoScalingConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_auto_scaling_configuration", "test_auto_scaling_configuration_compute_worker", acctest.Optional, acctest.Create, bdsOdhAutoScalingConfigurationRepresentationMetricHorizontalComputeWorker),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "auto_scaling_configuration_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_type", "COMPUTE_ONLY_WORKER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "policy_details.0.action_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.policy_type", "METRIC_BASED_HORIZONTAL_SCALING_POLICY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_in_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_in_config.0.metric.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_in_config.0.metric.0.metric_type", "CPU_UTILIZATION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_in_config.0.metric.0.threshold.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_in_config.0.metric.0.threshold.0.duration_in_minutes", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_in_config.0.metric.0.threshold.0.operator", "LT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_in_config.0.metric.0.threshold.0.value", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_in_config.0.min_node_count", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_in_config.0.step_size", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_out_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_out_config.0.max_node_count", "4"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_out_config.0.metric.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_out_config.0.metric.0.metric_type", "CPU_UTILIZATION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_out_config.0.metric.0.threshold.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_out_config.0.metric.0.threshold.0.duration_in_minutes", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_out_config.0.metric.0.threshold.0.operator", "GT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_out_config.0.metric.0.threshold.0.value", "80"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_details.0.scale_out_config.0.step_size", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "policy_details.0.trigger_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + BdsOdhAutoScalingConfigurationRequiredOnlyResource,
			ImportState:       true,
			ImportStateIdFunc: getBdsAutoScalingConfigurationCompositeId(resourceNameComputeWorker),
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"cluster_admin_password",
				"is_enabled",
				"time_updated",
			},
			ResourceName: resourceNameComputeWorker,
		},
	})
}
