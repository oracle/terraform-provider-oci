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

func CloudBridgeAssetsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCloudBridgeAssets,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"asset_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"asset_type": {
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
			"external_asset_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"inventory_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"asset_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(CloudBridgeAssetResource()),
						},
					},
				},
			},
		},
	}
}

func readCloudBridgeAssets(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeAssetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).InventoryClient()

	return tfresource.ReadResource(sync)
}

type CloudBridgeAssetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_bridge.InventoryClient
	Res    *oci_cloud_bridge.ListAssetsResponse
}

func (s *CloudBridgeAssetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudBridgeAssetsDataSourceCrud) Get() error {
	request := oci_cloud_bridge.ListAssetsRequest{}

	if assetId, ok := s.D.GetOkExists("id"); ok {
		tmp := assetId.(string)
		request.AssetId = &tmp
	}

	if assetType, ok := s.D.GetOkExists("asset_type"); ok {
		request.AssetType = oci_cloud_bridge.ListAssetsAssetTypeEnum(assetType.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if externalAssetKey, ok := s.D.GetOkExists("external_asset_key"); ok {
		tmp := externalAssetKey.(string)
		request.ExternalAssetKey = &tmp
	}

	if inventoryId, ok := s.D.GetOkExists("inventory_id"); ok {
		tmp := inventoryId.(string)
		request.InventoryId = &tmp
	}

	if sourceKey, ok := s.D.GetOkExists("source_key"); ok {
		tmp := sourceKey.(string)
		request.SourceKey = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_cloud_bridge.AssetLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_bridge")

	response, err := s.Client.ListAssets(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAssets(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CloudBridgeAssetsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CloudBridgeAssetsDataSource-", CloudBridgeAssetsDataSource(), s.D))
	resources := []map[string]interface{}{}
	asset := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AssetSummaryToMap(item))
	}
	asset["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CloudBridgeAssetsDataSource().Schema["asset_collection"].Elem.(*schema.Resource).Schema)
		asset["items"] = items
	}

	resources = append(resources, asset)
	if err := s.D.Set("asset_collection", resources); err != nil {
		return err
	}

	return nil
}
