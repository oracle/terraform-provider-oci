// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatasciencePipelineDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["pipeline_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatasciencePipelineResource(), fieldMap, readSingularDatasciencePipeline)
}

func readSingularDatasciencePipeline(d *schema.ResourceData, m interface{}) error {
	sync := &DatasciencePipelineDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

type DatasciencePipelineDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.GetPipelineResponse
}

func (s *DatasciencePipelineDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatasciencePipelineDataSourceCrud) Get() error {
	request := oci_datascience.GetPipelineRequest{}

	if pipelineId, ok := s.D.GetOkExists("pipeline_id"); ok {
		tmp := pipelineId.(string)
		request.PipelineId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datascience")

	response, err := s.Client.GetPipeline(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatasciencePipelineDataSourceCrud) SetData() error {
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

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InfrastructureConfigurationDetails != nil {
		s.D.Set("infrastructure_configuration_details", []interface{}{PipelineInfrastructureConfigurationDetailsToMap(s.Res.InfrastructureConfigurationDetails)})
	} else {
		s.D.Set("infrastructure_configuration_details", nil)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.LogConfigurationDetails != nil {
		s.D.Set("log_configuration_details", []interface{}{PipelineLogConfigurationDetailsToMap(s.Res.LogConfigurationDetails)})
	} else {
		s.D.Set("log_configuration_details", nil)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	stepDetails := []interface{}{}
	for _, item := range s.Res.StepDetails {
		stepDetails = append(stepDetails, PipelineStepDetailsToMap(item))
	}
	s.D.Set("step_details", stepDetails)

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
