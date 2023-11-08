// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package media_services

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_media_services "github.com/oracle/oci-go-sdk/v65/mediaservices"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MediaServicesStreamPackagingConfigResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createMediaServicesStreamPackagingConfig,
		Read:     readMediaServicesStreamPackagingConfig,
		Update:   updateMediaServicesStreamPackagingConfig,
		Delete:   deleteMediaServicesStreamPackagingConfig,
		Schema: map[string]*schema.Schema{
			// Required
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"distribution_channel_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"segment_time_in_seconds": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"stream_packaging_format": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"encryption": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"algorithm": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"AES128",
								"AES",
								"RSA",
								"NONE",
							}, true),
						},

						// Optional
						"kms_key_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createMediaServicesStreamPackagingConfig(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesStreamPackagingConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.CreateResource(d, sync)
}

func readMediaServicesStreamPackagingConfig(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesStreamPackagingConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.ReadResource(sync)
}

func updateMediaServicesStreamPackagingConfig(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesStreamPackagingConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteMediaServicesStreamPackagingConfig(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesStreamPackagingConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type MediaServicesStreamPackagingConfigResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_media_services.MediaServicesClient
	Res                    *oci_media_services.StreamPackagingConfig
	DisableNotFoundRetries bool
}

func (s *MediaServicesStreamPackagingConfigResourceCrud) ID() string {
	streamPackagingConfig := *s.Res
	return *streamPackagingConfig.GetId()
}

func (s *MediaServicesStreamPackagingConfigResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *MediaServicesStreamPackagingConfigResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_media_services.StreamPackagingConfigLifecycleStateActive),
		string(oci_media_services.StreamPackagingConfigLifecycleStateNeedsAttention),
	}
}

func (s *MediaServicesStreamPackagingConfigResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *MediaServicesStreamPackagingConfigResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_media_services.StreamPackagingConfigLifecycleStateDeleted),
	}
}

