// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func init() {
	RegisterResource("oci_core_volume_group", CoreVolumeGroupResource())
}

func CoreVolumeGroupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createCoreVolumeGroup,
		Read:     readCoreVolumeGroup,
		Update:   updateCoreVolumeGroup,
		Delete:   deleteCoreVolumeGroup,
		Schema: map[string]*schema.Schema{
			// Required
			"availability_domain": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"source_details": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"volumeGroupBackupId",
								"volumeGroupId",
								"volumeIds",
							}, true),
						},

						// Optional
						"volume_group_backup_id": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"volume_group_id": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"volume_ids": {
							Type:     schema.TypeSet,
							Optional: true,
							ForceNew: true,
							MaxItems: 64,
							MinItems: 0,
							Set:      literalTypeHashCodeForSets,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Computed
					},
				},
			},

			// Optional
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
			"is_hydrated": {
				Type:     schema.TypeBool,
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
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"volume_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func createCoreVolumeGroup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient()

	return CreateResource(d, sync)
}

func readCoreVolumeGroup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient()

	return ReadResource(sync)
}

func updateCoreVolumeGroup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient()

	return UpdateResource(d, sync)
}

func deleteCoreVolumeGroup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type CoreVolumeGroupResourceCrud struct {
	BaseCrud
	Client                 *oci_core.BlockstorageClient
	Res                    *oci_core.VolumeGroup
	DisableNotFoundRetries bool
}

func (s *CoreVolumeGroupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreVolumeGroupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.VolumeGroupLifecycleStateProvisioning),
	}
}

func (s *CoreVolumeGroupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.VolumeGroupLifecycleStateAvailable),
	}
}

func (s *CoreVolumeGroupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.VolumeGroupLifecycleStateTerminating),
	}
}

func (s *CoreVolumeGroupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.VolumeGroupLifecycleStateTerminated),
	}
}

func (s *CoreVolumeGroupResourceCrud) Create() error {
	request := oci_core.CreateVolumeGroupRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
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
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source_details", 0)
			tmp, err := s.mapToVolumeGroupSourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SourceDetails = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateVolumeGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VolumeGroup
	return nil
}

func (s *CoreVolumeGroupResourceCrud) Get() error {
	request := oci_core.GetVolumeGroupRequest{}

	tmp := s.D.Id()
	request.VolumeGroupId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetVolumeGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VolumeGroup
	return nil
}

func (s *CoreVolumeGroupResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdateVolumeGroupRequest{}

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
	request.VolumeGroupId = &tmp

	if volumeIds, ok := s.D.GetOkExists("volume_ids"); ok {
		interfaces := volumeIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("volume_ids") {
			request.VolumeIds = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateVolumeGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VolumeGroup
	return nil
}

func (s *CoreVolumeGroupResourceCrud) Delete() error {
	request := oci_core.DeleteVolumeGroupRequest{}

	tmp := s.D.Id()
	request.VolumeGroupId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteVolumeGroup(context.Background(), request)
	return err
}

func (s *CoreVolumeGroupResourceCrud) SetData() error {
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
		if sourceDetailsMap := VolumeGroupSourceDetailsToMap(&s.Res.SourceDetails, false); sourceDetailsMap != nil {
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

	s.D.Set("volume_ids", s.Res.VolumeIds)

	return nil
}

func (s *CoreVolumeGroupResourceCrud) mapToVolumeGroupSourceDetails(fieldKeyFormat string) (oci_core.VolumeGroupSourceDetails, error) {
	var baseObject oci_core.VolumeGroupSourceDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("volumeGroupBackupId"):
		details := oci_core.VolumeGroupSourceFromVolumeGroupBackupDetails{}
		if volumeGroupBackupId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "volume_group_backup_id")); ok {
			tmp := volumeGroupBackupId.(string)
			details.VolumeGroupBackupId = &tmp
		}
		baseObject = details
	case strings.ToLower("volumeGroupId"):
		details := oci_core.VolumeGroupSourceFromVolumeGroupDetails{}
		if volumeGroupId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "volume_group_id")); ok {
			tmp := volumeGroupId.(string)
			details.VolumeGroupId = &tmp
		}
		baseObject = details
	case strings.ToLower("volumeIds"):
		details := oci_core.VolumeGroupSourceFromVolumesDetails{}
		if volumeIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "volume_ids")); ok {
			set := volumeIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "volume_ids")) {
				details.VolumeIds = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func VolumeGroupSourceDetailsToMap(obj *oci_core.VolumeGroupSourceDetails, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_core.VolumeGroupSourceFromVolumeGroupBackupDetails:
		result["type"] = "volumeGroupBackupId"

		if v.VolumeGroupBackupId != nil {
			result["volume_group_backup_id"] = string(*v.VolumeGroupBackupId)
		}
	case oci_core.VolumeGroupSourceFromVolumeGroupDetails:
		result["type"] = "volumeGroupId"

		if v.VolumeGroupId != nil {
			result["volume_group_id"] = string(*v.VolumeGroupId)
		}
	case oci_core.VolumeGroupSourceFromVolumesDetails:
		result["type"] = "volumeIds"

		volumeIds := []interface{}{}
		for _, item := range v.VolumeIds {
			volumeIds = append(volumeIds, item)
		}
		if datasource {
			result["volume_ids"] = volumeIds
		} else {
			result["volume_ids"] = schema.NewSet(literalTypeHashCodeForSets, volumeIds)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *CoreVolumeGroupResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeVolumeGroupCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.VolumeGroupId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ChangeVolumeGroupCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}
