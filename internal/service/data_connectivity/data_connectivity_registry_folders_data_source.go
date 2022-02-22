// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_connectivity

import (
	"context"
	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_data_connectivity "github.com/oracle/oci-go-sdk/v58/dataconnectivity"
)

func DataConnectivityRegistryFoldersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataConnectivityRegistryFolders,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
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
			"folder_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     DataConnectivityRegistryFolderResource(),
						},
					},
				},
			},
		},
	}
}

func readDataConnectivityRegistryFolders(d *schema.ResourceData, m interface{}) error {
	sync := &DataConnectivityRegistryFoldersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataConnectivityManagementClient()

	return tfresource.ReadResource(sync)
}

type DataConnectivityRegistryFoldersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_connectivity.DataConnectivityManagementClient
	Res    *oci_data_connectivity.ListFoldersResponse
}

func (s *DataConnectivityRegistryFoldersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataConnectivityRegistryFoldersDataSourceCrud) Get() error {
	request := oci_data_connectivity.ListFoldersRequest{}

	if favoritesQueryParam, ok := s.D.GetOkExists("favorites_query_param"); ok {
		request.FavoritesQueryParam = oci_data_connectivity.ListFoldersFavoritesQueryParamEnum(favoritesQueryParam.(string))
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

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if registryId, ok := s.D.GetOkExists("registry_id"); ok {
		tmp := registryId.(string)
		request.RegistryId = &tmp
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		tmp := type_.(string)
		request.Type = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_connectivity")

	response, err := s.Client.ListFolders(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListFolders(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataConnectivityRegistryFoldersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataConnectivityRegistryFoldersDataSource-", DataConnectivityRegistryFoldersDataSource(), s.D))
	resources := []map[string]interface{}{}
	registryFolder := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, FolderSummaryToMap(item))
	}
	registryFolder["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataConnectivityRegistryFoldersDataSource().Schema["folder_summary_collection"].Elem.(*schema.Resource).Schema)
		registryFolder["items"] = items
	}

	resources = append(resources, registryFolder)
	if err := s.D.Set("folder_summary_collection", resources); err != nil {
		return err
	}

	return nil
}
