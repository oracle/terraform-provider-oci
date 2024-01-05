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

func OsManagementHubManagementStationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["management_station_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OsManagementHubManagementStationResource(), fieldMap, readSingularOsManagementHubManagementStation)
}

func readSingularOsManagementHubManagementStation(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubManagementStationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementStationClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubManagementStationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.ManagementStationClient
	Res    *oci_os_management_hub.GetManagementStationResponse
}

func (s *OsManagementHubManagementStationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubManagementStationDataSourceCrud) Get() error {
	request := oci_os_management_hub.GetManagementStationRequest{}

	if managementStationId, ok := s.D.GetOkExists("management_station_id"); ok {
		tmp := managementStationId.(string)
		request.ManagementStationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.GetManagementStation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsManagementHubManagementStationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

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

	if s.Res.Hostname != nil {
		s.D.Set("hostname", *s.Res.Hostname)
	}

	if s.Res.ManagedInstanceId != nil {
		s.D.Set("managed_instance_id", *s.Res.ManagedInstanceId)
	}

	if s.Res.Mirror != nil {
		s.D.Set("mirror", []interface{}{MirrorConfigurationToMap(s.Res.Mirror)})
	} else {
		s.D.Set("mirror", nil)
	}

	if s.Res.MirrorCapacity != nil {
		s.D.Set("mirror_capacity", *s.Res.MirrorCapacity)
	}

	if s.Res.MirrorSyncStatus != nil {
		s.D.Set("mirror_sync_status", []interface{}{MirrorSyncStatusToMap(s.Res.MirrorSyncStatus)})
	} else {
		s.D.Set("mirror_sync_status", nil)
	}

	if s.Res.OverallPercentage != nil {
		s.D.Set("overall_percentage", *s.Res.OverallPercentage)
	}

	s.D.Set("overall_state", s.Res.OverallState)

	if s.Res.ProfileId != nil {
		s.D.Set("profile_id", *s.Res.ProfileId)
	}

	if s.Res.Proxy != nil {
		s.D.Set("proxy", []interface{}{ProxyConfigurationToMap(s.Res.Proxy)})
	} else {
		s.D.Set("proxy", nil)
	}

	if s.Res.ScheduledJobId != nil {
		s.D.Set("scheduled_job_id", *s.Res.ScheduledJobId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TotalMirrors != nil {
		s.D.Set("total_mirrors", *s.Res.TotalMirrors)
	}

	return nil
}
