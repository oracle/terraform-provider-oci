// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_jms "github.com/oracle/oci-go-sdk/v65/jms"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsFleetErrorAnalyticsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readJmsFleetErrorAnalytics,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"fleet_error_aggregation_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"fleet_error_aggregations": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"fleet_error_analytic_count": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"reason": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"healthy_fleet_count": {
										Type:     schema.TypeInt,
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

func readJmsFleetErrorAnalytics(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetErrorAnalyticsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsFleetErrorAnalyticsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.SummarizeFleetErrorsResponse
}

func (s *JmsFleetErrorAnalyticsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsFleetErrorAnalyticsDataSourceCrud) Get() error {
	request := oci_jms.SummarizeFleetErrorsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms")

	response, err := s.Client.SummarizeFleetErrors(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.SummarizeFleetErrors(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *JmsFleetErrorAnalyticsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsFleetErrorAnalyticsDataSource-", JmsFleetErrorAnalyticsDataSource(), s.D))
	resources := []map[string]interface{}{}
	fleetErrorAnalytic := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, FleetErrorAggregationSummaryToMap(item))
	}
	fleetErrorAnalytic["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, JmsFleetErrorAnalyticsDataSource().Schema["fleet_error_aggregation_collection"].Elem.(*schema.Resource).Schema)
		fleetErrorAnalytic["items"] = items
	}

	resources = append(resources, fleetErrorAnalytic)
	if err := s.D.Set("fleet_error_aggregation_collection", resources); err != nil {
		return err
	}

	return nil
}

func FleetErrorAggregationToMap(obj oci_jms.FleetErrorAggregation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Count != nil {
		result["fleet_error_analytic_count"] = int(*obj.Count)
	}

	result["reason"] = string(obj.Reason)

	return result
}

func FleetErrorAggregationSummaryToMap(obj oci_jms.FleetErrorAggregationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	fleetErrorAggregations := []interface{}{}
	for _, item := range obj.FleetErrorAggregations {
		fleetErrorAggregations = append(fleetErrorAggregations, FleetErrorAggregationToMap(item))
	}
	result["fleet_error_aggregations"] = fleetErrorAggregations

	if obj.HealthyFleetCount != nil {
		result["healthy_fleet_count"] = int(*obj.HealthyFleetCount)
	}

	return result
}
