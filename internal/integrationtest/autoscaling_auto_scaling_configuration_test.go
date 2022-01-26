// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_auto_scaling "github.com/oracle/oci-go-sdk/v56/autoscaling"
	"github.com/oracle/oci-go-sdk/v56/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	AutoScalingConfigurationRequiredOnlyResource = AutoScalingConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_autoscaling_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Required, acctest.Create, autoScalingConfigurationRepresentation)

	AutoScalingConfigurationResourceConfig = AutoScalingConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_autoscaling_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Optional, acctest.Update, autoScalingConfigurationRepresentation)

	autoScalingConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"auto_scaling_configuration_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_autoscaling_auto_scaling_configuration.test_auto_scaling_configuration.id}`},
	}

	autoScalingConfigurationDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `example_threshold_autoscaling_configuration`, Update: `displayName2`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: autoScalingConfigurationDataSourceFilterRepresentation}}
	autoScalingConfigurationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_autoscaling_auto_scaling_configuration.test_auto_scaling_configuration.id}`}},
	}

	autoScalingConfigurationRepresentation = map[string]interface{}{
		"auto_scaling_resources": acctest.RepresentationGroup{RepType: acctest.Required, Group: autoScalingConfigurationAutoScalingResourcesRepresentation},
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"policies":               acctest.RepresentationGroup{RepType: acctest.Required, Group: autoScalingConfigurationPoliciesRepresentation},
		"cool_down_in_seconds":   acctest.Representation{RepType: acctest.Optional, Create: `300`, Update: `400`},
		"defined_tags":           acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `example_threshold_autoscaling_configuration`, Update: `displayName2`},
		"freeform_tags":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_enabled":             acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	autoScalingConfigurationAutoScalingResourcesRepresentation = map[string]interface{}{
		"id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance_pool.test_instance_pool.id}`},
		"type": acctest.Representation{RepType: acctest.Required, Create: `instancePool`},
	}
	autoScalingConfigurationPoliciesRepresentation = map[string]interface{}{
		"capacity":     acctest.RepresentationGroup{RepType: acctest.Required, Group: autoScalingConfigurationPoliciesCapacityRepresentation},
		"policy_type":  acctest.Representation{RepType: acctest.Required, Create: `threshold`, Update: `threshold`},
		"rules":        []acctest.RepresentationGroup{{RepType: acctest.Required, Group: autoScalingConfigurationPoliciesScaleOutRuleRepresentation}, {RepType: acctest.Required, Group: autoScalingConfigurationPoliciesScaleInRuleRepresentation}},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `example_autoscaling_configuration`, Update: `displayName2`},
	}
	autoScalingConfigurationPoliciesCapacityRepresentation = map[string]interface{}{
		"initial": acctest.Representation{RepType: acctest.Required, Create: `2`, Update: `4`},
		"max":     acctest.Representation{RepType: acctest.Required, Create: `3`, Update: `5`},
		"min":     acctest.Representation{RepType: acctest.Required, Create: `2`, Update: `3`},
	}
	autoScalingConfigurationPoliciesScaleOutRuleRepresentation = map[string]interface{}{
		"action":       acctest.RepresentationGroup{RepType: acctest.Required, Group: autoScalingConfigurationPoliciesScaleOutRuleActionRepresentation},
		"display_name": acctest.Representation{RepType: acctest.Required, Create: `scale out rule`, Update: `scale out rule - updated`},
		"metric":       acctest.RepresentationGroup{RepType: acctest.Required, Group: autoScalingConfigurationPoliciesScaleOutRuleMetricRepresentation},
	}
	autoScalingConfigurationPoliciesScaleOutRuleActionRepresentation = map[string]interface{}{
		"type":  acctest.Representation{RepType: acctest.Required, Create: `CHANGE_COUNT_BY`, Update: `CHANGE_COUNT_BY`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
	}
	autoScalingConfigurationPoliciesScaleOutRuleMetricRepresentation = map[string]interface{}{
		"metric_type": acctest.Representation{RepType: acctest.Required, Create: `CPU_UTILIZATION`, Update: `CPU_UTILIZATION`},
		"threshold":   acctest.RepresentationGroup{RepType: acctest.Required, Group: autoScalingConfigurationPoliciesScaleOutRuleMetricThresholdRepresentation},
	}
	autoScalingConfigurationPoliciesScaleOutRuleMetricThresholdRepresentation = map[string]interface{}{
		"operator": acctest.Representation{RepType: acctest.Required, Create: `GT`, Update: `GT`},
		"value":    acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `3`},
	}
	autoScalingConfigurationPoliciesScaleInRuleRepresentation = map[string]interface{}{
		"action":       acctest.RepresentationGroup{RepType: acctest.Required, Group: autoScalingConfigurationPoliciesScaleInRuleActionRepresentation},
		"metric":       acctest.RepresentationGroup{RepType: acctest.Required, Group: autoScalingConfigurationPoliciesScaleInRuleMetricRepresentation},
		"display_name": acctest.Representation{RepType: acctest.Required, Create: `scale in rule`, Update: `scale in rule - updated`},
	}
	autoScalingConfigurationPoliciesScaleInRuleActionRepresentation = map[string]interface{}{
		"type":  acctest.Representation{RepType: acctest.Required, Create: `CHANGE_COUNT_BY`, Update: `CHANGE_COUNT_BY`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `-1`, Update: `-3`},
	}
	autoScalingConfigurationPoliciesScaleInRuleMetricRepresentation = map[string]interface{}{
		"metric_type": acctest.Representation{RepType: acctest.Required, Create: `CPU_UTILIZATION`, Update: `CPU_UTILIZATION`},
		"threshold":   acctest.RepresentationGroup{RepType: acctest.Required, Group: autoScalingConfigurationPoliciesScaleInRuleMetricThresholdRepresentation},
	}
	autoScalingConfigurationPoliciesScaleInRuleMetricThresholdRepresentation = map[string]interface{}{
		"operator": acctest.Representation{RepType: acctest.Required, Create: `LT`, Update: `LT`},
		"value":    acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `3`},
	}

	AutoScalingConfigurationResourceDependencies = InstancePoolResourceDependenciesWithoutSecondaryVnic +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", acctest.Required, acctest.Create, instancePoolRepresentation)
)

