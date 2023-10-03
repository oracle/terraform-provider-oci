// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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

func DataSafeSqlFirewallViolationAnalyticsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeSqlFirewallViolationAnalytics,
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
			"query_time_zone": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scim_query": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"summary_field": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"time_ended": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_started": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sql_firewall_violation_analytics_collection": {
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
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"client_os_user_name": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"client_program": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"db_user_name": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"operation": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"operation_time": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"sql_level": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"target_id": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"target_name": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"violation_action": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"violation_cause": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"metric_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"sql_firewall_violation_analytic_count": {
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

func readDataSafeSqlFirewallViolationAnalytics(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSqlFirewallViolationAnalyticsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSqlFirewallViolationAnalyticsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListSqlFirewallViolationAnalyticsResponse
}

func (s *DataSafeSqlFirewallViolationAnalyticsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSqlFirewallViolationAnalyticsDataSourceCrud) Get() error {
	request := oci_data_safe.ListSqlFirewallViolationAnalyticsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListSqlFirewallViolationAnalyticsAccessLevelEnum(accessLevel.(string))
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
		tmp := make([]oci_data_safe.ListSqlFirewallViolationAnalyticsGroupByEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_data_safe.ListSqlFirewallViolationAnalyticsGroupByEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("group_by") {
			request.GroupBy = tmp
		}
	}

	if queryTimeZone, ok := s.D.GetOkExists("query_time_zone"); ok {
		tmp := queryTimeZone.(string)
		request.QueryTimeZone = &tmp
	}

	if scimQuery, ok := s.D.GetOkExists("scim_query"); ok {
		tmp := scimQuery.(string)
		request.ScimQuery = &tmp
	}

	if summaryField, ok := s.D.GetOkExists("summary_field"); ok {
		interfaces := summaryField.([]interface{})
		tmp := make([]oci_data_safe.ListSqlFirewallViolationAnalyticsSummaryFieldEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_data_safe.ListSqlFirewallViolationAnalyticsSummaryFieldEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("summary_field") {
			request.SummaryField = tmp
		}
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

	response, err := s.Client.ListSqlFirewallViolationAnalytics(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSqlFirewallViolationAnalytics(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSqlFirewallViolationAnalyticsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSqlFirewallViolationAnalyticsDataSource-", DataSafeSqlFirewallViolationAnalyticsDataSource(), s.D))
	resources := []map[string]interface{}{}
	sqlFirewallViolationAnalytic := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SqlFirewallViolationAggregationToMap(item))
	}
	sqlFirewallViolationAnalytic["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSqlFirewallViolationAnalyticsDataSource().Schema["sql_firewall_violation_analytics_collection"].Elem.(*schema.Resource).Schema)
		sqlFirewallViolationAnalytic["items"] = items
	}

	resources = append(resources, sqlFirewallViolationAnalytic)
	if err := s.D.Set("sql_firewall_violation_analytics_collection", resources); err != nil {
		return err
	}

	return nil
}

func SqlFirewallViolationAggregationToMap(obj oci_data_safe.SqlFirewallViolationAggregation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Dimensions != nil {
		result["dimensions"] = []interface{}{SqlFirewallViolationAggregationDimensionsToMap(obj.Dimensions)}
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.MetricName != nil {
		result["metric_name"] = string(*obj.MetricName)
	}

	if obj.Count != nil {
		result["sql_firewall_violation_analytic_count"] = strconv.FormatInt(*obj.Count, 10)
	}

	if obj.TimeEnded != nil {
		result["time_ended"] = obj.TimeEnded.String()
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	return result
}

func SqlFirewallViolationAggregationDimensionsToMap(obj *oci_data_safe.SqlFirewallViolationAggregationDimensions) map[string]interface{} {
	result := map[string]interface{}{}

	result["client_ip"] = obj.ClientIp
	result["client_ip"] = obj.ClientIp

	result["client_os_user_name"] = obj.ClientOsUserName
	result["client_os_user_name"] = obj.ClientOsUserName

	result["client_program"] = obj.ClientProgram
	result["client_program"] = obj.ClientProgram

	result["db_user_name"] = obj.DbUserName
	result["db_user_name"] = obj.DbUserName

	result["operation"] = obj.Operation
	result["operation"] = obj.Operation

	result["operation_time"] = obj.OperationTime
	result["operation_time"] = obj.OperationTime

	result["sql_level"] = obj.SqlLevel
	result["sql_level"] = obj.SqlLevel

	result["target_id"] = obj.TargetId
	result["target_id"] = obj.TargetId

	result["target_name"] = obj.TargetName
	result["target_name"] = obj.TargetName

	result["violation_action"] = obj.ViolationAction
	result["violation_action"] = obj.ViolationAction

	result["violation_cause"] = obj.ViolationCause
	result["violation_cause"] = obj.ViolationCause

	return result
}
