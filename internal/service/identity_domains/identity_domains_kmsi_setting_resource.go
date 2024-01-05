// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity_domains

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IdentityDomainsKmsiSettingResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityDomainsKmsiSetting,
		Read:     readIdentityDomainsKmsiSetting,
		Update:   updateIdentityDomainsKmsiSetting,
		Delete:   deleteIdentityDomainsKmsiSetting,
		Schema: map[string]*schema.Schema{
			// Required
			"idcs_endpoint": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"kmsi_setting_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"schemas": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"attribute_sets": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"attributes": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"authorization": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"kmsi_feature_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"kmsi_prompt_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"last_enabled_on": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_used_validity_in_days": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_allowed_sessions": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ocid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_type_schema_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"token_validity_in_days": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tou_prompt_disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// Computed
			"compartment_ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"delete_in_progress": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"domain_ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"idcs_created_by": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"display": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ocid": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"idcs_last_modified_by": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"display": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ocid": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"idcs_last_upgraded_in_release": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"idcs_prevented_operations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"meta": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"created": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"last_modified": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"location": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"resource_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"tenancy_ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createIdentityDomainsKmsiSetting(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsKmsiSettingResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpoint(d)
	if err != nil {
		return err
	}
	client, err := m.(*client.OracleClients).IdentityDomainsClientWithEndpoint(idcsEndpoint)
	if err != nil {
		return err
	}
	sync.Client = client
	return tfresource.CreateResource(d, sync)
}

func readIdentityDomainsKmsiSetting(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsKmsiSettingResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpointForRead(d, "kmsiSettings")
	if err != nil {
		return err
	}
	client, err := m.(*client.OracleClients).IdentityDomainsClientWithEndpoint(idcsEndpoint)
	if err != nil {
		return err
	}
	sync.Client = client
	return tfresource.ReadResource(sync)
}

func updateIdentityDomainsKmsiSetting(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsKmsiSettingResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpoint(d)
	if err != nil {
		return err
	}
	client, err := m.(*client.OracleClients).IdentityDomainsClientWithEndpoint(idcsEndpoint)
	if err != nil {
		return err
	}
	sync.Client = client
	return tfresource.UpdateResource(d, sync)
}

func deleteIdentityDomainsKmsiSetting(d *schema.ResourceData, m interface{}) error {
	return nil
}

type IdentityDomainsKmsiSettingResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity_domains.IdentityDomainsClient
	Res                    *oci_identity_domains.KmsiSetting
	DisableNotFoundRetries bool
}

func (s *IdentityDomainsKmsiSettingResourceCrud) ID() string {
	return *s.Res.Id
	//return GetKmsiSettingCompositeId(s.D.Get("id").(string))
}

