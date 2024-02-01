// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package media_services

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_media_services "github.com/oracle/oci-go-sdk/v65/mediaservices"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MediaServicesStreamCdnConfigResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createMediaServicesStreamCdnConfig,
		Read:     readMediaServicesStreamCdnConfig,
		Update:   updateMediaServicesStreamCdnConfig,
		Delete:   deleteMediaServicesStreamCdnConfig,
		Schema: map[string]*schema.Schema{
			// Required
			"config": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"AKAMAI_MANUAL",
								"EDGE",
							}, true),
						},

						// Optional
						"edge_hostname": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"edge_path_prefix": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"edge_token_key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"edge_token_salt": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"is_edge_token_auth": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"origin_auth_secret_key_a": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"origin_auth_secret_key_b": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"origin_auth_secret_key_nonce_a": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"origin_auth_secret_key_nonce_b": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"origin_auth_sign_encryption": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"origin_auth_sign_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"distribution_channel_id": {
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"locks": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"compartment_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"message": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"related_resource_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"time_created": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
						},

						// Computed
					},
				},
			},
			"is_lock_override": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecyle_details": {
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

func createMediaServicesStreamCdnConfig(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesStreamCdnConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.CreateResource(d, sync)
}

func readMediaServicesStreamCdnConfig(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesStreamCdnConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.ReadResource(sync)
}

func updateMediaServicesStreamCdnConfig(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesStreamCdnConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteMediaServicesStreamCdnConfig(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesStreamCdnConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type MediaServicesStreamCdnConfigResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_media_services.MediaServicesClient
	Res                    *oci_media_services.StreamCdnConfig
	DisableNotFoundRetries bool
}

func (s *MediaServicesStreamCdnConfigResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *MediaServicesStreamCdnConfigResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *MediaServicesStreamCdnConfigResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_media_services.StreamCdnConfigLifecycleStateActive),
		string(oci_media_services.StreamCdnConfigLifecycleStateNeedsAttention),
	}
}

func (s *MediaServicesStreamCdnConfigResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *MediaServicesStreamCdnConfigResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_media_services.StreamCdnConfigLifecycleStateDeleted),
	}
}

func (s *MediaServicesStreamCdnConfigResourceCrud) Create() error {
	request := oci_media_services.CreateStreamCdnConfigRequest{}

	if config, ok := s.D.GetOkExists("config"); ok {
		if tmpList := config.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "config", 0)
			tmp, err := s.mapToStreamCdnConfigSection(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Config = tmp
		}
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

	if distributionChannelId, ok := s.D.GetOkExists("distribution_channel_id"); ok {
		tmp := distributionChannelId.(string)
		request.DistributionChannelId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	if locks, ok := s.D.GetOkExists("locks"); ok {
		interfaces := locks.([]interface{})
		tmp := make([]oci_media_services.ResourceLock, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
			converted, err := s.mapToResourceLock(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("locks") {
			request.Locks = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "media_services")

	response, err := s.Client.CreateStreamCdnConfig(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.StreamCdnConfig
	return nil
}

func (s *MediaServicesStreamCdnConfigResourceCrud) Get() error {
	request := oci_media_services.GetStreamCdnConfigRequest{}

	tmp := s.D.Id()
	request.StreamCdnConfigId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "media_services")

	response, err := s.Client.GetStreamCdnConfig(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.StreamCdnConfig
	return nil
}

func (s *MediaServicesStreamCdnConfigResourceCrud) Update() error {
	request := oci_media_services.UpdateStreamCdnConfigRequest{}

	if config, ok := s.D.GetOkExists("config"); ok {
		if tmpList := config.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "config", 0)
			tmp, err := s.mapToStreamCdnConfigSection(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Config = tmp
		}
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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		request.IsLockOverride = &tmp
	}

	tmp := s.D.Id()
	request.StreamCdnConfigId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "media_services")

	response, err := s.Client.UpdateStreamCdnConfig(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.StreamCdnConfig
	return nil
}

func (s *MediaServicesStreamCdnConfigResourceCrud) Delete() error {
	request := oci_media_services.DeleteStreamCdnConfigRequest{}

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		request.IsLockOverride = &tmp
	}

	tmp := s.D.Id()
	request.StreamCdnConfigId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "media_services")

	_, err := s.Client.DeleteStreamCdnConfig(context.Background(), request)
	return err
}

func (s *MediaServicesStreamCdnConfigResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.Config != nil {
		configArray := []interface{}{}
		if configMap := StreamCdnConfigSectionToMap(&s.Res.Config); configMap != nil {
			configArray = append(configArray, configMap)
		}
		s.D.Set("config", configArray)
	} else {
		s.D.Set("config", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DistributionChannelId != nil {
		s.D.Set("distribution_channel_id", *s.Res.DistributionChannelId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	if s.Res.LifecyleDetails != nil {
		s.D.Set("lifecyle_details", *s.Res.LifecyleDetails)
	}

	locks := []interface{}{}
	for _, item := range s.Res.Locks {
		locks = append(locks, ResourceLockToMap(item))
	}
	s.D.Set("locks", locks)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *MediaServicesStreamCdnConfigResourceCrud) mapToResourceLock(fieldKeyFormat string) (oci_media_services.ResourceLock, error) {
	result := oci_media_services.ResourceLock{}

	if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
		tmp := compartmentId.(string)
		result.CompartmentId = &tmp
	}

	if message, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "message")); ok {
		tmp := message.(string)
		result.Message = &tmp
	}

	if relatedResourceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "related_resource_id")); ok {
		tmp := relatedResourceId.(string)
		result.RelatedResourceId = &tmp
	}

	if timeCreated, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_created")); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreated.(string))
		if err != nil {
			return result, err
		}
		result.TimeCreated = &oci_common.SDKTime{Time: tmp}
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_media_services.ResourceLockTypeEnum(type_.(string))
	}

	return result, nil
}

