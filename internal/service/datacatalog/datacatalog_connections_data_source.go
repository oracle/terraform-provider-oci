// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datacatalog

import (
	"context"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_datacatalog "github.com/oracle/oci-go-sdk/v58/datacatalog"
)

func DatacatalogConnectionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatacatalogConnections,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"catalog_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"created_by_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_asset_key": {
				Type:     schema.TypeString,
				Required: true,
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
				Set:      utils.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"is_default": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_status_updated": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"updated_by_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"connection_collection": {
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
							Elem:     DatacatalogConnectionResource(),
						},
					},
				},
			},
		},
	}
}

func readDatacatalogConnections(d *schema.ResourceData, m interface{}) error {
	sync := &DatacatalogConnectionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataCatalogClient()

	return tfresource.ReadResource(sync)
}

type DatacatalogConnectionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datacatalog.DataCatalogClient
	Res    *oci_datacatalog.ListConnectionsResponse
}

func (s *DatacatalogConnectionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatacatalogConnectionsDataSourceCrud) Get() error {
	request := oci_datacatalog.ListConnectionsRequest{}

	if catalogId, ok := s.D.GetOkExists("catalog_id"); ok {
		tmp := catalogId.(string)
		request.CatalogId = &tmp
	}

	if createdById, ok := s.D.GetOkExists("created_by_id"); ok {
		tmp := createdById.(string)
		request.CreatedById = &tmp
	}

	if dataAssetKey, ok := s.D.GetOkExists("data_asset_key"); ok {
		tmp := dataAssetKey.(string)
		request.DataAssetKey = &tmp
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
		tmp := make([]oci_datacatalog.ListConnectionsFieldsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_datacatalog.ListConnectionsFieldsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("fields") {
			request.Fields = tmp
		}
	}

	if isDefault, ok := s.D.GetOkExists("is_default"); ok {
		tmp := isDefault.(bool)
		request.IsDefault = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_datacatalog.ListConnectionsLifecycleStateEnum(state.(string))
	}

	if timeCreated, ok := s.D.GetOkExists("time_created"); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreated.(string))
		if err != nil {
			return err
		}
		request.TimeCreated = &oci_common.SDKTime{Time: tmp}
	}

	if timeStatusUpdated, ok := s.D.GetOkExists("time_status_updated"); ok {
		tmp, err := time.Parse(time.RFC3339, timeStatusUpdated.(string))
		if err != nil {
			return err
		}
		request.TimeStatusUpdated = &oci_common.SDKTime{Time: tmp}
	}

	if timeUpdated, ok := s.D.GetOkExists("time_updated"); ok {
		tmp, err := time.Parse(time.RFC3339, timeUpdated.(string))
		if err != nil {
			return err
		}
		request.TimeUpdated = &oci_common.SDKTime{Time: tmp}
	}

	if updatedById, ok := s.D.GetOkExists("updated_by_id"); ok {
		tmp := updatedById.(string)
		request.UpdatedById = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datacatalog")

	response, err := s.Client.ListConnections(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListConnections(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatacatalogConnectionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatacatalogConnectionsDataSource-", DatacatalogConnectionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	connection := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ConnectionSummaryToMap(item))
	}
	connection["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatacatalogConnectionsDataSource().Schema["connection_collection"].Elem.(*schema.Resource).Schema)
		connection["items"] = items
	}

	connection["count"] = *s.Res.Count

	resources = append(resources, connection)
	if err := s.D.Set("connection_collection", resources); err != nil {
		return err
	}

	return nil
}
