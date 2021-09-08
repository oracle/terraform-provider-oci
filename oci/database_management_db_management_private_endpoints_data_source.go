// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v47/databasemanagement"
)

func init() {
	RegisterDatasource("oci_database_management_db_management_private_endpoints", DatabaseManagementDbManagementPrivateEndpointsDataSource())
}

func DatabaseManagementDbManagementPrivateEndpointsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementDbManagementPrivateEndpoints,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_management_private_endpoint_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     GetDataSourceItemSchema(DatabaseManagementDbManagementPrivateEndpointResource()),
						},
					},
				},
			},
		},
	}
}

func readDatabaseManagementDbManagementPrivateEndpoints(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementDbManagementPrivateEndpointsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dbManagementClient()

	return ReadResource(sync)
}

type DatabaseManagementDbManagementPrivateEndpointsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListDbManagementPrivateEndpointsResponse
}

func (s *DatabaseManagementDbManagementPrivateEndpointsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementDbManagementPrivateEndpointsDataSourceCrud) Get() error {
	request := oci_database_management.ListDbManagementPrivateEndpointsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database_management.ListDbManagementPrivateEndpointsLifecycleStateEnum(state.(string))
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database_management")

	response, err := s.Client.ListDbManagementPrivateEndpoints(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDbManagementPrivateEndpoints(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementDbManagementPrivateEndpointsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("DatabaseManagementDbManagementPrivateEndpointsDataSource-", DatabaseManagementDbManagementPrivateEndpointsDataSource(), s.D))
	resources := []map[string]interface{}{}
	dbManagementPrivateEndpoint := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DbManagementPrivateEndpointSummaryToMap(item))
	}
	dbManagementPrivateEndpoint["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementDbManagementPrivateEndpointsDataSource().Schema["db_management_private_endpoint_collection"].Elem.(*schema.Resource).Schema)
		dbManagementPrivateEndpoint["items"] = items
	}

	resources = append(resources, dbManagementPrivateEndpoint)
	if err := s.D.Set("db_management_private_endpoint_collection", resources); err != nil {
		return err
	}

	return nil
}
