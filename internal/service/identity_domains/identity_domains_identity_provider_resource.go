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

func IdentityDomainsIdentityProviderResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityDomainsIdentityProvider,
		Read:     readIdentityDomainsIdentityProvider,
		Update:   updateIdentityDomainsIdentityProvider,
		Delete:   deleteIdentityDomainsIdentityProvider,
		Schema: map[string]*schema.Schema{
			// Required
			"enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"idcs_endpoint": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"partner_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"schemas": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"assertion_attribute": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"authn_request_binding": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authorization": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"correlation_policy": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
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

						// Computed
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encryption_certificate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icon_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"idp_sso_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"include_signing_cert_in_signature": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"jit_user_prov_assigned_groups": {
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
						"display": {
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
			"jit_user_prov_attribute_update_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"jit_user_prov_attributes": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
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
					},
				},
			},
			"jit_user_prov_create_user_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"jit_user_prov_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"jit_user_prov_group_assertion_attribute_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"jit_user_prov_group_assignment_method": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"jit_user_prov_group_mapping_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"jit_user_prov_group_mappings": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"idp_group": {
							Type:     schema.TypeString,
							Required: true,
						},
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
					},
				},
			},
			"jit_user_prov_group_saml_attribute_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"jit_user_prov_group_static_list_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"jit_user_prov_ignore_error_on_absent_groups": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"logout_binding": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logout_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"logout_request_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logout_response_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metadata": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name_id_format": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ocid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"partner_provider_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"requested_authentication_context": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"require_force_authn": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"requires_encrypted_assertion": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"resource_type_schema_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"saml_ho_krequired": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"service_instance_identifier": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"shown_on_login_page": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"signature_hash_algorithm": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"signing_certificate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"succinct_id": {
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
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
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
						"registration_enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"service_provider_name": {
							Type:     schema.TypeString,
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
						"authz_url": {
							Type:     schema.TypeString,
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
						"discovery_url": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id_attribute": {
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
						"scope": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"status": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"urnietfparamsscimschemasoracleidcsextensionx509identity_provider": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"cert_match_attribute": {
							Type:     schema.TypeString,
							Required: true,
						},
						"signing_certificate_chain": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"user_match_attribute": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"crl_check_on_ocsp_failure_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"crl_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"crl_location": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"crl_reload_duration": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"eku_validation_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"eku_values": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"ocsp_allow_unknown_response_status": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"ocsp_enable_signed_response": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"ocsp_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"ocsp_responder_url": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ocsp_revalidate_time": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ocsp_server_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ocsp_trust_cert_chain": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"other_cert_match_attribute": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"user_mapping_method": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_mapping_store_attribute": {
				Type:     schema.TypeString,
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
			"last_notification_sent_time": {
				Type:     schema.TypeString,
				Computed: true,
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
			"tenant_provider_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createIdentityDomainsIdentityProvider(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsIdentityProviderResourceCrud{}
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

func readIdentityDomainsIdentityProvider(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsIdentityProviderResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpointForRead(d, "identityProviders")
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

func updateIdentityDomainsIdentityProvider(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsIdentityProviderResourceCrud{}
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

func deleteIdentityDomainsIdentityProvider(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsIdentityProviderResourceCrud{}
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

type IdentityDomainsIdentityProviderResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity_domains.IdentityDomainsClient
	Res                    *oci_identity_domains.IdentityProvider
	DisableNotFoundRetries bool
}

func (s *IdentityDomainsIdentityProviderResourceCrud) ID() string {
	return *s.Res.Id
	//return GetIdentityProviderCompositeId(s.D.Get("id").(string))
}

func (s *IdentityDomainsIdentityProviderResourceCrud) Create() error {
	request := oci_identity_domains.CreateIdentityProviderRequest{}

	if assertionAttribute, ok := s.D.GetOkExists("assertion_attribute"); ok {
		tmp := assertionAttribute.(string)
		request.AssertionAttribute = &tmp
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

	if authnRequestBinding, ok := s.D.GetOkExists("authn_request_binding"); ok {
		request.AuthnRequestBinding = oci_identity_domains.IdentityProviderAuthnRequestBindingEnum(authnRequestBinding.(string))
	}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if correlationPolicy, ok := s.D.GetOkExists("correlation_policy"); ok {
		if tmpList := correlationPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "correlation_policy", 0)
			tmp, err := s.mapToIdentityProviderCorrelationPolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CorrelationPolicy = &tmp
		}
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if enabled, ok := s.D.GetOkExists("enabled"); ok {
		tmp := enabled.(bool)
		request.Enabled = &tmp
	}

	if encryptionCertificate, ok := s.D.GetOkExists("encryption_certificate"); ok {
		tmp := encryptionCertificate.(string)
		request.EncryptionCertificate = &tmp
	}

	if externalId, ok := s.D.GetOkExists("external_id"); ok {
		tmp := externalId.(string)
		request.ExternalId = &tmp
	}

	if iconUrl, ok := s.D.GetOkExists("icon_url"); ok {
		tmp := iconUrl.(string)
		request.IconUrl = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if idpSsoUrl, ok := s.D.GetOkExists("idp_sso_url"); ok {
		tmp := idpSsoUrl.(string)
		request.IdpSsoUrl = &tmp
	}

	if includeSigningCertInSignature, ok := s.D.GetOkExists("include_signing_cert_in_signature"); ok {
		tmp := includeSigningCertInSignature.(bool)
		request.IncludeSigningCertInSignature = &tmp
	}

	if jitUserProvAssignedGroups, ok := s.D.GetOkExists("jit_user_prov_assigned_groups"); ok {
		interfaces := jitUserProvAssignedGroups.([]interface{})
		tmp := make([]oci_identity_domains.IdentityProviderJitUserProvAssignedGroups, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "jit_user_prov_assigned_groups", stateDataIndex)
			converted, err := s.mapToIdentityProviderJitUserProvAssignedGroups(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("jit_user_prov_assigned_groups") {
			request.JitUserProvAssignedGroups = tmp
		}
	}

	if jitUserProvAttributeUpdateEnabled, ok := s.D.GetOkExists("jit_user_prov_attribute_update_enabled"); ok {
		tmp := jitUserProvAttributeUpdateEnabled.(bool)
		request.JitUserProvAttributeUpdateEnabled = &tmp
	}

	if jitUserProvAttributes, ok := s.D.GetOkExists("jit_user_prov_attributes"); ok {
		if tmpList := jitUserProvAttributes.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "jit_user_prov_attributes", 0)
			tmp, err := s.mapToIdentityProviderJitUserProvAttributes(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.JitUserProvAttributes = &tmp
		}
	}

	if jitUserProvCreateUserEnabled, ok := s.D.GetOkExists("jit_user_prov_create_user_enabled"); ok {
		tmp := jitUserProvCreateUserEnabled.(bool)
		request.JitUserProvCreateUserEnabled = &tmp
	}

	if jitUserProvEnabled, ok := s.D.GetOkExists("jit_user_prov_enabled"); ok {
		tmp := jitUserProvEnabled.(bool)
		request.JitUserProvEnabled = &tmp
	}

	if jitUserProvGroupAssertionAttributeEnabled, ok := s.D.GetOkExists("jit_user_prov_group_assertion_attribute_enabled"); ok {
		tmp := jitUserProvGroupAssertionAttributeEnabled.(bool)
		request.JitUserProvGroupAssertionAttributeEnabled = &tmp
	}

	if jitUserProvGroupAssignmentMethod, ok := s.D.GetOkExists("jit_user_prov_group_assignment_method"); ok {
		request.JitUserProvGroupAssignmentMethod = oci_identity_domains.IdentityProviderJitUserProvGroupAssignmentMethodEnum(jitUserProvGroupAssignmentMethod.(string))
	}

	if jitUserProvGroupMappingMode, ok := s.D.GetOkExists("jit_user_prov_group_mapping_mode"); ok {
		request.JitUserProvGroupMappingMode = oci_identity_domains.IdentityProviderJitUserProvGroupMappingModeEnum(jitUserProvGroupMappingMode.(string))
	}

	if jitUserProvGroupMappings, ok := s.D.GetOkExists("jit_user_prov_group_mappings"); ok {
		interfaces := jitUserProvGroupMappings.([]interface{})
		tmp := make([]oci_identity_domains.IdentityProviderJitUserProvGroupMappings, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "jit_user_prov_group_mappings", stateDataIndex)
			converted, err := s.mapToIdentityProviderJitUserProvGroupMappings(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("jit_user_prov_group_mappings") {
			request.JitUserProvGroupMappings = tmp
		}
	}

	if jitUserProvGroupSAMLAttributeName, ok := s.D.GetOkExists("jit_user_prov_group_saml_attribute_name"); ok {
		tmp := jitUserProvGroupSAMLAttributeName.(string)
		request.JitUserProvGroupSAMLAttributeName = &tmp
	}

	if jitUserProvGroupStaticListEnabled, ok := s.D.GetOkExists("jit_user_prov_group_static_list_enabled"); ok {
		tmp := jitUserProvGroupStaticListEnabled.(bool)
		request.JitUserProvGroupStaticListEnabled = &tmp
	}

	if jitUserProvIgnoreErrorOnAbsentGroups, ok := s.D.GetOkExists("jit_user_prov_ignore_error_on_absent_groups"); ok {
		tmp := jitUserProvIgnoreErrorOnAbsentGroups.(bool)
		request.JitUserProvIgnoreErrorOnAbsentGroups = &tmp
	}

	if logoutBinding, ok := s.D.GetOkExists("logout_binding"); ok {
		request.LogoutBinding = oci_identity_domains.IdentityProviderLogoutBindingEnum(logoutBinding.(string))
	}

	if logoutEnabled, ok := s.D.GetOkExists("logout_enabled"); ok {
		tmp := logoutEnabled.(bool)
		request.LogoutEnabled = &tmp
	}

	if logoutRequestUrl, ok := s.D.GetOkExists("logout_request_url"); ok {
		tmp := logoutRequestUrl.(string)
		request.LogoutRequestUrl = &tmp
	}

	if logoutResponseUrl, ok := s.D.GetOkExists("logout_response_url"); ok {
		tmp := logoutResponseUrl.(string)
		request.LogoutResponseUrl = &tmp
	}

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		tmp := metadata.(string)
		request.Metadata = &tmp
	}

	if nameIdFormat, ok := s.D.GetOkExists("name_id_format"); ok {
		tmp := nameIdFormat.(string)
		request.NameIdFormat = &tmp
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if partnerName, ok := s.D.GetOkExists("partner_name"); ok {
		tmp := partnerName.(string)
		request.PartnerName = &tmp
	}

	if partnerProviderId, ok := s.D.GetOkExists("partner_provider_id"); ok {
		tmp := partnerProviderId.(string)
		request.PartnerProviderId = &tmp
	}

	if requestedAuthenticationContext, ok := s.D.GetOkExists("requested_authentication_context"); ok {
		interfaces := requestedAuthenticationContext.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("requested_authentication_context") {
			request.RequestedAuthenticationContext = tmp
		}
	}

	if requireForceAuthn, ok := s.D.GetOkExists("require_force_authn"); ok {
		tmp := requireForceAuthn.(bool)
		request.RequireForceAuthn = &tmp
	}

	if requiresEncryptedAssertion, ok := s.D.GetOkExists("requires_encrypted_assertion"); ok {
		tmp := requiresEncryptedAssertion.(bool)
		request.RequiresEncryptedAssertion = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	if samlHoKRequired, ok := s.D.GetOkExists("saml_ho_krequired"); ok {
		tmp := samlHoKRequired.(bool)
		request.SamlHoKRequired = &tmp
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

	if serviceInstanceIdentifier, ok := s.D.GetOkExists("service_instance_identifier"); ok {
		tmp := serviceInstanceIdentifier.(string)
		request.ServiceInstanceIdentifier = &tmp
	}

	if shownOnLoginPage, ok := s.D.GetOkExists("shown_on_login_page"); ok {
		tmp := shownOnLoginPage.(bool)
		request.ShownOnLoginPage = &tmp
	}

	if signatureHashAlgorithm, ok := s.D.GetOkExists("signature_hash_algorithm"); ok {
		request.SignatureHashAlgorithm = oci_identity_domains.IdentityProviderSignatureHashAlgorithmEnum(signatureHashAlgorithm.(string))
	}

	if signingCertificate, ok := s.D.GetOkExists("signing_certificate"); ok {
		tmp := signingCertificate.(string)
		request.SigningCertificate = &tmp
	}

	if succinctId, ok := s.D.GetOkExists("succinct_id"); ok {
		tmp := succinctId.(string)
		request.SuccinctId = &tmp
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

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_identity_domains.IdentityProviderTypeEnum(type_.(string))
	}

	if urnietfparamsscimschemasoracleidcsextensionsocialIdentityProvider, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionsocialIdentityProvider.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider", 0)
			tmp, err := s.mapToExtensionSocialIdentityProvider(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionSocialIdentityProvider = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionx509IdentityProvider, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionx509identity_provider"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionx509IdentityProvider.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionx509identity_provider", 0)
			tmp, err := s.mapToExtensionX509IdentityProvider(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionX509IdentityProvider = &tmp
		}
	}

	if userMappingMethod, ok := s.D.GetOkExists("user_mapping_method"); ok {
		request.UserMappingMethod = oci_identity_domains.IdentityProviderUserMappingMethodEnum(userMappingMethod.(string))
	}

	if userMappingStoreAttribute, ok := s.D.GetOkExists("user_mapping_store_attribute"); ok {
		tmp := userMappingStoreAttribute.(string)
		request.UserMappingStoreAttribute = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.CreateIdentityProvider(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IdentityProvider
	return nil
}

func (s *IdentityDomainsIdentityProviderResourceCrud) Get() error {
	request := oci_identity_domains.GetIdentityProviderRequest{}

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
	request.IdentityProviderId = &tmp

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	identityProviderId, err := parseIdentityProviderCompositeId(s.D.Id())
	if err == nil {
		request.IdentityProviderId = &identityProviderId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.GetIdentityProvider(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IdentityProvider
	return nil
}

func (s *IdentityDomainsIdentityProviderResourceCrud) Update() error {
	request := oci_identity_domains.PutIdentityProviderRequest{}

	if assertionAttribute, ok := s.D.GetOkExists("assertion_attribute"); ok {
		tmp := assertionAttribute.(string)
		request.AssertionAttribute = &tmp
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

	if authnRequestBinding, ok := s.D.GetOkExists("authn_request_binding"); ok {
		request.AuthnRequestBinding = oci_identity_domains.IdentityProviderAuthnRequestBindingEnum(authnRequestBinding.(string))
	}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if correlationPolicy, ok := s.D.GetOkExists("correlation_policy"); ok {
		if tmpList := correlationPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "correlation_policy", 0)
			tmp, err := s.mapToIdentityProviderCorrelationPolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CorrelationPolicy = &tmp
		}
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if enabled, ok := s.D.GetOkExists("enabled"); ok {
		tmp := enabled.(bool)
		request.Enabled = &tmp
	}

	if encryptionCertificate, ok := s.D.GetOkExists("encryption_certificate"); ok {
		tmp := encryptionCertificate.(string)
		request.EncryptionCertificate = &tmp
	}

	if externalId, ok := s.D.GetOkExists("external_id"); ok {
		tmp := externalId.(string)
		request.ExternalId = &tmp
	}

	if iconUrl, ok := s.D.GetOkExists("icon_url"); ok {
		tmp := iconUrl.(string)
		request.IconUrl = &tmp
	}

	tmp := s.D.Id()
	request.Id = &tmp

	tmp = s.D.Id()
	request.IdentityProviderId = &tmp

	if idpSsoUrl, ok := s.D.GetOkExists("idp_sso_url"); ok {
		tmp := idpSsoUrl.(string)
		request.IdpSsoUrl = &tmp
	}

	if includeSigningCertInSignature, ok := s.D.GetOkExists("include_signing_cert_in_signature"); ok {
		tmp := includeSigningCertInSignature.(bool)
		request.IncludeSigningCertInSignature = &tmp
	}

	if jitUserProvAssignedGroups, ok := s.D.GetOkExists("jit_user_prov_assigned_groups"); ok {
		interfaces := jitUserProvAssignedGroups.([]interface{})
		tmp := make([]oci_identity_domains.IdentityProviderJitUserProvAssignedGroups, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "jit_user_prov_assigned_groups", stateDataIndex)
			converted, err := s.mapToIdentityProviderJitUserProvAssignedGroups(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("jit_user_prov_assigned_groups") {
			request.JitUserProvAssignedGroups = tmp
		}
	}

	if jitUserProvAttributeUpdateEnabled, ok := s.D.GetOkExists("jit_user_prov_attribute_update_enabled"); ok {
		tmp := jitUserProvAttributeUpdateEnabled.(bool)
		request.JitUserProvAttributeUpdateEnabled = &tmp
	}

	if jitUserProvAttributes, ok := s.D.GetOkExists("jit_user_prov_attributes"); ok {
		if tmpList := jitUserProvAttributes.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "jit_user_prov_attributes", 0)
			tmp, err := s.mapToIdentityProviderJitUserProvAttributes(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.JitUserProvAttributes = &tmp
		}
	}

	if jitUserProvCreateUserEnabled, ok := s.D.GetOkExists("jit_user_prov_create_user_enabled"); ok {
		tmp := jitUserProvCreateUserEnabled.(bool)
		request.JitUserProvCreateUserEnabled = &tmp
	}

	if jitUserProvEnabled, ok := s.D.GetOkExists("jit_user_prov_enabled"); ok {
		tmp := jitUserProvEnabled.(bool)
		request.JitUserProvEnabled = &tmp
	}

	if jitUserProvGroupAssertionAttributeEnabled, ok := s.D.GetOkExists("jit_user_prov_group_assertion_attribute_enabled"); ok {
		tmp := jitUserProvGroupAssertionAttributeEnabled.(bool)
		request.JitUserProvGroupAssertionAttributeEnabled = &tmp
	}

	if jitUserProvGroupAssignmentMethod, ok := s.D.GetOkExists("jit_user_prov_group_assignment_method"); ok {
		request.JitUserProvGroupAssignmentMethod = oci_identity_domains.IdentityProviderJitUserProvGroupAssignmentMethodEnum(jitUserProvGroupAssignmentMethod.(string))
	}

	if jitUserProvGroupMappingMode, ok := s.D.GetOkExists("jit_user_prov_group_mapping_mode"); ok {
		request.JitUserProvGroupMappingMode = oci_identity_domains.IdentityProviderJitUserProvGroupMappingModeEnum(jitUserProvGroupMappingMode.(string))
	}

	if jitUserProvGroupMappings, ok := s.D.GetOkExists("jit_user_prov_group_mappings"); ok {
		interfaces := jitUserProvGroupMappings.([]interface{})
		tmp := make([]oci_identity_domains.IdentityProviderJitUserProvGroupMappings, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "jit_user_prov_group_mappings", stateDataIndex)
			converted, err := s.mapToIdentityProviderJitUserProvGroupMappings(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("jit_user_prov_group_mappings") {
			request.JitUserProvGroupMappings = tmp
		}
	}

	if jitUserProvGroupSAMLAttributeName, ok := s.D.GetOkExists("jit_user_prov_group_saml_attribute_name"); ok {
		tmp := jitUserProvGroupSAMLAttributeName.(string)
		request.JitUserProvGroupSAMLAttributeName = &tmp
	}

	if jitUserProvGroupStaticListEnabled, ok := s.D.GetOkExists("jit_user_prov_group_static_list_enabled"); ok {
		tmp := jitUserProvGroupStaticListEnabled.(bool)
		request.JitUserProvGroupStaticListEnabled = &tmp
	}

	if jitUserProvIgnoreErrorOnAbsentGroups, ok := s.D.GetOkExists("jit_user_prov_ignore_error_on_absent_groups"); ok {
		tmp := jitUserProvIgnoreErrorOnAbsentGroups.(bool)
		request.JitUserProvIgnoreErrorOnAbsentGroups = &tmp
	}

	if logoutBinding, ok := s.D.GetOkExists("logout_binding"); ok {
		request.LogoutBinding = oci_identity_domains.IdentityProviderLogoutBindingEnum(logoutBinding.(string))
	}

	if logoutEnabled, ok := s.D.GetOkExists("logout_enabled"); ok {
		tmp := logoutEnabled.(bool)
		request.LogoutEnabled = &tmp
	}

	if logoutRequestUrl, ok := s.D.GetOkExists("logout_request_url"); ok {
		tmp := logoutRequestUrl.(string)
		request.LogoutRequestUrl = &tmp
	}

	if logoutResponseUrl, ok := s.D.GetOkExists("logout_response_url"); ok {
		tmp := logoutResponseUrl.(string)
		request.LogoutResponseUrl = &tmp
	}

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		tmp := metadata.(string)
		request.Metadata = &tmp
	}

	if nameIdFormat, ok := s.D.GetOkExists("name_id_format"); ok {
		tmp := nameIdFormat.(string)
		request.NameIdFormat = &tmp
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if partnerName, ok := s.D.GetOkExists("partner_name"); ok {
		tmp := partnerName.(string)
		request.PartnerName = &tmp
	}

	if partnerProviderId, ok := s.D.GetOkExists("partner_provider_id"); ok {
		tmp := partnerProviderId.(string)
		request.PartnerProviderId = &tmp
	}

	if requestedAuthenticationContext, ok := s.D.GetOkExists("requested_authentication_context"); ok {
		interfaces := requestedAuthenticationContext.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("requested_authentication_context") {
			request.RequestedAuthenticationContext = tmp
		}
	}

	if requireForceAuthn, ok := s.D.GetOkExists("require_force_authn"); ok {
		tmp := requireForceAuthn.(bool)
		request.RequireForceAuthn = &tmp
	}

	if requiresEncryptedAssertion, ok := s.D.GetOkExists("requires_encrypted_assertion"); ok {
		tmp := requiresEncryptedAssertion.(bool)
		request.RequiresEncryptedAssertion = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	if samlHoKRequired, ok := s.D.GetOkExists("saml_ho_krequired"); ok {
		tmp := samlHoKRequired.(bool)
		request.SamlHoKRequired = &tmp
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

	if serviceInstanceIdentifier, ok := s.D.GetOkExists("service_instance_identifier"); ok {
		tmp := serviceInstanceIdentifier.(string)
		request.ServiceInstanceIdentifier = &tmp
	}

	if shownOnLoginPage, ok := s.D.GetOkExists("shown_on_login_page"); ok {
		tmp := shownOnLoginPage.(bool)
		request.ShownOnLoginPage = &tmp
	}

	if signatureHashAlgorithm, ok := s.D.GetOkExists("signature_hash_algorithm"); ok {
		request.SignatureHashAlgorithm = oci_identity_domains.IdentityProviderSignatureHashAlgorithmEnum(signatureHashAlgorithm.(string))
	}

	if signingCertificate, ok := s.D.GetOkExists("signing_certificate"); ok {
		tmp := signingCertificate.(string)
		request.SigningCertificate = &tmp
	}

	if succinctId, ok := s.D.GetOkExists("succinct_id"); ok {
		tmp := succinctId.(string)
		request.SuccinctId = &tmp
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

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_identity_domains.IdentityProviderTypeEnum(type_.(string))
	}

	if urnietfparamsscimschemasoracleidcsextensionsocialIdentityProvider, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionsocialIdentityProvider.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider", 0)
			tmp, err := s.mapToExtensionSocialIdentityProvider(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionSocialIdentityProvider = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionx509IdentityProvider, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionx509identity_provider"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionx509IdentityProvider.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionx509identity_provider", 0)
			tmp, err := s.mapToExtensionX509IdentityProvider(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionX509IdentityProvider = &tmp
		}
	}

	if userMappingMethod, ok := s.D.GetOkExists("user_mapping_method"); ok {
		request.UserMappingMethod = oci_identity_domains.IdentityProviderUserMappingMethodEnum(userMappingMethod.(string))
	}

	if userMappingStoreAttribute, ok := s.D.GetOkExists("user_mapping_store_attribute"); ok {
		tmp := userMappingStoreAttribute.(string)
		request.UserMappingStoreAttribute = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.PutIdentityProvider(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IdentityProvider
	return nil
}

func (s *IdentityDomainsIdentityProviderResourceCrud) Delete() error {
	request := oci_identity_domains.DeleteIdentityProviderRequest{}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if forceDelete, ok := s.D.GetOkExists("force_delete"); ok {
		tmp := forceDelete.(bool)
		request.ForceDelete = &tmp
	}

	tmp := s.D.Id()
	request.IdentityProviderId = &tmp

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	_, err := s.Client.DeleteIdentityProvider(context.Background(), request)
	return err
}

func (s *IdentityDomainsIdentityProviderResourceCrud) SetData() error {

	identityProviderId, err := parseIdentityProviderCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(identityProviderId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.AssertionAttribute != nil {
		s.D.Set("assertion_attribute", *s.Res.AssertionAttribute)
	}

	s.D.Set("authn_request_binding", s.Res.AuthnRequestBinding)

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
	}

	if s.Res.CorrelationPolicy != nil {
		s.D.Set("correlation_policy", []interface{}{IdentityProviderCorrelationPolicyToMap(s.Res.CorrelationPolicy)})
	} else {
		s.D.Set("correlation_policy", nil)
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

	if s.Res.Enabled != nil {
		s.D.Set("enabled", *s.Res.Enabled)
	}

	if s.Res.EncryptionCertificate != nil {
		s.D.Set("encryption_certificate", *s.Res.EncryptionCertificate)
	}

	if s.Res.ExternalId != nil {
		s.D.Set("external_id", *s.Res.ExternalId)
	}

	if s.Res.IconUrl != nil {
		s.D.Set("icon_url", *s.Res.IconUrl)
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

	if s.Res.IdpSsoUrl != nil {
		s.D.Set("idp_sso_url", *s.Res.IdpSsoUrl)
	}

	if s.Res.IncludeSigningCertInSignature != nil {
		s.D.Set("include_signing_cert_in_signature", *s.Res.IncludeSigningCertInSignature)
	}

	jitUserProvAssignedGroups := []interface{}{}
	for _, item := range s.Res.JitUserProvAssignedGroups {
		jitUserProvAssignedGroups = append(jitUserProvAssignedGroups, IdentityProviderJitUserProvAssignedGroupsToMap(item))
	}
	s.D.Set("jit_user_prov_assigned_groups", jitUserProvAssignedGroups)

	if s.Res.JitUserProvAttributeUpdateEnabled != nil {
		s.D.Set("jit_user_prov_attribute_update_enabled", *s.Res.JitUserProvAttributeUpdateEnabled)
	}

	if s.Res.JitUserProvAttributes != nil {
		s.D.Set("jit_user_prov_attributes", []interface{}{IdentityProviderJitUserProvAttributesToMap(s.Res.JitUserProvAttributes)})
	} else {
		s.D.Set("jit_user_prov_attributes", nil)
	}

	if s.Res.JitUserProvCreateUserEnabled != nil {
		s.D.Set("jit_user_prov_create_user_enabled", *s.Res.JitUserProvCreateUserEnabled)
	}

	if s.Res.JitUserProvEnabled != nil {
		s.D.Set("jit_user_prov_enabled", *s.Res.JitUserProvEnabled)
	}

	if s.Res.JitUserProvGroupAssertionAttributeEnabled != nil {
		s.D.Set("jit_user_prov_group_assertion_attribute_enabled", *s.Res.JitUserProvGroupAssertionAttributeEnabled)
	}

	s.D.Set("jit_user_prov_group_assignment_method", s.Res.JitUserProvGroupAssignmentMethod)

	s.D.Set("jit_user_prov_group_mapping_mode", s.Res.JitUserProvGroupMappingMode)

	jitUserProvGroupMappings := []interface{}{}
	for _, item := range s.Res.JitUserProvGroupMappings {
		jitUserProvGroupMappings = append(jitUserProvGroupMappings, IdentityProviderJitUserProvGroupMappingsToMap(item))
	}
	s.D.Set("jit_user_prov_group_mappings", jitUserProvGroupMappings)

	if s.Res.JitUserProvGroupSAMLAttributeName != nil {
		s.D.Set("jit_user_prov_group_saml_attribute_name", *s.Res.JitUserProvGroupSAMLAttributeName)
	}

	if s.Res.JitUserProvGroupStaticListEnabled != nil {
		s.D.Set("jit_user_prov_group_static_list_enabled", *s.Res.JitUserProvGroupStaticListEnabled)
	}

	if s.Res.JitUserProvIgnoreErrorOnAbsentGroups != nil {
		s.D.Set("jit_user_prov_ignore_error_on_absent_groups", *s.Res.JitUserProvIgnoreErrorOnAbsentGroups)
	}

	if s.Res.LastNotificationSentTime != nil {
		s.D.Set("last_notification_sent_time", *s.Res.LastNotificationSentTime)
	}

	s.D.Set("logout_binding", s.Res.LogoutBinding)

	if s.Res.LogoutEnabled != nil {
		s.D.Set("logout_enabled", *s.Res.LogoutEnabled)
	}

	if s.Res.LogoutRequestUrl != nil {
		s.D.Set("logout_request_url", *s.Res.LogoutRequestUrl)
	}

	if s.Res.LogoutResponseUrl != nil {
		s.D.Set("logout_response_url", *s.Res.LogoutResponseUrl)
	}

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	if s.Res.Metadata != nil {
		s.D.Set("metadata", *s.Res.Metadata)
	}

	if s.Res.NameIdFormat != nil {
		s.D.Set("name_id_format", *s.Res.NameIdFormat)
	}

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
	}

	if s.Res.PartnerName != nil {
		s.D.Set("partner_name", *s.Res.PartnerName)
	}

	if s.Res.PartnerProviderId != nil {
		s.D.Set("partner_provider_id", *s.Res.PartnerProviderId)
	}

	s.D.Set("requested_authentication_context", s.Res.RequestedAuthenticationContext)

	if s.Res.RequireForceAuthn != nil {
		s.D.Set("require_force_authn", *s.Res.RequireForceAuthn)
	}

	if s.Res.RequiresEncryptedAssertion != nil {
		s.D.Set("requires_encrypted_assertion", *s.Res.RequiresEncryptedAssertion)
	}

	if s.Res.SamlHoKRequired != nil {
		s.D.Set("saml_ho_krequired", *s.Res.SamlHoKRequired)
	}

	s.D.Set("schemas", s.Res.Schemas)

	if s.Res.ServiceInstanceIdentifier != nil {
		s.D.Set("service_instance_identifier", *s.Res.ServiceInstanceIdentifier)
	}

	if s.Res.ShownOnLoginPage != nil {
		s.D.Set("shown_on_login_page", *s.Res.ShownOnLoginPage)
	}

	s.D.Set("signature_hash_algorithm", s.Res.SignatureHashAlgorithm)

	if s.Res.SigningCertificate != nil {
		s.D.Set("signing_certificate", *s.Res.SigningCertificate)
	}

	if s.Res.SuccinctId != nil {
		s.D.Set("succinct_id", *s.Res.SuccinctId)
	}

	tags := []interface{}{}
	for _, item := range s.Res.Tags {
		tags = append(tags, tagsToMap(item))
	}
	s.D.Set("tags", tags)

	if s.Res.TenancyOcid != nil {
		s.D.Set("tenancy_ocid", *s.Res.TenancyOcid)
	}

	if s.Res.TenantProviderId != nil {
		s.D.Set("tenant_provider_id", *s.Res.TenantProviderId)
	}

	s.D.Set("type", s.Res.Type)

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionSocialIdentityProvider != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider", []interface{}{ExtensionSocialIdentityProviderToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionSocialIdentityProvider)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionX509IdentityProvider != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionx509identity_provider", []interface{}{ExtensionX509IdentityProviderToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionX509IdentityProvider)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionx509identity_provider", nil)
	}

	s.D.Set("user_mapping_method", s.Res.UserMappingMethod)

	if s.Res.UserMappingStoreAttribute != nil {
		s.D.Set("user_mapping_store_attribute", *s.Res.UserMappingStoreAttribute)
	}

	return nil
}

//func GetIdentityProviderCompositeId(identityProviderId string) string {
//	id = url.PathEscape(id)
//	idcsEndpoint = url.PathEscape(idcsEndpoint)
//	identityProviderId = url.PathEscape(identityProviderId)
//	compositeId := "idcsEndpoint/" + idcsEndpoint + "/identityProviders/" + identityProviderId
//	return compositeId
//}

func parseIdentityProviderCompositeId(compositeId string) (identityProviderId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("idcsEndpoint/.*/identityProviders/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	//idcsEndpoint, _ = url.PathUnescape(parts[1])
	identityProviderId, _ = url.PathUnescape(parts[3])

	return
}

func (s *IdentityDomainsIdentityProviderResourceCrud) mapToExtensionSocialIdentityProvider(fieldKeyFormat string) (oci_identity_domains.ExtensionSocialIdentityProvider, error) {
	result := oci_identity_domains.ExtensionSocialIdentityProvider{}

	if accessTokenUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "access_token_url")); ok {
		tmp := accessTokenUrl.(string)
		result.AccessTokenUrl = &tmp
	}

	if accountLinkingEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "account_linking_enabled")); ok {
		tmp := accountLinkingEnabled.(bool)
		result.AccountLinkingEnabled = &tmp
	}

	if adminScope, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_scope")); ok {
		interfaces := adminScope.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "admin_scope")) {
			result.AdminScope = tmp
		}
	}

	if authzUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "authz_url")); ok {
		tmp := authzUrl.(string)
		result.AuthzUrl = &tmp
	}

	if clientCredentialInPayload, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "client_credential_in_payload")); ok {
		tmp := clientCredentialInPayload.(bool)
		result.ClientCredentialInPayload = &tmp
	}

	if clockSkewInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "clock_skew_in_seconds")); ok {
		tmp := clockSkewInSeconds.(int)
		result.ClockSkewInSeconds = &tmp
	}

	if consumerKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "consumer_key")); ok {
		tmp := consumerKey.(string)
		result.ConsumerKey = &tmp
	}

	if consumerSecret, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "consumer_secret")); ok {
		tmp := consumerSecret.(string)
		result.ConsumerSecret = &tmp
	}

	if discoveryUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "discovery_url")); ok {
		tmp := discoveryUrl.(string)
		result.DiscoveryUrl = &tmp
	}

	if idAttribute, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id_attribute")); ok {
		tmp := idAttribute.(string)
		result.IdAttribute = &tmp
	}

	if profileUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "profile_url")); ok {
		tmp := profileUrl.(string)
		result.ProfileUrl = &tmp
	}

	if redirectUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "redirect_url")); ok {
		tmp := redirectUrl.(string)
		result.RedirectUrl = &tmp
	}

	if registrationEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "registration_enabled")); ok {
		tmp := registrationEnabled.(bool)
		result.RegistrationEnabled = &tmp
	}

	if scope, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scope")); ok {
		interfaces := scope.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "scope")) {
			result.Scope = tmp
		}
	}

	if serviceProviderName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "service_provider_name")); ok {
		tmp := serviceProviderName.(string)
		result.ServiceProviderName = &tmp
	}

	if status, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "status")); ok {
		result.Status = oci_identity_domains.ExtensionSocialIdentityProviderStatusEnum(status.(string))
	}

	return result, nil
}

