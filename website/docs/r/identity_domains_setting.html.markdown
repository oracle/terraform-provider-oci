---
subcategory: "Identity Domains"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_domains_setting"
sidebar_current: "docs-oci-resource-identity_domains-setting"
description: |-
  Provides the Setting resource in Oracle Cloud Infrastructure Identity Domains service
---

# oci_identity_domains_setting
This resource provides the Setting resource in Oracle Cloud Infrastructure Identity Domains service.

Replace Settings

## Example Usage

```hcl
resource "oci_identity_domains_setting" "test_setting" {
	#Required
	csr_access = var.setting_csr_access
	idcs_endpoint = data.oci_identity_domain.test_domain.url
	schemas = ["urn:ietf:params:scim:schemas:oracle:idcs:Settings"]
	setting_id = "Settings"

	#Optional
	account_always_trust_scope = var.setting_account_always_trust_scope
	allowed_domains = ["test.com"]
	allowed_forgot_password_flow_return_urls = var.setting_allowed_forgot_password_flow_return_urls
	allowed_notification_redirect_urls = var.setting_allowed_notification_redirect_urls
	attribute_sets = ["all"]
	attributes = ""
	audit_event_retention_period = var.setting_audit_event_retention_period
	authorization = var.setting_authorization
	certificate_validation {

		#Optional
		crl_check_on_ocsp_failure_enabled = var.setting_certificate_validation_crl_check_on_ocsp_failure_enabled
		crl_enabled = var.setting_certificate_validation_crl_enabled
		crl_location = var.setting_certificate_validation_crl_location
		crl_refresh_interval = var.setting_certificate_validation_crl_refresh_interval
		ocsp_enabled = var.setting_certificate_validation_ocsp_enabled
		ocsp_responder_url = var.setting_certificate_validation_ocsp_responder_url
		ocsp_settings_responder_url_preferred = var.setting_certificate_validation_ocsp_settings_responder_url_preferred
		ocsp_signing_certificate_alias = var.setting_certificate_validation_ocsp_signing_certificate_alias
		ocsp_timeout_duration = var.setting_certificate_validation_ocsp_timeout_duration
		ocsp_unknown_response_status_allowed = var.setting_certificate_validation_ocsp_unknown_response_status_allowed
	}
	cloud_gate_cors_settings {

		#Optional
		cloud_gate_cors_allow_null_origin = var.setting_cloud_gate_cors_settings_cloud_gate_cors_allow_null_origin
		cloud_gate_cors_allowed_origins = ["https://test.com"]
		cloud_gate_cors_enabled = var.setting_cloud_gate_cors_settings_cloud_gate_cors_enabled
		cloud_gate_cors_exposed_headers = var.setting_cloud_gate_cors_settings_cloud_gate_cors_exposed_headers
		cloud_gate_cors_max_age = var.setting_cloud_gate_cors_settings_cloud_gate_cors_max_age
	}
	cloud_migration_custom_url = var.setting_cloud_migration_custom_url
	cloud_migration_url_enabled = var.setting_cloud_migration_url_enabled
	company_names {
		#Required
		locale = var.setting_company_names_locale
		value = var.setting_company_names_value
	}
	contact_emails = ["contactEmails@test.com"]
	custom_branding = var.setting_custom_branding
	custom_css_location = var.setting_custom_css_location
	custom_html_location = var.setting_custom_html_location
	custom_translation = var.setting_custom_translation
	default_trust_scope = var.setting_default_trust_scope
	diagnostic_level = var.setting_diagnostic_level
	diagnostic_record_for_search_identifies_returned_resources = var.setting_diagnostic_record_for_search_identifies_returned_resources
	enable_terms_of_use = var.setting_enable_terms_of_use
	external_id = "externalId"
	iam_upst_session_expiry = var.setting_iam_upst_session_expiry
	id = var.setting_id
	images {
		#Required
		type = var.setting_images_type
		value = var.setting_images_value

		#Optional
		display = var.setting_images_display
	}
	is_hosted_page = var.setting_is_hosted_page
	issuer = var.setting_issuer
	locale = var.setting_locale
	login_texts {
		#Required
		locale = var.setting_login_texts_locale
		value = var.setting_login_texts_value
	}
	max_no_of_app_cmva_to_return = var.setting_max_no_of_app_cmva_to_return
	max_no_of_app_role_members_to_return = var.setting_max_no_of_app_role_members_to_return
	ocid = var.setting_ocid
	preferred_language = var.setting_preferred_language
	prev_issuer = var.setting_prev_issuer
	privacy_policy_url = var.setting_privacy_policy_url
	purge_configs {
		#Required
		resource_name = "resourceName"
		retention_period = var.setting_purge_configs_retention_period
	}
	re_auth_factor = ["password"]
	re_auth_when_changing_my_authentication_factors = var.setting_re_auth_when_changing_my_authentication_factors
	resource_type_schema_version = var.setting_resource_type_schema_version
	service_admin_cannot_list_other_users = var.setting_service_admin_cannot_list_other_users
	signing_cert_public_access = var.setting_signing_cert_public_access
	sub_mapping_attr = var.setting_sub_mapping_attr
	tags {
		#Required
		key = var.setting_tags_key
		value = var.setting_tags_value
	}
	tenant_custom_claims {
		#Required
		all_scopes = var.setting_tenant_custom_claims_all_scopes
		expression = var.setting_tenant_custom_claims_expression
		mode = var.setting_tenant_custom_claims_mode
		name = var.setting_tenant_custom_claims_name
		token_type = var.setting_tenant_custom_claims_token_type
		value = var.setting_tenant_custom_claims_value

		#Optional
		scopes = ["scopes"]
	}
	terms_of_use_url = var.setting_terms_of_use_url
	timezone = var.setting_timezone
}
```

