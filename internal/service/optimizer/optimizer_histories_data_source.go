// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package optimizer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_optimizer "github.com/oracle/oci-go-sdk/v65/optimizer"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OptimizerHistoriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOptimizerHistories,
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
			"include_resource_metadata": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"recommendation_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"recommendation_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"history_collection": {
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
									"action": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"description": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"url": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"category_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compartment_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"estimated_cost_saving": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"extended_metadata": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"metadata": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"recommendation_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"recommendation_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_action_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
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

func readOptimizerHistories(d *schema.ResourceData, m interface{}) error {
	sync := &OptimizerHistoriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OptimizerClient()

	return tfresource.ReadResource(sync)
}

type OptimizerHistoriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_optimizer.OptimizerClient
	Res    *oci_optimizer.ListHistoriesResponse
}

func (s *OptimizerHistoriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OptimizerHistoriesDataSourceCrud) Get() error {
	request := oci_optimizer.ListHistoriesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if includeResourceMetadata, ok := s.D.GetOkExists("include_resource_metadata"); ok {
		tmp := includeResourceMetadata.(bool)
		request.IncludeResourceMetadata = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if recommendationId, ok := s.D.GetOkExists("recommendation_id"); ok {
		tmp := recommendationId.(string)
		request.RecommendationId = &tmp
	}

	if recommendationName, ok := s.D.GetOkExists("recommendation_name"); ok {
		tmp := recommendationName.(string)
		request.RecommendationName = &tmp
	}

	if resourceType, ok := s.D.GetOkExists("resource_type"); ok {
		tmp := resourceType.(string)
		request.ResourceType = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_optimizer.ListHistoriesLifecycleStateEnum(state.(string))
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_optimizer.ListHistoriesStatusEnum(status.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "optimizer")

	response, err := s.Client.ListHistories(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListHistories(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OptimizerHistoriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OptimizerHistoriesDataSource-", OptimizerHistoriesDataSource(), s.D))
	resources := []map[string]interface{}{}
	history := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, HistorySummaryToMap(item))
	}
	history["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OptimizerHistoriesDataSource().Schema["history_collection"].Elem.(*schema.Resource).Schema)
		history["items"] = items
	}

	resources = append(resources, history)
	if err := s.D.Set("history_collection", resources); err != nil {
		return err
	}

	return nil
}

func OptimizerActionToMap(obj *oci_optimizer.Action) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	result["type"] = string(obj.Type)

	if obj.Url != nil {
		result["url"] = string(*obj.Url)
	}

	return result
}

func HistorySummaryToMap(obj oci_optimizer.HistorySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Action != nil {
		result["action"] = []interface{}{OptimizerActionToMap(obj.Action)}
	}

	if obj.CategoryId != nil {
		result["category_id"] = string(*obj.CategoryId)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.CompartmentName != nil {
		result["compartment_name"] = string(*obj.CompartmentName)
	}

	if obj.EstimatedCostSaving != nil {
		result["estimated_cost_saving"] = float32(*obj.EstimatedCostSaving)
	}

	result["extended_metadata"] = tfresource.GenericMapToJsonMap(obj.ExtendedMetadata)

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["metadata"] = obj.Metadata

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.RecommendationId != nil {
		result["recommendation_id"] = string(*obj.RecommendationId)
	}

	if obj.RecommendationName != nil {
		result["recommendation_name"] = string(*obj.RecommendationName)
	}

	if obj.ResourceActionId != nil {
		result["resource_action_id"] = string(*obj.ResourceActionId)
	}

	if obj.ResourceId != nil {
		result["resource_id"] = string(*obj.ResourceId)
	}

	if obj.ResourceType != nil {
		result["resource_type"] = string(*obj.ResourceType)
	}

	result["state"] = string(obj.LifecycleState)

	result["status"] = string(obj.Status)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}
