---
subcategory: "Identity Domains"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_domains_identity_provider"
sidebar_current: "docs-oci-resource-identity_domains-identity_provider"
description: |-
  Provides the Identity Provider resource in Oracle Cloud Infrastructure Identity Domains service
---

# oci_identity_domains_identity_provider
This resource provides the Identity Provider resource in Oracle Cloud Infrastructure Identity Domains service.

Create an Identity Provider

## Example Usage

```hcl
resource "oci_identity_domains_identity_provider" "test_identity_provider" {
	#Required
	enabled = false
	idcs_endpoint = data.oci_identity_domain.test_domain.url
	partner_name = var.identity_provider_partner_name
	schemas = ["urn:ietf:params:scim:schemas:oracle:idcs:IdentityProvider"]

	#Optional
	assertion_attribute = var.identity_provider_assertion_attribute
	attribute_sets = []
	attributes = ""
	authn_request_binding = var.identity_provider_authn_request_binding
	authorization = var.identity_provider_authorization
	correlation_policy {
		#Required
		type = var.identity_provider_correlation_policy_type
		value = var.identity_provider_correlation_policy_value

		#Optional
		display = var.identity_provider_correlation_policy_display
	}
	description = var.identity_provider_description
	encryption_certificate = var.identity_provider_encryption_certificate
	external_id = "externalId"
	icon_url = var.identity_provider_icon_url
	id = var.identity_provider_id
	idp_sso_url = var.identity_provider_idp_sso_url
	include_signing_cert_in_signature = var.identity_provider_include_signing_cert_in_signature
	jit_user_prov_assigned_groups {
		#Required
		value = var.identity_provider_jit_user_prov_assigned_groups_value
	}
	jit_user_prov_attribute_update_enabled = var.identity_provider_jit_user_prov_attribute_update_enabled
	jit_user_prov_attributes {
		#Required
		value = var.identity_provider_jit_user_prov_attributes_value
	}
	jit_user_prov_create_user_enabled = var.identity_provider_jit_user_prov_create_user_enabled
	jit_user_prov_enabled = var.identity_provider_jit_user_prov_enabled
	jit_user_prov_group_assertion_attribute_enabled = var.identity_provider_jit_user_prov_group_assertion_attribute_enabled
	jit_user_prov_group_assignment_method = var.identity_provider_jit_user_prov_group_assignment_method
	jit_user_prov_group_mapping_mode = var.identity_provider_jit_user_prov_group_mapping_mode
	jit_user_prov_group_mappings {
		#Required
		idp_group = var.identity_provider_jit_user_prov_group_mappings_idp_group
		value = var.identity_provider_jit_user_prov_group_mappings_value
	}
	jit_user_prov_group_saml_attribute_name = var.identity_provider_jit_user_prov_group_saml_attribute_name
	jit_user_prov_group_static_list_enabled = var.identity_provider_jit_user_prov_group_static_list_enabled
	jit_user_prov_ignore_error_on_absent_groups = var.identity_provider_jit_user_prov_ignore_error_on_absent_groups
	logout_binding = var.identity_provider_logout_binding
	logout_enabled = var.identity_provider_logout_enabled
	logout_request_url = var.identity_provider_logout_request_url
	logout_response_url = var.identity_provider_logout_response_url
	metadata = var.identity_provider_metadata
	name_id_format = var.identity_provider_name_id_format
	ocid = var.identity_provider_ocid
	partner_provider_id = var.identity_provider_partner_provider_id
	requested_authentication_context = var.identity_provider_requested_authentication_context
	require_force_authn = var.identity_provider_require_force_authn
	requires_encrypted_assertion = var.identity_provider_requires_encrypted_assertion
	resource_type_schema_version = var.identity_provider_resource_type_schema_version
	saml_ho_krequired = var.identity_provider_saml_ho_krequired
	service_instance_identifier = var.identity_provider_service_instance_identifier
	shown_on_login_page = var.identity_provider_shown_on_login_page
	signature_hash_algorithm = var.identity_provider_signature_hash_algorithm
	signing_certificate = var.identity_provider_signing_certificate
	succinct_id = "succinctId"
	tags {
		#Required
		key = var.identity_provider_tags_key
		value = var.identity_provider_tags_value
	}
	type = var.identity_provider_type
	urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider {
		#Required
		account_linking_enabled = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_account_linking_enabled
		consumer_key = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_consumer_key
		consumer_secret = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_consumer_secret
		registration_enabled = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_registration_enabled
		service_provider_name = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider_service_provider_name

		#Optional
		access_token_url = var.identity_provider_urn_ietf_params_scim_schemas_oracle_idcs_extension_social_identity_provider_access_token_url
		admin_scope = var.identity_provider_urn_ietf_params_scim_schemas_oracle_idcs_extension_social_identity_provider_admin_scope
		authz_url = var.identity_provider_urn_ietf_params_scim_schemas_oracle_idcs_extension_social_identity_provider_authz_url
		auto_redirect_enabled = var.identity_provider_urn_ietf_params_scim_schemas_oracle_idcs_extension_social_identity_provider_auto_redirect_enabled
		client_credential_in_payload = var.identity_provider_urn_ietf_params_scim_schemas_oracle_idcs_extension_social_identity_provider_client_credential_in_payload
		clock_skew_in_seconds = var.identity_provider_urn_ietf_params_scim_schemas_oracle_idcs_extension_social_identity_provider_clock_skew_in_seconds
		discovery_url = var.identity_provider_urn_ietf_params_scim_schemas_oracle_idcs_extension_social_identity_provider_discovery_url
		id_attribute = var.identity_provider_urn_ietf_params_scim_schemas_oracle_idcs_extension_social_identity_provider_id_attribute
		jit_prov_assigned_groups {
			#Required
			value = var.identity_provider_urn_ietf_params_scim_schemas_oracle_idcs_extension_social_identity_provider_jit_prov_assigned_groups_value

			#Optional
			display = var.identity_provider_urn_ietf_params_scim_schemas_oracle_idcs_extension_social_identity_provider_jit_prov_assigned_groups_display
		}
		jit_prov_group_static_list_enabled = var.identity_provider_urn_ietf_params_scim_schemas_oracle_idcs_extension_social_identity_provider_jit_prov_group_static_list_enabled
		profile_url = var.identity_provider_urn_ietf_params_scim_schemas_oracle_idcs_extension_social_identity_provider_profile_url
		redirect_url = var.identity_provider_urn_ietf_params_scim_schemas_oracle_idcs_extension_social_identity_provider_redirect_url
		scope = var.identity_provider_urn_ietf_params_scim_schemas_oracle_idcs_extension_social_identity_provider_scope
		social_jit_provisioning_enabled = var.identity_provider_urn_ietf_params_scim_schemas_oracle_idcs_extension_social_identity_provider_social_jit_provisioning_enabled
		status = var.identity_provider_urn_ietf_params_scim_schemas_oracle_idcs_extension_social_identity_provider_status
	}
	urnietfparamsscimschemasoracleidcsextensionx509identity_provider {
		#Required
		cert_match_attribute = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_cert_match_attribute
		signing_certificate_chain = ["signingCertificateChain"]
		user_match_attribute = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_user_match_attribute

		#Optional
		crl_check_on_ocsp_failure_enabled = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_crl_check_on_ocsp_failure_enabled
		crl_enabled = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_crl_enabled
		crl_location = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_crl_location
		crl_reload_duration = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_crl_reload_duration
		eku_validation_enabled = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_eku_validation_enabled
		eku_values = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_eku_values
		ocsp_allow_unknown_response_status = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_ocsp_allow_unknown_response_status
		ocsp_enable_signed_response = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_ocsp_enable_signed_response
		ocsp_enabled = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_ocsp_enabled
		ocsp_responder_url = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_ocsp_responder_url
		ocsp_revalidate_time = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_ocsp_revalidate_time
		ocsp_server_name = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_ocsp_server_name
		ocsp_trust_cert_chain = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_ocsp_trust_cert_chain
		other_cert_match_attribute = var.identity_provider_urnietfparamsscimschemasoracleidcsextensionx509identity_provider_other_cert_match_attribute
	}
	user_mapping_method = var.identity_provider_user_mapping_method
	user_mapping_store_attribute = var.identity_provider_user_mapping_store_attribute
}
```

