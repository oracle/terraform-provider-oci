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

func IdentityDomainsAuthenticationFactorSettingResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityDomainsAuthenticationFactorSetting,
		Read:     readIdentityDomainsAuthenticationFactorSetting,
		Update:   updateIdentityDomainsAuthenticationFactorSetting,
		Delete:   deleteIdentityDomainsAuthenticationFactorSetting,
		Schema: map[string]*schema.Schema{
			// Required
			"authentication_factor_setting_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"bypass_code_enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"bypass_code_settings": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"help_desk_code_expiry_in_mins": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"help_desk_generation_enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"help_desk_max_usage": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"length": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"max_active": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"self_service_generation_enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"client_app_settings": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"device_protection_policy": {
							Type:     schema.TypeString,
							Required: true,
						},
						"initial_lockout_period_in_secs": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"key_pair_length": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"lockout_escalation_pattern": {
							Type:     schema.TypeString,
							Required: true,
						},
						"max_failures_before_lockout": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"max_failures_before_warning": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"max_lockout_interval_in_secs": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"min_pin_length": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"policy_update_freq_in_days": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"request_signing_algo": {
							Type:     schema.TypeString,
							Required: true,
						},
						"shared_secret_encoding": {
							Type:     schema.TypeString,
							Required: true,
						},
						"unlock_app_for_each_request_enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"unlock_app_interval_in_secs": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"unlock_on_app_foreground_enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"unlock_on_app_start_enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"compliance_policy": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"action": {
							Type:     schema.TypeString,
							Required: true,
						},
						"name": {
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
			"endpoint_restrictions": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"max_endpoint_trust_duration_in_days": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"max_enrolled_devices": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"max_incorrect_attempts": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"max_trusted_endpoints": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"trusted_endpoints_enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"idcs_endpoint": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"mfa_enrollment_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"notification_settings": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"pull_enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"push_enabled": {
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
			"security_questions_enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"sms_enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"totp_enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"totp_settings": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"email_otp_validity_duration_in_mins": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"email_passcode_length": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"hashing_algorithm": {
							Type:     schema.TypeString,
							Required: true,
						},
						"jwt_validity_duration_in_secs": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"key_refresh_interval_in_days": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"passcode_length": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"sms_otp_validity_duration_in_mins": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"sms_passcode_length": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"time_step_in_secs": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"time_step_tolerance": {
							Type:     schema.TypeInt,
							Required: true,
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
			"auto_enroll_email_factor_disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"email_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"email_settings": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"email_link_enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},

						// Optional
						"email_link_custom_url": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"fido_authenticator_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"hide_backup_factor_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"identity_store_settings": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"mobile_number_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"mobile_number_update_enabled": {
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
			"phone_call_enabled": {
				Type:     schema.TypeBool,
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
			"third_party_factor": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"duo_security": {
							Type:     schema.TypeBool,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"attestation": {
							Type:     schema.TypeString,
							Required: true,
						},
						"authenticator_selection_attachment": {
							Type:     schema.TypeString,
							Required: true,
						},
						"authenticator_selection_require_resident_key": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"authenticator_selection_resident_key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"authenticator_selection_user_verification": {
							Type:     schema.TypeString,
							Required: true,
						},
						"exclude_credentials": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"public_key_types": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"timeout": {
							Type:     schema.TypeInt,
							Required: true,
						},

						// Optional
						"domain_validation_level": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"duo_security_settings": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"api_hostname": {
										Type:     schema.TypeString,
										Required: true,
									},
									"integration_key": {
										Type:     schema.TypeString,
										Required: true,
									},
									"secret_key": {
										Type:     schema.TypeString,
										Required: true,
									},
									"user_mapping_attribute": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"attestation_key": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"user_enrollment_disabled_factors": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"yubico_otp_enabled": {
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
			"mfa_enabled_category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenancy_ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createIdentityDomainsAuthenticationFactorSetting(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsAuthenticationFactorSettingResourceCrud{}
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

func readIdentityDomainsAuthenticationFactorSetting(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsAuthenticationFactorSettingResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpointForRead(d, "authenticationFactorSettings")
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

func updateIdentityDomainsAuthenticationFactorSetting(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsAuthenticationFactorSettingResourceCrud{}
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

func deleteIdentityDomainsAuthenticationFactorSetting(d *schema.ResourceData, m interface{}) error {
	return nil
}

type IdentityDomainsAuthenticationFactorSettingResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity_domains.IdentityDomainsClient
	Res                    *oci_identity_domains.AuthenticationFactorSetting
	DisableNotFoundRetries bool
}

func (s *IdentityDomainsAuthenticationFactorSettingResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentityDomainsAuthenticationFactorSettingResourceCrud) Create() error {
	request := oci_identity_domains.PutAuthenticationFactorSettingRequest{}

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

	if authenticationFactorSettingId, ok := s.D.GetOkExists("authentication_factor_setting_id"); ok {
		tmp := authenticationFactorSettingId.(string)
		request.AuthenticationFactorSettingId = &tmp
	}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if autoEnrollEmailFactorDisabled, ok := s.D.GetOkExists("auto_enroll_email_factor_disabled"); ok {
		tmp := autoEnrollEmailFactorDisabled.(bool)
		request.AutoEnrollEmailFactorDisabled = &tmp
	}

	if bypassCodeEnabled, ok := s.D.GetOkExists("bypass_code_enabled"); ok {
		tmp := bypassCodeEnabled.(bool)
		request.BypassCodeEnabled = &tmp
	}

	if bypassCodeSettings, ok := s.D.GetOkExists("bypass_code_settings"); ok {
		if tmpList := bypassCodeSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "bypass_code_settings", 0)
			tmp, err := s.mapToAuthenticationFactorSettingsBypassCodeSettings(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.BypassCodeSettings = &tmp
		}
	}

	if clientAppSettings, ok := s.D.GetOkExists("client_app_settings"); ok {
		if tmpList := clientAppSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "client_app_settings", 0)
			tmp, err := s.mapToAuthenticationFactorSettingsClientAppSettings(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ClientAppSettings = &tmp
		}
	}

	if compliancePolicy, ok := s.D.GetOkExists("compliance_policy"); ok {
		interfaces := compliancePolicy.([]interface{})
		tmp := make([]oci_identity_domains.AuthenticationFactorSettingsCompliancePolicy, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "compliance_policy", stateDataIndex)
			converted, err := s.mapToAuthenticationFactorSettingsCompliancePolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("compliance_policy") {
			request.CompliancePolicy = tmp
		}
	}

	if emailEnabled, ok := s.D.GetOkExists("email_enabled"); ok {
		tmp := emailEnabled.(bool)
		request.EmailEnabled = &tmp
	}

	if emailSettings, ok := s.D.GetOkExists("email_settings"); ok {
		if tmpList := emailSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "email_settings", 0)
			tmp, err := s.mapToAuthenticationFactorSettingsEmailSettings(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.EmailSettings = &tmp
		}
	}

	if endpointRestrictions, ok := s.D.GetOkExists("endpoint_restrictions"); ok {
		if tmpList := endpointRestrictions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "endpoint_restrictions", 0)
			tmp, err := s.mapToAuthenticationFactorSettingsEndpointRestrictions(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.EndpointRestrictions = &tmp
		}
	}

	if fidoAuthenticatorEnabled, ok := s.D.GetOkExists("fido_authenticator_enabled"); ok {
		tmp := fidoAuthenticatorEnabled.(bool)
		request.FidoAuthenticatorEnabled = &tmp
	}

	if hideBackupFactorEnabled, ok := s.D.GetOkExists("hide_backup_factor_enabled"); ok {
		tmp := hideBackupFactorEnabled.(bool)
		request.HideBackupFactorEnabled = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if identityStoreSettings, ok := s.D.GetOkExists("identity_store_settings"); ok {
		if tmpList := identityStoreSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "identity_store_settings", 0)
			tmp, err := s.mapToAuthenticationFactorSettingsIdentityStoreSettings(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.IdentityStoreSettings = &tmp
		}
	}

	if mfaEnrollmentType, ok := s.D.GetOkExists("mfa_enrollment_type"); ok {
		tmp := mfaEnrollmentType.(string)
		request.MfaEnrollmentType = &tmp
	}

	if notificationSettings, ok := s.D.GetOkExists("notification_settings"); ok {
		if tmpList := notificationSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "notification_settings", 0)
			tmp, err := s.mapToAuthenticationFactorSettingsNotificationSettings(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NotificationSettings = &tmp
		}
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if phoneCallEnabled, ok := s.D.GetOkExists("phone_call_enabled"); ok {
		tmp := phoneCallEnabled.(bool)
		request.PhoneCallEnabled = &tmp
	}

	if pushEnabled, ok := s.D.GetOkExists("push_enabled"); ok {
		tmp := pushEnabled.(bool)
		request.PushEnabled = &tmp
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

	if securityQuestionsEnabled, ok := s.D.GetOkExists("security_questions_enabled"); ok {
		tmp := securityQuestionsEnabled.(bool)
		request.SecurityQuestionsEnabled = &tmp
	}

	if smsEnabled, ok := s.D.GetOkExists("sms_enabled"); ok {
		tmp := smsEnabled.(bool)
		request.SmsEnabled = &tmp
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

	if thirdPartyFactor, ok := s.D.GetOkExists("third_party_factor"); ok {
		if tmpList := thirdPartyFactor.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "third_party_factor", 0)
			tmp, err := s.mapToAuthenticationFactorSettingsThirdPartyFactor(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ThirdPartyFactor = &tmp
		}
	}

	if totpEnabled, ok := s.D.GetOkExists("totp_enabled"); ok {
		tmp := totpEnabled.(bool)
		request.TotpEnabled = &tmp
	}

	if totpSettings, ok := s.D.GetOkExists("totp_settings"); ok {
		if tmpList := totpSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "totp_settings", 0)
			tmp, err := s.mapToAuthenticationFactorSettingsTotpSettings(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TotpSettings = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionfidoAuthenticationFactorSettings, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionfidoAuthenticationFactorSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings", 0)
			tmp, err := s.mapToExtensionFidoAuthenticationFactorSettings(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnietfparamsscimschemasoracleidcsextensionfidoAuthenticationFactorSettings = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionthirdPartyAuthenticationFactorSettings, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionthirdPartyAuthenticationFactorSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings", 0)
			tmp, err := s.mapToExtensionThirdPartyAuthenticationFactorSettings(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnietfparamsscimschemasoracleidcsextensionthirdPartyAuthenticationFactorSettings = &tmp
		}
	}

	if userEnrollmentDisabledFactors, ok := s.D.GetOkExists("user_enrollment_disabled_factors"); ok {
		interfaces := userEnrollmentDisabledFactors.([]interface{})
		tmp := make([]oci_identity_domains.AuthenticationFactorSettingUserEnrollmentDisabledFactorsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_identity_domains.AuthenticationFactorSettingUserEnrollmentDisabledFactorsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("user_enrollment_disabled_factors") {
			request.UserEnrollmentDisabledFactors = tmp
		}
	}

	if yubicoOtpEnabled, ok := s.D.GetOkExists("yubico_otp_enabled"); ok {
		tmp := yubicoOtpEnabled.(bool)
		request.YubicoOtpEnabled = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.PutAuthenticationFactorSetting(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AuthenticationFactorSetting
	return nil
}

func (s *IdentityDomainsAuthenticationFactorSettingResourceCrud) Get() error {
	request := oci_identity_domains.GetAuthenticationFactorSettingRequest{}

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

	tmp := s.D.Id()
	request.AuthenticationFactorSettingId = &tmp

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	authenticationFactorSettingId, err := parseAuthenticationFactorSettingCompositeId(s.D.Id())
	if err == nil {
		request.AuthenticationFactorSettingId = &authenticationFactorSettingId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.GetAuthenticationFactorSetting(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AuthenticationFactorSetting
	return nil
}

func (s *IdentityDomainsAuthenticationFactorSettingResourceCrud) Update() error {
	request := oci_identity_domains.PutAuthenticationFactorSettingRequest{}

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

	tmp := s.D.Id()
	request.AuthenticationFactorSettingId = &tmp

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if autoEnrollEmailFactorDisabled, ok := s.D.GetOkExists("auto_enroll_email_factor_disabled"); ok {
		tmp := autoEnrollEmailFactorDisabled.(bool)
		request.AutoEnrollEmailFactorDisabled = &tmp
	}

	if bypassCodeEnabled, ok := s.D.GetOkExists("bypass_code_enabled"); ok {
		tmp := bypassCodeEnabled.(bool)
		request.BypassCodeEnabled = &tmp
	}

	if bypassCodeSettings, ok := s.D.GetOkExists("bypass_code_settings"); ok {
		if tmpList := bypassCodeSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "bypass_code_settings", 0)
			tmp, err := s.mapToAuthenticationFactorSettingsBypassCodeSettings(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.BypassCodeSettings = &tmp
		}
	}

	if clientAppSettings, ok := s.D.GetOkExists("client_app_settings"); ok {
		if tmpList := clientAppSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "client_app_settings", 0)
			tmp, err := s.mapToAuthenticationFactorSettingsClientAppSettings(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ClientAppSettings = &tmp
		}
	}

	if compliancePolicy, ok := s.D.GetOkExists("compliance_policy"); ok {
		interfaces := compliancePolicy.([]interface{})
		tmp := make([]oci_identity_domains.AuthenticationFactorSettingsCompliancePolicy, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "compliance_policy", stateDataIndex)
			converted, err := s.mapToAuthenticationFactorSettingsCompliancePolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("compliance_policy") {
			request.CompliancePolicy = tmp
		}
	}

	if emailEnabled, ok := s.D.GetOkExists("email_enabled"); ok {
		tmp := emailEnabled.(bool)
		request.EmailEnabled = &tmp
	}

	if emailSettings, ok := s.D.GetOkExists("email_settings"); ok {
		if tmpList := emailSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "email_settings", 0)
			tmp, err := s.mapToAuthenticationFactorSettingsEmailSettings(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.EmailSettings = &tmp
		}
	}

	if endpointRestrictions, ok := s.D.GetOkExists("endpoint_restrictions"); ok {
		if tmpList := endpointRestrictions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "endpoint_restrictions", 0)
			tmp, err := s.mapToAuthenticationFactorSettingsEndpointRestrictions(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.EndpointRestrictions = &tmp
		}
	}

	if fidoAuthenticatorEnabled, ok := s.D.GetOkExists("fido_authenticator_enabled"); ok {
		tmp := fidoAuthenticatorEnabled.(bool)
		request.FidoAuthenticatorEnabled = &tmp
	}

	if hideBackupFactorEnabled, ok := s.D.GetOkExists("hide_backup_factor_enabled"); ok {
		tmp := hideBackupFactorEnabled.(bool)
		request.HideBackupFactorEnabled = &tmp
	}

	tmp = s.D.Id()
	request.Id = &tmp

	if identityStoreSettings, ok := s.D.GetOkExists("identity_store_settings"); ok {
		if tmpList := identityStoreSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "identity_store_settings", 0)
			tmp, err := s.mapToAuthenticationFactorSettingsIdentityStoreSettings(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.IdentityStoreSettings = &tmp
		}
	}

	if mfaEnrollmentType, ok := s.D.GetOkExists("mfa_enrollment_type"); ok {
		tmp := mfaEnrollmentType.(string)
		request.MfaEnrollmentType = &tmp
	}

	if notificationSettings, ok := s.D.GetOkExists("notification_settings"); ok {
		if tmpList := notificationSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "notification_settings", 0)
			tmp, err := s.mapToAuthenticationFactorSettingsNotificationSettings(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NotificationSettings = &tmp
		}
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if phoneCallEnabled, ok := s.D.GetOkExists("phone_call_enabled"); ok {
		tmp := phoneCallEnabled.(bool)
		request.PhoneCallEnabled = &tmp
	}

	if pushEnabled, ok := s.D.GetOkExists("push_enabled"); ok {
		tmp := pushEnabled.(bool)
		request.PushEnabled = &tmp
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

	if securityQuestionsEnabled, ok := s.D.GetOkExists("security_questions_enabled"); ok {
		tmp := securityQuestionsEnabled.(bool)
		request.SecurityQuestionsEnabled = &tmp
	}

	if smsEnabled, ok := s.D.GetOkExists("sms_enabled"); ok {
		tmp := smsEnabled.(bool)
		request.SmsEnabled = &tmp
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

	if thirdPartyFactor, ok := s.D.GetOkExists("third_party_factor"); ok {
		if tmpList := thirdPartyFactor.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "third_party_factor", 0)
			tmp, err := s.mapToAuthenticationFactorSettingsThirdPartyFactor(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ThirdPartyFactor = &tmp
		}
	}

	if totpEnabled, ok := s.D.GetOkExists("totp_enabled"); ok {
		tmp := totpEnabled.(bool)
		request.TotpEnabled = &tmp
	}

	if totpSettings, ok := s.D.GetOkExists("totp_settings"); ok {
		if tmpList := totpSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "totp_settings", 0)
			tmp, err := s.mapToAuthenticationFactorSettingsTotpSettings(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TotpSettings = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionfidoAuthenticationFactorSettings, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionfidoAuthenticationFactorSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings", 0)
			tmp, err := s.mapToExtensionFidoAuthenticationFactorSettings(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnietfparamsscimschemasoracleidcsextensionfidoAuthenticationFactorSettings = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionthirdPartyAuthenticationFactorSettings, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionthirdPartyAuthenticationFactorSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings", 0)
			tmp, err := s.mapToExtensionThirdPartyAuthenticationFactorSettings(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnietfparamsscimschemasoracleidcsextensionthirdPartyAuthenticationFactorSettings = &tmp
		}
	}

	if userEnrollmentDisabledFactors, ok := s.D.GetOkExists("user_enrollment_disabled_factors"); ok {
		interfaces := userEnrollmentDisabledFactors.([]interface{})
		tmp := make([]oci_identity_domains.AuthenticationFactorSettingUserEnrollmentDisabledFactorsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_identity_domains.AuthenticationFactorSettingUserEnrollmentDisabledFactorsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("user_enrollment_disabled_factors") {
			request.UserEnrollmentDisabledFactors = tmp
		}
	}

	if yubicoOtpEnabled, ok := s.D.GetOkExists("yubico_otp_enabled"); ok {
		tmp := yubicoOtpEnabled.(bool)
		request.YubicoOtpEnabled = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.PutAuthenticationFactorSetting(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AuthenticationFactorSetting
	return nil
}

func (s *IdentityDomainsAuthenticationFactorSettingResourceCrud) SetData() error {

	authenticationFactorSettingId, err := parseAuthenticationFactorSettingCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(authenticationFactorSettingId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.AutoEnrollEmailFactorDisabled != nil {
		s.D.Set("auto_enroll_email_factor_disabled", *s.Res.AutoEnrollEmailFactorDisabled)
	}

	if s.Res.BypassCodeEnabled != nil {
		s.D.Set("bypass_code_enabled", *s.Res.BypassCodeEnabled)
	}

	if s.Res.BypassCodeSettings != nil {
		s.D.Set("bypass_code_settings", []interface{}{AuthenticationFactorSettingsBypassCodeSettingsToMap(s.Res.BypassCodeSettings)})
	} else {
		s.D.Set("bypass_code_settings", nil)
	}

	if s.Res.ClientAppSettings != nil {
		s.D.Set("client_app_settings", []interface{}{AuthenticationFactorSettingsClientAppSettingsToMap(s.Res.ClientAppSettings)})
	} else {
		s.D.Set("client_app_settings", nil)
	}

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
	}

	compliancePolicy := []interface{}{}
	for _, item := range s.Res.CompliancePolicy {
		compliancePolicy = append(compliancePolicy, AuthenticationFactorSettingsCompliancePolicyToMap(item))
	}
	s.D.Set("compliance_policy", compliancePolicy)

	if s.Res.DeleteInProgress != nil {
		s.D.Set("delete_in_progress", *s.Res.DeleteInProgress)
	}

	if s.Res.DomainOcid != nil {
		s.D.Set("domain_ocid", *s.Res.DomainOcid)
	}

	if s.Res.EmailEnabled != nil {
		s.D.Set("email_enabled", *s.Res.EmailEnabled)
	}

	if s.Res.EmailSettings != nil {
		s.D.Set("email_settings", []interface{}{AuthenticationFactorSettingsEmailSettingsToMap(s.Res.EmailSettings)})
	} else {
		s.D.Set("email_settings", nil)
	}

	if s.Res.EndpointRestrictions != nil {
		s.D.Set("endpoint_restrictions", []interface{}{AuthenticationFactorSettingsEndpointRestrictionsToMap(s.Res.EndpointRestrictions)})
	} else {
		s.D.Set("endpoint_restrictions", nil)
	}

	if s.Res.FidoAuthenticatorEnabled != nil {
		s.D.Set("fido_authenticator_enabled", *s.Res.FidoAuthenticatorEnabled)
	}

	if s.Res.HideBackupFactorEnabled != nil {
		s.D.Set("hide_backup_factor_enabled", *s.Res.HideBackupFactorEnabled)
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

	if s.Res.IdentityStoreSettings != nil {
		s.D.Set("identity_store_settings", []interface{}{AuthenticationFactorSettingsIdentityStoreSettingsToMap(s.Res.IdentityStoreSettings)})
	} else {
		s.D.Set("identity_store_settings", nil)
	}

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	if s.Res.MfaEnabledCategory != nil {
		s.D.Set("mfa_enabled_category", *s.Res.MfaEnabledCategory)
	}

	if s.Res.MfaEnrollmentType != nil {
		s.D.Set("mfa_enrollment_type", *s.Res.MfaEnrollmentType)
	}

	if s.Res.NotificationSettings != nil {
		s.D.Set("notification_settings", []interface{}{AuthenticationFactorSettingsNotificationSettingsToMap(s.Res.NotificationSettings)})
	} else {
		s.D.Set("notification_settings", nil)
	}

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
	}

	if s.Res.PhoneCallEnabled != nil {
		s.D.Set("phone_call_enabled", *s.Res.PhoneCallEnabled)
	}

	if s.Res.PushEnabled != nil {
		s.D.Set("push_enabled", *s.Res.PushEnabled)
	}

	s.D.Set("schemas", s.Res.Schemas)

	if s.Res.SecurityQuestionsEnabled != nil {
		s.D.Set("security_questions_enabled", *s.Res.SecurityQuestionsEnabled)
	}

	if s.Res.SmsEnabled != nil {
		s.D.Set("sms_enabled", *s.Res.SmsEnabled)
	}

	tags := []interface{}{}
	for _, item := range s.Res.Tags {
		tags = append(tags, tagsToMap(item))
	}
	s.D.Set("tags", tags)

	if s.Res.TenancyOcid != nil {
		s.D.Set("tenancy_ocid", *s.Res.TenancyOcid)
	}

	if s.Res.ThirdPartyFactor != nil {
		s.D.Set("third_party_factor", []interface{}{AuthenticationFactorSettingsThirdPartyFactorToMap(s.Res.ThirdPartyFactor)})
	} else {
		s.D.Set("third_party_factor", nil)
	}

	if s.Res.TotpEnabled != nil {
		s.D.Set("totp_enabled", *s.Res.TotpEnabled)
	}

	if s.Res.TotpSettings != nil {
		s.D.Set("totp_settings", []interface{}{AuthenticationFactorSettingsTotpSettingsToMap(s.Res.TotpSettings)})
	} else {
		s.D.Set("totp_settings", nil)
	}

	if s.Res.UrnietfparamsscimschemasoracleidcsextensionfidoAuthenticationFactorSettings != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings", []interface{}{ExtensionFidoAuthenticationFactorSettingsToMap(s.Res.UrnietfparamsscimschemasoracleidcsextensionfidoAuthenticationFactorSettings)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings", nil)
	}

	if s.Res.UrnietfparamsscimschemasoracleidcsextensionthirdPartyAuthenticationFactorSettings != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings", []interface{}{ExtensionThirdPartyAuthenticationFactorSettingsToMap(s.Res.UrnietfparamsscimschemasoracleidcsextensionthirdPartyAuthenticationFactorSettings)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings", nil)
	}

	s.D.Set("user_enrollment_disabled_factors", s.Res.UserEnrollmentDisabledFactors)

	if s.Res.YubicoOtpEnabled != nil {
		s.D.Set("yubico_otp_enabled", *s.Res.YubicoOtpEnabled)
	}

	return nil
}

//func GetAuthenticationFactorSettingCompositeId(authenticationFactorSettingId string) string {
//	authenticationFactorSettingId = url.PathEscape(authenticationFactorSettingId)
//	id = url.PathEscape(id)
//	idcsEndpoint = url.PathEscape(idcsEndpoint)
//	compositeId := "idcsEndpoint/" + idcsEndpoint + "/authenticationFactorSettings/" + authenticationFactorSettingId
//	return compositeId
//}

func parseAuthenticationFactorSettingCompositeId(compositeId string) (authenticationFactorSettingId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("idcsEndpoint/.*/authenticationFactorSettings/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	//idcsEndpoint, _ = url.PathUnescape(parts[1])
	authenticationFactorSettingId, _ = url.PathUnescape(parts[3])

	return
}

func AuthenticationFactorSettingToMap(obj oci_identity_domains.AuthenticationFactorSetting) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AutoEnrollEmailFactorDisabled != nil {
		result["auto_enroll_email_factor_disabled"] = bool(*obj.AutoEnrollEmailFactorDisabled)
	}

	if obj.BypassCodeEnabled != nil {
		result["bypass_code_enabled"] = bool(*obj.BypassCodeEnabled)
	}

	if obj.BypassCodeSettings != nil {
		result["bypass_code_settings"] = []interface{}{AuthenticationFactorSettingsBypassCodeSettingsToMap(obj.BypassCodeSettings)}
	}

	if obj.ClientAppSettings != nil {
		result["client_app_settings"] = []interface{}{AuthenticationFactorSettingsClientAppSettingsToMap(obj.ClientAppSettings)}
	}

	if obj.CompartmentOcid != nil {
		result["compartment_ocid"] = string(*obj.CompartmentOcid)
	}

	compliancePolicy := []interface{}{}
	for _, item := range obj.CompliancePolicy {
		compliancePolicy = append(compliancePolicy, AuthenticationFactorSettingsCompliancePolicyToMap(item))
	}
	result["compliance_policy"] = compliancePolicy

	if obj.DeleteInProgress != nil {
		result["delete_in_progress"] = bool(*obj.DeleteInProgress)
	}

	if obj.DomainOcid != nil {
		result["domain_ocid"] = string(*obj.DomainOcid)
	}

	if obj.EmailEnabled != nil {
		result["email_enabled"] = bool(*obj.EmailEnabled)
	}

	if obj.EmailSettings != nil {
		result["email_settings"] = []interface{}{AuthenticationFactorSettingsEmailSettingsToMap(obj.EmailSettings)}
	}

	if obj.EndpointRestrictions != nil {
		result["endpoint_restrictions"] = []interface{}{AuthenticationFactorSettingsEndpointRestrictionsToMap(obj.EndpointRestrictions)}
	}

	if obj.FidoAuthenticatorEnabled != nil {
		result["fido_authenticator_enabled"] = bool(*obj.FidoAuthenticatorEnabled)
	}

	if obj.HideBackupFactorEnabled != nil {
		result["hide_backup_factor_enabled"] = bool(*obj.HideBackupFactorEnabled)
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

	if obj.IdentityStoreSettings != nil {
		result["identity_store_settings"] = []interface{}{AuthenticationFactorSettingsIdentityStoreSettingsToMap(obj.IdentityStoreSettings)}
	}

	if obj.Meta != nil {
		result["meta"] = []interface{}{metaToMap(obj.Meta)}
	}

	if obj.MfaEnabledCategory != nil {
		result["mfa_enabled_category"] = string(*obj.MfaEnabledCategory)
	}

	if obj.MfaEnrollmentType != nil {
		result["mfa_enrollment_type"] = string(*obj.MfaEnrollmentType)
	}

	if obj.NotificationSettings != nil {
		result["notification_settings"] = []interface{}{AuthenticationFactorSettingsNotificationSettingsToMap(obj.NotificationSettings)}
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	if obj.PhoneCallEnabled != nil {
		result["phone_call_enabled"] = bool(*obj.PhoneCallEnabled)
	}

	if obj.PushEnabled != nil {
		result["push_enabled"] = bool(*obj.PushEnabled)
	}

	result["schemas"] = obj.Schemas

	if obj.SecurityQuestionsEnabled != nil {
		result["security_questions_enabled"] = bool(*obj.SecurityQuestionsEnabled)
	}

	if obj.SmsEnabled != nil {
		result["sms_enabled"] = bool(*obj.SmsEnabled)
	}

	tags := []interface{}{}
	for _, item := range obj.Tags {
		tags = append(tags, tagsToMap(item))
	}
	result["tags"] = tags

	if obj.TenancyOcid != nil {
		result["tenancy_ocid"] = string(*obj.TenancyOcid)
	}

	if obj.ThirdPartyFactor != nil {
		result["third_party_factor"] = []interface{}{AuthenticationFactorSettingsThirdPartyFactorToMap(obj.ThirdPartyFactor)}
	}

	if obj.TotpEnabled != nil {
		result["totp_enabled"] = bool(*obj.TotpEnabled)
	}

	if obj.TotpSettings != nil {
		result["totp_settings"] = []interface{}{AuthenticationFactorSettingsTotpSettingsToMap(obj.TotpSettings)}
	}

	if obj.UrnietfparamsscimschemasoracleidcsextensionfidoAuthenticationFactorSettings != nil {
		result["urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings"] = []interface{}{ExtensionFidoAuthenticationFactorSettingsToMap(obj.UrnietfparamsscimschemasoracleidcsextensionfidoAuthenticationFactorSettings)}
	}

	if obj.UrnietfparamsscimschemasoracleidcsextensionthirdPartyAuthenticationFactorSettings != nil {
		result["urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings"] = []interface{}{ExtensionThirdPartyAuthenticationFactorSettingsToMap(obj.UrnietfparamsscimschemasoracleidcsextensionthirdPartyAuthenticationFactorSettings)}
	}

	result["user_enrollment_disabled_factors"] = obj.UserEnrollmentDisabledFactors

	if obj.YubicoOtpEnabled != nil {
		result["yubico_otp_enabled"] = bool(*obj.YubicoOtpEnabled)
	}

	return result
}

func (s *IdentityDomainsAuthenticationFactorSettingResourceCrud) mapToAuthenticationFactorSettingsBypassCodeSettings(fieldKeyFormat string) (oci_identity_domains.AuthenticationFactorSettingsBypassCodeSettings, error) {
	result := oci_identity_domains.AuthenticationFactorSettingsBypassCodeSettings{}

	if helpDeskCodeExpiryInMins, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "help_desk_code_expiry_in_mins")); ok {
		tmp := helpDeskCodeExpiryInMins.(int)
		result.HelpDeskCodeExpiryInMins = &tmp
	}

	if helpDeskGenerationEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "help_desk_generation_enabled")); ok {
		tmp := helpDeskGenerationEnabled.(bool)
		result.HelpDeskGenerationEnabled = &tmp
	}

	if helpDeskMaxUsage, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "help_desk_max_usage")); ok {
		tmp := helpDeskMaxUsage.(int)
		result.HelpDeskMaxUsage = &tmp
	}

	if length, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "length")); ok {
		tmp := length.(int)
		result.Length = &tmp
	}

	if maxActive, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_active")); ok {
		tmp := maxActive.(int)
		result.MaxActive = &tmp
	}

	if selfServiceGenerationEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "self_service_generation_enabled")); ok {
		tmp := selfServiceGenerationEnabled.(bool)
		result.SelfServiceGenerationEnabled = &tmp
	}

	return result, nil
}

