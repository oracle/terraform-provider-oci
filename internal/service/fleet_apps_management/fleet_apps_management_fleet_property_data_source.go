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

func FleetAppsManagementFleetPropertyDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["fleet_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["fleet_property_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(FleetAppsManagementFleetPropertyResource(), fieldMap, readSingularFleetAppsManagementFleetProperty)
}

func readSingularFleetAppsManagementFleetProperty(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementFleetPropertyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.ReadResource(sync)
}

type FleetAppsManagementFleetPropertyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementClient
	Res    *oci_fleet_apps_management.GetFleetPropertyResponse
}

func (s *FleetAppsManagementFleetPropertyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementFleetPropertyDataSourceCrud) Get() error {
	request := oci_fleet_apps_management.GetFleetPropertyRequest{}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	if fleetPropertyId, ok := s.D.GetOkExists("fleet_property_id"); ok {
		tmp := fleetPropertyId.(string)
		request.FleetPropertyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.GetFleetProperty(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FleetAppsManagementFleetPropertyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("allowed_values", s.Res.AllowedValues)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	} else {
		s.D.Set("compartment_id", nil)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.PropertyId != nil {
		s.D.Set("property_id", *s.Res.PropertyId)
	} else {
		s.D.Set("property_id", nil)
	}

	s.D.Set("selection_type", s.Res.SelectionType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.Value != nil {
		s.D.Set("value", *s.Res.Value)
	}

	s.D.Set("value_type", s.Res.ValueType)

	return nil
}