## Argument Reference

The following arguments are supported:

* `assertion_attribute` - (Optional) (Updatable) Assertion attribute name.

	**Deprecated Since: 20.1.3**

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
	* idcsValuePersistedInOtherAttribute: true
* `attribute_sets` - (Optional) (Updatable) A multi-valued list of strings indicating the return type of attribute definition. The specified set of attributes can be fetched by the return type of the attribute. One or more values can be given together to fetch more than one group of attributes. If 'attributes' query parameter is also available, union of the two is fetched. Valid values - all, always, never, request, default. Values are case-insensitive.
* `attributes` - (Optional) (Updatable) A comma-delimited string that specifies the names of resource attributes that should be returned in the response. By default, a response that contains resource attributes contains only attributes that are defined in the schema for that resource type as returned=always or returned=default. An attribute that is defined as returned=request is returned in a response only if the request specifies its name in the value of this query parameter. If a request specifies this query parameter, the response contains the attributes that this query parameter specifies, as well as any attribute that is defined as returned=always.
* `authn_request_binding` - (Optional) (Updatable) HTTP binding to use for authentication requests.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `authorization` - (Optional) (Updatable) The Authorization field value consists of credentials containing the authentication information of the user agent for the realm of the resource being requested.
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
* `correlation_policy` - (Optional) (Updatable) Correlation policy

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: false
	* multiValued: false
	* mutability: immutable
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `display` - (Optional) (Updatable) Policy display name

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ref` - (Optional) (Updatable) Policy URI

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `type` - (Required) (Updatable) A label that indicates the type that this references.

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* idcsDefaultValue: Policy
		* idcsSearchable: false
		* multiValued: false
		* mutability: immutable
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - (Required) (Updatable) Policy identifier

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
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
* `description` - (Optional) (Updatable) Description

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
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
* `enabled` - (Required) (Updatable) Set to true to indicate Partner enabled.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: boolean
	* uniqueness: none
* `encryption_certificate` - (Optional) (Updatable) Encryption certificate

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
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
* `icon_url` - (Optional) (Updatable) Identity Provider Icon URL.

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
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
* `idp_sso_url` - (Optional) (Updatable) Identity Provider SSO URL

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `include_signing_cert_in_signature` - (Optional) (Updatable) Set to true to include the signing certificate in the signature.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `jit_user_prov_assigned_groups` - (Optional) (Updatable) Refers to every group of which a JIT-provisioned User should be a member.  Just-in-Time user-provisioning applies this static list when jitUserProvGroupStaticListEnabled:true.

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* idcsCompositeKey: [value]
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `display` - (Optional) (Updatable) A human readable name, primarily used for display purposes. READ-ONLY.

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ref` - (Optional) (Updatable) Group URI

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - (Required) (Updatable) Group identifier

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `jit_user_prov_attribute_update_enabled` - (Optional) (Updatable) Set to true to indicate JIT User Creation is enabled

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `jit_user_prov_attributes` - (Optional) (Updatable) Assertion To User Mapping

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCompositeKey: [value]
	* idcsSearchable: false
	* mutability: immutable
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `ref` - (Optional) (Updatable) Mapped Attribute URI

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* idcsSearchable: false
		* mutability: immutable
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - (Required) (Updatable) Mapped Attribute identifier

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* mutability: immutable
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `jit_user_prov_create_user_enabled` - (Optional) (Updatable) Set to true to indicate JIT User Creation is enabled

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `jit_user_prov_enabled` - (Optional) (Updatable) Set to true to indicate JIT User Provisioning is enabled

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `jit_user_prov_group_assertion_attribute_enabled` - (Optional) (Updatable) Set to true to indicate JIT User Provisioning Groups should be assigned based on assertion attribute

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `jit_user_prov_group_assignment_method` - (Optional) (Updatable) The default value is 'Overwrite', which tells Just-In-Time user-provisioning to replace any current group-assignments for a User with those assigned by assertions and/or those assigned statically. Specify 'Merge' if you want Just-In-Time user-provisioning to combine its group-assignments with those the user already has.

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `jit_user_prov_group_mapping_mode` - (Optional) (Updatable) Property to indicate the mode of group mapping

	**Added In:** 2205120021

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `jit_user_prov_group_mappings` - (Optional) (Updatable) The list of mappings between the Identity Domain Group and the IDP group.

	**Added In:** 2205120021

	**SCIM++ Properties:**
	* idcsCompositeKey: [idpGroup]
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `idp_group` - (Required) (Updatable) IDP Group Name

		**Added In:** 2205120021

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* type: string
	* `ref` - (Optional) (Updatable) Group URI

		**Added In:** 2205120021

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: reference
	* `value` - (Required) (Updatable) Domain Group

		**Added In:** 2205120021

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: true
		* idcsSearchable: true
		* type: string
