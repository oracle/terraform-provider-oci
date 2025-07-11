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

func DatabaseManagementCloudDbSystemConnectorsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementCloudDbSystemConnectors,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"cloud_db_system_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_db_system_connector_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DatabaseManagementCloudDbSystemConnectorResource()),
						},
					},
				},
			},
		},
	}
}

func readDatabaseManagementCloudDbSystemConnectors(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudDbSystemConnectorsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementCloudDbSystemConnectorsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListCloudDbSystemConnectorsResponse
}

func (s *DatabaseManagementCloudDbSystemConnectorsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementCloudDbSystemConnectorsDataSourceCrud) Get() error {
	request := oci_database_management.ListCloudDbSystemConnectorsRequest{}

	if cloudDbSystemId, ok := s.D.GetOkExists("cloud_db_system_id"); ok {
		tmp := cloudDbSystemId.(string)
		request.CloudDbSystemId = &tmp
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

	response, err := s.Client.ListCloudDbSystemConnectors(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCloudDbSystemConnectors(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementCloudDbSystemConnectorsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementCloudDbSystemConnectorsDataSource-", DatabaseManagementCloudDbSystemConnectorsDataSource(), s.D))
	resources := []map[string]interface{}{}
	cloudDbSystemConnector := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CloudDbSystemConnectorSummaryToMap(item))
	}
	cloudDbSystemConnector["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementCloudDbSystemConnectorsDataSource().Schema["cloud_db_system_connector_collection"].Elem.(*schema.Resource).Schema)
		cloudDbSystemConnector["items"] = items
	}

	resources = append(resources, cloudDbSystemConnector)
	if err := s.D.Set("cloud_db_system_connector_collection", resources); err != nil {
		return err
	}

	return nil
}
