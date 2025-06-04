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

func FleetAppsManagementFleetDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["fleet_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(FleetAppsManagementFleetResource(), fieldMap, readSingularFleetAppsManagementFleet)
}

func readSingularFleetAppsManagementFleet(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementFleetDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.ReadResource(sync)
}

type FleetAppsManagementFleetDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementClient
	Res    *oci_fleet_apps_management.GetFleetResponse
}

func (s *FleetAppsManagementFleetDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementFleetDataSourceCrud) Get() error {
	request := oci_fleet_apps_management.GetFleetRequest{}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.GetFleet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FleetAppsManagementFleetDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	} else {
		s.D.Set("compartment_id", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.Details != nil {
		detailsArray := []interface{}{}
		if detailsMap := FleetDetailsToMap(&s.Res.Details); detailsMap != nil {
			detailsArray = append(detailsArray, detailsMap)
		}
		s.D.Set("details", detailsArray)
	} else {
		s.D.Set("details", nil)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.EnvironmentType != nil {
		s.D.Set("environment_type", *s.Res.EnvironmentType)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	if s.Res.IsTargetAutoConfirm != nil {
		s.D.Set("is_target_auto_confirm", *s.Res.IsTargetAutoConfirm)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	notificationPreferences := []interface{}{}
	for _, item := range s.Res.NotificationPreferences {
		notificationPreferences = append(notificationPreferences, NotificationPreferenceToMap(item))
	}
	s.D.Set("notification_preferences", notificationPreferences)

	if s.Res.ParentFleetId != nil {
		s.D.Set("parent_fleet_id", *s.Res.ParentFleetId)
	}

	s.D.Set("products", s.Res.Products)

	if s.Res.ResourceRegion != nil {
		s.D.Set("resource_region", *s.Res.ResourceRegion)
	}

	if s.Res.ResourceSelection != nil {
		resourceSelectionArray := []interface{}{}
		if resourceSelectionMap := ResourceSelectionToMap(&s.Res.ResourceSelection); resourceSelectionMap != nil {
			resourceSelectionArray = append(resourceSelectionArray, resourceSelectionMap)
		}
		s.D.Set("resource_selection", resourceSelectionArray)
	} else {
		s.D.Set("resource_selection", nil)
	}

	resources := []interface{}{}
	for _, item := range s.Res.Resources {
		resources = append(resources, AssociatedFleetResourceDetailsToMap(item))
	}
	s.D.Set("resources", resources)

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

	return nil
}