func AuthenticationFactorSettingsBypassCodeSettingsToMap(obj *oci_identity_domains.AuthenticationFactorSettingsBypassCodeSettings) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.HelpDeskCodeExpiryInMins != nil {
		result["help_desk_code_expiry_in_mins"] = int(*obj.HelpDeskCodeExpiryInMins)
	}

	if obj.HelpDeskGenerationEnabled != nil {
		result["help_desk_generation_enabled"] = bool(*obj.HelpDeskGenerationEnabled)
	}

	if obj.HelpDeskMaxUsage != nil {
		result["help_desk_max_usage"] = int(*obj.HelpDeskMaxUsage)
	}

	if obj.Length != nil {
		result["length"] = int(*obj.Length)
	}

	if obj.MaxActive != nil {
		result["max_active"] = int(*obj.MaxActive)
	}

	if obj.SelfServiceGenerationEnabled != nil {
		result["self_service_generation_enabled"] = bool(*obj.SelfServiceGenerationEnabled)
	}

	return result
}

func (s *IdentityDomainsAuthenticationFactorSettingResourceCrud) mapToAuthenticationFactorSettingsClientAppSettings(fieldKeyFormat string) (oci_identity_domains.AuthenticationFactorSettingsClientAppSettings, error) {
	result := oci_identity_domains.AuthenticationFactorSettingsClientAppSettings{}

	if deviceProtectionPolicy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "device_protection_policy")); ok {
		tmp := deviceProtectionPolicy.(string)
		result.DeviceProtectionPolicy = &tmp
	}

	if initialLockoutPeriodInSecs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "initial_lockout_period_in_secs")); ok {
		tmp := initialLockoutPeriodInSecs.(int)
		result.InitialLockoutPeriodInSecs = &tmp
	}

	if keyPairLength, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_pair_length")); ok {
		tmp := keyPairLength.(int)
		result.KeyPairLength = &tmp
	}

	if lockoutEscalationPattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "lockout_escalation_pattern")); ok {
		tmp := lockoutEscalationPattern.(string)
		result.LockoutEscalationPattern = &tmp
	}

	if maxFailuresBeforeLockout, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_failures_before_lockout")); ok {
		tmp := maxFailuresBeforeLockout.(int)
		result.MaxFailuresBeforeLockout = &tmp
	}

	if maxFailuresBeforeWarning, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_failures_before_warning")); ok {
		tmp := maxFailuresBeforeWarning.(int)
		result.MaxFailuresBeforeWarning = &tmp
	}

	if maxLockoutIntervalInSecs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_lockout_interval_in_secs")); ok {
		tmp := maxLockoutIntervalInSecs.(int)
		result.MaxLockoutIntervalInSecs = &tmp
	}

	if minPinLength, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "min_pin_length")); ok {
		tmp := minPinLength.(int)
		result.MinPinLength = &tmp
	}

	if policyUpdateFreqInDays, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "policy_update_freq_in_days")); ok {
		tmp := policyUpdateFreqInDays.(int)
		result.PolicyUpdateFreqInDays = &tmp
	}

	if requestSigningAlgo, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "request_signing_algo")); ok {
		result.RequestSigningAlgo = oci_identity_domains.AuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoEnum(requestSigningAlgo.(string))
	}

	if sharedSecretEncoding, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shared_secret_encoding")); ok {
		result.SharedSecretEncoding = oci_identity_domains.AuthenticationFactorSettingsClientAppSettingsSharedSecretEncodingEnum(sharedSecretEncoding.(string))
	}

	if unlockAppForEachRequestEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "unlock_app_for_each_request_enabled")); ok {
		tmp := unlockAppForEachRequestEnabled.(bool)
		result.UnlockAppForEachRequestEnabled = &tmp
	}

	if unlockAppIntervalInSecs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "unlock_app_interval_in_secs")); ok {
		tmp := unlockAppIntervalInSecs.(int)
		result.UnlockAppIntervalInSecs = &tmp
	}

	if unlockOnAppForegroundEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "unlock_on_app_foreground_enabled")); ok {
		tmp := unlockOnAppForegroundEnabled.(bool)
		result.UnlockOnAppForegroundEnabled = &tmp
	}

	if unlockOnAppStartEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "unlock_on_app_start_enabled")); ok {
		tmp := unlockOnAppStartEnabled.(bool)
		result.UnlockOnAppStartEnabled = &tmp
	}

	return result, nil
}