## Argument Reference

The following arguments are supported:

* `account_always_trust_scope` - (Optional) (Updatable) Indicates whether all the Apps in this customer tenancy should trust each other. A value of true overrides the 'defaultTrustScope' attribute here in Settings, as well as any App-specific 'trustScope' attribute, to force in effect 'trustScope=Account' for every App in this customer tenancy.

	**Added In:** 18.1.6

	**SCIM++ Properties:**
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
* `allowed_domains` - (Optional) (Updatable) One or more email domains allowed in a user's email field. If unassigned, any domain is allowed.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `allowed_forgot_password_flow_return_urls` - (Optional) (Updatable) If specified, indicates the set of Urls which can be returned to after successful forgot password flow

	**Added In:** 19.3.3

	**SCIM++ Properties:**
	* type: string
	* multiValued: true
	* required: false
	* mutability: readWrite
	* returned: default
	* uniqueness: none
	* caseExact: false
* `allowed_notification_redirect_urls` - (Optional) (Updatable) If specified, indicates the set of allowed notification redirect Urls which can be specified as the value of \"notificationRedirectUrl\" in the POST .../admin/v1/MePasswordResetRequestor request payload, which will then be included in the reset password email notification sent to a user as part of the forgot password / password reset flow.

	**Added In:** 2009041201

	**SCIM++ Properties:**
	* type: string
	* multiValued: true
	* required: false
	* mutability: readWrite
	* returned: default
	* uniqueness: none
	* caseExact: false
