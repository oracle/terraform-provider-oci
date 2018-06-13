// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func VolumeBackupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createVolumeBackup,
		Read:     readVolumeBackup,
		Update:   updateVolumeBackup,
		Delete:   deleteVolumeBackup,
		Schema: map[string]*schema.Schema{
			// Required
			"volume_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"expiration_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			// @Deprecated 2017: size_in_mbs => size_in_gbs
			"size_in_mbs": {
				Type:       schema.TypeInt,
				Computed:   true,
				Deprecated: crud.FieldDeprecatedForAnother("size_in_mbs", "size_in_gbs"),
			},
			"source_type": {
				Type:     schema.TypeString,
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
			"time_request_received": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"unique_size_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			// @Deprecated 2017: unique_size_in_mbs => unique_size_in_gbs
			"unique_size_in_mbs": {
				Type:       schema.TypeInt,
				Computed:   true,
				Deprecated: crud.FieldDeprecatedForAnother("unique_size_in_mbs", "unique_size_in_gbs"),
			},
		},
	}
}

func createVolumeBackup(d *schema.ResourceData, m interface{}) error {
	sync := &VolumeBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient

	return crud.CreateResource(d, sync)
}

func readVolumeBackup(d *schema.ResourceData, m interface{}) error {
	sync := &VolumeBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient

	return crud.ReadResource(sync)
}

func updateVolumeBackup(d *schema.ResourceData, m interface{}) error {
	sync := &VolumeBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient

	return crud.UpdateResource(d, sync)
}

func deleteVolumeBackup(d *schema.ResourceData, m interface{}) error {
	sync := &VolumeBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type VolumeBackupResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_core.BlockstorageClient
	Res                    *oci_core.VolumeBackup
	DisableNotFoundRetries bool
}

func (s *VolumeBackupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *VolumeBackupResourceCrud) CreatedPending() []string {
	// Creating is considered "Created" because it can take some time to finish
	// actually creating and uploading the backup.
	return []string{
		string(oci_core.VolumeBackupLifecycleStateCreating),
		string(oci_core.VolumeBackupLifecycleStateRequestReceived),
	}
}

func (s *VolumeBackupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.VolumeBackupLifecycleStateAvailable),
	}
}

func (s *VolumeBackupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.VolumeBackupLifecycleStateTerminating),
	}
}

func (s *VolumeBackupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.VolumeBackupLifecycleStateTerminated),
	}
}

func (s *VolumeBackupResourceCrud) Create() error {
	request := oci_core.CreateVolumeBackupRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_core.CreateVolumeBackupDetailsTypeEnum(type_.(string))
	}

	if volumeId, ok := s.D.GetOkExists("volume_id"); ok {
		tmp := volumeId.(string)
		request.VolumeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateVolumeBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VolumeBackup
	return nil
}

func (s *VolumeBackupResourceCrud) Get() error {
	request := oci_core.GetVolumeBackupRequest{}

	tmp := s.D.Id()
	request.VolumeBackupId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetVolumeBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VolumeBackup
	return nil
}

func (s *VolumeBackupResourceCrud) Update() error {
	request := oci_core.UpdateVolumeBackupRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.VolumeBackupId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateVolumeBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VolumeBackup
	return nil
}

func (s *VolumeBackupResourceCrud) Delete() error {
	request := oci_core.DeleteVolumeBackupRequest{}

	tmp := s.D.Id()
	request.VolumeBackupId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteVolumeBackup(context.Background(), request)
	return err
}

func (s *VolumeBackupResourceCrud) SetData() {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ExpirationTime != nil {
		s.D.Set("expiration_time", s.Res.ExpirationTime.String())
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	if s.Res.SizeInGBs != nil {
		s.D.Set("size_in_gbs", *s.Res.SizeInGBs)
	}

	if s.Res.SizeInMBs != nil {
		s.D.Set("size_in_mbs", *s.Res.SizeInMBs)
	}

	s.D.Set("source_type", s.Res.SourceType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeRequestReceived != nil {
		s.D.Set("time_request_received", s.Res.TimeRequestReceived.String())
	}

	s.D.Set("type", s.Res.Type)

	if s.Res.UniqueSizeInGBs != nil {
		s.D.Set("unique_size_in_gbs", *s.Res.UniqueSizeInGBs)
	}

	if s.Res.UniqueSizeInMbs != nil {
		s.D.Set("unique_size_in_mbs", *s.Res.UniqueSizeInMbs)
	}

	if s.Res.VolumeId != nil {
		s.D.Set("volume_id", *s.Res.VolumeId)
	}

}
