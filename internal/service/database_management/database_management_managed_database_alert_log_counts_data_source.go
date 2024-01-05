// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementManagedDatabaseAlertLogCountsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementManagedDatabaseAlertLogCounts,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"group_by": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_regular_expression": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"level_filter": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_search_text": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"managed_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"time_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_less_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type_filter": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"alert_log_counts_collection": {
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
									"category": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"count": {
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

func readDatabaseManagementManagedDatabaseAlertLogCounts(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseAlertLogCountsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DiagnosabilityClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabaseAlertLogCountsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DiagnosabilityClient
	Res    *oci_database_management.SummarizeAlertLogCountsResponse
}

func (s *DatabaseManagementManagedDatabaseAlertLogCountsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseAlertLogCountsDataSourceCrud) Get() error {
	request := oci_database_management.SummarizeAlertLogCountsRequest{}

	if groupBy, ok := s.D.GetOkExists("group_by"); ok {
		request.GroupBy = oci_database_management.SummarizeAlertLogCountsGroupByEnum(groupBy.(string))
	}

	if isRegularExpression, ok := s.D.GetOkExists("is_regular_expression"); ok {
		tmp := isRegularExpression.(bool)
		request.IsRegularExpression = &tmp
	}

	if levelFilter, ok := s.D.GetOkExists("level_filter"); ok {
		request.LevelFilter = oci_database_management.SummarizeAlertLogCountsLevelFilterEnum(levelFilter.(string))
	}

	if logSearchText, ok := s.D.GetOkExists("log_search_text"); ok {
		tmp := logSearchText.(string)
		request.LogSearchText = &tmp
	}

	if managedDatabaseId, ok := s.D.GetOkExists("managed_database_id"); ok {
		tmp := managedDatabaseId.(string)
		request.ManagedDatabaseId = &tmp
	}

	if timeGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeLessThanOrEqualTo, ok := s.D.GetOkExists("time_less_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeLessThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeLessThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if typeFilter, ok := s.D.GetOkExists("type_filter"); ok {
		request.TypeFilter = oci_database_management.SummarizeAlertLogCountsTypeFilterEnum(typeFilter.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.SummarizeAlertLogCounts(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.SummarizeAlertLogCounts(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementManagedDatabaseAlertLogCountsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabaseAlertLogCountsDataSource-", DatabaseManagementManagedDatabaseAlertLogCountsDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedDatabaseAlertLogCount := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AlertLogCountSummaryToMap(item))
	}
	managedDatabaseAlertLogCount["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementManagedDatabaseAlertLogCountsDataSource().Schema["alert_log_counts_collection"].Elem.(*schema.Resource).Schema)
		managedDatabaseAlertLogCount["items"] = items
	}

	resources = append(resources, managedDatabaseAlertLogCount)
	if err := s.D.Set("alert_log_counts_collection", resources); err != nil {
		return err
	}

	return nil
}

func AlertLogCountSummaryToMap(obj oci_database_management.AlertLogCountSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["category"] = string(obj.Category)

	if obj.Count != nil {
		result["count"] = int(*obj.Count)
	}

	return result
}