* `attribute_sets` - (Optional) (Updatable) A multi-valued list of strings indicating the return type of attribute definition. The specified set of attributes can be fetched by the return type of the attribute. One or more values can be given together to fetch more than one group of attributes. If 'attributes' query parameter is also available, union of the two is fetched. Valid values - all, always, never, request, default. Values are case-insensitive.
* `attributes` - (Optional) (Updatable) A comma-delimited string that specifies the names of resource attributes that should be returned in the response. By default, a response that contains resource attributes contains only attributes that are defined in the schema for that resource type as returned=always or returned=default. An attribute that is defined as returned=request is returned in a response only if the request specifies its name in the value of this query parameter. If a request specifies this query parameter, the response contains the attributes that this query parameter specifies, as well as any attribute that is defined as returned=always.
* `audit_event_retention_period` - (Optional) (Updatable) Audit Event retention period. If set, overrides default of 30 days after which Audit Events will be purged

	**Added In:** 19.2.1

	**SCIM++ Properties:**
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
* `authorization` - (Optional) (Updatable) The Authorization field value consists of credentials containing the authentication information of the user agent for the realm of the resource being requested.
* `certificate_validation` - (Optional) (Updatable) Certificate Validation Config

	**Added In:** 2010242156

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `crl_check_on_ocsp_failure_enabled` - (Optional) (Updatable) Use CRL as Fallback.

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `crl_enabled` - (Optional) (Updatable) CRL is enabled Configuration

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `crl_location` - (Optional) (Updatable) CRL Location.

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `crl_refresh_interval` - (Optional) (Updatable) The CRL refresh interval in minutes

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `ocsp_enabled` - (Optional) (Updatable) OCSP is enabled Configuration

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `ocsp_responder_url` - (Optional) (Updatable) OCSP Responder URL

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ocsp_settings_responder_url_preferred` - (Optional) (Updatable) This setting says, OCSP Responder URL present in the issued certificate must be used. Otherwise, OCSP Responder URL from IDP or Settings.

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `ocsp_signing_certificate_alias` - (Optional) (Updatable) OCSP Signing Certificate Alias

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ocsp_timeout_duration` - (Optional) (Updatable) The OCSP Timeout duration in minutes

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* idcsMaxValue: 10
		* idcsMinValue: 1
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `ocsp_unknown_response_status_allowed` - (Optional) (Updatable) OCSP Accept unknown response status from ocsp responder.

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
* `cloud_account_name` - (Optional) (Updatable) The attribute to store the cloud account name

	**Deprecated Since: 2011192329**

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `cloud_gate_cors_settings` - (Optional) (Updatable) A complex attribute that specifies the Cloud Gate cross origin resource sharing settings.

	**Added In:** 2011192329

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `cloud_gate_cors_allow_null_origin` - (Optional) (Updatable) Allow Null Origin (CORS) for this tenant.

		**Added In:** 2011192329

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `cloud_gate_cors_allowed_origins` - (Optional) (Updatable) Cloud Gate Allowed Cross-Origin Resource Sharing (CORS) Origins for this tenant.

		**Added In:** 2011192329

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `cloud_gate_cors_enabled` - (Optional) (Updatable) Enable Cloud Gate Cross-Origin Resource Sharing (CORS) for this tenant.

		**Added In:** 2011192329

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `cloud_gate_cors_exposed_headers` - (Optional) (Updatable) List of Response Headers Cloud Gate is allowed to expose in the CORS Response Header: Access-Control-Expose-Headers.

		**Added In:** 2205182039

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `cloud_gate_cors_max_age` - (Optional) (Updatable) Maximum number of seconds a CORS Pre-flight Response may be cached by client.

		**Added In:** 2205182039

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
* `cloud_migration_custom_url` - (Optional) (Updatable) If specified, indicates the custom SIM Migrator Url which can be used while SIM to Oracle Identity Cloud Service CloudAccount Migration.

	**Added In:** 2012271618

	**SCIM++ Properties:**
	* type: string
	* multiValued: false
	* required: false
	* mutability: readWrite
	* returned: default
	* uniqueness: none
	* caseExact: false
* `cloud_migration_url_enabled` - (Optional) (Updatable) CloudAccountMigration: Enable Custom SIM Migrator Url.

	**Added In:** 2012271618

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `company_names` - (Optional) (Updatable) Name of the company in different locales

	**SCIM++ Properties:**
	* idcsCompositeKey: [locale]
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* `locale` - (Required) (Updatable) Locale

		**SCIM++ Properties:**
		* idcsCanonicalValueSourceFilter: attrName eq "locales" and attrValues.value eq "$(companyNames.locale)"
		* idcsCanonicalValueSourceResourceType: AllowedValue
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
	* `value` - (Required) (Updatable) Company name

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
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
* `contact_emails` - (Optional) (Updatable) Contact emails used to notify tenants. Can be one or more user or group alias emails.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `csr_access` - (Required) (Updatable) This value indicates whether Customer Service Representatives can login and have readOnly or readWrite access.  A value of 'none' means CSR cannot login to the services.

	**SCIM++ Properties:**
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: string
* `custom_branding` - (Optional) (Updatable) Indicates if the branding is default or custom

	**SCIM++ Properties:**
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
* `custom_css_location` - (Optional) (Updatable) Storage URL location where the sanitized custom css is located

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `custom_html_location` - (Optional) (Updatable) Storage URL location where the sanitized custom html is located

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `custom_translation` - (Optional) (Updatable) Custom translations (JSON String)

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `default_company_names` - (Optional) (Updatable) Default name of the Company in different locales

	**Added In:** 18.2.2

	**SCIM++ Properties:**
	* idcsCompositeKey: [locale]
	* multiValued: true
	* mutability: readOnly
	* required: false
	* returned: default
	* type: complex
	* `locale` - (Required) (Updatable) Locale

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* idcsCanonicalValueSourceFilter: attrName eq "locales" and attrValues.value eq "$(companyNames.locale)"
		* idcsCanonicalValueSourceResourceType: AllowedValue
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: string
	* `value` - (Required) (Updatable) Company name

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: string
* `default_images` - (Optional) (Updatable) References to various images

	**Added In:** 18.2.2

	**SCIM++ Properties:**
	* idcsCompositeKey: [type]
	* multiValued: true
	* mutability: readOnly
	* required: false
	* returned: default
	* type: complex
	* `display` - (Optional) (Updatable) A human-readable name, primarily used for display purposes

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
	* `type` - (Required) (Updatable) Indicates the image type

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: string
	* `value` - (Required) (Updatable) Image URI

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: reference
* `default_login_texts` - (Optional) (Updatable) Default Login text in different locales

	**Added In:** 18.2.2

	**SCIM++ Properties:**
	* idcsCompositeKey: [locale]
	* multiValued: true
	* mutability: readOnly
	* required: false
	* returned: default
	* type: complex
	* `locale` - (Required) (Updatable) Locale

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* idcsCanonicalValueSourceFilter: attrName eq "locales" and attrValues.value eq "$(loginTexts.locale)"
		* idcsCanonicalValueSourceResourceType: AllowedValue
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: string
	* `value` - (Required) (Updatable) Login text

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: string
		* idcsSanitize: true