func AuthenticationFactorSettingsClientAppSettingsToMap(obj *oci_identity_domains.AuthenticationFactorSettingsClientAppSettings) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DeviceProtectionPolicy != nil {
		result["device_protection_policy"] = string(*obj.DeviceProtectionPolicy)
	}

	if obj.InitialLockoutPeriodInSecs != nil {
		result["initial_lockout_period_in_secs"] = int(*obj.InitialLockoutPeriodInSecs)
	}

	if obj.KeyPairLength != nil {
		result["key_pair_length"] = int(*obj.KeyPairLength)
	}

	if obj.LockoutEscalationPattern != nil {
		result["lockout_escalation_pattern"] = string(*obj.LockoutEscalationPattern)
	}

	if obj.MaxFailuresBeforeLockout != nil {
		result["max_failures_before_lockout"] = int(*obj.MaxFailuresBeforeLockout)
	}

	if obj.MaxFailuresBeforeWarning != nil {
		result["max_failures_before_warning"] = int(*obj.MaxFailuresBeforeWarning)
	}

	if obj.MaxLockoutIntervalInSecs != nil {
		result["max_lockout_interval_in_secs"] = int(*obj.MaxLockoutIntervalInSecs)
	}

	if obj.MinPinLength != nil {
		result["min_pin_length"] = int(*obj.MinPinLength)
	}

	if obj.PolicyUpdateFreqInDays != nil {
		result["policy_update_freq_in_days"] = int(*obj.PolicyUpdateFreqInDays)
	}

	result["request_signing_algo"] = string(obj.RequestSigningAlgo)

	result["shared_secret_encoding"] = string(obj.SharedSecretEncoding)

	if obj.UnlockAppForEachRequestEnabled != nil {
		result["unlock_app_for_each_request_enabled"] = bool(*obj.UnlockAppForEachRequestEnabled)
	}

	if obj.UnlockAppIntervalInSecs != nil {
		result["unlock_app_interval_in_secs"] = int(*obj.UnlockAppIntervalInSecs)
	}

	if obj.UnlockOnAppForegroundEnabled != nil {
		result["unlock_on_app_foreground_enabled"] = bool(*obj.UnlockOnAppForegroundEnabled)
	}

	if obj.UnlockOnAppStartEnabled != nil {
		result["unlock_on_app_start_enabled"] = bool(*obj.UnlockOnAppStartEnabled)
	}

	return result
}

