// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementManagedDatabaseOptimizerStatisticsCollectionAggregationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementManagedDatabaseOptimizerStatisticsCollectionAggregations,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"end_time_less_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"managed_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"start_time_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"task_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"optimizer_statistics_collection_aggregations_collection": {
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
									"completed": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"failed": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"group_by": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"in_progress": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"pending": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"skipped": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"time_end": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_start": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"timed_out": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"total": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"unknown": {
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

func readDatabaseManagementManagedDatabaseOptimizerStatisticsCollectionAggregations(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseOptimizerStatisticsCollectionAggregationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabaseOptimizerStatisticsCollectionAggregationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListOptimizerStatisticsCollectionAggregationsResponse
}

func (s *DatabaseManagementManagedDatabaseOptimizerStatisticsCollectionAggregationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseOptimizerStatisticsCollectionAggregationsDataSourceCrud) Get() error {
	request := oci_database_management.ListOptimizerStatisticsCollectionAggregationsRequest{}

	if endTimeLessThanOrEqualTo, ok := s.D.GetOkExists("end_time_less_than_or_equal_to"); ok {
		tmp := endTimeLessThanOrEqualTo.(string)
		request.EndTimeLessThanOrEqualTo = &tmp
	}

	if groupType, ok := s.D.GetOkExists("group_type"); ok {
		request.GroupType = oci_database_management.ListOptimizerStatisticsCollectionAggregationsGroupTypeEnum(groupType.(string))
	}

	if managedDatabaseId, ok := s.D.GetOkExists("managed_database_id"); ok {
		tmp := managedDatabaseId.(string)
		request.ManagedDatabaseId = &tmp
	}

	if limit, ok := s.D.GetOkExists("limit"); ok {
		tmp := limit.(int)
		request.Limit = &tmp
	}

	if startTimeGreaterThanOrEqualTo, ok := s.D.GetOkExists("start_time_greater_than_or_equal_to"); ok {
		tmp := startTimeGreaterThanOrEqualTo.(string)
		request.StartTimeGreaterThanOrEqualTo = &tmp
	}

	if taskType, ok := s.D.GetOkExists("task_type"); ok {
		request.TaskType = oci_database_management.ListOptimizerStatisticsCollectionAggregationsTaskTypeEnum(taskType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListOptimizerStatisticsCollectionAggregations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOptimizerStatisticsCollectionAggregations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementManagedDatabaseOptimizerStatisticsCollectionAggregationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabaseOptimizerStatisticsCollectionAggregationsDataSource-", DatabaseManagementManagedDatabaseOptimizerStatisticsCollectionAggregationsDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedDatabaseOptimizerStatisticsCollectionAggregation := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OptimizerStatisticsCollectionAggregationSummaryToMap(item))
	}
	managedDatabaseOptimizerStatisticsCollectionAggregation["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementManagedDatabaseOptimizerStatisticsCollectionAggregationsDataSource().Schema["optimizer_statistics_collection_aggregations_collection"].Elem.(*schema.Resource).Schema)
		managedDatabaseOptimizerStatisticsCollectionAggregation["items"] = items
	}

	resources = append(resources, managedDatabaseOptimizerStatisticsCollectionAggregation)
	if err := s.D.Set("optimizer_statistics_collection_aggregations_collection", resources); err != nil {
		return err
	}

	return nil
}

func OptimizerStatisticsCollectionAggregationSummaryToMap(obj oci_database_management.OptimizerStatisticsCollectionAggregationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Completed != nil {
		result["completed"] = int(*obj.Completed)
	}

	if obj.Failed != nil {
		result["failed"] = int(*obj.Failed)
	}

	result["group_by"] = string(obj.GroupBy)

	if obj.InProgress != nil {
		result["in_progress"] = int(*obj.InProgress)
	}

	if obj.Pending != nil {
		result["pending"] = int(*obj.Pending)
	}

	if obj.Skipped != nil {
		result["skipped"] = int(*obj.Skipped)
	}

	if obj.TimeEnd != nil {
		result["time_end"] = obj.TimeEnd.String()
	}

	if obj.TimeStart != nil {
		result["time_start"] = obj.TimeStart.String()
	}

	if obj.TimedOut != nil {
		result["timed_out"] = int(*obj.TimedOut)
	}

	if obj.Total != nil {
		result["total"] = int(*obj.Total)
	}

	if obj.Unknown != nil {
		result["unknown"] = int(*obj.Unknown)
	}

	return result
}