func (s *MediaServicesStreamCdnConfigResourceCrud) mapToStreamCdnConfigSection(fieldKeyFormat string) (oci_media_services.StreamCdnConfigSection, error) {
	var baseObject oci_media_services.StreamCdnConfigSection
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("AKAMAI_MANUAL"):
		details := oci_media_services.AkamaiManualStreamCdnConfig{}
		if edgeHostname, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "edge_hostname")); ok {
			tmp := edgeHostname.(string)
			details.EdgeHostname = &tmp
		}
		if edgePathPrefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "edge_path_prefix")); ok {
			tmp := edgePathPrefix.(string)
			details.EdgePathPrefix = &tmp
		}
		if edgeTokenKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "edge_token_key")); ok {
			tmp := edgeTokenKey.(string)
			details.EdgeTokenKey = &tmp
		}
		if edgeTokenSalt, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "edge_token_salt")); ok {
			tmp := edgeTokenSalt.(string)
			details.EdgeTokenSalt = &tmp
		}
		if isEdgeTokenAuth, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_edge_token_auth")); ok {
			tmp := isEdgeTokenAuth.(bool)
			details.IsEdgeTokenAuth = &tmp
		}
		if originAuthSecretKeyA, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "origin_auth_secret_key_a")); ok {
			tmp := originAuthSecretKeyA.(string)
			details.OriginAuthSecretKeyA = &tmp
		}
		if originAuthSecretKeyB, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "origin_auth_secret_key_b")); ok {
			tmp := originAuthSecretKeyB.(string)
			details.OriginAuthSecretKeyB = &tmp
		}
		if originAuthSecretKeyNonceA, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "origin_auth_secret_key_nonce_a")); ok {
			tmp := originAuthSecretKeyNonceA.(string)
			details.OriginAuthSecretKeyNonceA = &tmp
		}
		if originAuthSecretKeyNonceB, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "origin_auth_secret_key_nonce_b")); ok {
			tmp := originAuthSecretKeyNonceB.(string)
			details.OriginAuthSecretKeyNonceB = &tmp
		}
		if originAuthSignEncryption, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "origin_auth_sign_encryption")); ok {
			details.OriginAuthSignEncryption = oci_media_services.AkamaiManualStreamCdnConfigOriginAuthSignEncryptionEnum(originAuthSignEncryption.(string))
		}
		if originAuthSignType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "origin_auth_sign_type")); ok {
			details.OriginAuthSignType = oci_media_services.AkamaiManualStreamCdnConfigOriginAuthSignTypeEnum(originAuthSignType.(string))
		}
		baseObject = details
	case strings.ToLower("EDGE"):
		details := oci_media_services.EdgeStreamCdnConfig{}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func StreamCdnConfigSectionToMap(obj *oci_media_services.StreamCdnConfigSection) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_media_services.AkamaiManualStreamCdnConfig:
		result["type"] = "AKAMAI_MANUAL"

		if v.EdgeHostname != nil {
			result["edge_hostname"] = string(*v.EdgeHostname)
		}

		if v.EdgePathPrefix != nil {
			result["edge_path_prefix"] = string(*v.EdgePathPrefix)
		}

		if v.EdgeTokenKey != nil {
			result["edge_token_key"] = string(*v.EdgeTokenKey)
		}

		if v.EdgeTokenSalt != nil {
			result["edge_token_salt"] = string(*v.EdgeTokenSalt)
		}

		if v.IsEdgeTokenAuth != nil {
			result["is_edge_token_auth"] = bool(*v.IsEdgeTokenAuth)
		}

		if v.OriginAuthSecretKeyA != nil {
			result["origin_auth_secret_key_a"] = string(*v.OriginAuthSecretKeyA)
		}

		if v.OriginAuthSecretKeyB != nil {
			result["origin_auth_secret_key_b"] = string(*v.OriginAuthSecretKeyB)
		}

		if v.OriginAuthSecretKeyNonceA != nil {
			result["origin_auth_secret_key_nonce_a"] = string(*v.OriginAuthSecretKeyNonceA)
		}

		if v.OriginAuthSecretKeyNonceB != nil {
			result["origin_auth_secret_key_nonce_b"] = string(*v.OriginAuthSecretKeyNonceB)
		}

		result["origin_auth_sign_encryption"] = string(v.OriginAuthSignEncryption)

		result["origin_auth_sign_type"] = string(v.OriginAuthSignType)
	case oci_media_services.EdgeStreamCdnConfig:
		result["type"] = "EDGE"
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func StreamCdnConfigSummaryToMap(obj oci_media_services.StreamCdnConfigSummary) map[string]interface{} {
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

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	locks := []interface{}{}
	for _, item := range obj.Locks {
		locks = append(locks, ResourceLockToMap(item))
	}
	result["locks"] = locks

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
