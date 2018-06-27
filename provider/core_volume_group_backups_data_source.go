// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/oracle/terraform-provider-oci/crud"
)

func VolumeGroupBackupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readVolumeGroupBackups,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"volume_group_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"volume_group_backups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     VolumeGroupBackupResource(),
			},
		},
	}
}

func readVolumeGroupBackups(d *schema.ResourceData, m interface{}) error {
	sync := &VolumeGroupBackupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient

	return crud.ReadResource(sync)
}

type VolumeGroupBackupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.BlockstorageClient
	Res    *oci_core.ListVolumeGroupBackupsResponse
}

func (s *VolumeGroupBackupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *VolumeGroupBackupsDataSourceCrud) Get() error {
	request := oci_core.ListVolumeGroupBackupsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if volumeGroupId, ok := s.D.GetOkExists("volume_group_id"); ok {
		tmp := volumeGroupId.(string)
		request.VolumeGroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListVolumeGroupBackups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListVolumeGroupBackups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *VolumeGroupBackupsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		volumeGroupBackup := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			volumeGroupBackup["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			volumeGroupBackup["display_name"] = *r.DisplayName
		}

		volumeGroupBackup["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			volumeGroupBackup["id"] = *r.Id
		}

		if r.SizeInMBs != nil {
			volumeGroupBackup["size_in_mbs"] = *r.SizeInMBs
		}

		volumeGroupBackup["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			volumeGroupBackup["time_created"] = r.TimeCreated.String()
		}

		if r.TimeRequestReceived != nil {
			volumeGroupBackup["time_request_received"] = r.TimeRequestReceived.String()
		}

		volumeGroupBackup["type"] = r.Type

		if r.UniqueSizeInMbs != nil {
			volumeGroupBackup["unique_size_in_mbs"] = *r.UniqueSizeInMbs
		}

		volumeGroupBackup["volume_backup_ids"] = r.VolumeBackupIds

		if r.VolumeGroupId != nil {
			volumeGroupBackup["volume_group_id"] = *r.VolumeGroupId
		}

		resources = append(resources, volumeGroupBackup)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, VolumeGroupBackupsDataSource().Schema["volume_group_backups"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("volume_group_backups", resources); err != nil {
		panic(err)
	}

	return
}
