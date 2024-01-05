// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import (
	"context"
	"log"
	"strconv"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_devops "github.com/oracle/oci-go-sdk/v65/devops"
)

func DevopsDeployStageDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["deploy_stage_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DevopsDeployStageResource(), fieldMap, readSingularDevopsDeployStage)
}

func readSingularDevopsDeployStage(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsDeployStageDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

type DevopsDeployStageDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.GetDeployStageResponse
}

func (s *DevopsDeployStageDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsDeployStageDataSourceCrud) Get() error {
	request := oci_devops.GetDeployStageRequest{}

	if deployStageId, ok := s.D.GetOkExists("deploy_stage_id"); ok {
		tmp := deployStageId.(string)
		request.DeployStageId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "devops")

	response, err := s.Client.GetDeployStage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DevopsDeployStageDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.GetId())
	switch v := (s.Res.DeployStage).(type) {
	case oci_devops.ComputeInstanceGroupBlueGreenDeployStage:
		s.D.Set("deploy_stage_type", "COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT")

		s.D.Set("deploy_artifact_ids", v.DeployArtifactIds)

		if v.DeployEnvironmentIdA != nil {
			s.D.Set("deploy_environment_id_a", *v.DeployEnvironmentIdA)
		}

		if v.DeployEnvironmentIdB != nil {
			s.D.Set("deploy_environment_id_b", *v.DeployEnvironmentIdB)
		}

		if v.DeploymentSpecDeployArtifactId != nil {
			s.D.Set("deployment_spec_deploy_artifact_id", *v.DeploymentSpecDeployArtifactId)
		}

		if v.FailurePolicy != nil {
			failurePolicyArray := []interface{}{}
			if failurePolicyMap := ComputeInstanceGroupFailurePolicyToMap(&v.FailurePolicy); failurePolicyMap != nil {
				failurePolicyArray = append(failurePolicyArray, failurePolicyMap)
			}
			s.D.Set("failure_policy", failurePolicyArray)
		} else {
			s.D.Set("failure_policy", nil)
		}

		if v.ProductionLoadBalancerConfig != nil {
			s.D.Set("production_load_balancer_config", []interface{}{LoadBalancerConfigToMap(v.ProductionLoadBalancerConfig)})
		} else {
			s.D.Set("production_load_balancer_config", nil)
		}

		if v.RolloutPolicy != nil {
			rolloutPolicyArray := []interface{}{}
			if rolloutPolicyMap := ComputeInstanceGroupRolloutPolicyToMap(&v.RolloutPolicy); rolloutPolicyMap != nil {
				rolloutPolicyArray = append(rolloutPolicyArray, rolloutPolicyMap)
			}
			s.D.Set("rollout_policy", rolloutPolicyArray)
		} else {
			s.D.Set("rollout_policy", nil)
		}

		if v.TestLoadBalancerConfig != nil {
			s.D.Set("test_load_balancer_config", []interface{}{LoadBalancerConfigToMap(v.TestLoadBalancerConfig)})
		} else {
			s.D.Set("test_load_balancer_config", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.ComputeInstanceGroupBlueGreenTrafficShiftDeployStage:
		s.D.Set("deploy_stage_type", "COMPUTE_INSTANCE_GROUP_BLUE_GREEN_TRAFFIC_SHIFT")

		if v.ComputeInstanceGroupBlueGreenDeploymentDeployStageId != nil {
			s.D.Set("compute_instance_group_blue_green_deployment_deploy_stage_id", *v.ComputeInstanceGroupBlueGreenDeploymentDeployStageId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.ComputeInstanceGroupCanaryApprovalDeployStage:
		s.D.Set("deploy_stage_type", "COMPUTE_INSTANCE_GROUP_CANARY_APPROVAL")

		if v.ApprovalPolicy != nil {
			approvalPolicyArray := []interface{}{}
			if approvalPolicyMap := ApprovalPolicyToMap(&v.ApprovalPolicy); approvalPolicyMap != nil {
				approvalPolicyArray = append(approvalPolicyArray, approvalPolicyMap)
			}
			s.D.Set("approval_policy", approvalPolicyArray)
		} else {
			s.D.Set("approval_policy", nil)
		}

		if v.ComputeInstanceGroupCanaryTrafficShiftDeployStageId != nil {
			s.D.Set("compute_instance_group_canary_traffic_shift_deploy_stage_id", *v.ComputeInstanceGroupCanaryTrafficShiftDeployStageId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.ComputeInstanceGroupCanaryDeployStage:
		s.D.Set("deploy_stage_type", "COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT")

		if v.ComputeInstanceGroupDeployEnvironmentId != nil {
			s.D.Set("compute_instance_group_deploy_environment_id", *v.ComputeInstanceGroupDeployEnvironmentId)
		}

		s.D.Set("deploy_artifact_ids", v.DeployArtifactIds)

		if v.DeploymentSpecDeployArtifactId != nil {
			s.D.Set("deployment_spec_deploy_artifact_id", *v.DeploymentSpecDeployArtifactId)
		}

		if v.ProductionLoadBalancerConfig != nil {
			s.D.Set("production_load_balancer_config", []interface{}{LoadBalancerConfigToMap(v.ProductionLoadBalancerConfig)})
		} else {
			s.D.Set("production_load_balancer_config", nil)
		}

		if v.RolloutPolicy != nil {
			rolloutPolicyArray := []interface{}{}
			if rolloutPolicyMap := ComputeInstanceGroupRolloutPolicyToMap(&v.RolloutPolicy); rolloutPolicyMap != nil {
				rolloutPolicyArray = append(rolloutPolicyArray, rolloutPolicyMap)
			}
			s.D.Set("rollout_policy", rolloutPolicyArray)
		} else {
			s.D.Set("rollout_policy", nil)
		}

		if v.TestLoadBalancerConfig != nil {
			s.D.Set("test_load_balancer_config", []interface{}{LoadBalancerConfigToMap(v.TestLoadBalancerConfig)})
		} else {
			s.D.Set("test_load_balancer_config", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.ComputeInstanceGroupCanaryTrafficShiftDeployStage:
		s.D.Set("deploy_stage_type", "COMPUTE_INSTANCE_GROUP_CANARY_TRAFFIC_SHIFT")

		if v.ComputeInstanceGroupCanaryDeployStageId != nil {
			s.D.Set("compute_instance_group_canary_deploy_stage_id", *v.ComputeInstanceGroupCanaryDeployStageId)
		}

		if v.RolloutPolicy != nil {
			s.D.Set("rollout_policy", []interface{}{LoadBalancerTrafficShiftRolloutPolicyToMap(v.RolloutPolicy)})
		} else {
			s.D.Set("rollout_policy", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.ComputeInstanceGroupDeployStage:
		s.D.Set("deploy_stage_type", "COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT")

		if v.ComputeInstanceGroupDeployEnvironmentId != nil {
			s.D.Set("compute_instance_group_deploy_environment_id", *v.ComputeInstanceGroupDeployEnvironmentId)
		}

		s.D.Set("deploy_artifact_ids", v.DeployArtifactIds)

		if v.DeploymentSpecDeployArtifactId != nil {
			s.D.Set("deployment_spec_deploy_artifact_id", *v.DeploymentSpecDeployArtifactId)
		}

		if v.FailurePolicy != nil {
			failurePolicyArray := []interface{}{}
			if failurePolicyMap := ComputeInstanceGroupFailurePolicyToMap(&v.FailurePolicy); failurePolicyMap != nil {
				failurePolicyArray = append(failurePolicyArray, failurePolicyMap)
			}
			s.D.Set("failure_policy", failurePolicyArray)
		} else {
			s.D.Set("failure_policy", nil)
		}

		if v.LoadBalancerConfig != nil {
			s.D.Set("load_balancer_config", []interface{}{LoadBalancerConfigToMap(v.LoadBalancerConfig)})
		} else {
			s.D.Set("load_balancer_config", nil)
		}

		if v.RollbackPolicy != nil {
			rollbackPolicyArray := []interface{}{}
			if rollbackPolicyMap := DeployStageRollbackPolicyToMap(&v.RollbackPolicy); rollbackPolicyMap != nil {
				rollbackPolicyArray = append(rollbackPolicyArray, rollbackPolicyMap)
			}
			s.D.Set("rollback_policy", rollbackPolicyArray)
		} else {
			s.D.Set("rollback_policy", nil)
		}

		if v.RolloutPolicy != nil {
			rolloutPolicyArray := []interface{}{}
			if rolloutPolicyMap := ComputeInstanceGroupRolloutPolicyToMap(&v.RolloutPolicy); rolloutPolicyMap != nil {
				rolloutPolicyArray = append(rolloutPolicyArray, rolloutPolicyMap)
			}
			s.D.Set("rollout_policy", rolloutPolicyArray)
		} else {
			s.D.Set("rollout_policy", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.FunctionDeployStage:
		s.D.Set("deploy_stage_type", "DEPLOY_FUNCTION")

		s.D.Set("config", v.Config)

		if v.DockerImageDeployArtifactId != nil {
			s.D.Set("docker_image_deploy_artifact_id", *v.DockerImageDeployArtifactId)
		}

		if v.FunctionDeployEnvironmentId != nil {
			s.D.Set("function_deploy_environment_id", *v.FunctionDeployEnvironmentId)
		}

		if v.FunctionTimeoutInSeconds != nil {
			s.D.Set("function_timeout_in_seconds", *v.FunctionTimeoutInSeconds)
		}

		if v.MaxMemoryInMBs != nil {
			s.D.Set("max_memory_in_mbs", strconv.FormatInt(*v.MaxMemoryInMBs, 10))
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.InvokeFunctionDeployStage:
		s.D.Set("deploy_stage_type", "INVOKE_FUNCTION")

		if v.DeployArtifactId != nil {
			s.D.Set("deploy_artifact_id", *v.DeployArtifactId)
		}

		if v.FunctionDeployEnvironmentId != nil {
			s.D.Set("function_deploy_environment_id", *v.FunctionDeployEnvironmentId)
		}

		if v.IsAsync != nil {
			s.D.Set("is_async", *v.IsAsync)
		}

		if v.IsValidationEnabled != nil {
			s.D.Set("is_validation_enabled", *v.IsValidationEnabled)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.LoadBalancerTrafficShiftDeployStage:
		s.D.Set("deploy_stage_type", "LOAD_BALANCER_TRAFFIC_SHIFT")

		if v.BlueBackendIps != nil {
			s.D.Set("blue_backend_ips", []interface{}{BackendSetIpCollectionToMap(v.BlueBackendIps)})
		} else {
			s.D.Set("blue_backend_ips", nil)
		}

		if v.GreenBackendIps != nil {
			s.D.Set("green_backend_ips", []interface{}{BackendSetIpCollectionToMap(v.GreenBackendIps)})
		} else {
			s.D.Set("green_backend_ips", nil)
		}

		if v.LoadBalancerConfig != nil {
			s.D.Set("load_balancer_config", []interface{}{LoadBalancerConfigToMap(v.LoadBalancerConfig)})
		} else {
			s.D.Set("load_balancer_config", nil)
		}

		if v.RollbackPolicy != nil {
			rollbackPolicyArray := []interface{}{}
			if rollbackPolicyMap := DeployStageRollbackPolicyToMap(&v.RollbackPolicy); rollbackPolicyMap != nil {
				rollbackPolicyArray = append(rollbackPolicyArray, rollbackPolicyMap)
			}
			s.D.Set("rollback_policy", rollbackPolicyArray)
		} else {
			s.D.Set("rollback_policy", nil)
		}

		if v.RolloutPolicy != nil {
			s.D.Set("rollout_policy", []interface{}{LoadBalancerTrafficShiftRolloutPolicyToMap(v.RolloutPolicy)})
		} else {
			s.D.Set("rollout_policy", nil)
		}

		s.D.Set("traffic_shift_target", v.TrafficShiftTarget)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.ManualApprovalDeployStage:
		s.D.Set("deploy_stage_type", "MANUAL_APPROVAL")

		if v.ApprovalPolicy != nil {
			approvalPolicyArray := []interface{}{}
			if approvalPolicyMap := ApprovalPolicyToMap(&v.ApprovalPolicy); approvalPolicyMap != nil {
				approvalPolicyArray = append(approvalPolicyArray, approvalPolicyMap)
			}
			s.D.Set("approval_policy", approvalPolicyArray)
		} else {
			s.D.Set("approval_policy", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.OkeBlueGreenDeployStage:
		s.D.Set("deploy_stage_type", "OKE_BLUE_GREEN_DEPLOYMENT")

		if v.BlueGreenStrategy != nil {
			blueGreenStrategyArray := []interface{}{}
			if blueGreenStrategyMap := OkeBlueGreenStrategyToMap(&v.BlueGreenStrategy); blueGreenStrategyMap != nil {
				blueGreenStrategyArray = append(blueGreenStrategyArray, blueGreenStrategyMap)
			}
			s.D.Set("blue_green_strategy", blueGreenStrategyArray)
		} else {
			s.D.Set("blue_green_strategy", nil)
		}

		s.D.Set("kubernetes_manifest_deploy_artifact_ids", v.KubernetesManifestDeployArtifactIds)

		if v.OkeClusterDeployEnvironmentId != nil {
			s.D.Set("oke_cluster_deploy_environment_id", *v.OkeClusterDeployEnvironmentId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.OkeBlueGreenTrafficShiftDeployStage:
		s.D.Set("deploy_stage_type", "OKE_BLUE_GREEN_TRAFFIC_SHIFT")

		if v.OkeBlueGreenDeployStageId != nil {
			s.D.Set("oke_blue_green_deploy_stage_id", *v.OkeBlueGreenDeployStageId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.OkeCanaryApprovalDeployStage:
		s.D.Set("deploy_stage_type", "OKE_CANARY_APPROVAL")

		if v.ApprovalPolicy != nil {
			approvalPolicyArray := []interface{}{}
			if approvalPolicyMap := ApprovalPolicyToMap(&v.ApprovalPolicy); approvalPolicyMap != nil {
				approvalPolicyArray = append(approvalPolicyArray, approvalPolicyMap)
			}
			s.D.Set("approval_policy", approvalPolicyArray)
		} else {
			s.D.Set("approval_policy", nil)
		}

		if v.OkeCanaryTrafficShiftDeployStageId != nil {
			s.D.Set("oke_canary_traffic_shift_deploy_stage_id", *v.OkeCanaryTrafficShiftDeployStageId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.OkeCanaryDeployStage:
		s.D.Set("deploy_stage_type", "OKE_CANARY_DEPLOYMENT")

		if v.CanaryStrategy != nil {
			canaryStrategyArray := []interface{}{}
			if canaryStrategyMap := OkeCanaryStrategyToMap(&v.CanaryStrategy); canaryStrategyMap != nil {
				canaryStrategyArray = append(canaryStrategyArray, canaryStrategyMap)
			}
			s.D.Set("canary_strategy", canaryStrategyArray)
		} else {
			s.D.Set("canary_strategy", nil)
		}

		s.D.Set("kubernetes_manifest_deploy_artifact_ids", v.KubernetesManifestDeployArtifactIds)

		if v.OkeClusterDeployEnvironmentId != nil {
			s.D.Set("oke_cluster_deploy_environment_id", *v.OkeClusterDeployEnvironmentId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.OkeCanaryTrafficShiftDeployStage:
		s.D.Set("deploy_stage_type", "OKE_CANARY_TRAFFIC_SHIFT")

		if v.OkeCanaryDeployStageId != nil {
			s.D.Set("oke_canary_deploy_stage_id", *v.OkeCanaryDeployStageId)
		}

		if v.RolloutPolicy != nil {
			s.D.Set("rollout_policy", []interface{}{LoadBalancerTrafficShiftRolloutPolicyToMap(v.RolloutPolicy)})
		} else {
			s.D.Set("rollout_policy", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.OkeDeployStage:
		s.D.Set("deploy_stage_type", "OKE_DEPLOYMENT")

		s.D.Set("kubernetes_manifest_deploy_artifact_ids", v.KubernetesManifestDeployArtifactIds)

		if v.Namespace != nil {
			s.D.Set("namespace", *v.Namespace)
		}

		if v.OkeClusterDeployEnvironmentId != nil {
			s.D.Set("oke_cluster_deploy_environment_id", *v.OkeClusterDeployEnvironmentId)
		}

		if v.RollbackPolicy != nil {
			rollbackPolicyArray := []interface{}{}
			if rollbackPolicyMap := DeployStageRollbackPolicyToMap(&v.RollbackPolicy); rollbackPolicyMap != nil {
				rollbackPolicyArray = append(rollbackPolicyArray, rollbackPolicyMap)
			}
			s.D.Set("rollback_policy", rollbackPolicyArray)
		} else {
			s.D.Set("rollback_policy", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.OkeHelmChartDeployStage:
		s.D.Set("deploy_stage_type", "OKE_HELM_CHART_DEPLOYMENT")

		if v.AreHooksEnabled != nil {
			s.D.Set("are_hooks_enabled", *v.AreHooksEnabled)
		}

		if v.HelmChartDeployArtifactId != nil {
			s.D.Set("helm_chart_deploy_artifact_id", *v.HelmChartDeployArtifactId)
		}

		if v.IsDebugEnabled != nil {
			s.D.Set("is_debug_enabled", *v.IsDebugEnabled)
		}

		if v.IsForceEnabled != nil {
			s.D.Set("is_force_enabled", *v.IsForceEnabled)
		}

		if v.MaxHistory != nil {
			s.D.Set("max_history", *v.MaxHistory)
		}

		if v.Namespace != nil {
			s.D.Set("namespace", *v.Namespace)
		}

		if v.OkeClusterDeployEnvironmentId != nil {
			s.D.Set("oke_cluster_deploy_environment_id", *v.OkeClusterDeployEnvironmentId)
		}

		if v.ReleaseName != nil {
			s.D.Set("release_name", *v.ReleaseName)
		}

		if v.RollbackPolicy != nil {
			rollbackPolicyArray := []interface{}{}
			if rollbackPolicyMap := DeployStageRollbackPolicyToMap(&v.RollbackPolicy); rollbackPolicyMap != nil {
				rollbackPolicyArray = append(rollbackPolicyArray, rollbackPolicyMap)
			}
			s.D.Set("rollback_policy", rollbackPolicyArray)
		} else {
			s.D.Set("rollback_policy", nil)
		}

		if v.SetString != nil {
			s.D.Set("set_string", []interface{}{HelmSetValueCollectionToMap(v.SetString)})
		} else {
			s.D.Set("set_string", nil)
		}

		if v.SetValues != nil {
			s.D.Set("set_values", []interface{}{HelmSetValueCollectionToMap(v.SetValues)})
		} else {
			s.D.Set("set_values", nil)
		}

		if v.ShouldCleanupOnFail != nil {
			s.D.Set("should_cleanup_on_fail", *v.ShouldCleanupOnFail)
		}

		if v.ShouldNotWait != nil {
			s.D.Set("should_not_wait", *v.ShouldNotWait)
		}

		if v.ShouldResetValues != nil {
			s.D.Set("should_reset_values", *v.ShouldResetValues)
		}

		if v.ShouldReuseValues != nil {
			s.D.Set("should_reuse_values", *v.ShouldReuseValues)
		}

		if v.ShouldSkipCrds != nil {
			s.D.Set("should_skip_crds", *v.ShouldSkipCrds)
		}

		if v.ShouldSkipRenderSubchartNotes != nil {
			s.D.Set("should_skip_render_subchart_notes", *v.ShouldSkipRenderSubchartNotes)
		}

		if v.TimeoutInSeconds != nil {
			s.D.Set("timeout_in_seconds", *v.TimeoutInSeconds)
		}

		s.D.Set("values_artifact_ids", v.ValuesArtifactIds)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.ShellDeployStage:
		s.D.Set("deploy_stage_type", "SHELL")

		if v.CommandSpecDeployArtifactId != nil {
			s.D.Set("command_spec_deploy_artifact_id", *v.CommandSpecDeployArtifactId)
		}

		if v.ContainerConfig != nil {
			containerConfigArray := []interface{}{}
			if containerConfigMap := ContainerConfigToMap(&v.ContainerConfig); containerConfigMap != nil {
				containerConfigArray = append(containerConfigArray, containerConfigMap)
			}
			s.D.Set("container_config", containerConfigArray)
		} else {
			s.D.Set("container_config", nil)
		}

		if v.TimeoutInSeconds != nil {
			s.D.Set("timeout_in_seconds", *v.TimeoutInSeconds)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.WaitDeployStage:
		s.D.Set("deploy_stage_type", "WAIT")

		if v.WaitCriteria != nil {
			waitCriteriaArray := []interface{}{}
			if waitCriteriaMap := WaitCriteriaToMap(&v.WaitCriteria); waitCriteriaMap != nil {
				waitCriteriaArray = append(waitCriteriaArray, waitCriteriaMap)
			}
			s.D.Set("wait_criteria", waitCriteriaArray)
		} else {
			s.D.Set("wait_criteria", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'deploy_stage_type' of unknown type %v", s.Res.DeployStage)
		return nil
	}

	return nil
}
