// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubProfileDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["profile_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OsManagementHubProfileResource(), fieldMap, readSingularOsManagementHubProfile)
}

func readSingularOsManagementHubProfile(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubProfileDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OnboardingClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubProfileDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.OnboardingClient
	Res    *oci_os_management_hub.GetProfileResponse
}

func (s *OsManagementHubProfileDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubProfileDataSourceCrud) Get() error {
	request := oci_os_management_hub.GetProfileRequest{}

	if profileId, ok := s.D.GetOkExists("profile_id"); ok {
		tmp := profileId.(string)
		request.ProfileId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.GetProfile(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsManagementHubProfileDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.GetId())
	switch v := (s.Res.Profile).(type) {
	case oci_os_management_hub.GroupProfile:
		s.D.Set("profile_type", "GROUP")

		if v.ManagedInstanceGroup != nil {
			s.D.Set("managed_instance_group", []interface{}{ManagedInstanceGroupDetailsToMap(v.ManagedInstanceGroup)})
		} else {
			s.D.Set("managed_instance_group", nil)
		}

		s.D.Set("arch_type", v.ArchType)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.ManagementStationId != nil {
			s.D.Set("management_station_id", *v.ManagementStationId)
		}

		s.D.Set("os_family", v.OsFamily)

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		s.D.Set("vendor_name", v.VendorName)
	case oci_os_management_hub.LifecycleProfile:
		s.D.Set("profile_type", "LIFECYCLE")

		if v.LifecycleEnvironment != nil {
			s.D.Set("lifecycle_environment", []interface{}{LifecycleEnvironmentDetailsToMap(v.LifecycleEnvironment)})
		} else {
			s.D.Set("lifecycle_environment", nil)
		}

		if v.LifecycleStage != nil {
			s.D.Set("lifecycle_stage", []interface{}{LifecycleStageDetailsToMap(v.LifecycleStage)})
		} else {
			s.D.Set("lifecycle_stage", nil)
		}

		s.D.Set("arch_type", v.ArchType)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.ManagementStationId != nil {
			s.D.Set("management_station_id", *v.ManagementStationId)
		}

		s.D.Set("os_family", v.OsFamily)

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		s.D.Set("vendor_name", v.VendorName)
	case oci_os_management_hub.SoftwareSourceProfile:
		s.D.Set("profile_type", "SOFTWARESOURCE")

		softwareSources := []interface{}{}
		for _, item := range v.SoftwareSources {
			softwareSources = append(softwareSources, SoftwareSourceDetailsToMap(item))
		}
		s.D.Set("software_sources", softwareSources)

		s.D.Set("arch_type", v.ArchType)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.ManagementStationId != nil {
			s.D.Set("management_station_id", *v.ManagementStationId)
		}

		s.D.Set("os_family", v.OsFamily)

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		s.D.Set("vendor_name", v.VendorName)
	case oci_os_management_hub.StationProfile:
		s.D.Set("profile_type", "STATION")

		s.D.Set("arch_type", v.ArchType)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.ManagementStationId != nil {
			s.D.Set("management_station_id", *v.ManagementStationId)
		}

		s.D.Set("os_family", v.OsFamily)

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		s.D.Set("vendor_name", v.VendorName)
	default:
		log.Printf("[WARN] Received 'profile_type' of unknown type %v", s.Res.Profile)
		return nil
	}

	return nil
}
