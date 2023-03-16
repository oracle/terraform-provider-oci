---
subcategory: "Identity Domains"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_domains_authentication_factor_setting"
sidebar_current: "docs-oci-resource-identity_domains-authentication_factor_setting"
description: |-
  Provides the Authentication Factor Setting resource in Oracle Cloud Infrastructure Identity Domains service
---

# oci_identity_domains_authentication_factor_setting
This resource provides the Authentication Factor Setting resource in Oracle Cloud Infrastructure Identity Domains service.

Replace Authentication Factor Settings

## Example Usage

```hcl
resource "oci_identity_domains_authentication_factor_setting" "test_authentication_factor_setting" {
	#Required
	authentication_factor_setting_id = oci_identity_domains_authentication_factor_setting.test_authentication_factor_setting.id
	bypass_code_enabled = var.authentication_factor_setting_bypass_code_enabled
	bypass_code_settings {
		#Required
		help_desk_code_expiry_in_mins = var.authentication_factor_setting_bypass_code_settings_help_desk_code_expiry_in_mins
		help_desk_generation_enabled = var.authentication_factor_setting_bypass_code_settings_help_desk_generation_enabled
		help_desk_max_usage = var.authentication_factor_setting_bypass_code_settings_help_desk_max_usage
		length = var.authentication_factor_setting_bypass_code_settings_length
		max_active = "6"
		self_service_generation_enabled = var.authentication_factor_setting_bypass_code_settings_self_service_generation_enabled
	}
	client_app_settings {
		#Required
		device_protection_policy = "NONE"
		initial_lockout_period_in_secs = "30"
		key_pair_length = "2048"
		lockout_escalation_pattern = "Constant"
		max_failures_before_lockout = "10"
		max_failures_before_warning = "5"
		max_lockout_interval_in_secs = "86400"
		min_pin_length = "6"
		policy_update_freq_in_days = var.authentication_factor_setting_client_app_settings_policy_update_freq_in_days
		request_signing_algo = var.authentication_factor_setting_client_app_settings_request_signing_algo
		shared_secret_encoding = var.authentication_factor_setting_client_app_settings_shared_secret_encoding
		unlock_app_for_each_request_enabled = "false"
		unlock_app_interval_in_secs = "300"
		unlock_on_app_foreground_enabled = "false"
		unlock_on_app_start_enabled = "false"
	}
	compliance_policy {
		#Required
		action = "Allow"
		name = "lockScreenRequired"
		value = "false"
	}
	endpoint_restrictions {
		#Required
		max_endpoint_trust_duration_in_days = "180"
		max_enrolled_devices = var.authentication_factor_setting_endpoint_restrictions_max_enrolled_devices
		max_incorrect_attempts = "20"
		max_trusted_endpoints = "20"
		trusted_endpoints_enabled = var.authentication_factor_setting_endpoint_restrictions_trusted_endpoints_enabled
	}
	idcs_endpoint = data.oci_identity_domain.test_domain.url
	mfa_enrollment_type = var.authentication_factor_setting_mfa_enrollment_type
	notification_settings {
		#Required
		pull_enabled = var.authentication_factor_setting_notification_settings_pull_enabled
	}
	push_enabled = var.authentication_factor_setting_push_enabled
	schemas = ["urn:ietf:params:scim:schemas:oracle:idcs:AuthenticationFactorSettings"]
	security_questions_enabled = var.authentication_factor_setting_security_questions_enabled
	sms_enabled = var.authentication_factor_setting_sms_enabled
	totp_enabled = var.authentication_factor_setting_totp_enabled
	totp_settings {
		#Required
		email_otp_validity_duration_in_mins = var.authentication_factor_setting_totp_settings_email_otp_validity_duration_in_mins
		email_passcode_length = "6"
		hashing_algorithm = var.authentication_factor_setting_totp_settings_hashing_algorithm
		jwt_validity_duration_in_secs = "300"
		key_refresh_interval_in_days = "60"
		passcode_length = "6"
		sms_otp_validity_duration_in_mins = "6"
		sms_passcode_length = "6"
		time_step_in_secs = "30"
		time_step_tolerance = "3"
	}

	#Optional
	attribute_sets = []
	attributes = ""
	authorization = var.authentication_factor_setting_authorization
	auto_enroll_email_factor_disabled = var.authentication_factor_setting_auto_enroll_email_factor_disabled
	email_enabled = var.authentication_factor_setting_email_enabled
	email_settings {
		#Required
		email_link_enabled = var.authentication_factor_setting_email_settings_email_link_enabled

		#Optional
		email_link_custom_url = var.authentication_factor_setting_email_settings_email_link_custom_url
	}
	fido_authenticator_enabled = var.authentication_factor_setting_fido_authenticator_enabled
	hide_backup_factor_enabled = var.authentication_factor_setting_hide_backup_factor_enabled
	id = var.authentication_factor_setting_id
	identity_store_settings {

		#Optional
		mobile_number_enabled = var.authentication_factor_setting_identity_store_settings_mobile_number_enabled
		mobile_number_update_enabled = var.authentication_factor_setting_identity_store_settings_mobile_number_update_enabled
	}
	ocid = var.authentication_factor_setting_ocid
	phone_call_enabled = var.authentication_factor_setting_phone_call_enabled
	resource_type_schema_version = var.authentication_factor_setting_resource_type_schema_version
	tags {
		#Required
		key = var.authentication_factor_setting_tags_key
		value = var.authentication_factor_setting_tags_value
	}
	third_party_factor {
		#Required
		duo_security = var.authentication_factor_setting_third_party_factor_duo_security
	}
	urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings {
		#Required
		attestation = "NONE"
		authenticator_selection_attachment = "BOTH"
		authenticator_selection_require_resident_key = "false"
		authenticator_selection_resident_key = "NONE"
		authenticator_selection_user_verification = var.authentication_factor_setting_urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings_authenticator_selection_user_verification
		exclude_credentials = "false"
		public_key_types = ["RS1"]
		timeout = "60000"

		#Optional
		domain_validation_level = "1"
	}
	urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings {

		#Optional
		duo_security_settings {
			#Required
			api_hostname = var.authentication_factor_setting_urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings_duo_security_settings_api_hostname
			integration_key = var.authentication_factor_setting_urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings_duo_security_settings_integration_key
			secret_key = var.authentication_factor_setting_urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings_duo_security_settings_secret_key
			user_mapping_attribute = var.authentication_factor_setting_urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings_duo_security_settings_user_mapping_attribute

			#Optional
			attestation_key = var.authentication_factor_setting_urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings_duo_security_settings_attestation_key
		}
	}
	user_enrollment_disabled_factors = var.authentication_factor_setting_user_enrollment_disabled_factors
	yubico_otp_enabled = var.authentication_factor_setting_yubico_otp_enabled
}
```

