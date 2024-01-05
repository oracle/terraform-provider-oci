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

func DataSafeSqlFirewallAllowedSqlAnalyticsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeSqlFirewallAllowedSqlAnalytics,
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
			"scim_query": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sql_firewall_allowed_sql_analytics_collection": {
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
												"db_user_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"sql_firewall_policy_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"sql_level": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"state": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"sql_firewall_allowed_sql_analytic_count": {
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

func readDataSafeSqlFirewallAllowedSqlAnalytics(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSqlFirewallAllowedSqlAnalyticsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSqlFirewallAllowedSqlAnalyticsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListSqlFirewallAllowedSqlAnalyticsResponse
}

func (s *DataSafeSqlFirewallAllowedSqlAnalyticsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSqlFirewallAllowedSqlAnalyticsDataSourceCrud) Get() error {
	request := oci_data_safe.ListSqlFirewallAllowedSqlAnalyticsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListSqlFirewallAllowedSqlAnalyticsAccessLevelEnum(accessLevel.(string))
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
		tmp := make([]oci_data_safe.ListSqlFirewallAllowedSqlAnalyticsGroupByEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(oci_data_safe.ListSqlFirewallAllowedSqlAnalyticsGroupByEnum)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("group_by") {
			request.GroupBy = tmp
		}
	}

	if scimQuery, ok := s.D.GetOkExists("scim_query"); ok {
		tmp := scimQuery.(string)
		request.ScimQuery = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListSqlFirewallAllowedSqlAnalytics(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSqlFirewallAllowedSqlAnalytics(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSqlFirewallAllowedSqlAnalyticsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSqlFirewallAllowedSqlAnalyticsDataSource-", DataSafeSqlFirewallAllowedSqlAnalyticsDataSource(), s.D))
	resources := []map[string]interface{}{}
	sqlFirewallAllowedSqlAnalytic := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SqlFirewallAllowedSqlAggregationToMap(item))
	}
	sqlFirewallAllowedSqlAnalytic["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSqlFirewallAllowedSqlAnalyticsDataSource().Schema["sql_firewall_allowed_sql_analytics_collection"].Elem.(*schema.Resource).Schema)
		sqlFirewallAllowedSqlAnalytic["items"] = items
	}

	resources = append(resources, sqlFirewallAllowedSqlAnalytic)
	if err := s.D.Set("sql_firewall_allowed_sql_analytics_collection", resources); err != nil {
		return err
	}

	return nil
}

func SqlFirewallAllowedSqlAggregationToMap(obj oci_data_safe.SqlFirewallAllowedSqlAggregation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Dimensions != nil {
		result["dimensions"] = []interface{}{SqlFirewallAllowedSqlDimensionsToMap(obj.Dimensions)}
	}

	if obj.Count != nil {
		result["sql_firewall_allowed_sql_analytic_count"] = strconv.FormatInt(*obj.Count, 10)
	}

	return result
}

func SqlFirewallAllowedSqlDimensionsToMap(obj *oci_data_safe.SqlFirewallAllowedSqlDimensions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DbUserName != nil {
		result["db_user_name"] = string(*obj.DbUserName)
	}

	if obj.SqlFirewallPolicyId != nil {
		result["sql_firewall_policy_id"] = string(*obj.SqlFirewallPolicyId)
	}

	result["sql_level"] = string(obj.SqlLevel)

	result["state"] = string(obj.LifecycleState)

	return result
}
