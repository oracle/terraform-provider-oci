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

func FleetAppsManagementPlatformConfigurationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["platform_configuration_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(FleetAppsManagementPlatformConfigurationResource(), fieldMap, readSingularFleetAppsManagementPlatformConfiguration)
}

func readSingularFleetAppsManagementPlatformConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementPlatformConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementAdminClient()

	return tfresource.ReadResource(sync)
}

type FleetAppsManagementPlatformConfigurationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementAdminClient
	Res    *oci_fleet_apps_management.GetPlatformConfigurationResponse
}

func (s *FleetAppsManagementPlatformConfigurationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementPlatformConfigurationDataSourceCrud) Get() error {
	request := oci_fleet_apps_management.GetPlatformConfigurationRequest{}

	if platformConfigurationId, ok := s.D.GetOkExists("platform_configuration_id"); ok {
		tmp := platformConfigurationId.(string)
		request.PlatformConfigurationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.GetPlatformConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FleetAppsManagementPlatformConfigurationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConfigCategoryDetails != nil {
		configCategoryDetailsArray := []interface{}{}
		if configCategoryDetailsMap := ConfigCategoryDetailsToMap(&s.Res.ConfigCategoryDetails); configCategoryDetailsMap != nil {
			configCategoryDetailsArray = append(configCategoryDetailsArray, configCategoryDetailsMap)
		}
		s.D.Set("config_category_details", configCategoryDetailsArray)
	} else {
		s.D.Set("config_category_details", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ResourceRegion != nil {
		s.D.Set("resource_region", *s.Res.ResourceRegion)
	}

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

	s.D.Set("type", s.Res.Type)

	return nil
}
