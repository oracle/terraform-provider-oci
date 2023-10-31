// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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

func IdentityDomainsOAuth2ClientCredentialResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityDomainsOAuth2ClientCredential,
		Read:     readIdentityDomainsOAuth2ClientCredential,
		Delete:   deleteIdentityDomainsOAuth2ClientCredential,
		Schema: map[string]*schema.Schema{
			// Required
			"idcs_endpoint": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"schemas": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"scopes": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"audience": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"scope": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},

			// Optional
			"attribute_sets": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"attributes": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"authorization": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"expires_on": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_reset_secret": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ocid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"resource_type_schema_version": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"key": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"urnietfparamsscimschemasoracleidcsextensionself_change_user": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"allow_self_change": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"user": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"ocid": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
						"display": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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
							ForceNew: true,
						},

						// Optional
						"display": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"ocid": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
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
							ForceNew: true,
						},

						// Optional
						"display": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"ocid": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
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
							ForceNew: true,
						},
						"last_modified": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"location": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"resource_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
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

func createIdentityDomainsOAuth2ClientCredential(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsOAuth2ClientCredentialResourceCrud{}
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

func readIdentityDomainsOAuth2ClientCredential(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsOAuth2ClientCredentialResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpointForRead(d, "oAuth2ClientCredentials")
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

func deleteIdentityDomainsOAuth2ClientCredential(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsOAuth2ClientCredentialResourceCrud{}
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
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type IdentityDomainsOAuth2ClientCredentialResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity_domains.IdentityDomainsClient
	Res                    *oci_identity_domains.OAuth2ClientCredential
	DisableNotFoundRetries bool
}

func (s *IdentityDomainsOAuth2ClientCredentialResourceCrud) ID() string {
	return *s.Res.Id
	//return GetOAuth2ClientCredentialCompositeId(s.D.Get("id").(string))
}

func (s *IdentityDomainsOAuth2ClientCredentialResourceCrud) Create() error {
	request := oci_identity_domains.CreateOAuth2ClientCredentialRequest{}

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

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if expiresOn, ok := s.D.GetOkExists("expires_on"); ok {
		tmp := expiresOn.(string)
		request.ExpiresOn = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if isResetSecret, ok := s.D.GetOkExists("is_reset_secret"); ok {
		tmp := isResetSecret.(bool)
		request.IsResetSecret = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
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

	if scopes, ok := s.D.GetOkExists("scopes"); ok {
		interfaces := scopes.([]interface{})
		tmp := make([]oci_identity_domains.OAuth2ClientCredentialScopes, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "scopes", stateDataIndex)
			converted, err := s.mapToOAuth2ClientCredentialScopes(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("scopes") {
			request.Scopes = tmp
		}
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_identity_domains.OAuth2ClientCredentialStatusEnum(status.(string))
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

	if urnietfparamsscimschemasoracleidcsextensionselfChangeUser, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionself_change_user"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionselfChangeUser.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionself_change_user", 0)
			tmp, err := s.mapToExtensionSelfChangeUser(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionSelfChangeUser = &tmp
		}
	}

	if user, ok := s.D.GetOkExists("user"); ok {
		if tmpList := user.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "user", 0)
			tmp, err := s.mapToOAuth2ClientCredentialUser(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.User = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.CreateOAuth2ClientCredential(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OAuth2ClientCredential
	return nil
}

func (s *IdentityDomainsOAuth2ClientCredentialResourceCrud) Get() error {
	request := oci_identity_domains.GetOAuth2ClientCredentialRequest{}

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
	request.OAuth2ClientCredentialId = &tmp

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	oAuth2ClientCredentialId, err := parseOAuth2ClientCredentialCompositeId(s.D.Id())
	if err == nil {
		request.OAuth2ClientCredentialId = &oAuth2ClientCredentialId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.GetOAuth2ClientCredential(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OAuth2ClientCredential
	return nil
}

func (s *IdentityDomainsOAuth2ClientCredentialResourceCrud) Delete() error {
	request := oci_identity_domains.DeleteOAuth2ClientCredentialRequest{}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if forceDelete, ok := s.D.GetOkExists("force_delete"); ok {
		tmp := forceDelete.(bool)
		request.ForceDelete = &tmp
	}

	tmp := s.D.Id()
	request.OAuth2ClientCredentialId = &tmp

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	_, err := s.Client.DeleteOAuth2ClientCredential(context.Background(), request)
	return err
}

func (s *IdentityDomainsOAuth2ClientCredentialResourceCrud) SetData() error {

	oAuth2ClientCredentialId, err := parseOAuth2ClientCredentialCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(oAuth2ClientCredentialId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
	}

	if s.Res.DeleteInProgress != nil {
		s.D.Set("delete_in_progress", *s.Res.DeleteInProgress)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DomainOcid != nil {
		s.D.Set("domain_ocid", *s.Res.DomainOcid)
	}

	if s.Res.ExpiresOn != nil {
		s.D.Set("expires_on", *s.Res.ExpiresOn)
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

	if s.Res.IsResetSecret != nil {
		s.D.Set("is_reset_secret", *s.Res.IsResetSecret)
	}

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
	}

	s.D.Set("schemas", s.Res.Schemas)

	scopes := []interface{}{}
	for _, item := range s.Res.Scopes {
		scopes = append(scopes, OAuth2ClientCredentialScopesToMap(item))
	}
	s.D.Set("scopes", scopes)

	s.D.Set("status", s.Res.Status)

	tags := []interface{}{}
	for _, item := range s.Res.Tags {
		tags = append(tags, tagsToMap(item))
	}
	s.D.Set("tags", tags)

	if s.Res.TenancyOcid != nil {
		s.D.Set("tenancy_ocid", *s.Res.TenancyOcid)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionSelfChangeUser != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionself_change_user", []interface{}{ExtensionSelfChangeUserToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionSelfChangeUser)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionself_change_user", nil)
	}

	if s.Res.User != nil {
		s.D.Set("user", []interface{}{OAuth2ClientCredentialUserToMap(s.Res.User)})
	} else {
		s.D.Set("user", nil)
	}

	return nil
}

//func GetOAuth2ClientCredentialCompositeId(oAuth2ClientCredentialId string) string {
//	id = url.PathEscape(id)
//	idcsEndpoint = url.PathEscape(idcsEndpoint)
//	oAuth2ClientCredentialId = url.PathEscape(oAuth2ClientCredentialId)
//	compositeId := "idcsEndpoint/" + idcsEndpoint + "/oAuth2ClientCredentials/" + oAuth2ClientCredentialId
//	return compositeId
//}

func parseOAuth2ClientCredentialCompositeId(compositeId string) (oAuth2ClientCredentialId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("idcsEndpoint/.*/oAuth2ClientCredentials/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	//idcsEndpoint, _ = url.PathUnescape(parts[1])
	oAuth2ClientCredentialId, _ = url.PathUnescape(parts[3])

	return
}

func (s *IdentityDomainsOAuth2ClientCredentialResourceCrud) mapToExtensionSelfChangeUser(fieldKeyFormat string) (oci_identity_domains.ExtensionSelfChangeUser, error) {
	result := oci_identity_domains.ExtensionSelfChangeUser{}

	if allowSelfChange, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "allow_self_change")); ok {
		tmp := allowSelfChange.(bool)
		result.AllowSelfChange = &tmp
	}

	return result, nil
}

func OAuth2ClientCredentialToMap(obj oci_identity_domains.OAuth2ClientCredential) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentOcid != nil {
		result["compartment_ocid"] = string(*obj.CompartmentOcid)
	}

	if obj.DeleteInProgress != nil {
		result["delete_in_progress"] = bool(*obj.DeleteInProgress)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DomainOcid != nil {
		result["domain_ocid"] = string(*obj.DomainOcid)
	}

	if obj.ExpiresOn != nil {
		result["expires_on"] = string(*obj.ExpiresOn)
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

	if obj.IsResetSecret != nil {
		result["is_reset_secret"] = bool(*obj.IsResetSecret)
	}

	if obj.Meta != nil {
		result["meta"] = []interface{}{metaToMap(obj.Meta)}
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	result["schemas"] = obj.Schemas

	scopes := []interface{}{}
	for _, item := range obj.Scopes {
		scopes = append(scopes, OAuth2ClientCredentialScopesToMap(item))
	}
	result["scopes"] = scopes

	result["status"] = string(obj.Status)

	tags := []interface{}{}
	for _, item := range obj.Tags {
		tags = append(tags, tagsToMap(item))
	}
	result["tags"] = tags

	if obj.TenancyOcid != nil {
		result["tenancy_ocid"] = string(*obj.TenancyOcid)
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionSelfChangeUser != nil {
		result["urnietfparamsscimschemasoracleidcsextensionself_change_user"] = []interface{}{ExtensionSelfChangeUserToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionSelfChangeUser)}
	}

	if obj.User != nil {
		result["user"] = []interface{}{OAuth2ClientCredentialUserToMap(obj.User)}
	}

	return result
}

func (s *IdentityDomainsOAuth2ClientCredentialResourceCrud) mapToOAuth2ClientCredentialScopes(fieldKeyFormat string) (oci_identity_domains.OAuth2ClientCredentialScopes, error) {
	result := oci_identity_domains.OAuth2ClientCredentialScopes{}

	if audience, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "audience")); ok {
		tmp := audience.(string)
		result.Audience = &tmp
	}

	if scope, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scope")); ok {
		tmp := scope.(string)
		result.Scope = &tmp
	}

	return result, nil
}

func OAuth2ClientCredentialScopesToMap(obj oci_identity_domains.OAuth2ClientCredentialScopes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Audience != nil {
		result["audience"] = string(*obj.Audience)
	}

	if obj.Scope != nil {
		result["scope"] = string(*obj.Scope)
	}

	return result
}

func (s *IdentityDomainsOAuth2ClientCredentialResourceCrud) mapToOAuth2ClientCredentialUser(fieldKeyFormat string) (oci_identity_domains.OAuth2ClientCredentialUser, error) {
	result := oci_identity_domains.OAuth2ClientCredentialUser{}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if ocid, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocid")); ok {
		tmp := ocid.(string)
		result.Ocid = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func OAuth2ClientCredentialUserToMap(obj *oci_identity_domains.OAuth2ClientCredentialUser) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsOAuth2ClientCredentialResourceCrud) mapTotags(fieldKeyFormat string) (oci_identity_domains.Tags, error) {
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