func (s *IdentityDomainsAuthenticationFactorSettingResourceCrud) mapToAuthenticationFactorSettingsCompliancePolicy(fieldKeyFormat string) (oci_identity_domains.AuthenticationFactorSettingsCompliancePolicy, error) {
	result := oci_identity_domains.AuthenticationFactorSettingsCompliancePolicy{}

	if action, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action")); ok {
		result.Action = oci_identity_domains.AuthenticationFactorSettingsCompliancePolicyActionEnum(action.(string))
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func AuthenticationFactorSettingsCompliancePolicyToMap(obj oci_identity_domains.AuthenticationFactorSettingsCompliancePolicy) map[string]interface{} {
	result := map[string]interface{}{}

	result["action"] = string(obj.Action)

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsAuthenticationFactorSettingResourceCrud) mapToAuthenticationFactorSettingsDuoSecuritySettings(fieldKeyFormat string) (oci_identity_domains.AuthenticationFactorSettingsDuoSecuritySettings, error) {
	result := oci_identity_domains.AuthenticationFactorSettingsDuoSecuritySettings{}

	if apiHostname, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "api_hostname")); ok {
		tmp := apiHostname.(string)
		result.ApiHostname = &tmp
	}

	if attestationKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "attestation_key")); ok {
		tmp := attestationKey.(string)
		result.AttestationKey = &tmp
	}

	if integrationKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "integration_key")); ok {
		tmp := integrationKey.(string)
		result.IntegrationKey = &tmp
	}

	if secretKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_key")); ok {
		tmp := secretKey.(string)
		result.SecretKey = &tmp
	}

	if userMappingAttribute, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_mapping_attribute")); ok {
		result.UserMappingAttribute = oci_identity_domains.AuthenticationFactorSettingsDuoSecuritySettingsUserMappingAttributeEnum(userMappingAttribute.(string))
	}

	return result, nil
}

