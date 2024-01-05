// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"
)

func CoreVolumeGroupBackupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreVolumeGroupBackup,
		Read:     readCoreVolumeGroupBackup,
		Update:   updateCoreVolumeGroupBackup,
		Delete:   deleteCoreVolumeGroupBackup,
		Schema: map[string]*schema.Schema{
			"volume_group_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				Computed:      true,
				ConflictsWith: []string{"source_details"},
			},

			"source_details": {
				Type:          schema.TypeList,
				Optional:      true,
				ForceNew:      true,
				MaxItems:      1,
				MinItems:      1,
				ConflictsWith: []string{"volume_group_id"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"region": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"volume_group_backup_id": {
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
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
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
	sync.Client = m.(*client.OracleClients).BlockstorageClient()

	// Issue logged with service team for `Create` not supporting non-default compartment_id
	// Remove custom code after issue is fixed.
	compartment, ok := sync.D.GetOkExists("compartment_id")

	err := tfresource.CreateResource(d, sync)
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
			log.Printf("error doing a Get() after compartment Update: %v", err)
		}
		err = sync.SetData()
		if err != nil {
			log.Printf("error doing a SetData() after compartment Update: %v", err)
		}
	}
	return nil
}

func readCoreVolumeGroupBackup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeGroupBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockstorageClient()

	return tfresource.ReadResource(sync)
}

func updateCoreVolumeGroupBackup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeGroupBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockstorageClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreVolumeGroupBackup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeGroupBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockstorageClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CoreVolumeGroupBackupResourceCrud struct {
	tfresource.BaseCrud
	SourceRegionClient     *oci_core.BlockstorageClient
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
		string(oci_core.VolumeGroupBackupLifecycleStateCommitted),
	}
}

func (s *CoreVolumeGroupBackupResourceCrud) CreatedTarget() []string {
	return []string{
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
	if s.isCopyCreate() {
		err := s.createVolumeGroupBackupCopy()
		if err != nil {
			return err
		}
		s.D.SetId(*s.Res.Id)
		err = tfresource.WaitForResourceCondition(s, func() bool { return s.Res.LifecycleState == oci_core.VolumeGroupBackupLifecycleStateAvailable }, s.D.Timeout(schema.TimeoutCreate))
		if err != nil {
			return err
		}
		// Update for some fields that can't be created by copy
		return s.Update()
	}

	return s.CreateVolumeGroupBackup()
}

func (s *CoreVolumeGroupBackupResourceCrud) isCopyCreate() bool {
	if sourceDetails, ok := s.D.GetOkExists("source_details"); ok {
		if tmpList := sourceDetails.([]interface{}); len(tmpList) > 0 {
			return true
		}
	}
	return false
}

func (s *CoreVolumeGroupBackupResourceCrud) CreateVolumeGroupBackup() error {
	request := oci_core.CreateVolumeGroupBackupRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_core.CreateVolumeGroupBackupDetailsTypeEnum(type_.(string))
	}

	if volumeGroupId, ok := s.D.GetOkExists("volume_group_id"); ok {
		tmp := volumeGroupId.(string)
		request.VolumeGroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateVolumeGroupBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VolumeGroupBackup
	return nil
}

func (s *CoreVolumeGroupBackupResourceCrud) createVolumeGroupBackupCopy() error {
	copyVolumeGroupBackupRequest := oci_core.CopyVolumeGroupBackupRequest{}

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
		copyVolumeGroupBackupRequest.DestinationRegion = &currentRegion

		if volumeGroupBackupId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "volume_group_backup_id")); ok {
			tmp := volumeGroupBackupId.(string)
			copyVolumeGroupBackupRequest.VolumeGroupBackupId = &tmp
		}

		if kmsKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_id")); ok {
			tmp := kmsKeyId.(string)
			copyVolumeGroupBackupRequest.KmsKeyId = &tmp
		}

	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		copyVolumeGroupBackupRequest.DisplayName = &tmp
	}

	response, err := s.SourceRegionClient.CopyVolumeGroupBackup(context.Background(), copyVolumeGroupBackupRequest)
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

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

	//check if there are any fields to Update (empty Update request is invalid)
	hasAttributeSet := false

	request := oci_core.UpdateVolumeGroupBackupRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
		hasAttributeSet = true
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
		hasAttributeSet = true
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		hasAttributeSet = true
	}
	if !hasAttributeSet {
		return nil
	}

	tmp := s.D.Id()
	request.VolumeGroupBackupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteVolumeGroupBackup(context.Background(), request)
	return err
}

func (s *CoreVolumeGroupBackupResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
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

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ChangeVolumeGroupBackupCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
