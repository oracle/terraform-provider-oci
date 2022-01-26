// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	oci_core "github.com/oracle/oci-go-sdk/v56/core"
	oci_work_requests "github.com/oracle/oci-go-sdk/v56/workrequests"
)

func CoreImageResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("2h"),
			Update: tfresource.GetTimeoutDuration("2h"),
			Delete: tfresource.GetTimeoutDuration("2h"),
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
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
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
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
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
						"operating_system": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"operating_system_version": {
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
			"agent_features": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"is_management_supported": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_monitoring_supported": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"base_image_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"billable_size_in_gbs": {
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
			"listing_type": {
				Type:     schema.TypeString,
				Computed: true,
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
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readCoreImage(d *schema.ResourceData, m interface{}) error {
	sync := &CoreImageResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

func updateCoreImage(d *schema.ResourceData, m interface{}) error {
	sync := &CoreImageResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreImage(d *schema.ResourceData, m interface{}) error {
	sync := &CoreImageResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type CoreImageResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.ComputeClient
	Res                    *oci_core.Image
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
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
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateImage(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.Res = &response.Image

	if workId != nil {
		identifier, err := tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "image", oci_work_requests.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
			return err
		}
	}

	return s.Get()
}

func (s *CoreImageResourceCrud) Get() error {
	request := oci_core.GetImageRequest{}

	tmp := s.D.Id()
	request.ImageId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetImage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Image
	return nil
}

func (s *CoreImageResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdateImageRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.ImageId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteImage(context.Background(), request)
	return err
}

func (s *CoreImageResourceCrud) SetData() error {
	if s.Res.AgentFeatures != nil {
		s.D.Set("agent_features", []interface{}{InstanceAgentFeaturesToMap(s.Res.AgentFeatures)})
	} else {
		s.D.Set("agent_features", nil)
	}

	if s.Res.BaseImageId != nil {
		s.D.Set("base_image_id", *s.Res.BaseImageId)
	}

	if s.Res.BillableSizeInGBs != nil {
		s.D.Set("billable_size_in_gbs", strconv.FormatInt(*s.Res.BillableSizeInGBs, 10))
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreateImageAllowed != nil {
		s.D.Set("create_image_allowed", *s.Res.CreateImageAllowed)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
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

	s.D.Set("listing_type", s.Res.ListingType)

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
		if operatingSystem, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operating_system")); ok {
			tmp := operatingSystem.(string)
			details.OperatingSystem = &tmp
		}
		if operatingSystemVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operating_system_version")); ok {
			tmp := operatingSystemVersion.(string)
			details.OperatingSystemVersion = &tmp
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
		if operatingSystem, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operating_system")); ok {
			tmp := operatingSystem.(string)
			details.OperatingSystem = &tmp
		}
		if operatingSystemVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operating_system_version")); ok {
			tmp := operatingSystemVersion.(string)
			details.OperatingSystemVersion = &tmp
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

func InstanceAgentFeaturesToMap(obj *oci_core.InstanceAgentFeatures) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsManagementSupported != nil {
		result["is_management_supported"] = bool(*obj.IsManagementSupported)
	}

	if obj.IsMonitoringSupported != nil {
		result["is_monitoring_supported"] = bool(*obj.IsMonitoringSupported)
	}

	return result
}

func (s *CoreImageResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeImageCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ImageId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ChangeImageCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