func AuthenticationFactorSettingsDuoSecuritySettingsToMap(obj *oci_identity_domains.AuthenticationFactorSettingsDuoSecuritySettings) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApiHostname != nil {
		result["api_hostname"] = string(*obj.ApiHostname)
	}

	if obj.AttestationKey != nil {
		result["attestation_key"] = string(*obj.AttestationKey)
	}

	if obj.IntegrationKey != nil {
		result["integration_key"] = string(*obj.IntegrationKey)
	}

	if obj.SecretKey != nil {
		result["secret_key"] = string(*obj.SecretKey)
	}

	result["user_mapping_attribute"] = string(obj.UserMappingAttribute)

	return result
}

func (s *IdentityDomainsAuthenticationFactorSettingResourceCrud) mapToAuthenticationFactorSettingsEmailSettings(fieldKeyFormat string) (oci_identity_domains.AuthenticationFactorSettingsEmailSettings, error) {
	result := oci_identity_domains.AuthenticationFactorSettingsEmailSettings{}

	if emailLinkCustomUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "email_link_custom_url")); ok {
		tmp := emailLinkCustomUrl.(string)
		result.EmailLinkCustomUrl = &tmp
	}

	if emailLinkEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "email_link_enabled")); ok {
		tmp := emailLinkEnabled.(bool)
		result.EmailLinkEnabled = &tmp
	}

	return result, nil
}