* `default_trust_scope` - (Optional) (Updatable) **Deprecated Since: 18.3.6**

	**SCIM++ Properties:**
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string Indicates the default trust scope for all apps 
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
* `diagnostic_level` - (Optional) (Updatable) The level of diagnostic logging that is currently in effect. A level of 0 (zero) indicates that diagnostic logging is disabled. A level of 1 (one) indicates that diagnostic logging is enabled.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `diagnostic_record_for_search_identifies_returned_resources` - (Optional) (Updatable) Controls whether DiagnosticRecords for external search-operations (against SCIM resource-types in the Admin service) identify returned resources.  If true, indicates that for each successful external search-operation at least one DiagnosticRecord will include at least one identifier for each matching resource that is returned in that search-response.  If false, no DiagnosticRecord should be expected to identify returned resources for a search-operation.  The default value is false.

	**Added In:** 2011192329

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `diagnostic_tracing_upto` - (Optional) (Updatable) The end time up to which diagnostic recording is switched on

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: dateTime
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
* `enable_terms_of_use` - (Optional) (Updatable) Indicates if Terms of Use is enabled in UI

	**Added In:** 18.2.4

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `external_id` - (Optional) (Updatable) An identifier for the Resource as defined by the Service Consumer. The externalId may simplify identification of the Resource between Service Consumer and Service Provider by allowing the Consumer to refer to the Resource with its own identifier, obviating the need to store a local mapping between the local identifier of the Resource and the identifier used by the Service Provider. Each Resource MAY include a non-empty externalId value. The value of the externalId attribute is always issued by the Service Consumer and can never be specified by the Service Provider. The Service Provider MUST always interpret the externalId as scoped to the Service Consumer's tenant.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `iam_upst_session_expiry` - (Optional) (Updatable) Maximum duration for IAM User Principal Session Token expiry

	**Added In:** 2307071836

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
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
* `images` - (Optional) (Updatable) References to various images

	**SCIM++ Properties:**
	* idcsCompositeKey: [type]
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* `display` - (Optional) (Updatable) A human-readable name, primarily used for display purposes

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
	* `type` - (Required) (Updatable) Indicates the image type

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
	* `value` - (Required) (Updatable) Image URI

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: reference
* `is_hosted_page` - (Optional) (Updatable) Indicates if 'hosted' option was selected

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `issuer` - (Optional) (Updatable) Tenant issuer.

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `locale` - (Optional) (Updatable) Default location for purposes of localizing items such as currency, date and time format, numerical representations, and so on.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCanonicalValueSourceFilter: attrName eq "locales" and attrValues.value eq "$(locale)"
	* idcsCanonicalValueSourceResourceType: AllowedValue
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `login_texts` - (Optional) (Updatable) Login text in different locales

	**SCIM++ Properties:**
	* idcsCompositeKey: [locale]
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* `locale` - (Required) (Updatable) Locale

		**SCIM++ Properties:**
		* idcsCanonicalValueSourceFilter: attrName eq "locales" and attrValues.value eq "$(loginTexts.locale)"
		* idcsCanonicalValueSourceResourceType: AllowedValue
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
	* `value` - (Required) (Updatable) Login text

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* idcsSanitize: true
		* type: string
