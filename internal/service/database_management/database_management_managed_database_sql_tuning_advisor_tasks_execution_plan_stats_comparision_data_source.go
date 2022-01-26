// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"
	"fmt"
	"strconv"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v56/databasemanagement"
)

func DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksExecutionPlanStatsComparisionDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksExecutionPlanStatsComparision,
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
			"modified": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"plan_stats": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"plan_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"plan_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"original": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"plan_stats": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"plan_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"plan_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksExecutionPlanStatsComparision(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksExecutionPlanStatsComparisionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SqlTuningClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksExecutionPlanStatsComparisionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.SqlTuningClient
	Res    *oci_database_management.GetExecutionPlanStatsComparisionResponse
}

func (s *DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksExecutionPlanStatsComparisionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksExecutionPlanStatsComparisionDataSourceCrud) Get() error {
	request := oci_database_management.GetExecutionPlanStatsComparisionRequest{}

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

	response, err := s.Client.GetExecutionPlanStatsComparision(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksExecutionPlanStatsComparisionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksExecutionPlanStatsComparisionDataSource-", DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksExecutionPlanStatsComparisionDataSource(), s.D))

	if s.Res.Modified != nil {
		s.D.Set("modified", []interface{}{SqlTuningTaskPlanStatsToMap(s.Res.Modified)})
	} else {
		s.D.Set("modified", nil)
	}

	if s.Res.Original != nil {
		s.D.Set("original", []interface{}{SqlTuningTaskPlanStatsToMap(s.Res.Original)})
	} else {
		s.D.Set("original", nil)
	}

	return nil
}

func SqlTuningTaskPlanStatsToMap(obj *oci_database_management.SqlTuningTaskPlanStats) map[string]interface{} {
	result := map[string]interface{}{}
	var planStats = make(map[string]string)
	for key, value := range obj.PlanStats {
		planStats[key] = fmt.Sprint(value)
	}

	result["plan_stats"] = planStats

	result["plan_status"] = string(obj.PlanStatus)

	if obj.PlanType != nil {
		result["plan_type"] = string(*obj.PlanType)
	}

	return result
}
