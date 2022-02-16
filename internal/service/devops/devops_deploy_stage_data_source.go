// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import (
	"context"
	"log"
	"strconv"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_devops "github.com/oracle/oci-go-sdk/v58/devops"
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

	if s.Res.GetCompartmentId() != nil {
		s.D.Set("compartment_id", *s.Res.GetCompartmentId())
	}

	if s.Res.GetDefinedTags() != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.GetDefinedTags()))
	}

	if s.Res.GetDeployPipelineId() != nil {
		s.D.Set("deploy_pipeline_id", *s.Res.GetDeployPipelineId())
	}

	if s.Res.GetDeployStagePredecessorCollection() != nil {
		s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(s.Res.GetDeployStagePredecessorCollection())})
	} else {
		s.D.Set("deploy_stage_predecessor_collection", nil)
	}

	if s.Res.GetDescription() != nil {
		s.D.Set("description", *s.Res.GetDescription())
	}

	if s.Res.GetDisplayName() != nil {
		s.D.Set("display_name", *s.Res.GetDisplayName())
	}

	if s.Res.GetLifecycleDetails() != nil {
		s.D.Set("lifecyle_details", *s.Res.GetLifecycleDetails())
	}

	if s.Res.GetProjectId() != nil {
		s.D.Set("project_id", *s.Res.GetProjectId())
	}

	s.D.Set("state", s.Res.GetLifecycleState())

	if s.Res.GetSystemTags() != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.GetSystemTags()))
	}

	if s.Res.GetTimeCreated() != nil {
		s.D.Set("time_created", s.Res.GetTimeCreated().String())
	}

	if s.Res.GetTimeUpdated() != nil {
		s.D.Set("time_updated", s.Res.GetTimeUpdated().String())
	}

	switch v := (s.Res.DeployStage).(type) {
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

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecyle_details", *v.LifecycleDetails)
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

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecyle_details", *v.LifecycleDetails)
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

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecyle_details", *v.LifecycleDetails)
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

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecyle_details", *v.LifecycleDetails)
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

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecyle_details", *v.LifecycleDetails)
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

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecyle_details", *v.LifecycleDetails)
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

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecyle_details", *v.LifecycleDetails)
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
		log.Printf("[WARN] Received 'deploy_stage_type' of unknown type %v", *s.Res)
		return nil

	}

	return nil
}