* `max_no_of_app_cmva_to_return` - (Optional) (Updatable) Limit the maximum return of CMVA for an App

	**Added In:** 2111112015

	**SCIM++ Properties:**
	* idcsMinValue: 0
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `max_no_of_app_role_members_to_return` - (Optional) (Updatable) Limit the maximum return of members for an AppRole

	**Added In:** 2111112015

	**SCIM++ Properties:**
	* idcsMinValue: 0
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
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
* `migration_status` - (Optional) (Updatable) Database Migration Status

	**Added In:** 19.2.1

	**SCIM++ Properties:**
	* caseExact: true
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: string
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
* `on_premises_provisioning` - (Optional) (Updatable) On-Premises provisioning feature toggle.

	**Added In:** 19.2.1

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `preferred_language` - (Optional) (Updatable) Preferred written or spoken language used for localized user interfaces

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCanonicalValueSourceFilter: attrName eq "languages" and attrValues.value eq "$(preferredLanguage)"
	* idcsCanonicalValueSourceResourceType: AllowedValue
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `prev_issuer` - (Optional) (Updatable) Previous Tenant issuer. This is an Oracle Identity Cloud Service internal attribute which is not meant to be directly modified by ID Admin. Even if the request body (Settings) contains this attribute, the actual value will be set according to the Oracle Identity Cloud Service internal logic rather than solely based on the value provided in the request payload.

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: request
	* type: string
	* uniqueness: none
* `privacy_policy_url` - (Optional) (Updatable) Privacy Policy URL

	**Added In:** 18.2.4

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `purge_configs` - (Optional) (Updatable) Purge Configs for different Resource Types

	**Deprecated Since: 19.1.6**

	**SCIM++ Properties:**
	* idcsCompositeKey: [resourceName]
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* `resource_name` - (Required) (Updatable) Resource Name

		**Deprecated Since: 19.1.6**

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
	* `retention_period` - (Required) (Updatable) Retention Period

		**Deprecated Since: 19.1.6**

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
* `re_auth_factor` - (Optional) (Updatable) If reAuthWhenChangingMyAuthenticationFactors is true (default), this attribute specifies which re-authentication factor to use. Allowed value is \"password\".

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
* `re_auth_when_changing_my_authentication_factors` - (Optional) (Updatable) Specifies whether re-authentication is required or not when a user changes one of their security factors such as password or email. Default is true to ensure more secure behavior.

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
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
* `service_admin_cannot_list_other_users` - (Optional) (Updatable) By default, a service admin can list all users in stripe. If true, a service admin cannot list other users.

	**Added In:** 2108190438

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
* `setting_id` - (Required) ID of the resource
* `signing_cert_public_access` - (Optional) (Updatable) Indicates if access on SigningCert is allowed to public or not

	**Added In:** 17.3.4

	**SCIM++ Properties:**
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
* `sub_mapping_attr` - (Optional) (Updatable) **Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none Subject mapping user profile attribute. The input format should be SCIM compliant. This attribute should be of type String and multivalued to false. 
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
* `tenant_custom_claims` - (Optional) (Updatable) Custom claims associated with the specific tenant

	**Added In:** 18.4.2

	**SCIM++ Properties:**
	* idcsCompositeKey: [name]
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `all_scopes` - (Required) (Updatable) Indicates if the custom claim is associated with all scopes

		**Added In:** 18.4.2

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: boolean
		* uniqueness: none
	* `expression` - (Required) (Updatable) Indicates if the custom claim is an expression

		**Added In:** 18.4.2

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: boolean
		* uniqueness: none
	* `mode` - (Required) (Updatable) Indicates under what scenario the custom claim will be return

		**Added In:** 18.4.2

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `name` - (Required) (Updatable) Custom claim name

		**Added In:** 18.4.2

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: server
	* `scopes` - (Optional) (Updatable) Scopes associated with a specific custom claim

		**Added In:** 18.4.2

		**SCIM++ Properties:**
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `token_type` - (Required) (Updatable) Indicates what type of token the custom claim will be embedded

		**Added In:** 18.4.2

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - (Required) (Updatable) Custom claim value

		**Added In:** 18.4.2

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `terms_of_use_url` - (Optional) (Updatable) Terms of Use URL

	**Added In:** 18.2.4

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `timezone` - (Optional) (Updatable) User's timezone

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCanonicalValueSourceFilter: attrName eq "timezones" and attrValues.value eq "$(timezone)"
	* idcsCanonicalValueSourceResourceType: AllowedValue
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `account_always_trust_scope` - Indicates whether all the Apps in this customer tenancy should trust each other. A value of true overrides the 'defaultTrustScope' attribute here in Settings, as well as any App-specific 'trustScope' attribute, to force in effect 'trustScope=Account' for every App in this customer tenancy.

	**Added In:** 18.1.6

	**SCIM++ Properties:**
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
* `allowed_domains` - One or more email domains allowed in a user's email field. If unassigned, any domain is allowed.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `allowed_forgot_password_flow_return_urls` - If specified, indicates the set of Urls which can be returned to after successful forgot password flow

	**Added In:** 19.3.3

	**SCIM++ Properties:**
	* type: string
	* multiValued: true
	* required: false
	* mutability: readWrite
	* returned: default
	* uniqueness: none
	* caseExact: false
