// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import (
	"context"
	"log"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_devops "github.com/oracle/oci-go-sdk/v58/devops"
)

func DevopsDeploymentDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["deployment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DevopsDeploymentResource(), fieldMap, readSingularDevopsDeployment)
}

func readSingularDevopsDeployment(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsDeploymentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

type DevopsDeploymentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.GetDeploymentResponse
}

func (s *DevopsDeploymentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsDeploymentDataSourceCrud) Get() error {
	request := oci_devops.GetDeploymentRequest{}

	if deploymentId, ok := s.D.GetOkExists("deployment_id"); ok {
		tmp := deploymentId.(string)
		request.DeploymentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "devops")

	response, err := s.Client.GetDeployment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DevopsDeploymentDataSourceCrud) SetData() error {
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

	if s.Res.GetDeployArtifactOverrideArguments() != nil {
		s.D.Set("deploy_artifact_override_arguments", []interface{}{DeployArtifactOverrideArgumentCollectionToMap(s.Res.GetDeployArtifactOverrideArguments())})
	} else {
		s.D.Set("deploy_artifact_override_arguments", nil)
	}

	if s.Res.GetDeployPipelineArtifacts() != nil {
		s.D.Set("deploy_pipeline_artifacts", []interface{}{DeployPipelineArtifactCollectionToMap(s.Res.GetDeployPipelineArtifacts())})
	} else {
		s.D.Set("deploy_pipeline_artifacts", nil)
	}

	if s.Res.GetDeployPipelineEnvironments() != nil {
		s.D.Set("deploy_pipeline_environments", []interface{}{DeployPipelineEnvironmentCollectionToMap(s.Res.GetDeployPipelineEnvironments())})
	} else {
		s.D.Set("deploy_pipeline_environments", nil)
	}

	if s.Res.GetDeployPipelineId() != nil {
		s.D.Set("deploy_pipeline_id", *s.Res.GetDeployPipelineId())
	}

	if s.Res.GetDeploymentArguments() != nil {
		s.D.Set("deployment_arguments", []interface{}{DeploymentArgumentCollectionToMap(s.Res.GetDeploymentArguments())})
	} else {
		s.D.Set("deployment_arguments", nil)
	}

	if s.Res.GetDeploymentExecutionProgress() != nil {
		s.D.Set("deployment_execution_progress", []interface{}{DeploymentExecutionProgressToMap(s.Res.GetDeploymentExecutionProgress())})
	} else {
		s.D.Set("deployment_execution_progress", nil)
	}

	if s.Res.GetDisplayName() != nil {
		s.D.Set("display_name", *s.Res.GetDisplayName())
	}

	s.D.Set("freeform_tags", s.Res.GetFreeformTags())

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

	switch v := (s.Res.Deployment).(type) {
	case oci_devops.DeployPipelineDeployment:
		s.D.Set("deployment_type", "PIPELINE_DEPLOYMENT")

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployArtifactOverrideArguments != nil {
			s.D.Set("deploy_artifact_override_arguments", []interface{}{DeployArtifactOverrideArgumentCollectionToMap(v.DeployArtifactOverrideArguments)})
		} else {
			s.D.Set("deploy_artifact_override_arguments", nil)
		}

		if v.DeployPipelineArtifacts != nil {
			s.D.Set("deploy_pipeline_artifacts", []interface{}{DeployPipelineArtifactCollectionToMap(v.DeployPipelineArtifacts)})
		} else {
			s.D.Set("deploy_pipeline_artifacts", nil)
		}

		if v.DeployPipelineEnvironments != nil {
			s.D.Set("deploy_pipeline_environments", []interface{}{DeployPipelineEnvironmentCollectionToMap(v.DeployPipelineEnvironments)})
		} else {
			s.D.Set("deploy_pipeline_environments", nil)
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeploymentArguments != nil {
			s.D.Set("deployment_arguments", []interface{}{DeploymentArgumentCollectionToMap(v.DeploymentArguments)})
		} else {
			s.D.Set("deployment_arguments", nil)
		}

		if v.DeploymentExecutionProgress != nil {
			s.D.Set("deployment_execution_progress", []interface{}{DeploymentExecutionProgressToMap(v.DeploymentExecutionProgress)})
		} else {
			s.D.Set("deployment_execution_progress", nil)
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
	case oci_devops.DeployPipelineRedeployment:
		s.D.Set("deployment_type", "PIPELINE_REDEPLOYMENT")

		if v.PreviousDeploymentId != nil {
			s.D.Set("previous_deployment_id", *v.PreviousDeploymentId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployArtifactOverrideArguments != nil {
			s.D.Set("deploy_artifact_override_arguments", []interface{}{DeployArtifactOverrideArgumentCollectionToMap(v.DeployArtifactOverrideArguments)})
		} else {
			s.D.Set("deploy_artifact_override_arguments", nil)
		}

		if v.DeployPipelineArtifacts != nil {
			s.D.Set("deploy_pipeline_artifacts", []interface{}{DeployPipelineArtifactCollectionToMap(v.DeployPipelineArtifacts)})
		} else {
			s.D.Set("deploy_pipeline_artifacts", nil)
		}

		if v.DeployPipelineEnvironments != nil {
			s.D.Set("deploy_pipeline_environments", []interface{}{DeployPipelineEnvironmentCollectionToMap(v.DeployPipelineEnvironments)})
		} else {
			s.D.Set("deploy_pipeline_environments", nil)
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeploymentArguments != nil {
			s.D.Set("deployment_arguments", []interface{}{DeploymentArgumentCollectionToMap(v.DeploymentArguments)})
		} else {
			s.D.Set("deployment_arguments", nil)
		}

		if v.DeploymentExecutionProgress != nil {
			s.D.Set("deployment_execution_progress", []interface{}{DeploymentExecutionProgressToMap(v.DeploymentExecutionProgress)})
		} else {
			s.D.Set("deployment_execution_progress", nil)
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
	case oci_devops.SingleDeployStageDeployment:
		s.D.Set("deployment_type", "SINGLE_STAGE_DEPLOYMENT")

		if v.DeployStageId != nil {
			s.D.Set("deploy_stage_id", *v.DeployStageId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployArtifactOverrideArguments != nil {
			s.D.Set("deploy_artifact_override_arguments", []interface{}{DeployArtifactOverrideArgumentCollectionToMap(v.DeployArtifactOverrideArguments)})
		} else {
			s.D.Set("deploy_artifact_override_arguments", nil)
		}

		if v.DeployPipelineArtifacts != nil {
			s.D.Set("deploy_pipeline_artifacts", []interface{}{DeployPipelineArtifactCollectionToMap(v.DeployPipelineArtifacts)})
		} else {
			s.D.Set("deploy_pipeline_artifacts", nil)
		}

		if v.DeployPipelineEnvironments != nil {
			s.D.Set("deploy_pipeline_environments", []interface{}{DeployPipelineEnvironmentCollectionToMap(v.DeployPipelineEnvironments)})
		} else {
			s.D.Set("deploy_pipeline_environments", nil)
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeploymentArguments != nil {
			s.D.Set("deployment_arguments", []interface{}{DeploymentArgumentCollectionToMap(v.DeploymentArguments)})
		} else {
			s.D.Set("deployment_arguments", nil)
		}

		if v.DeploymentExecutionProgress != nil {
			s.D.Set("deployment_execution_progress", []interface{}{DeploymentExecutionProgressToMap(v.DeploymentExecutionProgress)})
		} else {
			s.D.Set("deployment_execution_progress", nil)
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
		log.Printf("[WARN] Received 'deployment_type' of unknown type %v", *s.Res)
	}
	return nil
}
