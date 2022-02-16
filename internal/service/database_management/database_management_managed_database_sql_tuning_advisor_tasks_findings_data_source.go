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
	oci_database_management "github.com/oracle/oci-go-sdk/v58/databasemanagement"
)

func DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksFindingsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksFindings,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"begin_exec_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_exec_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"finding_filter": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"index_hash_filter": {
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
			"stats_hash_filter": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sql_tuning_advisor_task_finding_collection": {
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
									"db_time_benefit": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"is_alternative_plan_finding_present": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_error_finding_present": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_index_finding_present": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_miscellaneous_finding_present": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_restructure_sql_finding_present": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_sql_profile_finding_implemented": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_sql_profile_finding_present": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_stats_finding_present": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_timeout_finding_present": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"parsing_schema": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"per_execution_percentage": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"sql_key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"sql_text": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"sql_tuning_advisor_task_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"sql_tuning_advisor_task_object_execution_id": {
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
				},
			},
		},
	}
}

func readDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksFindings(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksFindingsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SqlTuningClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksFindingsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.SqlTuningClient
	Res    *oci_database_management.ListSqlTuningAdvisorTaskFindingsResponse
}

func (s *DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksFindingsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksFindingsDataSourceCrud) Get() error {
	request := oci_database_management.ListSqlTuningAdvisorTaskFindingsRequest{}

	if beginExecId, ok := s.D.GetOkExists("begin_exec_id"); ok {
		tmp := beginExecId.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert beginExecId string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.BeginExecId = &tmpInt64
	}

	if endExecId, ok := s.D.GetOkExists("end_exec_id"); ok {
		tmp := endExecId.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert endExecId string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.EndExecId = &tmpInt64
	}

	if findingFilter, ok := s.D.GetOkExists("finding_filter"); ok {
		request.FindingFilter = oci_database_management.ListSqlTuningAdvisorTaskFindingsFindingFilterEnum(findingFilter.(string))
	}

	if indexHashFilter, ok := s.D.GetOkExists("index_hash_filter"); ok {
		tmp := indexHashFilter.(string)
		request.IndexHashFilter = &tmp
	}

	if managedDatabaseId, ok := s.D.GetOkExists("managed_database_id"); ok {
		tmp := managedDatabaseId.(string)
		request.ManagedDatabaseId = &tmp
	}

	if searchPeriod, ok := s.D.GetOkExists("search_period"); ok {
		request.SearchPeriod = oci_database_management.ListSqlTuningAdvisorTaskFindingsSearchPeriodEnum(searchPeriod.(string))
	}

	if sqlTuningAdvisorTaskId, ok := s.D.GetOkExists("sql_tuning_advisor_task_id"); ok {
		tmp := sqlTuningAdvisorTaskId.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert sqlTuningAdvisorTaskId string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.SqlTuningAdvisorTaskId = &tmpInt64
	}

	if statsHashFilter, ok := s.D.GetOkExists("stats_hash_filter"); ok {
		tmp := statsHashFilter.(string)
		request.StatsHashFilter = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListSqlTuningAdvisorTaskFindings(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSqlTuningAdvisorTaskFindings(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksFindingsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksFindingsDataSource-", DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksFindingsDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedDatabaseSqlTuningAdvisorTasksFinding := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SqlTuningAdvisorTaskFindingSummaryToMap(item))
	}
	managedDatabaseSqlTuningAdvisorTasksFinding["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksFindingsDataSource().Schema["sql_tuning_advisor_task_finding_collection"].Elem.(*schema.Resource).Schema)
		managedDatabaseSqlTuningAdvisorTasksFinding["items"] = items
	}

	resources = append(resources, managedDatabaseSqlTuningAdvisorTasksFinding)
	if err := s.D.Set("sql_tuning_advisor_task_finding_collection", resources); err != nil {
		return err
	}

	return nil
}