* `allowed_notification_redirect_urls` - If specified, indicates the set of allowed notification redirect Urls which can be specified as the value of \"notificationRedirectUrl\" in the POST .../admin/v1/MePasswordResetRequestor request payload, which will then be included in the reset password email notification sent to a user as part of the forgot password / password reset flow.

	**Added In:** 2009041201

	**SCIM++ Properties:**
	* type: string
	* multiValued: true
	* required: false
	* mutability: readWrite
	* returned: default
	* uniqueness: none
	* caseExact: false
* `audit_event_retention_period` - Audit Event retention period. If set, overrides default of 30 days after which Audit Events will be purged

	**Added In:** 19.2.1

	**SCIM++ Properties:**
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
* `certificate_validation` - Certificate Validation Config

	**Added In:** 2010242156

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `crl_check_on_ocsp_failure_enabled` - Use CRL as Fallback.

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `crl_enabled` - CRL is enabled Configuration

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `crl_location` - CRL Location.

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `crl_refresh_interval` - The CRL refresh interval in minutes

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `ocsp_enabled` - OCSP is enabled Configuration

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `ocsp_responder_url` - OCSP Responder URL

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ocsp_settings_responder_url_preferred` - This setting says, OCSP Responder URL present in the issued certificate must be used. Otherwise, OCSP Responder URL from IDP or Settings.

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `ocsp_signing_certificate_alias` - OCSP Signing Certificate Alias

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ocsp_timeout_duration` - The OCSP Timeout duration in minutes

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* idcsMaxValue: 10
		* idcsMinValue: 1
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `ocsp_unknown_response_status_allowed` - OCSP Accept unknown response status from ocsp responder.

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
* `cloud_account_name` - The attribute to store the cloud account name

	**Deprecated Since: 2011192329**

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `cloud_gate_cors_settings` - A complex attribute that specifies the Cloud Gate cross origin resource sharing settings.

	**Added In:** 2011192329

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `cloud_gate_cors_allow_null_origin` - Allow Null Origin (CORS) for this tenant.

		**Added In:** 2011192329

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `cloud_gate_cors_allowed_origins` - Cloud Gate Allowed Cross-Origin Resource Sharing (CORS) Origins for this tenant.

		**Added In:** 2011192329

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `cloud_gate_cors_enabled` - Enable Cloud Gate Cross-Origin Resource Sharing (CORS) for this tenant.

		**Added In:** 2011192329

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `cloud_gate_cors_exposed_headers` - List of Response Headers Cloud Gate is allowed to expose in the CORS Response Header: Access-Control-Expose-Headers.

		**Added In:** 2205182039

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `cloud_gate_cors_max_age` - Maximum number of seconds a CORS Pre-flight Response may be cached by client.

		**Added In:** 2205182039

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
* `cloud_migration_custom_url` - If specified, indicates the custom SIM Migrator Url which can be used while SIM to Oracle Identity Cloud Service CloudAccount Migration.

	**Added In:** 2012271618

	**SCIM++ Properties:**
	* type: string
	* multiValued: false
	* required: false
	* mutability: readWrite
	* returned: default
	* uniqueness: none
	* caseExact: false
