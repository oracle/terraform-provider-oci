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

func IdentityDomainsSocialIdentityProviderResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityDomainsSocialIdentityProvider,
		Read:     readIdentityDomainsSocialIdentityProvider,
		Update:   updateIdentityDomainsSocialIdentityProvider,
		Delete:   deleteIdentityDomainsSocialIdentityProvider,
		Schema: map[string]*schema.Schema{
			// Required
			"account_linking_enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"consumer_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"consumer_secret": {
				Type:     schema.TypeString,
				Required: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"idcs_endpoint": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"registration_enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"schemas": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"service_provider_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"show_on_login": {
				Type:     schema.TypeBool,
				Required: true,
			},

			// Optional
			"access_token_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin_scope": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"authorization": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"authz_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_redirect_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"client_credential_in_payload": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"clock_skew_in_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"discovery_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icon_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"id_attribute": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				//ForceNew: true,
			},
			"jit_prov_assigned_groups": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"jit_prov_group_static_list_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ocid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"profile_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirect_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"refresh_token_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"relay_idp_param_mappings": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"relay_param_key": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"relay_param_value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"resource_type_schema_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scope": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"social_jit_provisioning_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
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

func createIdentityDomainsSocialIdentityProvider(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsSocialIdentityProviderResourceCrud{}
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

func readIdentityDomainsSocialIdentityProvider(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsSocialIdentityProviderResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpointForRead(d, "socialIdentityProviders")
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

func updateIdentityDomainsSocialIdentityProvider(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsSocialIdentityProviderResourceCrud{}
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

func deleteIdentityDomainsSocialIdentityProvider(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsSocialIdentityProviderResourceCrud{}
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

type IdentityDomainsSocialIdentityProviderResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity_domains.IdentityDomainsClient
	Res                    *oci_identity_domains.SocialIdentityProvider
	DisableNotFoundRetries bool
}

func (s *IdentityDomainsSocialIdentityProviderResourceCrud) ID() string {
	// return GetSocialIdentityProviderCompositeId(s.D.Get("id").(string))
	return *s.Res.Id
}

func (s *IdentityDomainsSocialIdentityProviderResourceCrud) Create() error {
	request := oci_identity_domains.CreateSocialIdentityProviderRequest{}

	if accessTokenUrl, ok := s.D.GetOkExists("access_token_url"); ok {
		tmp := accessTokenUrl.(string)
		request.AccessTokenUrl = &tmp
	}

	if accountLinkingEnabled, ok := s.D.GetOkExists("account_linking_enabled"); ok {
		tmp := accountLinkingEnabled.(bool)
		request.AccountLinkingEnabled = &tmp
	}

	if adminScope, ok := s.D.GetOkExists("admin_scope"); ok {
		interfaces := adminScope.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("admin_scope") {
			request.AdminScope = tmp
		}
	}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if authzUrl, ok := s.D.GetOkExists("authz_url"); ok {
		tmp := authzUrl.(string)
		request.AuthzUrl = &tmp
	}

	if autoRedirectEnabled, ok := s.D.GetOkExists("auto_redirect_enabled"); ok {
		tmp := autoRedirectEnabled.(bool)
		request.AutoRedirectEnabled = &tmp
	}

	if clientCredentialInPayload, ok := s.D.GetOkExists("client_credential_in_payload"); ok {
		tmp := clientCredentialInPayload.(bool)
		request.ClientCredentialInPayload = &tmp
	}

	if clockSkewInSeconds, ok := s.D.GetOkExists("clock_skew_in_seconds"); ok {
		tmp := clockSkewInSeconds.(int)
		request.ClockSkewInSeconds = &tmp
	}

	if consumerKey, ok := s.D.GetOkExists("consumer_key"); ok {
		tmp := consumerKey.(string)
		request.ConsumerKey = &tmp
	}

	if consumerSecret, ok := s.D.GetOkExists("consumer_secret"); ok {
		tmp := consumerSecret.(string)
		request.ConsumerSecret = &tmp
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if discoveryUrl, ok := s.D.GetOkExists("discovery_url"); ok {
		tmp := discoveryUrl.(string)
		request.DiscoveryUrl = &tmp
	}

	if enabled, ok := s.D.GetOkExists("enabled"); ok {
		tmp := enabled.(bool)
		request.Enabled = &tmp
	}

	if iconUrl, ok := s.D.GetOkExists("icon_url"); ok {
		tmp := iconUrl.(string)
		request.IconUrl = &tmp
	}

	if idAttribute, ok := s.D.GetOkExists("id_attribute"); ok {
		tmp := idAttribute.(string)
		request.IdAttribute = &tmp
	}

	if jitProvAssignedGroups, ok := s.D.GetOkExists("jit_prov_assigned_groups"); ok {
		interfaces := jitProvAssignedGroups.([]interface{})
		tmp := make([]oci_identity_domains.SocialIdentityProviderJitProvAssignedGroups, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "jit_prov_assigned_groups", stateDataIndex)
			converted, err := s.mapToSocialIdentityProviderJitProvAssignedGroups(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("jit_prov_assigned_groups") {
			request.JitProvAssignedGroups = tmp
		}
	}

	if jitProvGroupStaticListEnabled, ok := s.D.GetOkExists("jit_prov_group_static_list_enabled"); ok {
		tmp := jitProvGroupStaticListEnabled.(bool)
		request.JitProvGroupStaticListEnabled = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if profileUrl, ok := s.D.GetOkExists("profile_url"); ok {
		tmp := profileUrl.(string)
		request.ProfileUrl = &tmp
	}

	if redirectUrl, ok := s.D.GetOkExists("redirect_url"); ok {
		tmp := redirectUrl.(string)
		request.RedirectUrl = &tmp
	}

	if refreshTokenUrl, ok := s.D.GetOkExists("refresh_token_url"); ok {
		tmp := refreshTokenUrl.(string)
		request.RefreshTokenUrl = &tmp
	}

	if registrationEnabled, ok := s.D.GetOkExists("registration_enabled"); ok {
		tmp := registrationEnabled.(bool)
		request.RegistrationEnabled = &tmp
	}

	if relayIdpParamMappings, ok := s.D.GetOkExists("relay_idp_param_mappings"); ok {
		interfaces := relayIdpParamMappings.([]interface{})
		tmp := make([]oci_identity_domains.SocialIdentityProviderRelayIdpParamMappings, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "relay_idp_param_mappings", stateDataIndex)
			converted, err := s.mapToSocialIdentityProviderRelayIdpParamMappings(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("relay_idp_param_mappings") {
			request.RelayIdpParamMappings = tmp
		}
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

	if scope, ok := s.D.GetOkExists("scope"); ok {
		interfaces := scope.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("scope") {
			request.Scope = tmp
		}
	}

	if serviceProviderName, ok := s.D.GetOkExists("service_provider_name"); ok {
		tmp := serviceProviderName.(string)
		request.ServiceProviderName = &tmp
	}

	if showOnLogin, ok := s.D.GetOkExists("show_on_login"); ok {
		tmp := showOnLogin.(bool)
		request.ShowOnLogin = &tmp
	}

	if socialJitProvisioningEnabled, ok := s.D.GetOkExists("social_jit_provisioning_enabled"); ok {
		tmp := socialJitProvisioningEnabled.(bool)
		request.SocialJitProvisioningEnabled = &tmp
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_identity_domains.SocialIdentityProviderStatusEnum(status.(string))
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.CreateSocialIdentityProvider(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SocialIdentityProvider
	return nil
}

func (s *IdentityDomainsSocialIdentityProviderResourceCrud) Get() error {
	request := oci_identity_domains.GetSocialIdentityProviderRequest{}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	tmp := s.D.Id()
	request.SocialIdentityProviderId = &tmp

	socialIdentityProviderId, err := parseSocialIdentityProviderCompositeId(s.D.Id())
	if err == nil {
		request.SocialIdentityProviderId = &socialIdentityProviderId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.GetSocialIdentityProvider(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SocialIdentityProvider
	return nil
}

func (s *IdentityDomainsSocialIdentityProviderResourceCrud) Update() error {
	request := oci_identity_domains.PutSocialIdentityProviderRequest{}

	if accessTokenUrl, ok := s.D.GetOkExists("access_token_url"); ok {
		tmp := accessTokenUrl.(string)
		request.AccessTokenUrl = &tmp
	}

	if accountLinkingEnabled, ok := s.D.GetOkExists("account_linking_enabled"); ok {
		tmp := accountLinkingEnabled.(bool)
		request.AccountLinkingEnabled = &tmp
	}

	if adminScope, ok := s.D.GetOkExists("admin_scope"); ok {
		interfaces := adminScope.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("admin_scope") {
			request.AdminScope = tmp
		}
	}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if authzUrl, ok := s.D.GetOkExists("authz_url"); ok {
		tmp := authzUrl.(string)
		request.AuthzUrl = &tmp
	}

	if autoRedirectEnabled, ok := s.D.GetOkExists("auto_redirect_enabled"); ok {
		tmp := autoRedirectEnabled.(bool)
		request.AutoRedirectEnabled = &tmp
	}

	if clientCredentialInPayload, ok := s.D.GetOkExists("client_credential_in_payload"); ok {
		tmp := clientCredentialInPayload.(bool)
		request.ClientCredentialInPayload = &tmp
	}

	if clockSkewInSeconds, ok := s.D.GetOkExists("clock_skew_in_seconds"); ok {
		tmp := clockSkewInSeconds.(int)
		request.ClockSkewInSeconds = &tmp
	}

	if consumerKey, ok := s.D.GetOkExists("consumer_key"); ok {
		tmp := consumerKey.(string)
		request.ConsumerKey = &tmp
	}

	if consumerSecret, ok := s.D.GetOkExists("consumer_secret"); ok {
		tmp := consumerSecret.(string)
		request.ConsumerSecret = &tmp
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if discoveryUrl, ok := s.D.GetOkExists("discovery_url"); ok {
		tmp := discoveryUrl.(string)
		request.DiscoveryUrl = &tmp
	}

	if enabled, ok := s.D.GetOkExists("enabled"); ok {
		tmp := enabled.(bool)
		request.Enabled = &tmp
	}

	if iconUrl, ok := s.D.GetOkExists("icon_url"); ok {
		tmp := iconUrl.(string)
		request.IconUrl = &tmp
	}

	if jitProvAssignedGroups, ok := s.D.GetOkExists("jit_prov_assigned_groups"); ok {
		interfaces := jitProvAssignedGroups.([]interface{})
		tmp := make([]oci_identity_domains.SocialIdentityProviderJitProvAssignedGroups, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "jit_prov_assigned_groups", stateDataIndex)
			converted, err := s.mapToSocialIdentityProviderJitProvAssignedGroups(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("jit_prov_assigned_groups") {
			request.JitProvAssignedGroups = tmp
		}
	}

	if jitProvGroupStaticListEnabled, ok := s.D.GetOkExists("jit_prov_group_static_list_enabled"); ok {
		tmp := jitProvGroupStaticListEnabled.(bool)
		request.JitProvGroupStaticListEnabled = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if profileUrl, ok := s.D.GetOkExists("profile_url"); ok {
		tmp := profileUrl.(string)
		request.ProfileUrl = &tmp
	}

	if redirectUrl, ok := s.D.GetOkExists("redirect_url"); ok {
		tmp := redirectUrl.(string)
		request.RedirectUrl = &tmp
	}

	if refreshTokenUrl, ok := s.D.GetOkExists("refresh_token_url"); ok {
		tmp := refreshTokenUrl.(string)
		request.RefreshTokenUrl = &tmp
	}

	if registrationEnabled, ok := s.D.GetOkExists("registration_enabled"); ok {
		tmp := registrationEnabled.(bool)
		request.RegistrationEnabled = &tmp
	}

	if relayIdpParamMappings, ok := s.D.GetOkExists("relay_idp_param_mappings"); ok {
		interfaces := relayIdpParamMappings.([]interface{})
		tmp := make([]oci_identity_domains.SocialIdentityProviderRelayIdpParamMappings, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "relay_idp_param_mappings", stateDataIndex)
			converted, err := s.mapToSocialIdentityProviderRelayIdpParamMappings(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("relay_idp_param_mappings") {
			request.RelayIdpParamMappings = tmp
		}
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

	if scope, ok := s.D.GetOkExists("scope"); ok {
		interfaces := scope.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("scope") {
			request.Scope = tmp
		}
	}

	if showOnLogin, ok := s.D.GetOkExists("show_on_login"); ok {
		tmp := showOnLogin.(bool)
		request.ShowOnLogin = &tmp
	}
	tmp := s.D.Id()
	request.SocialIdentityProviderId = &tmp

	if socialJitProvisioningEnabled, ok := s.D.GetOkExists("social_jit_provisioning_enabled"); ok {
		tmp := socialJitProvisioningEnabled.(bool)
		request.SocialJitProvisioningEnabled = &tmp
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_identity_domains.SocialIdentityProviderStatusEnum(status.(string))
	}

	if serviceProviderName, ok := s.D.GetOkExists("service_provider_name"); ok {
		tmp := serviceProviderName.(string)
		request.ServiceProviderName = &tmp
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.PutSocialIdentityProvider(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SocialIdentityProvider
	return nil
}

func (s *IdentityDomainsSocialIdentityProviderResourceCrud) Delete() error {
	request := oci_identity_domains.DeleteSocialIdentityProviderRequest{}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if forceDelete, ok := s.D.GetOkExists("force_delete"); ok {
		tmp := forceDelete.(bool)
		request.ForceDelete = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	tmp := s.D.Id()
	request.SocialIdentityProviderId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	_, err := s.Client.DeleteSocialIdentityProvider(context.Background(), request)
	return err
}

func (s *IdentityDomainsSocialIdentityProviderResourceCrud) SetData() error {

	socialIdentityProviderId, err := parseSocialIdentityProviderCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(socialIdentityProviderId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.AccessTokenUrl != nil {
		s.D.Set("access_token_url", *s.Res.AccessTokenUrl)
	}

	if s.Res.AccountLinkingEnabled != nil {
		s.D.Set("account_linking_enabled", *s.Res.AccountLinkingEnabled)
	}

	s.D.Set("admin_scope", s.Res.AdminScope)

	if s.Res.AuthzUrl != nil {
		s.D.Set("authz_url", *s.Res.AuthzUrl)
	}

	if s.Res.AutoRedirectEnabled != nil {
		s.D.Set("auto_redirect_enabled", *s.Res.AutoRedirectEnabled)
	}

	if s.Res.ClientCredentialInPayload != nil {
		s.D.Set("client_credential_in_payload", *s.Res.ClientCredentialInPayload)
	}

	if s.Res.ClockSkewInSeconds != nil {
		s.D.Set("clock_skew_in_seconds", *s.Res.ClockSkewInSeconds)
	}

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
	}

	if s.Res.ConsumerKey != nil {
		s.D.Set("consumer_key", *s.Res.ConsumerKey)
	}

	if s.Res.ConsumerSecret != nil {
		s.D.Set("consumer_secret", *s.Res.ConsumerSecret)
	}

	if s.Res.DeleteInProgress != nil {
		s.D.Set("delete_in_progress", *s.Res.DeleteInProgress)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DiscoveryUrl != nil {
		s.D.Set("discovery_url", *s.Res.DiscoveryUrl)
	}

	if s.Res.DomainOcid != nil {
		s.D.Set("domain_ocid", *s.Res.DomainOcid)
	}

	if s.Res.Enabled != nil {
		s.D.Set("enabled", *s.Res.Enabled)
	}

	if s.Res.IconUrl != nil {
		s.D.Set("icon_url", *s.Res.IconUrl)
	}

	if s.Res.IdAttribute != nil {
		s.D.Set("id_attribute", *s.Res.IdAttribute)
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

	jitProvAssignedGroups := []interface{}{}
	for _, item := range s.Res.JitProvAssignedGroups {
		jitProvAssignedGroups = append(jitProvAssignedGroups, SocialIdentityProviderJitProvAssignedGroupsToMap(item))
	}
	s.D.Set("jit_prov_assigned_groups", jitProvAssignedGroups)

	if s.Res.JitProvGroupStaticListEnabled != nil {
		s.D.Set("jit_prov_group_static_list_enabled", *s.Res.JitProvGroupStaticListEnabled)
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

	if s.Res.ProfileUrl != nil {
		s.D.Set("profile_url", *s.Res.ProfileUrl)
	}

	if s.Res.RedirectUrl != nil {
		s.D.Set("redirect_url", *s.Res.RedirectUrl)
	}

	if s.Res.RefreshTokenUrl != nil {
		s.D.Set("refresh_token_url", *s.Res.RefreshTokenUrl)
	}

	if s.Res.RegistrationEnabled != nil {
		s.D.Set("registration_enabled", *s.Res.RegistrationEnabled)
	}

	relayIdpParamMappings := []interface{}{}
	for _, item := range s.Res.RelayIdpParamMappings {
		relayIdpParamMappings = append(relayIdpParamMappings, SocialIdentityProviderRelayIdpParamMappingsToMap(item))
	}
	s.D.Set("relay_idp_param_mappings", relayIdpParamMappings)

	s.D.Set("schemas", s.Res.Schemas)

	s.D.Set("scope", s.Res.Scope)

	if s.Res.ServiceProviderName != nil {
		s.D.Set("service_provider_name", *s.Res.ServiceProviderName)
	}

	if s.Res.ShowOnLogin != nil {
		s.D.Set("show_on_login", *s.Res.ShowOnLogin)
	}

	if s.Res.SocialJitProvisioningEnabled != nil {
		s.D.Set("social_jit_provisioning_enabled", *s.Res.SocialJitProvisioningEnabled)
	}

	s.D.Set("status", s.Res.Status)

	tags := []interface{}{}
	for _, item := range s.Res.Tags {
		tags = append(tags, tagsToMap(item))
	}
	s.D.Set("tags", tags)

	if s.Res.TenancyOcid != nil {
		s.D.Set("tenancy_ocid", *s.Res.TenancyOcid)
	}

	return nil
}

//func GetSocialIdentityProviderCompositeId(socialIdentityProviderId string) string {
//	id = url.PathEscape(id)
//	idcsEndpoint = url.PathEscape(idcsEndpoint)
//	socialIdentityProviderId = url.PathEscape(socialIdentityProviderId)
//	compositeId := "idcsEndpoint/" + idcsEndpoint + "/socialIdentityProviders/" + socialIdentityProviderId
//	return compositeId
//}

func parseSocialIdentityProviderCompositeId(compositeId string) (socialIdentityProviderId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("idcsEndpoint/.*/socialIdentityProviders/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	// idcsEndpoint, _ = url.PathUnescape(parts[1])
	socialIdentityProviderId, _ = url.PathUnescape(parts[3])

	return
}

func SocialIdentityProviderToMap(obj oci_identity_domains.SocialIdentityProvider) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AccessTokenUrl != nil {
		result["access_token_url"] = string(*obj.AccessTokenUrl)
	}

	if obj.AccountLinkingEnabled != nil {
		result["account_linking_enabled"] = bool(*obj.AccountLinkingEnabled)
	}

	result["admin_scope"] = obj.AdminScope

	if obj.AuthzUrl != nil {
		result["authz_url"] = string(*obj.AuthzUrl)
	}

	if obj.AutoRedirectEnabled != nil {
		result["auto_redirect_enabled"] = bool(*obj.AutoRedirectEnabled)
	}

	if obj.ClientCredentialInPayload != nil {
		result["client_credential_in_payload"] = bool(*obj.ClientCredentialInPayload)
	}

	if obj.ClockSkewInSeconds != nil {
		result["clock_skew_in_seconds"] = int(*obj.ClockSkewInSeconds)
	}

	if obj.CompartmentOcid != nil {
		result["compartment_ocid"] = string(*obj.CompartmentOcid)
	}

	if obj.ConsumerKey != nil {
		result["consumer_key"] = string(*obj.ConsumerKey)
	}

	if obj.ConsumerSecret != nil {
		result["consumer_secret"] = string(*obj.ConsumerSecret)
	}

	if obj.DeleteInProgress != nil {
		result["delete_in_progress"] = bool(*obj.DeleteInProgress)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DiscoveryUrl != nil {
		result["discovery_url"] = string(*obj.DiscoveryUrl)
	}

	if obj.DomainOcid != nil {
		result["domain_ocid"] = string(*obj.DomainOcid)
	}

	if obj.Enabled != nil {
		result["enabled"] = bool(*obj.Enabled)
	}

	if obj.IconUrl != nil {
		result["icon_url"] = string(*obj.IconUrl)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IdAttribute != nil {
		result["id_attribute"] = string(*obj.IdAttribute)
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

	jitProvAssignedGroups := []interface{}{}
	for _, item := range obj.JitProvAssignedGroups {
		jitProvAssignedGroups = append(jitProvAssignedGroups, SocialIdentityProviderJitProvAssignedGroupsToMap(item))
	}
	result["jit_prov_assigned_groups"] = jitProvAssignedGroups

	if obj.JitProvGroupStaticListEnabled != nil {
		result["jit_prov_group_static_list_enabled"] = bool(*obj.JitProvGroupStaticListEnabled)
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

	if obj.ProfileUrl != nil {
		result["profile_url"] = string(*obj.ProfileUrl)
	}

	if obj.RedirectUrl != nil {
		result["redirect_url"] = string(*obj.RedirectUrl)
	}

	if obj.RefreshTokenUrl != nil {
		result["refresh_token_url"] = string(*obj.RefreshTokenUrl)
	}

	if obj.RegistrationEnabled != nil {
		result["registration_enabled"] = bool(*obj.RegistrationEnabled)
	}

	relayIdpParamMappings := []interface{}{}
	for _, item := range obj.RelayIdpParamMappings {
		relayIdpParamMappings = append(relayIdpParamMappings, SocialIdentityProviderRelayIdpParamMappingsToMap(item))
	}
	result["relay_idp_param_mappings"] = relayIdpParamMappings

	result["schemas"] = obj.Schemas

	result["scope"] = obj.Scope

	if obj.ServiceProviderName != nil {
		result["service_provider_name"] = string(*obj.ServiceProviderName)
	}

	if obj.ShowOnLogin != nil {
		result["show_on_login"] = bool(*obj.ShowOnLogin)
	}

	if obj.SocialJitProvisioningEnabled != nil {
		result["social_jit_provisioning_enabled"] = bool(*obj.SocialJitProvisioningEnabled)
	}

	result["status"] = string(obj.Status)

	tags := []interface{}{}
	for _, item := range obj.Tags {
		tags = append(tags, tagsToMap(item))
	}
	result["tags"] = tags

	if obj.TenancyOcid != nil {
		result["tenancy_ocid"] = string(*obj.TenancyOcid)
	}

	return result
}

func (s *IdentityDomainsSocialIdentityProviderResourceCrud) mapToSocialIdentityProviderJitProvAssignedGroups(fieldKeyFormat string) (oci_identity_domains.SocialIdentityProviderJitProvAssignedGroups, error) {
	result := oci_identity_domains.SocialIdentityProviderJitProvAssignedGroups{}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func SocialIdentityProviderJitProvAssignedGroupsToMap(obj oci_identity_domains.SocialIdentityProviderJitProvAssignedGroups) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsSocialIdentityProviderResourceCrud) mapToSocialIdentityProviderRelayIdpParamMappings(fieldKeyFormat string) (oci_identity_domains.SocialIdentityProviderRelayIdpParamMappings, error) {
	result := oci_identity_domains.SocialIdentityProviderRelayIdpParamMappings{}

	if relayParamKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "relay_param_key")); ok {
		tmp := relayParamKey.(string)
		result.RelayParamKey = &tmp
	}

	if relayParamValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "relay_param_value")); ok {
		tmp := relayParamValue.(string)
		result.RelayParamValue = &tmp
	}

	return result, nil
}

func SocialIdentityProviderRelayIdpParamMappingsToMap(obj oci_identity_domains.SocialIdentityProviderRelayIdpParamMappings) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.RelayParamKey != nil {
		result["relay_param_key"] = string(*obj.RelayParamKey)
	}

	if obj.RelayParamValue != nil {
		result["relay_param_value"] = string(*obj.RelayParamValue)
	}

	return result
}

func (s *IdentityDomainsSocialIdentityProviderResourceCrud) mapTotags(fieldKeyFormat string) (oci_identity_domains.Tags, error) {
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