## Argument Reference

The following arguments are supported:

* `attribute_sets` - (Optional) (Updatable) A multi-valued list of strings indicating the return type of attribute definition. The specified set of attributes can be fetched by the return type of the attribute. One or more values can be given together to fetch more than one group of attributes. If 'attributes' query parameter is also available, union of the two is fetched. Valid values - all, always, never, request, default. Values are case-insensitive.
* `attributes` - (Optional) (Updatable) A comma-delimited string that specifies the names of resource attributes that should be returned in the response. By default, a response that contains resource attributes contains only attributes that are defined in the schema for that resource type as returned=always or returned=default. An attribute that is defined as returned=request is returned in a response only if the request specifies its name in the value of this query parameter. If a request specifies this query parameter, the response contains the attributes that this query parameter specifies, as well as any attribute that is defined as returned=always.
* `authentication_factor_setting_id` - (Required) ID of the resource
* `authorization` - (Optional) (Updatable) The Authorization field value consists of credentials containing the authentication information of the user agent for the realm of the resource being requested.
* `auto_enroll_email_factor_disabled` - (Optional) (Updatable) If true, indicates that email will not be enrolled as a MFA factor automatically if it a account recovery factor

	**Added In:** 2011192329

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `bypass_code_enabled` - (Required) (Updatable) If true, indicates that Bypass Code is enabled for authentication

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: boolean
	* uniqueness: none
* `bypass_code_settings` - (Required) (Updatable) Settings related to the bypass code, such as bypass code length, bypass code expiry, max active bypass codes, and so on

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: complex
	* uniqueness: none
	* `help_desk_code_expiry_in_mins` - (Required) (Updatable) Expiry (in minutes) of any bypass code that is generated by the help desk

		**SCIM++ Properties:**
		* idcsMaxValue: 9999999
		* idcsMinValue: 1
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `help_desk_generation_enabled` - (Required) (Updatable) If true, indicates that help desk bypass code generation is enabled

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: boolean
		* uniqueness: none
	* `help_desk_max_usage` - (Required) (Updatable) The maximum number of times that any bypass code that is generated by the help desk can be used

		**SCIM++ Properties:**
		* idcsMaxValue: 999
		* idcsMinValue: 1
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `length` - (Required) (Updatable) Exact length of the bypass code to be generated

		**SCIM++ Properties:**
		* idcsMaxValue: 20
		* idcsMinValue: 8
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `max_active` - (Required) (Updatable) The maximum number of bypass codes that can be issued to any user

		**SCIM++ Properties:**
		* idcsMaxValue: 6
		* idcsMinValue: 1
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `self_service_generation_enabled` - (Required) (Updatable) If true, indicates that self-service bypass code generation is enabled

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: boolean
		* uniqueness: none
* `client_app_settings` - (Required) (Updatable) Settings related to compliance, Personal Identification Number (PIN) policy, and so on

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: complex
	* uniqueness: none
	* `device_protection_policy` - (Required) (Updatable) Indicates what protection policy that the system applies on a device. By default, the value is NONE, which indicates that the system applies no protection policy. A value of APP_PIN indicates that the system requires a Personal Identification Number (PIN). A value of DEVICE_BIOMETRIC_OR_APP_PIN indicates that either a PIN or a biometric authentication factor is required.

		**SCIM++ Properties:**
		* idcsCanonicalValueSourceFilter: attrName eq "deviceProtectionPolicy" and attrValues.value eq "$(deviceProtectionPolicy)"
		* idcsCanonicalValueSourceResourceType: AllowedValue
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `initial_lockout_period_in_secs` - (Required) (Updatable) The period of time in seconds that the system will lock a user out of the service after that user exceeds the maximum number of login failures

		**SCIM++ Properties:**
		* idcsMaxValue: 86400
		* idcsMinValue: 30
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `key_pair_length` - (Required) (Updatable) The size of the key that the system uses to generate the public-private key pair

		**SCIM++ Properties:**
		* idcsMaxValue: 4000
		* idcsMinValue: 32
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `lockout_escalation_pattern` - (Required) (Updatable) The pattern of escalation that the system follows, in locking a particular user out of the service.

		**SCIM++ Properties:**
		* idcsCanonicalValueSourceFilter: attrName eq "lockoutEscalationPattern" and attrValues.value eq "$(lockoutEscalationPattern)"
		* idcsCanonicalValueSourceResourceType: AllowedValue
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `max_failures_before_lockout` - (Required) (Updatable) The maximum number of times that a particular user can fail to login before the system locks that user out of the service

		**SCIM++ Properties:**
		* idcsMaxValue: 10
		* idcsMinValue: 5
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `max_failures_before_warning` - (Required) (Updatable) The maximum number of login failures that the system will allow before raising a warning and sending an alert via email

		**SCIM++ Properties:**
		* idcsMaxValue: 10
		* idcsMinValue: 0
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `max_lockout_interval_in_secs` - (Required) (Updatable) The maximum period of time that the system will lock a particular user out of the service regardless of what the configured pattern of escalation would otherwise dictate

		**SCIM++ Properties:**
		* idcsMaxValue: 86400
		* idcsMinValue: 30
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `min_pin_length` - (Required) (Updatable) Minimum length of the Personal Identification Number (PIN)

		**SCIM++ Properties:**
		* idcsMaxValue: 10
		* idcsMinValue: 6
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `policy_update_freq_in_days` - (Required) (Updatable) The period of time in days after which a client should refresh its policy by re-reading that policy from the server

		**SCIM++ Properties:**
		* idcsMaxValue: 999
		* idcsMinValue: 1
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `request_signing_algo` - (Required) (Updatable) Indicates which algorithm the system will use to sign requests

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `shared_secret_encoding` - (Required) (Updatable) Indicates the type of encoding that the system should use to generate a shared secret

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `unlock_app_for_each_request_enabled` - (Required) (Updatable) If true, indicates that the system should require the user to unlock the client app for each request. In order to unlock the App, the user must supply a Personal Identification Number (PIN) or a biometric authentication-factor.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: boolean
		* uniqueness: none
	* `unlock_app_interval_in_secs` - (Required) (Updatable) Specifies the period of time in seconds after which the client App should require the user to unlock the App. In order to unlock the App, the user must supply a Personal Identification Number (PIN) or a biometric authentication-factor. A value of zero means that it is disabled.

		**SCIM++ Properties:**
		* idcsMaxValue: 9999999
		* idcsMinValue: 0
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `unlock_on_app_foreground_enabled` - (Required) (Updatable) If true, indicates that the system should require the user to unlock the client App, when the client App comes to the foreground in the display of the device. In order to unlock the App, the user must supply a Personal Identification Number (PIN) or a biometric authentication-factor.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: boolean
		* uniqueness: none
	* `unlock_on_app_start_enabled` - (Required) (Updatable) If true, indicates that the system should require the user to unlock the client App whenever the App is started. In order to unlock the App, the user must supply a Personal Identification Number (PIN) or a biometric authentication-factor.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: boolean
		* uniqueness: none
