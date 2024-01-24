// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"
)

func CoreVolumeAttachmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreVolumeAttachment,
		Read:     readCoreVolumeAttachment,
		Delete:   deleteCoreVolumeAttachment,
		Schema: map[string]*schema.Schema{
			// Required
			"attachment_type": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"emulated",
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
			"device": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"encryption_in_transit_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_agent_auto_iscsi_login_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_pv_encryption_in_transit_enabled": {
				Type:     schema.TypeBool,
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
			"is_shareable": {
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
			"chap_secret": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"chap_username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
				// Keep as optional to avoid validation during destroy for the legacy configurations
				Optional:   true,
				Deprecated: tfresource.FieldDeprecatedAndAvoidReferences("compartment_id"),
			},
			"ipv4": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"iqn": { // iSCSI Qualified Name per RFC 3720
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_multipath": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"iscsi_login_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"multipath_devices": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"ipv4": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"iqn": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"port": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"port": {
				Type:     schema.TypeInt,
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
		},
	}
}

func createCoreVolumeAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.CreateResource(d, sync)
}

func readCoreVolumeAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

func deleteCoreVolumeAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CoreVolumeAttachmentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.ComputeClient
	Res                    *oci_core.VolumeAttachment
	DisableNotFoundRetries bool
}

func (s *CoreVolumeAttachmentResourceCrud) ID() string {
	volumeAttachment := *s.Res
	return *volumeAttachment.GetId()
}

func (s *CoreVolumeAttachmentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.VolumeAttachmentLifecycleStateAttaching),
	}
}

func (s *CoreVolumeAttachmentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.VolumeAttachmentLifecycleStateAttached),
	}
}

func (s *CoreVolumeAttachmentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.VolumeAttachmentLifecycleStateDetaching),
	}
}

func (s *CoreVolumeAttachmentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.VolumeAttachmentLifecycleStateDetached),
	}
}

