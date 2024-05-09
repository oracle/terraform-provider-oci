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

func OsManagementHubManagedInstanceGroupDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["managed_instance_group_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OsManagementHubManagedInstanceGroupResource(), fieldMap, readSingularOsManagementHubManagedInstanceGroup)
}

func readSingularOsManagementHubManagedInstanceGroup(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubManagedInstanceGroupDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagedInstanceGroupClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubManagedInstanceGroupDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.ManagedInstanceGroupClient
	Res    *oci_os_management_hub.GetManagedInstanceGroupResponse
}

func (s *OsManagementHubManagedInstanceGroupDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubManagedInstanceGroupDataSourceCrud) Get() error {
	request := oci_os_management_hub.GetManagedInstanceGroupRequest{}

	if managedInstanceGroupId, ok := s.D.GetOkExists("managed_instance_group_id"); ok {
		tmp := managedInstanceGroupId.(string)
		request.ManagedInstanceGroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.GetManagedInstanceGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsManagementHubManagedInstanceGroupDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("arch_type", s.Res.ArchType)

	if s.Res.AutonomousSettings != nil {
		s.D.Set("autonomous_settings", []interface{}{AutonomousSettingsToMap(s.Res.AutonomousSettings)})
	} else {
		s.D.Set("autonomous_settings", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
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

	if s.Res.IsManagedByAutonomousLinux != nil {
		s.D.Set("is_managed_by_autonomous_linux", *s.Res.IsManagedByAutonomousLinux)
	}

	s.D.Set("location", s.Res.Location)

	if s.Res.ManagedInstanceCount != nil {
		s.D.Set("managed_instance_count", *s.Res.ManagedInstanceCount)
	}

	s.D.Set("managed_instance_ids", s.Res.ManagedInstanceIds)

	if s.Res.NotificationTopicId != nil {
		s.D.Set("notification_topic_id", *s.Res.NotificationTopicId)
	}

	s.D.Set("os_family", s.Res.OsFamily)

	if s.Res.PendingJobCount != nil {
		s.D.Set("pending_job_count", *s.Res.PendingJobCount)
	}

	softwareSources := []interface{}{}
	softwareSourceIds := []string{}
	for _, item := range s.Res.SoftwareSourceIds {
		softwareSources = append(softwareSources, SoftwareSourceDetailsToMap(item))
		softwareSourceIds = append(softwareSourceIds, *item.Id)
	}

	s.D.Set("software_source_ids", softwareSourceIds)

	s.D.Set("software_sources", softwareSources)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeModified != nil {
		s.D.Set("time_modified", s.Res.TimeModified.String())
	}

	s.D.Set("vendor_name", s.Res.VendorName)

	return nil
}
