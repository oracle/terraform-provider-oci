// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package lustre_file_storage

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_lustre_file_storage "github.com/oracle/oci-go-sdk/v65/lustrefilestorage"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LustreFileStorageLustreFileSystemDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["lustre_file_system_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(LustreFileStorageLustreFileSystemResource(), fieldMap, readSingularLustreFileStorageLustreFileSystem)
}

func readSingularLustreFileStorageLustreFileSystem(d *schema.ResourceData, m interface{}) error {
	sync := &LustreFileStorageLustreFileSystemDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LustreFileStorageClient()

	return tfresource.ReadResource(sync)
}

type LustreFileStorageLustreFileSystemDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_lustre_file_storage.LustreFileStorageClient
	Res    *oci_lustre_file_storage.GetLustreFileSystemResponse
}

func (s *LustreFileStorageLustreFileSystemDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LustreFileStorageLustreFileSystemDataSourceCrud) Get() error {
	request := oci_lustre_file_storage.GetLustreFileSystemRequest{}

	if lustreFileSystemId, ok := s.D.GetOkExists("lustre_file_system_id"); ok {
		tmp := lustreFileSystemId.(string)
		request.LustreFileSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "lustre_file_storage")

	response, err := s.Client.GetLustreFileSystem(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LustreFileStorageLustreFileSystemDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CapacityInGBs != nil {
		s.D.Set("capacity_in_gbs", *s.Res.CapacityInGBs)
	}

	if s.Res.ClusterPlacementGroupId != nil {
		s.D.Set("cluster_placement_group_id", *s.Res.ClusterPlacementGroupId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.FileSystemDescription != nil {
		s.D.Set("file_system_description", *s.Res.FileSystemDescription)
	}

	if s.Res.FileSystemName != nil {
		s.D.Set("file_system_name", *s.Res.FileSystemName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.KmsKeyId != nil {
		s.D.Set("kms_key_id", *s.Res.KmsKeyId)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Lnet != nil {
		s.D.Set("lnet", *s.Res.Lnet)
	}

	if s.Res.MaintenanceWindow != nil {
		s.D.Set("maintenance_window", []interface{}{MaintenanceWindowToMap(s.Res.MaintenanceWindow)})
	} else {
		s.D.Set("maintenance_window", nil)
	}

	if s.Res.MajorVersion != nil {
		s.D.Set("major_version", *s.Res.MajorVersion)
	}

	if s.Res.ManagementServiceAddress != nil {
		s.D.Set("management_service_address", *s.Res.ManagementServiceAddress)
	}

	s.D.Set("nsg_ids", s.Res.NsgIds)

	s.D.Set("performance_tier", s.Res.PerformanceTier)

	if s.Res.RootSquashConfiguration != nil {
		s.D.Set("root_squash_configuration", []interface{}{RootSquashConfigurationToMap(s.Res.RootSquashConfiguration)})
	} else {
		s.D.Set("root_squash_configuration", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeBillingCycleEnd != nil {
		s.D.Set("time_billing_cycle_end", s.Res.TimeBillingCycleEnd.String())
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
