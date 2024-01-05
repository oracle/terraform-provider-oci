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

func OptimizerCategoryDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularOptimizerCategory,
		Schema: map[string]*schema.Schema{
			"category_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
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
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"recommendation_counts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"importance": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"resource_counts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularOptimizerCategory(d *schema.ResourceData, m interface{}) error {
	sync := &OptimizerCategoryDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OptimizerClient()

	return tfresource.ReadResource(sync)
}

type OptimizerCategoryDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_optimizer.OptimizerClient
	Res    *oci_optimizer.GetCategoryResponse
}

func (s *OptimizerCategoryDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OptimizerCategoryDataSourceCrud) Get() error {
	request := oci_optimizer.GetCategoryRequest{}

	if categoryId, ok := s.D.GetOkExists("category_id"); ok {
		tmp := categoryId.(string)
		request.CategoryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "optimizer")

	response, err := s.Client.GetCategory(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OptimizerCategoryDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CompartmentName != nil {
		s.D.Set("compartment_name", *s.Res.CompartmentName)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.EstimatedCostSaving != nil {
		s.D.Set("estimated_cost_saving", *s.Res.EstimatedCostSaving)
	}

	s.D.Set("extended_metadata", s.Res.ExtendedMetadata)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	recommendationCounts := []interface{}{}
	for _, item := range s.Res.RecommendationCounts {
		recommendationCounts = append(recommendationCounts, RecommendationCountToMap(item))
	}
	s.D.Set("recommendation_counts", recommendationCounts)

	resourceCounts := []interface{}{}
	for _, item := range s.Res.ResourceCounts {
		resourceCounts = append(resourceCounts, ResourceCountToMap(item))
	}
	s.D.Set("resource_counts", resourceCounts)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func CategorySummaryToMap(obj oci_optimizer.CategorySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.EstimatedCostSaving != nil {
		result["estimated_cost_saving"] = float32(*obj.EstimatedCostSaving)
	}

	result["extended_metadata"] = obj.ExtendedMetadata

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	recommendationCounts := []interface{}{}
	for _, item := range obj.RecommendationCounts {
		recommendationCounts = append(recommendationCounts, RecommendationCountToMap(item))
	}
	result["recommendation_counts"] = recommendationCounts

	resourceCounts := []interface{}{}
	for _, item := range obj.ResourceCounts {
		resourceCounts = append(resourceCounts, ResourceCountToMap(item))
	}
	result["resource_counts"] = resourceCounts

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func RecommendationCountToMap(obj oci_optimizer.RecommendationCount) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Count != nil {
		result["count"] = int(*obj.Count)
	}

	result["importance"] = string(obj.Importance)

	return result
}

func ResourceCountToMap(obj oci_optimizer.ResourceCount) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Count != nil {
		result["count"] = int(*obj.Count)
	}

	result["status"] = string(obj.Status)

	return result
}
