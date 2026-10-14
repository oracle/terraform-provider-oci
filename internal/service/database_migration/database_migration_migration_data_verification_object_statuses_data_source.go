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

func DatabaseMigrationMigrationDataVerificationObjectStatusesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDatabaseMigrationMigrationDataVerificationObjectStatusesWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_match": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"migration_data_verification_object_status_filter": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"migration_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"object_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_verification_object_status_collection": {
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
									"object": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"object_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"owner": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status_in_source": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status_in_target": {
										Type:     schema.TypeString,
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

func readDatabaseMigrationMigrationDataVerificationObjectStatusesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseMigrationMigrationDataVerificationObjectStatusesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatabaseMigrationMigrationDataVerificationObjectStatusesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_migration.DatabaseMigrationClient
	Res    *oci_database_migration.ListDataVerificationObjectStatusesResponse
}

func (s *DatabaseMigrationMigrationDataVerificationObjectStatusesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseMigrationMigrationDataVerificationObjectStatusesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_migration.ListDataVerificationObjectStatusesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if isMatch, ok := s.D.GetOkExists("is_match"); ok {
		tmp := isMatch.(bool)
		request.IsMatch = &tmp
	}

	if migrationDataVerificationObjectStatusFilter, ok := s.D.GetOkExists("migration_data_verification_object_status_filter"); ok {
		tmp := migrationDataVerificationObjectStatusFilter.(string)
		request.Filter = &tmp
	}

	if migrationId, ok := s.D.GetOkExists("migration_id"); ok {
		tmp := migrationId.(string)
		request.MigrationId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if objectType, ok := s.D.GetOkExists("object_type"); ok {
		tmp := objectType.(string)
		request.ObjectType = &tmp
	}

	if owner, ok := s.D.GetOkExists("owner"); ok {
		tmp := owner.(string)
		request.Owner = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_migration")

	response, err := s.Client.ListDataVerificationObjectStatuses(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDataVerificationObjectStatuses(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseMigrationMigrationDataVerificationObjectStatusesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseMigrationMigrationDataVerificationObjectStatusesDataSource-", DatabaseMigrationMigrationDataVerificationObjectStatusesDataSource(), s.D))
	resources := []map[string]interface{}{}
	migrationDataVerificationObjectStatus := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DataVerificationObjectStatusSummaryToMap(item))
	}
	migrationDataVerificationObjectStatus["items"] = items

	if s.Res.TotalItems != nil {
		migrationDataVerificationObjectStatus["total_items"] = *s.Res.TotalItems
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseMigrationMigrationDataVerificationObjectStatusesDataSource().Schema["data_verification_object_status_collection"].Elem.(*schema.Resource).Schema)
		migrationDataVerificationObjectStatus["items"] = items
	}

	resources = append(resources, migrationDataVerificationObjectStatus)
	if err := s.D.Set("data_verification_object_status_collection", resources); err != nil {
		return err
	}

	return nil
}

func DataVerificationObjectStatusSummaryToMap(obj oci_database_migration.DataVerificationObjectStatusSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_database_migration.MySqlDataVerificationObjectStatusSummary:
		result["database_combination"] = "MYSQL"

		result["object_type"] = string(v.ObjectType)

		if v.ObjectName != nil {
			result["object"] = string(*v.ObjectName)
		}

		if v.Owner != nil {
			result["owner"] = string(*v.Owner)
		}

		if v.StatusInSource != nil {
			result["status_in_source"] = string(*v.StatusInSource)
		}

		if v.StatusInTarget != nil {
			result["status_in_target"] = string(*v.StatusInTarget)
		}
	case oci_database_migration.OracleDataVerificationObjectStatusSummary:
		result["database_combination"] = "ORACLE"

		result["object_type"] = string(v.ObjectType)

		if v.ObjectName != nil {
			result["object"] = string(*v.ObjectName)
		}

		if v.Owner != nil {
			result["owner"] = string(*v.Owner)
		}

		if v.StatusInSource != nil {
			result["status_in_source"] = string(*v.StatusInSource)
		}

		if v.StatusInTarget != nil {
			result["status_in_target"] = string(*v.StatusInTarget)
		}
	default:
		log.Printf("[WARN] Received 'database_combination' of unknown type %v", obj)
		return nil
	}

	return result
}
