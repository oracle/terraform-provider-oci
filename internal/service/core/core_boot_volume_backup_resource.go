// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"
)

func CoreBootVolumeBackupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createCoreBootVolumeBackupWithContext,
		ReadContext:   readCoreBootVolumeBackupWithContext,
		UpdateContext: updateCoreBootVolumeBackupWithContext,
		DeleteContext: deleteCoreBootVolumeBackupWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"boot_volume_id": {
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
				ConflictsWith: []string{"boot_volume_id"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"region": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"boot_volume_backup_id": {
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
			"kms_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"image_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size_in_gbs": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_boot_volume_backup_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_type": {
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
		},
	}
}

func createCoreBootVolumeBackupWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &CoreBootVolumeBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockstorageClient()
	sync.workRequestClient = m.(*client.OracleClients).WorkRequestClient

	compartment, ok := sync.D.GetOkExists("compartment_id")

	err := tfresource.CreateResourceWithContext(ctx, d, sync)
	if err != nil {
		return tfresource.HandleDiagError(m, err)
	}

	if ok && compartment != *sync.Res.CompartmentId {
		err = sync.updateCompartment(ctx, compartment)
		if err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		tmp := compartment.(string)
		sync.Res.CompartmentId = &tmp
		err := sync.GetWithContext(ctx)
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

func readCoreBootVolumeBackupWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &CoreBootVolumeBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockstorageClient()
	sync.workRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateCoreBootVolumeBackupWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &CoreBootVolumeBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockstorageClient()
	sync.workRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteCoreBootVolumeBackupWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &CoreBootVolumeBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockstorageClient()
	sync.workRequestClient = m.(*client.OracleClients).WorkRequestClient
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type CoreBootVolumeBackupResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.BlockstorageClient
	SourceRegionClient     *oci_core.BlockstorageClient
	workRequestClient      *oci_work_requests.WorkRequestClient
	Res                    *oci_core.BootVolumeBackup
	DisableNotFoundRetries bool
}

func (s *CoreBootVolumeBackupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreBootVolumeBackupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.BootVolumeBackupLifecycleStateCreating),
		string(oci_core.BootVolumeBackupLifecycleStateRequestReceived),
	}
}

func (s *CoreBootVolumeBackupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.BootVolumeBackupLifecycleStateAvailable),
	}
}

func (s *CoreBootVolumeBackupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.BootVolumeBackupLifecycleStateTerminating),
	}
}

func (s *CoreBootVolumeBackupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.BootVolumeBackupLifecycleStateTerminated),
	}
}

func (s *CoreBootVolumeBackupResourceCrud) CreateWithContext(ctx context.Context) error {
	if s.isCopyCreate() {
		err := s.createBootVolumeBackupCopy(ctx)
		if err != nil {
			return err
		}
		s.D.SetId(*s.Res.Id)
		err = tfresource.WaitForResourceConditionWithContext(ctx, s, func() bool { return s.Res.LifecycleState == oci_core.BootVolumeBackupLifecycleStateAvailable }, s.D.Timeout(schema.TimeoutCreate))
		if err != nil {
			return err
		}
		// Update for some fields that can't be created by copy
		return s.UpdateWithContext(ctx)
	}

	return s.createBootVolumeBackup(ctx)
}

func (s *CoreBootVolumeBackupResourceCrud) isCopyCreate() bool {
	if sourceDetails, ok := s.D.GetOkExists("source_details"); ok {
		if tmpList := sourceDetails.([]interface{}); len(tmpList) > 0 {
			return true
		}
	}
	return false
}

func (s *CoreBootVolumeBackupResourceCrud) createBootVolumeBackup(ctx context.Context) error {
	request := oci_core.CreateBootVolumeBackupRequest{}

	if bootVolumeId, ok := s.D.GetOkExists("boot_volume_id"); ok {
		tmp := bootVolumeId.(string)
		request.BootVolumeId = &tmp
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

	if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
		tmp := kmsKeyId.(string)
		request.KmsKeyId = &tmp
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_core.CreateBootVolumeBackupDetailsTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateBootVolumeBackup(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.BootVolumeBackup
	return nil
}

func (s *CoreBootVolumeBackupResourceCrud) createBootVolumeBackupCopy(ctx context.Context) error {
	copyBootVolumeBackupRequest := oci_core.CopyBootVolumeBackupRequest{}

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
		copyBootVolumeBackupRequest.DestinationRegion = &currentRegion

		if bootVolumeBackupId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "boot_volume_backup_id")); ok {
			tmp := bootVolumeBackupId.(string)
			copyBootVolumeBackupRequest.BootVolumeBackupId = &tmp
		}

		if kmsKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_id")); ok {
			tmp := kmsKeyId.(string)
			copyBootVolumeBackupRequest.KmsKeyId = &tmp
		}

	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		copyBootVolumeBackupRequest.DisplayName = &tmp
	}

	response, err := s.SourceRegionClient.CopyBootVolumeBackup(ctx, copyBootVolumeBackupRequest)
	if err != nil {
		return err
	}

	workRequestId := response.OpcWorkRequestId

	s.Res = &response.BootVolumeBackup

	if workRequestId != nil {
		_, err := tfresource.WaitForWorkRequestWithErrorHandlingAndContext(ctx, s.workRequestClient, workRequestId, "bootVolumeBackup", oci_work_requests.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *CoreBootVolumeBackupResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_core.GetBootVolumeBackupRequest{}

	tmp := s.D.Id()
	request.BootVolumeBackupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetBootVolumeBackup(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.BootVolumeBackup
	return nil
}

func (s *CoreBootVolumeBackupResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	//check if there are any fields is set (empty Update request is invalid)
	hasAttributeSet := false

	request := oci_core.UpdateBootVolumeBackupRequest{}

	tmp := s.D.Id()
	if tmp == "" && *s.Res.Id != "" {
		tmp = *s.Res.Id
	}
	request.BootVolumeBackupId = &tmp

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

	if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
		tmp := kmsKeyId.(string)
		request.KmsKeyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateBootVolumeBackup(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.BootVolumeBackup
	return nil
}

func (s *CoreBootVolumeBackupResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_core.DeleteBootVolumeBackupRequest{}

	tmp := s.D.Id()
	request.BootVolumeBackupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteBootVolumeBackup(ctx, request)
	return err
}

func (s *CoreBootVolumeBackupResourceCrud) SetData() error {
	if s.Res.BootVolumeId != nil {
		s.D.Set("boot_volume_id", *s.Res.BootVolumeId)
	}

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

	if s.Res.ImageId != nil {
		s.D.Set("image_id", *s.Res.ImageId)
	}

	if s.Res.KmsKeyId != nil {
		s.D.Set("kms_key_id", *s.Res.KmsKeyId)
	}

	if s.Res.SizeInGBs != nil {
		s.D.Set("size_in_gbs", strconv.FormatInt(*s.Res.SizeInGBs, 10))
	}

	if s.Res.SourceBootVolumeBackupId != nil {
		s.D.Set("source_boot_volume_backup_id", *s.Res.SourceBootVolumeBackupId)
	}

	s.D.Set("source_type", s.Res.SourceType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
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

	return nil
}

func (s *CoreBootVolumeBackupResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeBootVolumeBackupCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.BootVolumeBackupId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ChangeBootVolumeBackupCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
