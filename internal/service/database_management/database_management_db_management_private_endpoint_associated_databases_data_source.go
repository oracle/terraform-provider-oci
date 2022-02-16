// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v58/databasemanagement"
)

func DatabaseManagementDbManagementPrivateEndpointAssociatedDatabasesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementDbManagementPrivateEndpointAssociatedDatabases,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_management_private_endpoint_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"associated_database_collection": {
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
									"compartment_id": {
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
									"time_registered": {
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

func readDatabaseManagementDbManagementPrivateEndpointAssociatedDatabases(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementDbManagementPrivateEndpointAssociatedDatabasesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementDbManagementPrivateEndpointAssociatedDatabasesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListAssociatedDatabasesResponse
}

func (s *DatabaseManagementDbManagementPrivateEndpointAssociatedDatabasesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementDbManagementPrivateEndpointAssociatedDatabasesDataSourceCrud) Get() error {
	request := oci_database_management.ListAssociatedDatabasesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dbManagementPrivateEndpointId, ok := s.D.GetOkExists("db_management_private_endpoint_id"); ok {
		tmp := dbManagementPrivateEndpointId.(string)
		request.DbManagementPrivateEndpointId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListAssociatedDatabases(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAssociatedDatabases(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementDbManagementPrivateEndpointAssociatedDatabasesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementDbManagementPrivateEndpointAssociatedDatabasesDataSource-", DatabaseManagementDbManagementPrivateEndpointAssociatedDatabasesDataSource(), s.D))
	resources := []map[string]interface{}{}
	dbManagementPrivateEndpointAssociatedDatabase := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AssociatedDatabaseSummaryToMap(item))
	}
	dbManagementPrivateEndpointAssociatedDatabase["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementDbManagementPrivateEndpointAssociatedDatabasesDataSource().Schema["associated_database_collection"].Elem.(*schema.Resource).Schema)
		dbManagementPrivateEndpointAssociatedDatabase["items"] = items
	}

	resources = append(resources, dbManagementPrivateEndpointAssociatedDatabase)
	if err := s.D.Set("associated_database_collection", resources); err != nil {
		return err
	}

	return nil
}

func AssociatedDatabaseSummaryToMap(obj oci_database_management.AssociatedDatabaseSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.TimeRegistered != nil {
		result["time_registered"] = obj.TimeRegistered.String()
	}

	return result
}
