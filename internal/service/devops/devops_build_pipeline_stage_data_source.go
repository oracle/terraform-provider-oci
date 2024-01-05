// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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

	if s.Res.GetDescription() != nil {
		s.D.Set("description", s.Res.GetDescription())
	}

	s.D.Set("freeform_tags", s.Res.GetFreeformTags())

	s.D.Set("state", s.Res.GetLifecycleState())

	if s.Res.GetBuildPipelineId() != nil {
		s.D.Set("build_pipeline_id", s.Res.GetBuildPipelineId())
	}

	if s.Res.GetBuildPipelineStagePredecessorCollection() != nil {
		s.D.Set("build_pipeline_stage_predecessor_collection", []interface{}{BuildPipelineStagePredecessorCollectionToMap(s.Res.GetBuildPipelineStagePredecessorCollection())})
	} else {
		s.D.Set("build_pipeline_stage_predecessor_collection", nil)
	}

	if s.Res.GetCompartmentId() != nil {
		s.D.Set("compartment_id", s.Res.GetCompartmentId())
	}

	if s.Res.GetDefinedTags() != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.GetDefinedTags()))
	}

	if s.Res.GetDescription() != nil {
		s.D.Set("description", *s.Res.GetDescription())
	}

	if s.Res.GetDisplayName() != nil {
		s.D.Set("display_name", *s.Res.GetDisplayName())
	}

	s.D.Set("freeform_tags", s.Res.GetFreeformTags())

	if s.Res.GetLifecycleDetails() != nil {
		s.D.Set("lifecycle_details", *s.Res.GetLifecycleDetails())
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

	switch v := (s.Res.BuildPipelineStage).(type) {
	case oci_devops.BuildStage:
		s.D.Set("build_pipeline_stage_type", "BUILD")

		s.D.Set("image", v.Image)

		if v.PrimaryBuildSource != nil {
			s.D.Set("primary_build_source", v.PrimaryBuildSource)
		}

		if v.StageExecutionTimeoutInSeconds != nil {
			s.D.Set("stage_execution_timeout_in_seconds", v.StageExecutionTimeoutInSeconds)
		}

		if v.BuildSourceCollection != nil {
			s.D.Set("build_source_collection", []interface{}{BuildSourceCollectionToMap(v.BuildSourceCollection)})
		} else {
			s.D.Set("build_source_collection", nil)
		}

		if v.BuildSpecFile != nil {
			s.D.Set("build_spec_file", v.BuildSpecFile)
		}

		if v.BuildRunnerShapeConfig != nil {
			buildRunnerShapeConfigArray := []interface{}{}
			if buildRunnerShapeConfigMap := BuildRunnerShapeConfigToMap(&v.BuildRunnerShapeConfig); buildRunnerShapeConfigMap != nil {
				buildRunnerShapeConfigArray = append(buildRunnerShapeConfigArray, buildRunnerShapeConfigMap)
			}
			s.D.Set("build_runner_shape_config", buildRunnerShapeConfigArray)
		} else {
			s.D.Set("build_runner_shape_config", nil)
		}

		if v.PrivateAccessConfig != nil {
			privateAccessConfigArray := []interface{}{}
			if privateAccessConfigMap := NetworkChannelToMapForBuildStage(&v.PrivateAccessConfig, true); privateAccessConfigMap != nil {
				privateAccessConfigArray = append(privateAccessConfigArray, privateAccessConfigMap)
			}
			s.D.Set("private_access_config", privateAccessConfigArray)
		} else {
			s.D.Set("private_access_config", nil)
		}
	case oci_devops.DeliverArtifactStage:
		s.D.Set("build_pipeline_stage_type", "DELIVER_ARTIFACT")

		if v.DeliverArtifactCollection != nil {
			s.D.Set("deliver_artifact_collection", []interface{}{DeliverArtifactCollectionToMap(v.DeliverArtifactCollection)})
		} else {
			s.D.Set("deliver_artifact_collection", nil)
		}
	case oci_devops.TriggerDeploymentStage:
		s.D.Set("build_pipeline_stage_type", "TRIGGER_DEPLOYMENT_PIPELINE")

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", v.DeployPipelineId)
		}

		if v.IsPassAllParametersEnabled != nil {
			s.D.Set("is_pass_all_parameters_enabled", v.IsPassAllParametersEnabled)
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
	default:
		log.Printf("[WARN] Received 'connection_type' of unknown type %v", v)
		return nil
	}

	return nil
}
