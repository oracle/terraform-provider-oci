// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"strconv"

	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeAuditTrailAnalyticDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDataSafeAuditTrailAnalytic,
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
			"target_id": {
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
									"location": {
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
									"target_id": {
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

func readSingularDataSafeAuditTrailAnalytic(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAuditTrailAnalyticDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeAuditTrailAnalyticDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListAuditTrailAnalyticsResponse
}

func (s *DataSafeAuditTrailAnalyticDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeAuditTrailAnalyticDataSourceCrud) Get() error {
	request := oci_data_safe.ListAuditTrailAnalyticsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListAuditTrailAnalyticsAccessLevelEnum(accessLevel.(string))
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
		tmp := make([]oci_data_safe.ListAuditTrailAnalyticsGroupByEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_data_safe.ListAuditTrailAnalyticsGroupByEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("group_by") {
			request.GroupBy = tmp
		}
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListAuditTrailAnalytics(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeAuditTrailAnalyticDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeAuditTrailAnalyticDataSource-", DataSafeAuditTrailAnalyticDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AuditTrailAggregationItemsToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func AuditTrailAggregationItemsToMap(obj oci_data_safe.AuditTrailAggregationItems) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Count != nil {
		result["count"] = strconv.FormatInt(*obj.Count, 10)
	}

	if obj.Dimensions != nil {
		result["dimensions"] = []interface{}{AuditTrailDimensionsToMap(obj.Dimensions)}
	}

	return result
}

func AuditTrailDimensionsToMap(obj *oci_data_safe.AuditTrailDimensions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Location != nil {
		result["location"] = string(*obj.Location)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.Status != nil {
		result["status"] = string(*obj.Status)
	}

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	return result
}