* `jit_user_prov_group_saml_attribute_name` - (Optional) (Updatable) Name of the assertion attribute containing the users groups

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `jit_user_prov_group_static_list_enabled` - (Optional) (Updatable) Set to true to indicate JIT User Provisioning Groups should be assigned from a static list

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `jit_user_prov_ignore_error_on_absent_groups` - (Optional) (Updatable) Set to true to indicate ignoring absence of group while provisioning

	**Added In:** 2111112015

	**SCIM++ Properties:**
	* caseExact: false
	* idcsAddedSinceVersion: 30
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `last_notification_sent_time` - (Optional) (Updatable) Records the notification timestamp for the IdP whose signing certificate is about to expire

	**Added In:** 2302092332

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: dateTime
	* uniqueness: none
* `logout_binding` - (Optional) (Updatable) HTTP binding to use for logout.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `logout_enabled` - (Optional) (Updatable) Set to true to enable logout.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `logout_request_url` - (Optional) (Updatable) Logout request URL

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `logout_response_url` - (Optional) (Updatable) Logout response URL

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
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
* `metadata` - (Optional) (Updatable) Metadata

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `name_id_format` - (Optional) (Updatable) Default authentication request name ID format.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
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
* `partner_name` - (Required) (Updatable) Unique name of the trusted Identity Provider.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: always
	* type: string
	* uniqueness: server
