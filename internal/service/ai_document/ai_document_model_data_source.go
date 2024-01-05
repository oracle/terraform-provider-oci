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

func AiDocumentModelDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["model_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(AiDocumentModelResource(), fieldMap, readSingularAiDocumentModel)
}

func readSingularAiDocumentModel(d *schema.ResourceData, m interface{}) error {
	sync := &AiDocumentModelDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceDocumentClient()

	return tfresource.ReadResource(sync)
}

type AiDocumentModelDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ai_document.AIServiceDocumentClient
	Res    *oci_ai_document.GetModelResponse
}

func (s *AiDocumentModelDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AiDocumentModelDataSourceCrud) Get() error {
	request := oci_ai_document.GetModelRequest{}

	if modelId, ok := s.D.GetOkExists("model_id"); ok {
		tmp := modelId.(string)
		request.ModelId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ai_document")

	response, err := s.Client.GetModel(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AiDocumentModelDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AliasName != nil {
		s.D.Set("alias_name", *s.Res.AliasName)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	componentModels := []interface{}{}
	for _, item := range s.Res.ComponentModels {
		componentModels = append(componentModels, ComponentModelToMap(item))
	}
	s.D.Set("component_models", componentModels)

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
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsComposedModel != nil {
		s.D.Set("is_composed_model", *s.Res.IsComposedModel)
	}

	if s.Res.IsQuickMode != nil {
		s.D.Set("is_quick_mode", *s.Res.IsQuickMode)
	}

	s.D.Set("labels", s.Res.Labels)
	s.D.Set("labels", s.Res.Labels)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MaxTrainingTimeInHours != nil {
		s.D.Set("max_training_time_in_hours", *s.Res.MaxTrainingTimeInHours)
	}

	if s.Res.Metrics != nil {
		metricsArray := []interface{}{}
		if metricsMap := ModelMetricsToMap(&s.Res.Metrics); metricsMap != nil {
			metricsArray = append(metricsArray, metricsMap)
		}
		s.D.Set("metrics", metricsArray)
	} else {
		s.D.Set("metrics", nil)
	}

	s.D.Set("model_type", s.Res.ModelType)

	if s.Res.ModelVersion != nil {
		s.D.Set("model_version", *s.Res.ModelVersion)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TenancyId != nil {
		s.D.Set("tenancy_id", *s.Res.TenancyId)
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

	if s.Res.TrainedTimeInHours != nil {
		s.D.Set("trained_time_in_hours", *s.Res.TrainedTimeInHours)
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