* `cloud_migration_url_enabled` - CloudAccountMigration: Enable Custom SIM Migrator Url.

	**Added In:** 2012271618

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `company_names` - Name of the company in different locales

	**SCIM++ Properties:**
	* idcsCompositeKey: [locale]
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* `locale` - Locale

		**SCIM++ Properties:**
		* idcsCanonicalValueSourceFilter: attrName eq "locales" and attrValues.value eq "$(companyNames.locale)"
		* idcsCanonicalValueSourceResourceType: AllowedValue
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
	* `value` - Company name

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
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
* `contact_emails` - Contact emails used to notify tenants. Can be one or more user or group alias emails.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `csr_access` - This value indicates whether Customer Service Representatives can login and have readOnly or readWrite access.  A value of 'none' means CSR cannot login to the services.

	**SCIM++ Properties:**
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: string
* `custom_branding` - Indicates if the branding is default or custom

	**SCIM++ Properties:**
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
* `custom_css_location` - Storage URL location where the sanitized custom css is located

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `custom_html_location` - Storage URL location where the sanitized custom html is located

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `custom_translation` - Custom translations (JSON String)

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `default_company_names` - Default name of the Company in different locales

	**Added In:** 18.2.2

	**SCIM++ Properties:**
	* idcsCompositeKey: [locale]
	* multiValued: true
	* mutability: readOnly
	* required: false
	* returned: default
	* type: complex
	* `locale` - Locale

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* idcsCanonicalValueSourceFilter: attrName eq "locales" and attrValues.value eq "$(companyNames.locale)"
		* idcsCanonicalValueSourceResourceType: AllowedValue
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: string
	* `value` - Company name

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: string
* `default_images` - References to various images

	**Added In:** 18.2.2

	**SCIM++ Properties:**
	* idcsCompositeKey: [type]
	* multiValued: true
	* mutability: readOnly
	* required: false
	* returned: default
	* type: complex
	* `display` - A human-readable name, primarily used for display purposes

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
	* `type` - Indicates the image type

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: string
	* `value` - Image URI

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: reference
* `default_login_texts` - Default Login text in different locales

	**Added In:** 18.2.2

	**SCIM++ Properties:**
	* idcsCompositeKey: [locale]
	* multiValued: true
	* mutability: readOnly
	* required: false
	* returned: default
	* type: complex
	* `locale` - Locale

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* idcsCanonicalValueSourceFilter: attrName eq "locales" and attrValues.value eq "$(loginTexts.locale)"
		* idcsCanonicalValueSourceResourceType: AllowedValue
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: string
	* `value` - Login text

		**Added In:** 18.2.2

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: string
		* idcsSanitize: true
* `default_trust_scope` - **Deprecated Since: 18.3.6**

	**SCIM++ Properties:**
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string Indicates the default trust scope for all apps 
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
* `diagnostic_level` - The level of diagnostic logging that is currently in effect. A level of 0 (zero) indicates that diagnostic logging is disabled. A level of 1 (one) indicates that diagnostic logging is enabled.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `diagnostic_record_for_search_identifies_returned_resources` - Controls whether DiagnosticRecords for external search-operations (against SCIM resource-types in the Admin service) identify returned resources.  If true, indicates that for each successful external search-operation at least one DiagnosticRecord will include at least one identifier for each matching resource that is returned in that search-response.  If false, no DiagnosticRecord should be expected to identify returned resources for a search-operation.  The default value is false.

	**Added In:** 2011192329

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `diagnostic_tracing_upto` - The end time up to which diagnostic recording is switched on

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: dateTime
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
* `enable_terms_of_use` - Indicates if Terms of Use is enabled in UI

	**Added In:** 18.2.4

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `external_id` - An identifier for the Resource as defined by the Service Consumer. The externalId may simplify identification of the Resource between Service Consumer and Service Provider by allowing the Consumer to refer to the Resource with its own identifier, obviating the need to store a local mapping between the local identifier of the Resource and the identifier used by the Service Provider. Each Resource MAY include a non-empty externalId value. The value of the externalId attribute is always issued by the Service Consumer and can never be specified by the Service Provider. The Service Provider MUST always interpret the externalId as scoped to the Service Consumer's tenant.

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `iam_upst_session_expiry` - Maximum duration for IAM User Principal Session Token expiry

	**Added In:** 2307071836

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
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
* `images` - References to various images

	**SCIM++ Properties:**
	* idcsCompositeKey: [type]
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* `display` - A human-readable name, primarily used for display purposes

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
	* `type` - Indicates the image type

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
	* `value` - Image URI

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: reference
* `is_hosted_page` - Indicates if 'hosted' option was selected

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `issuer` - Tenant issuer.

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `locale` - Default location for purposes of localizing items such as currency, date and time format, numerical representations, and so on.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCanonicalValueSourceFilter: attrName eq "locales" and attrValues.value eq "$(locale)"
	* idcsCanonicalValueSourceResourceType: AllowedValue
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `login_texts` - Login text in different locales

	**SCIM++ Properties:**
	* idcsCompositeKey: [locale]
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* `locale` - Locale

		**SCIM++ Properties:**
		* idcsCanonicalValueSourceFilter: attrName eq "locales" and attrValues.value eq "$(loginTexts.locale)"
		* idcsCanonicalValueSourceResourceType: AllowedValue
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
	* `value` - Login text

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* idcsSanitize: true
		* type: string
