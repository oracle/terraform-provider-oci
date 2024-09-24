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

func FleetAppsManagementFleetResourcesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFleetAppsManagementFleetResources,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fleet_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"fleet_resource_type": {
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
			"tenancy_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fleet_resource_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(FleetAppsManagementFleetResourceResource()),
						},
					},
				},
			},
		},
	}
}

func readFleetAppsManagementFleetResources(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementFleetResourcesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.ReadResource(sync)
}

type FleetAppsManagementFleetResourcesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementClient
	Res    *oci_fleet_apps_management.ListFleetResourcesResponse
}

func (s *FleetAppsManagementFleetResourcesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementFleetResourcesDataSourceCrud) Get() error {
	request := oci_fleet_apps_management.ListFleetResourcesRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	if fleetResourceType, ok := s.D.GetOkExists("fleet_resource_type"); ok {
		tmp := fleetResourceType.(string)
		request.FleetResourceType = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_fleet_apps_management.FleetResourceLifecycleStateEnum(state.(string))
	}

	if tenancyId, ok := s.D.GetOkExists("tenancy_id"); ok {
		tmp := tenancyId.(string)
		request.TenancyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.ListFleetResources(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListFleetResources(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FleetAppsManagementFleetResourcesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FleetAppsManagementFleetResourcesDataSource-", FleetAppsManagementFleetResourcesDataSource(), s.D))
	resources := []map[string]interface{}{}
	fleetResource := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, FleetResourceSummaryToMap(item))
	}
	fleetResource["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FleetAppsManagementFleetResourcesDataSource().Schema["fleet_resource_collection"].Elem.(*schema.Resource).Schema)
		fleetResource["items"] = items
	}

	resources = append(resources, fleetResource)
	if err := s.D.Set("fleet_resource_collection", resources); err != nil {
		return err
	}

	return nil
}
