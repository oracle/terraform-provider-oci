// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementManagedDatabaseOptimizerStatisticsCollectionOperationDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseManagementManagedDatabaseOptimizerStatisticsCollectionOperation,
		Schema: map[string]*schema.Schema{
			"managed_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"optimizer_statistics_collection_operation_id": {
				Type:     schema.TypeFloat,
				Required: true,
			},
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
	}
}

func readSingularDatabaseManagementManagedDatabaseOptimizerStatisticsCollectionOperation(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseOptimizerStatisticsCollectionOperationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabaseOptimizerStatisticsCollectionOperationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.GetOptimizerStatisticsCollectionOperationResponse
}

func (s *DatabaseManagementManagedDatabaseOptimizerStatisticsCollectionOperationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseOptimizerStatisticsCollectionOperationDataSourceCrud) Get() error {
	request := oci_database_management.GetOptimizerStatisticsCollectionOperationRequest{}

	if managedDatabaseId, ok := s.D.GetOkExists("managed_database_id"); ok {
		tmp := managedDatabaseId.(string)
		request.ManagedDatabaseId = &tmp
	}

	if optimizerStatisticsCollectionOperationId, ok := s.D.GetOkExists("optimizer_statistics_collection_operation_id"); ok {
		tmp := float32(optimizerStatisticsCollectionOperationId.(float64))
		request.OptimizerStatisticsCollectionOperationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetOptimizerStatisticsCollectionOperation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementManagedDatabaseOptimizerStatisticsCollectionOperationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(strconv.Itoa(*s.Res.Id))

	if s.Res.CompletedCount != nil {
		s.D.Set("completed_count", *s.Res.CompletedCount)
	}

	if s.Res.Database != nil {
		s.D.Set("database", []interface{}{OperationOptimizerDatabaseToMap(s.Res.Database)})
	} else {
		s.D.Set("database", nil)
	}

	if s.Res.DurationInSeconds != nil {
		s.D.Set("duration_in_seconds", *s.Res.DurationInSeconds)
	}

	if s.Res.EndTime != nil {
		s.D.Set("end_time", *s.Res.EndTime)
	}

	if s.Res.FailedCount != nil {
		s.D.Set("failed_count", *s.Res.FailedCount)
	}

	if s.Res.InProgressCount != nil {
		s.D.Set("in_progress_count", *s.Res.InProgressCount)
	}

	if s.Res.JobName != nil {
		s.D.Set("job_name", *s.Res.JobName)
	}

	if s.Res.OperationName != nil {
		s.D.Set("operation_name", *s.Res.OperationName)
	}

	if s.Res.StartTime != nil {
		s.D.Set("start_time", *s.Res.StartTime)
	}

	s.D.Set("status", s.Res.Status)

	if s.Res.Target != nil {
		s.D.Set("target", *s.Res.Target)
	}

	tasks := []interface{}{}
	for _, item := range s.Res.Tasks {
		tasks = append(tasks, OperationOptimizerStatisticsOperationTaskToMap(item))
	}
	s.D.Set("tasks", tasks)

	if s.Res.TimedOutCount != nil {
		s.D.Set("timed_out_count", *s.Res.TimedOutCount)
	}

	if s.Res.TotalObjectsCount != nil {
		s.D.Set("total_objects_count", *s.Res.TotalObjectsCount)
	}

	return nil
}

func OperationOptimizerDatabaseToMap(obj *oci_database_management.OptimizerDatabase) map[string]interface{} {
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

func OperationOptimizerStatisticsCollectionOperationSummaryToMap(obj oci_database_management.OptimizerStatisticsCollectionOperationSummary) map[string]interface{} {
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

func OperationOptimizerStatisticsOperationTaskToMap(obj oci_database_management.OptimizerStatisticsOperationTask) map[string]interface{} {
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
