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

func DatabaseManagementManagedMySqlDatabaseSqlDataDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementManagedMySqlDatabaseSqlData,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"end_time": {
				Type:     schema.TypeString,
				Required: true,
			},
			"filter_column": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"managed_my_sql_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"start_time": {
				Type:     schema.TypeString,
				Required: true,
			},
			"my_sql_data_collection": {
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
									"avg_timer_wait": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"count_star": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"digest": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"digest_text": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"first_seen": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"last_seen": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"max_timer_wait": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"min_timer_wait": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"quantile95": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"quantile99": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"quantile999": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"schema_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"sum_created_temp_disk_tables": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"sum_created_temp_tables": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"sum_errors": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"sum_lock_time": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"sum_no_good_index_used": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"sum_no_index_used": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"sum_rows_affected": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"sum_rows_examined": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"sum_rows_sent": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"sum_select_full_join": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"sum_select_full_range_join": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"sum_select_range": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"sum_select_range_check": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"sum_select_scan": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"sum_sort_merge_passes": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"sum_sort_range": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"sum_sort_rows": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"sum_sort_scan": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"sum_timer_wait": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"sum_warnings": {
										Type:     schema.TypeFloat,
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

func readDatabaseManagementManagedMySqlDatabaseSqlData(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedMySqlDatabaseSqlDataDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedMySqlDatabasesClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedMySqlDatabaseSqlDataDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.ManagedMySqlDatabasesClient
	Res    *oci_database_management.ListManagedMySqlDatabaseSqlDataResponse
}

func (s *DatabaseManagementManagedMySqlDatabaseSqlDataDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedMySqlDatabaseSqlDataDataSourceCrud) Get() error {
	request := oci_database_management.ListManagedMySqlDatabaseSqlDataRequest{}

	if endTime, ok := s.D.GetOkExists("end_time"); ok {
		tmp := endTime.(string)
		request.EndTime = &tmp
	}

	if filterColumn, ok := s.D.GetOkExists("filter_column"); ok {
		tmp := filterColumn.(string)
		request.FilterColumn = &tmp
	}

	if managedMySqlDatabaseId, ok := s.D.GetOkExists("managed_my_sql_database_id"); ok {
		tmp := managedMySqlDatabaseId.(string)
		request.ManagedMySqlDatabaseId = &tmp
	}

	if startTime, ok := s.D.GetOkExists("start_time"); ok {
		tmp := startTime.(string)
		request.StartTime = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListManagedMySqlDatabaseSqlData(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListManagedMySqlDatabaseSqlData(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementManagedMySqlDatabaseSqlDataDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedMySqlDatabaseSqlDataDataSource-", DatabaseManagementManagedMySqlDatabaseSqlDataDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedMySqlDatabaseSqlData := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MySqlDataSummaryToMap(item))
	}
	managedMySqlDatabaseSqlData["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementManagedMySqlDatabaseSqlDataDataSource().Schema["my_sql_data_collection"].Elem.(*schema.Resource).Schema)
		managedMySqlDatabaseSqlData["items"] = items
	}

	resources = append(resources, managedMySqlDatabaseSqlData)
	if err := s.D.Set("my_sql_data_collection", resources); err != nil {
		return err
	}

	return nil
}

func MySqlDataSummaryToMap(obj oci_database_management.MySqlDataSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvgTimerWait != nil {
		result["avg_timer_wait"] = float32(*obj.AvgTimerWait)
	}

	if obj.CountStar != nil {
		result["count_star"] = float32(*obj.CountStar)
	}

	if obj.Digest != nil {
		result["digest"] = string(*obj.Digest)
	}

	if obj.DigestText != nil {
		result["digest_text"] = string(*obj.DigestText)
	}

	if obj.FirstSeen != nil {
		result["first_seen"] = obj.FirstSeen.String()
	}

	if obj.LastSeen != nil {
		result["last_seen"] = obj.LastSeen.String()
	}

	if obj.MaxTimerWait != nil {
		result["max_timer_wait"] = float32(*obj.MaxTimerWait)
	}

	if obj.MinTimerWait != nil {
		result["min_timer_wait"] = float32(*obj.MinTimerWait)
	}

	if obj.Quantile95 != nil {
		result["quantile95"] = float32(*obj.Quantile95)
	}

	if obj.Quantile99 != nil {
		result["quantile99"] = float32(*obj.Quantile99)
	}

	if obj.Quantile999 != nil {
		result["quantile999"] = float32(*obj.Quantile999)
	}

	if obj.SchemaName != nil {
		result["schema_name"] = string(*obj.SchemaName)
	}

	if obj.SumCreatedTempDiskTables != nil {
		result["sum_created_temp_disk_tables"] = float32(*obj.SumCreatedTempDiskTables)
	}

	if obj.SumCreatedTempTables != nil {
		result["sum_created_temp_tables"] = float32(*obj.SumCreatedTempTables)
	}

	if obj.SumErrors != nil {
		result["sum_errors"] = float32(*obj.SumErrors)
	}

	if obj.SumLockTime != nil {
		result["sum_lock_time"] = float32(*obj.SumLockTime)
	}

	if obj.SumNoGoodIndexUsed != nil {
		result["sum_no_good_index_used"] = float32(*obj.SumNoGoodIndexUsed)
	}

	if obj.SumNoIndexUsed != nil {
		result["sum_no_index_used"] = float32(*obj.SumNoIndexUsed)
	}

	if obj.SumRowsAffected != nil {
		result["sum_rows_affected"] = float32(*obj.SumRowsAffected)
	}

	if obj.SumRowsExamined != nil {
		result["sum_rows_examined"] = float32(*obj.SumRowsExamined)
	}

	if obj.SumRowsSent != nil {
		result["sum_rows_sent"] = float32(*obj.SumRowsSent)
	}

	if obj.SumSelectFullJoin != nil {
		result["sum_select_full_join"] = float32(*obj.SumSelectFullJoin)
	}

	if obj.SumSelectFullRangeJoin != nil {
		result["sum_select_full_range_join"] = float32(*obj.SumSelectFullRangeJoin)
	}

	if obj.SumSelectRange != nil {
		result["sum_select_range"] = float32(*obj.SumSelectRange)
	}

	if obj.SumSelectRangeCheck != nil {
		result["sum_select_range_check"] = float32(*obj.SumSelectRangeCheck)
	}

	if obj.SumSelectScan != nil {
		result["sum_select_scan"] = float32(*obj.SumSelectScan)
	}

	if obj.SumSortMergePasses != nil {
		result["sum_sort_merge_passes"] = float32(*obj.SumSortMergePasses)
	}

	if obj.SumSortRange != nil {
		result["sum_sort_range"] = float32(*obj.SumSortRange)
	}

	if obj.SumSortRows != nil {
		result["sum_sort_rows"] = float32(*obj.SumSortRows)
	}

	if obj.SumSortScan != nil {
		result["sum_sort_scan"] = float32(*obj.SumSortScan)
	}

	if obj.SumTimerWait != nil {
		result["sum_timer_wait"] = float32(*obj.SumTimerWait)
	}

	if obj.SumWarnings != nil {
		result["sum_warnings"] = float32(*obj.SumWarnings)
	}

	return result
}
