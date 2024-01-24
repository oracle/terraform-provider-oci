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

func DatabaseManagementManagedDatabaseOptimizerStatisticsCollectionOperationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementManagedDatabaseOptimizerStatisticsCollectionOperations,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"end_time_less_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"filter_by": {
				Type:     schema.TypeString,
				Optional: true,
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
			"optimizer_statistics_collection_operations_collection": {
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
									"completed_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"database": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"compartment_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"db_deployment_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"db_sub_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"db_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"db_version": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"duration_in_seconds": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"end_time": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"failed_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"in_progress_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"job_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"operation_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"start_time": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"tasks": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"status": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"target": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"target_type": {
													Type:     schema.TypeString,
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
											},
										},
									},
									"timed_out_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"total_objects_count": {
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

func readDatabaseManagementManagedDatabaseOptimizerStatisticsCollectionOperations(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseOptimizerStatisticsCollectionOperationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabaseOptimizerStatisticsCollectionOperationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListOptimizerStatisticsCollectionOperationsResponse
}

func (s *DatabaseManagementManagedDatabaseOptimizerStatisticsCollectionOperationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseOptimizerStatisticsCollectionOperationsDataSourceCrud) Get() error {
	request := oci_database_management.ListOptimizerStatisticsCollectionOperationsRequest{}

	if endTimeLessThanOrEqualTo, ok := s.D.GetOkExists("end_time_less_than_or_equal_to"); ok {
		tmp := endTimeLessThanOrEqualTo.(string)
		request.EndTimeLessThanOrEqualTo = &tmp
	}

	if filterBy, ok := s.D.GetOkExists("filter_by"); ok {
		tmp := filterBy.(string)
		request.FilterBy = &tmp
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
		request.TaskType = oci_database_management.ListOptimizerStatisticsCollectionOperationsTaskTypeEnum(taskType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListOptimizerStatisticsCollectionOperations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOptimizerStatisticsCollectionOperations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementManagedDatabaseOptimizerStatisticsCollectionOperationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabaseOptimizerStatisticsCollectionOperationsDataSource-", DatabaseManagementManagedDatabaseOptimizerStatisticsCollectionOperationsDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedDatabaseOptimizerStatisticsCollectionOperation := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OptimizerStatisticsCollectionOperationSummaryToMap(item))
	}
	managedDatabaseOptimizerStatisticsCollectionOperation["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementManagedDatabaseOptimizerStatisticsCollectionOperationsDataSource().Schema["optimizer_statistics_collection_operations_collection"].Elem.(*schema.Resource).Schema)
		managedDatabaseOptimizerStatisticsCollectionOperation["items"] = items
	}

	resources = append(resources, managedDatabaseOptimizerStatisticsCollectionOperation)
	if err := s.D.Set("optimizer_statistics_collection_operations_collection", resources); err != nil {
		return err
	}

	return nil
}

func OptimizerDatabaseToMap(obj *oci_database_management.OptimizerDatabase) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	result["db_deployment_type"] = string(obj.DbDeploymentType)

	result["db_sub_type"] = string(obj.DbSubType)

	result["db_type"] = string(obj.DbType)

	if obj.DbVersion != nil {
		result["db_version"] = string(*obj.DbVersion)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func OptimizerStatisticsCollectionOperationSummaryToMap(obj oci_database_management.OptimizerStatisticsCollectionOperationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompletedCount != nil {
		result["completed_count"] = int(*obj.CompletedCount)
	}

	if obj.DurationInSeconds != nil {
		result["duration_in_seconds"] = float32(*obj.DurationInSeconds)
	}

	if obj.EndTime != nil {
		result["end_time"] = string(*obj.EndTime)
	}

	if obj.FailedCount != nil {
		result["failed_count"] = int(*obj.FailedCount)
	}

	if obj.Id != nil {
		result["id"] = int(*obj.Id)
	}

	if obj.InProgressCount != nil {
		result["in_progress_count"] = int(*obj.InProgressCount)
	}

	if obj.JobName != nil {
		result["job_name"] = string(*obj.JobName)
	}

	if obj.OperationName != nil {
		result["operation_name"] = string(*obj.OperationName)
	}

	if obj.StartTime != nil {
		result["start_time"] = string(*obj.StartTime)
	}

	result["status"] = string(obj.Status)

	if obj.Target != nil {
		result["target"] = string(*obj.Target)
	}

	if obj.TimedOutCount != nil {
		result["timed_out_count"] = int(*obj.TimedOutCount)
	}

	if obj.TotalObjectsCount != nil {
		result["total_objects_count"] = int(*obj.TotalObjectsCount)
	}

	return result
}

func OptimizerStatisticsOperationTaskToMap(obj oci_database_management.OptimizerStatisticsOperationTask) map[string]interface{} {
	result := map[string]interface{}{}

	result["status"] = string(obj.Status)

	if obj.Target != nil {
		result["target"] = string(*obj.Target)
	}

	result["target_type"] = string(obj.TargetType)

	if obj.TimeEnd != nil {
		result["time_end"] = obj.TimeEnd.String()
	}

	if obj.TimeStart != nil {
		result["time_start"] = obj.TimeStart.String()
	}

	return result
}
