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

func CoreBootVolumesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreBootVolumes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"volume_group_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"boot_volumes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreBootVolumeResource()),
			},
		},
	}
}

func readCoreBootVolumes(d *schema.ResourceData, m interface{}) error {
	sync := &CoreBootVolumesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockstorageClient()

	return tfresource.ReadResource(sync)
}

type CoreBootVolumesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.BlockstorageClient
	Res    *oci_core.ListBootVolumesResponse
}

func (s *CoreBootVolumesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreBootVolumesDataSourceCrud) Get() error {
	request := oci_core.ListBootVolumesRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if volumeGroupId, ok := s.D.GetOkExists("volume_group_id"); ok {
		tmp := volumeGroupId.(string)
		request.VolumeGroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListBootVolumes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListBootVolumes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreBootVolumesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreBootVolumesDataSource-", CoreBootVolumesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		bootVolume := map[string]interface{}{}

		if r.AutoTunedVpusPerGB != nil {
			bootVolume["auto_tuned_vpus_per_gb"] = strconv.FormatInt(*r.AutoTunedVpusPerGB, 10)
		}

		if r.AvailabilityDomain != nil {
			bootVolume["availability_domain"] = *r.AvailabilityDomain
		}

		autotunePolicies := []interface{}{}
		for _, item := range r.AutotunePolicies {
			autotunePolicies = append(autotunePolicies, BootVolumeAutotunePolicyToMap(item))
		}
		bootVolume["autotune_policies"] = autotunePolicies

		bootVolumeReplicas := []interface{}{}
		for _, item := range r.BootVolumeReplicas {
			bootVolumeReplicas = append(bootVolumeReplicas, BootVolumeReplicaInfoToMap(item))
		}
		bootVolume["boot_volume_replicas"] = bootVolumeReplicas

		if r.ClusterPlacementGroupId != nil {
			bootVolume["cluster_placement_group_id"] = *r.ClusterPlacementGroupId
		}

		if r.CompartmentId != nil {
			bootVolume["compartment_id"] = *r.CompartmentId
		}

		if r.DefinedTags != nil {
			bootVolume["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			bootVolume["display_name"] = *r.DisplayName
		}

		bootVolume["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			bootVolume["id"] = *r.Id
		}

		if r.ImageId != nil {
			bootVolume["image_id"] = *r.ImageId
		}

		if r.IsAutoTuneEnabled != nil {
			bootVolume["is_auto_tune_enabled"] = *r.IsAutoTuneEnabled
		}

		if r.IsHydrated != nil {
			bootVolume["is_hydrated"] = *r.IsHydrated
		}

		if r.KmsKeyId != nil {
			bootVolume["kms_key_id"] = *r.KmsKeyId
		}

		if r.SizeInGBs != nil {
			bootVolume["size_in_gbs"] = strconv.FormatInt(*r.SizeInGBs, 10)
		}

		if r.SizeInMBs != nil {
			bootVolume["size_in_mbs"] = strconv.FormatInt(*r.SizeInMBs, 10)
		}

		if r.SourceDetails != nil {
			sourceDetailsArray := []interface{}{}
			if sourceDetailsMap := BootVolumeSourceDetailsToMap(&r.SourceDetails); sourceDetailsMap != nil {
				sourceDetailsArray = append(sourceDetailsArray, sourceDetailsMap)
			}
			bootVolume["source_details"] = sourceDetailsArray
		} else {
			bootVolume["source_details"] = nil
		}

		bootVolume["state"] = r.LifecycleState

		if r.SystemTags != nil {
			bootVolume["system_tags"] = tfresource.SystemTagsToMap(r.SystemTags)
		}

		if r.TimeCreated != nil {
			bootVolume["time_created"] = r.TimeCreated.String()
		}

		if r.VolumeGroupId != nil {
			bootVolume["volume_group_id"] = *r.VolumeGroupId
		}

		if r.VpusPerGB != nil {
			bootVolume["vpus_per_gb"] = strconv.FormatInt(*r.VpusPerGB, 10)
		}

		resources = append(resources, bootVolume)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreBootVolumesDataSource().Schema["boot_volumes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("boot_volumes", resources); err != nil {
		return err
	}

	return nil
}
