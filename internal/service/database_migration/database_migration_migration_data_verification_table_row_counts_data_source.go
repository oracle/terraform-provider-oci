// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_migration

import (
	"context"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_migration "github.com/oracle/oci-go-sdk/v65/databasemigration"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseMigrationMigrationDataVerificationTableRowCountsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDatabaseMigrationMigrationDataVerificationTableRowCountsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"migration_data_verification_table_row_count_filter": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"migration_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"min_abs_delta_percent": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"table_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_verification_table_row_count_collection": {
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
									"database_combination": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"owner": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"source_row_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"table_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_row_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"time_last_source_statistics_collection": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_last_target_statistics_collection": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"variance_percent": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},
						"total_items": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDatabaseMigrationMigrationDataVerificationTableRowCountsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseMigrationMigrationDataVerificationTableRowCountsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatabaseMigrationMigrationDataVerificationTableRowCountsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_migration.DatabaseMigrationClient
	Res    *oci_database_migration.ListDataVerificationTableRowCountsResponse
}

func (s *DatabaseMigrationMigrationDataVerificationTableRowCountsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseMigrationMigrationDataVerificationTableRowCountsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_migration.ListDataVerificationTableRowCountsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if migrationDataVerificationTableRowCountFilter, ok := s.D.GetOkExists("migration_data_verification_table_row_count_filter"); ok {
		tmp := migrationDataVerificationTableRowCountFilter.(string)
		request.Filter = &tmp
	}

	if migrationId, ok := s.D.GetOkExists("migration_id"); ok {
		tmp := migrationId.(string)
		request.MigrationId = &tmp
	}

	if minAbsDeltaPercent, ok := s.D.GetOkExists("min_abs_delta_percent"); ok {
		tmp := minAbsDeltaPercent.(float64)
		request.MinAbsDeltaPercent = &tmp
	}

	if owner, ok := s.D.GetOkExists("owner"); ok {
		tmp := owner.(string)
		request.Owner = &tmp
	}

	if tableName, ok := s.D.GetOkExists("table_name"); ok {
		tmp := tableName.(string)
		request.TableName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_migration")

	response, err := s.Client.ListDataVerificationTableRowCounts(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDataVerificationTableRowCounts(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseMigrationMigrationDataVerificationTableRowCountsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseMigrationMigrationDataVerificationTableRowCountsDataSource-", DatabaseMigrationMigrationDataVerificationTableRowCountsDataSource(), s.D))
	resources := []map[string]interface{}{}
	migrationDataVerificationTableRowCount := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DataVerificationTableRowCountSummaryToMap(item))
	}
	migrationDataVerificationTableRowCount["items"] = items

	if s.Res.TotalItems != nil {
		migrationDataVerificationTableRowCount["total_items"] = *s.Res.TotalItems
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseMigrationMigrationDataVerificationTableRowCountsDataSource().Schema["data_verification_table_row_count_collection"].Elem.(*schema.Resource).Schema)
		migrationDataVerificationTableRowCount["items"] = items
	}

	resources = append(resources, migrationDataVerificationTableRowCount)
	if err := s.D.Set("data_verification_table_row_count_collection", resources); err != nil {
		return err
	}

	return nil
}

func DataVerificationTableRowCountSummaryToMap(obj oci_database_migration.DataVerificationTableRowCountSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_database_migration.MySqlDataVerificationTableRowCountSummary:
		result["database_combination"] = "MYSQL"

		if v.TimeLastSourceStatisticsCollection != nil {
			result["time_last_source_statistics_collection"] = v.TimeLastSourceStatisticsCollection.Format(time.RFC3339Nano)
		}

		if v.TimeLastTargetStatisticsCollection != nil {
			result["time_last_target_statistics_collection"] = v.TimeLastTargetStatisticsCollection.Format(time.RFC3339Nano)
		}

		if v.Owner != nil {
			result["owner"] = string(*v.Owner)
		}

		if v.SourceRowCount != nil {
			result["source_row_count"] = int(*v.SourceRowCount)
		}

		if v.TableName != nil {
			result["table_name"] = string(*v.TableName)
		}

		if v.TargetRowCount != nil {
			result["target_row_count"] = int(*v.TargetRowCount)
		}

		if v.VariancePercent != nil {
			result["variance_percent"] = float64(*v.VariancePercent)
		}
	case oci_database_migration.OracleDataVerificationTableRowCountSummary:
		result["database_combination"] = "ORACLE"

		if v.TimeLastSourceStatisticsCollection != nil {
			result["time_last_source_statistics_collection"] = v.TimeLastSourceStatisticsCollection.Format(time.RFC3339Nano)
		}

		if v.TimeLastTargetStatisticsCollection != nil {
			result["time_last_target_statistics_collection"] = v.TimeLastTargetStatisticsCollection.Format(time.RFC3339Nano)
		}

		if v.Owner != nil {
			result["owner"] = string(*v.Owner)
		}

		if v.SourceRowCount != nil {
			result["source_row_count"] = int(*v.SourceRowCount)
		}

		if v.TableName != nil {
			result["table_name"] = string(*v.TableName)
		}

		if v.TargetRowCount != nil {
			result["target_row_count"] = int(*v.TargetRowCount)
		}

		if v.VariancePercent != nil {
			result["variance_percent"] = float64(*v.VariancePercent)
		}
	default:
		log.Printf("[WARN] Received 'database_combination' of unknown type %v", obj)
		return nil
	}

	return result
}
