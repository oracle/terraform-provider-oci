// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_migration

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_migration "github.com/oracle/oci-go-sdk/v65/databasemigration"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseMigrationMigrationDataVerificationObjectTypeCountsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDatabaseMigrationMigrationDataVerificationObjectTypeCountsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"migration_data_verification_object_type_count_filter": {
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
			"object_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"schema_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_verification_object_type_count_collection": {
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
									"delta_percent": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"object_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"schema_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"source_invalid_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"source_object_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"target_invalid_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"target_object_count": {
										Type:     schema.TypeInt,
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

func readDatabaseMigrationMigrationDataVerificationObjectTypeCountsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseMigrationMigrationDataVerificationObjectTypeCountsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatabaseMigrationMigrationDataVerificationObjectTypeCountsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_migration.DatabaseMigrationClient
	Res    *oci_database_migration.ListDataVerificationObjectTypeCountsResponse
}

func (s *DatabaseMigrationMigrationDataVerificationObjectTypeCountsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseMigrationMigrationDataVerificationObjectTypeCountsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_migration.ListDataVerificationObjectTypeCountsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if migrationDataVerificationObjectTypeCountFilter, ok := s.D.GetOkExists("migration_data_verification_object_type_count_filter"); ok {
		tmp := migrationDataVerificationObjectTypeCountFilter.(string)
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

	if objectType, ok := s.D.GetOkExists("object_type"); ok {
		tmp := objectType.(string)
		request.ObjectType = &tmp
	}

	if schemaName, ok := s.D.GetOkExists("schema_name"); ok {
		tmp := schemaName.(string)
		request.SchemaName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_migration")

	response, err := s.Client.ListDataVerificationObjectTypeCounts(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDataVerificationObjectTypeCounts(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseMigrationMigrationDataVerificationObjectTypeCountsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseMigrationMigrationDataVerificationObjectTypeCountsDataSource-", DatabaseMigrationMigrationDataVerificationObjectTypeCountsDataSource(), s.D))
	resources := []map[string]interface{}{}
	migrationDataVerificationObjectTypeCount := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DataVerificationObjectTypeCountSummaryToMap(item))
	}
	migrationDataVerificationObjectTypeCount["items"] = items

	if s.Res.TotalItems != nil {
		migrationDataVerificationObjectTypeCount["total_items"] = *s.Res.TotalItems
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseMigrationMigrationDataVerificationObjectTypeCountsDataSource().Schema["data_verification_object_type_count_collection"].Elem.(*schema.Resource).Schema)
		migrationDataVerificationObjectTypeCount["items"] = items
	}

	resources = append(resources, migrationDataVerificationObjectTypeCount)
	if err := s.D.Set("data_verification_object_type_count_collection", resources); err != nil {
		return err
	}

	return nil
}

func DataVerificationObjectTypeCountSummaryToMap(obj oci_database_migration.DataVerificationObjectTypeCountSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_database_migration.MySqlDataVerificationObjectTypeCountSummary:
		result["database_combination"] = "MYSQL"

		result["object_type"] = string(v.ObjectType)

		if v.SourceInvalidCount != nil {
			result["source_invalid_count"] = int(*v.SourceInvalidCount)
		}

		if v.TargetInvalidCount != nil {
			result["target_invalid_count"] = int(*v.TargetInvalidCount)
		}

		if v.DeltaPercent != nil {
			result["delta_percent"] = float64(*v.DeltaPercent)
		}

		if v.SchemaName != nil {
			result["schema_name"] = string(*v.SchemaName)
		}

		if v.SourceObjectCount != nil {
			result["source_object_count"] = int(*v.SourceObjectCount)
		}

		if v.TargetObjectCount != nil {
			result["target_object_count"] = int(*v.TargetObjectCount)
		}
	case oci_database_migration.OracleDataVerificationObjectTypeCountSummary:
		result["database_combination"] = "ORACLE"

		result["object_type"] = string(v.ObjectType)

		if v.SourceInvalidCount != nil {
			result["source_invalid_count"] = int(*v.SourceInvalidCount)
		}

		if v.TargetInvalidCount != nil {
			result["target_invalid_count"] = int(*v.TargetInvalidCount)
		}

		if v.DeltaPercent != nil {
			result["delta_percent"] = float64(*v.DeltaPercent)
		}

		if v.SchemaName != nil {
			result["schema_name"] = string(*v.SchemaName)
		}

		if v.SourceObjectCount != nil {
			result["source_object_count"] = int(*v.SourceObjectCount)
		}

		if v.TargetObjectCount != nil {
			result["target_object_count"] = int(*v.TargetObjectCount)
		}
	default:
		log.Printf("[WARN] Received 'database_combination' of unknown type %v", obj)
		return nil
	}

	return result
}
