// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package optimizer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_optimizer "github.com/oracle/oci-go-sdk/v56/optimizer"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func OptimizerProfileLevelsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOptimizerProfileLevels,
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
			"profile_level_collection": {
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
									"default_interval": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"metrics": {
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
												"statistic": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"target": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"threshold": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
											},
										},
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"recommendation_name": {
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
									"valid_intervals": {
										Type:     schema.TypeList,
										Computed: true,
										MinItems: 1,
										Elem: &schema.Schema{
											Type: schema.TypeInt,
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

func readOptimizerProfileLevels(d *schema.ResourceData, m interface{}) error {
	sync := &OptimizerProfileLevelsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OptimizerClient()

	return tfresource.ReadResource(sync)
}

type OptimizerProfileLevelsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_optimizer.OptimizerClient
	Res    *oci_optimizer.ListProfileLevelsResponse
}

func (s *OptimizerProfileLevelsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OptimizerProfileLevelsDataSourceCrud) Get() error {
	request := oci_optimizer.ListProfileLevelsRequest{}

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

	response, err := s.Client.ListProfileLevels(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListProfileLevels(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OptimizerProfileLevelsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OptimizerProfileLevelsDataSource-", OptimizerProfileLevelsDataSource(), s.D))
	resources := []map[string]interface{}{}
	profileLevel := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ProfileLevelSummaryToMap(item))
	}
	profileLevel["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OptimizerProfileLevelsDataSource().Schema["profile_level_collection"].Elem.(*schema.Resource).Schema)
		profileLevel["items"] = items
	}

	resources = append(resources, profileLevel)
	if err := s.D.Set("profile_level_collection", resources); err != nil {
		return err
	}

	return nil
}

func EvaluatedMetricToMap(obj oci_optimizer.EvaluatedMetric) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Statistic != nil {
		result["statistic"] = string(*obj.Statistic)
	}

	if obj.Target != nil {
		result["target"] = float64(*obj.Target)
	}

	if obj.Threshold != nil {
		result["threshold"] = float64(*obj.Threshold)
	}

	return result
}

func ProfileLevelSummaryToMap(obj oci_optimizer.ProfileLevelSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefaultInterval != nil {
		result["default_interval"] = int(*obj.DefaultInterval)
	}

	metrics := []interface{}{}
	for _, item := range obj.Metrics {
		metrics = append(metrics, EvaluatedMetricToMap(item))
	}
	result["metrics"] = metrics

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.RecommendationName != nil {
		result["recommendation_name"] = string(*obj.RecommendationName)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	result["valid_intervals"] = obj.ValidIntervals

	return result
}
