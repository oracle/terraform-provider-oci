// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/oracle/terraform-provider-oci/crud"
)

func BootVolumesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readBootVolumes,
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
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"availability_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"compartment_id": {
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
						"image_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"size_in_gbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"size_in_mbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"volume_group_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readBootVolumes(d *schema.ResourceData, m interface{}) error {
	sync := &BootVolumesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient

	return crud.ReadResource(sync)
}

type BootVolumesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.BlockstorageClient
	Res    *oci_core.ListBootVolumesResponse
}

func (s *BootVolumesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BootVolumesDataSourceCrud) Get() error {
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

func (s *BootVolumesDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		bootVolume := map[string]interface{}{
			"availability_domain": *r.AvailabilityDomain,
			"compartment_id":      *r.CompartmentId,
		}

		if r.DisplayName != nil {
			bootVolume["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			bootVolume["id"] = *r.Id
		}

		if r.ImageId != nil {
			bootVolume["image_id"] = *r.ImageId
		}

		if r.SizeInGBs != nil {
			bootVolume["size_in_gbs"] = *r.SizeInGBs
		}

		if r.SizeInMBs != nil {
			bootVolume["size_in_mbs"] = *r.SizeInMBs
		}

		bootVolume["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			bootVolume["time_created"] = r.TimeCreated.String()
		}

		if r.VolumeGroupId != nil {
			bootVolume["volume_group_id"] = *r.VolumeGroupId
		}

		resources = append(resources, bootVolume)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, BootVolumesDataSource().Schema["boot_volumes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("boot_volumes", resources); err != nil {
		panic(err)
	}

	return
}
