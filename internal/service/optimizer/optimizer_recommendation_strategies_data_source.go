// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package optimizer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_optimizer "github.com/oracle/oci-go-sdk/v58/optimizer"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func OptimizerRecommendationStrategiesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOptimizerRecommendationStrategies,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
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
			"recommendation_strategy_collection": {
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
				},
			},
		},
	}
}

func readOptimizerRecommendationStrategies(d *schema.ResourceData, m interface{}) error {
	sync := &OptimizerRecommendationStrategiesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OptimizerClient()

	return tfresource.ReadResource(sync)
}

type OptimizerRecommendationStrategiesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_optimizer.OptimizerClient
	Res    *oci_optimizer.ListRecommendationStrategiesResponse
}

func (s *OptimizerRecommendationStrategiesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OptimizerRecommendationStrategiesDataSourceCrud) Get() error {
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
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListRecommendationStrategies(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OptimizerRecommendationStrategiesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OptimizerRecommendationStrategiesDataSource-", OptimizerRecommendationStrategiesDataSource(), s.D))
	resources := []map[string]interface{}{}
	recommendationStrategy := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, RecommendationStrategySummaryToMap(item))
	}
	recommendationStrategy["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OptimizerRecommendationStrategiesDataSource().Schema["recommendation_strategy_collection"].Elem.(*schema.Resource).Schema)
		recommendationStrategy["items"] = items
	}

	resources = append(resources, recommendationStrategy)
	if err := s.D.Set("recommendation_strategy_collection", resources); err != nil {
		return err
	}

	return nil
}
