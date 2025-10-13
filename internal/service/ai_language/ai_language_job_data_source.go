// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_language

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_ai_language "github.com/oracle/oci-go-sdk/v65/ailanguage"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AiLanguageJobDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["job_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(AiLanguageJobResource(), fieldMap, readSingularAiLanguageJob)
}

func readSingularAiLanguageJob(d *schema.ResourceData, m interface{}) error {
	sync := &AiLanguageJobDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceLanguageClient()

	return tfresource.ReadResource(sync)
}

type AiLanguageJobDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ai_language.AIServiceLanguageClient
	Res    *oci_ai_language.GetJobResponse
}

func (s *AiLanguageJobDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AiLanguageJobDataSourceCrud) Get() error {
	request := oci_ai_language.GetJobRequest{}

	if jobId, ok := s.D.GetOkExists("job_id"); ok {
		tmp := jobId.(string)
		request.JobId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ai_language")

	response, err := s.Client.GetJob(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AiLanguageJobDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CompletedDocuments != nil {
		s.D.Set("completed_documents", *s.Res.CompletedDocuments)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.FailedDocuments != nil {
		s.D.Set("failed_documents", *s.Res.FailedDocuments)
	}

	if s.Res.InputConfiguration != nil {
		s.D.Set("input_configuration", []interface{}{InputConfigurationToMap(s.Res.InputConfiguration)})
	} else {
		s.D.Set("input_configuration", nil)
	}

	if s.Res.InputLocation != nil {
		inputLocationArray := []interface{}{}
		if inputLocationMap := InputLocationToMap(&s.Res.InputLocation); inputLocationMap != nil {
			inputLocationArray = append(inputLocationArray, inputLocationMap)
		}
		s.D.Set("input_location", inputLocationArray)
	} else {
		s.D.Set("input_location", nil)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	modelMetadataDetails := []interface{}{}
	for _, item := range s.Res.ModelMetadataDetails {
		modelMetadataDetails = append(modelMetadataDetails, ModelMetadataDetailsToMap(item))
	}
	s.D.Set("model_metadata_details", modelMetadataDetails)

	if s.Res.OutputLocation != nil {
		s.D.Set("output_location", []interface{}{ObjectPrefixOutputLocationToMap(s.Res.OutputLocation)})
	} else {
		s.D.Set("output_location", nil)
	}

	if s.Res.PendingDocuments != nil {
		s.D.Set("pending_documents", *s.Res.PendingDocuments)
	}

	if s.Res.PercentComplete != nil {
		s.D.Set("percent_complete", *s.Res.PercentComplete)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeAccepted != nil {
		s.D.Set("time_accepted", s.Res.TimeAccepted.String())
	}

	if s.Res.TimeCompleted != nil {
		s.D.Set("time_completed", s.Res.TimeCompleted.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	if s.Res.TotalDocuments != nil {
		s.D.Set("total_documents", *s.Res.TotalDocuments)
	}

	if s.Res.TtlInDays != nil {
		s.D.Set("ttl_in_days", *s.Res.TtlInDays)
	}

	if s.Res.WarningsCount != nil {
		s.D.Set("warnings_count", *s.Res.WarningsCount)
	}

	return nil
}
