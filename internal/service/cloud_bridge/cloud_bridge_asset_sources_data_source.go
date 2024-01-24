// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_bridge

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_cloud_bridge "github.com/oracle/oci-go-sdk/v65/cloudbridge"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudBridgeAssetSourcesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCloudBridgeAssetSources,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"asset_source_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"asset_source_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(CloudBridgeAssetSourceResource()),
						},
					},
				},
			},
		},
	}
}

func readCloudBridgeAssetSources(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeAssetSourcesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DiscoveryClient()

	return tfresource.ReadResource(sync)
}

type CloudBridgeAssetSourcesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_bridge.DiscoveryClient
	Res    *oci_cloud_bridge.ListAssetSourcesResponse
}

func (s *CloudBridgeAssetSourcesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudBridgeAssetSourcesDataSourceCrud) Get() error {
	request := oci_cloud_bridge.ListAssetSourcesRequest{}

	if assetSourceId, ok := s.D.GetOkExists("id"); ok {
		tmp := assetSourceId.(string)
		request.AssetSourceId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_cloud_bridge.ListAssetSourcesLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_bridge")

	response, err := s.Client.ListAssetSources(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAssetSources(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CloudBridgeAssetSourcesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CloudBridgeAssetSourcesDataSource-", CloudBridgeAssetSourcesDataSource(), s.D))
	resources := []map[string]interface{}{}
	assetSource := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AssetSourceSummaryToMap(item))
	}
	assetSource["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CloudBridgeAssetSourcesDataSource().Schema["asset_source_collection"].Elem.(*schema.Resource).Schema)
		assetSource["items"] = items
	}

	resources = append(resources, assetSource)
	if err := s.D.Set("asset_source_collection", resources); err != nil {
		return err
	}

	return nil
}
