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

func DatabaseManagementManagedDatabaseSqlTuningSetsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementManagedDatabaseSqlTuningSets,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"managed_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sql_tuning_set_collection": {
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
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"error_message": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
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
									"scheduled_job_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"statement_counts": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_last_modified": {
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

func readDatabaseManagementManagedDatabaseSqlTuningSets(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseSqlTuningSetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SqlTuningClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabaseSqlTuningSetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.SqlTuningClient
	Res    *oci_database_management.ListSqlTuningSetsResponse
}

func (s *DatabaseManagementManagedDatabaseSqlTuningSetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseSqlTuningSetsDataSourceCrud) Get() error {
	request := oci_database_management.ListSqlTuningSetsRequest{}

	if managedDatabaseId, ok := s.D.GetOkExists("managed_database_id"); ok {
		tmp := managedDatabaseId.(string)
		request.ManagedDatabaseId = &tmp
	}

	if nameContains, ok := s.D.GetOkExists("name_contains"); ok {
		tmp := nameContains.(string)
		request.NameContains = &tmp
	}

	if owner, ok := s.D.GetOkExists("owner"); ok {
		tmp := owner.(string)
		request.Owner = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListSqlTuningSets(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSqlTuningSets(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementManagedDatabaseSqlTuningSetsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabaseSqlTuningSetsDataSource-", DatabaseManagementManagedDatabaseSqlTuningSetsDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedDatabaseSqlTuningSet := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SqlTuningSetSummaryToMap(item))
	}
	managedDatabaseSqlTuningSet["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementManagedDatabaseSqlTuningSetsDataSource().Schema["sql_tuning_set_collection"].Elem.(*schema.Resource).Schema)
		managedDatabaseSqlTuningSet["items"] = items
	}

	resources = append(resources, managedDatabaseSqlTuningSet)
	if err := s.D.Set("sql_tuning_set_collection", resources); err != nil {
		return err
	}

	return nil
}

func SqlTuningSetSummaryToMap(obj oci_database_management.SqlTuningSetSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.ErrorMessage != nil {
		result["error_message"] = string(*obj.ErrorMessage)
	}

	if obj.Id != nil {
		result["id"] = int(*obj.Id)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Owner != nil {
		result["owner"] = string(*obj.Owner)
	}

	if obj.ScheduledJobName != nil {
		result["scheduled_job_name"] = string(*obj.ScheduledJobName)
	}

	if obj.StatementCounts != nil {
		result["statement_counts"] = int(*obj.StatementCounts)
	}

	result["status"] = string(obj.Status)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeLastModified != nil {
		result["time_last_modified"] = obj.TimeLastModified.String()
	}

	return result
}
