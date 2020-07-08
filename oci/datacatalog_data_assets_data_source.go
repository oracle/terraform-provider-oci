// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/common"
	oci_datacatalog "github.com/oracle/oci-go-sdk/datacatalog"
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
				Elem:     GetDataSourceItemSchema(DatacatalogDataAssetResource()),
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
	for _, item := range s.Res.Items {
		resources = append(resources, DataAssetSummaryToMap(item))
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, DatacatalogDataAssetsDataSource().Schema["data_asset_collection"].Elem.(*schema.Resource).Schema)
	}

	s.D.Set("data_asset_collection", resources)
	return nil
}