// issue-routing-tag: auto_scaling/default
func TestAutoScalingAutoScalingConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAutoScalingAutoScalingConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_autoscaling_auto_scaling_configuration.test_auto_scaling_configuration"
	datasourceName := "data.oci_autoscaling_auto_scaling_configurations.test_auto_scaling_configurations"
	singularDatasourceName := "data.oci_autoscaling_auto_scaling_configuration.test_auto_scaling_configuration"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+AutoScalingConfigurationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_autoscaling_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Optional, acctest.Create, autoScalingConfigurationRepresentation), "autoscaling", "autoScalingConfiguration", t)

	acctest.ResourceTest(t, testAccCheckAutoScalingAutoScalingConfigurationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + AutoScalingConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_autoscaling_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Required, acctest.Create, autoScalingConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "auto_scaling_resources.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "auto_scaling_resources.0.id"),
				resource.TestCheckResourceAttr(resourceName, "auto_scaling_resources.0.type", "instancePool"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "policies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.0.initial", "2"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.0.max", "3"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.0.min", "2"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.policy_type", "threshold"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.rules.#", "2"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "policies.0.rules", map[string]string{
					"action.#":                      "1",
					"action.0.type":                 "CHANGE_COUNT_BY",
					"action.0.value":                "1",
					"metric.#":                      "1",
					"metric.0.metric_type":          "CPU_UTILIZATION",
					"metric.0.threshold.#":          "1",
					"metric.0.threshold.0.operator": "GT",
					"metric.0.threshold.0.value":    "1",
				},
					[]string{}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "policies.0.rules", map[string]string{
					"action.#":                      "1",
					"action.0.type":                 "CHANGE_COUNT_BY",
					"action.0.value":                "-1",
					"metric.#":                      "1",
					"metric.0.metric_type":          "CPU_UTILIZATION",
					"metric.0.threshold.#":          "1",
					"metric.0.threshold.0.operator": "LT",
					"metric.0.threshold.0.value":    "1",
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
			Config: config + compartmentIdVariableStr + AutoScalingConfigurationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AutoScalingConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_autoscaling_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Optional, acctest.Create, autoScalingConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "auto_scaling_resources.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "auto_scaling_resources.0.id"),
				resource.TestCheckResourceAttr(resourceName, "auto_scaling_resources.0.type", "instancePool"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cool_down_in_seconds", "300"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_threshold_autoscaling_configuration"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "policies.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "policies.0.id"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.0.initial", "2"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.0.max", "3"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.0.min", "2"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.display_name", "example_autoscaling_configuration"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.policy_type", "threshold"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.rules.#", "2"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "policies.0.rules", map[string]string{
					"action.#":                      "1",
					"action.0.type":                 "CHANGE_COUNT_BY",
					"action.0.value":                "1",
					"display_name":                  "scale out rule",
					"metric.#":                      "1",
					"metric.0.metric_type":          "CPU_UTILIZATION",
					"metric.0.threshold.#":          "1",
					"metric.0.threshold.0.operator": "GT",
					"metric.0.threshold.0.value":    "1",
				},
					[]string{}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "policies.0.rules", map[string]string{
					"action.#":                      "1",
					"action.0.type":                 "CHANGE_COUNT_BY",
					"action.0.value":                "-1",
					"display_name":                  "scale in rule",
					"metric.#":                      "1",
					"metric.0.metric_type":          "CPU_UTILIZATION",
					"metric.0.threshold.#":          "1",
					"metric.0.threshold.0.operator": "LT",
					"metric.0.threshold.0.value":    "1",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "policies.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + AutoScalingConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_autoscaling_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(autoScalingConfigurationRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "auto_scaling_resources.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "auto_scaling_resources.0.id"),
				resource.TestCheckResourceAttr(resourceName, "auto_scaling_resources.0.type", "instancePool"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "cool_down_in_seconds", "300"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_threshold_autoscaling_configuration"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "policies.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "policies.0.id"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.0.initial", "2"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.0.max", "3"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.0.min", "2"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.display_name", "example_autoscaling_configuration"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.policy_type", "threshold"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.rules.#", "2"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "policies.0.rules", map[string]string{
					"action.#":                      "1",
					"action.0.type":                 "CHANGE_COUNT_BY",
					"action.0.value":                "1",
					"display_name":                  "scale out rule",
					"metric.#":                      "1",
					"metric.0.metric_type":          "CPU_UTILIZATION",
					"metric.0.threshold.#":          "1",
					"metric.0.threshold.0.operator": "GT",
					"metric.0.threshold.0.value":    "1",
				},
					[]string{}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "policies.0.rules", map[string]string{
					"action.#":                      "1",
					"action.0.type":                 "CHANGE_COUNT_BY",
					"action.0.value":                "-1",
					"display_name":                  "scale in rule",
					"metric.#":                      "1",
					"metric.0.metric_type":          "CPU_UTILIZATION",
					"metric.0.threshold.#":          "1",
					"metric.0.threshold.0.operator": "LT",
					"metric.0.threshold.0.value":    "1",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "policies.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + AutoScalingConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_autoscaling_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Optional, acctest.Update, autoScalingConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "auto_scaling_resources.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "auto_scaling_resources.0.id"),
				resource.TestCheckResourceAttr(resourceName, "auto_scaling_resources.0.type", "instancePool"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cool_down_in_seconds", "400"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "policies.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "policies.0.id"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.0.initial", "4"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.0.max", "5"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.0.min", "3"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.policy_type", "threshold"),
				resource.TestCheckResourceAttr(resourceName, "policies.0.rules.#", "2"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "policies.0.rules", map[string]string{
					"action.#":                      "1",
					"action.0.type":                 "CHANGE_COUNT_BY",
					"action.0.value":                "2",
					"display_name":                  "scale out rule - updated",
					"metric.#":                      "1",
					"metric.0.metric_type":          "CPU_UTILIZATION",
					"metric.0.threshold.#":          "1",
					"metric.0.threshold.0.operator": "GT",
					"metric.0.threshold.0.value":    "3",
				},
					[]string{}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "policies.0.rules", map[string]string{
					"action.#":                      "1",
					"action.0.type":                 "CHANGE_COUNT_BY",
					"action.0.value":                "-3",
					"display_name":                  "scale in rule - updated",
					"metric.#":                      "1",
					"metric.0.metric_type":          "CPU_UTILIZATION",
					"metric.0.threshold.#":          "1",
					"metric.0.threshold.0.operator": "LT",
					"metric.0.threshold.0.value":    "3",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "policies.0.time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId == resId2 {
						return fmt.Errorf("Resource updated when it was supposed to be recreated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_autoscaling_auto_scaling_configurations", "test_auto_scaling_configurations", acctest.Optional, acctest.Update, autoScalingConfigurationDataSourceRepresentation) +
				compartmentIdVariableStr + AutoScalingConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_autoscaling_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Optional, acctest.Update, autoScalingConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.auto_scaling_resources.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "auto_scaling_configurations.0.auto_scaling_resources.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.auto_scaling_resources.0.type", "instancePool"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.cool_down_in_seconds", "400"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "auto_scaling_configurations.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.is_enabled", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "auto_scaling_configurations.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_autoscaling_auto_scaling_configuration", "test_auto_scaling_configuration", acctest.Required, acctest.Create, autoScalingConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AutoScalingConfigurationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "auto_scaling_configuration_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "auto_scaling_resources.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "auto_scaling_resources.0.id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "auto_scaling_resources.0.type", "instancePool"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "cool_down_in_seconds", "400"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policies.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policies.0.capacity.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policies.0.capacity.0.initial", "4"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policies.0.capacity.0.max", "5"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policies.0.capacity.0.min", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policies.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policies.0.policy_type", "threshold"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policies.0.rules.#", "2"),
				acctest.CheckResourceSetContainsElementWithProperties(singularDatasourceName, "policies.0.rules", map[string]string{
					"action.#":                      "1",
					"action.0.type":                 "CHANGE_COUNT_BY",
					"action.0.value":                "2",
					"display_name":                  "scale out rule - updated",
					"metric.#":                      "1",
					"metric.0.metric_type":          "CPU_UTILIZATION",
					"metric.0.threshold.#":          "1",
					"metric.0.threshold.0.operator": "GT",
					"metric.0.threshold.0.value":    "3",
				},
					[]string{}),
				acctest.CheckResourceSetContainsElementWithProperties(singularDatasourceName, "policies.0.rules", map[string]string{
					"action.#":                      "1",
					"action.0.type":                 "CHANGE_COUNT_BY",
					"action.0.value":                "-3",
					"display_name":                  "scale in rule - updated",
					"metric.#":                      "1",
					"metric.0.metric_type":          "CPU_UTILIZATION",
					"metric.0.threshold.#":          "1",
					"metric.0.threshold.0.operator": "LT",
					"metric.0.threshold.0.value":    "3",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "policies.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + AutoScalingConfigurationResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckAutoScalingAutoScalingConfigurationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).AutoScalingClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_autoscaling_auto_scaling_configuration" {
			noResourceFound = false
			request := oci_auto_scaling.GetAutoScalingConfigurationRequest{}

			tmp := rs.Primary.ID
			request.AutoScalingConfigurationId = &tmp

			_, err := client.GetAutoScalingConfiguration(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("AutoScalingAutoScalingConfiguration") {
		resource.AddTestSweepers("AutoScalingAutoScalingConfiguration", &resource.Sweeper{
			Name:         "AutoScalingAutoScalingConfiguration",
			Dependencies: acctest.DependencyGraph["autoScalingConfiguration"],
			F:            sweepAutoScalingAutoScalingConfigurationResource,
		})
	}
}

func sweepAutoScalingAutoScalingConfigurationResource(compartment string) error {
	autoScalingClient := acctest.GetTestClients(&schema.ResourceData{}).AutoScalingClient()
	autoScalingConfigurationIds, err := getAutoScalingConfigurationIds(compartment)
	if err != nil {
		return err
	}
	for _, autoScalingConfigurationId := range autoScalingConfigurationIds {
		if ok := acctest.SweeperDefaultResourceId[autoScalingConfigurationId]; !ok {
			deleteAutoScalingConfigurationRequest := oci_auto_scaling.DeleteAutoScalingConfigurationRequest{}

			deleteAutoScalingConfigurationRequest.AutoScalingConfigurationId = &autoScalingConfigurationId

			deleteAutoScalingConfigurationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "auto_scaling")
			_, error := autoScalingClient.DeleteAutoScalingConfiguration(context.Background(), deleteAutoScalingConfigurationRequest)
			if error != nil {
				fmt.Printf("Error deleting AutoScalingConfiguration %s %s, It is possible that the resource is already deleted. Please verify manually \n", autoScalingConfigurationId, error)
				continue
			}
		}
	}
	return nil
}

func getAutoScalingConfigurationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AutoScalingConfigurationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	autoScalingClient := acctest.GetTestClients(&schema.ResourceData{}).AutoScalingClient()

	listAutoScalingConfigurationsRequest := oci_auto_scaling.ListAutoScalingConfigurationsRequest{}
	listAutoScalingConfigurationsRequest.CompartmentId = &compartmentId
	listAutoScalingConfigurationsResponse, err := autoScalingClient.ListAutoScalingConfigurations(context.Background(), listAutoScalingConfigurationsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting AutoScalingConfiguration list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, autoScalingConfiguration := range listAutoScalingConfigurationsResponse.Items {
		id := *autoScalingConfiguration.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AutoScalingConfigurationId", id)
	}
	return resourceIds, nil
}
