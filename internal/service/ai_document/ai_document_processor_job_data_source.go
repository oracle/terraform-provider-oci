// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_document

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_ai_document "github.com/oracle/oci-go-sdk/v65/aidocument"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AiDocumentProcessorJobDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["processor_job_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(AiDocumentProcessorJobResource(), fieldMap, readSingularAiDocumentProcessorJob)
}

func readSingularAiDocumentProcessorJob(d *schema.ResourceData, m interface{}) error {
	sync := &AiDocumentProcessorJobDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceDocumentClient()

	return tfresource.ReadResource(sync)
}

type AiDocumentProcessorJobDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ai_document.AIServiceDocumentClient
	Res    *oci_ai_document.GetProcessorJobResponse
}

func (s *AiDocumentProcessorJobDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AiDocumentProcessorJobDataSourceCrud) Get() error {
	request := oci_ai_document.GetProcessorJobRequest{}

	if processorJobId, ok := s.D.GetOkExists("processor_job_id"); ok {
		tmp := processorJobId.(string)
		request.ProcessorJobId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ai_document")

	response, err := s.Client.GetProcessorJob(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AiDocumentProcessorJobDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
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

	s.D.Set("lifecycle_details", s.Res.LifecycleDetails)

	if s.Res.OutputLocation != nil {
		s.D.Set("output_location", []interface{}{OutputLocationToMap(s.Res.OutputLocation)})
	} else {
		s.D.Set("output_location", nil)
	}

	if s.Res.PercentComplete != nil {
		s.D.Set("percent_complete", *s.Res.PercentComplete)
	}

	if s.Res.ProcessorConfig != nil {
		processorConfigArray := []interface{}{}
		if processorConfigMap := ProcessorConfigToMap(&s.Res.ProcessorConfig); processorConfigMap != nil {
			processorConfigArray = append(processorConfigArray, processorConfigMap)
		}
		s.D.Set("processor_config", processorConfigArray)
	} else {
		s.D.Set("processor_config", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeAccepted != nil {
		s.D.Set("time_accepted", s.Res.TimeAccepted.String())
	}

	if s.Res.TimeFinished != nil {
		s.D.Set("time_finished", s.Res.TimeFinished.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	return nil
}