func ExtensionSocialIdentityProviderToMap(obj *oci_identity_domains.ExtensionSocialIdentityProvider) map[string]interface{} {
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

	if obj.ClientCredentialInPayload != nil {
		result["client_credential_in_payload"] = bool(*obj.ClientCredentialInPayload)
	}

	if obj.ClockSkewInSeconds != nil {
		result["clock_skew_in_seconds"] = int(*obj.ClockSkewInSeconds)
	}

	if obj.ConsumerKey != nil {
		result["consumer_key"] = string(*obj.ConsumerKey)
	}

	if obj.ConsumerSecret != nil {
		result["consumer_secret"] = string(*obj.ConsumerSecret)
	}

	if obj.DiscoveryUrl != nil {
		result["discovery_url"] = string(*obj.DiscoveryUrl)
	}

	if obj.IdAttribute != nil {
		result["id_attribute"] = string(*obj.IdAttribute)
	}

	if obj.ProfileUrl != nil {
		result["profile_url"] = string(*obj.ProfileUrl)
	}

	if obj.RedirectUrl != nil {
		result["redirect_url"] = string(*obj.RedirectUrl)
	}

	if obj.RegistrationEnabled != nil {
		result["registration_enabled"] = bool(*obj.RegistrationEnabled)
	}

	result["scope"] = obj.Scope

	if obj.ServiceProviderName != nil {
		result["service_provider_name"] = string(*obj.ServiceProviderName)
	}

	result["status"] = string(obj.Status)

	return result
}

