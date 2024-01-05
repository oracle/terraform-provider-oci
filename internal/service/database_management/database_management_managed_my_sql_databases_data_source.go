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

func DatabaseManagementManagedMySqlDatabasesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementManagedMySqlDatabases,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"managed_my_sql_database_collection": {
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
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"db_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"db_version": {
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
									"time_created": {
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

func readDatabaseManagementManagedMySqlDatabases(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedMySqlDatabasesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedMySqlDatabasesClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedMySqlDatabasesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.ManagedMySqlDatabasesClient
	Res    *oci_database_management.ListManagedMySqlDatabasesResponse
}

func (s *DatabaseManagementManagedMySqlDatabasesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedMySqlDatabasesDataSourceCrud) Get() error {
	request := oci_database_management.ListManagedMySqlDatabasesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListManagedMySqlDatabases(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListManagedMySqlDatabases(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementManagedMySqlDatabasesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedMySqlDatabasesDataSource-", DatabaseManagementManagedMySqlDatabasesDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedMySqlDatabase := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ManagedMySqlDatabaseSummaryToMap(item))
	}
	managedMySqlDatabase["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementManagedMySqlDatabasesDataSource().Schema["managed_my_sql_database_collection"].Elem.(*schema.Resource).Schema)
		managedMySqlDatabase["items"] = items
	}

	resources = append(resources, managedMySqlDatabase)
	if err := s.D.Set("managed_my_sql_database_collection", resources); err != nil {
		return err
	}

	return nil
}

func ManagedMySqlDatabaseSummaryToMap(obj oci_database_management.ManagedMySqlDatabaseSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DbName != nil {
		result["db_name"] = string(*obj.DbName)
	}

	if obj.DbVersion != nil {
		result["db_version"] = string(*obj.DbVersion)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}
