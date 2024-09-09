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

func FleetAppsManagementPropertiesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFleetAppsManagementProperties,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
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
			"scope": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"property_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(FleetAppsManagementPropertyResource()),
						},
					},
				},
			},
		},
	}
}

func readFleetAppsManagementProperties(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementPropertiesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementAdminClient()

	return tfresource.ReadResource(sync)
}

type FleetAppsManagementPropertiesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementAdminClient
	Res    *oci_fleet_apps_management.ListPropertiesResponse
}

func (s *FleetAppsManagementPropertiesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementPropertiesDataSourceCrud) Get() error {
	request := oci_fleet_apps_management.ListPropertiesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_fleet_apps_management.ListPropertiesScopeEnum(scope.(string))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_fleet_apps_management.PropertyLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.ListProperties(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListProperties(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FleetAppsManagementPropertiesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FleetAppsManagementPropertiesDataSource-", FleetAppsManagementPropertiesDataSource(), s.D))
	resources := []map[string]interface{}{}
	property := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, PropertySummaryToMap(item))
	}
	property["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FleetAppsManagementPropertiesDataSource().Schema["property_collection"].Elem.(*schema.Resource).Schema)
		property["items"] = items
	}

	resources = append(resources, property)
	if err := s.D.Set("property_collection", resources); err != nil {
		return err
	}

	return nil
}
