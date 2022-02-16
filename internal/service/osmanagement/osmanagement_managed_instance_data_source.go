// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package osmanagement

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_osmanagement "github.com/oracle/oci-go-sdk/v58/osmanagement"
)

func OsmanagementManagedInstanceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["managed_instance_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OsmanagementManagedInstanceResource(), fieldMap, readSingularOsmanagementManagedInstance)
}

func readSingularOsmanagementManagedInstance(d *schema.ResourceData, m interface{}) error {
	sync := &OsmanagementManagedInstanceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OsManagementClient()

	return tfresource.ReadResource(sync)
}

type OsmanagementManagedInstanceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_osmanagement.OsManagementClient
	Res    *oci_osmanagement.GetManagedInstanceResponse
}

func (s *OsmanagementManagedInstanceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsmanagementManagedInstanceDataSourceCrud) Get() error {
	request := oci_osmanagement.GetManagedInstanceRequest{}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "osmanagement")

	response, err := s.Client.GetManagedInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsmanagementManagedInstanceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.Autonomous != nil {
		s.D.Set("autonomous", []interface{}{AutonomousSettingsToMap(s.Res.Autonomous)})
	} else {
		s.D.Set("autonomous", nil)
	}

	if s.Res.BugUpdatesAvailable != nil {
		s.D.Set("bug_updates_available", *s.Res.BugUpdatesAvailable)
	}

	childSoftwareSources := []interface{}{}
	for _, item := range s.Res.ChildSoftwareSources {
		childSoftwareSources = append(childSoftwareSources, SoftwareSourceIdToMap(&item))
	}
	s.D.Set("child_software_sources", childSoftwareSources)

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

	if s.Res.IsDataCollectionAuthorized != nil {
		s.D.Set("is_data_collection_authorized", *s.Res.IsDataCollectionAuthorized)
	}

	if s.Res.IsRebootRequired != nil {
		s.D.Set("is_reboot_required", *s.Res.IsRebootRequired)
	}

	if s.Res.KspliceEffectiveKernelVersion != nil {
		s.D.Set("ksplice_effective_kernel_version", *s.Res.KspliceEffectiveKernelVersion)
	}

	if s.Res.LastBoot != nil {
		s.D.Set("last_boot", *s.Res.LastBoot)
	}

	if s.Res.LastCheckin != nil {
		s.D.Set("last_checkin", *s.Res.LastCheckin)
	}

	managedInstanceGroups := []interface{}{}
	for _, item := range s.Res.ManagedInstanceGroups {
		managedInstanceGroups = append(managedInstanceGroups, IdToMap(item))
	}
	s.D.Set("managed_instance_groups", managedInstanceGroups)

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

	if s.Res.ParentSoftwareSource != nil {
		s.D.Set("parent_software_source", []interface{}{SoftwareSourceIdToMap(s.Res.ParentSoftwareSource)})
	} else {
		s.D.Set("parent_software_source", nil)
	}

	if s.Res.ScheduledJobCount != nil {
		s.D.Set("scheduled_job_count", *s.Res.ScheduledJobCount)
	}

	if s.Res.SecurityUpdatesAvailable != nil {
		s.D.Set("security_updates_available", *s.Res.SecurityUpdatesAvailable)
	}

	s.D.Set("status", s.Res.Status)

	if s.Res.UpdatesAvailable != nil {
		s.D.Set("updates_available", *s.Res.UpdatesAvailable)
	}

	if s.Res.WorkRequestCount != nil {
		s.D.Set("work_request_count", *s.Res.WorkRequestCount)
	}

	return nil
}
