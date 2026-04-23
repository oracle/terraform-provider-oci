// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatascienceComputeTargetRequiredOnlyResource = DatascienceComputeTargetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_compute_target", "test_compute_target", acctest.Required, acctest.Create, DatascienceComputeTargetRepresentation)

	DatascienceComputeTargetResourceConfig = DatascienceComputeTargetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_compute_target", "test_compute_target", acctest.Optional, acctest.Update, DatascienceComputeTargetRepresentation)

	DatascienceComputeTargetSingularDataSourceRepresentation = map[string]interface{}{
		"compute_target_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_compute_target.test_compute_target.id}`},
	}

	DatascienceComputeTargetDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_compute_target.test_compute_target.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceComputeTargetDataSourceFilterRepresentation}}
	DatascienceComputeTargetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_datascience_compute_target.test_compute_target.id}`}},
	}

	DatascienceComputeTargetRepresentation = map[string]interface{}{
		"compartment_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compute_configuration_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceComputeTargetComputeConfigurationDetailsRepresentation},
		"defined_tags":                  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                   acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                 acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"metadata":                      acctest.Representation{RepType: acctest.Required, Create: map[string]string{"skipImageVerification": "true"}, Update: map[string]string{"skipImageVerification": "true"}},
	}
	DatascienceComputeTargetComputeConfigurationDetailsRepresentation = map[string]interface{}{
		"compute_type":           acctest.Representation{RepType: acctest.Required, Create: `MANAGED_COMPUTE_CLUSTER`},
		"instance_configuration": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceComputeTargetComputeConfigurationDetailsInstanceConfigurationRepresentation},
		"scaling_policy":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatascienceComputeTargetComputeConfigurationDetailsScalingPolicyRepresentation},
	}
	DatascienceComputeTargetComputeConfigurationDetailsInstanceConfigurationRepresentation = map[string]interface{}{
		"instance_shape":          acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E4.Flex`, Update: `VM.Standard.E4.Flex`},
		"boot_volume_size_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `100`, Update: `100`},
		"capacity_reservation_id": acctest.Representation{RepType: acctest.Optional, Create: nil, Update: nil},
		"instance_shape_details":  acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceComputeTargetComputeConfigurationDetailsInstanceConfigurationInstanceShapeDetailsRepresentation},
	}
	DatascienceComputeTargetComputeConfigurationDetailsScalingPolicyRepresentation = map[string]interface{}{
		"policy_type":           acctest.Representation{RepType: acctest.Required, Create: `AUTOSCALING`, Update: `AUTOSCALING`},
		"auto_scaling_policies": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatascienceComputeTargetComputeConfigurationDetailsScalingPolicyAutoScalingPoliciesRepresentation},
		"cool_down_in_seconds":  acctest.Representation{RepType: acctest.Optional, Create: `600`, Update: `600`},
		//"instance_count":        acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `1`},
		"is_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
	}
	DatascienceComputeTargetComputeConfigurationDetailsInstanceConfigurationInstanceShapeDetailsRepresentation = map[string]interface{}{
		"memory_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `10`},
		"ocpus":         acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `10`},
	}
	DatascienceComputeTargetComputeConfigurationDetailsScalingPolicyAutoScalingPoliciesRepresentation = map[string]interface{}{
		"auto_scaling_policy_type": acctest.Representation{RepType: acctest.Required, Create: `THRESHOLD`},
		"initial_instance_count":   acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `1`},
		"maximum_instance_count":   acctest.Representation{RepType: acctest.Required, Create: `2`, Update: `2`},
		"minimum_instance_count":   acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `1`},
		"rules":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceComputeTargetComputeConfigurationDetailsScalingPolicyAutoScalingPoliciesRulesRepresentation},
	}
	DatascienceComputeTargetComputeConfigurationDetailsScalingPolicyAutoScalingPoliciesRulesRepresentation = map[string]interface{}{
		"metric_expression_rule_type": acctest.Representation{RepType: acctest.Required, Create: `CUSTOM_EXPRESSION`, Update: `CUSTOM_EXPRESSION`},
		"scale_in_configuration":      acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceComputeTargetComputeConfigurationDetailsScalingPolicyAutoScalingPoliciesRulesScaleInConfigurationRepresentation},
		"scale_out_configuration":     acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceComputeTargetComputeConfigurationDetailsScalingPolicyAutoScalingPoliciesRulesScaleOutConfigurationRepresentation},
		//"metric_type":                 acctest.Representation{RepType: acctest.Optional, Create: `CPU_UTILIZATION`, Update: `MEMORY_UTILIZATION`},
	}
	DatascienceComputeTargetComputeConfigurationDetailsScalingPolicyAutoScalingPoliciesRulesScaleInConfigurationRepresentation = map[string]interface{}{
		"instance_count_adjustment":  acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `1`},
		"pending_duration":           acctest.Representation{RepType: acctest.Optional, Create: `PT3M`, Update: `PT3M`},
		"query":                      acctest.Representation{RepType: acctest.Optional, Create: `ComputeTargetCPUUtilization[1m]{resourceId = \"COMPUTE_TARGET_OCID\"}.mean() < 1`, Update: `ComputeTargetCPUUtilization[1m]{resourceId = \"COMPUTE_TARGET_OCID\"}.mean() < 1`},
		"scaling_configuration_type": acctest.Representation{RepType: acctest.Optional, Create: `QUERY`, Update: `QUERY`},
		//"threshold":                  acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `1`},
	}
	DatascienceComputeTargetComputeConfigurationDetailsScalingPolicyAutoScalingPoliciesRulesScaleOutConfigurationRepresentation = map[string]interface{}{
		"instance_count_adjustment":  acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `1`},
		"pending_duration":           acctest.Representation{RepType: acctest.Optional, Create: `PT3M`, Update: `PT3M`},
		"query":                      acctest.Representation{RepType: acctest.Optional, Create: `ComputeTargetCPUUtilization[1m]{resourceId = \"COMPUTE_TARGET_OCID\"}.mean() > 94`, Update: `ComputeTargetCPUUtilization[1m]{resourceId = \"COMPUTE_TARGET_OCID\"}.mean() > 94`},
		"scaling_configuration_type": acctest.Representation{RepType: acctest.Optional, Create: `QUERY`, Update: `QUERY`},
		//"threshold":                  acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}

	DatascienceComputeTargetResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: datascience/default
func TestDatascienceComputeTargetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceComputeTargetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_datascience_compute_target.test_compute_target"
	datasourceName := "data.oci_datascience_compute_targets.test_compute_targets"
	singularDatasourceName := "data.oci_datascience_compute_target.test_compute_target"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatascienceComputeTargetResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_compute_target", "test_compute_target", acctest.Optional, acctest.Create, DatascienceComputeTargetRepresentation), "datascience", "computeTarget", t)

	acctest.ResourceTest(t, testAccCheckDatascienceComputeTargetDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatascienceComputeTargetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_compute_target", "test_compute_target", acctest.Required, acctest.Create, DatascienceComputeTargetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.compute_type", "MANAGED_COMPUTE_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.instance_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.instance_configuration.0.instance_shape", "VM.Standard.E4.Flex"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatascienceComputeTargetResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatascienceComputeTargetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_compute_target", "test_compute_target", acctest.Optional, acctest.Create, DatascienceComputeTargetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.compute_type", "MANAGED_COMPUTE_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.instance_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.instance_configuration.0.boot_volume_size_in_gbs", "100"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.instance_configuration.0.instance_shape", "VM.Standard.E4.Flex"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.instance_configuration.0.instance_shape_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.instance_configuration.0.instance_shape_details.0.memory_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.instance_configuration.0.instance_shape_details.0.ocpus", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.auto_scaling_policy_type", "THRESHOLD"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.initial_instance_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.maximum_instance_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.minimum_instance_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.metric_expression_rule_type", "CUSTOM_EXPRESSION"),
				//resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.metric_type", "CPU_UTILIZATION"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_in_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_in_configuration.0.instance_count_adjustment", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_in_configuration.0.pending_duration", "PT3M"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_in_configuration.0.query", "ComputeTargetCPUUtilization[1m]{resourceId = \"COMPUTE_TARGET_OCID\"}.mean() < 1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_in_configuration.0.scaling_configuration_type", "QUERY"),
				//resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_in_configuration.0.threshold", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_out_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_out_configuration.0.instance_count_adjustment", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_out_configuration.0.pending_duration", "PT3M"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_out_configuration.0.query", "ComputeTargetCPUUtilization[1m]{resourceId = \"COMPUTE_TARGET_OCID\"}.mean() > 94"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_out_configuration.0.scaling_configuration_type", "QUERY"),
				//resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_out_configuration.0.threshold", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.cool_down_in_seconds", "600"),
				//resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.instance_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.is_enabled", "false"),
				//resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.policy_type", "FIXED_SIZE"),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatascienceComputeTargetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_compute_target", "test_compute_target", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatascienceComputeTargetRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.compute_type", "MANAGED_COMPUTE_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.instance_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.instance_configuration.0.boot_volume_size_in_gbs", "100"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.instance_configuration.0.instance_shape", "VM.Standard.E4.Flex"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.instance_configuration.0.instance_shape_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.instance_configuration.0.instance_shape_details.0.memory_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.instance_configuration.0.instance_shape_details.0.ocpus", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.auto_scaling_policy_type", "THRESHOLD"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.initial_instance_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.maximum_instance_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.minimum_instance_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.metric_expression_rule_type", "CUSTOM_EXPRESSION"),
				//resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.metric_type", "CPU_UTILIZATION"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_in_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_in_configuration.0.instance_count_adjustment", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_in_configuration.0.pending_duration", "PT3M"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_in_configuration.0.query", "ComputeTargetCPUUtilization[1m]{resourceId = \"COMPUTE_TARGET_OCID\"}.mean() < 1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_in_configuration.0.scaling_configuration_type", "QUERY"),
				//resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_in_configuration.0.threshold", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_out_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_out_configuration.0.instance_count_adjustment", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_out_configuration.0.pending_duration", "PT3M"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_out_configuration.0.query", "ComputeTargetCPUUtilization[1m]{resourceId = \"COMPUTE_TARGET_OCID\"}.mean() > 94"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_out_configuration.0.scaling_configuration_type", "QUERY"),
				//resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_out_configuration.0.threshold", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.cool_down_in_seconds", "600"),
				//resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.instance_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.is_enabled", "false"),
				//resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.policy_type", "FIXED_SIZE"),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DatascienceComputeTargetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_compute_target", "test_compute_target", acctest.Optional, acctest.Update, DatascienceComputeTargetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.compute_type", "MANAGED_COMPUTE_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.instance_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.instance_configuration.0.boot_volume_size_in_gbs", "100"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.instance_configuration.0.instance_shape", "VM.Standard.E4.Flex"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.instance_configuration.0.instance_shape_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.instance_configuration.0.instance_shape_details.0.memory_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.instance_configuration.0.instance_shape_details.0.ocpus", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.auto_scaling_policy_type", "THRESHOLD"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.initial_instance_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.maximum_instance_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.minimum_instance_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.metric_expression_rule_type", "CUSTOM_EXPRESSION"),
				//resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.metric_type", "MEMORY_UTILIZATION"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_in_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_in_configuration.0.instance_count_adjustment", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_in_configuration.0.pending_duration", "PT3M"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_in_configuration.0.query", "ComputeTargetCPUUtilization[1m]{resourceId = \"COMPUTE_TARGET_OCID\"}.mean() < 1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_in_configuration.0.scaling_configuration_type", "QUERY"),
				//resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_in_configuration.0.threshold", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_out_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_out_configuration.0.instance_count_adjustment", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_out_configuration.0.pending_duration", "PT3M"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_out_configuration.0.query", "ComputeTargetCPUUtilization[1m]{resourceId = \"COMPUTE_TARGET_OCID\"}.mean() > 94"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_out_configuration.0.scaling_configuration_type", "QUERY"),
				//resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_out_configuration.0.threshold", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.cool_down_in_seconds", "600"),
				//resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.instance_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.is_enabled", "false"),
				//resource.TestCheckResourceAttr(resourceName, "compute_configuration_details.0.scaling_policy.0.policy_type", "AUTOSCALING"),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_compute_targets", "test_compute_targets", acctest.Optional, acctest.Update, DatascienceComputeTargetDataSourceRepresentation) +
				compartmentIdVariableStr + DatascienceComputeTargetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_compute_target", "test_compute_target", acctest.Optional, acctest.Update, DatascienceComputeTargetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compute_targets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "compute_targets.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_targets.0.created_by"),
				resource.TestCheckResourceAttr(datasourceName, "compute_targets.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "compute_targets.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_targets.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_targets.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_targets.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_compute_target", "test_compute_target", acctest.Required, acctest.Create, DatascienceComputeTargetSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatascienceComputeTargetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_target_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_configuration_details.0.compute_type", "MANAGED_COMPUTE_CLUSTER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_configuration_details.0.instance_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_configuration_details.0.instance_configuration.0.boot_volume_size_in_gbs", "100"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_configuration_details.0.instance_configuration.0.instance_shape", "VM.Standard.E4.Flex"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_configuration_details.0.instance_configuration.0.instance_shape_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_configuration_details.0.instance_configuration.0.instance_shape_details.0.memory_in_gbs", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_configuration_details.0.instance_configuration.0.instance_shape_details.0.ocpus", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.auto_scaling_policy_type", "THRESHOLD"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.initial_instance_count", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.maximum_instance_count", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.minimum_instance_count", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.metric_expression_rule_type", "CUSTOM_EXPRESSION"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.metric_type", "MEMORY_UTILIZATION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_in_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_in_configuration.0.instance_count_adjustment", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_in_configuration.0.pending_duration", "PT3M"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_in_configuration.0.query", "ComputeTargetCPUUtilization[1m]{resourceId = \"COMPUTE_TARGET_OCID\"}.mean() < 1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_in_configuration.0.scaling_configuration_type", "QUERY"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_in_configuration.0.threshold", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_out_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_out_configuration.0.instance_count_adjustment", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_out_configuration.0.pending_duration", "PT3M"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_out_configuration.0.query", "ComputeTargetCPUUtilization[1m]{resourceId = \"COMPUTE_TARGET_OCID\"}.mean() > 94"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_out_configuration.0.scaling_configuration_type", "QUERY"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "compute_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_out_configuration.0.threshold", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_configuration_details.0.scaling_policy.0.cool_down_in_seconds", "600"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "compute_configuration_details.0.scaling_policy.0.instance_count", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_configuration_details.0.scaling_policy.0.is_enabled", "false"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "compute_configuration_details.0.scaling_policy.0.policy_type", "AUTOSCALING"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "compute_target_system_data.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + DatascienceComputeTargetRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatascienceComputeTargetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataScienceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datascience_compute_target" {
			noResourceFound = false
			request := oci_datascience.GetComputeTargetRequest{}

			tmp := rs.Primary.ID
			request.ComputeTargetId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")

			response, err := client.GetComputeTarget(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_datascience.ComputeTargetLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
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
	if !acctest.InSweeperExcludeList("DatascienceComputeTarget") {
		resource.AddTestSweepers("DatascienceComputeTarget", &resource.Sweeper{
			Name:         "DatascienceComputeTarget",
			Dependencies: acctest.DependencyGraph["computeTarget"],
			F:            sweepDatascienceComputeTargetResource,
		})
	}
}

func sweepDatascienceComputeTargetResource(compartment string) error {
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()
	computeTargetIds, err := getDatascienceComputeTargetIds(compartment)
	if err != nil {
		return err
	}
	for _, computeTargetId := range computeTargetIds {
		if ok := acctest.SweeperDefaultResourceId[computeTargetId]; !ok {
			deleteComputeTargetRequest := oci_datascience.DeleteComputeTargetRequest{}

			deleteComputeTargetRequest.ComputeTargetId = &computeTargetId

			deleteComputeTargetRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")
			_, error := dataScienceClient.DeleteComputeTarget(context.Background(), deleteComputeTargetRequest)
			if error != nil {
				fmt.Printf("Error deleting ComputeTarget %s %s, It is possible that the resource is already deleted. Please verify manually \n", computeTargetId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &computeTargetId, DatascienceComputeTargetSweepWaitCondition, time.Duration(3*time.Minute),
				DatascienceComputeTargetSweepResponseFetchOperation, "datascience", true)
		}
	}
	return nil
}

func getDatascienceComputeTargetIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ComputeTargetId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()

	listComputeTargetsRequest := oci_datascience.ListComputeTargetsRequest{}
	listComputeTargetsRequest.CompartmentId = &compartmentId
	listComputeTargetsRequest.LifecycleState = oci_datascience.ListComputeTargetsLifecycleStateActive
	listComputeTargetsResponse, err := dataScienceClient.ListComputeTargets(context.Background(), listComputeTargetsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ComputeTarget list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, computeTarget := range listComputeTargetsResponse.Items {
		id := *computeTarget.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ComputeTargetId", id)
	}
	return resourceIds, nil
}

func DatascienceComputeTargetSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if computeTargetResponse, ok := response.Response.(oci_datascience.GetComputeTargetResponse); ok {
		return computeTargetResponse.LifecycleState != oci_datascience.ComputeTargetLifecycleStateDeleted
	}
	return false
}

func DatascienceComputeTargetSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataScienceClient().GetComputeTarget(context.Background(), oci_datascience.GetComputeTargetRequest{
		ComputeTargetId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
