// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai_agent

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_generative_ai_agent "github.com/oracle/oci-go-sdk/v65/generativeaiagent"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GenerativeAiAgentDataIngestionJobDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["data_ingestion_job_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(GenerativeAiAgentDataIngestionJobResource(), fieldMap, readSingularGenerativeAiAgentDataIngestionJob)
}

func readSingularGenerativeAiAgentDataIngestionJob(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiAgentDataIngestionJobDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiAgentClient()

	return tfresource.ReadResource(sync)
}

type GenerativeAiAgentDataIngestionJobDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_generative_ai_agent.GenerativeAiAgentClient
	Res    *oci_generative_ai_agent.GetDataIngestionJobResponse
}

func (s *GenerativeAiAgentDataIngestionJobDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GenerativeAiAgentDataIngestionJobDataSourceCrud) Get() error {
	request := oci_generative_ai_agent.GetDataIngestionJobRequest{}

	if dataIngestionJobId, ok := s.D.GetOkExists("data_ingestion_job_id"); ok {
		tmp := dataIngestionJobId.(string)
		request.DataIngestionJobId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "generative_ai_agent")

	response, err := s.Client.GetDataIngestionJob(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *GenerativeAiAgentDataIngestionJobDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DataIngestionJobStatistics != nil {
		s.D.Set("data_ingestion_job_statistics", []interface{}{DataIngestionJobStatisticsToMap(s.Res.DataIngestionJobStatistics)})
	} else {
		s.D.Set("data_ingestion_job_statistics", nil)
	}

	if s.Res.DataSourceId != nil {
		s.D.Set("data_source_id", *s.Res.DataSourceId)
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

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
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
