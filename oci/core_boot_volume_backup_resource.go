// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func CoreBootVolumeBackupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createCoreBootVolumeBackup,
		Read:     readCoreBootVolumeBackup,
		Update:   updateCoreBootVolumeBackup,
		Delete:   deleteCoreBootVolumeBackup,
		Schema: map[string]*schema.Schema{
			// Required
			"boot_volume_id": {
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
			"image_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size_in_gbs": {
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

func createCoreBootVolumeBackup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreBootVolumeBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient

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

func readCoreBootVolumeBackup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreBootVolumeBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient

	return ReadResource(sync)
}

func updateCoreBootVolumeBackup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreBootVolumeBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient

	return UpdateResource(d, sync)
}

func deleteCoreBootVolumeBackup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreBootVolumeBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type CoreBootVolumeBackupResourceCrud struct {
	BaseCrud
	Client                 *oci_core.BlockstorageClient
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

func (s *CoreBootVolumeBackupResourceCrud) Create() error {
	request := oci_core.CreateBootVolumeBackupRequest{}

	if bootVolumeId, ok := s.D.GetOkExists("boot_volume_id"); ok {
		tmp := bootVolumeId.(string)
		request.BootVolumeId = &tmp
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
		request.Type = oci_core.CreateBootVolumeBackupDetailsTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateBootVolumeBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.BootVolumeBackup
	return nil
}

func (s *CoreBootVolumeBackupResourceCrud) Get() error {
	request := oci_core.GetBootVolumeBackupRequest{}

	tmp := s.D.Id()
	request.BootVolumeBackupId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetBootVolumeBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.BootVolumeBackup
	return nil
}

func (s *CoreBootVolumeBackupResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdateBootVolumeBackupRequest{}

	tmp := s.D.Id()
	request.BootVolumeBackupId = &tmp

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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateBootVolumeBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.BootVolumeBackup
	return nil
}

func (s *CoreBootVolumeBackupResourceCrud) Delete() error {
	request := oci_core.DeleteBootVolumeBackupRequest{}

	tmp := s.D.Id()
	request.BootVolumeBackupId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteBootVolumeBackup(context.Background(), request)
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
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
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

	if s.Res.SizeInGBs != nil {
		s.D.Set("size_in_gbs", strconv.FormatInt(*s.Res.SizeInGBs, 10))
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
		s.D.Set("unique_size_in_gbs", strconv.FormatInt(*s.Res.UniqueSizeInGBs, 10))
	}

	return nil
}

func (s *CoreBootVolumeBackupResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeBootVolumeBackupCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.BootVolumeBackupId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ChangeBootVolumeBackupCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}
