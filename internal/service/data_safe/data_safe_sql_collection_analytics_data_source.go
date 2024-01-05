// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSqlCollectionAnalyticsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeSqlCollectionAnalytics,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
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
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_ended": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_started": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sql_collection_analytics_collection": {
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
									"dimensions": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"state": {
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
									"sql_collection_analytic_count": {
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

func readDataSafeSqlCollectionAnalytics(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSqlCollectionAnalyticsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSqlCollectionAnalyticsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListSqlCollectionAnalyticsResponse
}

func (s *DataSafeSqlCollectionAnalyticsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSqlCollectionAnalyticsDataSourceCrud) Get() error {
	request := oci_data_safe.ListSqlCollectionAnalyticsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListSqlCollectionAnalyticsAccessLevelEnum(accessLevel.(string))
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
		tmp := make([]oci_data_safe.ListSqlCollectionAnalyticsGroupByEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_data_safe.ListSqlCollectionAnalyticsGroupByEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("group_by") {
			request.GroupBy = tmp
		}
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_data_safe.ListSqlCollectionAnalyticsLifecycleStateEnum(state.(string))
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	if timeEnded, ok := s.D.GetOkExists("time_ended"); ok {
		tmp, err := time.Parse(time.RFC3339, timeEnded.(string))
		if err != nil {
			return err
		}
		request.TimeEnded = &oci_common.SDKTime{Time: tmp}
	}

	if timeStarted, ok := s.D.GetOkExists("time_started"); ok {
		tmp, err := time.Parse(time.RFC3339, timeStarted.(string))
		if err != nil {
			return err
		}
		request.TimeStarted = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListSqlCollectionAnalytics(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSqlCollectionAnalytics(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSqlCollectionAnalyticsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSqlCollectionAnalyticsDataSource-", DataSafeSqlCollectionAnalyticsDataSource(), s.D))
	resources := []map[string]interface{}{}
	sqlCollectionAnalytic := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SqlCollectionAggregationToMap(item))
	}
	sqlCollectionAnalytic["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSqlCollectionAnalyticsDataSource().Schema["sql_collection_analytics_collection"].Elem.(*schema.Resource).Schema)
		sqlCollectionAnalytic["items"] = items
	}

	resources = append(resources, sqlCollectionAnalytic)
	if err := s.D.Set("sql_collection_analytics_collection", resources); err != nil {
		return err
	}

	return nil
}

func SqlCollectionAggregationToMap(obj oci_data_safe.SqlCollectionAggregation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Dimensions != nil {
		result["dimensions"] = []interface{}{SqlCollectionDimensionsToMap(obj.Dimensions)}
	}

	if obj.Count != nil {
		result["sql_collection_analytic_count"] = strconv.FormatInt(*obj.Count, 10)
	}

	return result
}

func SqlCollectionDimensionsToMap(obj *oci_data_safe.SqlCollectionDimensions) map[string]interface{} {
	result := map[string]interface{}{}

	result["state"] = string(obj.LifecycleState)

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	return result
}
