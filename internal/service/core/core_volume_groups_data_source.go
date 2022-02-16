// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v58/core"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func CoreVolumeGroupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreVolumeGroups,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
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
				Elem:     tfresource.GetDataSourceItemSchema(CoreVolumeGroupResource()),
			},
		},
	}
}

func readCoreVolumeGroups(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeGroupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockstorageClient()

	return tfresource.ReadResource(sync)
}

type CoreVolumeGroupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.BlockstorageClient
	Res    *oci_core.ListVolumeGroupsResponse
}

func (s *CoreVolumeGroupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreVolumeGroupsDataSourceCrud) Get() error {
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

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

func (s *CoreVolumeGroupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreVolumeGroupsDataSource-", CoreVolumeGroupsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		volumeGroup := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AvailabilityDomain != nil {
			volumeGroup["availability_domain"] = *r.AvailabilityDomain
		}

		if r.DefinedTags != nil {
			volumeGroup["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			volumeGroup["display_name"] = *r.DisplayName
		}

		volumeGroup["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			volumeGroup["id"] = *r.Id
		}

		if r.IsHydrated != nil {
			volumeGroup["is_hydrated"] = *r.IsHydrated
		}

		if r.SizeInGBs != nil {
			volumeGroup["size_in_gbs"] = strconv.FormatInt(*r.SizeInGBs, 10)
		}

		if r.SizeInMBs != nil {
			volumeGroup["size_in_mbs"] = strconv.FormatInt(*r.SizeInMBs, 10)
		}

		if r.SourceDetails != nil {
			sourceDetailsArray := []interface{}{}
			if sourceDetailsMap := VolumeGroupSourceDetailsToMap(&r.SourceDetails, true); sourceDetailsMap != nil {
				sourceDetailsArray = append(sourceDetailsArray, sourceDetailsMap)
			}
			volumeGroup["source_details"] = sourceDetailsArray
		} else {
			volumeGroup["source_details"] = nil
		}

		volumeGroup["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			volumeGroup["time_created"] = r.TimeCreated.String()
		}

		volumeGroupReplicas := []interface{}{}
		for _, item := range r.VolumeGroupReplicas {
			volumeGroupReplicas = append(volumeGroupReplicas, VolumeGroupReplicaInfoToMap(item))
		}
		volumeGroup["volume_group_replicas"] = volumeGroupReplicas

		volumeGroup["volume_ids"] = r.VolumeIds

		resources = append(resources, volumeGroup)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreVolumeGroupsDataSource().Schema["volume_groups"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("volume_groups", resources); err != nil {
		return err
	}

	return nil
}
