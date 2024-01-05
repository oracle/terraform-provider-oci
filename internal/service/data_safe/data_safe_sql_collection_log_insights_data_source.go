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

func DataSafeSqlCollectionLogInsightsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeSqlCollectionLogInsights,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"group_by": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sql_collection_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"time_ended": {
				Type:     schema.TypeString,
				Required: true,
			},
			"time_started": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sql_collection_log_insights_collection": {
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
												"client_ip": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"client_os_user_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"client_program": {
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
									"sql_collection_log_insight_count": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_ended": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_started": {
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

func readDataSafeSqlCollectionLogInsights(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSqlCollectionLogInsightsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSqlCollectionLogInsightsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListSqlCollectionLogInsightsResponse
}

func (s *DataSafeSqlCollectionLogInsightsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSqlCollectionLogInsightsDataSourceCrud) Get() error {
	request := oci_data_safe.ListSqlCollectionLogInsightsRequest{}

	if groupBy, ok := s.D.GetOkExists("group_by"); ok {
		request.GroupBy = oci_data_safe.ListSqlCollectionLogInsightsGroupByEnum(groupBy.(string))
	}

	if sqlCollectionId, ok := s.D.GetOkExists("sql_collection_id"); ok {
		tmp := sqlCollectionId.(string)
		request.SqlCollectionId = &tmp
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

	response, err := s.Client.ListSqlCollectionLogInsights(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSqlCollectionLogInsights(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSqlCollectionLogInsightsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSqlCollectionLogInsightsDataSource-", DataSafeSqlCollectionLogInsightsDataSource(), s.D))
	resources := []map[string]interface{}{}
	sqlCollectionLogInsight := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SqlCollectionLogAggregationToMap(item))
	}
	sqlCollectionLogInsight["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSqlCollectionLogInsightsDataSource().Schema["sql_collection_log_insights_collection"].Elem.(*schema.Resource).Schema)
		sqlCollectionLogInsight["items"] = items
	}

	resources = append(resources, sqlCollectionLogInsight)
	if err := s.D.Set("sql_collection_log_insights_collection", resources); err != nil {
		return err
	}

	return nil
}

func SqlCollectionLogAggregationToMap(obj oci_data_safe.SqlCollectionLogAggregation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Dimensions != nil {
		result["dimensions"] = []interface{}{SqlCollectionLogDimensionsToMap(obj.Dimensions)}
	}

	if obj.MetricName != nil {
		result["metric_name"] = string(*obj.MetricName)
	}

	if obj.Count != nil {
		result["sql_collection_log_insight_count"] = strconv.FormatInt(*obj.Count, 10)
	}

	if obj.TimeEnded != nil {
		result["time_ended"] = obj.TimeEnded.String()
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	return result
}

func SqlCollectionLogDimensionsToMap(obj *oci_data_safe.SqlCollectionLogDimensions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ClientIp != nil {
		result["client_ip"] = string(*obj.ClientIp)
	}

	if obj.ClientOsUserName != nil {
		result["client_os_user_name"] = string(*obj.ClientOsUserName)
	}

	if obj.ClientProgram != nil {
		result["client_program"] = string(*obj.ClientProgram)
	}

	return result
}
