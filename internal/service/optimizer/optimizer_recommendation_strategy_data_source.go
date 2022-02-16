//Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
//Licensed under the Mozilla Public License v2.0

package optimizer

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_optimizer "github.com/oracle/oci-go-sdk/v58/optimizer"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func OptimizerRecommendationStrategyDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularOptimizerRecommendationStrategy,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"recommendation_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"strategies": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"is_default": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"parameters_definition": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"default_value": {
													Type:     schema.TypeList,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem:     schema.TypeString,
												},
												"description": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"is_required": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"possible_values": {
													Type:     schema.TypeList,
													Computed: true,
													Elem:     schema.TypeString,
												},
												"type": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"strategy_name": {
										Type:     schema.TypeString,
										Computed: true,
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

func readSingularOptimizerRecommendationStrategy(d *schema.ResourceData, m interface{}) error {
	sync := &OptimizerRecommendationStrategyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OptimizerClient()

	return tfresource.ReadResource(sync)
}

type OptimizerRecommendationStrategyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_optimizer.OptimizerClient
	Res    *oci_optimizer.ListRecommendationStrategiesResponse
}

func (s *OptimizerRecommendationStrategyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OptimizerRecommendationStrategyDataSourceCrud) Get() error {
	request := oci_optimizer.ListRecommendationStrategiesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if recommendationName, ok := s.D.GetOkExists("recommendation_name"); ok {
		tmp := recommendationName.(string)
		request.RecommendationName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "optimizer")

	response, err := s.Client.ListRecommendationStrategies(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OptimizerRecommendationStrategyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OptimizerRecommendationStrategyDataSource-", OptimizerRecommendationStrategyDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, RecommendationStrategySummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func RecommendationStrategySummaryToMap(obj oci_optimizer.RecommendationStrategySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	strategies := []interface{}{}
	for _, item := range obj.Strategies {
		strategies = append(strategies, StrategyToMap(item))
	}
	result["strategies"] = strategies

	return result
}

func StrategyToMap(obj oci_optimizer.Strategy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsDefault != nil {
		result["is_default"] = bool(*obj.IsDefault)
	}

	parametersDefinition := []interface{}{}
	for _, item := range obj.ParametersDefinition {
		parametersDefinition = append(parametersDefinition, StrategyParameterToMap(item))
	}
	result["parameters_definition"] = parametersDefinition

	if obj.StrategyName != nil {
		result["strategy_name"] = string(*obj.StrategyName)
	}

	return result
}

func StrategyParameterToMap(obj oci_optimizer.StrategyParameter) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefaultValue != nil {
		result["default_value"] = []interface{}{fmt.Sprintf("%v", obj.DefaultValue)}
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.IsRequired != nil {
		result["is_required"] = bool(*obj.IsRequired)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	possibleValues := []interface{}{}
	for _, item := range obj.PossibleValues {
		possibleValues = append(possibleValues, fmt.Sprintf("%v", item))
	}
	result["possible_values"] = possibleValues

	result["type"] = string(obj.Type)

	return result
}
