// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v46/core"
)

func init() {
	RegisterDatasource("oci_core_volume_backups", CoreVolumeBackupsDataSource())
}

func CoreVolumeBackupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreVolumeBackups,
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
			"source_volume_backup_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"volume_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"volume_backups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(CoreVolumeBackupResource()),
			},
		},
	}
}

func readCoreVolumeBackups(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeBackupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient()

	return ReadResource(sync)
}

type CoreVolumeBackupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.BlockstorageClient
	Res    *oci_core.ListVolumeBackupsResponse
}

func (s *CoreVolumeBackupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreVolumeBackupsDataSourceCrud) Get() error {
	request := oci_core.ListVolumeBackupsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if sourceVolumeBackupId, ok := s.D.GetOkExists("source_volume_backup_id"); ok {
		tmp := sourceVolumeBackupId.(string)
		request.SourceVolumeBackupId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.VolumeBackupLifecycleStateEnum(state.(string))
	}

	if volumeId, ok := s.D.GetOkExists("volume_id"); ok {
		tmp := volumeId.(string)
		request.VolumeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListVolumeBackups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListVolumeBackups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreVolumeBackupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("CoreVolumeBackupsDataSource-", CoreVolumeBackupsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		volumeBackup := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			volumeBackup["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			volumeBackup["display_name"] = *r.DisplayName
		}

		if r.ExpirationTime != nil {
			volumeBackup["expiration_time"] = r.ExpirationTime.String()
		}

		volumeBackup["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			volumeBackup["id"] = *r.Id
		}

		if r.KmsKeyId != nil {
			volumeBackup["kms_key_id"] = *r.KmsKeyId
		}

		if r.SizeInGBs != nil {
			volumeBackup["size_in_gbs"] = strconv.FormatInt(*r.SizeInGBs, 10)
		}

		if r.SizeInMBs != nil {
			volumeBackup["size_in_mbs"] = strconv.FormatInt(*r.SizeInMBs, 10)
		}

		volumeBackup["source_type"] = r.SourceType

		if r.SourceVolumeBackupId != nil {
			volumeBackup["source_volume_backup_id"] = *r.SourceVolumeBackupId
		}

		volumeBackup["state"] = r.LifecycleState

		if r.SystemTags != nil {
			volumeBackup["system_tags"] = systemTagsToMap(r.SystemTags)
		}

		if r.TimeCreated != nil {
			volumeBackup["time_created"] = r.TimeCreated.String()
		}

		if r.TimeRequestReceived != nil {
			volumeBackup["time_request_received"] = r.TimeRequestReceived.String()
		}

		volumeBackup["type"] = r.Type

		if r.UniqueSizeInGBs != nil {
			volumeBackup["unique_size_in_gbs"] = strconv.FormatInt(*r.UniqueSizeInGBs, 10)
		}

		if r.UniqueSizeInMbs != nil {
			volumeBackup["unique_size_in_mbs"] = strconv.FormatInt(*r.UniqueSizeInMbs, 10)
		}

		if r.VolumeId != nil {
			volumeBackup["volume_id"] = *r.VolumeId
		}

		resources = append(resources, volumeBackup)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, CoreVolumeBackupsDataSource().Schema["volume_backups"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("volume_backups", resources); err != nil {
		return err
	}

	return nil
}
