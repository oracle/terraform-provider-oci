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

func DatabaseManagementCloudExadataStorageConnectorsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementCloudExadataStorageConnectors,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"cloud_exadata_infrastructure_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_exadata_storage_connector_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DatabaseManagementCloudExadataStorageConnectorResource()),
						},
					},
				},
			},
		},
	}
}

func readDatabaseManagementCloudExadataStorageConnectors(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudExadataStorageConnectorsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementCloudExadataStorageConnectorsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListCloudExadataStorageConnectorsResponse
}

func (s *DatabaseManagementCloudExadataStorageConnectorsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementCloudExadataStorageConnectorsDataSourceCrud) Get() error {
	request := oci_database_management.ListCloudExadataStorageConnectorsRequest{}

	if cloudExadataInfrastructureId, ok := s.D.GetOkExists("cloud_exadata_infrastructure_id"); ok {
		tmp := cloudExadataInfrastructureId.(string)
		request.CloudExadataInfrastructureId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListCloudExadataStorageConnectors(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCloudExadataStorageConnectors(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementCloudExadataStorageConnectorsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementCloudExadataStorageConnectorsDataSource-", DatabaseManagementCloudExadataStorageConnectorsDataSource(), s.D))
	resources := []map[string]interface{}{}
	cloudExadataStorageConnector := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CloudExadataStorageConnectorSummaryToMap(item))
	}
	cloudExadataStorageConnector["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementCloudExadataStorageConnectorsDataSource().Schema["cloud_exadata_storage_connector_collection"].Elem.(*schema.Resource).Schema)
		cloudExadataStorageConnector["items"] = items
	}

	resources = append(resources, cloudExadataStorageConnector)
	if err := s.D.Set("cloud_exadata_storage_connector_collection", resources); err != nil {
		return err
	}

	return nil
}
