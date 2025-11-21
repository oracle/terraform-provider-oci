// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dbmulticloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_dbmulticloud "github.com/oracle/oci-go-sdk/v65/dbmulticloud"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DbmulticloudMultiCloudResourceDiscoveriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDbmulticloudMultiCloudResourceDiscoveries,
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
			"multi_cloud_resource_discovery_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oracle_db_azure_connector_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resources_filter": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"multi_cloud_resource_discovery_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DbmulticloudMultiCloudResourceDiscoveryResource()),
						},
					},
				},
			},
		},
	}
}

func readDbmulticloudMultiCloudResourceDiscoveries(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudMultiCloudResourceDiscoveriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MultiCloudResourceDiscoveryClient()

	return tfresource.ReadResource(sync)
}

type DbmulticloudMultiCloudResourceDiscoveriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dbmulticloud.MultiCloudResourceDiscoveryClient
	Res    *oci_dbmulticloud.ListMultiCloudResourceDiscoveriesResponse
}

func (s *DbmulticloudMultiCloudResourceDiscoveriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DbmulticloudMultiCloudResourceDiscoveriesDataSourceCrud) Get() error {
	request := oci_dbmulticloud.ListMultiCloudResourceDiscoveriesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if multiCloudResourceDiscoveryId, ok := s.D.GetOkExists("id"); ok {
		tmp := multiCloudResourceDiscoveryId.(string)
		request.MultiCloudResourceDiscoveryId = &tmp
	}

	if oracleDbAzureConnectorId, ok := s.D.GetOkExists("oracle_db_azure_connector_id"); ok {
		tmp := oracleDbAzureConnectorId.(string)
		request.OracleDbAzureConnectorId = &tmp
	}

	if resourceType, ok := s.D.GetOkExists("resource_type"); ok {
		request.ResourceType = oci_dbmulticloud.MultiCloudResourceDiscoveryResourceTypeEnum(resourceType.(string))
	}

	if resourcesFilter, ok := s.D.GetOkExists("resources_filter"); ok {
		arr := []string{resourcesFilter.(string)}
		//request.ResourcesFilter = tfresource.ObjectMapToStringMap(resourcesFilter.(map[string]interface{}))
		request.ResourcesFilter = arr
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_dbmulticloud.MultiCloudResourceDiscoveryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dbmulticloud")

	response, err := s.Client.ListMultiCloudResourceDiscoveries(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMultiCloudResourceDiscoveries(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DbmulticloudMultiCloudResourceDiscoveriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DbmulticloudMultiCloudResourceDiscoveriesDataSource-", DbmulticloudMultiCloudResourceDiscoveriesDataSource(), s.D))
	resources := []map[string]interface{}{}
	multiCloudResourceDiscovery := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MultiCloudResourceDiscoverySummaryToMap(item))
	}
	multiCloudResourceDiscovery["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DbmulticloudMultiCloudResourceDiscoveriesDataSource().Schema["multi_cloud_resource_discovery_summary_collection"].Elem.(*schema.Resource).Schema)
		multiCloudResourceDiscovery["items"] = items
	}

	resources = append(resources, multiCloudResourceDiscovery)
	if err := s.D.Set("multi_cloud_resource_discovery_summary_collection", resources); err != nil {
		return err
	}

	return nil
}
