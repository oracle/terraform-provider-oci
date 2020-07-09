// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v25/common"
	oci_datacatalog "github.com/oracle/oci-go-sdk/v25/datacatalog"
)

func init() {
	RegisterDatasource("oci_datacatalog_data_assets", DatacatalogDataAssetsDataSource())
}

func DatacatalogDataAssetsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatacatalogDataAssets,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"catalog_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"created_by_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fields": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      literalTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_asset_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"count": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     DatacatalogDataAssetResource(),
						},
					},
				},
			},
		},
	}
}

func readDatacatalogDataAssets(d *schema.ResourceData, m interface{}) error {
	sync := &DatacatalogDataAssetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataCatalogClient()

	return ReadResource(sync)
}

type DatacatalogDataAssetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datacatalog.DataCatalogClient
	Res    *oci_datacatalog.ListDataAssetsResponse
}

func (s *DatacatalogDataAssetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatacatalogDataAssetsDataSourceCrud) Get() error {
	request := oci_datacatalog.ListDataAssetsRequest{}

	if catalogId, ok := s.D.GetOkExists("catalog_id"); ok {
		tmp := catalogId.(string)
		request.CatalogId = &tmp
	}

	if createdById, ok := s.D.GetOkExists("created_by_id"); ok {
		tmp := createdById.(string)
		request.CreatedById = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if displayNameContains, ok := s.D.GetOkExists("display_name_contains"); ok {
		tmp := displayNameContains.(string)
		request.DisplayNameContains = &tmp
	}

	if externalKey, ok := s.D.GetOkExists("external_key"); ok {
		tmp := externalKey.(string)
		request.ExternalKey = &tmp
	}

	if fields, ok := s.D.GetOkExists("fields"); ok {
		set := fields.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_datacatalog.ListDataAssetsFieldsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_datacatalog.ListDataAssetsFieldsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("fields") {
			request.Fields = tmp
		}
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_datacatalog.ListDataAssetsLifecycleStateEnum(state.(string))
	}

	if timeCreated, ok := s.D.GetOkExists("time_created"); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreated.(string))
		if err != nil {
			return err
		}
		request.TimeCreated = &oci_common.SDKTime{Time: tmp}
	}

	if timeUpdated, ok := s.D.GetOkExists("time_updated"); ok {
		tmp, err := time.Parse(time.RFC3339, timeUpdated.(string))
		if err != nil {
			return err
		}
		request.TimeUpdated = &oci_common.SDKTime{Time: tmp}
	}

	if typeKey, ok := s.D.GetOkExists("type_key"); ok {
		tmp := typeKey.(string)
		request.TypeKey = &tmp
	}

	if updatedById, ok := s.D.GetOkExists("updated_by_id"); ok {
		tmp := updatedById.(string)
		request.UpdatedById = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "datacatalog")

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

func (s *DatacatalogDataAssetsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}
	dataAsset := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DataAssetSummaryToMap(item))
	}
	dataAsset["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, DatacatalogDataAssetsDataSource().Schema["data_asset_collection"].Elem.(*schema.Resource).Schema)
		dataAsset["items"] = items
	}

	dataAsset["count"] = *s.Res.Count

	resources = append(resources, dataAsset)
	if err := s.D.Set("data_asset_collection", resources); err != nil {
		return err
	}

	return nil
}
