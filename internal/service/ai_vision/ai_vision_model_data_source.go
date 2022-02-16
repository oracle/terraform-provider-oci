// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_vision

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_ai_vision "github.com/oracle/oci-go-sdk/v58/aivision"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func AiVisionModelDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["model_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(AiVisionModelResource(), fieldMap, readSingularAiVisionModel)
}

func readSingularAiVisionModel(d *schema.ResourceData, m interface{}) error {
	sync := &AiVisionModelDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceVisionClient()

	return tfresource.ReadResource(sync)
}

type AiVisionModelDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ai_vision.AIServiceVisionClient
	Res    *oci_ai_vision.GetModelResponse
}

func (s *AiVisionModelDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AiVisionModelDataSourceCrud) Get() error {
	request := oci_ai_vision.GetModelRequest{}

	if modelId, ok := s.D.GetOkExists("model_id"); ok {
		tmp := modelId.(string)
		request.ModelId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ai_vision")

	response, err := s.Client.GetModel(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AiVisionModelDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AveragePrecision != nil {
		s.D.Set("average_precision", *s.Res.AveragePrecision)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConfidenceThreshold != nil {
		s.D.Set("confidence_threshold", *s.Res.ConfidenceThreshold)
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

	if s.Res.IsQuickMode != nil {
		s.D.Set("is_quick_mode", *s.Res.IsQuickMode)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MaxTrainingDurationInHours != nil {
		s.D.Set("max_training_duration_in_hours", *s.Res.MaxTrainingDurationInHours)
	}

	if s.Res.Metrics != nil {
		s.D.Set("metrics", *s.Res.Metrics)
	}

	s.D.Set("model_type", s.Res.ModelType)

	if s.Res.ModelVersion != nil {
		s.D.Set("model_version", *s.Res.ModelVersion)
	}

	if s.Res.Precision != nil {
		s.D.Set("precision", *s.Res.Precision)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	if s.Res.Recall != nil {
		s.D.Set("recall", *s.Res.Recall)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TestImageCount != nil {
		s.D.Set("test_image_count", *s.Res.TestImageCount)
	}

	if s.Res.TestingDataset != nil {
		testingDatasetArray := []interface{}{}
		if testingDatasetMap := DatasetToMap(&s.Res.TestingDataset); testingDatasetMap != nil {
			testingDatasetArray = append(testingDatasetArray, testingDatasetMap)
		}
		s.D.Set("testing_dataset", testingDatasetArray)
	} else {
		s.D.Set("testing_dataset", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TotalImageCount != nil {
		s.D.Set("total_image_count", *s.Res.TotalImageCount)
	}

	if s.Res.TrainedDurationInHours != nil {
		s.D.Set("trained_duration_in_hours", *s.Res.TrainedDurationInHours)
	}

	if s.Res.TrainingDataset != nil {
		trainingDatasetArray := []interface{}{}
		if trainingDatasetMap := DatasetToMap(&s.Res.TrainingDataset); trainingDatasetMap != nil {
			trainingDatasetArray = append(trainingDatasetArray, trainingDatasetMap)
		}
		s.D.Set("training_dataset", trainingDatasetArray)
	} else {
		s.D.Set("training_dataset", nil)
	}

	if s.Res.ValidationDataset != nil {
		validationDatasetArray := []interface{}{}
		if validationDatasetMap := DatasetToMap(&s.Res.ValidationDataset); validationDatasetMap != nil {
			validationDatasetArray = append(validationDatasetArray, validationDatasetMap)
		}
		s.D.Set("validation_dataset", validationDatasetArray)
	} else {
		s.D.Set("validation_dataset", nil)
	}

	return nil
}