func (s *IdentityDomainsKmsiSettingResourceCrud) Create() error {
	request := oci_identity_domains.PutKmsiSettingRequest{}

	if attributeSets, ok := s.D.GetOkExists("attribute_sets"); ok {
		interfaces := attributeSets.([]interface{})
		tmp := make([]oci_identity_domains.AttributeSetsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_identity_domains.AttributeSetsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("attribute_sets") {
			request.AttributeSets = tmp
		}
	}

	if attributes, ok := s.D.GetOkExists("attributes"); ok {
		tmp := attributes.(string)
		request.Attributes = &tmp
	}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if externalId, ok := s.D.GetOkExists("external_id"); ok {
		tmp := externalId.(string)
		request.ExternalId = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if kmsiFeatureEnabled, ok := s.D.GetOkExists("kmsi_feature_enabled"); ok {
		tmp := kmsiFeatureEnabled.(bool)
		request.KmsiFeatureEnabled = &tmp
	}

	if kmsiPromptEnabled, ok := s.D.GetOkExists("kmsi_prompt_enabled"); ok {
		tmp := kmsiPromptEnabled.(bool)
		request.KmsiPromptEnabled = &tmp
	}

	if kmsiSettingId, ok := s.D.GetOkExists("kmsi_setting_id"); ok {
		tmp := kmsiSettingId.(string)
		request.KmsiSettingId = &tmp
	}

	if lastEnabledOn, ok := s.D.GetOkExists("last_enabled_on"); ok {
		tmp := lastEnabledOn.(string)
		request.LastEnabledOn = &tmp
	}

	if lastUsedValidityInDays, ok := s.D.GetOkExists("last_used_validity_in_days"); ok {
		tmp := lastUsedValidityInDays.(int)
		request.LastUsedValidityInDays = &tmp
	}

	if maxAllowedSessions, ok := s.D.GetOkExists("max_allowed_sessions"); ok {
		tmp := maxAllowedSessions.(int)
		request.MaxAllowedSessions = &tmp
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	if schemas, ok := s.D.GetOkExists("schemas"); ok {
		interfaces := schemas.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("schemas") {
			request.Schemas = tmp
		}
	}

	if tags, ok := s.D.GetOkExists("tags"); ok {
		interfaces := tags.([]interface{})
		tmp := make([]oci_identity_domains.Tags, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tags", stateDataIndex)
			converted, err := s.mapTotags(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("tags") {
			request.Tags = tmp
		}
	}

	if tokenValidityInDays, ok := s.D.GetOkExists("token_validity_in_days"); ok {
		tmp := tokenValidityInDays.(int)
		request.TokenValidityInDays = &tmp
	}

	if touPromptDisabled, ok := s.D.GetOkExists("tou_prompt_disabled"); ok {
		tmp := touPromptDisabled.(bool)
		request.TouPromptDisabled = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.PutKmsiSetting(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.KmsiSetting
	return nil
}

func (s *IdentityDomainsKmsiSettingResourceCrud) Get() error {
	request := oci_identity_domains.GetKmsiSettingRequest{}

	if attributeSets, ok := s.D.GetOkExists("attribute_sets"); ok {
		interfaces := attributeSets.([]interface{})
		tmp := make([]oci_identity_domains.AttributeSetsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_identity_domains.AttributeSetsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("attribute_sets") {
			request.AttributeSets = tmp
		}
	}

	if attributes, ok := s.D.GetOkExists("attributes"); ok {
		tmp := attributes.(string)
		request.Attributes = &tmp
	}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	tmp := s.D.Id()
	request.KmsiSettingId = &tmp

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	kmsiSettingId, err := parseKmsiSettingCompositeId(s.D.Id())
	if err == nil {
		request.KmsiSettingId = &kmsiSettingId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.GetKmsiSetting(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.KmsiSetting
	return nil
}

func (s *IdentityDomainsKmsiSettingResourceCrud) Update() error {
	request := oci_identity_domains.PutKmsiSettingRequest{}

	if attributeSets, ok := s.D.GetOkExists("attribute_sets"); ok {
		interfaces := attributeSets.([]interface{})
		tmp := make([]oci_identity_domains.AttributeSetsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_identity_domains.AttributeSetsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("attribute_sets") {
			request.AttributeSets = tmp
		}
	}

	if attributes, ok := s.D.GetOkExists("attributes"); ok {
		tmp := attributes.(string)
		request.Attributes = &tmp
	}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if externalId, ok := s.D.GetOkExists("external_id"); ok {
		tmp := externalId.(string)
		request.ExternalId = &tmp
	}

	tmp := s.D.Id()
	request.Id = &tmp

	if kmsiFeatureEnabled, ok := s.D.GetOkExists("kmsi_feature_enabled"); ok {
		tmp := kmsiFeatureEnabled.(bool)
		request.KmsiFeatureEnabled = &tmp
	}

	if kmsiPromptEnabled, ok := s.D.GetOkExists("kmsi_prompt_enabled"); ok {
		tmp := kmsiPromptEnabled.(bool)
		request.KmsiPromptEnabled = &tmp
	}

	if kmsiSettingId, ok := s.D.GetOkExists("kmsi_setting_id"); ok {
		tmp := kmsiSettingId.(string)
		request.KmsiSettingId = &tmp
	}

	if lastEnabledOn, ok := s.D.GetOkExists("last_enabled_on"); ok {
		tmp := lastEnabledOn.(string)
		request.LastEnabledOn = &tmp
	}

	if lastUsedValidityInDays, ok := s.D.GetOkExists("last_used_validity_in_days"); ok {
		tmp := lastUsedValidityInDays.(int)
		request.LastUsedValidityInDays = &tmp
	}

	if maxAllowedSessions, ok := s.D.GetOkExists("max_allowed_sessions"); ok {
		tmp := maxAllowedSessions.(int)
		request.MaxAllowedSessions = &tmp
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	if schemas, ok := s.D.GetOkExists("schemas"); ok {
		interfaces := schemas.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("schemas") {
			request.Schemas = tmp
		}
	}

	if tags, ok := s.D.GetOkExists("tags"); ok {
		interfaces := tags.([]interface{})
		tmp := make([]oci_identity_domains.Tags, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tags", stateDataIndex)
			converted, err := s.mapTotags(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("tags") {
			request.Tags = tmp
		}
	}

	if tokenValidityInDays, ok := s.D.GetOkExists("token_validity_in_days"); ok {
		tmp := tokenValidityInDays.(int)
		request.TokenValidityInDays = &tmp
	}

	if touPromptDisabled, ok := s.D.GetOkExists("tou_prompt_disabled"); ok {
		tmp := touPromptDisabled.(bool)
		request.TouPromptDisabled = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.PutKmsiSetting(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.KmsiSetting
	return nil
}

func (s *IdentityDomainsKmsiSettingResourceCrud) SetData() error {

	kmsiSettingId, err := parseKmsiSettingCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(kmsiSettingId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
	}

	if s.Res.DeleteInProgress != nil {
		s.D.Set("delete_in_progress", *s.Res.DeleteInProgress)
	}

	if s.Res.DomainOcid != nil {
		s.D.Set("domain_ocid", *s.Res.DomainOcid)
	}

	if s.Res.ExternalId != nil {
		s.D.Set("external_id", *s.Res.ExternalId)
	}

	if s.Res.IdcsCreatedBy != nil {
		s.D.Set("idcs_created_by", []interface{}{idcsCreatedByToMap(s.Res.IdcsCreatedBy)})
	} else {
		s.D.Set("idcs_created_by", nil)
	}

	if s.Res.IdcsLastModifiedBy != nil {
		s.D.Set("idcs_last_modified_by", []interface{}{idcsLastModifiedByToMap(s.Res.IdcsLastModifiedBy)})
	} else {
		s.D.Set("idcs_last_modified_by", nil)
	}

	if s.Res.IdcsLastUpgradedInRelease != nil {
		s.D.Set("idcs_last_upgraded_in_release", *s.Res.IdcsLastUpgradedInRelease)
	}

	s.D.Set("idcs_prevented_operations", s.Res.IdcsPreventedOperations)

	if s.Res.KmsiFeatureEnabled != nil {
		s.D.Set("kmsi_feature_enabled", *s.Res.KmsiFeatureEnabled)
	}

	if s.Res.KmsiPromptEnabled != nil {
		s.D.Set("kmsi_prompt_enabled", *s.Res.KmsiPromptEnabled)
	}

	if s.Res.LastEnabledOn != nil {
		s.D.Set("last_enabled_on", *s.Res.LastEnabledOn)
	}

	if s.Res.LastUsedValidityInDays != nil {
		s.D.Set("last_used_validity_in_days", *s.Res.LastUsedValidityInDays)
	}

	if s.Res.MaxAllowedSessions != nil {
		s.D.Set("max_allowed_sessions", *s.Res.MaxAllowedSessions)
	}

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
	}

	s.D.Set("schemas", s.Res.Schemas)

	tags := []interface{}{}
	for _, item := range s.Res.Tags {
		tags = append(tags, tagsToMap(item))
	}
	s.D.Set("tags", tags)

	if s.Res.TenancyOcid != nil {
		s.D.Set("tenancy_ocid", *s.Res.TenancyOcid)
	}

	if s.Res.TokenValidityInDays != nil {
		s.D.Set("token_validity_in_days", *s.Res.TokenValidityInDays)
	}

	if s.Res.TouPromptDisabled != nil {
		s.D.Set("tou_prompt_disabled", *s.Res.TouPromptDisabled)
	}

	return nil
}

//func GetKmsiSettingCompositeId(kmsiSettingId string) string {
//	id = url.PathEscape(id)
//	idcsEndpoint = url.PathEscape(idcsEndpoint)
//	kmsiSettingId = url.PathEscape(kmsiSettingId)
//	compositeId := "idcsEndpoint/" + idcsEndpoint + "/kmsiSettings/" + kmsiSettingId
//	return compositeId
//}

func parseKmsiSettingCompositeId(compositeId string) (kmsiSettingId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("idcsEndpoint/.*/kmsiSettings/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	//idcsEndpoint, _ = url.PathUnescape(parts[1])
	kmsiSettingId, _ = url.PathUnescape(parts[3])

	return
}

func KmsiSettingToMap(obj oci_identity_domains.KmsiSetting) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentOcid != nil {
		result["compartment_ocid"] = string(*obj.CompartmentOcid)
	}

	if obj.DeleteInProgress != nil {
		result["delete_in_progress"] = bool(*obj.DeleteInProgress)
	}

	if obj.DomainOcid != nil {
		result["domain_ocid"] = string(*obj.DomainOcid)
	}

	if obj.ExternalId != nil {
		result["external_id"] = string(*obj.ExternalId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IdcsCreatedBy != nil {
		result["idcs_created_by"] = []interface{}{idcsCreatedByToMap(obj.IdcsCreatedBy)}
	}

	if obj.IdcsLastModifiedBy != nil {
		result["idcs_last_modified_by"] = []interface{}{idcsLastModifiedByToMap(obj.IdcsLastModifiedBy)}
	}

	if obj.IdcsLastUpgradedInRelease != nil {
		result["idcs_last_upgraded_in_release"] = string(*obj.IdcsLastUpgradedInRelease)
	}

	result["idcs_prevented_operations"] = obj.IdcsPreventedOperations

	if obj.KmsiFeatureEnabled != nil {
		result["kmsi_feature_enabled"] = bool(*obj.KmsiFeatureEnabled)
	}

	if obj.KmsiPromptEnabled != nil {
		result["kmsi_prompt_enabled"] = bool(*obj.KmsiPromptEnabled)
	}

	if obj.LastEnabledOn != nil {
		result["last_enabled_on"] = string(*obj.LastEnabledOn)
	}

	if obj.LastUsedValidityInDays != nil {
		result["last_used_validity_in_days"] = int(*obj.LastUsedValidityInDays)
	}

	if obj.MaxAllowedSessions != nil {
		result["max_allowed_sessions"] = int(*obj.MaxAllowedSessions)
	}

	if obj.Meta != nil {
		result["meta"] = []interface{}{metaToMap(obj.Meta)}
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	result["schemas"] = obj.Schemas

	tags := []interface{}{}
	for _, item := range obj.Tags {
		tags = append(tags, tagsToMap(item))
	}
	result["tags"] = tags

	if obj.TenancyOcid != nil {
		result["tenancy_ocid"] = string(*obj.TenancyOcid)
	}

	if obj.TokenValidityInDays != nil {
		result["token_validity_in_days"] = int(*obj.TokenValidityInDays)
	}

	if obj.TouPromptDisabled != nil {
		result["tou_prompt_disabled"] = bool(*obj.TouPromptDisabled)
	}

	return result
}

func (s *IdentityDomainsKmsiSettingResourceCrud) mapTotags(fieldKeyFormat string) (oci_identity_domains.Tags, error) {
	result := oci_identity_domains.Tags{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}
