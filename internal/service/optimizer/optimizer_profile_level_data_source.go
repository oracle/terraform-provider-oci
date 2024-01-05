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

func OptimizerProfileLevelDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularOptimizerProfileLevel,
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
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
					},
				},
			},
		},
		DeprecationMessage: tfresource.DatasourceDeprecatedForAnother("oci_optimizer_profile_level", "oci_optimizer_profile_levels"),
	}
}

func readSingularOptimizerProfileLevel(d *schema.ResourceData, m interface{}) error {
	sync := &OptimizerProfileLevelDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OptimizerClient()

	return tfresource.ReadResource(sync)
}

type OptimizerProfileLevelDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_optimizer.OptimizerClient
	Res    *oci_optimizer.ListProfileLevelsResponse
}

func (s *OptimizerProfileLevelDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OptimizerProfileLevelDataSourceCrud) Get() error {
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
	return nil
}

func (s *OptimizerProfileLevelDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OptimizerProfileLevelDataSource-", OptimizerProfileLevelDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ProfileLevelSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func optimizerEvaluatedMetricToMap(obj oci_optimizer.EvaluatedMetric) map[string]interface{} {
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

func optimizerProfileLevelSummaryToMap(obj oci_optimizer.ProfileLevelSummary) map[string]interface{} {
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