* `compartment_ocid` - (Optional) (Updatable) Oracle Cloud Infrastructure Compartment Id (ocid) in which the resource lives.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `compliance_policy` - (Required) (Updatable) Compliance Policy that defines actions to be taken when a condition is violated

	**SCIM++ Properties:**
	* idcsCompositeKey: [name]
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
	* required: true
	* returned: default
	* type: complex
	* uniqueness: none
	* `action` - (Required) (Updatable) The action to be taken if the value of the attribute is not as expected

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `name` - (Required) (Updatable) The name of the attribute being evaluated

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - (Required) (Updatable) The value of the attribute to be evaluated

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `delete_in_progress` - (Optional) (Updatable) A boolean flag indicating this resource in the process of being deleted. Usually set to true when synchronous deletion of the resource would take too long.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `domain_ocid` - (Optional) (Updatable) Oracle Cloud Infrastructure Domain Id (ocid) in which the resource lives.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `email_enabled` - (Optional) (Updatable) If true, indicates that the EMAIL channel is enabled for authentication

	**Added In:** 18.1.2

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `email_settings` - (Optional) (Updatable) Settings related to Email Factor, such as enabled email magic link factor, custom url for Email Link

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `email_link_custom_url` - (Optional) (Updatable) Custom redirect Url which will be used in email link

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `email_link_enabled` - (Required) (Updatable) Specifies whether Email link is enabled or not.

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: boolean
		* uniqueness: none
* `endpoint_restrictions` - (Required) (Updatable) Settings that describe the set of restrictions that the system should apply to devices and trusted endpoints of a user

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: complex
	* uniqueness: none
	* `max_endpoint_trust_duration_in_days` - (Required) (Updatable) Maximum number of days until an endpoint can be trusted

		**SCIM++ Properties:**
		* idcsMaxValue: 180
		* idcsMinValue: 1
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `max_enrolled_devices` - (Required) (Updatable) Maximum number of enrolled devices per user

		**SCIM++ Properties:**
		* idcsMaxValue: 20
		* idcsMinValue: 1
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `max_incorrect_attempts` - (Required) (Updatable) An integer that represents the maximum number of failed MFA logins before an account is locked

		**SCIM++ Properties:**
		* idcsMaxValue: 20
		* idcsMinValue: 5
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `max_trusted_endpoints` - (Required) (Updatable) Max number of trusted endpoints per user

		**SCIM++ Properties:**
		* idcsMaxValue: 20
		* idcsMinValue: 1
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `trusted_endpoints_enabled` - (Required) (Updatable) Specify if trusted endpoints are enabled

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: boolean
		* uniqueness: none
* `fido_authenticator_enabled` - (Optional) (Updatable) If true, indicates that the Fido Authenticator channels are enabled for authentication

	**Added In:** 2009232244

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `hide_backup_factor_enabled` - (Optional) (Updatable) If true, indicates that 'Show backup factor(s)' button will be hidden during authentication

	**Added In:** 19.3.3

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `id` - (Optional) (Updatable) Unique identifier for the SCIM Resource as defined by the Service Provider. Each representation of the Resource MUST include a non-empty id value. This identifier MUST be unique across the Service Provider's entire set of Resources. It MUST be a stable, non-reassignable identifier that does not change when the same Resource is returned in subsequent requests. The value of the id attribute is always issued by the Service Provider and MUST never be specified by the Service Consumer. bulkId: is a reserved keyword and MUST NOT be used in the unique identifier.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: always
	* type: string
	* uniqueness: global
