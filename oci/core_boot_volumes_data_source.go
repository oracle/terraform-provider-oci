// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v33/core"
)

func init() {
	RegisterDatasource("oci_core_boot_volumes", CoreBootVolumesDataSource())
}

func CoreBootVolumesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreBootVolumes,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"volume_group_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"boot_volumes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(CoreBootVolumeResource()),
			},
		},
	}
}

func readCoreBootVolumes(d *schema.ResourceData, m interface{}) error {
	sync := &CoreBootVolumesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient()

	return ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

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

	s.D.SetId(GenerateDataSourceHashID("CoreBootVolumesDataSource-", CoreBootVolumesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		bootVolume := map[string]interface{}{
			"availability_domain": *r.AvailabilityDomain,
			"compartment_id":      *r.CompartmentId,
		}

		if r.AutoTunedVpusPerGB != nil {
			bootVolume["auto_tuned_vpus_per_gb"] = strconv.FormatInt(*r.AutoTunedVpusPerGB, 10)
		}

		if r.DefinedTags != nil {
			bootVolume["defined_tags"] = definedTagsToMap(r.DefinedTags)
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
			bootVolume["system_tags"] = systemTagsToMap(r.SystemTags)
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
		resources = ApplyFilters(f.(*schema.Set), resources, CoreBootVolumesDataSource().Schema["boot_volumes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("boot_volumes", resources); err != nil {
		return err
	}

	return nil
}
