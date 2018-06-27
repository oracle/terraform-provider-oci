// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/oracle/terraform-provider-oci/crud"
)

func VolumeGroupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readVolumeGroups,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"volume_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     VolumeGroupResource(),
			},
		},
	}
}

func readVolumeGroups(d *schema.ResourceData, m interface{}) error {
	sync := &VolumeGroupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient

	return crud.ReadResource(sync)
}

type VolumeGroupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.BlockstorageClient
	Res    *oci_core.ListVolumeGroupsResponse
}

func (s *VolumeGroupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *VolumeGroupsDataSourceCrud) Get() error {
	request := oci_core.ListVolumeGroupsRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.VolumeGroupLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListVolumeGroups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListVolumeGroups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *VolumeGroupsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		volumeGroup := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AvailabilityDomain != nil {
			volumeGroup["availability_domain"] = *r.AvailabilityDomain
		}

		if r.DefinedTags != nil {
			volumeGroup["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			volumeGroup["display_name"] = *r.DisplayName
		}

		volumeGroup["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			volumeGroup["id"] = *r.Id
		}

		if r.SizeInMBs != nil {
			volumeGroup["size_in_mbs"] = *r.SizeInMBs
		}

		volumeGroup["source_details"] = VolumeGroupSourceDetailsToMap(r.SourceDetails)

		volumeGroup["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			volumeGroup["time_created"] = r.TimeCreated.String()
		}

		volumeGroup["volume_ids"] = r.VolumeIds

		resources = append(resources, volumeGroup)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, VolumeGroupsDataSource().Schema["volume_groups"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("volume_groups", resources); err != nil {
		panic(err)
	}

	return
}