* `idcs_created_by` - (Optional) (Updatable) The User or App who created the Resource

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* required: true
	* returned: default
	* type: complex
	* `display` - (Optional) (Updatable) The displayName of the User or App who created this Resource

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ocid` - (Optional) (Updatable) The OCID of the SCIM resource that represents the User or App who created this Resource

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* returned: default
		* type: string
		* uniqueness: none
	* `ref` - (Optional) (Updatable) The URI of the SCIM resource that represents the User or App who created this Resource

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `type` - (Optional) (Updatable) The type of resource, User or App, that created this Resource

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - (Required) (Updatable) The ID of the SCIM resource that represents the User or App who created this Resource

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `idcs_endpoint` - (Required) The basic endpoint for the identity domain
* `idcs_last_modified_by` - (Optional) (Updatable) The User or App who modified the Resource

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: complex
	* `display` - (Optional) (Updatable) The displayName of the User or App who modified this Resource

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ocid` - (Optional) (Updatable) The OCID of the SCIM resource that represents the User or App who modified this Resource

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* returned: default
		* type: string
		* uniqueness: none
	* `ref` - (Optional) (Updatable) The URI of the SCIM resource that represents the User or App who modified this Resource

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `type` - (Optional) (Updatable) The type of resource, User or App, that modified this Resource

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - (Required) (Updatable) The ID of the SCIM resource that represents the User or App who modified this Resource

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `idcs_last_upgraded_in_release` - (Optional) (Updatable) The release number when the resource was upgraded.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: request
	* type: string
	* uniqueness: none
* `idcs_prevented_operations` - (Optional) (Updatable) Each value of this attribute specifies an operation that only an internal client may perform on this particular resource.

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: true
	* mutability: readOnly
	* required: false
	* returned: request
	* type: string
	* uniqueness: none
* `identity_store_settings` - (Optional) (Updatable) Settings related to the use of a user's profile details from the identity store

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `mobile_number_enabled` - (Optional) (Updatable) If true, indicates that Multi-Factor Authentication should use the mobile number in the identity store

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `mobile_number_update_enabled` - (Optional) (Updatable) If true, indicates that the user can update the mobile number in the user's Multi-Factor Authentication profile

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
* `meta` - (Optional) (Updatable) A complex attribute that contains resource metadata. All sub-attributes are OPTIONAL.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Created Date, mapsTo:meta.created]]
	* type: complex
	* `created` - (Optional) (Updatable) The DateTime the Resource was added to the Service Provider

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: dateTime
		* uniqueness: none
	* `last_modified` - (Optional) (Updatable) The most recent DateTime that the details of this Resource were updated at the Service Provider. If this Resource has never been modified since its initial creation, the value MUST be the same as the value of created. The attribute MUST be a DateTime.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: dateTime
		* uniqueness: none
	* `location` - (Optional) (Updatable) The URI of the Resource being returned. This value MUST be the same as the Location HTTP response header.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `resource_type` - (Optional) (Updatable) Name of the resource type of the resource--for example, Users or Groups

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `version` - (Optional) (Updatable) The version of the Resource being returned. This value must be the same as the ETag HTTP response header.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
* `mfa_enabled_category` - (Optional) (Updatable) Specifies the category of people for whom Multi-Factor Authentication is enabled. This is a readOnly attribute which reflects the value of mfaEnabledCategory attribute in SsoSettings

	**Deprecated Since: 18.1.2**

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `mfa_enrollment_type` - (Required) (Updatable) Specifies if Multi-Factor Authentication enrollment is mandatory or optional for a user

	**Deprecated Since: 18.1.2**

	**SCIM++ Properties:**
	* idcsCanonicalValueSourceFilter: attrName eq "mfaEnrollmentType" and attrValues.value eq "$(mfaEnrollmentType)"
	* idcsCanonicalValueSourceResourceType: AllowedValue
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: string
	* uniqueness: none
* `notification_settings` - (Required) (Updatable) Settings related to the Mobile App Notification channel, such as pull

	**Added In:** 17.4.2

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: complex
	* uniqueness: none
	* `pull_enabled` - (Required) (Updatable) If true, indicates that the Mobile App Pull Notification channel is enabled for authentication

		**Added In:** 17.4.2

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: boolean
		* uniqueness: none
* `ocid` - (Optional) (Updatable) Unique Oracle Cloud Infrastructure identifier for the SCIM Resource.

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: immutable
	* required: false
	* returned: default
	* type: string
	* uniqueness: global
* `phone_call_enabled` - (Optional) (Updatable) If true, indicates that the phone (PHONE_CALL) channel is enabled for authentication

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `push_enabled` - (Required) (Updatable) If true, indicates that the Mobile App Push Notification channel is enabled for authentication

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: boolean
	* uniqueness: none
* `resource_type_schema_version` - (Optional) (Updatable) An endpoint-specific schema version number to use in the Request. Allowed version values are Earliest Version or Latest Version as specified in each REST API endpoint description, or any sequential number inbetween. All schema attributes/body parameters are a part of version 1. After version 1, any attributes added or deprecated will be tagged with the version that they were added to or deprecated in. If no version is provided, the latest schema version is returned.
* `schemas` - (Required) (Updatable) REQUIRED. The schemas attribute is an array of Strings which allows introspection of the supported schema version for a SCIM representation as well any schema extensions supported by that representation. Each String value must be a unique URI. This specification defines URIs for User, Group, and a standard \"enterprise\" extension. All representations of SCIM schema MUST include a non-zero value array with value(s) of the URIs supported by that representation. Duplicate values MUST NOT be included. Value order is not specified and MUST not impact behavior.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
	* required: true
	* returned: default
	* type: string
	* uniqueness: none
* `security_questions_enabled` - (Required) (Updatable) If true, indicates that Security Questions are enabled for authentication

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: boolean
	* uniqueness: none
* `sms_enabled` - (Required) (Updatable) If true, indicates that the Short Message Service (SMS) channel is enabled for authentication

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: boolean
	* uniqueness: none
* `tags` - (Optional) (Updatable) A list of tags on this resource.

	**SCIM++ Properties:**
	* idcsCompositeKey: [key, value]
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: request
	* type: complex
	* uniqueness: none
	* `key` - (Required) (Updatable) Key or name of the tag.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - (Required) (Updatable) Value of the tag.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `tenancy_ocid` - (Optional) (Updatable) Oracle Cloud Infrastructure Tenant Id (ocid) in which the resource lives.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `third_party_factor` - (Optional) (Updatable) Settings related to third-party factor

	**Added In:** 19.2.1

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `duo_security` - (Required) (Updatable) To enable Duo Security factor

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: boolean
		* uniqueness: none
