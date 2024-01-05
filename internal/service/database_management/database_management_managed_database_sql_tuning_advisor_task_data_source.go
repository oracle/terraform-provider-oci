// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"
)

func DatabaseManagementManagedDatabaseSqlTuningAdvisorTaskDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseManagementManagedDatabaseSqlTuningAdvisorTask,
		Schema: map[string]*schema.Schema{
			"managed_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
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
			// Computed
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
		DeprecationMessage: tfresource.DatasourceDeprecatedForAnother("oci_database_management_managed_database_sql_tuning_advisor_task", "oci_database_management_managed_database_sql_tuning_advisor_tasks"),
	}
}

func readSingularDatabaseManagementManagedDatabaseSqlTuningAdvisorTask(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseSqlTuningAdvisorTaskDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SqlTuningClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabaseSqlTuningAdvisorTaskDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.SqlTuningClient
	Res    *oci_database_management.ListSqlTuningAdvisorTasksResponse
}

func (s *DatabaseManagementManagedDatabaseSqlTuningAdvisorTaskDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseSqlTuningAdvisorTaskDataSourceCrud) Get() error {
	request := oci_database_management.ListSqlTuningAdvisorTasksRequest{}

	if managedDatabaseId, ok := s.D.GetOkExists("managed_database_id"); ok {
		tmp := managedDatabaseId.(string)
		request.ManagedDatabaseId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
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
	return nil
}

func (s *DatabaseManagementManagedDatabaseSqlTuningAdvisorTaskDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabaseSqlTuningAdvisorTaskDataSource-", DatabaseManagementManagedDatabaseSqlTuningAdvisorTaskDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SqlTuningAdvisorTaskSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}
