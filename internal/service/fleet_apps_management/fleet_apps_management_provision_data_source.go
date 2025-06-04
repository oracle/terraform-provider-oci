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

func FleetAppsManagementProvisionDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["provision_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(FleetAppsManagementProvisionResource(), fieldMap, readSingularFleetAppsManagementProvision)
}

func readSingularFleetAppsManagementProvision(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementProvisionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementProvisionClient()

	return tfresource.ReadResource(sync)
}

type FleetAppsManagementProvisionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementProvisionClient
	Res    *oci_fleet_apps_management.GetProvisionResponse
}

func (s *FleetAppsManagementProvisionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementProvisionDataSourceCrud) Get() error {
	request := oci_fleet_apps_management.GetProvisionRequest{}

	if provisionId, ok := s.D.GetOkExists("provision_id"); ok {
		tmp := provisionId.(string)
		request.ProvisionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.GetProvision(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FleetAppsManagementProvisionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConfigCatalogItemDisplayName != nil {
		s.D.Set("config_catalog_item_display_name", *s.Res.ConfigCatalogItemDisplayName)
	}

	if s.Res.ConfigCatalogItemId != nil {
		s.D.Set("config_catalog_item_id", *s.Res.ConfigCatalogItemId)
	}

	if s.Res.ConfigCatalogItemListingId != nil {
		s.D.Set("config_catalog_item_listing_id", *s.Res.ConfigCatalogItemListingId)
	}

	if s.Res.ConfigCatalogItemListingVersion != nil {
		s.D.Set("config_catalog_item_listing_version", *s.Res.ConfigCatalogItemListingVersion)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	deployedResources := []interface{}{}
	for _, item := range s.Res.DeployedResources {
		deployedResources = append(deployedResources, DeployedResourceDetailsToMap(item))
	}
	s.D.Set("deployed_resources", deployedResources)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.FleetId != nil {
		s.D.Set("fleet_id", *s.Res.FleetId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.PackageCatalogItemDisplayName != nil {
		s.D.Set("package_catalog_item_display_name", *s.Res.PackageCatalogItemDisplayName)
	}

	if s.Res.PackageCatalogItemId != nil {
		s.D.Set("package_catalog_item_id", *s.Res.PackageCatalogItemId)
	}

	if s.Res.PackageCatalogItemListingId != nil {
		s.D.Set("package_catalog_item_listing_id", *s.Res.PackageCatalogItemListingId)
	}

	if s.Res.PackageCatalogItemListingVersion != nil {
		s.D.Set("package_catalog_item_listing_version", *s.Res.PackageCatalogItemListingVersion)
	}

	if s.Res.ProvisionDescription != nil {
		s.D.Set("provision_description", *s.Res.ProvisionDescription)
	}

	if s.Res.RmsApplyJobId != nil {
		s.D.Set("rms_apply_job_id", *s.Res.RmsApplyJobId)
	}

	if s.Res.StackId != nil {
		s.D.Set("stack_id", *s.Res.StackId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	tfOutputs := []interface{}{}
	for _, item := range s.Res.TfOutputs {
		tfOutputs = append(tfOutputs, JobExecutionDetailsToMap(item))
	}
	s.D.Set("tf_outputs", tfOutputs)

	if s.Res.TfVariableCompartmentId != nil {
		s.D.Set("tf_variable_compartment_id", *s.Res.TfVariableCompartmentId)
	}

	if s.Res.TfVariableCurrentUserId != nil {
		s.D.Set("tf_variable_current_user_id", *s.Res.TfVariableCurrentUserId)
	}

	if s.Res.TfVariableRegionId != nil {
		s.D.Set("tf_variable_region_id", *s.Res.TfVariableRegionId)
	}

	if s.Res.TfVariableTenancyId != nil {
		s.D.Set("tf_variable_tenancy_id", *s.Res.TfVariableTenancyId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
