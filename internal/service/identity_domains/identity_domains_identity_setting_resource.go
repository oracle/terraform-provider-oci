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

func IdentityDomainsIdentitySettingResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityDomainsIdentitySetting,
		Read:     readIdentityDomainsIdentitySetting,
		Update:   updateIdentityDomainsIdentitySetting,
		Delete:   deleteIdentityDomainsIdentitySetting,
		Schema: map[string]*schema.Schema{
			// Required
			"idcs_endpoint": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"identity_setting_id": {
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
			"posix_gid": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"manual_assignment_ends_at": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"manual_assignment_starts_from": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"posix_uid": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"manual_assignment_ends_at": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"manual_assignment_starts_from": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
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
			"emit_locked_message_when_user_is_locked": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"my_profile": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"allow_end_users_to_change_their_password": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"allow_end_users_to_link_their_support_account": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"allow_end_users_to_manage_their_capabilities": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"allow_end_users_to_update_their_security_settings": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"ocid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"primary_email_required": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"resource_type_schema_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"return_inactive_over_locked_message": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
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
			"tokens": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"expires_after": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"user_allowed_to_set_recovery_email": {
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
			"remove_invalid_emails": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"tenancy_ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createIdentityDomainsIdentitySetting(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsIdentitySettingResourceCrud{}
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

func readIdentityDomainsIdentitySetting(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsIdentitySettingResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpointForRead(d, "identitySettings")
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

func updateIdentityDomainsIdentitySetting(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsIdentitySettingResourceCrud{}
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

func deleteIdentityDomainsIdentitySetting(d *schema.ResourceData, m interface{}) error {
	return nil
}

type IdentityDomainsIdentitySettingResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity_domains.IdentityDomainsClient
	Res                    *oci_identity_domains.IdentitySetting
	DisableNotFoundRetries bool
}

func (s *IdentityDomainsIdentitySettingResourceCrud) ID() string {
	return *s.Res.Id
	//return GetIdentitySettingCompositeId(s.D.Get("id").(string))
}

func (s *IdentityDomainsIdentitySettingResourceCrud) Create() error {
	request := oci_identity_domains.PutIdentitySettingRequest{}

	if posixGid, ok := s.D.GetOkExists("posix_gid"); ok {
		if tmpList := posixGid.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "posix_gid", 0)
			tmp, err := s.mapToIdentitySettingsPOSIXGid(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.POSIXGid = &tmp
		}
	}

	if posixUid, ok := s.D.GetOkExists("posix_uid"); ok {
		if tmpList := posixUid.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "posix_uid", 0)
			tmp, err := s.mapToIdentitySettingsPOSIXUid(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.POSIXUid = &tmp
		}
	}

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

	if emitLockedMessageWhenUserIsLocked, ok := s.D.GetOkExists("emit_locked_message_when_user_is_locked"); ok {
		tmp := emitLockedMessageWhenUserIsLocked.(bool)
		request.EmitLockedMessageWhenUserIsLocked = &tmp
	}

	if externalId, ok := s.D.GetOkExists("external_id"); ok {
		tmp := externalId.(string)
		request.ExternalId = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if identitySettingId, ok := s.D.GetOkExists("identity_setting_id"); ok {
		tmp := identitySettingId.(string)
		request.IdentitySettingId = &tmp
	}

	if myProfile, ok := s.D.GetOkExists("my_profile"); ok {
		if tmpList := myProfile.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "my_profile", 0)
			tmp, err := s.mapToIdentitySettingsMyProfile(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MyProfile = &tmp
		}
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if primaryEmailRequired, ok := s.D.GetOkExists("primary_email_required"); ok {
		tmp := primaryEmailRequired.(bool)
		request.PrimaryEmailRequired = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	if returnInactiveOverLockedMessage, ok := s.D.GetOkExists("return_inactive_over_locked_message"); ok {
		tmp := returnInactiveOverLockedMessage.(bool)
		request.ReturnInactiveOverLockedMessage = &tmp
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

	if tokens, ok := s.D.GetOkExists("tokens"); ok {
		interfaces := tokens.([]interface{})
		tmp := make([]oci_identity_domains.IdentitySettingsTokens, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tokens", stateDataIndex)
			converted, err := s.mapToIdentitySettingsTokens(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("tokens") {
			request.Tokens = tmp
		}
	}

	if userAllowedToSetRecoveryEmail, ok := s.D.GetOkExists("user_allowed_to_set_recovery_email"); ok {
		tmp := userAllowedToSetRecoveryEmail.(bool)
		request.UserAllowedToSetRecoveryEmail = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.PutIdentitySetting(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IdentitySetting
	return nil
}

func (s *IdentityDomainsIdentitySettingResourceCrud) Get() error {
	request := oci_identity_domains.GetIdentitySettingRequest{}

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
	request.IdentitySettingId = &tmp

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	identitySettingId, err := parseIdentitySettingCompositeId(s.D.Id())
	if err == nil {
		request.IdentitySettingId = &identitySettingId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.GetIdentitySetting(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IdentitySetting
	return nil
}

func (s *IdentityDomainsIdentitySettingResourceCrud) Update() error {
	request := oci_identity_domains.PutIdentitySettingRequest{}

	if posixGid, ok := s.D.GetOkExists("posix_gid"); ok {
		if tmpList := posixGid.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "posix_gid", 0)
			tmp, err := s.mapToIdentitySettingsPOSIXGid(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.POSIXGid = &tmp
		}
	}

	if posixUid, ok := s.D.GetOkExists("posix_uid"); ok {
		if tmpList := posixUid.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "posix_uid", 0)
			tmp, err := s.mapToIdentitySettingsPOSIXUid(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.POSIXUid = &tmp
		}
	}

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

	if emitLockedMessageWhenUserIsLocked, ok := s.D.GetOkExists("emit_locked_message_when_user_is_locked"); ok {
		tmp := emitLockedMessageWhenUserIsLocked.(bool)
		request.EmitLockedMessageWhenUserIsLocked = &tmp
	}

	if externalId, ok := s.D.GetOkExists("external_id"); ok {
		tmp := externalId.(string)
		request.ExternalId = &tmp
	}

	tmp := s.D.Id()
	request.Id = &tmp

	request.IdentitySettingId = &tmp

	if myProfile, ok := s.D.GetOkExists("my_profile"); ok {
		if tmpList := myProfile.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "my_profile", 0)
			tmp, err := s.mapToIdentitySettingsMyProfile(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MyProfile = &tmp
		}
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if primaryEmailRequired, ok := s.D.GetOkExists("primary_email_required"); ok {
		tmp := primaryEmailRequired.(bool)
		request.PrimaryEmailRequired = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	if returnInactiveOverLockedMessage, ok := s.D.GetOkExists("return_inactive_over_locked_message"); ok {
		tmp := returnInactiveOverLockedMessage.(bool)
		request.ReturnInactiveOverLockedMessage = &tmp
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

	if tokens, ok := s.D.GetOkExists("tokens"); ok {
		interfaces := tokens.([]interface{})
		tmp := make([]oci_identity_domains.IdentitySettingsTokens, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tokens", stateDataIndex)
			converted, err := s.mapToIdentitySettingsTokens(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("tokens") {
			request.Tokens = tmp
		}
	}

	if userAllowedToSetRecoveryEmail, ok := s.D.GetOkExists("user_allowed_to_set_recovery_email"); ok {
		tmp := userAllowedToSetRecoveryEmail.(bool)
		request.UserAllowedToSetRecoveryEmail = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.PutIdentitySetting(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IdentitySetting
	return nil
}

func (s *IdentityDomainsIdentitySettingResourceCrud) SetData() error {

	identitySettingId, err := parseIdentitySettingCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(identitySettingId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.POSIXGid != nil {
		s.D.Set("posix_gid", []interface{}{IdentitySettingsPOSIXGidToMap(s.Res.POSIXGid)})
	} else {
		s.D.Set("posix_gid", nil)
	}

	if s.Res.POSIXUid != nil {
		s.D.Set("posix_uid", []interface{}{IdentitySettingsPOSIXUidToMap(s.Res.POSIXUid)})
	} else {
		s.D.Set("posix_uid", nil)
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

	if s.Res.EmitLockedMessageWhenUserIsLocked != nil {
		s.D.Set("emit_locked_message_when_user_is_locked", *s.Res.EmitLockedMessageWhenUserIsLocked)
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

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	if s.Res.MyProfile != nil {
		s.D.Set("my_profile", []interface{}{IdentitySettingsMyProfileToMap(s.Res.MyProfile)})
	} else {
		s.D.Set("my_profile", nil)
	}

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
	}

	if s.Res.PrimaryEmailRequired != nil {
		s.D.Set("primary_email_required", *s.Res.PrimaryEmailRequired)
	}

	if s.Res.RemoveInvalidEmails != nil {
		s.D.Set("remove_invalid_emails", *s.Res.RemoveInvalidEmails)
	}

	if s.Res.ReturnInactiveOverLockedMessage != nil {
		s.D.Set("return_inactive_over_locked_message", *s.Res.ReturnInactiveOverLockedMessage)
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

	tokens := []interface{}{}
	for _, item := range s.Res.Tokens {
		tokens = append(tokens, IdentitySettingsTokensToMap(item))
	}
	s.D.Set("tokens", tokens)

	if s.Res.UserAllowedToSetRecoveryEmail != nil {
		s.D.Set("user_allowed_to_set_recovery_email", *s.Res.UserAllowedToSetRecoveryEmail)
	}

	return nil
}

//func GetIdentitySettingCompositeId(identitySettingId string) string {
//	id = url.PathEscape(id)
//	idcsEndpoint = url.PathEscape(idcsEndpoint)
//	identitySettingId = url.PathEscape(identitySettingId)
//	compositeId := "idcsEndpoint/" + idcsEndpoint + "/identitySettings/" + identitySettingId
//	return compositeId
//}

func parseIdentitySettingCompositeId(compositeId string) (identitySettingId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("idcsEndpoint/.*/identitySettings/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	//idcsEndpoint, _ = url.PathUnescape(parts[1])
	identitySettingId, _ = url.PathUnescape(parts[3])

	return
}

func IdentitySettingToMap(obj oci_identity_domains.IdentitySetting) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.POSIXGid != nil {
		result["posix_gid"] = []interface{}{IdentitySettingsPOSIXGidToMap(obj.POSIXGid)}
	}

	if obj.POSIXUid != nil {
		result["posix_uid"] = []interface{}{IdentitySettingsPOSIXUidToMap(obj.POSIXUid)}
	}

	if obj.CompartmentOcid != nil {
		result["compartment_ocid"] = string(*obj.CompartmentOcid)
	}

	if obj.DeleteInProgress != nil {
		result["delete_in_progress"] = bool(*obj.DeleteInProgress)
	}

	if obj.DomainOcid != nil {
		result["domain_ocid"] = string(*obj.DomainOcid)
	}

	if obj.EmitLockedMessageWhenUserIsLocked != nil {
		result["emit_locked_message_when_user_is_locked"] = bool(*obj.EmitLockedMessageWhenUserIsLocked)
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

	if obj.Meta != nil {
		result["meta"] = []interface{}{metaToMap(obj.Meta)}
	}

	if obj.MyProfile != nil {
		result["my_profile"] = []interface{}{IdentitySettingsMyProfileToMap(obj.MyProfile)}
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	if obj.PrimaryEmailRequired != nil {
		result["primary_email_required"] = bool(*obj.PrimaryEmailRequired)
	}

	if obj.RemoveInvalidEmails != nil {
		result["remove_invalid_emails"] = bool(*obj.RemoveInvalidEmails)
	}

	if obj.ReturnInactiveOverLockedMessage != nil {
		result["return_inactive_over_locked_message"] = bool(*obj.ReturnInactiveOverLockedMessage)
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

	tokens := []interface{}{}
	for _, item := range obj.Tokens {
		tokens = append(tokens, IdentitySettingsTokensToMap(item))
	}
	result["tokens"] = tokens

	if obj.UserAllowedToSetRecoveryEmail != nil {
		result["user_allowed_to_set_recovery_email"] = bool(*obj.UserAllowedToSetRecoveryEmail)
	}

	return result
}

func (s *IdentityDomainsIdentitySettingResourceCrud) mapToIdentitySettingsMyProfile(fieldKeyFormat string) (oci_identity_domains.IdentitySettingsMyProfile, error) {
	result := oci_identity_domains.IdentitySettingsMyProfile{}

	if allowEndUsersToChangeTheirPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "allow_end_users_to_change_their_password")); ok {
		tmp := allowEndUsersToChangeTheirPassword.(bool)
		result.AllowEndUsersToChangeTheirPassword = &tmp
	}

	if allowEndUsersToLinkTheirSupportAccount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "allow_end_users_to_link_their_support_account")); ok {
		tmp := allowEndUsersToLinkTheirSupportAccount.(bool)
		result.AllowEndUsersToLinkTheirSupportAccount = &tmp
	}

	if allowEndUsersToManageTheirCapabilities, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "allow_end_users_to_manage_their_capabilities")); ok {
		tmp := allowEndUsersToManageTheirCapabilities.(bool)
		result.AllowEndUsersToManageTheirCapabilities = &tmp
	}

	if allowEndUsersToUpdateTheirSecuritySettings, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "allow_end_users_to_update_their_security_settings")); ok {
		tmp := allowEndUsersToUpdateTheirSecuritySettings.(bool)
		result.AllowEndUsersToUpdateTheirSecuritySettings = &tmp
	}

	return result, nil
}

func IdentitySettingsMyProfileToMap(obj *oci_identity_domains.IdentitySettingsMyProfile) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AllowEndUsersToChangeTheirPassword != nil {
		result["allow_end_users_to_change_their_password"] = bool(*obj.AllowEndUsersToChangeTheirPassword)
	}

	if obj.AllowEndUsersToLinkTheirSupportAccount != nil {
		result["allow_end_users_to_link_their_support_account"] = bool(*obj.AllowEndUsersToLinkTheirSupportAccount)
	}

	if obj.AllowEndUsersToManageTheirCapabilities != nil {
		result["allow_end_users_to_manage_their_capabilities"] = bool(*obj.AllowEndUsersToManageTheirCapabilities)
	}

	if obj.AllowEndUsersToUpdateTheirSecuritySettings != nil {
		result["allow_end_users_to_update_their_security_settings"] = bool(*obj.AllowEndUsersToUpdateTheirSecuritySettings)
	}

	return result
}

func (s *IdentityDomainsIdentitySettingResourceCrud) mapToIdentitySettingsPOSIXGid(fieldKeyFormat string) (oci_identity_domains.IdentitySettingsPosixGid, error) {
	result := oci_identity_domains.IdentitySettingsPosixGid{}

	if manualAssignmentEndsAt, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "manual_assignment_ends_at")); ok {
		tmp := manualAssignmentEndsAt.(int)
		result.ManualAssignmentEndsAt = &tmp
	}

	if manualAssignmentStartsFrom, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "manual_assignment_starts_from")); ok {
		tmp := manualAssignmentStartsFrom.(int)
		result.ManualAssignmentStartsFrom = &tmp
	}

	return result, nil
}

func IdentitySettingsPOSIXGidToMap(obj *oci_identity_domains.IdentitySettingsPosixGid) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ManualAssignmentEndsAt != nil {
		result["manual_assignment_ends_at"] = int(*obj.ManualAssignmentEndsAt)
	}

	if obj.ManualAssignmentStartsFrom != nil {
		result["manual_assignment_starts_from"] = int(*obj.ManualAssignmentStartsFrom)
	}

	return result
}

func (s *IdentityDomainsIdentitySettingResourceCrud) mapToIdentitySettingsPOSIXUid(fieldKeyFormat string) (oci_identity_domains.IdentitySettingsPosixUid, error) {
	result := oci_identity_domains.IdentitySettingsPosixUid{}

	if manualAssignmentEndsAt, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "manual_assignment_ends_at")); ok {
		tmp := manualAssignmentEndsAt.(int)
		result.ManualAssignmentEndsAt = &tmp
	}

	if manualAssignmentStartsFrom, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "manual_assignment_starts_from")); ok {
		tmp := manualAssignmentStartsFrom.(int)
		result.ManualAssignmentStartsFrom = &tmp
	}

	return result, nil
}

func IdentitySettingsPOSIXUidToMap(obj *oci_identity_domains.IdentitySettingsPosixUid) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ManualAssignmentEndsAt != nil {
		result["manual_assignment_ends_at"] = int(*obj.ManualAssignmentEndsAt)
	}

	if obj.ManualAssignmentStartsFrom != nil {
		result["manual_assignment_starts_from"] = int(*obj.ManualAssignmentStartsFrom)
	}

	return result
}

func (s *IdentityDomainsIdentitySettingResourceCrud) mapToIdentitySettingsTokens(fieldKeyFormat string) (oci_identity_domains.IdentitySettingsTokens, error) {
	result := oci_identity_domains.IdentitySettingsTokens{}

	if expiresAfter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "expires_after")); ok {
		tmp := expiresAfter.(int)
		result.ExpiresAfter = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_identity_domains.IdentitySettingsTokensTypeEnum(type_.(string))
	}

	return result, nil
}

func IdentitySettingsTokensToMap(obj oci_identity_domains.IdentitySettingsTokens) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ExpiresAfter != nil {
		result["expires_after"] = int(*obj.ExpiresAfter)
	}

	result["type"] = string(obj.Type)

	return result
}

func (s *IdentityDomainsIdentitySettingResourceCrud) mapTotags(fieldKeyFormat string) (oci_identity_domains.Tags, error) {
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
