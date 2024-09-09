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

func FleetAppsManagementFleetCredentialsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFleetAppsManagementFleetCredentials,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"credential_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fleet_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fleet_credential_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(FleetAppsManagementFleetCredentialResource()),
						},
					},
				},
			},
		},
	}
}

func readFleetAppsManagementFleetCredentials(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementFleetCredentialsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.ReadResource(sync)
}

type FleetAppsManagementFleetCredentialsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementClient
	Res    *oci_fleet_apps_management.ListFleetCredentialsResponse
}

func (s *FleetAppsManagementFleetCredentialsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementFleetCredentialsDataSourceCrud) Get() error {
	request := oci_fleet_apps_management.ListFleetCredentialsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if credentialLevel, ok := s.D.GetOkExists("credential_level"); ok {
		request.CredentialLevel = oci_fleet_apps_management.CredentialEntitySpecificDetailsCredentialLevelEnum(credentialLevel.(string))
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_fleet_apps_management.FleetCredentialLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.ListFleetCredentials(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListFleetCredentials(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FleetAppsManagementFleetCredentialsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FleetAppsManagementFleetCredentialsDataSource-", FleetAppsManagementFleetCredentialsDataSource(), s.D))
	resources := []map[string]interface{}{}
	fleetCredential := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, FleetCredentialSummaryToMap(item))
	}
	fleetCredential["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FleetAppsManagementFleetCredentialsDataSource().Schema["fleet_credential_collection"].Elem.(*schema.Resource).Schema)
		fleetCredential["items"] = items
	}

	resources = append(resources, fleetCredential)
	if err := s.D.Set("fleet_credential_collection", resources); err != nil {
		return err
	}

	return nil
}
