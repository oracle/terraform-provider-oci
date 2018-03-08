// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

const IScsiVolumeAttachmentDiscriminator = "iscsi"

func VolumeAttachmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createVolumeAttachment,
		Read:     readVolumeAttachment,
		Delete:   deleteVolumeAttachment,
		Schema: map[string]*schema.Schema{
			// Required
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"attachment_type": { // => "type" (polymorphic discriminator)
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: crud.EqualIgnoreCaseSuppressDiff,
				ValidateFunc:     validation.StringInSlice([]string{"iscsi"}, true),
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
			"id": {
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

			// VolumeAttachment is a polymorphic type itself. However, it has only one subtype
			// IScsiVolumeAttachment. Hence, the following are only computed if type == "iscsi".
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

	return crud.CreateResource(d, sync)
}

func readVolumeAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &VolumeAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient

	return crud.ReadResource(sync)
}

func deleteVolumeAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &VolumeAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type VolumeAttachmentResourceCrud struct {
	crud.BaseCrud
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

	// VolumeDetails is a polymorphic type, but will always be "iSCSI".
	var details oci_core.AttachIScsiVolumeDetails

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		details.DisplayName = &tmp
	}

	if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
		tmp := instanceId.(string)
		details.InstanceId = &tmp
	}

	if type_, ok := s.D.GetOkExists("attachment_type"); ok {
		tmp := type_.(string)
		// The only supported AttachmentType is "iSCSI". Panic otherwise.
		if strings.ToLower(tmp) != strings.ToLower(IScsiVolumeAttachmentDiscriminator) {
			panic("unable to convert configuration to a VolumeAttachment.")
		}
	}

	if volumeId, ok := s.D.GetOkExists("volume_id"); ok {
		tmp := volumeId.(string)
		details.VolumeId = &tmp
	}

	request.AttachVolumeDetails = details

	response, err := s.Client.AttachVolume(context.Background(), request, getRetryOptions(s.DisableNotFoundRetries, "core")...)
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

	response, err := s.Client.GetVolumeAttachment(context.Background(), request, getRetryOptions(s.DisableNotFoundRetries, "core")...)
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

	_, err := s.Client.DetachVolume(context.Background(), request, getRetryOptions(s.DisableNotFoundRetries, "core")...)
	return err
}

func (s *VolumeAttachmentResourceCrud) SetData() {
	// VolumeDetails is a polymorphic type, but will always be "iSCSI". Panic otherwise.
	volumeAttachment := *s.Res
	iscsiVolumeAttachment, castOk := volumeAttachment.(oci_core.IScsiVolumeAttachment)
	if !castOk {
		panic("unexpected VolumeAttachment type.")
	}

	s.D.Set("attachment_type", IScsiVolumeAttachmentDiscriminator)

	if iscsiVolumeAttachment.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *iscsiVolumeAttachment.AvailabilityDomain)
	}

	if iscsiVolumeAttachment.CompartmentId != nil {
		s.D.Set("compartment_id", *iscsiVolumeAttachment.CompartmentId)
	}

	if iscsiVolumeAttachment.DisplayName != nil {
		s.D.Set("display_name", *iscsiVolumeAttachment.DisplayName)
	}

	if iscsiVolumeAttachment.Id != nil {
		s.D.Set("id", *iscsiVolumeAttachment.Id)
	}

	if iscsiVolumeAttachment.InstanceId != nil {
		s.D.Set("instance_id", *iscsiVolumeAttachment.InstanceId)
	}

	s.D.Set("state", iscsiVolumeAttachment.LifecycleState)

	s.D.Set("time_created", iscsiVolumeAttachment.TimeCreated.String())

	if iscsiVolumeAttachment.VolumeId != nil {
		s.D.Set("volume_id", *iscsiVolumeAttachment.VolumeId)
	}

	// IScsiVolumeAttachment-specific fields:
	if iscsiVolumeAttachment.ChapSecret != nil {
		s.D.Set("chap_secret", *iscsiVolumeAttachment.ChapSecret)
	}

	if iscsiVolumeAttachment.ChapUsername != nil {
		s.D.Set("chap_username", *iscsiVolumeAttachment.ChapUsername)
	}

	if iscsiVolumeAttachment.Ipv4 != nil {
		s.D.Set("ipv4", *iscsiVolumeAttachment.Ipv4)
	}

	if iscsiVolumeAttachment.Iqn != nil {
		s.D.Set("iqn", *iscsiVolumeAttachment.Iqn)
	}

	if iscsiVolumeAttachment.Port != nil {
		s.D.Set("port", *iscsiVolumeAttachment.Port)
	}

}
