// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_anomaly_detection

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_ai_anomaly_detection "github.com/oracle/oci-go-sdk/v65/aianomalydetection"
)

func AiAnomalyDetectionModelDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["model_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(AiAnomalyDetectionModelResource(), fieldMap, readSingularAiAnomalyDetectionModel)
}

func readSingularAiAnomalyDetectionModel(d *schema.ResourceData, m interface{}) error {
	sync := &AiAnomalyDetectionModelDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnomalyDetectionClient()

	return tfresource.ReadResource(sync)
}

type AiAnomalyDetectionModelDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ai_anomaly_detection.AnomalyDetectionClient
	Res    *oci_ai_anomaly_detection.GetModelResponse
}

func (s *AiAnomalyDetectionModelDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AiAnomalyDetectionModelDataSourceCrud) Get() error {
	request := oci_ai_anomaly_detection.GetModelRequest{}

	if modelId, ok := s.D.GetOkExists("model_id"); ok {
		tmp := modelId.(string)
		request.ModelId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ai_anomaly_detection")

	response, err := s.Client.GetModel(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AiAnomalyDetectionModelDataSourceCrud) SetData() error {
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

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ModelTrainingDetails != nil {
		s.D.Set("model_training_details", []interface{}{ModelTrainingDetailsToMap(s.Res.ModelTrainingDetails)})
	} else {
		s.D.Set("model_training_details", nil)
	}

	if s.Res.ModelTrainingResults != nil {
		s.D.Set("model_training_results", []interface{}{ModelTrainingResultsToMap(s.Res.ModelTrainingResults)})
	} else {
		s.D.Set("model_training_results", nil)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	// if s.Res.SystemTags != nil {
	// 	s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	// }

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
