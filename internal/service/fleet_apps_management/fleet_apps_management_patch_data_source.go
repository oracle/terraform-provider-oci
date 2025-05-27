// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_apps_management

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetAppsManagementPatchDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["patch_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(FleetAppsManagementPatchResource(), fieldMap, readSingularFleetAppsManagementPatch)
}

func readSingularFleetAppsManagementPatch(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementPatchDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementOperationsClient()

	return tfresource.ReadResource(sync)
}

type FleetAppsManagementPatchDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementOperationsClient
	Res    *oci_fleet_apps_management.GetPatchResponse
}

func (s *FleetAppsManagementPatchDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementPatchDataSourceCrud) Get() error {
	request := oci_fleet_apps_management.GetPatchRequest{}

	if patchId, ok := s.D.GetOkExists("patch_id"); ok {
		tmp := patchId.(string)
		request.PatchId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.GetPatch(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FleetAppsManagementPatchDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.ArtifactDetails != nil {
		artifactDetailsArray := []interface{}{}
		if artifactDetailsMap := ArtifactDetailsToMap(&s.Res.ArtifactDetails); artifactDetailsMap != nil {
			artifactDetailsArray = append(artifactDetailsArray, artifactDetailsMap)
		}
		s.D.Set("artifact_details", artifactDetailsArray)
	} else {
		s.D.Set("artifact_details", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	dependentPatches := []interface{}{}
	for _, item := range s.Res.DependentPatches {
		dependentPatches = append(dependentPatches, DependentPatchDetailsToMap(item))
	}
	s.D.Set("dependent_patches", dependentPatches)

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.PatchType != nil {
		s.D.Set("patch_type", []interface{}{PatchTypeToMap(s.Res.PatchType)})
	} else {
		s.D.Set("patch_type", nil)
	}

	if s.Res.Product != nil {
		s.D.Set("product", []interface{}{PatchProductToMap(s.Res.Product)})
	} else {
		s.D.Set("product", nil)
	}

	if s.Res.ResourceRegion != nil {
		s.D.Set("resource_region", *s.Res.ResourceRegion)
	}

	s.D.Set("severity", s.Res.Severity)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeReleased != nil {
		s.D.Set("time_released", s.Res.TimeReleased.Format(time.RFC3339Nano))
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("type", s.Res.Type)

	return nil
}
