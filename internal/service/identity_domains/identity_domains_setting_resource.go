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

func IdentityDomainsSettingResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityDomainsSetting,
		Read:     readIdentityDomainsSetting,
		Update:   updateIdentityDomainsSetting,
		Delete:   deleteIdentityDomainsSetting,
		Schema: map[string]*schema.Schema{
			// Required
			"csr_access": {
				Type:     schema.TypeString,
				Required: true,
			},
			"idcs_endpoint": {
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
			"setting_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"account_always_trust_scope": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"allowed_domains": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"allowed_forgot_password_flow_return_urls": {
				Type:      schema.TypeList,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"allowed_notification_redirect_urls": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
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
			"audit_event_retention_period": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"authorization": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"certificate_validation": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

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
						"crl_refresh_interval": {
							Type:     schema.TypeInt,
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
						"ocsp_settings_responder_url_preferred": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"ocsp_signing_certificate_alias": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ocsp_timeout_duration": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ocsp_unknown_response_status_allowed": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"cloud_gate_cors_settings": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"cloud_gate_cors_allow_null_origin": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"cloud_gate_cors_allowed_origins": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"cloud_gate_cors_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"cloud_gate_cors_exposed_headers": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"cloud_gate_cors_max_age": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"cloud_migration_custom_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cloud_migration_url_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"company_names": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"locale": {
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
			"contact_emails": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"custom_branding": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"custom_css_location": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"custom_html_location": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"custom_translation": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_trust_scope": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diagnostic_level": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"diagnostic_record_for_search_identifies_returned_resources": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"enable_terms_of_use": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"iam_upst_session_expiry": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"images": {
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
					},
				},
			},
			"is_hosted_page": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"issuer": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"locale": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"login_texts": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"locale": {
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
			"max_no_of_app_cmva_to_return": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_no_of_app_role_members_to_return": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ocid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preferred_language": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prev_issuer": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"privacy_policy_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"purge_configs": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"resource_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"retention_period": {
							Type:     schema.TypeInt,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"re_auth_factor": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"re_auth_when_changing_my_authentication_factors": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"resource_type_schema_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_admin_cannot_list_other_users": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"signing_cert_public_access": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"sub_mapping_attr": {
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
			"tenant_custom_claims": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"all_scopes": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"expression": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"mode": {
							Type:     schema.TypeString,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"token_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"scopes": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Computed
					},
				},
			},
			"terms_of_use_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timezone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"cloud_account_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_company_names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"locale": {
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
			"default_images": {
				Type:     schema.TypeList,
				Computed: true,
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
					},
				},
			},
			"default_login_texts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"locale": {
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
			"delete_in_progress": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"diagnostic_tracing_upto": {
				Type:     schema.TypeString,
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
			"migration_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"on_premises_provisioning": {
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

func createIdentityDomainsSetting(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsSettingResourceCrud{}
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

func readIdentityDomainsSetting(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsSettingResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpointForRead(d, "settings")
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

func updateIdentityDomainsSetting(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsSettingResourceCrud{}
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

func deleteIdentityDomainsSetting(d *schema.ResourceData, m interface{}) error {
	return nil
}

type IdentityDomainsSettingResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity_domains.IdentityDomainsClient
	Res                    *oci_identity_domains.Setting
	DisableNotFoundRetries bool
}

func (s *IdentityDomainsSettingResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentityDomainsSettingResourceCrud) Create() error {
	request := oci_identity_domains.PutSettingRequest{}

	if accountAlwaysTrustScope, ok := s.D.GetOkExists("account_always_trust_scope"); ok {
		tmp := accountAlwaysTrustScope.(bool)
		request.AccountAlwaysTrustScope = &tmp
	}

	if allowedDomains, ok := s.D.GetOkExists("allowed_domains"); ok {
		interfaces := allowedDomains.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("allowed_domains") {
			request.AllowedDomains = tmp
		}
	}

	if allowedForgotPasswordFlowReturnUrls, ok := s.D.GetOkExists("allowed_forgot_password_flow_return_urls"); ok {
		interfaces := allowedForgotPasswordFlowReturnUrls.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("allowed_forgot_password_flow_return_urls") {
			request.AllowedForgotPasswordFlowReturnUrls = tmp
		}
	}

	if allowedNotificationRedirectUrls, ok := s.D.GetOkExists("allowed_notification_redirect_urls"); ok {
		interfaces := allowedNotificationRedirectUrls.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("allowed_notification_redirect_urls") {
			request.AllowedNotificationRedirectUrls = tmp
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

	if auditEventRetentionPeriod, ok := s.D.GetOkExists("audit_event_retention_period"); ok {
		// enum of type integer
		tmp := auditEventRetentionPeriod.(int)
		request.AuditEventRetentionPeriod = &tmp
	}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if certificateValidation, ok := s.D.GetOkExists("certificate_validation"); ok {
		if tmpList := certificateValidation.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "certificate_validation", 0)
			tmp, err := s.mapToSettingsCertificateValidation(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CertificateValidation = &tmp
		}
	}

	if cloudGateCorsSettings, ok := s.D.GetOkExists("cloud_gate_cors_settings"); ok {
		if tmpList := cloudGateCorsSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "cloud_gate_cors_settings", 0)
			tmp, err := s.mapToSettingsCloudGateCorsSettings(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CloudGateCorsSettings = &tmp
		}
	}

	if cloudMigrationCustomUrl, ok := s.D.GetOkExists("cloud_migration_custom_url"); ok {
		tmp := cloudMigrationCustomUrl.(string)
		request.CloudMigrationCustomUrl = &tmp
	}

	if cloudMigrationUrlEnabled, ok := s.D.GetOkExists("cloud_migration_url_enabled"); ok {
		tmp := cloudMigrationUrlEnabled.(bool)
		request.CloudMigrationUrlEnabled = &tmp
	}

	if companyNames, ok := s.D.GetOkExists("company_names"); ok {
		interfaces := companyNames.([]interface{})
		tmp := make([]oci_identity_domains.SettingsCompanyNames, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "company_names", stateDataIndex)
			converted, err := s.mapToSettingsCompanyNames(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("company_names") {
			request.CompanyNames = tmp
		}
	}

	if contactEmails, ok := s.D.GetOkExists("contact_emails"); ok {
		interfaces := contactEmails.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("contact_emails") {
			request.ContactEmails = tmp
		}
	}

	if csrAccess, ok := s.D.GetOkExists("csr_access"); ok {
		request.CsrAccess = oci_identity_domains.SettingCsrAccessEnum(csrAccess.(string))
	}

	if customBranding, ok := s.D.GetOkExists("custom_branding"); ok {
		tmp := customBranding.(bool)
		request.CustomBranding = &tmp
	}

	if customCssLocation, ok := s.D.GetOkExists("custom_css_location"); ok {
		tmp := customCssLocation.(string)
		request.CustomCssLocation = &tmp
	}

	if customHtmlLocation, ok := s.D.GetOkExists("custom_html_location"); ok {
		tmp := customHtmlLocation.(string)
		request.CustomHtmlLocation = &tmp
	}

	if customTranslation, ok := s.D.GetOkExists("custom_translation"); ok {
		tmp := customTranslation.(string)
		request.CustomTranslation = &tmp
	}

	if defaultTrustScope, ok := s.D.GetOkExists("default_trust_scope"); ok {
		request.DefaultTrustScope = oci_identity_domains.SettingDefaultTrustScopeEnum(defaultTrustScope.(string))
	}

	if diagnosticLevel, ok := s.D.GetOkExists("diagnostic_level"); ok {
		tmp := diagnosticLevel.(int)
		request.DiagnosticLevel = &tmp
	}

	if diagnosticRecordForSearchIdentifiesReturnedResources, ok := s.D.GetOkExists("diagnostic_record_for_search_identifies_returned_resources"); ok {
		tmp := diagnosticRecordForSearchIdentifiesReturnedResources.(bool)
		request.DiagnosticRecordForSearchIdentifiesReturnedResources = &tmp
	}

	if enableTermsOfUse, ok := s.D.GetOkExists("enable_terms_of_use"); ok {
		tmp := enableTermsOfUse.(bool)
		request.EnableTermsOfUse = &tmp
	}

	if externalId, ok := s.D.GetOkExists("external_id"); ok {
		tmp := externalId.(string)
		request.ExternalId = &tmp
	}

	if iamUpstSessionExpiry, ok := s.D.GetOkExists("iam_upst_session_expiry"); ok {
		tmp := iamUpstSessionExpiry.(int)
		request.IamUpstSessionExpiry = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if images, ok := s.D.GetOkExists("images"); ok {
		interfaces := images.([]interface{})
		tmp := make([]oci_identity_domains.SettingsImages, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "images", stateDataIndex)
			converted, err := s.mapToSettingsImages(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("images") {
			request.Images = tmp
		}
	}

	if isHostedPage, ok := s.D.GetOkExists("is_hosted_page"); ok {
		tmp := isHostedPage.(bool)
		request.IsHostedPage = &tmp
	}

	if issuer, ok := s.D.GetOkExists("issuer"); ok {
		tmp := issuer.(string)
		request.Issuer = &tmp
	}

	if locale, ok := s.D.GetOkExists("locale"); ok {
		tmp := locale.(string)
		request.Locale = &tmp
	}

	if loginTexts, ok := s.D.GetOkExists("login_texts"); ok {
		interfaces := loginTexts.([]interface{})
		tmp := make([]oci_identity_domains.SettingsLoginTexts, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "login_texts", stateDataIndex)
			converted, err := s.mapToSettingsLoginTexts(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("login_texts") {
			request.LoginTexts = tmp
		}
	}

	if maxNoOfAppCMVAToReturn, ok := s.D.GetOkExists("max_no_of_app_cmva_to_return"); ok {
		tmp := maxNoOfAppCMVAToReturn.(int)
		request.MaxNoOfAppCMVAToReturn = &tmp
	}

	if maxNoOfAppRoleMembersToReturn, ok := s.D.GetOkExists("max_no_of_app_role_members_to_return"); ok {
		tmp := maxNoOfAppRoleMembersToReturn.(int)
		request.MaxNoOfAppRoleMembersToReturn = &tmp
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if preferredLanguage, ok := s.D.GetOkExists("preferred_language"); ok {
		tmp := preferredLanguage.(string)
		request.PreferredLanguage = &tmp
	}

	if prevIssuer, ok := s.D.GetOkExists("prev_issuer"); ok {
		tmp := prevIssuer.(string)
		request.PrevIssuer = &tmp
	}

	if privacyPolicyUrl, ok := s.D.GetOkExists("privacy_policy_url"); ok {
		tmp := privacyPolicyUrl.(string)
		request.PrivacyPolicyUrl = &tmp
	}

	if purgeConfigs, ok := s.D.GetOkExists("purge_configs"); ok {
		interfaces := purgeConfigs.([]interface{})
		tmp := make([]oci_identity_domains.SettingsPurgeConfigs, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "purge_configs", stateDataIndex)
			converted, err := s.mapToSettingsPurgeConfigs(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("purge_configs") {
			request.PurgeConfigs = tmp
		}
	}

	if reAuthFactor, ok := s.D.GetOkExists("re_auth_factor"); ok {
		interfaces := reAuthFactor.([]interface{})
		tmp := make([]oci_identity_domains.SettingReAuthFactorEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_identity_domains.SettingReAuthFactorEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("re_auth_factor") {
			request.ReAuthFactor = tmp
		}
	}

	if reAuthWhenChangingMyAuthenticationFactors, ok := s.D.GetOkExists("re_auth_when_changing_my_authentication_factors"); ok {
		tmp := reAuthWhenChangingMyAuthenticationFactors.(bool)
		request.ReAuthWhenChangingMyAuthenticationFactors = &tmp
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

	if serviceAdminCannotListOtherUsers, ok := s.D.GetOkExists("service_admin_cannot_list_other_users"); ok {
		tmp := serviceAdminCannotListOtherUsers.(bool)
		request.ServiceAdminCannotListOtherUsers = &tmp
	}

	if settingId, ok := s.D.GetOkExists("setting_id"); ok {
		tmp := settingId.(string)
		request.SettingId = &tmp
	}

	if signingCertPublicAccess, ok := s.D.GetOkExists("signing_cert_public_access"); ok {
		tmp := signingCertPublicAccess.(bool)
		request.SigningCertPublicAccess = &tmp
	}

	if subMappingAttr, ok := s.D.GetOkExists("sub_mapping_attr"); ok {
		tmp := subMappingAttr.(string)
		request.SubMappingAttr = &tmp
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

	if tenantCustomClaims, ok := s.D.GetOkExists("tenant_custom_claims"); ok {
		interfaces := tenantCustomClaims.([]interface{})
		tmp := make([]oci_identity_domains.SettingsTenantCustomClaims, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tenant_custom_claims", stateDataIndex)
			converted, err := s.mapToSettingsTenantCustomClaims(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("tenant_custom_claims") {
			request.TenantCustomClaims = tmp
		}
	}

	if termsOfUseUrl, ok := s.D.GetOkExists("terms_of_use_url"); ok {
		tmp := termsOfUseUrl.(string)
		request.TermsOfUseUrl = &tmp
	}

	if timezone, ok := s.D.GetOkExists("timezone"); ok {
		tmp := timezone.(string)
		request.Timezone = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.PutSetting(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Setting
	return nil
}

func (s *IdentityDomainsSettingResourceCrud) Get() error {
	request := oci_identity_domains.GetSettingRequest{}

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

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	tmp := s.D.Id()
	request.SettingId = &tmp

	settingId, err := parseSettingCompositeId(s.D.Id())
	if err == nil {
		request.SettingId = &settingId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.GetSetting(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Setting
	return nil
}

func (s *IdentityDomainsSettingResourceCrud) Update() error {
	request := oci_identity_domains.PutSettingRequest{}

	if accountAlwaysTrustScope, ok := s.D.GetOkExists("account_always_trust_scope"); ok {
		tmp := accountAlwaysTrustScope.(bool)
		request.AccountAlwaysTrustScope = &tmp
	}

	if allowedDomains, ok := s.D.GetOkExists("allowed_domains"); ok {
		interfaces := allowedDomains.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("allowed_domains") {
			request.AllowedDomains = tmp
		}
	}

	if allowedForgotPasswordFlowReturnUrls, ok := s.D.GetOkExists("allowed_forgot_password_flow_return_urls"); ok {
		interfaces := allowedForgotPasswordFlowReturnUrls.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("allowed_forgot_password_flow_return_urls") {
			request.AllowedForgotPasswordFlowReturnUrls = tmp
		}
	}

	if allowedNotificationRedirectUrls, ok := s.D.GetOkExists("allowed_notification_redirect_urls"); ok {
		interfaces := allowedNotificationRedirectUrls.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("allowed_notification_redirect_urls") {
			request.AllowedNotificationRedirectUrls = tmp
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

	if auditEventRetentionPeriod, ok := s.D.GetOkExists("audit_event_retention_period"); ok {
		// enum of type integer
		tmp := auditEventRetentionPeriod.(int)
		request.AuditEventRetentionPeriod = &tmp
	}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if certificateValidation, ok := s.D.GetOkExists("certificate_validation"); ok {
		if tmpList := certificateValidation.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "certificate_validation", 0)
			tmp, err := s.mapToSettingsCertificateValidation(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CertificateValidation = &tmp
		}
	}

	if cloudGateCorsSettings, ok := s.D.GetOkExists("cloud_gate_cors_settings"); ok {
		if tmpList := cloudGateCorsSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "cloud_gate_cors_settings", 0)
			tmp, err := s.mapToSettingsCloudGateCorsSettings(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CloudGateCorsSettings = &tmp
		}
	}

	if cloudMigrationCustomUrl, ok := s.D.GetOkExists("cloud_migration_custom_url"); ok {
		tmp := cloudMigrationCustomUrl.(string)
		request.CloudMigrationCustomUrl = &tmp
	}

	if cloudMigrationUrlEnabled, ok := s.D.GetOkExists("cloud_migration_url_enabled"); ok {
		tmp := cloudMigrationUrlEnabled.(bool)
		request.CloudMigrationUrlEnabled = &tmp
	}

	if companyNames, ok := s.D.GetOkExists("company_names"); ok {
		interfaces := companyNames.([]interface{})
		tmp := make([]oci_identity_domains.SettingsCompanyNames, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "company_names", stateDataIndex)
			converted, err := s.mapToSettingsCompanyNames(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("company_names") {
			request.CompanyNames = tmp
		}
	}

	if contactEmails, ok := s.D.GetOkExists("contact_emails"); ok {
		interfaces := contactEmails.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("contact_emails") {
			request.ContactEmails = tmp
		}
	}

	if csrAccess, ok := s.D.GetOkExists("csr_access"); ok {
		request.CsrAccess = oci_identity_domains.SettingCsrAccessEnum(csrAccess.(string))
	}

	if customBranding, ok := s.D.GetOkExists("custom_branding"); ok {
		tmp := customBranding.(bool)
		request.CustomBranding = &tmp
	}

	if customCssLocation, ok := s.D.GetOkExists("custom_css_location"); ok {
		tmp := customCssLocation.(string)
		request.CustomCssLocation = &tmp
	}

	if customHtmlLocation, ok := s.D.GetOkExists("custom_html_location"); ok {
		tmp := customHtmlLocation.(string)
		request.CustomHtmlLocation = &tmp
	}

	if customTranslation, ok := s.D.GetOkExists("custom_translation"); ok {
		tmp := customTranslation.(string)
		request.CustomTranslation = &tmp
	}

	if defaultTrustScope, ok := s.D.GetOkExists("default_trust_scope"); ok {
		request.DefaultTrustScope = oci_identity_domains.SettingDefaultTrustScopeEnum(defaultTrustScope.(string))
	}

	if diagnosticLevel, ok := s.D.GetOkExists("diagnostic_level"); ok {
		tmp := diagnosticLevel.(int)
		request.DiagnosticLevel = &tmp
	}

	if diagnosticRecordForSearchIdentifiesReturnedResources, ok := s.D.GetOkExists("diagnostic_record_for_search_identifies_returned_resources"); ok {
		tmp := diagnosticRecordForSearchIdentifiesReturnedResources.(bool)
		request.DiagnosticRecordForSearchIdentifiesReturnedResources = &tmp
	}

	if enableTermsOfUse, ok := s.D.GetOkExists("enable_terms_of_use"); ok {
		tmp := enableTermsOfUse.(bool)
		request.EnableTermsOfUse = &tmp
	}

	if externalId, ok := s.D.GetOkExists("external_id"); ok {
		tmp := externalId.(string)
		request.ExternalId = &tmp
	}

	if iamUpstSessionExpiry, ok := s.D.GetOkExists("iam_upst_session_expiry"); ok {
		tmp := iamUpstSessionExpiry.(int)
		request.IamUpstSessionExpiry = &tmp
	}

	tmp := s.D.Id()
	request.Id = &tmp

	if images, ok := s.D.GetOkExists("images"); ok {
		interfaces := images.([]interface{})
		tmp := make([]oci_identity_domains.SettingsImages, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "images", stateDataIndex)
			converted, err := s.mapToSettingsImages(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("images") {
			request.Images = tmp
		}
	}

	if isHostedPage, ok := s.D.GetOkExists("is_hosted_page"); ok {
		tmp := isHostedPage.(bool)
		request.IsHostedPage = &tmp
	}

	if issuer, ok := s.D.GetOkExists("issuer"); ok {
		tmp := issuer.(string)
		request.Issuer = &tmp
	}

	if locale, ok := s.D.GetOkExists("locale"); ok {
		tmp := locale.(string)
		request.Locale = &tmp
	}

	if loginTexts, ok := s.D.GetOkExists("login_texts"); ok {
		interfaces := loginTexts.([]interface{})
		tmp := make([]oci_identity_domains.SettingsLoginTexts, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "login_texts", stateDataIndex)
			converted, err := s.mapToSettingsLoginTexts(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("login_texts") {
			request.LoginTexts = tmp
		}
	}

	if maxNoOfAppCMVAToReturn, ok := s.D.GetOkExists("max_no_of_app_cmva_to_return"); ok {
		tmp := maxNoOfAppCMVAToReturn.(int)
		request.MaxNoOfAppCMVAToReturn = &tmp
	}

	if maxNoOfAppRoleMembersToReturn, ok := s.D.GetOkExists("max_no_of_app_role_members_to_return"); ok {
		tmp := maxNoOfAppRoleMembersToReturn.(int)
		request.MaxNoOfAppRoleMembersToReturn = &tmp
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if preferredLanguage, ok := s.D.GetOkExists("preferred_language"); ok {
		tmp := preferredLanguage.(string)
		request.PreferredLanguage = &tmp
	}

	if prevIssuer, ok := s.D.GetOkExists("prev_issuer"); ok {
		tmp := prevIssuer.(string)
		request.PrevIssuer = &tmp
	}

	if privacyPolicyUrl, ok := s.D.GetOkExists("privacy_policy_url"); ok {
		tmp := privacyPolicyUrl.(string)
		request.PrivacyPolicyUrl = &tmp
	}

	if purgeConfigs, ok := s.D.GetOkExists("purge_configs"); ok {
		interfaces := purgeConfigs.([]interface{})
		tmp := make([]oci_identity_domains.SettingsPurgeConfigs, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "purge_configs", stateDataIndex)
			converted, err := s.mapToSettingsPurgeConfigs(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("purge_configs") {
			request.PurgeConfigs = tmp
		}
	}

	if reAuthFactor, ok := s.D.GetOkExists("re_auth_factor"); ok {
		interfaces := reAuthFactor.([]interface{})
		tmp := make([]oci_identity_domains.SettingReAuthFactorEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_identity_domains.SettingReAuthFactorEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("re_auth_factor") {
			request.ReAuthFactor = tmp
		}
	}

	if reAuthWhenChangingMyAuthenticationFactors, ok := s.D.GetOkExists("re_auth_when_changing_my_authentication_factors"); ok {
		tmp := reAuthWhenChangingMyAuthenticationFactors.(bool)
		request.ReAuthWhenChangingMyAuthenticationFactors = &tmp
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

	if serviceAdminCannotListOtherUsers, ok := s.D.GetOkExists("service_admin_cannot_list_other_users"); ok {
		tmp := serviceAdminCannotListOtherUsers.(bool)
		request.ServiceAdminCannotListOtherUsers = &tmp
	}

	tmp = s.D.Id()
	request.SettingId = &tmp

	if signingCertPublicAccess, ok := s.D.GetOkExists("signing_cert_public_access"); ok {
		tmp := signingCertPublicAccess.(bool)
		request.SigningCertPublicAccess = &tmp
	}

	if subMappingAttr, ok := s.D.GetOkExists("sub_mapping_attr"); ok {
		tmp := subMappingAttr.(string)
		request.SubMappingAttr = &tmp
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

	if tenantCustomClaims, ok := s.D.GetOkExists("tenant_custom_claims"); ok {
		interfaces := tenantCustomClaims.([]interface{})
		tmp := make([]oci_identity_domains.SettingsTenantCustomClaims, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tenant_custom_claims", stateDataIndex)
			converted, err := s.mapToSettingsTenantCustomClaims(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("tenant_custom_claims") {
			request.TenantCustomClaims = tmp
		}
	}

	if termsOfUseUrl, ok := s.D.GetOkExists("terms_of_use_url"); ok {
		tmp := termsOfUseUrl.(string)
		request.TermsOfUseUrl = &tmp
	}

	if timezone, ok := s.D.GetOkExists("timezone"); ok {
		tmp := timezone.(string)
		request.Timezone = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.PutSetting(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Setting
	return nil
}

func (s *IdentityDomainsSettingResourceCrud) SetData() error {

	settingId, err := parseSettingCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(settingId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.AccountAlwaysTrustScope != nil {
		s.D.Set("account_always_trust_scope", *s.Res.AccountAlwaysTrustScope)
	}

	s.D.Set("allowed_domains", s.Res.AllowedDomains)

	s.D.Set("allowed_forgot_password_flow_return_urls", s.Res.AllowedForgotPasswordFlowReturnUrls)

	s.D.Set("allowed_notification_redirect_urls", s.Res.AllowedNotificationRedirectUrls)

	s.D.Set("audit_event_retention_period", s.Res.AuditEventRetentionPeriod)

	if s.Res.CertificateValidation != nil {
		s.D.Set("certificate_validation", []interface{}{SettingsCertificateValidationToMap(s.Res.CertificateValidation)})
	} else {
		s.D.Set("certificate_validation", nil)
	}

	if s.Res.CloudAccountName != nil {
		s.D.Set("cloud_account_name", *s.Res.CloudAccountName)
	}

	if s.Res.CloudGateCorsSettings != nil {
		s.D.Set("cloud_gate_cors_settings", []interface{}{SettingsCloudGateCorsSettingsToMap(s.Res.CloudGateCorsSettings)})
	} else {
		s.D.Set("cloud_gate_cors_settings", nil)
	}

	if s.Res.CloudMigrationCustomUrl != nil {
		s.D.Set("cloud_migration_custom_url", *s.Res.CloudMigrationCustomUrl)
	}

	if s.Res.CloudMigrationUrlEnabled != nil {
		s.D.Set("cloud_migration_url_enabled", *s.Res.CloudMigrationUrlEnabled)
	}

	companyNames := []interface{}{}
	for _, item := range s.Res.CompanyNames {
		companyNames = append(companyNames, SettingsCompanyNamesToMap(item))
	}
	s.D.Set("company_names", companyNames)

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
	}

	s.D.Set("contact_emails", s.Res.ContactEmails)

	s.D.Set("csr_access", s.Res.CsrAccess)

	if s.Res.CustomBranding != nil {
		s.D.Set("custom_branding", *s.Res.CustomBranding)
	}

	if s.Res.CustomCssLocation != nil {
		s.D.Set("custom_css_location", *s.Res.CustomCssLocation)
	}

	if s.Res.CustomHtmlLocation != nil {
		s.D.Set("custom_html_location", *s.Res.CustomHtmlLocation)
	}

	if s.Res.CustomTranslation != nil {
		s.D.Set("custom_translation", *s.Res.CustomTranslation)
	}

	defaultCompanyNames := []interface{}{}
	for _, item := range s.Res.DefaultCompanyNames {
		defaultCompanyNames = append(defaultCompanyNames, SettingsDefaultCompanyNamesToMap(item))
	}
	s.D.Set("default_company_names", defaultCompanyNames)

	defaultImages := []interface{}{}
	for _, item := range s.Res.DefaultImages {
		defaultImages = append(defaultImages, SettingsDefaultImagesToMap(item))
	}
	s.D.Set("default_images", defaultImages)

	defaultLoginTexts := []interface{}{}
	for _, item := range s.Res.DefaultLoginTexts {
		defaultLoginTexts = append(defaultLoginTexts, SettingsDefaultLoginTextsToMap(item))
	}
	s.D.Set("default_login_texts", defaultLoginTexts)

	s.D.Set("default_trust_scope", s.Res.DefaultTrustScope)

	if s.Res.DeleteInProgress != nil {
		s.D.Set("delete_in_progress", *s.Res.DeleteInProgress)
	}

	if s.Res.DiagnosticLevel != nil {
		s.D.Set("diagnostic_level", *s.Res.DiagnosticLevel)
	}

	if s.Res.DiagnosticRecordForSearchIdentifiesReturnedResources != nil {
		s.D.Set("diagnostic_record_for_search_identifies_returned_resources", *s.Res.DiagnosticRecordForSearchIdentifiesReturnedResources)
	}

	if s.Res.DiagnosticTracingUpto != nil {
		s.D.Set("diagnostic_tracing_upto", *s.Res.DiagnosticTracingUpto)
	}

	if s.Res.DomainOcid != nil {
		s.D.Set("domain_ocid", *s.Res.DomainOcid)
	}

	if s.Res.EnableTermsOfUse != nil {
		s.D.Set("enable_terms_of_use", *s.Res.EnableTermsOfUse)
	}

	if s.Res.ExternalId != nil {
		s.D.Set("external_id", *s.Res.ExternalId)
	}

	if s.Res.IamUpstSessionExpiry != nil {
		s.D.Set("iam_upst_session_expiry", *s.Res.IamUpstSessionExpiry)
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

	images := []interface{}{}
	for _, item := range s.Res.Images {
		images = append(images, SettingsImagesToMap(item))
	}
	s.D.Set("images", images)

	if s.Res.IsHostedPage != nil {
		s.D.Set("is_hosted_page", *s.Res.IsHostedPage)
	}

	if s.Res.Issuer != nil {
		s.D.Set("issuer", *s.Res.Issuer)
	}

	if s.Res.Locale != nil {
		s.D.Set("locale", *s.Res.Locale)
	}

	loginTexts := []interface{}{}
	for _, item := range s.Res.LoginTexts {
		loginTexts = append(loginTexts, SettingsLoginTextsToMap(item))
	}
	s.D.Set("login_texts", loginTexts)

	if s.Res.MaxNoOfAppCMVAToReturn != nil {
		s.D.Set("max_no_of_app_cmva_to_return", *s.Res.MaxNoOfAppCMVAToReturn)
	}

	if s.Res.MaxNoOfAppRoleMembersToReturn != nil {
		s.D.Set("max_no_of_app_role_members_to_return", *s.Res.MaxNoOfAppRoleMembersToReturn)
	}

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	if s.Res.MigrationStatus != nil {
		s.D.Set("migration_status", *s.Res.MigrationStatus)
	}

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
	}

	if s.Res.OnPremisesProvisioning != nil {
		s.D.Set("on_premises_provisioning", *s.Res.OnPremisesProvisioning)
	}

	if s.Res.PreferredLanguage != nil {
		s.D.Set("preferred_language", *s.Res.PreferredLanguage)
	}

	if s.Res.PrevIssuer != nil {
		s.D.Set("prev_issuer", *s.Res.PrevIssuer)
	}

	if s.Res.PrivacyPolicyUrl != nil {
		s.D.Set("privacy_policy_url", *s.Res.PrivacyPolicyUrl)
	}

	purgeConfigs := []interface{}{}
	for _, item := range s.Res.PurgeConfigs {
		purgeConfigs = append(purgeConfigs, SettingsPurgeConfigsToMap(item))
	}
	s.D.Set("purge_configs", purgeConfigs)

	s.D.Set("re_auth_factor", s.Res.ReAuthFactor)

	if s.Res.ReAuthWhenChangingMyAuthenticationFactors != nil {
		s.D.Set("re_auth_when_changing_my_authentication_factors", *s.Res.ReAuthWhenChangingMyAuthenticationFactors)
	}

	s.D.Set("schemas", s.Res.Schemas)

	if s.Res.ServiceAdminCannotListOtherUsers != nil {
		s.D.Set("service_admin_cannot_list_other_users", *s.Res.ServiceAdminCannotListOtherUsers)
	}

	if s.Res.SigningCertPublicAccess != nil {
		s.D.Set("signing_cert_public_access", *s.Res.SigningCertPublicAccess)
	}

	if s.Res.SubMappingAttr != nil {
		s.D.Set("sub_mapping_attr", *s.Res.SubMappingAttr)
	}

	tags := []interface{}{}
	for _, item := range s.Res.Tags {
		tags = append(tags, tagsToMap(item))
	}
	s.D.Set("tags", tags)

	if s.Res.TenancyOcid != nil {
		s.D.Set("tenancy_ocid", *s.Res.TenancyOcid)
	}

	tenantCustomClaims := []interface{}{}
	for _, item := range s.Res.TenantCustomClaims {
		tenantCustomClaims = append(tenantCustomClaims, SettingsTenantCustomClaimsToMap(item))
	}
	s.D.Set("tenant_custom_claims", tenantCustomClaims)

	if s.Res.TermsOfUseUrl != nil {
		s.D.Set("terms_of_use_url", *s.Res.TermsOfUseUrl)
	}

	if s.Res.Timezone != nil {
		s.D.Set("timezone", *s.Res.Timezone)
	}

	return nil
}

func parseSettingCompositeId(compositeId string) (settingId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("idcsEndpoint/.*/settings/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}

	settingId, _ = url.PathUnescape(parts[3])

	return
}

func SettingToMap(obj oci_identity_domains.Setting) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AccountAlwaysTrustScope != nil {
		result["account_always_trust_scope"] = bool(*obj.AccountAlwaysTrustScope)
	}

	result["allowed_domains"] = obj.AllowedDomains

	result["allowed_forgot_password_flow_return_urls"] = obj.AllowedForgotPasswordFlowReturnUrls

	result["allowed_notification_redirect_urls"] = obj.AllowedNotificationRedirectUrls

	if obj.AuditEventRetentionPeriod != nil {
		result["audit_event_retention_period"] = int(*obj.AuditEventRetentionPeriod)
	}

	if obj.CertificateValidation != nil {
		result["certificate_validation"] = []interface{}{SettingsCertificateValidationToMap(obj.CertificateValidation)}
	}

	if obj.CloudAccountName != nil {
		result["cloud_account_name"] = string(*obj.CloudAccountName)
	}

	if obj.CloudGateCorsSettings != nil {
		result["cloud_gate_cors_settings"] = []interface{}{SettingsCloudGateCorsSettingsToMap(obj.CloudGateCorsSettings)}
	}

	if obj.CloudMigrationCustomUrl != nil {
		result["cloud_migration_custom_url"] = string(*obj.CloudMigrationCustomUrl)
	}

	if obj.CloudMigrationUrlEnabled != nil {
		result["cloud_migration_url_enabled"] = bool(*obj.CloudMigrationUrlEnabled)
	}

	companyNames := []interface{}{}
	for _, item := range obj.CompanyNames {
		companyNames = append(companyNames, SettingsCompanyNamesToMap(item))
	}
	result["company_names"] = companyNames

	if obj.CompartmentOcid != nil {
		result["compartment_ocid"] = string(*obj.CompartmentOcid)
	}

	result["contact_emails"] = obj.ContactEmails

	result["csr_access"] = string(obj.CsrAccess)

	if obj.CustomBranding != nil {
		result["custom_branding"] = bool(*obj.CustomBranding)
	}

	if obj.CustomCssLocation != nil {
		result["custom_css_location"] = string(*obj.CustomCssLocation)
	}

	if obj.CustomHtmlLocation != nil {
		result["custom_html_location"] = string(*obj.CustomHtmlLocation)
	}

	if obj.CustomTranslation != nil {
		result["custom_translation"] = string(*obj.CustomTranslation)
	}

	defaultCompanyNames := []interface{}{}
	for _, item := range obj.DefaultCompanyNames {
		defaultCompanyNames = append(defaultCompanyNames, SettingsDefaultCompanyNamesToMap(item))
	}
	result["default_company_names"] = defaultCompanyNames

	defaultImages := []interface{}{}
	for _, item := range obj.DefaultImages {
		defaultImages = append(defaultImages, SettingsDefaultImagesToMap(item))
	}
	result["default_images"] = defaultImages

	defaultLoginTexts := []interface{}{}
	for _, item := range obj.DefaultLoginTexts {
		defaultLoginTexts = append(defaultLoginTexts, SettingsDefaultLoginTextsToMap(item))
	}
	result["default_login_texts"] = defaultLoginTexts

	result["default_trust_scope"] = string(obj.DefaultTrustScope)

	if obj.DeleteInProgress != nil {
		result["delete_in_progress"] = bool(*obj.DeleteInProgress)
	}

	if obj.DiagnosticLevel != nil {
		result["diagnostic_level"] = int(*obj.DiagnosticLevel)
	}

	if obj.DiagnosticRecordForSearchIdentifiesReturnedResources != nil {
		result["diagnostic_record_for_search_identifies_returned_resources"] = bool(*obj.DiagnosticRecordForSearchIdentifiesReturnedResources)
	}

	if obj.DiagnosticTracingUpto != nil {
		result["diagnostic_tracing_upto"] = string(*obj.DiagnosticTracingUpto)
	}

	if obj.DomainOcid != nil {
		result["domain_ocid"] = string(*obj.DomainOcid)
	}

	if obj.EnableTermsOfUse != nil {
		result["enable_terms_of_use"] = bool(*obj.EnableTermsOfUse)
	}

	if obj.ExternalId != nil {
		result["external_id"] = string(*obj.ExternalId)
	}

	if obj.IamUpstSessionExpiry != nil {
		result["iam_upst_session_expiry"] = int(*obj.IamUpstSessionExpiry)
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

	images := []interface{}{}
	for _, item := range obj.Images {
		images = append(images, SettingsImagesToMap(item))
	}
	result["images"] = images

	if obj.IsHostedPage != nil {
		result["is_hosted_page"] = bool(*obj.IsHostedPage)
	}

	if obj.Issuer != nil {
		result["issuer"] = string(*obj.Issuer)
	}

	if obj.Locale != nil {
		result["locale"] = string(*obj.Locale)
	}

	loginTexts := []interface{}{}
	for _, item := range obj.LoginTexts {
		loginTexts = append(loginTexts, SettingsLoginTextsToMap(item))
	}
	result["login_texts"] = loginTexts

	if obj.MaxNoOfAppCMVAToReturn != nil {
		result["max_no_of_app_cmva_to_return"] = int(*obj.MaxNoOfAppCMVAToReturn)
	}

	if obj.MaxNoOfAppRoleMembersToReturn != nil {
		result["max_no_of_app_role_members_to_return"] = int(*obj.MaxNoOfAppRoleMembersToReturn)
	}

	if obj.Meta != nil {
		result["meta"] = []interface{}{metaToMap(obj.Meta)}
	}

	if obj.MigrationStatus != nil {
		result["migration_status"] = string(*obj.MigrationStatus)
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	if obj.OnPremisesProvisioning != nil {
		result["on_premises_provisioning"] = bool(*obj.OnPremisesProvisioning)
	}

	if obj.PreferredLanguage != nil {
		result["preferred_language"] = string(*obj.PreferredLanguage)
	}

	if obj.PrevIssuer != nil {
		result["prev_issuer"] = string(*obj.PrevIssuer)
	}

	if obj.PrivacyPolicyUrl != nil {
		result["privacy_policy_url"] = string(*obj.PrivacyPolicyUrl)
	}

	purgeConfigs := []interface{}{}
	for _, item := range obj.PurgeConfigs {
		purgeConfigs = append(purgeConfigs, SettingsPurgeConfigsToMap(item))
	}
	result["purge_configs"] = purgeConfigs

	result["re_auth_factor"] = obj.ReAuthFactor

	if obj.ReAuthWhenChangingMyAuthenticationFactors != nil {
		result["re_auth_when_changing_my_authentication_factors"] = bool(*obj.ReAuthWhenChangingMyAuthenticationFactors)
	}

	result["schemas"] = obj.Schemas

	if obj.ServiceAdminCannotListOtherUsers != nil {
		result["service_admin_cannot_list_other_users"] = bool(*obj.ServiceAdminCannotListOtherUsers)
	}

	if obj.SigningCertPublicAccess != nil {
		result["signing_cert_public_access"] = bool(*obj.SigningCertPublicAccess)
	}

	if obj.SubMappingAttr != nil {
		result["sub_mapping_attr"] = string(*obj.SubMappingAttr)
	}

	tags := []interface{}{}
	for _, item := range obj.Tags {
		tags = append(tags, tagsToMap(item))
	}
	result["tags"] = tags

	if obj.TenancyOcid != nil {
		result["tenancy_ocid"] = string(*obj.TenancyOcid)
	}

	tenantCustomClaims := []interface{}{}
	for _, item := range obj.TenantCustomClaims {
		tenantCustomClaims = append(tenantCustomClaims, SettingsTenantCustomClaimsToMap(item))
	}
	result["tenant_custom_claims"] = tenantCustomClaims

	if obj.TermsOfUseUrl != nil {
		result["terms_of_use_url"] = string(*obj.TermsOfUseUrl)
	}

	if obj.Timezone != nil {
		result["timezone"] = string(*obj.Timezone)
	}

	return result
}

func (s *IdentityDomainsSettingResourceCrud) mapToSettingsCertificateValidation(fieldKeyFormat string) (oci_identity_domains.SettingsCertificateValidation, error) {
	result := oci_identity_domains.SettingsCertificateValidation{}

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

	if crlRefreshInterval, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "crl_refresh_interval")); ok {
		tmp := crlRefreshInterval.(int)
		result.CrlRefreshInterval = &tmp
	}

	if ocspEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocsp_enabled")); ok {
		tmp := ocspEnabled.(bool)
		result.OcspEnabled = &tmp
	}

	if ocspResponderURL, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocsp_responder_url")); ok {
		tmp := ocspResponderURL.(string)
		result.OcspResponderURL = &tmp
	}

	if ocspSettingsResponderURLPreferred, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocsp_settings_responder_url_preferred")); ok {
		tmp := ocspSettingsResponderURLPreferred.(bool)
		result.OcspSettingsResponderURLPreferred = &tmp
	}

	if ocspSigningCertificateAlias, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocsp_signing_certificate_alias")); ok {
		tmp := ocspSigningCertificateAlias.(string)
		result.OcspSigningCertificateAlias = &tmp
	}

	if ocspTimeoutDuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocsp_timeout_duration")); ok {
		tmp := ocspTimeoutDuration.(int)
		result.OcspTimeoutDuration = &tmp
	}

	if ocspUnknownResponseStatusAllowed, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocsp_unknown_response_status_allowed")); ok {
		tmp := ocspUnknownResponseStatusAllowed.(bool)
		result.OcspUnknownResponseStatusAllowed = &tmp
	}

	return result, nil
}

func SettingsCertificateValidationToMap(obj *oci_identity_domains.SettingsCertificateValidation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CrlCheckOnOCSPFailureEnabled != nil {
		result["crl_check_on_ocsp_failure_enabled"] = bool(*obj.CrlCheckOnOCSPFailureEnabled)
	}

	if obj.CrlEnabled != nil {
		result["crl_enabled"] = bool(*obj.CrlEnabled)
	}

	if obj.CrlLocation != nil {
		result["crl_location"] = string(*obj.CrlLocation)
	}

	if obj.CrlRefreshInterval != nil {
		result["crl_refresh_interval"] = int(*obj.CrlRefreshInterval)
	}

	if obj.OcspEnabled != nil {
		result["ocsp_enabled"] = bool(*obj.OcspEnabled)
	}

	if obj.OcspResponderURL != nil {
		result["ocsp_responder_url"] = string(*obj.OcspResponderURL)
	}

	if obj.OcspSettingsResponderURLPreferred != nil {
		result["ocsp_settings_responder_url_preferred"] = bool(*obj.OcspSettingsResponderURLPreferred)
	}

	if obj.OcspSigningCertificateAlias != nil {
		result["ocsp_signing_certificate_alias"] = string(*obj.OcspSigningCertificateAlias)
	}

	if obj.OcspTimeoutDuration != nil {
		result["ocsp_timeout_duration"] = int(*obj.OcspTimeoutDuration)
	}

	if obj.OcspUnknownResponseStatusAllowed != nil {
		result["ocsp_unknown_response_status_allowed"] = bool(*obj.OcspUnknownResponseStatusAllowed)
	}

	return result
}

func (s *IdentityDomainsSettingResourceCrud) mapToSettingsCloudGateCorsSettings(fieldKeyFormat string) (oci_identity_domains.SettingsCloudGateCorsSettings, error) {
	result := oci_identity_domains.SettingsCloudGateCorsSettings{}

	if cloudGateCorsAllowNullOrigin, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cloud_gate_cors_allow_null_origin")); ok {
		tmp := cloudGateCorsAllowNullOrigin.(bool)
		result.CloudGateCorsAllowNullOrigin = &tmp
	}

	if cloudGateCorsAllowedOrigins, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cloud_gate_cors_allowed_origins")); ok {
		interfaces := cloudGateCorsAllowedOrigins.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "cloud_gate_cors_allowed_origins")) {
			result.CloudGateCorsAllowedOrigins = tmp
		}
	}

	if cloudGateCorsEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cloud_gate_cors_enabled")); ok {
		tmp := cloudGateCorsEnabled.(bool)
		result.CloudGateCorsEnabled = &tmp
	}

	if cloudGateCorsExposedHeaders, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cloud_gate_cors_exposed_headers")); ok {
		interfaces := cloudGateCorsExposedHeaders.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "cloud_gate_cors_exposed_headers")) {
			result.CloudGateCorsExposedHeaders = tmp
		}
	}

	if cloudGateCorsMaxAge, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cloud_gate_cors_max_age")); ok {
		tmp := cloudGateCorsMaxAge.(int)
		result.CloudGateCorsMaxAge = &tmp
	}

	return result, nil
}

func SettingsCloudGateCorsSettingsToMap(obj *oci_identity_domains.SettingsCloudGateCorsSettings) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CloudGateCorsAllowNullOrigin != nil {
		result["cloud_gate_cors_allow_null_origin"] = bool(*obj.CloudGateCorsAllowNullOrigin)
	}

	result["cloud_gate_cors_allowed_origins"] = obj.CloudGateCorsAllowedOrigins

	if obj.CloudGateCorsEnabled != nil {
		result["cloud_gate_cors_enabled"] = bool(*obj.CloudGateCorsEnabled)
	}

	result["cloud_gate_cors_exposed_headers"] = obj.CloudGateCorsExposedHeaders

	if obj.CloudGateCorsMaxAge != nil {
		result["cloud_gate_cors_max_age"] = int(*obj.CloudGateCorsMaxAge)
	}

	return result
}

func (s *IdentityDomainsSettingResourceCrud) mapToSettingsCompanyNames(fieldKeyFormat string) (oci_identity_domains.SettingsCompanyNames, error) {
	result := oci_identity_domains.SettingsCompanyNames{}

	if locale, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "locale")); ok {
		tmp := locale.(string)
		result.Locale = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func SettingsCompanyNamesToMap(obj oci_identity_domains.SettingsCompanyNames) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Locale != nil {
		result["locale"] = string(*obj.Locale)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func SettingsDefaultCompanyNamesToMap(obj oci_identity_domains.SettingsDefaultCompanyNames) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Locale != nil {
		result["locale"] = string(*obj.Locale)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func SettingsDefaultImagesToMap(obj oci_identity_domains.SettingsDefaultImages) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func SettingsDefaultLoginTextsToMap(obj oci_identity_domains.SettingsDefaultLoginTexts) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Locale != nil {
		result["locale"] = string(*obj.Locale)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsSettingResourceCrud) mapToSettingsImages(fieldKeyFormat string) (oci_identity_domains.SettingsImages, error) {
	result := oci_identity_domains.SettingsImages{}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		tmp := type_.(string)
		result.Type = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func SettingsImagesToMap(obj oci_identity_domains.SettingsImages) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsSettingResourceCrud) mapToSettingsLoginTexts(fieldKeyFormat string) (oci_identity_domains.SettingsLoginTexts, error) {
	result := oci_identity_domains.SettingsLoginTexts{}

	if locale, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "locale")); ok {
		tmp := locale.(string)
		result.Locale = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func SettingsLoginTextsToMap(obj oci_identity_domains.SettingsLoginTexts) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Locale != nil {
		result["locale"] = string(*obj.Locale)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsSettingResourceCrud) mapToSettingsPurgeConfigs(fieldKeyFormat string) (oci_identity_domains.SettingsPurgeConfigs, error) {
	result := oci_identity_domains.SettingsPurgeConfigs{}

	if resourceName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_name")); ok {
		tmp := resourceName.(string)
		result.ResourceName = &tmp
	}

	if retentionPeriod, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "retention_period")); ok {
		tmp := retentionPeriod.(int)
		result.RetentionPeriod = &tmp
	}

	return result, nil
}

func SettingsPurgeConfigsToMap(obj oci_identity_domains.SettingsPurgeConfigs) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ResourceName != nil {
		result["resource_name"] = string(*obj.ResourceName)
	}

	if obj.RetentionPeriod != nil {
		result["retention_period"] = int(*obj.RetentionPeriod)
	}

	return result
}

func (s *IdentityDomainsSettingResourceCrud) mapToSettingsTenantCustomClaims(fieldKeyFormat string) (oci_identity_domains.SettingsTenantCustomClaims, error) {
	result := oci_identity_domains.SettingsTenantCustomClaims{}

	if allScopes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "all_scopes")); ok {
		tmp := allScopes.(bool)
		result.AllScopes = &tmp
	}

	if expression, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "expression")); ok {
		tmp := expression.(bool)
		result.Expression = &tmp
	}

	if mode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mode")); ok {
		result.Mode = oci_identity_domains.SettingsTenantCustomClaimsModeEnum(mode.(string))
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if scopes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scopes")); ok {
		interfaces := scopes.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "scopes")) {
			result.Scopes = tmp
		}
	}

	if tokenType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "token_type")); ok {
		result.TokenType = oci_identity_domains.SettingsTenantCustomClaimsTokenTypeEnum(tokenType.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func SettingsTenantCustomClaimsToMap(obj oci_identity_domains.SettingsTenantCustomClaims) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AllScopes != nil {
		result["all_scopes"] = bool(*obj.AllScopes)
	}

	if obj.Expression != nil {
		result["expression"] = bool(*obj.Expression)
	}

	result["mode"] = string(obj.Mode)

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["scopes"] = obj.Scopes

	result["token_type"] = string(obj.TokenType)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsSettingResourceCrud) mapTotags(fieldKeyFormat string) (oci_identity_domains.Tags, error) {
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