* `partner_provider_id` - (Optional) (Updatable) Provider ID

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: server
* `requested_authentication_context` - (Optional) (Updatable) SAML SP authentication type.

	**Added In:** 2102181953

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `require_force_authn` - (Optional) (Updatable) This SP requires requests SAML IdP to enforce re-authentication.

	**Added In:** 2102181953

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `requires_encrypted_assertion` - (Optional) (Updatable) SAML SP must accept encrypted assertion only.

	**Added In:** 2102181953

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `resource_type_schema_version` - (Optional) (Updatable) An endpoint-specific schema version number to use in the Request. Allowed version values are Earliest Version or Latest Version as specified in each REST API endpoint description, or any sequential number inbetween. All schema attributes/body parameters are a part of version 1. After version 1, any attributes added or deprecated will be tagged with the version that they were added to or deprecated in. If no version is provided, the latest schema version is returned.
* `saml_ho_krequired` - (Optional) (Updatable) SAML SP HoK Enabled.

	**Added In:** 2102181953

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
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
* `service_instance_identifier` - (Optional) (Updatable) The serviceInstanceIdentifier of the App that hosts this IdP. This value will match the opcServiceInstanceGUID of any service-instance that the IdP represents.

	**Added In:** 18.2.6

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: never
	* type: string
	* uniqueness: server
