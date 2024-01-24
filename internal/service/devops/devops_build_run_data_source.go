// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_devops "github.com/oracle/oci-go-sdk/v65/devops"
)

func DevopsBuildRunDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["build_run_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DevopsBuildRunResource(), fieldMap, readSingularDevopsBuildRun)
}

func readSingularDevopsBuildRun(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsBuildRunDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

type DevopsBuildRunDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.GetBuildRunResponse
}

func (s *DevopsBuildRunDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsBuildRunDataSourceCrud) Get() error {
	request := oci_devops.GetBuildRunRequest{}

	if buildRunId, ok := s.D.GetOkExists("build_run_id"); ok {
		tmp := buildRunId.(string)
		request.BuildRunId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "devops")

	response, err := s.Client.GetBuildRun(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DevopsBuildRunDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.BuildOutputs != nil {
		s.D.Set("build_outputs", []interface{}{BuildOutputsToMap(s.Res.BuildOutputs)})
	} else {
		s.D.Set("build_outputs", nil)
	}

	if s.Res.BuildPipelineId != nil {
		s.D.Set("build_pipeline_id", *s.Res.BuildPipelineId)
	}

	if s.Res.BuildRunArguments != nil {
		s.D.Set("build_run_arguments", []interface{}{BuildRunArgumentCollectionToMap(s.Res.BuildRunArguments)})
	} else {
		s.D.Set("build_run_arguments", nil)
	}

	if s.Res.BuildRunProgress != nil {
		s.D.Set("build_run_progress", []interface{}{BuildRunProgressToMap(s.Res.BuildRunProgress)})
	} else {
		s.D.Set("build_run_progress", nil)
	}

	if s.Res.BuildRunSource != nil {
		buildRunSourceArray := []interface{}{}
		if buildRunSourceMap := BuildRunSourceToMap(&s.Res.BuildRunSource); buildRunSourceMap != nil {
			buildRunSourceArray = append(buildRunSourceArray, buildRunSourceMap)
		}
		s.D.Set("build_run_source", buildRunSourceArray)
	} else {
		s.D.Set("build_run_source", nil)
	}

	if s.Res.CommitInfo != nil {
		s.D.Set("commit_info", []interface{}{CommitInfoToMap(s.Res.CommitInfo)})
	} else {
		s.D.Set("commit_info", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
