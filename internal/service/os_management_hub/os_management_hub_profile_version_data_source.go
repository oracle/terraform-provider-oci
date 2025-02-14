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

func OsManagementHubProfileVersionDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularOsManagementHubProfileVersion,
		Schema: map[string]*schema.Schema{
			"profile_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"profile_version": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"arch_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_default_profile": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_service_provided_profile": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_environment": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"lifecycle_stage": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"managed_instance_group": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"management_station_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"os_family": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"profile_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"registration_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"software_sources": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_mandatory_for_autonomous_linux": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"software_source_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vendor_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularOsManagementHubProfileVersion(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubProfileVersionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OnboardingClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubProfileVersionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.OnboardingClient
	Res    *oci_os_management_hub.GetProfileVersionResponse
}

func (s *OsManagementHubProfileVersionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubProfileVersionDataSourceCrud) Get() error {
	request := oci_os_management_hub.GetProfileVersionRequest{}

	if profileId, ok := s.D.GetOkExists("profile_id"); ok {
		tmp := profileId.(string)
		request.ProfileId = &tmp
	}

	if profileVersion, ok := s.D.GetOkExists("profile_version"); ok {
		tmp := profileVersion.(string)
		request.ProfileVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.GetProfileVersion(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsManagementHubProfileVersionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("arch_type", s.Res.ArchType)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.IsDefaultProfile != nil {
		s.D.Set("is_default_profile", *s.Res.IsDefaultProfile)
	}

	if s.Res.IsServiceProvidedProfile != nil {
		s.D.Set("is_service_provided_profile", *s.Res.IsServiceProvidedProfile)
	}

	if s.Res.LifecycleEnvironment != nil {
		s.D.Set("lifecycle_environment", []interface{}{LifecycleEnvironmentDetailsToMap(s.Res.LifecycleEnvironment)})
	} else {
		s.D.Set("lifecycle_environment", nil)
	}

	if s.Res.LifecycleStage != nil {
		s.D.Set("lifecycle_stage", []interface{}{LifecycleStageDetailsToMap(s.Res.LifecycleStage)})
	} else {
		s.D.Set("lifecycle_stage", nil)
	}

	if s.Res.ManagedInstanceGroup != nil {
		s.D.Set("managed_instance_group", []interface{}{ManagedInstanceGroupDetailsToMap(s.Res.ManagedInstanceGroup)})
	} else {
		s.D.Set("managed_instance_group", nil)
	}

	if s.Res.ManagementStationId != nil {
		s.D.Set("management_station_id", *s.Res.ManagementStationId)
	}

	s.D.Set("os_family", s.Res.OsFamily)

	s.D.Set("profile_type", s.Res.ProfileType)

	s.D.Set("registration_type", s.Res.RegistrationType)

	softwareSources := []interface{}{}
	for _, item := range s.Res.SoftwareSources {
		softwareSources = append(softwareSources, SoftwareSourceDetailsToMap(item))
	}
	s.D.Set("software_sources", softwareSources)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeModified != nil {
		s.D.Set("time_modified", s.Res.TimeModified.String())
	}

	s.D.Set("vendor_name", s.Res.VendorName)

	return nil
}