func (s *IdentityDomainsIdentityProviderResourceCrud) mapToExtensionX509IdentityProvider(fieldKeyFormat string) (oci_identity_domains.ExtensionX509IdentityProvider, error) {
	result := oci_identity_domains.ExtensionX509IdentityProvider{}

	if certMatchAttribute, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cert_match_attribute")); ok {
		tmp := certMatchAttribute.(string)
		result.CertMatchAttribute = &tmp
	}

	if crlCheckOnOCSPFailureEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "crl_check_on_ocsp_failure_enabled")); ok {
		tmp := crlCheckOnOCSPFailureEnabled.(bool)
		result.CrlCheckOnOCSPFailureEnabled = &tmp
	}

	if crlEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "crl_enabled")); ok {
		tmp := crlEnabled.(bool)
		result.CrlEnabled = &tmp
	}

	if crlLocation, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "crl_location")); ok {
		tmp := crlLocation.(string)
		result.CrlLocation = &tmp
	}

	if crlReloadDuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "crl_reload_duration")); ok {
		tmp := crlReloadDuration.(int)
		result.CrlReloadDuration = &tmp
	}

	if ekuValidationEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "eku_validation_enabled")); ok {
		tmp := ekuValidationEnabled.(bool)
		result.EkuValidationEnabled = &tmp
	}

	if ekuValues, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "eku_values")); ok {
		interfaces := ekuValues.([]interface{})
		tmp := make([]oci_identity_domains.ExtensionX509IdentityProviderEkuValuesEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_identity_domains.ExtensionX509IdentityProviderEkuValuesEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "eku_values")) {
			result.EkuValues = tmp
		}
	}

	if ocspAllowUnknownResponseStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocsp_allow_unknown_response_status")); ok {
		tmp := ocspAllowUnknownResponseStatus.(bool)
		result.OcspAllowUnknownResponseStatus = &tmp
	}

	if ocspEnableSignedResponse, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocsp_enable_signed_response")); ok {
		tmp := ocspEnableSignedResponse.(bool)
		result.OcspEnableSignedResponse = &tmp
	}

	if ocspEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocsp_enabled")); ok {
		tmp := ocspEnabled.(bool)
		result.OcspEnabled = &tmp
	}

	if ocspResponderURL, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocsp_responder_url")); ok {
		tmp := ocspResponderURL.(string)
		result.OcspResponderURL = &tmp
	}

	if ocspRevalidateTime, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocsp_revalidate_time")); ok {
		tmp := ocspRevalidateTime.(int)
		result.OcspRevalidateTime = &tmp
	}

	if ocspServerName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocsp_server_name")); ok {
		tmp := ocspServerName.(string)
		result.OcspServerName = &tmp
	}

	if ocspTrustCertChain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocsp_trust_cert_chain")); ok {
		interfaces := ocspTrustCertChain.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "ocsp_trust_cert_chain")) {
			result.OcspTrustCertChain = tmp
		}
	}

	if otherCertMatchAttribute, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "other_cert_match_attribute")); ok {
		tmp := otherCertMatchAttribute.(string)
		result.OtherCertMatchAttribute = &tmp
	}

	if signingCertificateChain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "signing_certificate_chain")); ok {
		interfaces := signingCertificateChain.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "signing_certificate_chain")) {
			result.SigningCertificateChain = tmp
		}
	}

	if userMatchAttribute, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_match_attribute")); ok {
		tmp := userMatchAttribute.(string)
		result.UserMatchAttribute = &tmp
	}

	return result, nil
}

