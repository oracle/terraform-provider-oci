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

func AiLanguageModelDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(AiLanguageModelResource(), fieldMap, readSingularAiLanguageModel)
}

func readSingularAiLanguageModel(d *schema.ResourceData, m interface{}) error {
	sync := &AiLanguageModelDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceLanguageClient()

	return tfresource.ReadResource(sync)
}

type AiLanguageModelDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ai_language.AIServiceLanguageClient
	Res    *oci_ai_language.GetModelResponse
}

func (s *AiLanguageModelDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AiLanguageModelDataSourceCrud) Get() error {
	request := oci_ai_language.GetModelRequest{}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.ModelId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ai_language")

	response, err := s.Client.GetModel(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AiLanguageModelDataSourceCrud) SetData() error {
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

	if s.Res.EvaluationResults != nil {
		evaluationResultsArray := []interface{}{}
		if evaluationResultsMap := EvaluationResultsToMap(&s.Res.EvaluationResults); evaluationResultsMap != nil {
			evaluationResultsArray = append(evaluationResultsArray, evaluationResultsMap)
		}
		s.D.Set("evaluation_results", evaluationResultsArray)
	} else {
		s.D.Set("evaluation_results", nil)
	}

	if s.Res.FreeformTags != nil {
		s.D.Set("freeform_tags", s.Res.FreeformTags)
	}
	// s.D.Set("freeform_tags", s.Res.FreeformTags)
	// s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ModelDetails != nil {
		modelDetailsArray := []interface{}{}
		if modelDetailsMap := ModelDetailsToMap(&s.Res.ModelDetails); modelDetailsMap != nil {
			modelDetailsArray = append(modelDetailsArray, modelDetailsMap)
		}
		s.D.Set("model_details", modelDetailsArray)
	} else {
		s.D.Set("model_details", nil)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", s.Res.SystemTags)
	}

	if s.Res.TestStrategy != nil {
		testStrategyArray := []interface{}{}
		if testStrategyMap := TestStrategyToMap(&s.Res.TestStrategy); testStrategyMap != nil {
			testStrategyArray = append(testStrategyArray, testStrategyMap)
		}
		s.D.Set("test_strategy", testStrategyArray)
	} else {
		s.D.Set("test_strategy", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TrainingDataset != nil {
		trainingDatasetArray := []interface{}{}
		if trainingDatasetMap := DatasetDetailsToMap(&s.Res.TrainingDataset); trainingDatasetMap != nil {
			trainingDatasetArray = append(trainingDatasetArray, trainingDatasetMap)
		}
		s.D.Set("training_dataset", trainingDatasetArray)
	} else {
		s.D.Set("training_dataset", nil)
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}
