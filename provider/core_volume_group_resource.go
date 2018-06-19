// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	"strings"

	"fmt"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

const (
	VolumeGroupSourceDetailsVolumeGroupBackupDiscriminator = "volumeGroupBackup"
	VolumeGroupSourceDetailsVolumesDiscriminator           = "volumeIds"
	VolumeGroupSourceDetailsVolumeGroupDiscriminator       = "volumeGroup"
)

func VolumeGroupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createVolumeGroup,
		Read:     readVolumeGroup,
		Update:   updateVolumeGroup,
		Delete:   deleteVolumeGroup,
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
					// Polymorphic type with 3 subtypes. Individual types have different fields
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: crud.EqualIgnoreCaseSuppressDiff,
						},

						// Optional
						"volume_group_id": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"volume_group_backup_id": {
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
							Set:      schema.HashString,
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
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size_in_mbs": {
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

func createVolumeGroup(d *schema.ResourceData, m interface{}) error {
	sync := &VolumeGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient

	return crud.CreateResource(d, sync)
}

func readVolumeGroup(d *schema.ResourceData, m interface{}) error {
	sync := &VolumeGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient

	return crud.ReadResource(sync)
}

func updateVolumeGroup(d *schema.ResourceData, m interface{}) error {
	sync := &VolumeGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient

	return crud.UpdateResource(d, sync)
}

func deleteVolumeGroup(d *schema.ResourceData, m interface{}) error {
	sync := &VolumeGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type VolumeGroupResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_core.BlockstorageClient
	Res                    *oci_core.VolumeGroup
	DisableNotFoundRetries bool
}

func (s *VolumeGroupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *VolumeGroupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.VolumeGroupLifecycleStateProvisioning),
	}
}

func (s *VolumeGroupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.VolumeGroupLifecycleStateAvailable),
	}
}

func (s *VolumeGroupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.VolumeGroupLifecycleStateTerminating),
	}
}

func (s *VolumeGroupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.VolumeGroupLifecycleStateTerminated),
	}
}

func (s *VolumeGroupResourceCrud) Create() error {
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
		tmp := mapToVolumeGroupSourceDetails(sourceDetails.([]interface{}))
		request.SourceDetails = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateVolumeGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VolumeGroup
	return nil
}

func (s *VolumeGroupResourceCrud) Get() error {
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

func (s *VolumeGroupResourceCrud) Update() error {
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

	request.VolumeIds = []string{}
	if volumeIds, ok := s.D.GetOkExists("volume_ids"); ok {
		interfaces := volumeIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i, toBeConverted := range interfaces {
			tmp[i] = toBeConverted.(string)
		}
		request.VolumeIds = tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateVolumeGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VolumeGroup
	return nil
}

func (s *VolumeGroupResourceCrud) Delete() error {
	request := oci_core.DeleteVolumeGroupRequest{}

	tmp := s.D.Id()
	request.VolumeGroupId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteVolumeGroup(context.Background(), request)
	return err
}

func (s *VolumeGroupResourceCrud) SetData() {
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

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	if s.Res.SizeInMBs != nil {
		s.D.Set("size_in_mbs", *s.Res.SizeInMBs)
	}

	s.D.Set("source_details", VolumeGroupSourceDetailsToMap(s.Res.SourceDetails))

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	s.D.Set("volume_ids", s.Res.VolumeIds)

}

func mapToVolumeGroupSourceDetails(rawList []interface{}) oci_core.VolumeGroupSourceDetails {
	var item oci_core.VolumeGroupSourceDetails

	if len(rawList) > 0 {
		rawItem := rawList[0].(map[string]interface{})

		var sourceType string
		if rawType, ok := rawItem["type"]; ok {
			sourceType = strings.ToLower(rawType.(string))
		}

		switch sourceType {
		case strings.ToLower(VolumeGroupSourceDetailsVolumesDiscriminator):
			volumeIdsSet, assertOk := rawItem["volume_ids"].(*schema.Set)
			if !assertOk {
				return fmt.Errorf("could not assert volume_ids as type schema.Set")
			}
			item = oci_core.VolumeGroupSourceFromVolumesDetails{
				VolumeIds: crud.SetToStrings(volumeIdsSet),
			}
		case strings.ToLower(VolumeGroupSourceDetailsVolumeGroupBackupDiscriminator):
			volumeGroupBackupId := rawItem["volume_group_backup_id"].(string)
			item = oci_core.VolumeGroupSourceFromVolumeGroupBackupDetails{
				VolumeGroupBackupId: &volumeGroupBackupId,
			}
		case strings.ToLower(VolumeGroupSourceDetailsVolumeGroupDiscriminator):
			volumeGroupId := rawItem["volume_group_id"].(string)
			item = oci_core.VolumeGroupSourceFromVolumeGroupDetails{
				VolumeGroupId: &volumeGroupId,
			}
		}
	}

	return item
}

func VolumeGroupSourceDetailsToMap(obj oci_core.VolumeGroupSourceDetails) []interface{} {
	var sourceDetails []interface{}
	var item map[string]interface{}

	if details, ok := obj.(oci_core.VolumeGroupSourceFromVolumesDetails); ok {
		item = map[string]interface{}{
			"type":       VolumeGroupSourceDetailsVolumesDiscriminator,
			"volume_ids": crud.StringsToSet(details.VolumeIds),
		}
	} else if details, ok := obj.(oci_core.VolumeGroupSourceFromVolumeGroupBackupDetails); ok {
		item = map[string]interface{}{
			"type": VolumeGroupSourceDetailsVolumeGroupBackupDiscriminator,
			"volume_group_backup_id": *details.VolumeGroupBackupId,
		}
	} else if details, ok := obj.(oci_core.VolumeGroupSourceFromVolumeGroupDetails); ok {
		item = map[string]interface{}{
			"type":            VolumeGroupSourceDetailsVolumeGroupDiscriminator,
			"volume_group_id": *details.VolumeGroupId,
		}
	}

	if item != nil {
		sourceDetails = append(sourceDetails, item)
	}

	return sourceDetails
}