func AuthenticationFactorSettingsEmailSettingsToMap(obj *oci_identity_domains.AuthenticationFactorSettingsEmailSettings) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EmailLinkCustomUrl != nil {
		result["email_link_custom_url"] = string(*obj.EmailLinkCustomUrl)
	}

	if obj.EmailLinkEnabled != nil {
		result["email_link_enabled"] = bool(*obj.EmailLinkEnabled)
	}

	return result
}

func (s *IdentityDomainsAuthenticationFactorSettingResourceCrud) mapToAuthenticationFactorSettingsEndpointRestrictions(fieldKeyFormat string) (oci_identity_domains.AuthenticationFactorSettingsEndpointRestrictions, error) {
	result := oci_identity_domains.AuthenticationFactorSettingsEndpointRestrictions{}

	if maxEndpointTrustDurationInDays, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_endpoint_trust_duration_in_days")); ok {
		tmp := maxEndpointTrustDurationInDays.(int)
		result.MaxEndpointTrustDurationInDays = &tmp
	}

	if maxEnrolledDevices, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_enrolled_devices")); ok {
		tmp := maxEnrolledDevices.(int)
		result.MaxEnrolledDevices = &tmp
	}

	if maxIncorrectAttempts, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_incorrect_attempts")); ok {
		tmp := maxIncorrectAttempts.(int)
		result.MaxIncorrectAttempts = &tmp
	}

	if maxTrustedEndpoints, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_trusted_endpoints")); ok {
		tmp := maxTrustedEndpoints.(int)
		result.MaxTrustedEndpoints = &tmp
	}

	if trustedEndpointsEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "trusted_endpoints_enabled")); ok {
		tmp := trustedEndpointsEnabled.(bool)
		result.TrustedEndpointsEnabled = &tmp
	}

	return result, nil
}

