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

func DatabaseManagementCloudDbNodesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementCloudDbNodes,
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
			"cloud_db_node_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DatabaseManagementCloudDbNodeResource()),
						},
					},
				},
			},
		},
	}
}

func readDatabaseManagementCloudDbNodes(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementCloudDbNodesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementCloudDbNodesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListCloudDbNodesResponse
}

func (s *DatabaseManagementCloudDbNodesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementCloudDbNodesDataSourceCrud) Get() error {
	request := oci_database_management.ListCloudDbNodesRequest{}

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

	response, err := s.Client.ListCloudDbNodes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCloudDbNodes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementCloudDbNodesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementCloudDbNodesDataSource-", DatabaseManagementCloudDbNodesDataSource(), s.D))
	resources := []map[string]interface{}{}
	cloudDbNode := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CloudDbNodeSummaryToMap(item))
	}
	cloudDbNode["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementCloudDbNodesDataSource().Schema["cloud_db_node_collection"].Elem.(*schema.Resource).Schema)
		cloudDbNode["items"] = items
	}

	resources = append(resources, cloudDbNode)
	if err := s.D.Set("cloud_db_node_collection", resources); err != nil {
		return err
	}

	return nil
}
