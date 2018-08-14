// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

const (
	IScsiVolumeAttachmentDiscriminator           = "iscsi"
	ParavirtualizedVolumeAttachmentDiscriminator = "paravirtualized"
)

func VolumeAttachmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createVolumeAttachment,
		Read:     readVolumeAttachment,
		Delete:   deleteVolumeAttachment,
		Schema: map[string]*schema.Schema{
			// Required
			"attachment_type": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"iscsi",
					"paravirtualized",
				}, true),
			},
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"volume_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_read_only": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"use_chap": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
				// The legacy provider required this, but the API no longer accepts it. Keep as optional
				// to avoid a breaking change. The value will be ignored if defined in the config.
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},

			// VolumeAttachment is a polymorphic type itself. The following are only computed if attachment_type == "iscsi".
			"chap_secret": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"chap_username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv4": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"iqn": { // iSCSI Qualified Name per RFC 3720
				Type:     schema.TypeString,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func createVolumeAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &VolumeAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient

	return CreateResource(d, sync)
}

func readVolumeAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &VolumeAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient

	return ReadResource(sync)
}

func deleteVolumeAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &VolumeAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type VolumeAttachmentResourceCrud struct {
	BaseCrud
	Client                 *oci_core.ComputeClient
	Res                    *oci_core.VolumeAttachment
	DisableNotFoundRetries bool
}

func (s *VolumeAttachmentResourceCrud) ID() string {
	volumeAttachment := *s.Res
	return *volumeAttachment.GetId()
}

func (s *VolumeAttachmentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.VolumeAttachmentLifecycleStateAttaching),
	}
}

func (s *VolumeAttachmentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.VolumeAttachmentLifecycleStateAttached),
	}
}

func (s *VolumeAttachmentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.VolumeAttachmentLifecycleStateDetaching),
	}
}

func (s *VolumeAttachmentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.VolumeAttachmentLifecycleStateDetached),
	}
}

func (s *VolumeAttachmentResourceCrud) Create() error {
	request := oci_core.AttachVolumeRequest{}
	err := s.populateTopLevelPolymorphicAttachVolumeRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.AttachVolume(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VolumeAttachment
	return nil
}

func (s *VolumeAttachmentResourceCrud) Get() error {
	request := oci_core.GetVolumeAttachmentRequest{}

	tmp := s.D.Id()
	request.VolumeAttachmentId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetVolumeAttachment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VolumeAttachment
	return nil
}

func (s *VolumeAttachmentResourceCrud) Delete() error {
	request := oci_core.DetachVolumeRequest{}

	tmp := s.D.Id()
	request.VolumeAttachmentId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DetachVolume(context.Background(), request)
	return err
}

func (s *VolumeAttachmentResourceCrud) SetData() error {
	volumeAttachment := *s.Res

	if availabilityDomain := volumeAttachment.GetAvailabilityDomain(); availabilityDomain != nil {
		s.D.Set("availability_domain", *availabilityDomain)
	}

	if compartmentId := volumeAttachment.GetCompartmentId(); compartmentId != nil {
		s.D.Set("compartment_id", *compartmentId)
	}

	if displayName := volumeAttachment.GetDisplayName(); displayName != nil {
		s.D.Set("display_name", *displayName)
	}

	if instanceId := volumeAttachment.GetInstanceId(); instanceId != nil {
		s.D.Set("instance_id", *instanceId)
	}

	if isReadOnly := volumeAttachment.GetIsReadOnly(); isReadOnly != nil {
		s.D.Set("is_read_only", *isReadOnly)
	}

	s.D.Set("state", volumeAttachment.GetLifecycleState())

	if timeCreated := volumeAttachment.GetTimeCreated(); timeCreated != nil {
		s.D.Set("time_created", timeCreated.String())
	}

	if volumeId := volumeAttachment.GetVolumeId(); volumeId != nil {
		s.D.Set("volume_id", *volumeId)
	}

	switch v := volumeAttachment.(type) {
	case oci_core.IScsiVolumeAttachment:
		s.D.Set("attachment_type", IScsiVolumeAttachmentDiscriminator)

		// IScsiVolumeAttachment-specific fields:
		if v.ChapSecret != nil {
			s.D.Set("chap_secret", *v.ChapSecret)
		}

		if v.ChapUsername != nil {
			s.D.Set("chap_username", *v.ChapUsername)
		}

		if v.Ipv4 != nil {
			s.D.Set("ipv4", *v.Ipv4)
		}

		if v.Iqn != nil {
			s.D.Set("iqn", *v.Iqn)
		}

		if v.Port != nil {
			s.D.Set("port", *v.Port)
		}
	case oci_core.ParavirtualizedVolumeAttachment:
		s.D.Set("attachment_type", ParavirtualizedVolumeAttachmentDiscriminator)
	default:
		log.Printf("[WARN] Received volume attachment of unknown type")
	}

	return nil
}

func (s *VolumeAttachmentResourceCrud) populateTopLevelPolymorphicAttachVolumeRequest(request *oci_core.AttachVolumeRequest) error {
	//discriminator
	attachmentTypeRaw, ok := s.D.GetOkExists("attachment_type")
	var attachmentType string
	if ok {
		attachmentType = attachmentTypeRaw.(string)
	} else {
		attachmentType = "" // default value
	}
	switch strings.ToLower(attachmentType) {
	case strings.ToLower("iscsi"):
		details := oci_core.AttachIScsiVolumeDetails{}
		if useChap, ok := s.D.GetOkExists("use_chap"); ok {
			tmp := useChap.(bool)
			details.UseChap = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
			tmp := instanceId.(string)
			details.InstanceId = &tmp
		}
		if isReadOnly, ok := s.D.GetOkExists("is_read_only"); ok {
			tmp := isReadOnly.(bool)
			details.IsReadOnly = &tmp
		}
		if volumeId, ok := s.D.GetOkExists("volume_id"); ok {
			tmp := volumeId.(string)
			details.VolumeId = &tmp
		}
		request.AttachVolumeDetails = details
	case strings.ToLower("paravirtualized"):
		details := oci_core.AttachParavirtualizedVolumeDetails{}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
			tmp := instanceId.(string)
			details.InstanceId = &tmp
		}
		if isReadOnly, ok := s.D.GetOkExists("is_read_only"); ok {
			tmp := isReadOnly.(bool)
			details.IsReadOnly = &tmp
		}
		if volumeId, ok := s.D.GetOkExists("volume_id"); ok {
			tmp := volumeId.(string)
			details.VolumeId = &tmp
		}
		request.AttachVolumeDetails = details
	default:
		return fmt.Errorf("unknown attachment_type '%v' was specified", attachmentType)
	}
	return nil
}
