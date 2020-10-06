// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"

	oci_core "github.com/oracle/oci-go-sdk/v26/core"
)

func init() {
	RegisterResource("oci_core_volume_backup", CoreVolumeBackupResource())
}

func CoreVolumeBackupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createCoreVolumeBackup,
		Read:     readCoreVolumeBackup,
		Update:   updateCoreVolumeBackup,
		Delete:   deleteCoreVolumeBackup,
		Schema: map[string]*schema.Schema{
			// Optional
			"volume_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"source_details": {
				Type:          schema.TypeList,
				Optional:      true,
				ForceNew:      true,
				MaxItems:      1,
				MinItems:      1,
				ConflictsWith: []string{"volume_id", "defined_tags", "freeform_tags", "type"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"region": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						// Optional
						"kms_key_id": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						// Required
						"volume_backup_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
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
			"kms_key_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size_in_gbs": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size_in_mbs": {
				Type:       schema.TypeString,
				Computed:   true,
				Deprecated: FieldDeprecatedForAnother("size_in_mbs", "size_in_gbs"),
			},
			"source_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_volume_backup_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
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
				Type:       schema.TypeString,
				Computed:   true,
				Deprecated: FieldDeprecatedForAnother("unique_size_in_mbs", "unique_size_in_gbs"),
			},
		},
	}
}

func createCoreVolumeBackup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient()

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

func readCoreVolumeBackup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient()

	return ReadResource(sync)
}

func updateCoreVolumeBackup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient()

	return UpdateResource(d, sync)
}

func deleteCoreVolumeBackup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type CoreVolumeBackupResourceCrud struct {
	BaseCrud
	Client                 *oci_core.BlockstorageClient
	SourceRegionClient     *oci_core.BlockstorageClient
	Res                    *oci_core.VolumeBackup
	DisableNotFoundRetries bool
}

func (s *CoreVolumeBackupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreVolumeBackupResourceCrud) CreatedPending() []string {
	// Creating is considered "Created" because it can take some time to finish
	// actually creating and uploading the backup.
	return []string{
		string(oci_core.VolumeBackupLifecycleStateCreating),
		string(oci_core.VolumeBackupLifecycleStateRequestReceived),
	}
}

func (s *CoreVolumeBackupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.VolumeBackupLifecycleStateAvailable),
	}
}

func (s *CoreVolumeBackupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.VolumeBackupLifecycleStateTerminating),
	}
}

func (s *CoreVolumeBackupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.VolumeBackupLifecycleStateTerminated),
	}
}

func (s *CoreVolumeBackupResourceCrud) Create() error {
	if s.isCopyCreate() {
		return s.createVolumeBackupCopy()
	}

	return s.CreateVolumeBackup()
}

func (s *CoreVolumeBackupResourceCrud) isCopyCreate() bool {
	if sourceDetails, ok := s.D.GetOkExists("source_details"); ok {
		if tmpList := sourceDetails.([]interface{}); len(tmpList) > 0 {
			return true
		}
	}
	return false
}

func (s *CoreVolumeBackupResourceCrud) createVolumeBackupCopy() error {
	copyVolumeBackupRequest := oci_core.CopyVolumeBackupRequest{}

	configProvider := *s.Client.ConfigurationProvider()
	if configProvider == nil {
		return fmt.Errorf("cannot access ConfigurationProvider")
	}
	currentRegion, error := configProvider.Region()
	if error != nil {
		return fmt.Errorf("cannot access Region for the current ConfigurationProvider")
	}

	if sourceDetails, ok := s.D.GetOkExists("source_details"); ok && sourceDetails != nil {
		fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source_details", 0)

		if region, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			err := s.createBlockStorageSourceRegionClient(tmp)
			if err != nil {
				return err
			}
		}
		copyVolumeBackupRequest.DestinationRegion = &currentRegion

		if volumeBackupId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "volume_backup_id")); ok {
			tmp := volumeBackupId.(string)
			copyVolumeBackupRequest.VolumeBackupId = &tmp
		}

		if kmsKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_id")); ok {
			tmp := kmsKeyId.(string)
			copyVolumeBackupRequest.KmsKeyId = &tmp
		}

	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		copyVolumeBackupRequest.DisplayName = &tmp
	}

	response, err := s.SourceRegionClient.CopyVolumeBackup(context.Background(), copyVolumeBackupRequest)
	if err != nil {
		return err
	}

	s.Res = &response.VolumeBackup
	return nil
}

func (s *CoreVolumeBackupResourceCrud) CreateVolumeBackup() error {
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

func (s *CoreVolumeBackupResourceCrud) Get() error {
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

func (s *CoreVolumeBackupResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
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

func (s *CoreVolumeBackupResourceCrud) Delete() error {
	request := oci_core.DeleteVolumeBackupRequest{}

	tmp := s.D.Id()
	request.VolumeBackupId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteVolumeBackup(context.Background(), request)
	return err
}

func (s *CoreVolumeBackupResourceCrud) SetData() error {
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

	if s.Res.KmsKeyId != nil {
		s.D.Set("kms_key_id", *s.Res.KmsKeyId)
	}

	if s.Res.SizeInGBs != nil {
		s.D.Set("size_in_gbs", strconv.FormatInt(*s.Res.SizeInGBs, 10))
	}

	if s.Res.SizeInMBs != nil {
		s.D.Set("size_in_mbs", strconv.FormatInt(*s.Res.SizeInMBs, 10))
	}

	s.D.Set("source_type", s.Res.SourceType)

	if s.Res.SourceVolumeBackupId != nil {
		s.D.Set("source_volume_backup_id", *s.Res.SourceVolumeBackupId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", systemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeRequestReceived != nil {
		s.D.Set("time_request_received", s.Res.TimeRequestReceived.String())
	}

	s.D.Set("type", s.Res.Type)

	if s.Res.UniqueSizeInGBs != nil {
		s.D.Set("unique_size_in_gbs", strconv.FormatInt(*s.Res.UniqueSizeInGBs, 10))
	}

	if s.Res.UniqueSizeInMbs != nil {
		s.D.Set("unique_size_in_mbs", strconv.FormatInt(*s.Res.UniqueSizeInMbs, 10))
	}

	if s.Res.VolumeId != nil {
		s.D.Set("volume_id", *s.Res.VolumeId)
	}

	return nil
}

func (s *CoreVolumeBackupResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeVolumeBackupCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.VolumeBackupId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ChangeVolumeBackupCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}
