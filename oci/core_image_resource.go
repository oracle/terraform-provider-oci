// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func CoreImageResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: &TwoHours,
			Update: &TwoHours,
			Delete: &TwoHours,
		},
		Create: createCoreImage,
		Read:   readCoreImage,
		Update: updateCoreImage,
		Delete: deleteCoreImage,
		Schema: map[string]*schema.Schema{
			// Required
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
			"image_source_details": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"source_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"objectStorageTuple",
								"objectStorageUri",
							}, true),
						},

						// Optional
						"bucket_name": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"namespace_name": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"object_name": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"source_image_type": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"source_uri": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"instance_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"launch_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"base_image_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"create_image_allowed": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"launch_options": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"boot_volume_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"firmware": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_consistent_volume_naming_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_pv_encryption_in_transit_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"network_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"remote_data_volume_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"operating_system": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"operating_system_version": {
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
		},
	}
}

func createCoreImage(d *schema.ResourceData, m interface{}) error {
	sync := &CoreImageResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient

	return CreateResource(d, sync)
}

func readCoreImage(d *schema.ResourceData, m interface{}) error {
	sync := &CoreImageResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient

	return ReadResource(sync)
}

func updateCoreImage(d *schema.ResourceData, m interface{}) error {
	sync := &CoreImageResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient

	return UpdateResource(d, sync)
}

func deleteCoreImage(d *schema.ResourceData, m interface{}) error {
	sync := &CoreImageResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type CoreImageResourceCrud struct {
	BaseCrud
	Client                 *oci_core.ComputeClient
	Res                    *oci_core.Image
	DisableNotFoundRetries bool
}

func (s *CoreImageResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreImageResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.ImageLifecycleStateProvisioning),
		string(oci_core.ImageLifecycleStateImporting),
	}
}

func (s *CoreImageResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.ImageLifecycleStateAvailable),
	}
}

func (s *CoreImageResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.ImageLifecycleStateDisabled),
	}
}

func (s *CoreImageResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.ImageLifecycleStateDeleted),
	}
}

func (s *CoreImageResourceCrud) Create() error {
	request := oci_core.CreateImageRequest{}

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

	if imageSourceDetails, ok := s.D.GetOkExists("image_source_details"); ok {
		if tmpList := imageSourceDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "image_source_details", 0)
			tmp, err := s.mapToImageSourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ImageSourceDetails = tmp
		}
	}

	if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
		tmp := instanceId.(string)
		request.InstanceId = &tmp
	}

	if launchMode, ok := s.D.GetOkExists("launch_mode"); ok {
		request.LaunchMode = oci_core.CreateImageDetailsLaunchModeEnum(launchMode.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateImage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Image
	return nil
}

func (s *CoreImageResourceCrud) Get() error {
	request := oci_core.GetImageRequest{}

	tmp := s.D.Id()
	request.ImageId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetImage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Image
	return nil
}

func (s *CoreImageResourceCrud) Update() error {
	request := oci_core.UpdateImageRequest{}

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
	request.ImageId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateImage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Image
	return nil
}

func (s *CoreImageResourceCrud) Delete() error {
	request := oci_core.DeleteImageRequest{}

	tmp := s.D.Id()
	request.ImageId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteImage(context.Background(), request)
	return err
}

func (s *CoreImageResourceCrud) SetData() error {
	if s.Res.BaseImageId != nil {
		s.D.Set("base_image_id", *s.Res.BaseImageId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreateImageAllowed != nil {
		s.D.Set("create_image_allowed", *s.Res.CreateImageAllowed)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("launch_mode", s.Res.LaunchMode)

	if s.Res.LaunchOptions != nil {
		s.D.Set("launch_options", []interface{}{LaunchOptionsToMap(s.Res.LaunchOptions)})
	} else {
		s.D.Set("launch_options", nil)
	}

	if s.Res.OperatingSystem != nil {
		s.D.Set("operating_system", *s.Res.OperatingSystem)
	}

	if s.Res.OperatingSystemVersion != nil {
		s.D.Set("operating_system_version", *s.Res.OperatingSystemVersion)
	}

	if s.Res.SizeInMBs != nil {
		s.D.Set("size_in_mbs", strconv.FormatInt(*s.Res.SizeInMBs, 10))
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *CoreImageResourceCrud) mapToImageSourceDetails(fieldKeyFormat string) (oci_core.ImageSourceDetails, error) {
	var baseObject oci_core.ImageSourceDetails
	//discriminator
	sourceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_type"))
	var sourceType string
	if ok {
		sourceType = sourceTypeRaw.(string)
	} else {
		sourceType = "" // default value
	}
	switch strings.ToLower(sourceType) {
	case strings.ToLower("objectStorageTuple"):
		details := oci_core.ImageSourceViaObjectStorageTupleDetails{}
		if bucketName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket_name")); ok {
			tmp := bucketName.(string)
			details.BucketName = &tmp
		}
		if namespaceName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace_name")); ok {
			tmp := namespaceName.(string)
			details.NamespaceName = &tmp
		}
		if objectName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_name")); ok {
			tmp := objectName.(string)
			details.ObjectName = &tmp
		}
		if sourceImageType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_image_type")); ok {
			details.SourceImageType = oci_core.ImageSourceDetailsSourceImageTypeEnum(sourceImageType.(string))
		}
		baseObject = details
	case strings.ToLower("objectStorageUri"):
		details := oci_core.ImageSourceViaObjectStorageUriDetails{}
		if sourceUri, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_uri")); ok {
			tmp := sourceUri.(string)
			details.SourceUri = &tmp
		}
		if sourceImageType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_image_type")); ok {
			details.SourceImageType = oci_core.ImageSourceDetailsSourceImageTypeEnum(sourceImageType.(string))
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown source_type '%v' was specified", sourceType)
	}
	return baseObject, nil
}
