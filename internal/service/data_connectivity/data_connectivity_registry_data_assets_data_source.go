// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_connectivity

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_connectivity "github.com/oracle/oci-go-sdk/v64/dataconnectivity"
)

func DataConnectivityRegistryDataAssetsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataConnectivityRegistryDataAssets,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"endpoint_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"exclude_endpoint_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"exclude_types": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"favorites_query_param": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fields": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"folder_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"include_types": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"registry_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_asset_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     DataConnectivityRegistryDataAssetResource(),
						},
					},
				},
			},
		},
	}
}

func readDataConnectivityRegistryDataAssets(d *schema.ResourceData, m interface{}) error {
	sync := &DataConnectivityRegistryDataAssetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataConnectivityManagementClient()

	return tfresource.ReadResource(sync)
}

type DataConnectivityRegistryDataAssetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_connectivity.DataConnectivityManagementClient
	Res    *oci_data_connectivity.ListDataAssetsResponse
}

func (s *DataConnectivityRegistryDataAssetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataConnectivityRegistryDataAssetsDataSourceCrud) Get() error {
	request := oci_data_connectivity.ListDataAssetsRequest{}

	if endpointIds, ok := s.D.GetOkExists("endpoint_ids"); ok {
		interfaces := endpointIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("endpoint_ids") {
			request.EndpointIds = tmp
		}
	}

	if excludeEndpointIds, ok := s.D.GetOkExists("exclude_endpoint_ids"); ok {
		interfaces := excludeEndpointIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("exclude_endpoint_ids") {
			request.ExcludeEndpointIds = tmp
		}
	}

	if excludeTypes, ok := s.D.GetOkExists("exclude_types"); ok {
		interfaces := excludeTypes.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("exclude_types") {
			request.ExcludeTypes = tmp
		}
	}

	if favoritesQueryParam, ok := s.D.GetOkExists("favorites_query_param"); ok {
		request.FavoritesQueryParam = oci_data_connectivity.ListDataAssetsFavoritesQueryParamEnum(favoritesQueryParam.(string))
	}

	if fields, ok := s.D.GetOkExists("fields"); ok {
		interfaces := fields.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("fields") {
			request.Fields = tmp
		}
	}

	if folderId, ok := s.D.GetOkExists("folder_id"); ok {
		tmp := folderId.(string)
		request.FolderId = &tmp
	}

	if includeTypes, ok := s.D.GetOkExists("include_types"); ok {
		interfaces := includeTypes.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("include_types") {
			request.IncludeTypes = tmp
		}
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if registryId, ok := s.D.GetOkExists("registry_id"); ok {
		tmp := registryId.(string)
		request.RegistryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_connectivity")

	response, err := s.Client.ListDataAssets(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDataAssets(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataConnectivityRegistryDataAssetsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataConnectivityRegistryDataAssetsDataSource-", DataConnectivityRegistryDataAssetsDataSource(), s.D))
	resources := []map[string]interface{}{}
	registryDataAsset := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DataConnectivityDataAssetSummaryToMap(item))
	}
	registryDataAsset["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataConnectivityRegistryDataAssetsDataSource().Schema["data_asset_summary_collection"].Elem.(*schema.Resource).Schema)
		registryDataAsset["items"] = items
	}

	resources = append(resources, registryDataAsset)
	if err := s.D.Set("data_asset_summary_collection", resources); err != nil {
		return err
	}

	return nil
}
