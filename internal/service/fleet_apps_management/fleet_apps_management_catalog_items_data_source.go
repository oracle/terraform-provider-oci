// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_apps_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetAppsManagementCatalogItemsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFleetAppsManagementCatalogItems,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"catalog_listing_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"catalog_listing_version_criteria": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"config_source_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"package_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"should_list_public_items": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"catalog_item_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(FleetAppsManagementCatalogItemResource()),
						},
					},
				},
			},
		},
	}
}

func readFleetAppsManagementCatalogItems(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementCatalogItemsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementCatalogClient()

	return tfresource.ReadResource(sync)
}

type FleetAppsManagementCatalogItemsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementCatalogClient
	Res    *oci_fleet_apps_management.ListCatalogItemsResponse
}

func (s *FleetAppsManagementCatalogItemsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementCatalogItemsDataSourceCrud) Get() error {
	request := oci_fleet_apps_management.ListCatalogItemsRequest{}

	if catalogListingId, ok := s.D.GetOkExists("catalog_listing_id"); ok {
		tmp := catalogListingId.(string)
		request.CatalogListingId = &tmp
	}

	if catalogListingVersionCriteria, ok := s.D.GetOkExists("catalog_listing_version_criteria"); ok {
		request.CatalogListingVersionCriteria = oci_fleet_apps_management.ListCatalogItemsCatalogListingVersionCriteriaEnum(catalogListingVersionCriteria.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if configSourceType, ok := s.D.GetOkExists("config_source_type"); ok {
		tmp := string(oci_fleet_apps_management.CatalogItemConfigSourceTypeEnum(configSourceType.(string)))
		request.ConfigSourceType = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if packageType, ok := s.D.GetOkExists("package_type"); ok {
		request.PackageType = oci_fleet_apps_management.CatalogItemPackageTypeEnum(packageType.(string))
	}

	if shouldListPublicItems, ok := s.D.GetOkExists("should_list_public_items"); ok {
		tmp := shouldListPublicItems.(bool)
		request.ShouldListPublicItems = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_fleet_apps_management.CatalogItemLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.ListCatalogItems(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCatalogItems(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FleetAppsManagementCatalogItemsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FleetAppsManagementCatalogItemsDataSource-", FleetAppsManagementCatalogItemsDataSource(), s.D))
	resources := []map[string]interface{}{}
	catalogItem := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CatalogItemSummaryToMap(item))
	}
	catalogItem["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FleetAppsManagementCatalogItemsDataSource().Schema["catalog_item_collection"].Elem.(*schema.Resource).Schema)
		catalogItem["items"] = items
	}

	resources = append(resources, catalogItem)
	if err := s.D.Set("catalog_item_collection", resources); err != nil {
		return err
	}

	return nil
}
