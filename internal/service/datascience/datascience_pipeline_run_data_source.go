// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatasciencePipelineRunDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["pipeline_run_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatasciencePipelineRunResource(), fieldMap, readSingularDatasciencePipelineRun)
}

func readSingularDatasciencePipelineRun(d *schema.ResourceData, m interface{}) error {
	sync := &DatasciencePipelineRunDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

type DatasciencePipelineRunDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.GetPipelineRunResponse
}

func (s *DatasciencePipelineRunDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatasciencePipelineRunDataSourceCrud) Get() error {
	request := oci_datascience.GetPipelineRunRequest{}

	if pipelineRunId, ok := s.D.GetOkExists("pipeline_run_id"); ok {
		tmp := pipelineRunId.(string)
		request.PipelineRunId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datascience")

	response, err := s.Client.GetPipelineRun(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatasciencePipelineRunDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConfigurationDetails != nil {
		configurationDetailsArray := []interface{}{}
		if configurationDetailsMap := PipelineConfigurationDetailsToMap(&s.Res.ConfigurationDetails); configurationDetailsMap != nil {
			configurationDetailsArray = append(configurationDetailsArray, configurationDetailsMap)
		}
		s.D.Set("configuration_details", configurationDetailsArray)
	} else {
		s.D.Set("configuration_details", nil)
	}

	if s.Res.ConfigurationOverrideDetails != nil {
		configurationOverrideDetailsArray := []interface{}{}
		if configurationOverrideDetailsMap := PipelineConfigurationDetailsToMap(&s.Res.ConfigurationOverrideDetails); configurationOverrideDetailsMap != nil {
			configurationOverrideDetailsArray = append(configurationOverrideDetailsArray, configurationOverrideDetailsMap)
		}
		s.D.Set("configuration_override_details", configurationOverrideDetailsArray)
	} else {
		s.D.Set("configuration_override_details", nil)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
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

	if s.Res.LogConfigurationOverrideDetails != nil {
		s.D.Set("log_configuration_override_details", []interface{}{PipelineLogConfigurationDetailsToMap(s.Res.LogConfigurationOverrideDetails)})
	} else {
		s.D.Set("log_configuration_override_details", nil)
	}

	if s.Res.LogDetails != nil {
		s.D.Set("log_details", []interface{}{PipelineRunLogDetailsToMap(s.Res.LogDetails)})
	} else {
		s.D.Set("log_details", nil)
	}

	if s.Res.PipelineId != nil {
		s.D.Set("pipeline_id", *s.Res.PipelineId)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	stepOverrideDetails := []interface{}{}
	for _, item := range s.Res.StepOverrideDetails {
		stepOverrideDetails = append(stepOverrideDetails, PipelineStepOverrideDetailsToMap(item))
	}
	s.D.Set("step_override_details", stepOverrideDetails)

	stepRuns := []interface{}{}
	for _, item := range s.Res.StepRuns {
		stepRuns = append(stepRuns, PipelineStepRunToMap(item))
	}
	s.D.Set("step_runs", stepRuns)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeAccepted != nil {
		s.D.Set("time_accepted", s.Res.TimeAccepted.String())
	}

	if s.Res.TimeFinished != nil {
		s.D.Set("time_finished", s.Res.TimeFinished.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