func (s *CoreVolumeAttachmentResourceCrud) Create() error {
	request := oci_core.AttachVolumeRequest{}
	err := s.populateTopLevelPolymorphicAttachVolumeRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.AttachVolume(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VolumeAttachment
	return nil
}

func (s *CoreVolumeAttachmentResourceCrud) Get() error {
	request := oci_core.GetVolumeAttachmentRequest{}

	tmp := s.D.Id()
	request.VolumeAttachmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetVolumeAttachment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VolumeAttachment
	return nil
}

func (s *CoreVolumeAttachmentResourceCrud) Delete() error {
	request := oci_core.DetachVolumeRequest{}

	tmp := s.D.Id()
	request.VolumeAttachmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DetachVolume(context.Background(), request)
	return err
}

func (s *CoreVolumeAttachmentResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_core.EmulatedVolumeAttachment:
		s.D.Set("attachment_type", "emulated")

		if v.AvailabilityDomain != nil {
			s.D.Set("availability_domain", *v.AvailabilityDomain)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.Device != nil {
			s.D.Set("device", *v.Device)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.InstanceId != nil {
			s.D.Set("instance_id", *v.InstanceId)
		}

		if v.IsMultipath != nil {
			s.D.Set("is_multipath", *v.IsMultipath)
		}

		if v.IsPvEncryptionInTransitEnabled != nil {
			s.D.Set("is_pv_encryption_in_transit_enabled", *v.IsPvEncryptionInTransitEnabled)
		}

		if v.IsReadOnly != nil {
			s.D.Set("is_read_only", *v.IsReadOnly)
		}

		s.D.Set("iscsi_login_state", v.IscsiLoginState)

		if v.IsShareable != nil {
			s.D.Set("is_shareable", *v.IsShareable)
		}

		s.D.Set("state", v.LifecycleState)

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.VolumeId != nil {
			s.D.Set("volume_id", *v.VolumeId)
		}
	case oci_core.IScsiVolumeAttachment:
		s.D.Set("attachment_type", "iscsi")

		if v.ChapSecret != nil {
			s.D.Set("chap_secret", *v.ChapSecret)
		}

		if v.ChapUsername != nil {
			s.D.Set("chap_username", *v.ChapUsername)
		}

		s.D.Set("encryption_in_transit_type", v.EncryptionInTransitType)

		if v.Ipv4 != nil {
			s.D.Set("ipv4", *v.Ipv4)
		}

		if v.Iqn != nil {
			s.D.Set("iqn", *v.Iqn)
		}

		if v.IsAgentAutoIscsiLoginEnabled != nil {
			s.D.Set("is_agent_auto_iscsi_login_enabled", *v.IsAgentAutoIscsiLoginEnabled)
		}

		multipathDevices := []interface{}{}
		for _, item := range v.MultipathDevices {
			multipathDevices = append(multipathDevices, MultipathDeviceToMap(item))
		}
		s.D.Set("multipath_devices", multipathDevices)

		if v.Port != nil {
			s.D.Set("port", *v.Port)
		}

		if v.AvailabilityDomain != nil {
			s.D.Set("availability_domain", *v.AvailabilityDomain)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.Device != nil {
			s.D.Set("device", *v.Device)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.InstanceId != nil {
			s.D.Set("instance_id", *v.InstanceId)
		}

		if v.IsMultipath != nil {
			s.D.Set("is_multipath", *v.IsMultipath)
		}

		if v.IsPvEncryptionInTransitEnabled != nil {
			s.D.Set("is_pv_encryption_in_transit_enabled", *v.IsPvEncryptionInTransitEnabled)
		}

		if v.IsReadOnly != nil {
			s.D.Set("is_read_only", *v.IsReadOnly)
		}

		s.D.Set("iscsi_login_state", v.IscsiLoginState)

		if v.IsShareable != nil {
			s.D.Set("is_shareable", *v.IsShareable)
		}

		s.D.Set("state", v.LifecycleState)

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.VolumeId != nil {
			s.D.Set("volume_id", *v.VolumeId)
		}
	case oci_core.ParavirtualizedVolumeAttachment:
		s.D.Set("attachment_type", "paravirtualized")

		if v.AvailabilityDomain != nil {
			s.D.Set("availability_domain", *v.AvailabilityDomain)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.Device != nil {
			s.D.Set("device", *v.Device)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.InstanceId != nil {
			s.D.Set("instance_id", *v.InstanceId)
		}

		if v.IsMultipath != nil {
			s.D.Set("is_multipath", *v.IsMultipath)
		}

		multipathDevices := []interface{}{}
		s.D.Set("multipath_devices", multipathDevices)

		if v.IsPvEncryptionInTransitEnabled != nil {
			s.D.Set("is_pv_encryption_in_transit_enabled", *v.IsPvEncryptionInTransitEnabled)
		}

		if v.IsReadOnly != nil {
			s.D.Set("is_read_only", *v.IsReadOnly)
		}

		s.D.Set("iscsi_login_state", v.IscsiLoginState)

		if v.IsShareable != nil {
			s.D.Set("is_shareable", *v.IsShareable)
		}

		s.D.Set("state", v.LifecycleState)

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.VolumeId != nil {
			s.D.Set("volume_id", *v.VolumeId)
		}
	default:
		log.Printf("[WARN] Received 'attachment_type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func (s *CoreVolumeAttachmentResourceCrud) mapToMultipathDevice(fieldKeyFormat string) (oci_core.MultipathDevice, error) {
	result := oci_core.MultipathDevice{}

	if ipv4, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ipv4")); ok {
		tmp := ipv4.(string)
		result.Ipv4 = &tmp
	}

	if iqn, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "iqn")); ok {
		tmp := iqn.(string)
		result.Iqn = &tmp
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(int)
		result.Port = &tmp
	}

	return result, nil
}

func MultipathDeviceToMap(obj oci_core.MultipathDevice) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ipv4 != nil {
		result["ipv4"] = string(*obj.Ipv4)
	}

	if obj.Iqn != nil {
		result["iqn"] = string(*obj.Iqn)
	}

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	return result
}

func (s *CoreVolumeAttachmentResourceCrud) populateTopLevelPolymorphicAttachVolumeRequest(request *oci_core.AttachVolumeRequest) error {
	//discriminator
	attachmentTypeRaw, ok := s.D.GetOkExists("attachment_type")
	var attachmentType string
	if ok {
		attachmentType = attachmentTypeRaw.(string)
	} else {
		attachmentType = "" // default value
	}
	switch strings.ToLower(attachmentType) {
	case strings.ToLower("emulated"):
		details := oci_core.AttachEmulatedVolumeDetails{}
		if device, ok := s.D.GetOkExists("device"); ok {
			tmp := device.(string)
			details.Device = &tmp
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
		if isShareable, ok := s.D.GetOkExists("is_shareable"); ok {
			tmp := isShareable.(bool)
			details.IsShareable = &tmp
		}
		if volumeId, ok := s.D.GetOkExists("volume_id"); ok {
			tmp := volumeId.(string)
			details.VolumeId = &tmp
		}
		request.AttachVolumeDetails = details
	case strings.ToLower("iscsi"):
		details := oci_core.AttachIScsiVolumeDetails{}
		if encryptionInTransitType, ok := s.D.GetOkExists("encryption_in_transit_type"); ok {
			details.EncryptionInTransitType = oci_core.EncryptionInTransitTypeEnum(encryptionInTransitType.(string))
		}
		if isAgentAutoIscsiLoginEnabled, ok := s.D.GetOkExists("is_agent_auto_iscsi_login_enabled"); ok {
			tmp := isAgentAutoIscsiLoginEnabled.(bool)
			details.IsAgentAutoIscsiLoginEnabled = &tmp
		}
		if useChap, ok := s.D.GetOkExists("use_chap"); ok {
			tmp := useChap.(bool)
			details.UseChap = &tmp
		}
		if device, ok := s.D.GetOkExists("device"); ok {
			tmp := device.(string)
			details.Device = &tmp
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
		if isShareable, ok := s.D.GetOkExists("is_shareable"); ok {
			tmp := isShareable.(bool)
			details.IsShareable = &tmp
		}
		if volumeId, ok := s.D.GetOkExists("volume_id"); ok {
			tmp := volumeId.(string)
			details.VolumeId = &tmp
		}
		request.AttachVolumeDetails = details
	case strings.ToLower("paravirtualized"):
		details := oci_core.AttachParavirtualizedVolumeDetails{}
		if device, ok := s.D.GetOkExists("device"); ok {
			tmp := device.(string)
			details.Device = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
			tmp := instanceId.(string)
			details.InstanceId = &tmp
		}
		if isPvEncryptionInTransitEnabled, ok := s.D.GetOkExists("is_pv_encryption_in_transit_enabled"); ok {
			tmp := isPvEncryptionInTransitEnabled.(bool)
			details.IsPvEncryptionInTransitEnabled = &tmp
		}
		if isReadOnly, ok := s.D.GetOkExists("is_read_only"); ok {
			tmp := isReadOnly.(bool)
			details.IsReadOnly = &tmp
		}
		if isShareable, ok := s.D.GetOkExists("is_shareable"); ok {
			tmp := isShareable.(bool)
			details.IsShareable = &tmp
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
