// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_devops "github.com/oracle/oci-go-sdk/v65/devops"
)

func DevopsBuildPipelineStageDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["build_pipeline_stage_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DevopsBuildPipelineStageResource(), fieldMap, readSingularDevopsBuildPipelineStage)
}

func readSingularDevopsBuildPipelineStage(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsBuildPipelineStageDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

type DevopsBuildPipelineStageDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.GetBuildPipelineStageResponse
}

func (s *DevopsBuildPipelineStageDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsBuildPipelineStageDataSourceCrud) Get() error {
	request := oci_devops.GetBuildPipelineStageRequest{}

	if buildPipelineStageId, ok := s.D.GetOkExists("build_pipeline_stage_id"); ok {
		tmp := buildPipelineStageId.(string)
		request.BuildPipelineStageId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "devops")

	response, err := s.Client.GetBuildPipelineStage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DevopsBuildPipelineStageDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.GetId())
	switch v := (s.Res.BuildPipelineStage).(type) {
	case oci_devops.BuildStage:
		s.D.Set("build_pipeline_stage_type", "BUILD")

		if v.BuildSpecFile != nil {
			s.D.Set("build_spec_file", *v.BuildSpecFile)
		}

		s.D.Set("image", v.Image)

		if v.PrimaryBuildSource != nil {
			s.D.Set("primary_build_source", *v.PrimaryBuildSource)
		}

		if v.StageExecutionTimeoutInSeconds != nil {
			s.D.Set("stage_execution_timeout_in_seconds", *v.StageExecutionTimeoutInSeconds)
		}

		if v.BuildPipelineId != nil {
			s.D.Set("build_pipeline_id", *v.BuildPipelineId)
		}

		if v.BuildPipelineStagePredecessorCollection != nil {
			s.D.Set("build_pipeline_stage_predecessor_collection", []interface{}{BuildPipelineStagePredecessorCollectionToMap(v.BuildPipelineStagePredecessorCollection)})
		} else {
			s.D.Set("build_pipeline_stage_predecessor_collection", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
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
	case oci_devops.DeliverArtifactStage:
		s.D.Set("build_pipeline_stage_type", "DELIVER_ARTIFACT")

		if v.DeliverArtifactCollection != nil {
			s.D.Set("deliver_artifact_collection", []interface{}{DeliverArtifactCollectionToMap(v.DeliverArtifactCollection)})
		} else {
			s.D.Set("deliver_artifact_collection", nil)
		}

		if v.BuildPipelineId != nil {
			s.D.Set("build_pipeline_id", *v.BuildPipelineId)
		}

		if v.BuildPipelineStagePredecessorCollection != nil {
			s.D.Set("build_pipeline_stage_predecessor_collection", []interface{}{BuildPipelineStagePredecessorCollectionToMap(v.BuildPipelineStagePredecessorCollection)})
		} else {
			s.D.Set("build_pipeline_stage_predecessor_collection", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
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
	case oci_devops.TriggerDeploymentStage:
		s.D.Set("build_pipeline_stage_type", "TRIGGER_DEPLOYMENT_PIPELINE")

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.IsPassAllParametersEnabled != nil {
			s.D.Set("is_pass_all_parameters_enabled", *v.IsPassAllParametersEnabled)
		}

		if v.BuildPipelineId != nil {
			s.D.Set("build_pipeline_id", *v.BuildPipelineId)
		}

		if v.BuildPipelineStagePredecessorCollection != nil {
			s.D.Set("build_pipeline_stage_predecessor_collection", []interface{}{BuildPipelineStagePredecessorCollectionToMap(v.BuildPipelineStagePredecessorCollection)})
		} else {
			s.D.Set("build_pipeline_stage_predecessor_collection", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
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
	case oci_devops.WaitStage:
		s.D.Set("build_pipeline_stage_type", "WAIT")

		if v.WaitCriteria != nil {
			waitCriteriaArray := []interface{}{}
			if waitCriteriaMap := WaitCriteriaToMap(&v.WaitCriteria); waitCriteriaMap != nil {
				waitCriteriaArray = append(waitCriteriaArray, waitCriteriaMap)
			}
			s.D.Set("wait_criteria", waitCriteriaArray)
		} else {
			s.D.Set("wait_criteria", nil)
		}

		if v.BuildPipelineId != nil {
			s.D.Set("build_pipeline_id", *v.BuildPipelineId)
		}

		if v.BuildPipelineStagePredecessorCollection != nil {
			s.D.Set("build_pipeline_stage_predecessor_collection", []interface{}{BuildPipelineStagePredecessorCollectionToMap(v.BuildPipelineStagePredecessorCollection)})
		} else {
			s.D.Set("build_pipeline_stage_predecessor_collection", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
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
		log.Printf("[WARN] Received 'build_pipeline_stage_type' of unknown type %v", s.Res.BuildPipelineStage)
		return nil
	}

	return nil
}