* `totp_enabled` - (Required) (Updatable) If true, indicates that the Mobile App One Time Passcode channel is enabled for authentication

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: boolean
	* uniqueness: none
* `totp_settings` - (Required) (Updatable) Settings related to Time-Based One-Time Passcodes (TOTP), such as hashing algo, totp time step, passcode length, and so on

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: complex
	* uniqueness: none
	* `email_otp_validity_duration_in_mins` - (Required) (Updatable) The period of time (in minutes) that a one-time passcode remains valid that the system sends by email.

		**Added In:** 18.1.2

		**SCIM++ Properties:**
		* idcsMaxValue: 60
		* idcsMinValue: 2
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `email_passcode_length` - (Required) (Updatable) Exact length of the email one-time passcode.

		**Added In:** 18.1.2

		**SCIM++ Properties:**
		* idcsMaxValue: 10
		* idcsMinValue: 4
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `hashing_algorithm` - (Required) (Updatable) The hashing algorithm to be used to calculate a One-Time Passcode. By default, the system uses SHA1.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `jwt_validity_duration_in_secs` - (Required) (Updatable) The period of time (in seconds) that a JSON Web Token (JWT) is valid

		**SCIM++ Properties:**
		* idcsMaxValue: 99999
		* idcsMinValue: 30
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `key_refresh_interval_in_days` - (Required) (Updatable) The duration of time (in days) after which the shared secret has to be refreshed

		**SCIM++ Properties:**
		* idcsMaxValue: 999
		* idcsMinValue: 30
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `passcode_length` - (Required) (Updatable) Exact length of the One-Time Passcode that the system should generate

		**SCIM++ Properties:**
		* idcsMaxValue: 10
		* idcsMinValue: 4
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `sms_otp_validity_duration_in_mins` - (Required) (Updatable) The period of time (in minutes) for which a One-Time Passcode that the system sends by Short Message Service (SMS) or by voice remains valid

		**SCIM++ Properties:**
		* idcsMaxValue: 60
		* idcsMinValue: 2
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `sms_passcode_length` - (Required) (Updatable) Exact length of the Short Message Service (SMS) One-Time Passcode

		**SCIM++ Properties:**
		* idcsMaxValue: 10
		* idcsMinValue: 4
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `time_step_in_secs` - (Required) (Updatable) Time (in secs) to be used as the time step

		**SCIM++ Properties:**
		* idcsMaxValue: 300
		* idcsMinValue: 30
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `time_step_tolerance` - (Required) (Updatable) The tolerance/step-size that the system should use when validating a One-Time Passcode

		**SCIM++ Properties:**
		* idcsMaxValue: 3
		* idcsMinValue: 2
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings` - (Optional) (Updatable) This extension defines attributes used to manage Multi-Factor Authentication settings of fido authentication
	* `attestation` - (Required) (Updatable) Attribute used to define the type of attestation required.

		**Added In:** 2009232244

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `authenticator_selection_attachment` - (Required) (Updatable) Attribute used to define authenticator selection attachment.

		**Added In:** 2009232244

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `authenticator_selection_require_resident_key` - (Required) (Updatable) Flag used to indicate authenticator selection is required or not

		**Added In:** 2009232244

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: boolean
		* uniqueness: none
	* `authenticator_selection_resident_key` - (Required) (Updatable) Attribute used to define authenticator selection resident key requirement.

		**Added In:** 2009232244

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `authenticator_selection_user_verification` - (Required) (Updatable) Attribute used to define authenticator selection verification.

		**Added In:** 2009232244

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `domain_validation_level` - (Optional) (Updatable) Number of domain levels IDCS should use for origin comparision

		**Added In:** 2109020413

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* idcsMaxValue: 2
		* idcsMinValue: 0
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `exclude_credentials` - (Required) (Updatable) Flag used to indicate whether we need to restrict creation of multiple credentials in same authenticator

		**Added In:** 2009232244

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: boolean
		* uniqueness: none
	* `public_key_types` - (Required) (Updatable) List of server supported public key algorithms

		**Added In:** 2009232244

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: true
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `timeout` - (Required) (Updatable) Timeout for the fido authentication to complete

		**Added In:** 2009232244

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* idcsMaxValue: 600000
		* idcsMinValue: 10000
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings` - (Optional) (Updatable) This extension defines attributes used to manage Multi-Factor Authentication settings of third party provider
	* `duo_security_settings` - (Optional) (Updatable) Settings related to Duo Security

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `api_hostname` - (Required) (Updatable) Hostname to access the Duo security account

			**Added In:** 19.2.1

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `attestation_key` - (Optional) (Updatable) Attestation key to attest the request and response between Duo Security

			**Added In:** 19.2.1

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: never
			* type: string
			* uniqueness: none
		* `integration_key` - (Required) (Updatable) Integration key from Duo Security authenticator

			**Added In:** 19.2.1

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `secret_key` - (Required) (Updatable) Secret key from Duo Security authenticator

			**Added In:** 19.2.1

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `user_mapping_attribute` - (Required) (Updatable) User attribute mapping value

			**Added In:** 19.2.1

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
* `user_enrollment_disabled_factors` - (Optional) (Updatable) Factors for which enrollment should be blocked for End User

	**Added In:** 2012271618

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `yubico_otp_enabled` - (Optional) (Updatable) If true, indicates that the Yubico OTP is enabled for authentication

	**Added In:** 2109090424

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `auto_enroll_email_factor_disabled` - If true, indicates that email will not be enrolled as a MFA factor automatically if it a account recovery factor

	**Added In:** 2011192329

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `bypass_code_enabled` - If true, indicates that Bypass Code is enabled for authentication

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: boolean
	* uniqueness: none
