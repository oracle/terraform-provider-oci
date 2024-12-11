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

func FleetAppsManagementPlatformConfigurationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFleetAppsManagementPlatformConfigurations,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"config_category": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"platform_configuration_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(FleetAppsManagementPlatformConfigurationResource()),
						},
					},
				},
			},
		},
	}
}

func readFleetAppsManagementPlatformConfigurations(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementPlatformConfigurationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementAdminClient()

	return tfresource.ReadResource(sync)
}

type FleetAppsManagementPlatformConfigurationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementAdminClient
	Res    *oci_fleet_apps_management.ListPlatformConfigurationsResponse
}

func (s *FleetAppsManagementPlatformConfigurationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementPlatformConfigurationsDataSourceCrud) Get() error {
	request := oci_fleet_apps_management.ListPlatformConfigurationsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if configCategory, ok := s.D.GetOkExists("config_category"); ok {
		request.ConfigCategory = oci_fleet_apps_management.ConfigCategoryDetailsConfigCategoryEnum(configCategory.(string))
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_fleet_apps_management.PlatformConfigurationLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.ListPlatformConfigurations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPlatformConfigurations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FleetAppsManagementPlatformConfigurationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FleetAppsManagementPlatformConfigurationsDataSource-", FleetAppsManagementPlatformConfigurationsDataSource(), s.D))
	resources := []map[string]interface{}{}
	platformConfiguration := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, PlatformConfigurationSummaryToMap(item))
	}
	platformConfiguration["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FleetAppsManagementPlatformConfigurationsDataSource().Schema["platform_configuration_collection"].Elem.(*schema.Resource).Schema)
		platformConfiguration["items"] = items
	}

	resources = append(resources, platformConfiguration)
	if err := s.D.Set("platform_configuration_collection", resources); err != nil {
		return err
	}

	return nil
}