func (s *MediaServicesStreamPackagingConfigResourceCrud) Create() error {
	request := oci_media_services.CreateStreamPackagingConfigRequest{}

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

	if distributionChannelId, ok := s.D.GetOkExists("distribution_channel_id"); ok {
		tmp := distributionChannelId.(string)
		request.DistributionChannelId = &tmp
	}

	if encryption, ok := s.D.GetOkExists("encryption"); ok {
		if tmpList := encryption.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "encryption", 0)
			tmp, err := s.mapToStreamPackagingConfigEncryption(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Encryption = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if segmentTimeInSeconds, ok := s.D.GetOkExists("segment_time_in_seconds"); ok {
		tmp := segmentTimeInSeconds.(int)
		request.SegmentTimeInSeconds = &tmp
	}

	if streamPackagingFormat, ok := s.D.GetOkExists("stream_packaging_format"); ok {
		request.StreamPackagingFormat = oci_media_services.CreateStreamPackagingConfigDetailsStreamPackagingFormatEnum(streamPackagingFormat.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "media_services")

	response, err := s.Client.CreateStreamPackagingConfig(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.StreamPackagingConfig
	return nil
}

func (s *MediaServicesStreamPackagingConfigResourceCrud) Get() error {
	request := oci_media_services.GetStreamPackagingConfigRequest{}

	tmp := s.D.Id()
	request.StreamPackagingConfigId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "media_services")

	response, err := s.Client.GetStreamPackagingConfig(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.StreamPackagingConfig
	return nil
}

func (s *MediaServicesStreamPackagingConfigResourceCrud) Update() error {
	request := oci_media_services.UpdateStreamPackagingConfigRequest{}

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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.StreamPackagingConfigId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "media_services")

	response, err := s.Client.UpdateStreamPackagingConfig(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.StreamPackagingConfig
	return nil
}

func (s *MediaServicesStreamPackagingConfigResourceCrud) Delete() error {
	request := oci_media_services.DeleteStreamPackagingConfigRequest{}

	tmp := s.D.Id()
	request.StreamPackagingConfigId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "media_services")

	_, err := s.Client.DeleteStreamPackagingConfig(context.Background(), request)
	return err
}

func (s *MediaServicesStreamPackagingConfigResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_media_services.DashStreamPackagingConfig:
		s.D.Set("stream_packaging_format", "DASH")

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DistributionChannelId != nil {
			s.D.Set("distribution_channel_id", *v.DistributionChannelId)
		}

		if v.Encryption != nil {
			encryptionArray := []interface{}{}
			if encryptionMap := StreamPackagingConfigEncryptionToMap(&v.Encryption); encryptionMap != nil {
				encryptionArray = append(encryptionArray, encryptionMap)
			}
			s.D.Set("encryption", encryptionArray)
		} else {
			s.D.Set("encryption", nil)
		}

		s.D.Set("freeform_tags", v.FreeformTags)
		s.D.Set("freeform_tags", v.FreeformTags)

		if v.SegmentTimeInSeconds != nil {
			s.D.Set("segment_time_in_seconds", *v.SegmentTimeInSeconds)
		}

		s.D.Set("state", v.LifecycleState)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}
	case oci_media_services.HlsStreamPackagingConfig:
		s.D.Set("stream_packaging_format", "HLS")

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DistributionChannelId != nil {
			s.D.Set("distribution_channel_id", *v.DistributionChannelId)
		}

		if v.Encryption != nil {
			encryptionArray := []interface{}{}
			if encryptionMap := StreamPackagingConfigEncryptionToMap(&v.Encryption); encryptionMap != nil {
				encryptionArray = append(encryptionArray, encryptionMap)
			}
			s.D.Set("encryption", encryptionArray)
		} else {
			s.D.Set("encryption", nil)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.SegmentTimeInSeconds != nil {
			s.D.Set("segment_time_in_seconds", *v.SegmentTimeInSeconds)
		}

		s.D.Set("state", v.LifecycleState)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}
	default:
		log.Printf("[WARN] Received 'stream_packaging_format' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func (s *MediaServicesStreamPackagingConfigResourceCrud) mapToStreamPackagingConfigEncryption(fieldKeyFormat string) (oci_media_services.StreamPackagingConfigEncryption, error) {
	var baseObject oci_media_services.StreamPackagingConfigEncryption
	//discriminator
	algorithmRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "algorithm"))
	var algorithm string
	if ok {
		algorithm = algorithmRaw.(string)
	} else {
		algorithm = "" // default value
	}
	switch strings.ToLower(algorithm) {
	case strings.ToLower("AES128"):
	case strings.ToLower("AES"):
		details := oci_media_services.StreamPackagingConfigEncryptionAes128{}
		if kmsKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_id")); ok {
			tmp := kmsKeyId.(string)
			details.KmsKeyId = &tmp
		}
		baseObject = details
	case strings.ToLower("NONE"):
		details := oci_media_services.StreamPackagingConfigEncryptionNone{}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown algorithm '%v' was specified", algorithm)
	}
	return baseObject, nil
}

func StreamPackagingConfigEncryptionToMap(obj *oci_media_services.StreamPackagingConfigEncryption) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_media_services.StreamPackagingConfigEncryptionAes128:
		result["algorithm"] = "AES"

		if v.KmsKeyId != nil {
			result["kms_key_id"] = string(*v.KmsKeyId)
		}
	case oci_media_services.StreamPackagingConfigEncryptionNone:
		result["algorithm"] = "NONE"
	default:
		log.Printf("[WARN] Received 'algorithm' of unknown type %v", *obj)
		return nil
	}

	return result
}

func StreamPackagingConfigSummaryToMap(obj oci_media_services.StreamPackagingConfigSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.DistributionChannelId != nil {
		result["distribution_channel_id"] = string(*obj.DistributionChannelId)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