* `bypass_code_settings` - Settings related to the bypass code, such as bypass code length, bypass code expiry, max active bypass codes, and so on

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: complex
	* uniqueness: none
	* `help_desk_code_expiry_in_mins` - Expiry (in minutes) of any bypass code that is generated by the help desk

		**SCIM++ Properties:**
		* idcsMaxValue: 9999999
		* idcsMinValue: 1
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `help_desk_generation_enabled` - If true, indicates that help desk bypass code generation is enabled

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: boolean
		* uniqueness: none
	* `help_desk_max_usage` - The maximum number of times that any bypass code that is generated by the help desk can be used

		**SCIM++ Properties:**
		* idcsMaxValue: 999
		* idcsMinValue: 1
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `length` - Exact length of the bypass code to be generated

		**SCIM++ Properties:**
		* idcsMaxValue: 20
		* idcsMinValue: 8
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `max_active` - The maximum number of bypass codes that can be issued to any user

		**SCIM++ Properties:**
		* idcsMaxValue: 6
		* idcsMinValue: 1
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `self_service_generation_enabled` - If true, indicates that self-service bypass code generation is enabled

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: boolean
		* uniqueness: none
* `client_app_settings` - Settings related to compliance, Personal Identification Number (PIN) policy, and so on

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: complex
	* uniqueness: none
	* `device_protection_policy` - Indicates what protection policy that the system applies on a device. By default, the value is NONE, which indicates that the system applies no protection policy. A value of APP_PIN indicates that the system requires a Personal Identification Number (PIN). A value of DEVICE_BIOMETRIC_OR_APP_PIN indicates that either a PIN or a biometric authentication factor is required.

		**SCIM++ Properties:**
		* idcsCanonicalValueSourceFilter: attrName eq "deviceProtectionPolicy" and attrValues.value eq "$(deviceProtectionPolicy)"
		* idcsCanonicalValueSourceResourceType: AllowedValue
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `initial_lockout_period_in_secs` - The period of time in seconds that the system will lock a user out of the service after that user exceeds the maximum number of login failures

		**SCIM++ Properties:**
		* idcsMaxValue: 86400
		* idcsMinValue: 30
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `key_pair_length` - The size of the key that the system uses to generate the public-private key pair

		**SCIM++ Properties:**
		* idcsMaxValue: 4000
		* idcsMinValue: 32
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `lockout_escalation_pattern` - The pattern of escalation that the system follows, in locking a particular user out of the service.

		**SCIM++ Properties:**
		* idcsCanonicalValueSourceFilter: attrName eq "lockoutEscalationPattern" and attrValues.value eq "$(lockoutEscalationPattern)"
		* idcsCanonicalValueSourceResourceType: AllowedValue
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `max_failures_before_lockout` - The maximum number of times that a particular user can fail to login before the system locks that user out of the service

		**SCIM++ Properties:**
		* idcsMaxValue: 10
		* idcsMinValue: 5
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `max_failures_before_warning` - The maximum number of login failures that the system will allow before raising a warning and sending an alert via email

		**SCIM++ Properties:**
		* idcsMaxValue: 10
		* idcsMinValue: 0
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `max_lockout_interval_in_secs` - The maximum period of time that the system will lock a particular user out of the service regardless of what the configured pattern of escalation would otherwise dictate

		**SCIM++ Properties:**
		* idcsMaxValue: 86400
		* idcsMinValue: 30
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `min_pin_length` - Minimum length of the Personal Identification Number (PIN)

		**SCIM++ Properties:**
		* idcsMaxValue: 10
		* idcsMinValue: 6
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `policy_update_freq_in_days` - The period of time in days after which a client should refresh its policy by re-reading that policy from the server

		**SCIM++ Properties:**
		* idcsMaxValue: 999
		* idcsMinValue: 1
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `request_signing_algo` - Indicates which algorithm the system will use to sign requests

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `shared_secret_encoding` - Indicates the type of encoding that the system should use to generate a shared secret

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `unlock_app_for_each_request_enabled` - If true, indicates that the system should require the user to unlock the client app for each request. In order to unlock the App, the user must supply a Personal Identification Number (PIN) or a biometric authentication-factor.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: boolean
		* uniqueness: none
	* `unlock_app_interval_in_secs` - Specifies the period of time in seconds after which the client App should require the user to unlock the App. In order to unlock the App, the user must supply a Personal Identification Number (PIN) or a biometric authentication-factor. A value of zero means that it is disabled.

		**SCIM++ Properties:**
		* idcsMaxValue: 9999999
		* idcsMinValue: 0
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `unlock_on_app_foreground_enabled` - If true, indicates that the system should require the user to unlock the client App, when the client App comes to the foreground in the display of the device. In order to unlock the App, the user must supply a Personal Identification Number (PIN) or a biometric authentication-factor.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: boolean
		* uniqueness: none
	* `unlock_on_app_start_enabled` - If true, indicates that the system should require the user to unlock the client App whenever the App is started. In order to unlock the App, the user must supply a Personal Identification Number (PIN) or a biometric authentication-factor.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: boolean
		* uniqueness: none
* `compartment_ocid` - Oracle Cloud Infrastructure Compartment Id (ocid) in which the resource lives.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `compliance_policy` - Compliance Policy that defines actions to be taken when a condition is violated

	**SCIM++ Properties:**
	* idcsCompositeKey: [name]
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
	* required: true
	* returned: default
	* type: complex
	* uniqueness: none
	* `action` - The action to be taken if the value of the attribute is not as expected

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `name` - The name of the attribute being evaluated

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - The value of the attribute to be evaluated

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `delete_in_progress` - A boolean flag indicating this resource in the process of being deleted. Usually set to true when synchronous deletion of the resource would take too long.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `domain_ocid` - Oracle Cloud Infrastructure Domain Id (ocid) in which the resource lives.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `email_enabled` - If true, indicates that the EMAIL channel is enabled for authentication

	**Added In:** 18.1.2

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `email_settings` - Settings related to Email Factor, such as enabled email magic link factor, custom url for Email Link

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `email_link_custom_url` - Custom redirect Url which will be used in email link

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `email_link_enabled` - Specifies whether Email link is enabled or not.

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: boolean
		* uniqueness: none
