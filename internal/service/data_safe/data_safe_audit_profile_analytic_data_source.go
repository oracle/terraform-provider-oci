// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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

func DataSafeAuditProfileAnalyticDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDataSafeAuditProfileAnalytic,
		Schema: map[string]*schema.Schema{
			"access_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"group_by": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
									"is_paid_usage_enabled": {
										Type:     schema.TypeBool,
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

func readSingularDataSafeAuditProfileAnalytic(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAuditProfileAnalyticDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeAuditProfileAnalyticDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListAuditProfileAnalyticsResponse
}

func (s *DataSafeAuditProfileAnalyticDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeAuditProfileAnalyticDataSourceCrud) Get() error {
	request := oci_data_safe.ListAuditProfileAnalyticsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListAuditProfileAnalyticsAccessLevelEnum(accessLevel.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if groupBy, ok := s.D.GetOkExists("group_by"); ok {
		interfaces := groupBy.([]interface{})
		tmp := make([]oci_data_safe.ListAuditProfileAnalyticsGroupByEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(oci_data_safe.ListAuditProfileAnalyticsGroupByEnum)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("group_by") {
			request.GroupBy = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListAuditProfileAnalytics(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeAuditProfileAnalyticDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeAuditProfileAnalyticDataSource-", DataSafeAuditProfileAnalyticDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AuditProfileAggregationItemsToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func AuditProfileAggregationItemsToMap(obj oci_data_safe.AuditProfileAggregationItems) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Count != nil {
		result["count"] = strconv.FormatInt(*obj.Count, 10)
	}

	if obj.Dimensions != nil {
		result["dimensions"] = []interface{}{AuditProfileDimensionsToMap(obj.Dimensions)}
	}

	return result
}

func AuditProfileDimensionsToMap(obj *oci_data_safe.AuditProfileDimensions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsPaidUsageEnabled != nil {
		result["is_paid_usage_enabled"] = bool(*obj.IsPaidUsageEnabled)
	}

	return result
}
