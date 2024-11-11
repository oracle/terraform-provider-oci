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

func JmsPluginErrorAnalyticsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readJmsPluginErrorAnalytics,
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
			"plugin_error_aggregation_collection": {
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
									"healthy_plugin_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"plugin_error_aggregations": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"plugin_error_analytic_count": {
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
								},
							},
						},
					},
				},
			},
		},
	}
}

func readJmsPluginErrorAnalytics(d *schema.ResourceData, m interface{}) error {
	sync := &JmsPluginErrorAnalyticsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsPluginErrorAnalyticsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.SummarizePluginErrorsResponse
}

func (s *JmsPluginErrorAnalyticsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsPluginErrorAnalyticsDataSourceCrud) Get() error {
	request := oci_jms.SummarizePluginErrorsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms")

	response, err := s.Client.SummarizePluginErrors(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.SummarizePluginErrors(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *JmsPluginErrorAnalyticsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsPluginErrorAnalyticsDataSource-", JmsPluginErrorAnalyticsDataSource(), s.D))
	resources := []map[string]interface{}{}
	pluginErrorAnalytic := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, PluginErrorAggregationSummaryToMap(item))
	}
	pluginErrorAnalytic["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, JmsPluginErrorAnalyticsDataSource().Schema["plugin_error_aggregation_collection"].Elem.(*schema.Resource).Schema)
		pluginErrorAnalytic["items"] = items
	}

	resources = append(resources, pluginErrorAnalytic)
	if err := s.D.Set("plugin_error_aggregation_collection", resources); err != nil {
		return err
	}

	return nil
}

func PluginErrorAggregationToMap(obj oci_jms.PluginErrorAggregation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Count != nil {
		result["plugin_error_analytic_count"] = int(*obj.Count)
	}

	result["reason"] = string(obj.Reason)

	return result
}

func PluginErrorAggregationSummaryToMap(obj oci_jms.PluginErrorAggregationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.HealthyPluginCount != nil {
		result["healthy_plugin_count"] = int(*obj.HealthyPluginCount)
	}

	pluginErrorAggregations := []interface{}{}
	for _, item := range obj.PluginErrorAggregations {
		pluginErrorAggregations = append(pluginErrorAggregations, PluginErrorAggregationToMap(item))
	}
	result["plugin_error_aggregations"] = pluginErrorAggregations

	return result
}
