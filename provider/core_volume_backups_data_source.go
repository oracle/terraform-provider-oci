// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/oracle/terraform-provider-oci/crud"
)

func VolumeBackupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readVolumeBackups,
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
			"limit": {
				Type:       schema.TypeInt,
				Optional:   true,
				Deprecated: crud.FieldDeprecated("limit"),
			},
			"page": {
				Type:       schema.TypeString,
				Optional:   true,
				Deprecated: crud.FieldDeprecated("page"),
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
				Elem:     VolumeBackupResource(),
			},
		},
	}
}

func readVolumeBackups(d *schema.ResourceData, m interface{}) error {
	sync := &VolumeBackupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient

	return crud.ReadResource(sync)
}

type VolumeBackupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.BlockstorageClient
	Res    *oci_core.ListVolumeBackupsResponse
}

func (s *VolumeBackupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *VolumeBackupsDataSourceCrud) Get() error {
	request := oci_core.ListVolumeBackupsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if limit, ok := s.D.GetOkExists("limit"); ok {
		tmp := limit.(int)
		request.Limit = &tmp
	}

	if page, ok := s.D.GetOkExists("page"); ok {
		tmp := page.(string)
		request.Page = &tmp
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

func (s *VolumeBackupsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
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

		if r.SizeInGBs != nil {
			volumeBackup["size_in_gbs"] = *r.SizeInGBs
		}

		if r.SizeInMBs != nil {
			volumeBackup["size_in_mbs"] = *r.SizeInMBs
		}

		volumeBackup["source_type"] = r.SourceType

		volumeBackup["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			volumeBackup["time_created"] = r.TimeCreated.String()
		}

		if r.TimeRequestReceived != nil {
			volumeBackup["time_request_received"] = r.TimeRequestReceived.String()
		}

		volumeBackup["type"] = r.Type

		if r.UniqueSizeInGBs != nil {
			volumeBackup["unique_size_in_gbs"] = *r.UniqueSizeInGBs
		}

		if r.UniqueSizeInMbs != nil {
			volumeBackup["unique_size_in_mbs"] = *r.UniqueSizeInMbs
		}

		if r.VolumeId != nil {
			volumeBackup["volume_id"] = *r.VolumeId
		}

		resources = append(resources, volumeBackup)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, VolumeBackupsDataSource().Schema["volume_backups"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("volume_backups", resources); err != nil {
		panic(err)
	}

	return
}
