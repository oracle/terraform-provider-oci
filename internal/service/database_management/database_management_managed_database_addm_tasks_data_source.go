// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementManagedDatabaseAddmTasksDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementManagedDatabaseAddmTasks,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"managed_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"time_end": {
				Type:     schema.TypeString,
				Required: true,
			},
			"time_start": {
				Type:     schema.TypeString,
				Required: true,
			},
			"addm_tasks_collection": {
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
									"begin_snapshot_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"db_user": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"end_snapshot_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"end_snapshot_time": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"findings": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"how_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"start_snapshot_time": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"task_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"task_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
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

func readDatabaseManagementManagedDatabaseAddmTasks(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseAddmTasksDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabaseAddmTasksDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.AddmTasksResponse
}

func (s *DatabaseManagementManagedDatabaseAddmTasksDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseAddmTasksDataSourceCrud) Get() error {
	request := oci_database_management.AddmTasksRequest{}

	if managedDatabaseId, ok := s.D.GetOkExists("managed_database_id"); ok {
		tmp := managedDatabaseId.(string)
		request.ManagedDatabaseId = &tmp
	}

	if timeEnd, ok := s.D.GetOkExists("time_end"); ok {
		tmp, err := time.Parse(time.RFC3339, timeEnd.(string))
		if err != nil {
			return err
		}
		request.TimeEnd = &oci_common.SDKTime{Time: tmp}
	}

	if timeStart, ok := s.D.GetOkExists("time_start"); ok {
		tmp, err := time.Parse(time.RFC3339, timeStart.(string))
		if err != nil {
			return err
		}
		request.TimeStart = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.AddmTasks(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.AddmTasks(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementManagedDatabaseAddmTasksDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabaseAddmTasksDataSource-", DatabaseManagementManagedDatabaseAddmTasksDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedDatabaseAddmTask := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AddmTaskSummaryToMap(item))
	}
	managedDatabaseAddmTask["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementManagedDatabaseAddmTasksDataSource().Schema["addm_tasks_collection"].Elem.(*schema.Resource).Schema)
		managedDatabaseAddmTask["items"] = items
	}

	resources = append(resources, managedDatabaseAddmTask)
	if err := s.D.Set("addm_tasks_collection", resources); err != nil {
		return err
	}

	return nil
}

func AddmTaskSummaryToMap(obj oci_database_management.AddmTaskSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BeginSnapshotId != nil {
		result["begin_snapshot_id"] = strconv.FormatInt(*obj.BeginSnapshotId, 10)
	}

	if obj.DbUser != nil {
		result["db_user"] = string(*obj.DbUser)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.EndSnapshotId != nil {
		result["end_snapshot_id"] = strconv.FormatInt(*obj.EndSnapshotId, 10)
	}

	if obj.EndSnapshotTime != nil {
		result["end_snapshot_time"] = obj.EndSnapshotTime.String()
	}

	if obj.Findings != nil {
		result["findings"] = strconv.FormatInt(*obj.Findings, 10)
	}

	result["how_created"] = string(obj.HowCreated)

	if obj.StartSnapshotTime != nil {
		result["start_snapshot_time"] = obj.StartSnapshotTime.String()
	}

	result["status"] = string(obj.Status)

	if obj.TaskId != nil {
		result["task_id"] = strconv.FormatInt(*obj.TaskId, 10)
	}

	if obj.TaskName != nil {
		result["task_name"] = string(*obj.TaskName)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}
