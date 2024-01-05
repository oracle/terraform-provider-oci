// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"strconv"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"
)

func CoreBootVolumeDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["boot_volume_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CoreBootVolumeResource(), fieldMap, readSingularCoreBootVolume)
}

func readSingularCoreBootVolume(d *schema.ResourceData, m interface{}) error {
	sync := &CoreBootVolumeDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockstorageClient()

	return tfresource.ReadResource(sync)
}

type CoreBootVolumeDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.BlockstorageClient
	Res    *oci_core.GetBootVolumeResponse
}

func (s *CoreBootVolumeDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreBootVolumeDataSourceCrud) Get() error {
	request := oci_core.GetBootVolumeRequest{}

	if bootVolumeId, ok := s.D.GetOkExists("boot_volume_id"); ok {
		tmp := bootVolumeId.(string)
		request.BootVolumeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetBootVolume(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreBootVolumeDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AutoTunedVpusPerGB != nil {
		s.D.Set("auto_tuned_vpus_per_gb", strconv.FormatInt(*s.Res.AutoTunedVpusPerGB, 10))
	}

	autotunePolicies := []interface{}{}
	for _, item := range s.Res.AutotunePolicies {
		autotunePolicies = append(autotunePolicies, BootVolumeAutotunePolicyToMap(item))
	}
	s.D.Set("autotune_policies", autotunePolicies)

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	bootVolumeReplicas := []interface{}{}
	for _, item := range s.Res.BootVolumeReplicas {
		bootVolumeReplicas = append(bootVolumeReplicas, BootVolumeReplicaInfoToMap(item))
	}
	s.D.Set("boot_volume_replicas", bootVolumeReplicas)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.ImageId != nil {
		s.D.Set("image_id", *s.Res.ImageId)
	}

	if s.Res.IsAutoTuneEnabled != nil {
		s.D.Set("is_auto_tune_enabled", *s.Res.IsAutoTuneEnabled)
	}

	if s.Res.IsHydrated != nil {
		s.D.Set("is_hydrated", *s.Res.IsHydrated)
	}

	if s.Res.KmsKeyId != nil {
		s.D.Set("kms_key_id", *s.Res.KmsKeyId)
	}

	if s.Res.SizeInGBs != nil {
		s.D.Set("size_in_gbs", strconv.FormatInt(*s.Res.SizeInGBs, 10))
	}

	if s.Res.SizeInMBs != nil {
		s.D.Set("size_in_mbs", strconv.FormatInt(*s.Res.SizeInMBs, 10))
	}

	if s.Res.SourceDetails != nil {
		sourceDetailsArray := []interface{}{}
		if sourceDetailsMap := BootVolumeSourceDetailsToMap(&s.Res.SourceDetails); sourceDetailsMap != nil {
			sourceDetailsArray = append(sourceDetailsArray, sourceDetailsMap)
		}
		s.D.Set("source_details", sourceDetailsArray)
	} else {
		s.D.Set("source_details", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VolumeGroupId != nil {
		s.D.Set("volume_group_id", *s.Res.VolumeGroupId)
	}

	if s.Res.VpusPerGB != nil {
		s.D.Set("vpus_per_gb", strconv.FormatInt(*s.Res.VpusPerGB, 10))
	}

	return nil
}
