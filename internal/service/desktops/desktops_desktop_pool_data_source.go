// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package desktops

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_desktops "github.com/oracle/oci-go-sdk/v65/desktops"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DesktopsDesktopPoolDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["desktop_pool_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DesktopsDesktopPoolResource(), fieldMap, readSingularDesktopsDesktopPool)
}

func readSingularDesktopsDesktopPool(d *schema.ResourceData, m interface{}) error {
	sync := &DesktopsDesktopPoolDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DesktopServiceClient()

	return tfresource.ReadResource(sync)
}

type DesktopsDesktopPoolDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_desktops.DesktopServiceClient
	Res    *oci_desktops.GetDesktopPoolResponse
}

func (s *DesktopsDesktopPoolDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DesktopsDesktopPoolDataSourceCrud) Get() error {
	request := oci_desktops.GetDesktopPoolRequest{}

	if desktopPoolId, ok := s.D.GetOkExists("desktop_pool_id"); ok {
		tmp := desktopPoolId.(string)
		request.DesktopPoolId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "desktops")

	response, err := s.Client.GetDesktopPool(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DesktopsDesktopPoolDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.ArePrivilegedUsers != nil {
		s.D.Set("are_privileged_users", *s.Res.ArePrivilegedUsers)
	}

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.AvailabilityPolicy != nil {
		s.D.Set("availability_policy", []interface{}{DesktopAvailabilityPolicyToMap(s.Res.AvailabilityPolicy)})
	} else {
		s.D.Set("availability_policy", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ContactDetails != nil {
		s.D.Set("contact_details", *s.Res.ContactDetails)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DevicePolicy != nil {
		s.D.Set("device_policy", []interface{}{DesktopDevicePolicyToMap(s.Res.DevicePolicy)})
	} else {
		s.D.Set("device_policy", nil)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Image != nil {
		s.D.Set("image", []interface{}{DesktopImageToMap(s.Res.Image)})
	} else {
		s.D.Set("image", nil)
	}

	if s.Res.IsStorageEnabled != nil {
		s.D.Set("is_storage_enabled", *s.Res.IsStorageEnabled)
	}

	if s.Res.MaximumSize != nil {
		s.D.Set("maximum_size", *s.Res.MaximumSize)
	}

	if s.Res.NetworkConfiguration != nil {
		s.D.Set("network_configuration", []interface{}{DesktopNetworkConfigurationToMap(s.Res.NetworkConfiguration)})
	} else {
		s.D.Set("network_configuration", nil)
	}

	s.D.Set("nsg_ids", s.Res.NsgIds)

	if s.Res.ShapeConfig != nil {
		s.D.Set("shape_config", []interface{}{DesktopPoolShapeConfigToMap(s.Res.ShapeConfig)})
	} else {
		s.D.Set("shape_config", nil)
	}

	if s.Res.PrivateAccessDetails != nil {
		s.D.Set("private_access_details", []interface{}{DesktopPoolPrivateAccessDetailsToMap(s.Res.PrivateAccessDetails, true)})
	} else {
		s.D.Set("private_access_details", nil)
	}

	if s.Res.SessionLifecycleActions != nil {
		s.D.Set("session_lifecycle_actions", []interface{}{DesktopSessionLifecycleActionsToMap(s.Res.SessionLifecycleActions)})
	} else {
		s.D.Set("session_lifecycle_actions", nil)
	}

	if s.Res.ShapeName != nil {
		s.D.Set("shape_name", *s.Res.ShapeName)
	}

	if s.Res.StandbySize != nil {
		s.D.Set("standby_size", *s.Res.StandbySize)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StorageBackupPolicyId != nil {
		s.D.Set("storage_backup_policy_id", *s.Res.StorageBackupPolicyId)
	}

	if s.Res.StorageSizeInGBs != nil {
		s.D.Set("storage_size_in_gbs", *s.Res.StorageSizeInGBs)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeStartScheduled != nil {
		s.D.Set("time_start_scheduled", s.Res.TimeStartScheduled.Format(time.RFC3339Nano))
	}

	if s.Res.TimeStopScheduled != nil {
		s.D.Set("time_stop_scheduled", s.Res.TimeStopScheduled.Format(time.RFC3339Nano))
	}

	s.D.Set("use_dedicated_vm_host", s.Res.UseDedicatedVmHost)

	return nil
}
