// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

const (
	ImageSourceViaObjectStorageUriDiscriminator   = "objectStorageUri"
	ImageSourceViaObjectStorageTupleDiscriminator = "objectStorageTuple"
)

func ImageResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: &crud.TwoHours,
			Update: &crud.TwoHours,
			Delete: &crud.TwoHours,
		},
		Create: createImage,
		Read:   readImage,
		Update: updateImage,
		Delete: deleteImage,
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
							DiffSuppressFunc: crud.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								ImageSourceViaObjectStorageUriDiscriminator,
								ImageSourceViaObjectStorageTupleDiscriminator,
							}, true),
						},

						// Optional
						"source_image_type": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},

						// ImageSourceViaObjectStorageUriDetails
						"source_uri": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},

						// ImageSourceViaObjectStorageTupleDetails
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
			"id": {
				Type:     schema.TypeString,
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

func createImage(d *schema.ResourceData, m interface{}) error {
	sync := &ImageResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient

	return crud.CreateResource(d, sync)
}

func readImage(d *schema.ResourceData, m interface{}) error {
	sync := &ImageResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient

	return crud.ReadResource(sync)
}

func updateImage(d *schema.ResourceData, m interface{}) error {
	sync := &ImageResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient

	return crud.UpdateResource(d, sync)
}

func deleteImage(d *schema.ResourceData, m interface{}) error {
	sync := &ImageResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type ImageResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_core.ComputeClient
	Res                    *oci_core.Image
	DisableNotFoundRetries bool
}

func (s *ImageResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ImageResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.ImageLifecycleStateProvisioning),
		string(oci_core.ImageLifecycleStateImporting),
	}
}

func (s *ImageResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.ImageLifecycleStateAvailable),
	}
}

func (s *ImageResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.ImageLifecycleStateDisabled),
	}
}

func (s *ImageResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.ImageLifecycleStateDeleted),
	}
}

func (s *ImageResourceCrud) Create() error {
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
			request.ImageSourceDetails = mapToImageSourceDetails(tmpList[0].(map[string]interface{}))
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

func (s *ImageResourceCrud) Get() error {
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

func (s *ImageResourceCrud) Update() error {
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

func (s *ImageResourceCrud) Delete() error {
	request := oci_core.DeleteImageRequest{}

	tmp := s.D.Id()
	request.ImageId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteImage(context.Background(), request)
	return err
}

func (s *ImageResourceCrud) SetData() {
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

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	s.D.Set("launch_mode", s.Res.LaunchMode)

	if s.Res.LaunchOptions != nil {
		s.D.Set("launch_options", []interface{}{LaunchOptionsToMap(s.Res.LaunchOptions)})
	}

	if s.Res.OperatingSystem != nil {
		s.D.Set("operating_system", *s.Res.OperatingSystem)
	}

	if s.Res.OperatingSystemVersion != nil {
		s.D.Set("operating_system_version", *s.Res.OperatingSystemVersion)
	}

	if s.Res.SizeInMBs != nil {
		s.D.Set("size_in_mbs", *s.Res.SizeInMBs)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

}

func mapToImageSourceDetails(raw map[string]interface{}) oci_core.ImageSourceDetails {
	sourceType := raw["source_type"].(string)

	switch strings.ToLower(sourceType) {
	case strings.ToLower(ImageSourceViaObjectStorageUriDiscriminator):
		result := oci_core.ImageSourceViaObjectStorageUriDetails{}
		if sourceImageType, ok := raw["source_image_type"]; ok {
			tmp := sourceImageType.(string)
			if tmp != "" {
				result.SourceImageType = oci_core.ImageSourceDetailsSourceImageTypeEnum(tmp)
			}
		}
		if sourceUri, ok := raw["source_uri"]; ok {
			tmp := sourceUri.(string)
			if tmp != "" {
				result.SourceUri = &tmp
			}
		}
		return result
	case strings.ToLower(ImageSourceViaObjectStorageTupleDiscriminator):
		result := oci_core.ImageSourceViaObjectStorageTupleDetails{}

		if sourceImageType, ok := raw["source_image_type"]; ok {
			tmp := sourceImageType.(string)
			if tmp != "" {
				result.SourceImageType = oci_core.ImageSourceDetailsSourceImageTypeEnum(tmp)
			}
		}

		if bucketName, ok := raw["bucket_name"]; ok {
			tmp := bucketName.(string)
			if tmp != "" {
				result.BucketName = &tmp
			}
		}

		if namespaceName, ok := raw["namespace_name"]; ok {
			tmp := namespaceName.(string)
			if tmp != "" {
				result.NamespaceName = &tmp
			}
		}

		if objectName, ok := raw["object_name"]; ok {
			tmp := objectName.(string)
			if tmp != "" {
				result.ObjectName = &tmp
			}
		}
		return result
	default:
		log.Printf("[WARN] Unknown source_type '%v' was specified", sourceType)
	}

	return nil
}
