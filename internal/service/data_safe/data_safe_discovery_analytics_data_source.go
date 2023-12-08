// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeDiscoveryAnalyticsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeDiscoveryAnalytics,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"group_by": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_common": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"sensitive_data_model_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sensitive_type_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"discovery_analytics_collection": {
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
									// Required

									// Optional

									// Computed
									"count": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"dimensions": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"sensitive_data_model_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"sensitive_type_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"target_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"metric_name": {
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

func readDataSafeDiscoveryAnalytics(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeDiscoveryAnalyticsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeDiscoveryAnalyticsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListDiscoveryAnalyticsResponse
}

func (s *DataSafeDiscoveryAnalyticsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeDiscoveryAnalyticsDataSourceCrud) Get() error {
	request := oci_data_safe.ListDiscoveryAnalyticsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if groupBy, ok := s.D.GetOkExists("group_by"); ok {
		request.GroupBy = oci_data_safe.ListDiscoveryAnalyticsGroupByEnum(groupBy.(string))
	}

	if isCommon, ok := s.D.GetOkExists("is_common"); ok {
		tmp := isCommon.(bool)
		request.IsCommon = &tmp
	}

	if sensitiveDataModelId, ok := s.D.GetOkExists("sensitive_data_model_id"); ok {
		tmp := sensitiveDataModelId.(string)
		request.SensitiveDataModelId = &tmp
	}

	if sensitiveTypeId, ok := s.D.GetOkExists("sensitive_type_id"); ok {
		tmp := sensitiveTypeId.(string)
		request.SensitiveTypeId = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListDiscoveryAnalytics(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDiscoveryAnalytics(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeDiscoveryAnalyticsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeDiscoveryAnalyticsDataSource-", DataSafeDiscoveryAnalyticsDataSource(), s.D))
	resources := []map[string]interface{}{}
	discoveryAnalytic := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DiscoveryAnalyticsSummaryToMap(item))
	}
	discoveryAnalytic["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeDiscoveryAnalyticsDataSource().Schema["discovery_analytics_collection"].Elem.(*schema.Resource).Schema)
		discoveryAnalytic["items"] = items
	}

	resources = append(resources, discoveryAnalytic)
	if err := s.D.Set("discovery_analytics_collection", resources); err != nil {
		return err
	}

	return nil
}

func DataSafeDimensionsToMap(obj *oci_data_safe.Dimensions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.SensitiveDataModelId != nil {
		result["sensitive_data_model_id"] = string(*obj.SensitiveDataModelId)
	}

	if obj.SensitiveTypeId != nil {
		result["sensitive_type_id"] = string(*obj.SensitiveTypeId)
	}

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	return result
}

func DataSafeDiscoveryAnalyticsSummaryToMap(obj oci_data_safe.DiscoveryAnalyticsSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Count != nil {
		result["count"] = strconv.FormatInt(*obj.Count, 10)
	}

	if obj.Dimensions != nil {
		result["dimensions"] = []interface{}{DimensionsToMap(obj.Dimensions)}
	}

	result["metric_name"] = string(obj.MetricName)

	return result
}

func DimensionsToMap(obj *oci_data_safe.Dimensions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.SensitiveDataModelId != nil {
		result["sensitive_data_model_id"] = string(*obj.SensitiveDataModelId)
	}

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	return result
}

func DiscoveryAnalyticsSummaryToMap(obj oci_data_safe.DiscoveryAnalyticsSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Count != nil {
		result["count"] = strconv.FormatInt(*obj.Count, 10)
	}

	if obj.Dimensions != nil {
		result["dimensions"] = []interface{}{DimensionsToMap(obj.Dimensions)}
	}

	result["metric_name"] = string(obj.MetricName)

	return result
}
