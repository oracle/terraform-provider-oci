// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatascienceModelDeploymentRequiredOnlyResource = DatascienceModelDeploymentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Required, acctest.Create, DatascienceModelDeploymentRepresentation)

	DatascienceModelDeploymentResourceConfig = DatascienceModelDeploymentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Optional, acctest.Update, DatascienceModelDeploymentRepresentation)

	DatascienceDatascienceModelDeploymentSingularDataSourceRepresentation = map[string]interface{}{
		"model_deployment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_model_deployment.test_model_deployment.id}`},
	}

	modelForModelDeploymentRepresentation = map[string]interface{}{
		"artifact_content_length":      acctest.Representation{RepType: acctest.Required, Create: `7532`},
		"model_artifact":               acctest.Representation{RepType: acctest.Required, Create: `../../examples/datascience/model_deployment/artifact.zip`},
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"project_id":                   acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.test_project.id}`},
		"artifact_content_disposition": acctest.Representation{RepType: acctest.Optional, Create: `attachment; filename=tfTestArtifact.zip`},
	}

	modelForUpdateModelDeploymentRepresentation = map[string]interface{}{
		"artifact_content_length":      acctest.Representation{RepType: acctest.Required, Create: `7532`},
		"model_artifact":               acctest.Representation{RepType: acctest.Required, Create: `../../examples/datascience/model_deployment/artifact.zip`},
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"project_id":                   acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.test_project.id}`},
		"artifact_content_disposition": acctest.Representation{RepType: acctest.Optional, Create: `attachment; filename=tfTestArtifact.zip`},
	}

	DatascienceDatascienceModelDeploymentDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"created_by":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_model_deployment.test_model_deployment.created_by}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_model_deployment.test_model_deployment.id}`},
		"project_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.test_project.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceModelDeploymentDataSourceFilterRepresentation}}

	DatascienceModelDeploymentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_datascience_model_deployment.test_model_deployment.id}`}},
	}

	DatascienceModelDeploymentRepresentation = map[string]interface{}{
		"compartment_id":                         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"model_deployment_configuration_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceModelDeploymentModelDeploymentConfigurationDetailsRepresentation},
		"project_id":                             acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.test_project.id}`},
		"category_log_details":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatascienceModelDeploymentCategoryLogDetailsRepresentation},
		"defined_tags":                           acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                            acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                           acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                              acctest.RepresentationGroup{RepType: acctest.Optional, Group: definedTagsIgnoreMDRepresentation},
	}

	definedTagsIgnoreMDRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	DatascienceModelDeploymentModelDeploymentConfigurationDetailsRepresentation = map[string]interface{}{
		"deployment_type":             acctest.Representation{RepType: acctest.Required, Create: `SINGLE_MODEL`},
		"model_configuration_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceModelDeploymentModelDeploymentConfigurationDetailsModelConfigurationDetailsRepresentation},
	}
	DatascienceModelDeploymentCategoryLogDetailsRepresentation = map[string]interface{}{
		"access":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatascienceModelDeploymentCategoryLogDetailsAccessRepresentation},
		"predict": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatascienceModelDeploymentCategoryLogDetailsPredictRepresentation},
	}
	DatascienceModelDeploymentModelDeploymentConfigurationDetailsModelConfigurationDetailsRepresentation = map[string]interface{}{
		"instance_configuration": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceModelDeploymentModelDeploymentConfigurationDetailsModelConfigurationDetailsInstanceConfigurationRepresentation},
		"model_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_model.test_model.id}`, Update: `${oci_datascience_model.test_model_update.id}`},
		"bandwidth_mbps":         acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"maximum_bandwidth_mbps": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"scaling_policy":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatascienceModelDeploymentModelDeploymentConfigurationDetailsModelConfigurationDetailsScalingPolicyRepresentation},
	}
	DatascienceModelDeploymentCategoryLogDetailsAccessRepresentation = map[string]interface{}{
		"log_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log_group.test_log_group.id}`},
		"log_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log.test_access_log.id}`},
	}
	DatascienceModelDeploymentCategoryLogDetailsPredictRepresentation = map[string]interface{}{
		"log_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log_group.test_log_group.id}`},
		"log_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log.test_predict_log.id}`},
	}
	DatascienceModelDeploymentModelDeploymentConfigurationDetailsModelConfigurationDetailsInstanceConfigurationRepresentation = map[string]interface{}{
		"subnet_id":           acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"instance_shape_name": acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E4.Flex`, Update: `VM.Standard.E3.Flex`},
		"model_deployment_instance_shape_config_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceModelDeploymentModelDeploymentConfigurationDetailsModelConfigurationDetailsInstanceConfigurationModelDeploymentInstanceShapeConfigDetailsRepresentation},
	}
	DatascienceModelDeploymentModelDeploymentConfigurationDetailsModelConfigurationDetailsScalingPolicyRepresentation = map[string]interface{}{
		"policy_type":    acctest.Representation{RepType: acctest.Required, Create: `FIXED_SIZE`},
		"instance_count": acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
	}
	DatascienceModelDeploymentModelDeploymentConfigurationDetailsModelConfigurationDetailsInstanceConfigurationModelDeploymentInstanceShapeConfigDetailsRepresentation = map[string]interface{}{
		"cpu_baseline":  acctest.Representation{RepType: acctest.Optional, Create: `BASELINE_1_8`, Update: `BASELINE_1_2`},
		"ocpus":         acctest.Representation{RepType: acctest.Required, Create: `1.0`, Update: `2.0`},
		"memory_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `6.0`, Update: `12.0`},
	}
	logGroupMDRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `tf_testing_log_group`, Update: `tf_testing_log_group_update`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	logGroupUpdateMDRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `tf_update_testing_log_group`, Update: `tf_update_testing_log_group_update`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	customAccessLogRepresentation = map[string]interface{}{
		"display_name":       acctest.Representation{RepType: acctest.Required, Create: `tf-testing-access-log`, Update: `tf-testing-Update-access-log`},
		"log_group_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log_group.test_log_group.id}`, Update: `${oci_logging_log_group.test_update_log_group.id}`},
		"log_type":           acctest.Representation{RepType: acctest.Required, Create: `CUSTOM`},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_enabled":         acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"retention_duration": acctest.Representation{RepType: acctest.Optional, Create: `30`, Update: `60`},
	}

	customPredictLogRepresentation = map[string]interface{}{
		"display_name":       acctest.Representation{RepType: acctest.Required, Create: `tf-testing-predict-log`, Update: `tf-testing-Update-predict-log`},
		"log_group_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log_group.test_log_group.id}`, Update: `${oci_logging_log_group.test_update_log_group.id}`},
		"log_type":           acctest.Representation{RepType: acctest.Required, Create: `CUSTOM`},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_enabled":         acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"retention_duration": acctest.Representation{RepType: acctest.Optional, Create: `30`, Update: `60`},
	}

	ModelDeploymentCoreSubnetRepresentation = map[string]interface{}{
		"cidr_block":                 acctest.Representation{RepType: acctest.Required, Create: `10.0.0.0/16`},
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"availability_domain":        acctest.Representation{RepType: acctest.Optional, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"dhcp_options_id":            acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn.default_dhcp_options_id}`, Update: `${oci_core_dhcp_options.test_dhcp_options.id}`},
		"display_name":               acctest.Representation{RepType: acctest.Optional, Create: `subnetDisplayName`, Update: `subnetDisplayName2`},
		"dns_label":                  acctest.Representation{RepType: acctest.Optional, Create: `dnslabel`},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"prohibit_public_ip_on_vnic": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"prohibit_internet_ingress":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"route_table_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn.default_route_table_id}`, Update: `${oci_core_route_table.test_route_table.id}`},
		"security_list_ids":          acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_vcn.test_vcn.default_security_list_id}`}, Update: []string{`${oci_core_security_list.test_security_list.id}`}},
		"lifecycle":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRep},
	}

	DatascienceModelDeploymentResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model", acctest.Optional, acctest.Create, modelForModelDeploymentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model_update", acctest.Optional, acctest.Create, modelForUpdateModelDeploymentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, ModelDeploymentCoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", acctest.Required, acctest.Create, DatascienceProjectRepresentation) +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", acctest.Required, acctest.Create, logGroupMDRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_access_log", acctest.Required, acctest.Create, customAccessLogRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_predict_log", acctest.Required, acctest.Create, customPredictLogRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_update_log_group", acctest.Required, acctest.Create, logGroupUpdateMDRepresentation)

	DatascienceMCCModelDeploymentRepresentation = map[string]interface{}{
		"compartment_id":                         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"model_deployment_configuration_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceMCCModelDeploymentModelDeploymentConfigurationDetailsRepresentation},
		"project_id":                             acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.test_project.id}`},
		"defined_tags":                           acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                            acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                           acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                              acctest.RepresentationGroup{RepType: acctest.Optional, Group: definedTagsIgnoreMDRepresentation},
	}

	DatascienceMCCModelDeploymentModelDeploymentConfigurationDetailsRepresentation = map[string]interface{}{
		"deployment_type":                      acctest.Representation{RepType: acctest.Required, Create: `SINGLE_MODEL_FLEX`},
		"model_configuration_details":          acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceMCCModelDeploymentModelDeploymentConfigurationDetailsModelConfigurationDetailsRepresentation},
		"infrastructure_configuration_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceMCCModelDeploymentModelDeploymentConfigurationDetailsInfrastructureConfigurationDetailsRepresentation},
		"environment_configuration_details":    acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceMCCModelDeploymentModelDeploymentConfigurationDetailsEnvironmentConfigurationDetailsRepresentation},
	}

	DatascienceMCCModelDeploymentModelDeploymentConfigurationDetailsInfrastructureConfigurationDetailsRepresentation = map[string]interface{}{
		"compute_target_id":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_compute_target.test_compute_target.id}`},
		"infrastructure_type":                     acctest.Representation{RepType: acctest.Required, Create: `MANAGED_COMPUTE_CLUSTER`},
		"model_deployment_resource_configuration": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceMCCModelDeploymentInfrastructureConfigurationDetailsModelDeploymentResourceConfigurationRepresentation},
		"scaling_policy":                          acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceMCCModelDeploymentInfrastructureConfigurationDetailsScalingPolicyRepresentation},
	}

	DatascienceMCCModelDeploymentModelDeploymentConfigurationDetailsModelConfigurationDetailsRepresentation = map[string]interface{}{
		"model_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_model.test_model.id}`},
	}

	DatascienceMCCModelDeploymentModelDeploymentConfigurationDetailsEnvironmentConfigurationDetailsRepresentation = map[string]interface{}{
		"environment_configuration_type": acctest.Representation{RepType: acctest.Required, Create: `OCIR_CONTAINER`},
		"image":                          acctest.Representation{RepType: acctest.Required, Create: `iad.ocir.io/ociodscdev/sample-byoc:cpu`},
		"image_digest":                   acctest.Representation{RepType: acctest.Required, Create: `sha256:2b90a418013aae4422177e24c6fd2269931272efc4fda3c54e82ab1da60219af`},
		"server_port":                    acctest.Representation{RepType: acctest.Required, Create: `8080`},
		"health_check_port":              acctest.Representation{RepType: acctest.Required, Create: `8080`},
	}

	DatascienceMCCModelDeploymentInfrastructureConfigurationDetailsModelDeploymentResourceConfigurationRepresentation = map[string]interface{}{
		"resource_request_configuration": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceMCCModelDeploymentInfrastructureConfigurationDetailsModelDeploymentResourceConfigurationResourceRequestConfigurationRepresentation},
		"resource_limit_configuration":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatascienceMCCModelDeploymentInfrastructureConfigurationDetailsModelDeploymentResourceConfigurationResourceLimitConfigurationRepresentation},
	}

	DatascienceMCCModelDeploymentInfrastructureConfigurationDetailsModelDeploymentResourceConfigurationResourceRequestConfigurationRepresentation = map[string]interface{}{
		"memory_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `6`, Update: `6`},
		"ocpus":         acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `1`},
	}

	DatascienceMCCModelDeploymentInfrastructureConfigurationDetailsModelDeploymentResourceConfigurationResourceLimitConfigurationRepresentation = map[string]interface{}{
		"memory_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `6`, Update: `6`},
		"ocpus":         acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `1`},
	}

	DatascienceMCCModelDeploymentInfrastructureConfigurationDetailsScalingPolicyRepresentation = map[string]interface{}{
		"policy_type":    acctest.Representation{RepType: acctest.Required, Create: `FIXED_SIZE`},
		"instance_count": acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `2`},
	}

	DatascienceMCCModelDeploymentAutoscalingPolicyUpdateRepresentation = map[string]interface{}{
		"policy_type":           acctest.Representation{RepType: acctest.Required, Create: `AUTOSCALING`, Update: `AUTOSCALING`},
		"is_enabled":            acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `true`},
		"auto_scaling_policies": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceMCCModelDeploymentInfrastructureConfigurationDetailsScalingPolicyAutoScalingPoliciesCanaryRepresentation},
	}

	DatascienceMCCModelDeploymentInfrastructureConfigurationDetailsScalingPolicyAutoScalingPoliciesCanaryRepresentation = map[string]interface{}{
		"auto_scaling_policy_type": acctest.Representation{RepType: acctest.Required, Create: `THRESHOLD`, Update: `THRESHOLD`},
		"initial_instance_count":   acctest.Representation{RepType: acctest.Required, Create: `1`, Update: `1`},
		"maximum_instance_count":   acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `10`},
		"minimum_instance_count":   acctest.Representation{RepType: acctest.Required, Create: `0`, Update: `0`},
		"rules":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceMCCModelDeploymentInfrastructureConfigurationDetailsScalingPolicyAutoScalingPoliciesRulesCanaryRepresentation},
		"scale_in_policy":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatascienceMCCModelDeploymentInfrastructureConfigurationDetailsScalingPolicyAutoScalingPoliciesScaleInPolicyRepresentation},
		"scale_out_policy":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatascienceMCCModelDeploymentInfrastructureConfigurationDetailsScalingPolicyAutoScalingPoliciesScaleOutPolicyRepresentation},
	}

	DatascienceMCCModelDeploymentInfrastructureConfigurationDetailsScalingPolicyAutoScalingPoliciesRulesCanaryRepresentation = map[string]interface{}{
		"metric_expression_rule_type": acctest.Representation{RepType: acctest.Required, Create: `TARGET_CUSTOM_EXPRESSION`, Update: `TARGET_CUSTOM_EXPRESSION`},
		"scale_configuration":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceMCCModelDeploymentInfrastructureConfigurationDetailsScalingPolicyAutoScalingPoliciesRulesScaleConfigurationCanaryRepresentation},
	}

	DatascienceMCCModelDeploymentInfrastructureConfigurationDetailsScalingPolicyAutoScalingPoliciesRulesScaleConfigurationCanaryRepresentation = map[string]interface{}{
		"target_scaling_configuration_type": acctest.Representation{RepType: acctest.Required, Create: `QUERY`, Update: `QUERY`},
		"query":                             acctest.Representation{RepType: acctest.Required, Create: `PredictRequestCount[5m]{resourceId = \"MODEL_DEPLOYMENT_OCID\"}.count()`, Update: `PredictRequestCount[5m]{resourceId = \"MODEL_DEPLOYMENT_OCID\"}.count()`},
		"threshold":                         acctest.Representation{RepType: acctest.Required, Create: `2`, Update: `2`},
	}

	DatascienceMCCModelDeploymentInfrastructureConfigurationDetailsScalingPolicyAutoScalingPoliciesScaleInPolicyRepresentation = map[string]interface{}{
		"cool_down_in_seconds":      acctest.Representation{RepType: acctest.Optional, Create: `300`, Update: `300`},
		"instance_count_adjustment": acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `1`},
		"pending_duration":          acctest.Representation{RepType: acctest.Optional, Create: `PT1M`, Update: `PT1M`},
	}

	DatascienceMCCModelDeploymentInfrastructureConfigurationDetailsScalingPolicyAutoScalingPoliciesScaleOutPolicyRepresentation = map[string]interface{}{
		"cool_down_in_seconds":      acctest.Representation{RepType: acctest.Optional, Create: `30`, Update: `30`},
		"instance_count_adjustment": acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `1`},
		"pending_duration":          acctest.Representation{RepType: acctest.Optional, Create: `PT1M`, Update: `PT1M`},
	}

	DatascienceMCCModelDeploymentUpdateToAutoscalingRepresentation = acctest.RepresentationCopyWithNewProperties(DatascienceMCCModelDeploymentRepresentation, map[string]interface{}{
		"model_deployment_configuration_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: acctest.RepresentationCopyWithNewProperties(DatascienceMCCModelDeploymentModelDeploymentConfigurationDetailsRepresentation, map[string]interface{}{
			"infrastructure_configuration_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: acctest.RepresentationCopyWithNewProperties(DatascienceMCCModelDeploymentModelDeploymentConfigurationDetailsInfrastructureConfigurationDetailsRepresentation, map[string]interface{}{
				"model_deployment_resource_configuration": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceMCCModelDeploymentInfrastructureConfigurationDetailsModelDeploymentResourceConfigurationRepresentation},
				"scaling_policy": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceMCCModelDeploymentAutoscalingPolicyUpdateRepresentation},
			})},
		})},
	})

	DatascienceMCCModelDeploymentResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model", acctest.Optional, acctest.Create, modelForModelDeploymentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model_update", acctest.Optional, acctest.Create, modelForUpdateModelDeploymentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_compute_target", "test_compute_target", acctest.Required, acctest.Create, DatascienceComputeTargetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", acctest.Required, acctest.Create, DatascienceProjectRepresentation) +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", acctest.Required, acctest.Create, logGroupMDRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_access_log", acctest.Required, acctest.Create, customAccessLogRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_predict_log", acctest.Required, acctest.Create, customPredictLogRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_update_log_group", acctest.Required, acctest.Create, logGroupUpdateMDRepresentation)

	DatascienceMCCModelDeploymentRequiredOnlyResource = DatascienceMCCModelDeploymentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Required, acctest.Create, DatascienceMCCModelDeploymentRepresentation)

	DatascienceMCCModelDeploymentResourceConfig = DatascienceMCCModelDeploymentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Optional, acctest.Update, DatascienceMCCModelDeploymentRepresentation)
)

