// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"

	"strings"

	"github.com/hashicorp/terraform/helper/validation"
	oci_core "github.com/oracle/oci-go-sdk/core"
	"log"
)

func BootVolumeResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createBootVolume,
		Read:     readBootVolume,
		Update:   updateBootVolume,
		Delete:   deleteBootVolume,
		Schema: map[string]*schema.Schema{
			// Required
			"availability_domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source_details": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					// Polymorphic type with 2 subtypes. Both subtypes have the exact schema (required type & required id).
					Schema: map[string]*schema.Schema{
						// Required
						"id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"bootVolumeBackup",
								"bootVolume",
							}, true),
						},

						// Optional

						// Computed
					},
				},
			},

			// Optional
			"backup_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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

			// Computed
			"image_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_hydrated": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			// Add it till it is resolved that passing this as input will not cause an error
			"size_in_gbs": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size_in_mbs": {
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
			"volume_group_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createBootVolume(d *schema.ResourceData, m interface{}) error {
	sync := &BootVolumeResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient

	return CreateResource(d, sync)
}

func readBootVolume(d *schema.ResourceData, m interface{}) error {
	sync := &BootVolumeResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient

	return ReadResource(sync)
}

func updateBootVolume(d *schema.ResourceData, m interface{}) error {
	sync := &BootVolumeResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient

	return UpdateResource(d, sync)
}

func deleteBootVolume(d *schema.ResourceData, m interface{}) error {
	sync := &BootVolumeResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type BootVolumeResourceCrud struct {
	BaseCrud
	Client                 *oci_core.BlockstorageClient
	Res                    *oci_core.BootVolume
	DisableNotFoundRetries bool
}

func (s *BootVolumeResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *BootVolumeResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.BootVolumeLifecycleStateProvisioning),
		string(oci_core.BootVolumeLifecycleStateRestoring),
	}
}

func (s *BootVolumeResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.BootVolumeLifecycleStateAvailable),
	}
}

func (s *BootVolumeResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.BootVolumeLifecycleStateTerminating),
	}
}

func (s *BootVolumeResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.BootVolumeLifecycleStateTerminated),
	}
}

func (s *BootVolumeResourceCrud) Create() error {
	request := oci_core.CreateBootVolumeRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if backupPolicyId, ok := s.D.GetOkExists("backup_policy_id"); ok {
		tmp := backupPolicyId.(string)
		request.BackupPolicyId = &tmp
	}

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

	if sourceDetails, ok := s.D.GetOkExists("source_details"); ok {
		if tmpList := sourceDetails.([]interface{}); len(tmpList) > 0 {
			tmp := mapToBootVolumeSourceDetails(tmpList[0].(map[string]interface{}))
			request.SourceDetails = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateBootVolume(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.BootVolume
	return nil
}

func (s *BootVolumeResourceCrud) Get() error {
	request := oci_core.GetBootVolumeRequest{}

	tmp := s.D.Id()
	request.BootVolumeId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetBootVolume(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.BootVolume
	return nil
}

func (s *BootVolumeResourceCrud) Update() error {
	request := oci_core.UpdateBootVolumeRequest{}

	tmp := s.D.Id()
	request.BootVolumeId = &tmp

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

	response, err := s.Client.UpdateBootVolume(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.BootVolume
	return nil
}

func (s *BootVolumeResourceCrud) Delete() error {
	request := oci_core.DeleteBootVolumeRequest{}

	tmp := s.D.Id()
	request.BootVolumeId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteBootVolume(context.Background(), request)
	return err
}

func (s *BootVolumeResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
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

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.ImageId != nil {
		s.D.Set("image_id", *s.Res.ImageId)
	}

	if s.Res.IsHydrated != nil {
		s.D.Set("is_hydrated", *s.Res.IsHydrated)
	}

	if s.Res.SizeInGBs != nil {
		s.D.Set("size_in_gbs", strconv.FormatInt(*s.Res.SizeInGBs, 10))
	}

	if s.Res.SizeInMBs != nil {
		s.D.Set("size_in_mbs", strconv.FormatInt(*s.Res.SizeInMBs, 10))
	}

	if s.Res.SourceDetails != nil {
		var sourceDetailsArray []interface{}
		if sourceDetailsMap := BootVolumeSourceDetailsToMap(&s.Res.SourceDetails); sourceDetailsMap != nil {
			sourceDetailsArray = append(sourceDetailsArray, sourceDetailsMap)
		}
		s.D.Set("source_details", sourceDetailsArray)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VolumeGroupId != nil {
		s.D.Set("volume_group_id", *s.Res.VolumeGroupId)
	}

	return nil
}

func mapToBootVolumeSourceDetails(raw map[string]interface{}) oci_core.BootVolumeSourceDetails {
	//discriminator
	typeRaw, ok := raw["type"]
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}

	var baseObject oci_core.BootVolumeSourceDetails
	switch strings.ToLower(type_) {
	case strings.ToLower("bootVolumeBackup"):
		details := oci_core.BootVolumeSourceFromBootVolumeBackupDetails{}
		if id, ok := raw["id"]; ok {
			tmp := id.(string)
			details.Id = &tmp
		}
		baseObject = details
	case strings.ToLower("bootVolume"):
		details := oci_core.BootVolumeSourceFromBootVolumeDetails{}
		if id, ok := raw["id"]; ok {
			tmp := id.(string)
			details.Id = &tmp
		}
		baseObject = details
	default:
		log.Printf("[WARN] Unknown type '%v' was specified", type_)
	}
	return baseObject
}

func BootVolumeSourceDetailsToMap(obj *oci_core.BootVolumeSourceDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_core.BootVolumeSourceFromBootVolumeBackupDetails:
		result["type"] = "bootVolumeBackup"

		if v.Id != nil {
			result["id"] = string(*v.Id)
		}
	case oci_core.BootVolumeSourceFromBootVolumeDetails:
		result["type"] = "bootVolume"

		if v.Id != nil {
			result["id"] = string(*v.Id)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}
