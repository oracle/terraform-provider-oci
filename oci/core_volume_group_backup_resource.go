// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_core "github.com/oracle/oci-go-sdk/v37/core"
)

func init() {
	RegisterResource("oci_core_volume_group_backup", CoreVolumeGroupBackupResource())
}

func CoreVolumeGroupBackupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createCoreVolumeGroupBackup,
		Read:     readCoreVolumeGroupBackup,
		Update:   updateCoreVolumeGroupBackup,
		Delete:   deleteCoreVolumeGroupBackup,
		Schema: map[string]*schema.Schema{
			// Required
			"volume_group_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
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
			"expiration_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size_in_gbs": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size_in_mbs": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_volume_group_backup_id": {
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
				Type:     schema.TypeString,
				Computed: true,
			},
			"unique_size_in_mbs": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"volume_backup_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func createCoreVolumeGroupBackup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeGroupBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient()

	// Issue logged with service team for `create` not supporting non-default compartment_id
	// Remove custom code after issue is fixed.
	compartment, ok := sync.D.GetOkExists("compartment_id")

	err := CreateResource(d, sync)
	if err != nil {
		return err
	}
	if ok && compartment != *sync.Res.CompartmentId {
		err = sync.updateCompartment(compartment)
		if err != nil {
			return err
		}
		tmp := compartment.(string)
		sync.Res.CompartmentId = &tmp
		err := sync.Get()
		if err != nil {
			log.Printf("error doing a Get() after compartment update: %v", err)
		}
		err = sync.SetData()
		if err != nil {
			log.Printf("error doing a SetData() after compartment update: %v", err)
		}
	}
	return nil
}

func readCoreVolumeGroupBackup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeGroupBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient()

	return ReadResource(sync)
}

func updateCoreVolumeGroupBackup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeGroupBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient()

	return UpdateResource(d, sync)
}

func deleteCoreVolumeGroupBackup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeGroupBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type CoreVolumeGroupBackupResourceCrud struct {
	BaseCrud
	Client                 *oci_core.BlockstorageClient
	Res                    *oci_core.VolumeGroupBackup
	DisableNotFoundRetries bool
}

func (s *CoreVolumeGroupBackupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreVolumeGroupBackupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.VolumeGroupBackupLifecycleStateCreating),
		string(oci_core.VolumeGroupBackupLifecycleStateRequestReceived),
	}
}

func (s *CoreVolumeGroupBackupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.VolumeGroupBackupLifecycleStateCommitted),
		string(oci_core.VolumeGroupBackupLifecycleStateAvailable),
	}
}

func (s *CoreVolumeGroupBackupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.VolumeGroupBackupLifecycleStateTerminating),
	}
}

func (s *CoreVolumeGroupBackupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.VolumeGroupBackupLifecycleStateTerminated),
	}
}

func (s *CoreVolumeGroupBackupResourceCrud) Create() error {
	request := oci_core.CreateVolumeGroupBackupRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

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
		request.Type = oci_core.CreateVolumeGroupBackupDetailsTypeEnum(type_.(string))
	}

	if volumeGroupId, ok := s.D.GetOkExists("volume_group_id"); ok {
		tmp := volumeGroupId.(string)
		request.VolumeGroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateVolumeGroupBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VolumeGroupBackup
	return nil
}

func (s *CoreVolumeGroupBackupResourceCrud) Get() error {
	request := oci_core.GetVolumeGroupBackupRequest{}

	tmp := s.D.Id()
	request.VolumeGroupBackupId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetVolumeGroupBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VolumeGroupBackup
	return nil
}

func (s *CoreVolumeGroupBackupResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdateVolumeGroupBackupRequest{}

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
	request.VolumeGroupBackupId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateVolumeGroupBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VolumeGroupBackup
	return nil
}

func (s *CoreVolumeGroupBackupResourceCrud) Delete() error {
	request := oci_core.DeleteVolumeGroupBackupRequest{}

	tmp := s.D.Id()
	request.VolumeGroupBackupId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteVolumeGroupBackup(context.Background(), request)
	return err
}

func (s *CoreVolumeGroupBackupResourceCrud) SetData() error {
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

	if s.Res.SizeInGBs != nil {
		s.D.Set("size_in_gbs", strconv.FormatInt(*s.Res.SizeInGBs, 10))
	}

	if s.Res.SizeInMBs != nil {
		s.D.Set("size_in_mbs", strconv.FormatInt(*s.Res.SizeInMBs, 10))
	}

	s.D.Set("source_type", s.Res.SourceType)

	if s.Res.SourceVolumeGroupBackupId != nil {
		s.D.Set("source_volume_group_backup_id", *s.Res.SourceVolumeGroupBackupId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeRequestReceived != nil {
		s.D.Set("time_request_received", s.Res.TimeRequestReceived.String())
	}

	s.D.Set("type", s.Res.Type)

	if s.Res.UniqueSizeInGbs != nil {
		s.D.Set("unique_size_in_gbs", strconv.FormatInt(*s.Res.UniqueSizeInGbs, 10))
	}

	if s.Res.UniqueSizeInMbs != nil {
		s.D.Set("unique_size_in_mbs", strconv.FormatInt(*s.Res.UniqueSizeInMbs, 10))
	}

	s.D.Set("volume_backup_ids", s.Res.VolumeBackupIds)

	if s.Res.VolumeGroupId != nil {
		s.D.Set("volume_group_id", *s.Res.VolumeGroupId)
	}

	return nil
}

func (s *CoreVolumeGroupBackupResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeVolumeGroupBackupCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.VolumeGroupBackupId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ChangeVolumeGroupBackupCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}