* `max_no_of_app_cmva_to_return` - Limit the maximum return of CMVA for an App

	**Added In:** 2111112015

	**SCIM++ Properties:**
	* idcsMinValue: 0
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
	* uniqueness: none
* `max_no_of_app_role_members_to_return` - Limit the maximum return of members for an AppRole

	**Added In:** 2111112015

	**SCIM++ Properties:**
	* idcsMinValue: 0
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
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
* `migration_status` - Database Migration Status

	**Added In:** 19.2.1

	**SCIM++ Properties:**
	* caseExact: true
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: string
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
* `on_premises_provisioning` - On-Premises provisioning feature toggle.

	**Added In:** 19.2.1

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `preferred_language` - Preferred written or spoken language used for localized user interfaces

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCanonicalValueSourceFilter: attrName eq "languages" and attrValues.value eq "$(preferredLanguage)"
	* idcsCanonicalValueSourceResourceType: AllowedValue
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `prev_issuer` - Previous Tenant issuer. This is an Oracle Identity Cloud Service internal attribute which is not meant to be directly modified by ID Admin. Even if the request body (Settings) contains this attribute, the actual value will be set according to the Oracle Identity Cloud Service internal logic rather than solely based on the value provided in the request payload.

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: request
	* type: string
	* uniqueness: none
* `privacy_policy_url` - Privacy Policy URL

	**Added In:** 18.2.4

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `purge_configs` - Purge Configs for different Resource Types

	**Deprecated Since: 19.1.6**

	**SCIM++ Properties:**
	* idcsCompositeKey: [resourceName]
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* `resource_name` - Resource Name

		**Deprecated Since: 19.1.6**

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
	* `retention_period` - Retention Period

		**Deprecated Since: 19.1.6**

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: integer
* `re_auth_factor` - If reAuthWhenChangingMyAuthenticationFactors is true (default), this attribute specifies which re-authentication factor to use. Allowed value is \"password\".

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
* `re_auth_when_changing_my_authentication_factors` - Specifies whether re-authentication is required or not when a user changes one of their security factors such as password or email. Default is true to ensure more secure behavior.

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
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
* `service_admin_cannot_list_other_users` - By default, a service admin can list all users in stripe. If true, a service admin cannot list other users.

	**Added In:** 2108190438

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
* `signing_cert_public_access` - Indicates if access on SigningCert is allowed to public or not

	**Added In:** 17.3.4

	**SCIM++ Properties:**
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
* `sub_mapping_attr` - **Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none Subject mapping user profile attribute. The input format should be SCIM compliant. This attribute should be of type String and multivalued to false. 
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
* `tenant_custom_claims` - Custom claims associated with the specific tenant

	**Added In:** 18.4.2

	**SCIM++ Properties:**
	* idcsCompositeKey: [name]
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `all_scopes` - Indicates if the custom claim is associated with all scopes

		**Added In:** 18.4.2

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: boolean
		* uniqueness: none
	* `expression` - Indicates if the custom claim is an expression

		**Added In:** 18.4.2

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: boolean
		* uniqueness: none
	* `mode` - Indicates under what scenario the custom claim will be return

		**Added In:** 18.4.2

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `name` - Custom claim name

		**Added In:** 18.4.2

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: server
	* `scopes` - Scopes associated with a specific custom claim

		**Added In:** 18.4.2

		**SCIM++ Properties:**
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `token_type` - Indicates what type of token the custom claim will be embedded

		**Added In:** 18.4.2

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - Custom claim value

		**Added In:** 18.4.2

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `terms_of_use_url` - Terms of Use URL

	**Added In:** 18.2.4

	**SCIM++ Properties:**
	* caseExact: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `timezone` - User's timezone

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCanonicalValueSourceFilter: attrName eq "timezones" and attrValues.value eq "$(timezone)"
	* idcsCanonicalValueSourceResourceType: AllowedValue
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Setting
	* `update` - (Defaults to 20 minutes), when updating the Setting
	* `delete` - (Defaults to 20 minutes), when destroying the Setting


## Import

Settings can be imported using the `id`, e.g.

```
$ terraform import oci_identity_domains_setting.test_setting "idcsEndpoint/{idcsEndpoint}/settings/{settingId}" 
```