* `shown_on_login_page` - (Optional) (Updatable) Set to true to indicate whether to show IdP in login page or not.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `signature_hash_algorithm` - (Optional) (Updatable) Signature hash algorithm.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `signing_certificate` - (Optional) (Updatable) Signing certificate

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `succinct_id` - (Optional) (Updatable) Succinct ID

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: server
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
* `tenant_provider_id` - (Optional) (Updatable) The alternate Provider ID to be used as the Oracle Identity Cloud Service providerID (instead of the one in SamlSettings) when interacting with this IdP.

	**Added In:** 19.2.1

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: false
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `type` - (Optional) (Updatable) Identity Provider Type

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: immutable
	* required: false
	* returned: always
	* type: string
	* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider` - (Optional) (Updatable) Social Identity Provider Extension Schema
	* `access_token_url` - (Optional) (Updatable) Social IDP Access token URL

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `account_linking_enabled` - (Required) (Updatable) Whether account linking is enabled

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: boolean
		* uniqueness: none
	* `admin_scope` - (Optional) (Updatable) Admin scope to request

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `authz_url` - (Optional) (Updatable) Social IDP Authorization URL

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `auto_redirect_enabled` - (Optional) (Updatable) Whether social auto redirect is enabled. The IDP policy should be configured with only one Social IDP, and without username/password selected.

		**Added In:** 2310202314

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `client_credential_in_payload` - (Optional) (Updatable) Whether the client credential is contained in payload

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `clock_skew_in_seconds` - (Optional) (Updatable) Social IDP allowed clock skew time

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `consumer_key` - (Required) (Updatable) Social IDP Client Application Client ID

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `consumer_secret` - (Required) (Updatable) Social IDP Client Application Client Secret

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* idcsSensitive: encrypt
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `discovery_url` - (Optional) (Updatable) Discovery URL

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `id_attribute` - (Optional) (Updatable) Id attribute used for account linking

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: immutable
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `jit_prov_assigned_groups` - (Optional) (Updatable) Lists the groups each social JIT-provisioned user is a member. Just-in-Time user-provisioning applies this static list when jitProvGroupStaticListEnabled:true.

		**Added In:** 2310202314

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* idcsSearchable: false
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `_ref` - (Optional) (Updatable) Group URI

			**Added In:** 2310202314

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `display` - (Optional) (Updatable) A human readable name, primarily used for display purposes. READ-ONLY.

			**Added In:** 2310202314

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: request
			* type: string
			* uniqueness: none
		* `value` - (Required) (Updatable) Group identifier

			**Added In:** 2310202314

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
	* `jit_prov_group_static_list_enabled` - (Optional) (Updatable) Set to true to indicate Social JIT User Provisioning Groups should be assigned from a static list

		**Added In:** 2310202314

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `profile_url` - (Optional) (Updatable) Social IDP User profile URL

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `redirect_url` - (Optional) (Updatable) redirect URL for social idp

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `registration_enabled` - (Required) (Updatable) Whether registration is enabled

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: boolean
		* uniqueness: none
	* `scope` - (Optional) (Updatable) Scope to request

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `service_provider_name` - (Required) (Updatable) Service Provider Name

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `social_jit_provisioning_enabled` - (Optional) (Updatable) Whether Social JIT Provisioning is enabled

		**Added In:** 2307282043

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `status` - (Optional) (Updatable) Status

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionx509identity_provider` - (Optional) (Updatable) X509 Identity Provider Extension Schema
	* `cert_match_attribute` - (Required) (Updatable) X509 Certificate Matching Attribute

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `crl_check_on_ocsp_failure_enabled` - (Optional) (Updatable) Fallback on CRL Validation if OCSP fails.

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `crl_enabled` - (Optional) (Updatable) Set to true to enable CRL Validation

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `crl_location` - (Optional) (Updatable) CRL Location URL

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `crl_reload_duration` - (Optional) (Updatable) Fetch the CRL contents every X minutes

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `eku_validation_enabled` - (Optional) (Updatable) Set to true to enable EKU Validation

		**Added In:** 2304270343

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `eku_values` - (Optional) (Updatable) List of EKU which needs to be validated

		**Added In:** 2304270343

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ocsp_allow_unknown_response_status` - (Optional) (Updatable) Allow access if OCSP response is UNKNOWN or OCSP Responder does not respond within the timeout duration

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `ocsp_enable_signed_response` - (Optional) (Updatable) Describes if the OCSP response is signed

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `ocsp_enabled` - (Optional) (Updatable) Set to true to enable OCSP Validation

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `ocsp_responder_url` - (Optional) (Updatable) This property specifies OCSP Responder URL.

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ocsp_revalidate_time` - (Optional) (Updatable) Revalidate OCSP status for user after X hours

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* idcsMaxValue: 24
		* idcsMinValue: 0
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `ocsp_server_name` - (Optional) (Updatable) This property specifies the OCSP Server alias name

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ocsp_trust_cert_chain` - (Optional) (Updatable) OCSP Trusted Certificate Chain

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `other_cert_match_attribute` - (Optional) (Updatable) Check for specific conditions of other certificate attributes

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `signing_certificate_chain` - (Required) (Updatable) Certificate alias list to create a chain for the incoming client certificate

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: true
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `user_match_attribute` - (Required) (Updatable) This property specifies the userstore attribute value that must match the incoming certificate attribute.

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `user_mapping_method` - (Optional) (Updatable) User mapping method.

	**Deprecated Since: 20.1.3**

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
	* idcsValuePersistedInOtherAttribute: true
* `user_mapping_store_attribute` - (Optional) (Updatable) This property specifies the userstore attribute value that must match the incoming assertion attribute value or the incoming nameid attribute value in order to identify the user during SSO.<br>You can construct the userMappingStoreAttribute value by specifying attributes from the Oracle Identity Cloud Service Core Users schema. For examples of how to construct the userMappingStoreAttribute value, see the <b>Example of a Request Body</b> section of the Examples tab for the <a href='./op-admin-v1-identityproviders-post.html'>POST</a> and <a href='./op-admin-v1-identityproviders-id-put.html'>PUT</a> methods of the /IdentityProviders endpoint.

	**Deprecated Since: 20.1.3**

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
	* idcsValuePersistedInOtherAttribute: true


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `assertion_attribute` - Assertion attribute name.

	**Deprecated Since: 20.1.3**

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
	* idcsValuePersistedInOtherAttribute: true
* `authn_request_binding` - HTTP binding to use for authentication requests.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
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
* `correlation_policy` - Correlation policy

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: false
	* multiValued: false
	* mutability: immutable
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `display` - Policy display name

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ref` - Policy URI

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `type` - A label that indicates the type that this references.

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* idcsDefaultValue: Policy
		* idcsSearchable: false
		* multiValued: false
		* mutability: immutable
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - Policy identifier

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
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
* `description` - Description

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
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
* `enabled` - Set to true to indicate Partner enabled.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: boolean
	* uniqueness: none