func AuthenticationFactorSettingsEndpointRestrictionsToMap(obj *oci_identity_domains.AuthenticationFactorSettingsEndpointRestrictions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MaxEndpointTrustDurationInDays != nil {
		result["max_endpoint_trust_duration_in_days"] = int(*obj.MaxEndpointTrustDurationInDays)
	}

	if obj.MaxEnrolledDevices != nil {
		result["max_enrolled_devices"] = int(*obj.MaxEnrolledDevices)
	}

	if obj.MaxIncorrectAttempts != nil {
		result["max_incorrect_attempts"] = int(*obj.MaxIncorrectAttempts)
	}

	if obj.MaxTrustedEndpoints != nil {
		result["max_trusted_endpoints"] = int(*obj.MaxTrustedEndpoints)
	}

	if obj.TrustedEndpointsEnabled != nil {
		result["trusted_endpoints_enabled"] = bool(*obj.TrustedEndpointsEnabled)
	}

	return result
}

func (s *IdentityDomainsAuthenticationFactorSettingResourceCrud) mapToAuthenticationFactorSettingsIdentityStoreSettings(fieldKeyFormat string) (oci_identity_domains.AuthenticationFactorSettingsIdentityStoreSettings, error) {
	result := oci_identity_domains.AuthenticationFactorSettingsIdentityStoreSettings{}

	if mobileNumberEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mobile_number_enabled")); ok {
		tmp := mobileNumberEnabled.(bool)
		result.MobileNumberEnabled = &tmp
	}

	if mobileNumberUpdateEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mobile_number_update_enabled")); ok {
		tmp := mobileNumberUpdateEnabled.(bool)
		result.MobileNumberUpdateEnabled = &tmp
	}

	return result, nil
}

func AuthenticationFactorSettingsIdentityStoreSettingsToMap(obj *oci_identity_domains.AuthenticationFactorSettingsIdentityStoreSettings) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MobileNumberEnabled != nil {
		result["mobile_number_enabled"] = bool(*obj.MobileNumberEnabled)
	}

	if obj.MobileNumberUpdateEnabled != nil {
		result["mobile_number_update_enabled"] = bool(*obj.MobileNumberUpdateEnabled)
	}

	return result
}

func (s *IdentityDomainsAuthenticationFactorSettingResourceCrud) mapToAuthenticationFactorSettingsNotificationSettings(fieldKeyFormat string) (oci_identity_domains.AuthenticationFactorSettingsNotificationSettings, error) {
	result := oci_identity_domains.AuthenticationFactorSettingsNotificationSettings{}

	if pullEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pull_enabled")); ok {
		tmp := pullEnabled.(bool)
		result.PullEnabled = &tmp
	}

	return result, nil
}

func AuthenticationFactorSettingsNotificationSettingsToMap(obj *oci_identity_domains.AuthenticationFactorSettingsNotificationSettings) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.PullEnabled != nil {
		result["pull_enabled"] = bool(*obj.PullEnabled)
	}

	return result
}

func (s *IdentityDomainsAuthenticationFactorSettingResourceCrud) mapToAuthenticationFactorSettingsThirdPartyFactor(fieldKeyFormat string) (oci_identity_domains.AuthenticationFactorSettingsThirdPartyFactor, error) {
	result := oci_identity_domains.AuthenticationFactorSettingsThirdPartyFactor{}

	if duoSecurity, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "duo_security")); ok {
		tmp := duoSecurity.(bool)
		result.DuoSecurity = &tmp
	}

	return result, nil
}

func AuthenticationFactorSettingsThirdPartyFactorToMap(obj *oci_identity_domains.AuthenticationFactorSettingsThirdPartyFactor) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DuoSecurity != nil {
		result["duo_security"] = bool(*obj.DuoSecurity)
	}

	return result
}

func (s *IdentityDomainsAuthenticationFactorSettingResourceCrud) mapToAuthenticationFactorSettingsTotpSettings(fieldKeyFormat string) (oci_identity_domains.AuthenticationFactorSettingsTotpSettings, error) {
	result := oci_identity_domains.AuthenticationFactorSettingsTotpSettings{}

	if emailOtpValidityDurationInMins, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "email_otp_validity_duration_in_mins")); ok {
		tmp := emailOtpValidityDurationInMins.(int)
		result.EmailOtpValidityDurationInMins = &tmp
	}

	if emailPasscodeLength, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "email_passcode_length")); ok {
		tmp := emailPasscodeLength.(int)
		result.EmailPasscodeLength = &tmp
	}

	if hashingAlgorithm, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hashing_algorithm")); ok {
		result.HashingAlgorithm = oci_identity_domains.AuthenticationFactorSettingsTotpSettingsHashingAlgorithmEnum(hashingAlgorithm.(string))
	}

	if jwtValidityDurationInSecs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "jwt_validity_duration_in_secs")); ok {
		tmp := jwtValidityDurationInSecs.(int)
		result.JwtValidityDurationInSecs = &tmp
	}

	if keyRefreshIntervalInDays, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_refresh_interval_in_days")); ok {
		tmp := keyRefreshIntervalInDays.(int)
		result.KeyRefreshIntervalInDays = &tmp
	}

	if passcodeLength, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "passcode_length")); ok {
		tmp := passcodeLength.(int)
		result.PasscodeLength = &tmp
	}

	if smsOtpValidityDurationInMins, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sms_otp_validity_duration_in_mins")); ok {
		tmp := smsOtpValidityDurationInMins.(int)
		result.SmsOtpValidityDurationInMins = &tmp
	}

	if smsPasscodeLength, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sms_passcode_length")); ok {
		tmp := smsPasscodeLength.(int)
		result.SmsPasscodeLength = &tmp
	}

	if timeStepInSecs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_step_in_secs")); ok {
		tmp := timeStepInSecs.(int)
		result.TimeStepInSecs = &tmp
	}

	if timeStepTolerance, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_step_tolerance")); ok {
		tmp := timeStepTolerance.(int)
		result.TimeStepTolerance = &tmp
	}

	return result, nil
}

