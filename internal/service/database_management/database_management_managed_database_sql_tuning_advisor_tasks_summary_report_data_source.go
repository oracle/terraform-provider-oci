// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v56/common"
	oci_database_management "github.com/oracle/oci-go-sdk/v56/databasemanagement"
)

func DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSummaryReportDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSummaryReport,
		Schema: map[string]*schema.Schema{
			"begin_exec_id_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_exec_id_less_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"managed_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"search_period": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sql_tuning_advisor_task_id": {
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
			// Computed
			"index_findings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"index_columns": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"index_hash_value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"index_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"reference_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"schema": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"table_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"object_stat_findings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"object": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"object_hash_value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"object_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"problem_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"reference_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"schema": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"statistics": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"finding_benefits": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"db_time_after_implemented": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"db_time_after_recommended": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"db_time_before_implemented": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"db_time_before_recommended": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"finding_counts": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"alternate_plan": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"implemented_sql_profile": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"index": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"recommended_sql_profile": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"restructure": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"statistics": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"statement_counts": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"distinct_sql": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"error_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"finding_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"total_sql": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"task_info": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"description": {
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
						"owner": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"running_time": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_ended": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_started": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSummaryReport(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSummaryReportDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SqlTuningClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSummaryReportDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.SqlTuningClient
	Res    *oci_database_management.GetSqlTuningAdvisorTaskSummaryReportResponse
}

func (s *DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSummaryReportDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSummaryReportDataSourceCrud) Get() error {
	request := oci_database_management.GetSqlTuningAdvisorTaskSummaryReportRequest{}

	if beginExecIdGreaterThanOrEqualTo, ok := s.D.GetOkExists("begin_exec_id_greater_than_or_equal_to"); ok {
		tmp := beginExecIdGreaterThanOrEqualTo.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert beginExecIdGreaterThanOrEqualTo string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.BeginExecIdGreaterThanOrEqualTo = &tmpInt64
	}

	if endExecIdLessThanOrEqualTo, ok := s.D.GetOkExists("end_exec_id_less_than_or_equal_to"); ok {
		tmp := endExecIdLessThanOrEqualTo.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert endExecIdLessThanOrEqualTo string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.EndExecIdLessThanOrEqualTo = &tmpInt64
	}

	if managedDatabaseId, ok := s.D.GetOkExists("managed_database_id"); ok {
		tmp := managedDatabaseId.(string)
		request.ManagedDatabaseId = &tmp
	}

	if searchPeriod, ok := s.D.GetOkExists("search_period"); ok {
		request.SearchPeriod = oci_database_management.GetSqlTuningAdvisorTaskSummaryReportSearchPeriodEnum(searchPeriod.(string))
	}

	if sqlTuningAdvisorTaskId, ok := s.D.GetOkExists("sql_tuning_advisor_task_id"); ok {
		tmp := sqlTuningAdvisorTaskId.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert sqlTuningAdvisorTaskId string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.SqlTuningAdvisorTaskId = &tmpInt64
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

	response, err := s.Client.GetSqlTuningAdvisorTaskSummaryReport(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSummaryReportDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSummaryReportDataSource-", DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSummaryReportDataSource(), s.D))

	indexFindings := []interface{}{}
	for _, item := range s.Res.IndexFindings {
		indexFindings = append(indexFindings, SqlTuningAdvisorTaskSummaryReportIndexFindingSummaryToMap(item))
	}
	s.D.Set("index_findings", indexFindings)

	objectStatFindings := []interface{}{}
	for _, item := range s.Res.ObjectStatFindings {
		objectStatFindings = append(objectStatFindings, SqlTuningAdvisorTaskSummaryReportObjectStatFindingSummaryToMap(item))
	}
	s.D.Set("object_stat_findings", objectStatFindings)

	if s.Res.Statistics != nil {
		s.D.Set("statistics", []interface{}{SqlTuningAdvisorTaskSummaryReportStatisticsToMap(s.Res.Statistics)})
	} else {
		s.D.Set("statistics", nil)
	}

	if s.Res.TaskInfo != nil {
		s.D.Set("task_info", []interface{}{SqlTuningAdvisorTaskSummaryReportTaskInfoToMap(s.Res.TaskInfo)})
	} else {
		s.D.Set("task_info", nil)
	}

	return nil
}

func SqlTuningAdvisorTaskSummaryFindingBenefitsToMap(obj *oci_database_management.SqlTuningAdvisorTaskSummaryFindingBenefits) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DbTimeAfterImplemented != nil {
		result["db_time_after_implemented"] = int(*obj.DbTimeAfterImplemented)
	}

	if obj.DbTimeAfterRecommended != nil {
		result["db_time_after_recommended"] = int(*obj.DbTimeAfterRecommended)
	}

	if obj.DbTimeBeforeImplemented != nil {
		result["db_time_before_implemented"] = int(*obj.DbTimeBeforeImplemented)
	}

	if obj.DbTimeBeforeRecommended != nil {
		result["db_time_before_recommended"] = int(*obj.DbTimeBeforeRecommended)
	}

	return result
}

func SqlTuningAdvisorTaskSummaryFindingCountsToMap(obj *oci_database_management.SqlTuningAdvisorTaskSummaryFindingCounts) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AlternatePlan != nil {
		result["alternate_plan"] = int(*obj.AlternatePlan)
	}

	if obj.ImplementedSqlProfile != nil {
		result["implemented_sql_profile"] = int(*obj.ImplementedSqlProfile)
	}

	if obj.Index != nil {
		result["index"] = int(*obj.Index)
	}

	if obj.RecommendedSqlProfile != nil {
		result["recommended_sql_profile"] = int(*obj.RecommendedSqlProfile)
	}

	if obj.Restructure != nil {
		result["restructure"] = int(*obj.Restructure)
	}

	if obj.Statistics != nil {
		result["statistics"] = int(*obj.Statistics)
	}

	return result
}

func SqlTuningAdvisorTaskSummaryReportIndexFindingSummaryToMap(obj oci_database_management.SqlTuningAdvisorTaskSummaryReportIndexFindingSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["index_columns"] = obj.IndexColumns

	if obj.IndexHashValue != nil {
		result["index_hash_value"] = strconv.FormatInt(*obj.IndexHashValue, 10)
	}

	if obj.IndexName != nil {
		result["index_name"] = string(*obj.IndexName)
	}

	if obj.ReferenceCount != nil {
		result["reference_count"] = int(*obj.ReferenceCount)
	}

	if obj.Schema != nil {
		result["schema"] = string(*obj.Schema)
	}

	if obj.TableName != nil {
		result["table_name"] = string(*obj.TableName)
	}

	return result
}

func SqlTuningAdvisorTaskSummaryReportObjectStatFindingSummaryToMap(obj oci_database_management.SqlTuningAdvisorTaskSummaryReportObjectStatFindingSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ObjectName != nil {
		result["object"] = string(*obj.ObjectName)
	}

	if obj.ObjectHashValue != nil {
		result["object_hash_value"] = strconv.FormatInt(*obj.ObjectHashValue, 10)
	}

	if obj.ObjectType != nil {
		result["object_type"] = string(*obj.ObjectType)
	}

	result["problem_type"] = string(obj.ProblemType)

	if obj.ReferenceCount != nil {
		result["reference_count"] = int(*obj.ReferenceCount)
	}

	if obj.Schema != nil {
		result["schema"] = string(*obj.Schema)
	}

	return result
}

func SqlTuningAdvisorTaskSummaryReportStatementCountsToMap(obj *oci_database_management.SqlTuningAdvisorTaskSummaryReportStatementCounts) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DistinctSql != nil {
		result["distinct_sql"] = int(*obj.DistinctSql)
	}

	if obj.ErrorCount != nil {
		result["error_count"] = int(*obj.ErrorCount)
	}

	if obj.FindingCount != nil {
		result["finding_count"] = int(*obj.FindingCount)
	}

	if obj.TotalSql != nil {
		result["total_sql"] = int(*obj.TotalSql)
	}

	return result
}

func SqlTuningAdvisorTaskSummaryReportStatisticsToMap(obj *oci_database_management.SqlTuningAdvisorTaskSummaryReportStatistics) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.FindingBenefits != nil {
		result["finding_benefits"] = []interface{}{SqlTuningAdvisorTaskSummaryFindingBenefitsToMap(obj.FindingBenefits)}
	}

	if obj.FindingCounts != nil {
		result["finding_counts"] = []interface{}{SqlTuningAdvisorTaskSummaryFindingCountsToMap(obj.FindingCounts)}
	}

	if obj.StatementCounts != nil {
		result["statement_counts"] = []interface{}{SqlTuningAdvisorTaskSummaryReportStatementCountsToMap(obj.StatementCounts)}
	}

	return result
}

func SqlTuningAdvisorTaskSummaryReportTaskInfoToMap(obj *oci_database_management.SqlTuningAdvisorTaskSummaryReportTaskInfo) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Id != nil {
		result["id"] = strconv.FormatInt(*obj.Id, 10)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Owner != nil {
		result["owner"] = string(*obj.Owner)
	}

	if obj.RunningTime != nil {
		result["running_time"] = int(*obj.RunningTime)
	}

	result["status"] = string(obj.Status)

	if obj.TimeEnded != nil {
		result["time_ended"] = obj.TimeEnded.String()
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	return result
}
