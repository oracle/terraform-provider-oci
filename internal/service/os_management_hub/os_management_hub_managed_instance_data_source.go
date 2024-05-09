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

func OsManagementHubManagedInstanceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["managed_instance_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OsManagementHubManagedInstanceResource(), fieldMap, readSingularOsManagementHubManagedInstance)
}

func readSingularOsManagementHubManagedInstance(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubManagedInstanceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedInstanceClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubManagedInstanceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.ManagedInstanceClient
	Res    *oci_os_management_hub.GetManagedInstanceResponse
}

func (s *OsManagementHubManagedInstanceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubManagedInstanceDataSourceCrud) Get() error {
	request := oci_os_management_hub.GetManagedInstanceRequest{}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.GetManagedInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsManagementHubManagedInstanceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("architecture", s.Res.Architecture)

	if s.Res.AutonomousSettings != nil {
		s.D.Set("autonomous_settings", []interface{}{AutonomousSettingsToMap(s.Res.AutonomousSettings)})
	} else {
		s.D.Set("autonomous_settings", nil)
	}

	if s.Res.BugUpdatesAvailable != nil {
		s.D.Set("bug_updates_available", *s.Res.BugUpdatesAvailable)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.EnhancementUpdatesAvailable != nil {
		s.D.Set("enhancement_updates_available", *s.Res.EnhancementUpdatesAvailable)
	}

	if s.Res.InstalledPackages != nil {
		s.D.Set("installed_packages", *s.Res.InstalledPackages)
	}

	if s.Res.InstalledWindowsUpdates != nil {
		s.D.Set("installed_windows_updates", *s.Res.InstalledWindowsUpdates)
	}

	if s.Res.IsManagedByAutonomousLinux != nil {
		s.D.Set("is_managed_by_autonomous_linux", *s.Res.IsManagedByAutonomousLinux)
	}

	if s.Res.IsManagementStation != nil {
		s.D.Set("is_management_station", *s.Res.IsManagementStation)
	}

	if s.Res.IsRebootRequired != nil {
		s.D.Set("is_reboot_required", *s.Res.IsRebootRequired)
	}

	if s.Res.KspliceEffectiveKernelVersion != nil {
		s.D.Set("ksplice_effective_kernel_version", *s.Res.KspliceEffectiveKernelVersion)
	}

	if s.Res.LifecycleEnvironment != nil {
		s.D.Set("lifecycle_environment", []interface{}{IdToMap(s.Res.LifecycleEnvironment)})
	} else {
		s.D.Set("lifecycle_environment", nil)
	}

	if s.Res.LifecycleStage != nil {
		s.D.Set("lifecycle_stage", []interface{}{IdToMap(s.Res.LifecycleStage)})
	} else {
		s.D.Set("lifecycle_stage", nil)
	}

	s.D.Set("location", s.Res.Location)

	if s.Res.ManagedInstanceGroup != nil {
		s.D.Set("managed_instance_group", []interface{}{IdToMap(s.Res.ManagedInstanceGroup)})
	} else {
		s.D.Set("managed_instance_group", nil)
	}

	if s.Res.NotificationTopicId != nil {
		s.D.Set("notification_topic_id", *s.Res.NotificationTopicId)
	}

	s.D.Set("os_family", s.Res.OsFamily)

	if s.Res.OsKernelVersion != nil {
		s.D.Set("os_kernel_version", *s.Res.OsKernelVersion)
	}

	if s.Res.OsName != nil {
		s.D.Set("os_name", *s.Res.OsName)
	}

	if s.Res.OsVersion != nil {
		s.D.Set("os_version", *s.Res.OsVersion)
	}

	if s.Res.OtherUpdatesAvailable != nil {
		s.D.Set("other_updates_available", *s.Res.OtherUpdatesAvailable)
	}

	if s.Res.PrimaryManagementStationId != nil {
		s.D.Set("primary_management_station_id", *s.Res.PrimaryManagementStationId)
	}

	if s.Res.Profile != nil {
		s.D.Set("profile", *s.Res.Profile)
	}

	if s.Res.ScheduledJobCount != nil {
		s.D.Set("scheduled_job_count", *s.Res.ScheduledJobCount)
	}

	if s.Res.SecondaryManagementStationId != nil {
		s.D.Set("secondary_management_station_id", *s.Res.SecondaryManagementStationId)
	}

	if s.Res.SecurityUpdatesAvailable != nil {
		s.D.Set("security_updates_available", *s.Res.SecurityUpdatesAvailable)
	}

	softwareSources := []interface{}{}
	for _, item := range s.Res.SoftwareSources {
		softwareSources = append(softwareSources, SoftwareSourceDetailsToMap(item))
	}
	s.D.Set("software_sources", softwareSources)

	s.D.Set("status", s.Res.Status)

	if s.Res.TenancyId != nil {
		s.D.Set("tenancy_id", *s.Res.TenancyId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastBoot != nil {
		s.D.Set("time_last_boot", s.Res.TimeLastBoot.String())
	}

	if s.Res.TimeLastCheckin != nil {
		s.D.Set("time_last_checkin", s.Res.TimeLastCheckin.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.UpdatesAvailable != nil {
		s.D.Set("updates_available", *s.Res.UpdatesAvailable)
	}

	if s.Res.WorkRequestCount != nil {
		s.D.Set("work_request_count", *s.Res.WorkRequestCount)
	}

	return nil
}