// issue-routing-tag: datascience/default
func TestDatascienceModelDeploymentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceModelDeploymentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_datascience_model_deployment.test_model_deployment"
	datasourceName := "data.oci_datascience_model_deployments.test_model_deployments"
	singularDatasourceName := "data.oci_datascience_model_deployment.test_model_deployment"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckDatascienceModelDeploymentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatascienceModelDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Required, acctest.Create, DatascienceModelDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.deployment_type", "SINGLE_MODEL"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.instance_shape_name"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatascienceModelDeploymentResourceDependencies,
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatascienceModelDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Optional, acctest.Create, DatascienceModelDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "category_log_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "category_log_details.0.access.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.access.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.access.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "category_log_details.0.predict.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.predict.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.predict.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.deployment_type", "SINGLE_MODEL"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.bandwidth_mbps", "10"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.instance_shape_name"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.0.cpu_baseline", "BASELINE_1_8"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.maximum_bandwidth_mbps", "10"),

				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.0.memory_in_gbs", "6"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.0.ocpus", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.subnet_id"),

				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.instance_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.policy_type", "FIXED_SIZE"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_url"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
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
			)},

		// verify Deactivate model deployment
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatascienceModelDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatascienceModelDeploymentRepresentation, map[string]interface{}{
						"state": acctest.Representation{RepType: acctest.Optional, Create: `INACTIVE`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "category_log_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "category_log_details.0.access.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.access.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.access.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "category_log_details.0.predict.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.predict.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.predict.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.deployment_type", "SINGLE_MODEL"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.bandwidth_mbps", "10"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.maximum_bandwidth_mbps", "10"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.instance_shape_name"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.0.cpu_baseline", "BASELINE_1_8"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.instance_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.policy_type", "FIXED_SIZE"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_url"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "state", "INACTIVE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify Activate model deployment
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatascienceModelDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatascienceModelDeploymentRepresentation, map[string]interface{}{
						"state": acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "category_log_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "category_log_details.0.access.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.access.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.access.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "category_log_details.0.predict.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.predict.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.predict.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.deployment_type", "SINGLE_MODEL"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.bandwidth_mbps", "10"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.maximum_bandwidth_mbps", "10"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.instance_shape_name"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.0.cpu_baseline", "BASELINE_1_8"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.instance_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.policy_type", "FIXED_SIZE"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_url"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatascienceModelDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatascienceModelDeploymentRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "category_log_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "category_log_details.0.access.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.access.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.access.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "category_log_details.0.predict.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.predict.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.predict.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.deployment_type", "SINGLE_MODEL"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.bandwidth_mbps", "10"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.instance_shape_name"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.0.cpu_baseline", "BASELINE_1_8"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.maximum_bandwidth_mbps", "10"),

				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.0.memory_in_gbs", "6"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.0.ocpus", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.subnet_id"),

				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.instance_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.policy_type", "FIXED_SIZE"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_url"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
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
			Config: config + compartmentIdVariableStr + DatascienceModelDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Optional, acctest.Update, DatascienceModelDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "category_log_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "category_log_details.0.access.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.access.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.access.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "category_log_details.0.predict.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.predict.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "category_log_details.0.predict.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.deployment_type", "SINGLE_MODEL"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.bandwidth_mbps", "10"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.instance_shape_name"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.0.cpu_baseline", "BASELINE_1_2"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.maximum_bandwidth_mbps", "10"),

				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.0.memory_in_gbs", "12"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.0.ocpus", "2"),

				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.subnet_id"),

				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.instance_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.policy_type", "FIXED_SIZE"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_url"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
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

		// verify datasource - list model deployments
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_model_deployments", "test_model_deployments", acctest.Optional, acctest.Update, DatascienceDatascienceModelDeploymentDataSourceRepresentation) +
				compartmentIdVariableStr + DatascienceModelDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Optional, acctest.Update, DatascienceModelDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "project_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.category_log_details.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.category_log_details.0.access.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_deployments.0.category_log_details.0.access.0.log_group_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_deployments.0.category_log_details.0.access.0.log_id"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.category_log_details.0.predict.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_deployments.0.category_log_details.0.predict.0.log_group_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_deployments.0.category_log_details.0.predict.0.log_id"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "model_deployments.0.created_by"),
				//resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_deployments.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.model_deployment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.deployment_type", "SINGLE_MODEL"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.model_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.model_configuration_details.0.bandwidth_mbps", "10"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.instance_shape_name"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.0.cpu_baseline", "BASELINE_1_2"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.model_configuration_details.0.maximum_bandwidth_mbps", "10"),

				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.0.memory_in_gbs", "12"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.0.ocpus", "2"),

				resource.TestCheckResourceAttrSet(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.subnet_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.instance_count", "2"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.policy_type", "FIXED_SIZE"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_deployments.0.model_deployment_url"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_deployments.0.project_id"),
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.0.state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_deployments.0.time_created"),
			),
		},

		// verify singular datasource - get model deployment
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Required, acctest.Create, DatascienceDatascienceModelDeploymentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatascienceModelDeploymentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "model_deployment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "category_log_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "category_log_details.0.access.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "category_log_details.0.predict.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.deployment_type", "SINGLE_MODEL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.model_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.model_configuration_details.0.bandwidth_mbps", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.model_configuration_details.0.instance_configuration.0.model_deployment_instance_shape_config_details.0.cpu_baseline", "BASELINE_1_2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.model_configuration_details.0.maximum_bandwidth_mbps", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.instance_count", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_deployment_configuration_details.0.model_configuration_details.0.scaling_policy.0.policy_type", "FIXED_SIZE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "model_deployment_url"),
				resource.TestCheckNoResourceAttr(singularDatasourceName, "model_deployment_system_data.#"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},

		// verify resource import
		{
			Config:            config + DatascienceModelDeploymentRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"description",
			},
			ResourceName: resourceName,
		},
	})
}

