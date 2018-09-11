// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func VolumeResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createVolume,
		Read:     readVolume,
		Update:   updateVolume,
		Delete:   deleteVolume,
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
			"size_in_gbs": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				ValidateFunc:     validateInt64TypeString,
				DiffSuppressFunc: int64StringDiffSuppressFunction,
			},
			"size_in_mbs": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				ValidateFunc:     validateInt64TypeString,
				DiffSuppressFunc: int64StringDiffSuppressFunction,
				Deprecated:       FieldDeprecatedForAnother("size_in_mbs", "size_in_gbs"),
			},
			"source_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
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
								"volume",
								"volumeBackup",
							}, true),
						},

						// Optional

						// Computed
					},
				},
			},
			"volume_backup_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"is_hydrated": {
				Type:     schema.TypeBool,
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

func createVolume(d *schema.ResourceData, m interface{}) error {
	sync := &VolumeResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient

	return CreateResource(d, sync)
}

func readVolume(d *schema.ResourceData, m interface{}) error {
	sync := &VolumeResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient

	return ReadResource(sync)
}

func updateVolume(d *schema.ResourceData, m interface{}) error {
	sync := &VolumeResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient

	return UpdateResource(d, sync)
}

func deleteVolume(d *schema.ResourceData, m interface{}) error {
	sync := &VolumeResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type VolumeResourceCrud struct {
	BaseCrud
	Client                 *oci_core.BlockstorageClient
	Res                    *oci_core.Volume
	DisableNotFoundRetries bool
}

func (s *VolumeResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *VolumeResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.VolumeLifecycleStateProvisioning),
		string(oci_core.VolumeLifecycleStateRestoring),
	}
}

func (s *VolumeResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.VolumeLifecycleStateAvailable),
	}
}

func (s *VolumeResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.VolumeLifecycleStateTerminating),
	}
}

func (s *VolumeResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.VolumeLifecycleStateTerminated),
	}
}

func (s *VolumeResourceCrud) Create() error {
	request := oci_core.CreateVolumeRequest{}

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

	if sizeInGBs, ok := s.D.GetOkExists("size_in_gbs"); ok {
		tmp := sizeInGBs.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert sizeInGBs string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.SizeInGBs = &tmpInt64
	}

	if sizeInMBs, ok := s.D.GetOkExists("size_in_mbs"); ok {
		tmp := sizeInMBs.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert sizeInMBs string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.SizeInMBs = &tmpInt64
	}

	if request.SizeInMBs != nil && request.SizeInGBs != nil &&
		*request.SizeInMBs > 0 && *request.SizeInGBs > 0 {
		return fmt.Errorf("both size in Megabytes and Gigabytes cannot be set. Specify one or the other, or leave both undefined to use the default size")
	}

	if sourceDetails, ok := s.D.GetOkExists("source_details"); ok {
		if tmpList := sourceDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source_details", 0)
			tmp, err := s.mapToVolumeSourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SourceDetails = tmp
		}
	}

	if volumeBackupId, ok := s.D.GetOkExists("volume_backup_id"); ok {
		tmp := volumeBackupId.(string)
		request.VolumeBackupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateVolume(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Volume
	return nil
}

func (s *VolumeResourceCrud) Get() error {
	request := oci_core.GetVolumeRequest{}

	tmp := s.D.Id()
	request.VolumeId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetVolume(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Volume
	return nil
}

func (s *VolumeResourceCrud) Update() error {
	request := oci_core.UpdateVolumeRequest{}

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
	request.VolumeId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateVolume(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Volume
	return nil
}

func (s *VolumeResourceCrud) Delete() error {
	request := oci_core.DeleteVolumeRequest{}

	tmp := s.D.Id()
	request.VolumeId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteVolume(context.Background(), request)
	return err
}

func (s *VolumeResourceCrud) SetData() error {
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
		sourceDetailsArray := []interface{}{}
		if sourceDetailsMap := VolumeSourceDetailsToMap(&s.Res.SourceDetails); sourceDetailsMap != nil {
			sourceDetailsArray = append(sourceDetailsArray, sourceDetailsMap)
		}
		s.D.Set("source_details", sourceDetailsArray)
	} else {
		s.D.Set("source_details", nil)
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

func (s *VolumeResourceCrud) mapToVolumeSourceDetails(fieldKeyFormat string) (oci_core.VolumeSourceDetails, error) {
	var baseObject oci_core.VolumeSourceDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("volume"):
		details := oci_core.VolumeSourceFromVolumeDetails{}
		if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
			tmp := id.(string)
			details.Id = &tmp
		}
		baseObject = details
	case strings.ToLower("volumeBackup"):
		details := oci_core.VolumeSourceFromVolumeBackupDetails{}
		if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
			tmp := id.(string)
			details.Id = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func VolumeSourceDetailsToMap(obj *oci_core.VolumeSourceDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_core.VolumeSourceFromVolumeDetails:
		result["type"] = "volume"

		if v.Id != nil {
			result["id"] = string(*v.Id)
		}
	case oci_core.VolumeSourceFromVolumeBackupDetails:
		result["type"] = "volumeBackup"

		if v.Id != nil {
			result["id"] = string(*v.Id)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}
