// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"
	"fmt"
	"strconv"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"
)

func DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksRecommendationDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksRecommendation,
		Schema: map[string]*schema.Schema{
			"execution_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"managed_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sql_object_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sql_tuning_advisor_task_id": {
				Type:     schema.TypeString,
				Required: true,
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
						"benefit": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"finding": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"implement_action_sql": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_parallel_execution": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"rationale": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"recommendation": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"recommendation_key": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"recommendation_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"sql_tuning_advisor_task_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"sql_tuning_advisor_task_object_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
		DeprecationMessage: tfresource.DatasourceDeprecatedForAnother("oci_database_management_managed_database_sql_tuning_advisor_tasks_recommendation", "oci_database_management_managed_database_sql_tuning_advisor_tasks_recommendations"),
	}
}

func readSingularDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksRecommendation(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksRecommendationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SqlTuningClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksRecommendationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.SqlTuningClient
	Res    *oci_database_management.ListSqlTuningAdvisorTaskRecommendationsResponse
}

func (s *DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksRecommendationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksRecommendationDataSourceCrud) Get() error {
	request := oci_database_management.ListSqlTuningAdvisorTaskRecommendationsRequest{}

	if executionId, ok := s.D.GetOkExists("execution_id"); ok {
		tmp := executionId.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert executionId string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.ExecutionId = &tmpInt64
	}

	if managedDatabaseId, ok := s.D.GetOkExists("managed_database_id"); ok {
		tmp := managedDatabaseId.(string)
		request.ManagedDatabaseId = &tmp
	}

	if sqlObjectId, ok := s.D.GetOkExists("sql_object_id"); ok {
		tmp := sqlObjectId.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert sqlObjectId string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.SqlObjectId = &tmpInt64
	}

	if sqlTuningAdvisorTaskId, ok := s.D.GetOkExists("sql_tuning_advisor_task_id"); ok {
		tmp := sqlTuningAdvisorTaskId.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert sqlTuningAdvisorTaskId string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.SqlTuningAdvisorTaskId = &tmpInt64
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListSqlTuningAdvisorTaskRecommendations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksRecommendationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksRecommendationDataSource-", DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksRecommendationDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SqlTuningAdvisorTaskRecommendationSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}
