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

func DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSqlExecutionPlanDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSqlExecutionPlan,
		Schema: map[string]*schema.Schema{
			"attribute": {
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
			"plan": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"access_predicates": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"attribute": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"bytes": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cardinality": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cost": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"cpu_cost": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"filter_predicates": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"io_cost": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"number_of_search_column": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"object": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"object_node": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"object_owner": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"object_position": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"object_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"operation": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"optimizer_mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"options": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"other": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"other_tag": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"parent_step_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"partition_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"partition_start": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"partition_stop": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"plan_hash_value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"position": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"remarks": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"step_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"temp_space": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSqlExecutionPlan(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSqlExecutionPlanDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SqlTuningClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSqlExecutionPlanDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.SqlTuningClient
	Res    *oci_database_management.GetSqlExecutionPlanResponse
}

func (s *DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSqlExecutionPlanDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSqlExecutionPlanDataSourceCrud) Get() error {
	request := oci_database_management.GetSqlExecutionPlanRequest{}

	if attribute, ok := s.D.GetOkExists("attribute"); ok {
		request.Attribute = oci_database_management.GetSqlExecutionPlanAttributeEnum(attribute.(string))
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

	response, err := s.Client.GetSqlExecutionPlan(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSqlExecutionPlanDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSqlExecutionPlanDataSource-", DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSqlExecutionPlanDataSource(), s.D))

	plan := []interface{}{}
	for _, item := range s.Res.Plan {
		plan = append(plan, SqlTuningTaskSqlExecutionPlanStepToMap(item))
	}
	s.D.Set("plan", plan)

	return nil
}

func SqlTuningTaskSqlExecutionPlanStepToMap(obj oci_database_management.SqlTuningTaskSqlExecutionPlanStep) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AccessPredicates != nil {
		result["access_predicates"] = string(*obj.AccessPredicates)
	}

	if obj.Attribute != nil {
		result["attribute"] = string(*obj.Attribute)
	}

	if obj.Bytes != nil {
		result["bytes"] = strconv.FormatInt(*obj.Bytes, 10)
	}

	if obj.Cardinality != nil {
		result["cardinality"] = strconv.FormatInt(*obj.Cardinality, 10)
	}

	if obj.Cost != nil {
		result["cost"] = float64(*obj.Cost)
	}

	if obj.CpuCost != nil {
		result["cpu_cost"] = float64(*obj.CpuCost)
	}

	if obj.FilterPredicates != nil {
		result["filter_predicates"] = string(*obj.FilterPredicates)
	}

	if obj.IoCost != nil {
		result["io_cost"] = float64(*obj.IoCost)
	}

	if obj.NumberOfSearchColumn != nil {
		result["number_of_search_column"] = int(*obj.NumberOfSearchColumn)
	}

	if obj.ObjectName != nil {
		result["object"] = string(*obj.ObjectName)
	}

	if obj.ObjectNode != nil {
		result["object_node"] = string(*obj.ObjectNode)
	}

	if obj.ObjectOwner != nil {
		result["object_owner"] = string(*obj.ObjectOwner)
	}

	if obj.ObjectPosition != nil {
		result["object_position"] = int(*obj.ObjectPosition)
	}

	if obj.ObjectType != nil {
		result["object_type"] = string(*obj.ObjectType)
	}

	if obj.Operation != nil {
		result["operation"] = string(*obj.Operation)
	}

	if obj.OptimizerMode != nil {
		result["optimizer_mode"] = string(*obj.OptimizerMode)
	}

	if obj.Options != nil {
		result["options"] = string(*obj.Options)
	}

	if obj.Other != nil {
		result["other"] = string(*obj.Other)
	}

	if obj.OtherTag != nil {
		result["other_tag"] = string(*obj.OtherTag)
	}

	if obj.ParentStepId != nil {
		result["parent_step_id"] = int(*obj.ParentStepId)
	}

	if obj.PartitionId != nil {
		result["partition_id"] = int(*obj.PartitionId)
	}

	if obj.PartitionStart != nil {
		result["partition_start"] = string(*obj.PartitionStart)
	}

	if obj.PartitionStop != nil {
		result["partition_stop"] = string(*obj.PartitionStop)
	}

	if obj.PlanHashValue != nil {
		result["plan_hash_value"] = strconv.FormatInt(*obj.PlanHashValue, 10)
	}

	if obj.Position != nil {
		result["position"] = int(*obj.Position)
	}

	if obj.Remarks != nil {
		result["remarks"] = string(*obj.Remarks)
	}

	if obj.StepId != nil {
		result["step_id"] = int(*obj.StepId)
	}

	if obj.TempSpace != nil {
		result["temp_space"] = strconv.FormatInt(*obj.TempSpace, 10)
	}

	if obj.Time != nil {
		result["time"] = strconv.FormatInt(*obj.Time, 10)
	}

	return result
}
