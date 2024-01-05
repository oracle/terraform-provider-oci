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

func DatabaseManagementExternalDbSystemConnectorsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementExternalDbSystemConnectors,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_db_system_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_db_system_connector_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DatabaseManagementExternalDbSystemConnectorResource()),
						},
					},
				},
			},
		},
	}
}

func readDatabaseManagementExternalDbSystemConnectors(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalDbSystemConnectorsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementExternalDbSystemConnectorsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListExternalDbSystemConnectorsResponse
}

func (s *DatabaseManagementExternalDbSystemConnectorsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementExternalDbSystemConnectorsDataSourceCrud) Get() error {
	request := oci_database_management.ListExternalDbSystemConnectorsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if externalDbSystemId, ok := s.D.GetOkExists("external_db_system_id"); ok {
		tmp := externalDbSystemId.(string)
		request.ExternalDbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListExternalDbSystemConnectors(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListExternalDbSystemConnectors(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementExternalDbSystemConnectorsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementExternalDbSystemConnectorsDataSource-", DatabaseManagementExternalDbSystemConnectorsDataSource(), s.D))
	resources := []map[string]interface{}{}
	externalDbSystemConnector := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ExternalDbSystemConnectorSummaryToMap(item))
	}
	externalDbSystemConnector["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementExternalDbSystemConnectorsDataSource().Schema["external_db_system_connector_collection"].Elem.(*schema.Resource).Schema)
		externalDbSystemConnector["items"] = items
	}

	resources = append(resources, externalDbSystemConnector)
	if err := s.D.Set("external_db_system_connector_collection", resources); err != nil {
		return err
	}

	return nil
}