func ExtensionX509IdentityProviderToMap(obj *oci_identity_domains.ExtensionX509IdentityProvider) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CertMatchAttribute != nil {
		result["cert_match_attribute"] = string(*obj.CertMatchAttribute)
	}

	if obj.CrlCheckOnOCSPFailureEnabled != nil {
		result["crl_check_on_ocsp_failure_enabled"] = bool(*obj.CrlCheckOnOCSPFailureEnabled)
	}

	if obj.CrlEnabled != nil {
		result["crl_enabled"] = bool(*obj.CrlEnabled)
	}

	if obj.CrlLocation != nil {
		result["crl_location"] = string(*obj.CrlLocation)
	}

	if obj.CrlReloadDuration != nil {
		result["crl_reload_duration"] = int(*obj.CrlReloadDuration)
	}

	if obj.EkuValidationEnabled != nil {
		result["eku_validation_enabled"] = bool(*obj.EkuValidationEnabled)
	}

	result["eku_values"] = obj.EkuValues

	if obj.OcspAllowUnknownResponseStatus != nil {
		result["ocsp_allow_unknown_response_status"] = bool(*obj.OcspAllowUnknownResponseStatus)
	}

	if obj.OcspEnableSignedResponse != nil {
		result["ocsp_enable_signed_response"] = bool(*obj.OcspEnableSignedResponse)
	}

	if obj.OcspEnabled != nil {
		result["ocsp_enabled"] = bool(*obj.OcspEnabled)
	}

	if obj.OcspResponderURL != nil {
		result["ocsp_responder_url"] = string(*obj.OcspResponderURL)
	}

	if obj.OcspRevalidateTime != nil {
		result["ocsp_revalidate_time"] = int(*obj.OcspRevalidateTime)
	}

	if obj.OcspServerName != nil {
		result["ocsp_server_name"] = string(*obj.OcspServerName)
	}

	result["ocsp_trust_cert_chain"] = obj.OcspTrustCertChain

	if obj.OtherCertMatchAttribute != nil {
		result["other_cert_match_attribute"] = string(*obj.OtherCertMatchAttribute)
	}

	result["signing_certificate_chain"] = obj.SigningCertificateChain

	if obj.UserMatchAttribute != nil {
		result["user_match_attribute"] = string(*obj.UserMatchAttribute)
	}

	return result
}

