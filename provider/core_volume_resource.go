// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

const (
	VolumeSourceDetailsVolumeBackupDiscriminator = "volumeBackup"
	VolumeSourceDetailsVolumeDiscriminator       = "volume"
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
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			// @Deprecated 2017: size_in_mbs => size_in_gbs
			"size_in_mbs": {
				Type:       schema.TypeString,
				Optional:   true,
				ForceNew:   true,
				Computed:   true,
				Deprecated: FieldDeprecatedForAnother("size_in_mbs", "size_in_gbs"),
			},
			"source_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
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
		tmp_i, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert sizeInGBs string: %s to an int64", tmp)
		}
		request.SizeInGBs = &tmp_i
	}

	if sizeInMBs, ok := s.D.GetOkExists("size_in_mbs"); ok {
		tmp := sizeInMBs.(string)
		tmp_i, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert sizeInMBs string: %s to an int64", tmp)
		}
		request.SizeInMBs = &tmp_i
	}

	// @Deprecated 2017: size_in_mbs => size_in_gbs
	if request.SizeInMBs != nil && request.SizeInGBs != nil &&
		*request.SizeInMBs > 0 && *request.SizeInGBs > 0 {
		return fmt.Errorf("both size in Megabytes and Gigabytes cannot be set. Specify one or the other, or leave both undefined to use the default size")
	}

	if sourceDetails, ok := s.D.GetOkExists("source_details"); ok {
		tmp := mapToVolumeSourceDetails(sourceDetails.([]interface{}))
		request.SourceDetails = &tmp
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

	s.D.Set("source_details", VolumeSourceDetailsToMap(s.Res.SourceDetails))

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VolumeGroupId != nil {
		s.D.Set("volume_group_id", *s.Res.VolumeGroupId)
	}

	return nil
}

func mapToVolumeSourceDetails(rawList []interface{}) oci_core.VolumeSourceDetails {
	var item oci_core.VolumeSourceDetails

	if len(rawList) > 0 {
		rawItem := rawList[0].(map[string]interface{})

		var sourceType string
		if _type, ok := rawItem["type"]; ok {
			sourceType = strings.ToLower(_type.(string))
		}

		id := rawItem["id"].(string)

		switch sourceType {
		case strings.ToLower(VolumeSourceDetailsVolumeDiscriminator):
			item = oci_core.VolumeSourceFromVolumeDetails{
				Id: &id,
			}
		case strings.ToLower(VolumeSourceDetailsVolumeBackupDiscriminator):
			item = oci_core.VolumeSourceFromVolumeBackupDetails{
				Id: &id,
			}
		}
	}

	return item
}

func VolumeSourceDetailsToMap(obj oci_core.VolumeSourceDetails) []interface{} {
	sourceDetails := []interface{}{}
	var item map[string]interface{}

	if details, ok := obj.(oci_core.VolumeSourceFromVolumeDetails); ok {
		item = map[string]interface{}{
			"type": VolumeSourceDetailsVolumeDiscriminator,
			"id":   *details.Id,
		}
	} else if details, ok := obj.(oci_core.VolumeSourceFromVolumeBackupDetails); ok {
		item = map[string]interface{}{
			"type": VolumeSourceDetailsVolumeBackupDiscriminator,
			"id":   *details.Id,
		}
	}

	if item != nil {
		sourceDetails = append(sourceDetails, item)
	}

	return sourceDetails
}
