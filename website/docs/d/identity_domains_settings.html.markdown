---
subcategory: "Identity Domains"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_domains_settings"
sidebar_current: "docs-oci-datasource-identity_domains-settings"
description: |-
  Provides the list of Settings in Oracle Cloud Infrastructure Identity Domains service
---

# Data Source: oci_identity_domains_settings
This data source provides the list of Settings in Oracle Cloud Infrastructure Identity Domains service.

Search Settings

## Example Usage

```hcl
data "oci_identity_domains_settings" "test_settings" {
	#Required
	idcs_endpoint = data.oci_identity_domain.test_domain.url

	#Optional
	attribute_sets = ["all"]
	attributes = ""
	authorization = var.setting_authorization
	resource_type_schema_version = var.setting_resource_type_schema_version
}
```

## Argument Reference

The following arguments are supported:

* `attribute_sets` - (Optional) A multi-valued list of strings indicating the return type of attribute definition. The specified set of attributes can be fetched by the return type of the attribute. One or more values can be given together to fetch more than one group of attributes. If 'attributes' query parameter is also available, union of the two is fetched. Valid values - all, always, never, request, default. Values are case-insensitive.
* `attributes` - (Optional) A comma-delimited string that specifies the names of resource attributes that should be returned in the response. By default, a response that contains resource attributes contains only attributes that are defined in the schema for that resource type as returned=always or returned=default. An attribute that is defined as returned=request is returned in a response only if the request specifies its name in the value of this query parameter. If a request specifies this query parameter, the response contains the attributes that this query parameter specifies, as well as any attribute that is defined as returned=always.
* `authorization` - (Optional) The Authorization field value consists of credentials containing the authentication information of the user agent for the realm of the resource being requested.
* `idcs_endpoint` - (Required) The basic endpoint for the identity domain
* `resource_type_schema_version` - (Optional) An endpoint-specific schema version number to use in the Request. Allowed version values are Earliest Version or Latest Version as specified in each REST API endpoint description, or any sequential number inbetween. All schema attributes/body parameters are a part of version 1. After version 1, any attributes added or deprecated will be tagged with the version that they were added to or deprecated in. If no version is provided, the latest schema version is returned.


## Attributes Reference

The following attributes are exported:

* `settings` - The list of settings.

### Setting Reference

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