func AuthenticationFactorSettingsTotpSettingsToMap(obj *oci_identity_domains.AuthenticationFactorSettingsTotpSettings) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EmailOtpValidityDurationInMins != nil {
		result["email_otp_validity_duration_in_mins"] = int(*obj.EmailOtpValidityDurationInMins)
	}

	if obj.EmailPasscodeLength != nil {
		result["email_passcode_length"] = int(*obj.EmailPasscodeLength)
	}

	result["hashing_algorithm"] = string(obj.HashingAlgorithm)

	if obj.JwtValidityDurationInSecs != nil {
		result["jwt_validity_duration_in_secs"] = int(*obj.JwtValidityDurationInSecs)
	}

	if obj.KeyRefreshIntervalInDays != nil {
		result["key_refresh_interval_in_days"] = int(*obj.KeyRefreshIntervalInDays)
	}

	if obj.PasscodeLength != nil {
		result["passcode_length"] = int(*obj.PasscodeLength)
	}

	if obj.SmsOtpValidityDurationInMins != nil {
		result["sms_otp_validity_duration_in_mins"] = int(*obj.SmsOtpValidityDurationInMins)
	}

	if obj.SmsPasscodeLength != nil {
		result["sms_passcode_length"] = int(*obj.SmsPasscodeLength)
	}

	if obj.TimeStepInSecs != nil {
		result["time_step_in_secs"] = int(*obj.TimeStepInSecs)
	}

	if obj.TimeStepTolerance != nil {
		result["time_step_tolerance"] = int(*obj.TimeStepTolerance)
	}

	return result
}

func (s *IdentityDomainsAuthenticationFactorSettingResourceCrud) mapToExtensionFidoAuthenticationFactorSettings(fieldKeyFormat string) (oci_identity_domains.ExtensionFidoAuthenticationFactorSettings, error) {
	result := oci_identity_domains.ExtensionFidoAuthenticationFactorSettings{}

	if attestation, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "attestation")); ok {
		result.Attestation = oci_identity_domains.ExtensionFidoAuthenticationFactorSettingsAttestationEnum(attestation.(string))
	}

	if authenticatorSelectionAttachment, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "authenticator_selection_attachment")); ok {
		result.AuthenticatorSelectionAttachment = oci_identity_domains.ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentEnum(authenticatorSelectionAttachment.(string))
	}

	if authenticatorSelectionRequireResidentKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "authenticator_selection_require_resident_key")); ok {
		tmp := authenticatorSelectionRequireResidentKey.(bool)
		result.AuthenticatorSelectionRequireResidentKey = &tmp
	}

	if authenticatorSelectionResidentKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "authenticator_selection_resident_key")); ok {
		result.AuthenticatorSelectionResidentKey = oci_identity_domains.ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyEnum(authenticatorSelectionResidentKey.(string))
	}

	if authenticatorSelectionUserVerification, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "authenticator_selection_user_verification")); ok {
		result.AuthenticatorSelectionUserVerification = oci_identity_domains.ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationEnum(authenticatorSelectionUserVerification.(string))
	}

	if domainValidationLevel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "domain_validation_level")); ok {
		tmp := domainValidationLevel.(int)
		result.DomainValidationLevel = &tmp
	}

	if excludeCredentials, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exclude_credentials")); ok {
		tmp := excludeCredentials.(bool)
		result.ExcludeCredentials = &tmp
	}

	if publicKeyTypes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "public_key_types")); ok {
		interfaces := publicKeyTypes.([]interface{})
		tmp := make([]oci_identity_domains.ExtensionFidoAuthenticationFactorSettingsPublicKeyTypesEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_identity_domains.ExtensionFidoAuthenticationFactorSettingsPublicKeyTypesEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "public_key_types")) {
			result.PublicKeyTypes = tmp
		}
	}

	if timeout, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout")); ok {
		tmp := timeout.(int)
		result.Timeout = &tmp
	}

	return result, nil
}

func ExtensionFidoAuthenticationFactorSettingsToMap(obj *oci_identity_domains.ExtensionFidoAuthenticationFactorSettings) map[string]interface{} {
	result := map[string]interface{}{}

	result["attestation"] = string(obj.Attestation)

	result["authenticator_selection_attachment"] = string(obj.AuthenticatorSelectionAttachment)

	if obj.AuthenticatorSelectionRequireResidentKey != nil {
		result["authenticator_selection_require_resident_key"] = bool(*obj.AuthenticatorSelectionRequireResidentKey)
	}

	result["authenticator_selection_resident_key"] = string(obj.AuthenticatorSelectionResidentKey)

	result["authenticator_selection_user_verification"] = string(obj.AuthenticatorSelectionUserVerification)

	if obj.DomainValidationLevel != nil {
		result["domain_validation_level"] = int(*obj.DomainValidationLevel)
	}

	if obj.ExcludeCredentials != nil {
		result["exclude_credentials"] = bool(*obj.ExcludeCredentials)
	}

	result["public_key_types"] = obj.PublicKeyTypes

	if obj.Timeout != nil {
		result["timeout"] = int(*obj.Timeout)
	}

	return result
}

func (s *IdentityDomainsAuthenticationFactorSettingResourceCrud) mapToExtensionThirdPartyAuthenticationFactorSettings(fieldKeyFormat string) (oci_identity_domains.ExtensionThirdPartyAuthenticationFactorSettings, error) {
	result := oci_identity_domains.ExtensionThirdPartyAuthenticationFactorSettings{}

	if duoSecuritySettings, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "duo_security_settings")); ok {
		if tmpList := duoSecuritySettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "duo_security_settings"), 0)
			tmp, err := s.mapToAuthenticationFactorSettingsDuoSecuritySettings(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert duo_security_settings, encountered error: %v", err)
			}
			result.DuoSecuritySettings = &tmp
		}
	}

	return result, nil
}

func ExtensionThirdPartyAuthenticationFactorSettingsToMap(obj *oci_identity_domains.ExtensionThirdPartyAuthenticationFactorSettings) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DuoSecuritySettings != nil {
		result["duo_security_settings"] = []interface{}{AuthenticationFactorSettingsDuoSecuritySettingsToMap(obj.DuoSecuritySettings)}
	}

	return result
}

func (s *IdentityDomainsAuthenticationFactorSettingResourceCrud) mapTotags(fieldKeyFormat string) (oci_identity_domains.Tags, error) {
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
