// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"
	"strconv"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"
)

func DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementManagedDatabaseSqlTuningAdvisorTasks,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"managed_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"opc_named_credential_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_less_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sql_tuning_advisor_task_collection": {
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
									"days_to_expire": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"instance_id": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"owner": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"recommendation_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"sql_tuning_advisor_task_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"task_status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_execution_ended": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_execution_started": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"total_sql_statements": {
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

func readDatabaseManagementManagedDatabaseSqlTuningAdvisorTasks(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SqlTuningClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.SqlTuningClient
	Res    *oci_database_management.ListSqlTuningAdvisorTasksResponse
}

func (s *DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksDataSourceCrud) Get() error {
	request := oci_database_management.ListSqlTuningAdvisorTasksRequest{}

	if managedDatabaseId, ok := s.D.GetOkExists("managed_database_id"); ok {
		tmp := managedDatabaseId.(string)
		request.ManagedDatabaseId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if opcNamedCredentialId, ok := s.D.GetOkExists("opc_named_credential_id"); ok {
		tmp := opcNamedCredentialId.(string)
		request.OpcNamedCredentialId = &tmp
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_database_management.ListSqlTuningAdvisorTasksStatusEnum(status.(string))
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListSqlTuningAdvisorTasks(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSqlTuningAdvisorTasks(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksDataSource-", DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedDatabaseSqlTuningAdvisorTask := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SqlTuningAdvisorTaskSummaryToMap(item))
	}
	managedDatabaseSqlTuningAdvisorTask["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksDataSource().Schema["sql_tuning_advisor_task_collection"].Elem.(*schema.Resource).Schema)
		managedDatabaseSqlTuningAdvisorTask["items"] = items
	}

	resources = append(resources, managedDatabaseSqlTuningAdvisorTask)
	if err := s.D.Set("sql_tuning_advisor_task_collection", resources); err != nil {
		return err
	}

	return nil
}

func SqlTuningAdvisorTaskSummaryToMap(obj oci_database_management.SqlTuningAdvisorTaskSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DaysToExpire != nil {
		result["days_to_expire"] = int(*obj.DaysToExpire)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.InstanceId != nil {
		result["instance_id"] = int(*obj.InstanceId)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Owner != nil {
		result["owner"] = string(*obj.Owner)
	}

	if obj.RecommendationCount != nil {
		result["recommendation_count"] = int(*obj.RecommendationCount)
	}

	if obj.SqlTuningAdvisorTaskId != nil {
		result["sql_tuning_advisor_task_id"] = strconv.FormatInt(*obj.SqlTuningAdvisorTaskId, 10)
	}

	result["task_status"] = string(obj.TaskStatus)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeExecutionEnded != nil {
		result["time_execution_ended"] = obj.TimeExecutionEnded.String()
	}

	if obj.TimeExecutionStarted != nil {
		result["time_execution_started"] = obj.TimeExecutionStarted.String()
	}

	if obj.TotalSqlStatements != nil {
		result["total_sql_statements"] = int(*obj.TotalSqlStatements)
	}

	return result
}
