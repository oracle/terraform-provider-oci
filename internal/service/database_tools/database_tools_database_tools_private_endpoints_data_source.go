// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_tools

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database_tools "github.com/oracle/oci-go-sdk/v56/databasetools"
)

func DatabaseToolsDatabaseToolsPrivateEndpointsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseToolsDatabaseToolsPrivateEndpoints,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"endpoint_service_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"database_tools_private_endpoint_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DatabaseToolsDatabaseToolsPrivateEndpointResource()),
						},
					},
				},
			},
		},
	}
}

func readDatabaseToolsDatabaseToolsPrivateEndpoints(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseToolsDatabaseToolsPrivateEndpointsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()

	return tfresource.ReadResource(sync)
}

type DatabaseToolsDatabaseToolsPrivateEndpointsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_tools.DatabaseToolsClient
	Res    *oci_database_tools.ListDatabaseToolsPrivateEndpointsResponse
}

func (s *DatabaseToolsDatabaseToolsPrivateEndpointsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseToolsDatabaseToolsPrivateEndpointsDataSourceCrud) Get() error {
	request := oci_database_tools.ListDatabaseToolsPrivateEndpointsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if endpointServiceId, ok := s.D.GetOkExists("endpoint_service_id"); ok {
		tmp := endpointServiceId.(string)
		request.EndpointServiceId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database_tools.ListDatabaseToolsPrivateEndpointsLifecycleStateEnum(state.(string))
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_tools")

	response, err := s.Client.ListDatabaseToolsPrivateEndpoints(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDatabaseToolsPrivateEndpoints(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseToolsDatabaseToolsPrivateEndpointsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseToolsDatabaseToolsPrivateEndpointsDataSource-", DatabaseToolsDatabaseToolsPrivateEndpointsDataSource(), s.D))
	resources := []map[string]interface{}{}
	databaseToolsPrivateEndpoint := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DatabaseToolsPrivateEndpointSummaryToMap(item, true))
	}
	databaseToolsPrivateEndpoint["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseToolsDatabaseToolsPrivateEndpointsDataSource().Schema["database_tools_private_endpoint_collection"].Elem.(*schema.Resource).Schema)
		databaseToolsPrivateEndpoint["items"] = items
	}

	resources = append(resources, databaseToolsPrivateEndpoint)
	if err := s.D.Set("database_tools_private_endpoint_collection", resources); err != nil {
		return err
	}

	return nil
}
