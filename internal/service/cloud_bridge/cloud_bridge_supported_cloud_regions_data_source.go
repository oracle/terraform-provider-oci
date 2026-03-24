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

func CloudBridgeSupportedCloudRegionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCloudBridgeSupportedCloudRegions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"asset_source_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"supported_cloud_region_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"asset_source_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readCloudBridgeSupportedCloudRegions(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeSupportedCloudRegionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DiscoveryClient()

	return tfresource.ReadResource(sync)
}

type CloudBridgeSupportedCloudRegionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_bridge.DiscoveryClient
	Res    *oci_cloud_bridge.ListSupportedCloudRegionsResponse
}

func (s *CloudBridgeSupportedCloudRegionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudBridgeSupportedCloudRegionsDataSourceCrud) Get() error {
	request := oci_cloud_bridge.ListSupportedCloudRegionsRequest{}

	if assetSourceType, ok := s.D.GetOkExists("asset_source_type"); ok {
		request.AssetSourceType = oci_cloud_bridge.ListSupportedCloudRegionsAssetSourceTypeEnum(assetSourceType.(string))
	}

	if nameContains, ok := s.D.GetOkExists("name_contains"); ok {
		tmp := nameContains.(string)
		request.NameContains = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_bridge")

	response, err := s.Client.ListSupportedCloudRegions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSupportedCloudRegions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CloudBridgeSupportedCloudRegionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CloudBridgeSupportedCloudRegionsDataSource-", CloudBridgeSupportedCloudRegionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	supportedCloudRegion := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SupportedCloudRegionSummaryToMap(item))
	}
	supportedCloudRegion["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CloudBridgeSupportedCloudRegionsDataSource().Schema["supported_cloud_region_collection"].Elem.(*schema.Resource).Schema)
		supportedCloudRegion["items"] = items
	}

	resources = append(resources, supportedCloudRegion)
	if err := s.D.Set("supported_cloud_region_collection", resources); err != nil {
		return err
	}

	return nil
}

func SupportedCloudRegionSummaryToMap(obj oci_cloud_bridge.SupportedCloudRegionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["asset_source_type"] = string(obj.AssetSourceType)

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["state"] = string(obj.LifecycleState)

	return result
}