func TestDatascienceModelDeploymentResource_single_model_flex_only(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceModelDeploymentResource_single_model_flex_only")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_datascience_model_deployment.test_model_deployment"
	datasourceName := "data.oci_datascience_model_deployments.test_model_deployments"
	singularDatasourceName := "data.oci_datascience_model_deployment.test_model_deployment"

	runDatascienceModelDeploymentMCCLifecycleTest(t, config, compartmentIdVariableStr, compartmentIdUVariableStr, resourceName, datasourceName, singularDatasourceName, compartmentId, compartmentIdU)
}

func runDatascienceModelDeploymentMCCLifecycleTest(t *testing.T, config, compartmentIdVariableStr, compartmentIdUVariableStr, resourceName, datasourceName, singularDatasourceName, compartmentId, compartmentIdU string) {
	var mccResId, mccResId2 string

	acctest.ResourceTest(t, testAccCheckDatascienceModelDeploymentDestroy, []resource.TestStep{
		// ------------------------- MCC Block: Create (Required) -------------------------
		{
			Config: config + compartmentIdVariableStr + DatascienceMCCModelDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Required, acctest.Create, DatascienceMCCModelDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.deployment_type", "SINGLE_MODEL_FLEX"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.compute_target_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.infrastructure_type", "MANAGED_COMPUTE_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.environment_configuration_type", "OCIR_CONTAINER"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.image", "iad.ocir.io/ociodscdev/sample-byoc:cpu"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.image_digest", "sha256:2b90a418013aae4422177e24c6fd2269931272efc4fda3c54e82ab1da60219af"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.server_port", "8080"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.health_check_port", "8080"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.0.resource_request_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.0.resource_request_configuration.0.memory_in_gbs", "6"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.0.resource_request_configuration.0.ocpus", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.policy_type", "FIXED_SIZE"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.instance_count", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_url"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				func(s *terraform.State) (err error) {
					mccResId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatascienceMCCModelDeploymentResourceDependencies,
		},

		// ------------------------- MCC Block: Create (Optional) -------------------------
		{
			Config: config + compartmentIdVariableStr + DatascienceMCCModelDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Optional, acctest.Create, DatascienceMCCModelDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.deployment_type", "SINGLE_MODEL_FLEX"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.compute_target_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.infrastructure_type", "MANAGED_COMPUTE_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.environment_configuration_type", "OCIR_CONTAINER"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.image", "iad.ocir.io/ociodscdev/sample-byoc:cpu"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.image_digest", "sha256:2b90a418013aae4422177e24c6fd2269931272efc4fda3c54e82ab1da60219af"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.server_port", "8080"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.health_check_port", "8080"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.0.resource_request_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.0.resource_request_configuration.0.memory_in_gbs", "6"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.0.resource_request_configuration.0.ocpus", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.policy_type", "FIXED_SIZE"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.instance_count", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_url"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				func(s *terraform.State) (err error) {
					mccResId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatascienceMCCModelDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatascienceMCCModelDeploymentRepresentation, map[string]interface{}{
						"state": acctest.Representation{RepType: acctest.Optional, Create: `INACTIVE`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "state", "INACTIVE"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.deployment_type", "SINGLE_MODEL_FLEX"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.compute_target_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.infrastructure_type", "MANAGED_COMPUTE_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.environment_configuration_type", "OCIR_CONTAINER"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.image", "iad.ocir.io/ociodscdev/sample-byoc:cpu"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.image_digest", "sha256:2b90a418013aae4422177e24c6fd2269931272efc4fda3c54e82ab1da60219af"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.server_port", "8080"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.health_check_port", "8080"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.0.resource_request_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.0.resource_request_configuration.0.memory_in_gbs", "6"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.0.resource_request_configuration.0.ocpus", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.policy_type", "FIXED_SIZE"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.instance_count", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_url"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				func(s *terraform.State) (err error) {
					mccResId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if mccResId != mccResId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatascienceMCCModelDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatascienceMCCModelDeploymentRepresentation, map[string]interface{}{
						"state": acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.deployment_type", "SINGLE_MODEL_FLEX"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.compute_target_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.infrastructure_type", "MANAGED_COMPUTE_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.environment_configuration_type", "OCIR_CONTAINER"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.image", "iad.ocir.io/ociodscdev/sample-byoc:cpu"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.image_digest", "sha256:2b90a418013aae4422177e24c6fd2269931272efc4fda3c54e82ab1da60219af"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.server_port", "8080"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.health_check_port", "8080"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.0.resource_request_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.0.resource_request_configuration.0.memory_in_gbs", "6"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.0.resource_request_configuration.0.ocpus", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.0.resource_limit_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.0.resource_limit_configuration.0.memory_in_gbs", "6"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.0.resource_limit_configuration.0.ocpus", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.policy_type", "FIXED_SIZE"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.instance_count", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_url"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
			),
		},

		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatascienceMCCModelDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatascienceMCCModelDeploymentRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.deployment_type", "SINGLE_MODEL_FLEX"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.compute_target_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.infrastructure_type", "MANAGED_COMPUTE_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.environment_configuration_type", "OCIR_CONTAINER"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.image", "iad.ocir.io/ociodscdev/sample-byoc:cpu"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.image_digest", "sha256:2b90a418013aae4422177e24c6fd2269931272efc4fda3c54e82ab1da60219af"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.server_port", "8080"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.health_check_port", "8080"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.0.resource_request_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.0.resource_request_configuration.0.memory_in_gbs", "6"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.0.resource_request_configuration.0.ocpus", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.0.resource_limit_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.0.resource_limit_configuration.0.memory_in_gbs", "6"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.0.resource_limit_configuration.0.ocpus", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.policy_type", "FIXED_SIZE"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.instance_count", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_url"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
			),
		},

		{
			Config: config + compartmentIdVariableStr + DatascienceMCCModelDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Optional, acctest.Update, DatascienceMCCModelDeploymentUpdateToAutoscalingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.deployment_type", "SINGLE_MODEL_FLEX"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.compute_target_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.infrastructure_type", "MANAGED_COMPUTE_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.environment_configuration_type", "OCIR_CONTAINER"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.image", "iad.ocir.io/ociodscdev/sample-byoc:cpu"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.image_digest", "sha256:2b90a418013aae4422177e24c6fd2269931272efc4fda3c54e82ab1da60219af"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.server_port", "8080"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.health_check_port", "8080"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.policy_type", "AUTOSCALING"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.auto_scaling_policy_type", "THRESHOLD"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.minimum_instance_count", "0"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.maximum_instance_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.initial_instance_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.metric_expression_rule_type", "TARGET_CUSTOM_EXPRESSION"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_configuration.0.target_scaling_configuration_type", "QUERY"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_configuration.0.query", "PredictRequestCount[5m]{resourceId = \"MODEL_DEPLOYMENT_OCID\"}.count()"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.scale_in_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.scale_in_policy.0.cool_down_in_seconds", "300"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.scale_in_policy.0.instance_count_adjustment", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.scale_in_policy.0.pending_duration", "PT1M"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.scale_out_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.scale_out_policy.0.cool_down_in_seconds", "30"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.scale_out_policy.0.instance_count_adjustment", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.scale_out_policy.0.pending_duration", "PT1M"),
				func(s *terraform.State) (err error) {
					mccResId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if mccResId != mccResId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_model_deployments", "test_model_deployments", acctest.Optional, acctest.Update, DatascienceDatascienceModelDeploymentDataSourceRepresentation) +
				compartmentIdVariableStr + DatascienceMCCModelDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Optional, acctest.Update, DatascienceMCCModelDeploymentUpdateToAutoscalingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "model_deployments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.deployment_type", "SINGLE_MODEL_FLEX"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.compute_target_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.infrastructure_type", "MANAGED_COMPUTE_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.environment_configuration_type", "OCIR_CONTAINER"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.image", "iad.ocir.io/ociodscdev/sample-byoc:cpu"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.image_digest", "sha256:2b90a418013aae4422177e24c6fd2269931272efc4fda3c54e82ab1da60219af"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.server_port", "8080"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.health_check_port", "8080"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.0.resource_request_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.0.resource_request_configuration.0.memory_in_gbs", "6"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.0.resource_request_configuration.0.ocpus", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.0.resource_limit_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.0.resource_limit_configuration.0.memory_in_gbs", "6"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.0.resource_limit_configuration.0.ocpus", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.policy_type", "AUTOSCALING"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.auto_scaling_policy_type", "THRESHOLD"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.minimum_instance_count", "0"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.maximum_instance_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.initial_instance_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.metric_expression_rule_type", "TARGET_CUSTOM_EXPRESSION"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_configuration.0.target_scaling_configuration_type", "QUERY"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_configuration.0.query", "PredictRequestCount[5m]{resourceId = \"MODEL_DEPLOYMENT_OCID\"}.count()"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.scale_in_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.scale_in_policy.0.cool_down_in_seconds", "300"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.scale_in_policy.0.instance_count_adjustment", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.scale_in_policy.0.pending_duration", "PT1M"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.scale_out_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.scale_out_policy.0.cool_down_in_seconds", "30"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.scale_out_policy.0.instance_count_adjustment", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.scale_out_policy.0.pending_duration", "PT1M"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_url"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
			),
		},

		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Required, acctest.Create, DatascienceDatascienceModelDeploymentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatascienceMCCModelDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_deployment", "test_model_deployment", acctest.Optional, acctest.Update, DatascienceMCCModelDeploymentUpdateToAutoscalingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "model_deployment_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.deployment_type", "SINGLE_MODEL_FLEX"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.model_configuration_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.compute_target_id"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.infrastructure_type", "MANAGED_COMPUTE_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.environment_configuration_type", "OCIR_CONTAINER"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.image", "iad.ocir.io/ociodscdev/sample-byoc:cpu"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.image_digest", "sha256:2b90a418013aae4422177e24c6fd2269931272efc4fda3c54e82ab1da60219af"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.server_port", "8080"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.environment_configuration_details.0.health_check_port", "8080"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.0.resource_request_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.0.resource_request_configuration.0.memory_in_gbs", "6"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.model_deployment_resource_configuration.0.resource_request_configuration.0.ocpus", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.policy_type", "AUTOSCALING"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.auto_scaling_policy_type", "THRESHOLD"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.minimum_instance_count", "0"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.maximum_instance_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.initial_instance_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.metric_expression_rule_type", "TARGET_CUSTOM_EXPRESSION"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_configuration.0.target_scaling_configuration_type", "QUERY"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.rules.0.scale_configuration.0.query", "PredictRequestCount[5m]{resourceId = \"MODEL_DEPLOYMENT_OCID\"}.count()"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.scale_in_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.scale_in_policy.0.cool_down_in_seconds", "300"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.scale_in_policy.0.instance_count_adjustment", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.scale_in_policy.0.pending_duration", "PT1M"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.scale_out_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.scale_out_policy.0.cool_down_in_seconds", "30"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.scale_out_policy.0.instance_count_adjustment", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_deployment_configuration_details.0.infrastructure_configuration_details.0.scaling_policy.0.auto_scaling_policies.0.scale_out_policy.0.pending_duration", "PT1M"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_configuration_details.0.model_configuration_details.0.model_id"),
				resource.TestCheckResourceAttrSet(resourceName, "model_deployment_url"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
			),
		},

		{
			Config:            config + DatascienceMCCModelDeploymentRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"description",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatascienceModelDeploymentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataScienceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datascience_model_deployment" {
			noResourceFound = false
			request := oci_datascience.GetModelDeploymentRequest{}

			tmp := rs.Primary.ID
			request.ModelDeploymentId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")

			response, err := client.GetModelDeployment(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_datascience.ModelDeploymentLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatascienceModelDeployment") {
		resource.AddTestSweepers("DatascienceModelDeployment", &resource.Sweeper{
			Name:         "DatascienceModelDeployment",
			Dependencies: acctest.DependencyGraph["modelDeployment"],
			F:            sweepDatascienceModelDeploymentResource,
		})
	}
}

func sweepDatascienceModelDeploymentResource(compartment string) error {
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()
	modelDeploymentIds, err := getDatascienceModelDeploymentIds(compartment)
	if err != nil {
		return err
	}
	for _, modelDeploymentId := range modelDeploymentIds {
		if ok := acctest.SweeperDefaultResourceId[modelDeploymentId]; !ok {
			deleteModelDeploymentRequest := oci_datascience.DeleteModelDeploymentRequest{}

			deleteModelDeploymentRequest.ModelDeploymentId = &modelDeploymentId

			deleteModelDeploymentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")
			_, error := dataScienceClient.DeleteModelDeployment(context.Background(), deleteModelDeploymentRequest)
			if error != nil {
				fmt.Printf("Error deleting ModelDeployment %s %s, It is possible that the resource is already deleted. Please verify manually \n", modelDeploymentId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &modelDeploymentId, DatascienceModelDeploymentSweepWaitCondition, time.Duration(3*time.Minute),
				DatascienceModelDeploymentSweepResponseFetchOperation, "datascience", true)
		}
	}
	return nil
}

func getDatascienceModelDeploymentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ModelDeploymentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()

	listModelDeploymentsRequest := oci_datascience.ListModelDeploymentsRequest{}
	listModelDeploymentsRequest.CompartmentId = &compartmentId
	listModelDeploymentsRequest.LifecycleState = oci_datascience.ListModelDeploymentsLifecycleStateNeedsAttention
	// listModelDeploymentsRequest.LifecycleState = oci_datascience.ListModelDeploymentsLifecycleStateActiveNeedsAttention
	listModelDeploymentsResponse, err := dataScienceClient.ListModelDeployments(context.Background(), listModelDeploymentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ModelDeployment list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, modelDeployment := range listModelDeploymentsResponse.Items {
		id := *modelDeployment.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ModelDeploymentId", id)
	}
	return resourceIds, nil
}

func DatascienceModelDeploymentSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if modelDeploymentResponse, ok := response.Response.(oci_datascience.GetModelDeploymentResponse); ok {
		return modelDeploymentResponse.LifecycleState != oci_datascience.ModelDeploymentLifecycleStateDeleted
	}
	return false
}

func DatascienceModelDeploymentSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataScienceClient().GetModelDeployment(context.Background(), oci_datascience.GetModelDeploymentRequest{
		ModelDeploymentId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