* `encryption_certificate` - Encryption certificate

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
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
* `icon_url` - Identity Provider Icon URL.

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
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
* `idp_sso_url` - Identity Provider SSO URL

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `include_signing_cert_in_signature` - Set to true to include the signing certificate in the signature.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `jit_user_prov_assigned_groups` - Refers to every group of which a JIT-provisioned User should be a member.  Just-in-Time user-provisioning applies this static list when jitUserProvGroupStaticListEnabled:true.

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* idcsCompositeKey: [value]
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `display` - A human readable name, primarily used for display purposes. READ-ONLY.

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ref` - Group URI

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - Group identifier

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `jit_user_prov_attribute_update_enabled` - Set to true to indicate JIT User Creation is enabled

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `jit_user_prov_attributes` - Assertion To User Mapping

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCompositeKey: [value]
	* idcsSearchable: false
	* mutability: immutable
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `ref` - Mapped Attribute URI

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* idcsSearchable: false
		* mutability: immutable
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - Mapped Attribute identifier

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* mutability: immutable
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `jit_user_prov_create_user_enabled` - Set to true to indicate JIT User Creation is enabled

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `jit_user_prov_enabled` - Set to true to indicate JIT User Provisioning is enabled

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `jit_user_prov_group_assertion_attribute_enabled` - Set to true to indicate JIT User Provisioning Groups should be assigned based on assertion attribute

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `jit_user_prov_group_assignment_method` - The default value is 'Overwrite', which tells Just-In-Time user-provisioning to replace any current group-assignments for a User with those assigned by assertions and/or those assigned statically. Specify 'Merge' if you want Just-In-Time user-provisioning to combine its group-assignments with those the user already has.

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `jit_user_prov_group_mapping_mode` - Property to indicate the mode of group mapping

	**Added In:** 2205120021

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `jit_user_prov_group_mappings` - The list of mappings between the Identity Domain Group and the IDP group.

	**Added In:** 2205120021

	**SCIM++ Properties:**
	* idcsCompositeKey: [idpGroup]
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `idp_group` - IDP Group Name

		**Added In:** 2205120021

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* type: string
	* `ref` - Group URI

		**Added In:** 2205120021

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: reference
	* `value` - Domain Group

		**Added In:** 2205120021

		**SCIM++ Properties:**
		* multiValued: false
		* mutability: readWrite
		* required: true
		* idcsSearchable: true
		* type: string
* `jit_user_prov_group_saml_attribute_name` - Name of the assertion attribute containing the users groups

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `jit_user_prov_group_static_list_enabled` - Set to true to indicate JIT User Provisioning Groups should be assigned from a static list

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `jit_user_prov_ignore_error_on_absent_groups` - Set to true to indicate ignoring absence of group while provisioning

	**Added In:** 2111112015

	**SCIM++ Properties:**
	* caseExact: false
	* idcsAddedSinceVersion: 30
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `last_notification_sent_time` - Records the notification timestamp for the IdP whose signing certificate is about to expire

	**Added In:** 2302092332

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: dateTime
	* uniqueness: none
* `logout_binding` - HTTP binding to use for logout.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `logout_enabled` - Set to true to enable logout.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `logout_request_url` - Logout request URL

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `logout_response_url` - Logout response URL

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
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
* `metadata` - Metadata

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `name_id_format` - Default authentication request name ID format.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
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
* `partner_name` - Unique name of the trusted Identity Provider.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: always
	* type: string
	* uniqueness: server
* `partner_provider_id` - Provider ID

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: server
* `requested_authentication_context` - SAML SP authentication type.

	**Added In:** 2102181953

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `require_force_authn` - This SP requires requests SAML IdP to enforce re-authentication.

	**Added In:** 2102181953

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `requires_encrypted_assertion` - SAML SP must accept encrypted assertion only.

	**Added In:** 2102181953

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `saml_ho_krequired` - SAML SP HoK Enabled.

	**Added In:** 2102181953

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: false
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
* `service_instance_identifier` - The serviceInstanceIdentifier of the App that hosts this IdP. This value will match the opcServiceInstanceGUID of any service-instance that the IdP represents.

	**Added In:** 18.2.6

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: never
	* type: string
	* uniqueness: server
* `shown_on_login_page` - Set to true to indicate whether to show IdP in login page or not.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: boolean
	* uniqueness: none
* `signature_hash_algorithm` - Signature hash algorithm.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `signing_certificate` - Signing certificate

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `succinct_id` - Succinct ID

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: server
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
* `tenant_provider_id` - The alternate Provider ID to be used as the Oracle Identity Cloud Service providerID (instead of the one in SamlSettings) when interacting with this IdP.

	**Added In:** 19.2.1

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: false
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `type` - Identity Provider Type

	**Added In:** 20.1.3

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: immutable
	* required: false
	* returned: always
	* type: string
	* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionsocial_identity_provider` - Social Identity Provider Extension Schema
	* `access_token_url` - Social IDP Access token URL

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `account_linking_enabled` - Whether account linking is enabled

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: boolean
		* uniqueness: none
	* `admin_scope` - Admin scope to request

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `authz_url` - Social IDP Authorization URL

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `auto_redirect_enabled` - Whether social auto redirect is enabled. The IDP policy should be configured with only one Social IDP, and without username/password selected.

		**Added In:** 2310202314

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `client_credential_in_payload` - Whether the client credential is contained in payload

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `clock_skew_in_seconds` - Social IDP allowed clock skew time

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `consumer_key` - Social IDP Client Application Client ID

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `consumer_secret` - Social IDP Client Application Client Secret

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* idcsSensitive: encrypt
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `discovery_url` - Discovery URL

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `id_attribute` - Id attribute used for account linking

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: immutable
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `jit_prov_assigned_groups` - Lists the groups each social JIT-provisioned user is a member. Just-in-Time user-provisioning applies this static list when jitProvGroupStaticListEnabled:true.

		**Added In:** 2310202314

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* idcsSearchable: false
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `_ref` - Group URI

			**Added In:** 2310202314

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `display` - A human readable name, primarily used for display purposes. READ-ONLY.

			**Added In:** 2310202314

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: request
			* type: string
			* uniqueness: none
		* `value` - Group identifier

			**Added In:** 2310202314

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
	* `jit_prov_group_static_list_enabled` - Set to true to indicate Social JIT User Provisioning Groups should be assigned from a static list

		**Added In:** 2310202314

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `profile_url` - Social IDP User profile URL

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `redirect_url` - redirect URL for social idp

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `registration_enabled` - Whether registration is enabled

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: boolean
		* uniqueness: none
	* `scope` - Scope to request

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `service_provider_name` - Service Provider Name

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `social_jit_provisioning_enabled` - Whether Social JIT Provisioning is enabled

		**Added In:** 2307282043

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `status` - Status

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionx509identity_provider` - X509 Identity Provider Extension Schema
	* `cert_match_attribute` - X509 Certificate Matching Attribute

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `crl_check_on_ocsp_failure_enabled` - Fallback on CRL Validation if OCSP fails.

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `crl_enabled` - Set to true to enable CRL Validation

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `crl_location` - CRL Location URL

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `crl_reload_duration` - Fetch the CRL contents every X minutes

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `eku_validation_enabled` - Set to true to enable EKU Validation

		**Added In:** 2304270343

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `eku_values` - List of EKU which needs to be validated

		**Added In:** 2304270343

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ocsp_allow_unknown_response_status` - Allow access if OCSP response is UNKNOWN or OCSP Responder does not respond within the timeout duration

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `ocsp_enable_signed_response` - Describes if the OCSP response is signed

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `ocsp_enabled` - Set to true to enable OCSP Validation

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `ocsp_responder_url` - This property specifies OCSP Responder URL.

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ocsp_revalidate_time` - Revalidate OCSP status for user after X hours

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* idcsMaxValue: 24
		* idcsMinValue: 0
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: integer
		* uniqueness: none
	* `ocsp_server_name` - This property specifies the OCSP Server alias name

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ocsp_trust_cert_chain` - OCSP Trusted Certificate Chain

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `other_cert_match_attribute` - Check for specific conditions of other certificate attributes

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `signing_certificate_chain` - Certificate alias list to create a chain for the incoming client certificate

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: true
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `user_match_attribute` - This property specifies the userstore attribute value that must match the incoming certificate attribute.

		**Added In:** 2010242156

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `user_mapping_method` - User mapping method.

	**Deprecated Since: 20.1.3**

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
	* idcsValuePersistedInOtherAttribute: true
* `user_mapping_store_attribute` - This property specifies the userstore attribute value that must match the incoming assertion attribute value or the incoming nameid attribute value in order to identify the user during SSO.<br>You can construct the userMappingStoreAttribute value by specifying attributes from the Oracle Identity Cloud Service Core Users schema. For examples of how to construct the userMappingStoreAttribute value, see the <b>Example of a Request Body</b> section of the Examples tab for the <a href='./op-admin-v1-identityproviders-post.html'>POST</a> and <a href='./op-admin-v1-identityproviders-id-put.html'>PUT</a> methods of the /IdentityProviders endpoint.

	**Deprecated Since: 20.1.3**

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
	* idcsValuePersistedInOtherAttribute: true

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Identity Provider
	* `update` - (Defaults to 20 minutes), when updating the Identity Provider
	* `delete` - (Defaults to 20 minutes), when destroying the Identity Provider


## Import

IdentityProviders can be imported using the `id`, e.g.

```
$ terraform import oci_identity_domains_identity_provider.test_identity_provider "idcsEndpoint/{idcsEndpoint}/identityProviders/{identityProviderId}" 
```

