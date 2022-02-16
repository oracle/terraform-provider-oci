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

func DatabaseManagementDbManagementPrivateEndpointAssociatedDatabaseDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseManagementDbManagementPrivateEndpointAssociatedDatabase,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_management_private_endpoint_id": {
				Type:     schema.TypeString,
				Required: true,
			},
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
	}
}

func readSingularDatabaseManagementDbManagementPrivateEndpointAssociatedDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementDbManagementPrivateEndpointAssociatedDatabaseDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementDbManagementPrivateEndpointAssociatedDatabaseDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListAssociatedDatabasesResponse
}

func (s *DatabaseManagementDbManagementPrivateEndpointAssociatedDatabaseDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementDbManagementPrivateEndpointAssociatedDatabaseDataSourceCrud) Get() error {
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
	return nil
}

func (s *DatabaseManagementDbManagementPrivateEndpointAssociatedDatabaseDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementDbManagementPrivateEndpointAssociatedDatabaseDataSource-", DatabaseManagementDbManagementPrivateEndpointAssociatedDatabaseDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AssociatedDatabaseSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}
