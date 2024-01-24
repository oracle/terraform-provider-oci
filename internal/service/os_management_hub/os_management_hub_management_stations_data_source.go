// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubManagementStationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsManagementHubManagementStations,
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
			"display_name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"managed_instance_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"management_station_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(OsManagementHubManagementStationResource()),
						},
					},
				},
			},
		},
	}
}

func readOsManagementHubManagementStations(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubManagementStationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementStationClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubManagementStationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.ManagementStationClient
	Res    *oci_os_management_hub.ListManagementStationsResponse
}

func (s *OsManagementHubManagementStationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubManagementStationsDataSourceCrud) Get() error {
	request := oci_os_management_hub.ListManagementStationsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if displayNameContains, ok := s.D.GetOkExists("display_name_contains"); ok {
		tmp := displayNameContains.(string)
		request.DisplayNameContains = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_os_management_hub.ManagementStationLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.ListManagementStations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListManagementStations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsManagementHubManagementStationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsManagementHubManagementStationsDataSource-", OsManagementHubManagementStationsDataSource(), s.D))
	resources := []map[string]interface{}{}
	managementStation := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ManagementStationSummaryToMap(item))
	}
	managementStation["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OsManagementHubManagementStationsDataSource().Schema["management_station_collection"].Elem.(*schema.Resource).Schema)
		managementStation["items"] = items
	}

	resources = append(resources, managementStation)
	if err := s.D.Set("management_station_collection", resources); err != nil {
		return err
	}

	return nil
}
