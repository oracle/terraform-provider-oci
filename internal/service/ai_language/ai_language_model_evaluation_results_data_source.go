// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_language

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_ai_language "github.com/oracle/oci-go-sdk/v65/ailanguage"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AiLanguageModelEvaluationResultsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readAiLanguageModelEvaluationResults,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"model_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"evaluation_result_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"location": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"model_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"predicted_entities": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"length": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"offset": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"type": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"predicted_labels": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"record": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"true_entities": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"length": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"offset": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"type": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"true_labels": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readAiLanguageModelEvaluationResults(d *schema.ResourceData, m interface{}) error {
	sync := &AiLanguageModelEvaluationResultsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceLanguageClient()

	return tfresource.ReadResource(sync)
}

type AiLanguageModelEvaluationResultsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ai_language.AIServiceLanguageClient
	Res    *oci_ai_language.ListEvaluationResultsResponse
}

func (s *AiLanguageModelEvaluationResultsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AiLanguageModelEvaluationResultsDataSourceCrud) Get() error {
	request := oci_ai_language.ListEvaluationResultsRequest{}

	if modelId, ok := s.D.GetOkExists("model_id"); ok {
		tmp := modelId.(string)
		request.ModelId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ai_language")

	response, err := s.Client.ListEvaluationResults(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListEvaluationResults(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *AiLanguageModelEvaluationResultsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("AiLanguageModelEvaluationResultsDataSource-", AiLanguageModelEvaluationResultsDataSource(), s.D))
	resources := []map[string]interface{}{}
	modelEvaluationResult := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, EvaluationResultSummaryToMap(item))
	}
	modelEvaluationResult["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, AiLanguageModelEvaluationResultsDataSource().Schema["evaluation_result_collection"].Elem.(*schema.Resource).Schema)
		modelEvaluationResult["items"] = items
	}

	resources = append(resources, modelEvaluationResult)
	if err := s.D.Set("evaluation_result_collection", resources); err != nil {
		return err
	}

	return nil
}

func EntityLabelErrorAnalysisToMap(obj oci_ai_language.EntityLabelErrorAnalysis) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Length != nil {
		result["length"] = int(*obj.Length)
	}

	if obj.Offset != nil {
		result["offset"] = int(*obj.Offset)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}

func EvaluationResultSummaryToMap(obj oci_ai_language.EvaluationResultSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_ai_language.NamedEntityRecognitionEvaluationResult:
		result["model_type"] = "NAMED_ENTITY_RECOGNITION"

		predictedEntities := []interface{}{}
		for _, item := range v.PredictedEntities {
			predictedEntities = append(predictedEntities, EntityLabelErrorAnalysisToMap(item))
		}
		result["predicted_entities"] = predictedEntities

		if v.Record != nil {
			result["record"] = string(*v.Record)
		}

		trueEntities := []interface{}{}
		for _, item := range v.TrueEntities {
			trueEntities = append(trueEntities, EntityLabelErrorAnalysisToMap(item))
		}
		result["true_entities"] = trueEntities

		if v.DefinedTags != nil {
			result["defined_tags"] = tfresource.DefinedTagsToMap(v.DefinedTags)
		}

		result["freeform_tags"] = v.FreeformTags
	case oci_ai_language.TextClassificationModelEvaluationResult:
		result["model_type"] = "TEXT_CLASSIFICATION"

		if v.Location != nil {
			result["location"] = string(*v.Location)
		}

		result["predicted_labels"] = v.PredictedLabels

		result["true_labels"] = v.TrueLabels

		if v.DefinedTags != nil {
			result["defined_tags"] = tfresource.DefinedTagsToMap(v.DefinedTags)
		}

		result["freeform_tags"] = v.FreeformTags
	default:
		log.Printf("[WARN] Received 'model_type' of unknown type %v", obj)
		return nil
	}

	return result
}