* `endpoint_restrictions` - Settings that describe the set of restrictions that the system should apply to devices and trusted endpoints of a user

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: complex
	* uniqueness: none
	* `max_endpoint_trust_duration_in_days` - Maximum number of days until an endpoint can be trusted

		**SCIM++ Properties:**
		* idcsMaxValue: 180
		* idcsMinValue: 1
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `max_enrolled_devices` - Maximum number of enrolled devices per user

		**SCIM++ Properties:**
		* idcsMaxValue: 20
		* idcsMinValue: 1
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `max_incorrect_attempts` - An integer that represents the maximum number of failed MFA logins before an account is locked

		**SCIM++ Properties:**
		* idcsMaxValue: 20
		* idcsMinValue: 5
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `max_trusted_endpoints` - Max number of trusted endpoints per user

		**SCIM++ Properties:**
		* idcsMaxValue: 20
		* idcsMinValue: 1
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `trusted_endpoints_enabled` - Specify if trusted endpoints are enabled

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: boolean
		* uniqueness: none
* `fido_authenticator_enabled` - If true, indicates that the Fido Authenticator channels are enabled for authentication

	**Added In:** 2009232244

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `hide_backup_factor_enabled` - If true, indicates that 'Show backup factor(s)' button will be hidden during authentication

	**Added In:** 19.3.3

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `id` - Unique identifier for the SCIM Resource as defined by the Service Provider. Each representation of the Resource MUST include a non-empty id value. This identifier MUST be unique across the Service Provider's entire set of Resources. It MUST be a stable, non-reassignable identifier that does not change when the same Resource is returned in subsequent requests. The value of the id attribute is always issued by the Service Provider and MUST never be specified by the Service Consumer. bulkId: is a reserved keyword and MUST NOT be used in the unique identifier.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: always
	* type: string
	* uniqueness: global
* `idcs_created_by` - The User or App who created the Resource

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* required: true
	* returned: default
	* type: complex
	* `display` - The displayName of the User or App who created this Resource

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ocid` - The OCID of the SCIM resource that represents the User or App who created this Resource

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* returned: default
		* type: string
		* uniqueness: none
	* `ref` - The URI of the SCIM resource that represents the User or App who created this Resource

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `type` - The type of resource, User or App, that created this Resource

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - The ID of the SCIM resource that represents the User or App who created this Resource

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `idcs_last_modified_by` - The User or App who modified the Resource

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: complex
	* `display` - The displayName of the User or App who modified this Resource

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ocid` - The OCID of the SCIM resource that represents the User or App who modified this Resource

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* returned: default
		* type: string
		* uniqueness: none
	* `ref` - The URI of the SCIM resource that represents the User or App who modified this Resource

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `type` - The type of resource, User or App, that modified this Resource

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - The ID of the SCIM resource that represents the User or App who modified this Resource

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `idcs_last_upgraded_in_release` - The release number when the resource was upgraded.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: request
	* type: string
	* uniqueness: none
* `idcs_prevented_operations` - Each value of this attribute specifies an operation that only an internal client may perform on this particular resource.

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: true
	* mutability: readOnly
	* required: false
	* returned: request
	* type: string
	* uniqueness: none
* `identity_store_settings` - Settings related to the use of a user's profile details from the identity store

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `mobile_number_enabled` - If true, indicates that Multi-Factor Authentication should use the mobile number in the identity store

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `mobile_number_update_enabled` - If true, indicates that the user can update the mobile number in the user's Multi-Factor Authentication profile

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
* `meta` - A complex attribute that contains resource metadata. All sub-attributes are OPTIONAL.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Created Date, mapsTo:meta.created]]
	* type: complex
	* `created` - The DateTime the Resource was added to the Service Provider

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: dateTime
		* uniqueness: none
	* `last_modified` - The most recent DateTime that the details of this Resource were updated at the Service Provider. If this Resource has never been modified since its initial creation, the value MUST be the same as the value of created. The attribute MUST be a DateTime.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: dateTime
		* uniqueness: none
	* `location` - The URI of the Resource being returned. This value MUST be the same as the Location HTTP response header.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `resource_type` - Name of the resource type of the resource--for example, Users or Groups

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `version` - The version of the Resource being returned. This value must be the same as the ETag HTTP response header.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
* `mfa_enabled_category` - Specifies the category of people for whom Multi-Factor Authentication is enabled. This is a readOnly attribute which reflects the value of mfaEnabledCategory attribute in SsoSettings

	**Deprecated Since: 18.1.2**

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `mfa_enrollment_type` - Specifies if Multi-Factor Authentication enrollment is mandatory or optional for a user

	**Deprecated Since: 18.1.2**

	**SCIM++ Properties:**
	* idcsCanonicalValueSourceFilter: attrName eq "mfaEnrollmentType" and attrValues.value eq "$(mfaEnrollmentType)"
	* idcsCanonicalValueSourceResourceType: AllowedValue
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: string
	* uniqueness: none
* `notification_settings` - Settings related to the Mobile App Notification channel, such as pull

	**Added In:** 17.4.2

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: complex
	* uniqueness: none
	* `pull_enabled` - If true, indicates that the Mobile App Pull Notification channel is enabled for authentication

		**Added In:** 17.4.2

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: boolean
		* uniqueness: none
* `ocid` - Unique Oracle Cloud Infrastructure identifier for the SCIM Resource.

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: immutable
	* required: false
	* returned: default
	* type: string
	* uniqueness: global
* `phone_call_enabled` - If true, indicates that the phone (PHONE_CALL) channel is enabled for authentication

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `push_enabled` - If true, indicates that the Mobile App Push Notification channel is enabled for authentication

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: boolean
	* uniqueness: none
* `schemas` - REQUIRED. The schemas attribute is an array of Strings which allows introspection of the supported schema version for a SCIM representation as well any schema extensions supported by that representation. Each String value must be a unique URI. This specification defines URIs for User, Group, and a standard \"enterprise\" extension. All representations of SCIM schema MUST include a non-zero value array with value(s) of the URIs supported by that representation. Duplicate values MUST NOT be included. Value order is not specified and MUST not impact behavior.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
	* required: true
	* returned: default
	* type: string
	* uniqueness: none
