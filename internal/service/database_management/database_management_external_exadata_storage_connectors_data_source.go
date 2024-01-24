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

func DatabaseManagementExternalExadataStorageConnectorsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementExternalExadataStorageConnectors,
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
			"external_exadata_infrastructure_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"external_exadata_storage_connector_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DatabaseManagementExternalExadataStorageConnectorResource()),
						},
					},
				},
			},
		},
	}
}

func readDatabaseManagementExternalExadataStorageConnectors(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExternalExadataStorageConnectorsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementExternalExadataStorageConnectorsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListExternalExadataStorageConnectorsResponse
}

func (s *DatabaseManagementExternalExadataStorageConnectorsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementExternalExadataStorageConnectorsDataSourceCrud) Get() error {
	request := oci_database_management.ListExternalExadataStorageConnectorsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if externalExadataInfrastructureId, ok := s.D.GetOkExists("external_exadata_infrastructure_id"); ok {
		tmp := externalExadataInfrastructureId.(string)
		request.ExternalExadataInfrastructureId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListExternalExadataStorageConnectors(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListExternalExadataStorageConnectors(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementExternalExadataStorageConnectorsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementExternalExadataStorageConnectorsDataSource-", DatabaseManagementExternalExadataStorageConnectorsDataSource(), s.D))
	resources := []map[string]interface{}{}
	externalExadataStorageConnector := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ExternalExadataStorageConnectorSummaryToMap(item))
	}
	externalExadataStorageConnector["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementExternalExadataStorageConnectorsDataSource().Schema["external_exadata_storage_connector_collection"].Elem.(*schema.Resource).Schema)
		externalExadataStorageConnector["items"] = items
	}

	resources = append(resources, externalExadataStorageConnector)
	if err := s.D.Set("external_exadata_storage_connector_collection", resources); err != nil {
		return err
	}

	return nil
}
