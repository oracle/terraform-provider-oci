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

func DatabaseManagementManagedDatabaseAttentionLogCountsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementManagedDatabaseAttentionLogCounts,
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
			"urgency_filter": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"attention_log_counts_collection": {
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

func readDatabaseManagementManagedDatabaseAttentionLogCounts(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseAttentionLogCountsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DiagnosabilityClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabaseAttentionLogCountsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DiagnosabilityClient
	Res    *oci_database_management.SummarizeAttentionLogCountsResponse
}

func (s *DatabaseManagementManagedDatabaseAttentionLogCountsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseAttentionLogCountsDataSourceCrud) Get() error {
	request := oci_database_management.SummarizeAttentionLogCountsRequest{}

	if groupBy, ok := s.D.GetOkExists("group_by"); ok {
		request.GroupBy = oci_database_management.SummarizeAttentionLogCountsGroupByEnum(groupBy.(string))
	}

	if isRegularExpression, ok := s.D.GetOkExists("is_regular_expression"); ok {
		tmp := isRegularExpression.(bool)
		request.IsRegularExpression = &tmp
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
		request.TypeFilter = oci_database_management.SummarizeAttentionLogCountsTypeFilterEnum(typeFilter.(string))
	}

	if urgencyFilter, ok := s.D.GetOkExists("urgency_filter"); ok {
		request.UrgencyFilter = oci_database_management.SummarizeAttentionLogCountsUrgencyFilterEnum(urgencyFilter.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.SummarizeAttentionLogCounts(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.SummarizeAttentionLogCounts(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementManagedDatabaseAttentionLogCountsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabaseAttentionLogCountsDataSource-", DatabaseManagementManagedDatabaseAttentionLogCountsDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedDatabaseAttentionLogCount := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AttentionLogCountSummaryToMap(item))
	}
	managedDatabaseAttentionLogCount["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementManagedDatabaseAttentionLogCountsDataSource().Schema["attention_log_counts_collection"].Elem.(*schema.Resource).Schema)
		managedDatabaseAttentionLogCount["items"] = items
	}

	resources = append(resources, managedDatabaseAttentionLogCount)
	if err := s.D.Set("attention_log_counts_collection", resources); err != nil {
		return err
	}

	return nil
}

func AttentionLogCountSummaryToMap(obj oci_database_management.AttentionLogCountSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["category"] = string(obj.Category)

	if obj.Count != nil {
		result["count"] = int(*obj.Count)
	}

	return result
}