* `security_questions_enabled` - If true, indicates that Security Questions are enabled for authentication

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: boolean
	* uniqueness: none
* `sms_enabled` - If true, indicates that the Short Message Service (SMS) channel is enabled for authentication

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: boolean
	* uniqueness: none
* `tags` - A list of tags on this resource.

	**SCIM++ Properties:**
	* idcsCompositeKey: [key, value]
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: request
	* type: complex
	* uniqueness: none
	* `key` - Key or name of the tag.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - Value of the tag.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `tenancy_ocid` - Oracle Cloud Infrastructure Tenant Id (ocid) in which the resource lives.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `third_party_factor` - Settings related to third-party factor

	**Added In:** 19.2.1

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `duo_security` - To enable Duo Security factor

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: boolean
		* uniqueness: none
* `totp_enabled` - If true, indicates that the Mobile App One Time Passcode channel is enabled for authentication

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: boolean
	* uniqueness: none
* `totp_settings` - Settings related to Time-Based One-Time Passcodes (TOTP), such as hashing algo, totp time step, passcode length, and so on

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: complex
	* uniqueness: none
	* `email_otp_validity_duration_in_mins` - The period of time (in minutes) that a one-time passcode remains valid that the system sends by email.

		**Added In:** 18.1.2

		**SCIM++ Properties:**
		* idcsMaxValue: 60
		* idcsMinValue: 2
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `email_passcode_length` - Exact length of the email one-time passcode.

		**Added In:** 18.1.2

		**SCIM++ Properties:**
		* idcsMaxValue: 10
		* idcsMinValue: 4
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `hashing_algorithm` - The hashing algorithm to be used to calculate a One-Time Passcode. By default, the system uses SHA1.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `jwt_validity_duration_in_secs` - The period of time (in seconds) that a JSON Web Token (JWT) is valid

		**SCIM++ Properties:**
		* idcsMaxValue: 99999
		* idcsMinValue: 30
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `key_refresh_interval_in_days` - The duration of time (in days) after which the shared secret has to be refreshed

		**SCIM++ Properties:**
		* idcsMaxValue: 999
		* idcsMinValue: 30
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `passcode_length` - Exact length of the One-Time Passcode that the system should generate

		**SCIM++ Properties:**
		* idcsMaxValue: 10
		* idcsMinValue: 4
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `sms_otp_validity_duration_in_mins` - The period of time (in minutes) for which a One-Time Passcode that the system sends by Short Message Service (SMS) or by voice remains valid

		**SCIM++ Properties:**
		* idcsMaxValue: 60
		* idcsMinValue: 2
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `sms_passcode_length` - Exact length of the Short Message Service (SMS) One-Time Passcode

		**SCIM++ Properties:**
		* idcsMaxValue: 10
		* idcsMinValue: 4
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `time_step_in_secs` - Time (in secs) to be used as the time step

		**SCIM++ Properties:**
		* idcsMaxValue: 300
		* idcsMinValue: 30
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
	* `time_step_tolerance` - The tolerance/step-size that the system should use when validating a One-Time Passcode

		**SCIM++ Properties:**
		* idcsMaxValue: 3
		* idcsMinValue: 2
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionfido_authentication_factor_settings` - This extension defines attributes used to manage Multi-Factor Authentication settings of fido authentication
	* `attestation` - Attribute used to define the type of attestation required.

		**Added In:** 2009232244

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `authenticator_selection_attachment` - Attribute used to define authenticator selection attachment.

		**Added In:** 2009232244

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `authenticator_selection_require_resident_key` - Flag used to indicate authenticator selection is required or not

		**Added In:** 2009232244

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: boolean
		* uniqueness: none
	* `authenticator_selection_resident_key` - Attribute used to define authenticator selection resident key requirement.

		**Added In:** 2009232244

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `authenticator_selection_user_verification` - Attribute used to define authenticator selection verification.

		**Added In:** 2009232244

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `domain_validation_level` - Number of domain levels IDCS should use for origin comparision

		**Added In:** 2109020413

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* idcsMaxValue: 2
		* idcsMinValue: 0
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `exclude_credentials` - Flag used to indicate whether we need to restrict creation of multiple credentials in same authenticator

		**Added In:** 2009232244

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: boolean
		* uniqueness: none
	* `public_key_types` - List of server supported public key algorithms

		**Added In:** 2009232244

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: true
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `timeout` - Timeout for the fido authentication to complete

		**Added In:** 2009232244

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* idcsMaxValue: 600000
		* idcsMinValue: 10000
		* required: true
		* returned: default
		* type: integer
		* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionthird_party_authentication_factor_settings` - This extension defines attributes used to manage Multi-Factor Authentication settings of third party provider
	* `duo_security_settings` - Settings related to Duo Security

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `api_hostname` - Hostname to access the Duo security account

			**Added In:** 19.2.1

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `attestation_key` - Attestation key to attest the request and response between Duo Security

			**Added In:** 19.2.1

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: false
			* returned: never
			* type: string
			* uniqueness: none
		* `integration_key` - Integration key from Duo Security authenticator

			**Added In:** 19.2.1

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `secret_key` - Secret key from Duo Security authenticator

			**Added In:** 19.2.1

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `user_mapping_attribute` - User attribute mapping value

			**Added In:** 19.2.1

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
* `user_enrollment_disabled_factors` - Factors for which enrollment should be blocked for End User

	**Added In:** 2012271618

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `yubico_otp_enabled` - If true, indicates that the Yubico OTP is enabled for authentication

	**Added In:** 2109090424

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Authentication Factor Setting
	* `update` - (Defaults to 20 minutes), when updating the Authentication Factor Setting
	* `delete` - (Defaults to 20 minutes), when destroying the Authentication Factor Setting


## Import

AuthenticationFactorSettings can be imported using the `id`, e.g.

```
$ terraform import oci_identity_domains_authentication_factor_setting.test_authentication_factor_setting "idcsEndpoint/{idcsEndpoint}/authenticationFactorSettings/{authenticationFactorSettingId}" 
```