func IdentityProviderToMap(obj oci_identity_domains.IdentityProvider) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AssertionAttribute != nil {
		result["assertion_attribute"] = string(*obj.AssertionAttribute)
	}

	result["authn_request_binding"] = string(obj.AuthnRequestBinding)

	if obj.CompartmentOcid != nil {
		result["compartment_ocid"] = string(*obj.CompartmentOcid)
	}

	if obj.CorrelationPolicy != nil {
		result["correlation_policy"] = []interface{}{IdentityProviderCorrelationPolicyToMap(obj.CorrelationPolicy)}
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

	if obj.Enabled != nil {
		result["enabled"] = bool(*obj.Enabled)
	}

	if obj.EncryptionCertificate != nil {
		result["encryption_certificate"] = string(*obj.EncryptionCertificate)
	}

	if obj.ExternalId != nil {
		result["external_id"] = string(*obj.ExternalId)
	}

	if obj.IconUrl != nil {
		result["icon_url"] = string(*obj.IconUrl)
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

	if obj.IdpSsoUrl != nil {
		result["idp_sso_url"] = string(*obj.IdpSsoUrl)
	}

	if obj.IncludeSigningCertInSignature != nil {
		result["include_signing_cert_in_signature"] = bool(*obj.IncludeSigningCertInSignature)
	}

	jitUserProvAssignedGroups := []interface{}{}
	for _, item := range obj.JitUserProvAssignedGroups {
		jitUserProvAssignedGroups = append(jitUserProvAssignedGroups, IdentityProviderJitUserProvAssignedGroupsToMap(item))
	}
	result["jit_user_prov_assigned_groups"] = jitUserProvAssignedGroups

	if obj.JitUserProvAttributeUpdateEnabled != nil {
		result["jit_user_prov_attribute_update_enabled"] = bool(*obj.JitUserProvAttributeUpdateEnabled)
	}

	if obj.JitUserProvAttributes != nil {
		result["jit_user_prov_attributes"] = []interface{}{IdentityProviderJitUserProvAttributesToMap(obj.JitUserProvAttributes)}
	}

	if obj.JitUserProvCreateUserEnabled != nil {
		result["jit_user_prov_create_user_enabled"] = bool(*obj.JitUserProvCreateUserEnabled)
	}

	if obj.JitUserProvEnabled != nil {
		result["jit_user_prov_enabled"] = bool(*obj.JitUserProvEnabled)
	}

	if obj.JitUserProvGroupAssertionAttributeEnabled != nil {
		result["jit_user_prov_group_assertion_attribute_enabled"] = bool(*obj.JitUserProvGroupAssertionAttributeEnabled)
	}

	result["jit_user_prov_group_assignment_method"] = string(obj.JitUserProvGroupAssignmentMethod)

	result["jit_user_prov_group_mapping_mode"] = string(obj.JitUserProvGroupMappingMode)

	jitUserProvGroupMappings := []interface{}{}
	for _, item := range obj.JitUserProvGroupMappings {
		jitUserProvGroupMappings = append(jitUserProvGroupMappings, IdentityProviderJitUserProvGroupMappingsToMap(item))
	}
	result["jit_user_prov_group_mappings"] = jitUserProvGroupMappings

	if obj.JitUserProvGroupSAMLAttributeName != nil {
		result["jit_user_prov_group_saml_attribute_name"] = string(*obj.JitUserProvGroupSAMLAttributeName)
	}

	if obj.JitUserProvGroupStaticListEnabled != nil {
		result["jit_user_prov_group_static_list_enabled"] = bool(*obj.JitUserProvGroupStaticListEnabled)
	}

	if obj.JitUserProvIgnoreErrorOnAbsentGroups != nil {
		result["jit_user_prov_ignore_error_on_absent_groups"] = bool(*obj.JitUserProvIgnoreErrorOnAbsentGroups)
	}

	if obj.LastNotificationSentTime != nil {
		result["last_notification_sent_time"] = string(*obj.LastNotificationSentTime)
	}

	result["logout_binding"] = string(obj.LogoutBinding)

	if obj.LogoutEnabled != nil {
		result["logout_enabled"] = bool(*obj.LogoutEnabled)
	}

	if obj.LogoutRequestUrl != nil {
		result["logout_request_url"] = string(*obj.LogoutRequestUrl)
	}

	if obj.LogoutResponseUrl != nil {
		result["logout_response_url"] = string(*obj.LogoutResponseUrl)
	}

	if obj.Meta != nil {
		result["meta"] = []interface{}{metaToMap(obj.Meta)}
	}

	if obj.Metadata != nil {
		result["metadata"] = string(*obj.Metadata)
	}

	if obj.NameIdFormat != nil {
		result["name_id_format"] = string(*obj.NameIdFormat)
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	if obj.PartnerName != nil {
		result["partner_name"] = string(*obj.PartnerName)
	}

	if obj.PartnerProviderId != nil {
		result["partner_provider_id"] = string(*obj.PartnerProviderId)
	}

	result["requested_authentication_context"] = obj.RequestedAuthenticationContext

	if obj.RequireForceAuthn != nil {
		result["require_force_authn"] = bool(*obj.RequireForceAuthn)
	}

	if obj.RequiresEncryptedAssertion != nil {
		result["requires_encrypted_assertion"] = bool(*obj.RequiresEncryptedAssertion)
	}

	if obj.SamlHoKRequired != nil {
		result["saml_ho_krequired"] = bool(*obj.SamlHoKRequired)
	}

	result["schemas"] = obj.Schemas

	if obj.ServiceInstanceIdentifier != nil {
		result["service_instance_identifier"] = string(*obj.ServiceInstanceIdentifier)
	}

	if obj.ShownOnLoginPage != nil {
		result["shown_on_login_page"] = bool(*obj.ShownOnLoginPage)
	}

	result["signature_hash_algorithm"] = string(obj.SignatureHashAlgorithm)

	if obj.SigningCertificate != nil {
		result["signing_certificate"] = string(*obj.SigningCertificate)
	}

	if obj.SuccinctId != nil {
		result["succinct_id"] = string(*obj.SuccinctId)
	}

	tags := []interface{}{}
	for _, item := range obj.Tags {
		tags = append(tags, tagsToMap(item))
	}
	result["tags"] = tags

	if obj.TenancyOcid != nil {
		result["tenancy_ocid"] = string(*obj.TenancyOcid)
	}

	if obj.TenantProviderId != nil {
		result["tenant_provider_id"] = string(*obj.TenantProviderId)
	}

	result["type"] = string(obj.Type)

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionSocialIdentityProvider != nil {
		result["urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider"] = []interface{}{ExtensionSocialIdentityProviderToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionSocialIdentityProvider)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionX509IdentityProvider != nil {
		result["urnietfparamsscimschemasoracleidcsextensionx509identity_provider"] = []interface{}{ExtensionX509IdentityProviderToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionX509IdentityProvider)}
	}

	result["user_mapping_method"] = string(obj.UserMappingMethod)

	if obj.UserMappingStoreAttribute != nil {
		result["user_mapping_store_attribute"] = string(*obj.UserMappingStoreAttribute)
	}

	return result
}

func (s *IdentityDomainsIdentityProviderResourceCrud) mapToIdentityProviderCorrelationPolicy(fieldKeyFormat string) (oci_identity_domains.IdentityProviderCorrelationPolicy, error) {
	result := oci_identity_domains.IdentityProviderCorrelationPolicy{}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_identity_domains.IdentityProviderCorrelationPolicyTypeEnum(type_.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func IdentityProviderCorrelationPolicyToMap(obj *oci_identity_domains.IdentityProviderCorrelationPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsIdentityProviderResourceCrud) mapToIdentityProviderJitUserProvAssignedGroups(fieldKeyFormat string) (oci_identity_domains.IdentityProviderJitUserProvAssignedGroups, error) {
	result := oci_identity_domains.IdentityProviderJitUserProvAssignedGroups{}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
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

func IdentityProviderJitUserProvAssignedGroupsToMap(obj oci_identity_domains.IdentityProviderJitUserProvAssignedGroups) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsIdentityProviderResourceCrud) mapToIdentityProviderJitUserProvAttributes(fieldKeyFormat string) (oci_identity_domains.IdentityProviderJitUserProvAttributes, error) {
	result := oci_identity_domains.IdentityProviderJitUserProvAttributes{}

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

func IdentityProviderJitUserProvAttributesToMap(obj *oci_identity_domains.IdentityProviderJitUserProvAttributes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsIdentityProviderResourceCrud) mapToIdentityProviderJitUserProvGroupMappings(fieldKeyFormat string) (oci_identity_domains.IdentityProviderJitUserProvGroupMappings, error) {
	result := oci_identity_domains.IdentityProviderJitUserProvGroupMappings{}

	if idpGroup, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "idp_group")); ok {
		tmp := idpGroup.(string)
		result.IdpGroup = &tmp
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

func IdentityProviderJitUserProvGroupMappingsToMap(obj oci_identity_domains.IdentityProviderJitUserProvGroupMappings) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IdpGroup != nil {
		result["idp_group"] = string(*obj.IdpGroup)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsIdentityProviderResourceCrud) mapTotags(fieldKeyFormat string) (oci_identity_domains.Tags, error) {
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
