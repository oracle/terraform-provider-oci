// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_anomaly_detection

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_ai_anomaly_detection "github.com/oracle/oci-go-sdk/v65/aianomalydetection"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AiAnomalyDetectionDetectAnomalyJobDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["detect_anomaly_job_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(AiAnomalyDetectionDetectAnomalyJobResource(), fieldMap, readSingularAiAnomalyDetectionDetectAnomalyJob)
}

func readSingularAiAnomalyDetectionDetectAnomalyJob(d *schema.ResourceData, m interface{}) error {
	sync := &AiAnomalyDetectionDetectAnomalyJobDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnomalyDetectionClient()

	return tfresource.ReadResource(sync)
}

type AiAnomalyDetectionDetectAnomalyJobDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ai_anomaly_detection.AnomalyDetectionClient
	Res    *oci_ai_anomaly_detection.GetDetectAnomalyJobResponse
}

func (s *AiAnomalyDetectionDetectAnomalyJobDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AiAnomalyDetectionDetectAnomalyJobDataSourceCrud) Get() error {
	request := oci_ai_anomaly_detection.GetDetectAnomalyJobRequest{}

	if detectAnomalyJobId, ok := s.D.GetOkExists("detect_anomaly_job_id"); ok {
		tmp := detectAnomalyJobId.(string)
		request.DetectAnomalyJobId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ai_anomaly_detection")

	response, err := s.Client.GetDetectAnomalyJob(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AiAnomalyDetectionDetectAnomalyJobDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
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

	if s.Res.InputDetails != nil {
		inputDetailsArray := []interface{}{}
		if inputDetailsMap := InputJobDetailsToMap(&s.Res.InputDetails); inputDetailsMap != nil {
			inputDetailsArray = append(inputDetailsArray, inputDetailsMap)
		}
		s.D.Set("input_details", inputDetailsArray)
	} else {
		s.D.Set("input_details", nil)
	}

	if s.Res.LifecycleStateDetails != nil {
		s.D.Set("lifecycle_state_details", *s.Res.LifecycleStateDetails)
	}

	if s.Res.ModelId != nil {
		s.D.Set("model_id", *s.Res.ModelId)
	}

	if s.Res.OutputDetails != nil {
		outputDetailsArray := []interface{}{}
		if outputDetailsMap := OutputJobDetailsToMap(&s.Res.OutputDetails); outputDetailsMap != nil {
			outputDetailsArray = append(outputDetailsArray, outputDetailsMap)
		}
		s.D.Set("output_details", outputDetailsArray)
	} else {
		s.D.Set("output_details", nil)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	if s.Res.Sensitivity != nil {
		s.D.Set("sensitivity", *s.Res.Sensitivity)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		//s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
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

	return nil
}
