---
subcategory: "Identity Domains"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_domains_identity_propagation_trust"
sidebar_current: "docs-oci-resource-identity_domains-identity_propagation_trust"
description: |-
  Provides the Identity Propagation Trust resource in Oracle Cloud Infrastructure Identity Domains service
---

# oci_identity_domains_identity_propagation_trust
This resource provides the Identity Propagation Trust resource in Oracle Cloud Infrastructure Identity Domains service.

Register a new Identity Propagation Trust configuration.

## Example Usage

```hcl
resource "oci_identity_domains_identity_propagation_trust" "test_identity_propagation_trust" {
	#Required
	idcs_endpoint = data.oci_identity_domain.test_domain.url
	issuer = var.identity_propagation_trust_issuer
	name = var.identity_propagation_trust_name
	schemas = ["urn:ietf:params:scim:schemas:oracle:idcs:IdentityPropagationTrust"]
	type = var.identity_propagation_trust_type

	#Optional
	account_id = "accountId"
	active = var.identity_propagation_trust_active
	allow_impersonation = var.identity_propagation_trust_allow_impersonation
	attribute_sets = ["all"]
	attributes = ""
	authorization = var.identity_propagation_trust_authorization
	client_claim_name = var.identity_propagation_trust_client_claim_name
	client_claim_values = ["clientClaimValues"]
	clock_skew_seconds = var.identity_propagation_trust_clock_skew_seconds
	description = var.identity_propagation_trust_description
	impersonation_service_users {
		#Required
		rule = var.identity_propagation_trust_impersonation_service_users_rule
		value = oci_identity_domains_user.test_identity_propagation_trust_user.id

		#Optional
		ocid = var.identity_propagation_trust_impersonation_service_users_ocid
	}
	keytab {
		#Required
		secret_ocid = var.identity_propagation_trust_keytab_secret_ocid

		#Optional
		secret_version = var.identity_propagation_trust_keytab_secret_version
	}
	oauth_clients = ["oauthClients"]
	ocid = var.identity_propagation_trust_ocid
	public_certificate = var.identity_propagation_trust_public_certificate
	public_key_endpoint = var.identity_propagation_trust_public_key_endpoint
	resource_type_schema_version = var.identity_propagation_trust_resource_type_schema_version
	subject_claim_name = var.identity_propagation_trust_subject_claim_name
	subject_mapping_attribute = var.identity_propagation_trust_subject_mapping_attribute
	subject_type = var.identity_propagation_trust_subject_type
	tags {
		#Required
		key = var.identity_propagation_trust_tags_key
		value = var.identity_propagation_trust_tags_value
	}
}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Optional) (Updatable) The Identity cloud provider service identifier, for example, the Azure Tenancy ID, AWS Account ID, or GCP Project ID.

	**SCIM++ Properties:**
	* type: string
	* multiValued: false
	* required: false
	* mutability: readWrite
	* returned: default
	* caseExact: true
	* idcsSearchable: true
	* uniqueness: none
* `active` - (Optional) (Updatable) If true, specifies that this Identity Propagation Trust is in an enabled state. The default value is false.

	**SCIM++ Properties:**
	* type: boolean
	* multiValued: false
	* required: false
	* mutability: readWrite
	* returned: default
	* uniqueness: none
	* idcsSearchable: true
* `allow_impersonation` - (Optional) (Updatable) Allow customers to define whether the resulting token should contain the authenticated user as the subject or whether the token should impersonate another Application Principal in IAM.

	**SCIM++ Properties:**
	* type: boolean
	* multiValued: false
	* required: false
	* mutability: readWrite
	* returned: default
	* uniqueness: none
	* idcsSearchable: false
* `attribute_sets` - (Optional) (Updatable) A multi-valued list of strings indicating the return type of attribute definition. The specified set of attributes can be fetched by the return type of the attribute. One or more values can be given together to fetch more than one group of attributes. If 'attributes' query parameter is also available, union of the two is fetched. Valid values - all, always, never, request, default. Values are case-insensitive.
* `attributes` - (Optional) (Updatable) A comma-delimited string that specifies the names of resource attributes that should be returned in the response. By default, a response that contains resource attributes contains only attributes that are defined in the schema for that resource type as returned=always or returned=default. An attribute that is defined as returned=request is returned in a response only if the request specifies its name in the value of this query parameter. If a request specifies this query parameter, the response contains the attributes that this query parameter specifies, as well as any attribute that is defined as returned=always.
* `authorization` - (Optional) (Updatable) The Authorization field value consists of credentials containing the authentication information of the user agent for the realm of the resource being requested.
* `client_claim_name` - (Optional) (Updatable) The claim name that identifies to whom the JWT/SAML token is issued. If AWS, then \"aud\" or \"client_id\". If Azure, then \"appid\". If GCP, then \"aud\".

	**SCIM++ Properties:**
	* type: string
	* multiValued: false
	* required: false
	* mutability: readWrite
	* returned: default
	* uniqueness: none
	* idcsSearchable: false
* `client_claim_values` - (Optional) (Updatable) The value that corresponds to the client claim name used to identify to whom the token is issued.

	**SCIM++ Properties:**
	* type: string
	* multiValued: true
	* required: false
	* mutability: readWrite
	* returned: default
	* uniqueness: none
	* caseExact: true
	* idcsSearchable: false
* `clock_skew_seconds` - (Optional) (Updatable) The clock skew (in secs) that's allowed for the token issue and expiry time.

	**Added In:** 2308181911

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
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
* `description` - (Optional) (Updatable) The description of the Identity Propagation Trust.

	**SCIM++ Properties:**
	* type: string
	* multiValued: false
	* required: false
	* mutability: readWrite
	* returned: default
	* uniqueness: none
	* caseExact: false
	* idcsSearchable: false
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
	* `_ref` - (Optional) (Updatable) The URI of the SCIM resource that represents the User or App who created this Resource

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
	* `_ref` - (Optional) (Updatable) The URI of the SCIM resource that represents the User or App who modified this Resource

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
* `impersonation_service_users` - (Optional) (Updatable) The Impersonating Principal.

	**SCIM++ Properties:**
	* idcsCompositeKey: [rule, value]
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: request
	* type: complex
	* uniqueness: none
	* `_ref` - (Optional) (Updatable) The URI that corresponds to the Service User.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `ocid` - (Optional) (Updatable) The OCID of the Service User.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `rule` - (Required) (Updatable) The rule expression to be used for matching the inbound token for impersonation.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - (Required) (Updatable) The ID of the Service User.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `issuer` - (Required) (Updatable) The issuer claim of the Identity provider.

	**SCIM++ Properties:**
	* type: string
	* multiValued: false
	* required: true
	* mutability: readWrite
	* returned: always
	* caseExact: true
	* idcsSearchable: true
	* uniqueness: server
* `keytab` - (Optional) (Updatable) The keytab stored in the tenancy's Vault. This is required if the identity propagation type is 'SPNEGO'.

	**SCIM++ Properties:**
	* idcsCompositeKey: [secretOcid]
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `secret_ocid` - (Required) (Updatable) The OCID of the secret. The secret content corresponding to the OCID is expected to be in Base64 encoded content type.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `secret_version` - (Optional) (Updatable) The version of the secret. When the version is not specified, then the latest secret version is used during runtime.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
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
* `name` - (Required) The name of the the Identity Propagation Trust.

	**SCIM++ Properties:**
	* type: string
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* required: true
	* mutability: immutable
	* returned: default
	* uniqueness: none
* `oauth_clients` - (Optional) (Updatable) The value of all the authorized OAuth Clients.

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: true
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
* `public_certificate` - (Optional) (Updatable) Store the public key if public key cert.

	**SCIM++ Properties:**
	* type: string
	* multiValued: false
	* required: false
	* mutability: readWrite
	* returned: default
	* uniqueness: none
	* idcsSearchable: false
* `public_key_endpoint` - (Optional) (Updatable) The cloud provider's public key API of SAML and OIDC providers for signature validation.

	**SCIM++ Properties:**
	* type: string
	* multiValued: false
	* required: false
	* mutability: readWrite
	* returned: default
	* uniqueness: none
	* caseExact: false
	* idcsSearchable: false
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
* `subject_claim_name` - (Optional) (Updatable) Used for locating the subject claim from the incoming token.

	**SCIM++ Properties:**
	* type: string
	* multiValued: false
	* required: false
	* mutability: readWrite
	* returned: default
	* uniqueness: none
	* caseExact: true
	* idcsSearchable: false
* `subject_mapping_attribute` - (Optional) (Updatable) Subject Mapping Attribute to which the value from subject claim name value would be used for identity lookup.

	**SCIM++ Properties:**
	* type: string
	* multiValued: false
	* idcsSearchable: false
	* required: false
	* mutability: readWrite
	* returned: default
	* uniqueness: none
* `subject_type` - (Optional) (Updatable) The type of the resource against which lookup will be made in the identity domain in IAM for the incoming subject claim value.

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
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
* `type` - (Required) (Updatable) The type of the inbound token from the Identity cloud provider.

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: false
	* required: true
	* mutability: readWrite
	* returned: default
	* type: string
	* multiValued: false
	* uniqueness: none


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `account_id` - The Identity cloud provider service identifier, for example, the Azure Tenancy ID, AWS Account ID, or GCP Project ID.

	**SCIM++ Properties:**
	* type: string
	* multiValued: false
	* required: false
	* mutability: readWrite
	* returned: default
	* caseExact: true
	* idcsSearchable: true
	* uniqueness: none
* `active` - If true, specifies that this Identity Propagation Trust is in an enabled state. The default value is false.

	**SCIM++ Properties:**
	* type: boolean
	* multiValued: false
	* required: false
	* mutability: readWrite
	* returned: default
	* uniqueness: none
	* idcsSearchable: true
* `allow_impersonation` - Allow customers to define whether the resulting token should contain the authenticated user as the subject or whether the token should impersonate another Application Principal in IAM.

	**SCIM++ Properties:**
	* type: boolean
	* multiValued: false
	* required: false
	* mutability: readWrite
	* returned: default
	* uniqueness: none
	* idcsSearchable: false
* `client_claim_name` - The claim name that identifies to whom the JWT/SAML token is issued. If AWS, then \"aud\" or \"client_id\". If Azure, then \"appid\". If GCP, then \"aud\".

	**SCIM++ Properties:**
	* type: string
	* multiValued: false
	* required: false
	* mutability: readWrite
	* returned: default
	* uniqueness: none
	* idcsSearchable: false
* `client_claim_values` - The value that corresponds to the client claim name used to identify to whom the token is issued.

	**SCIM++ Properties:**
	* type: string
	* multiValued: true
	* required: false
	* mutability: readWrite
	* returned: default
	* uniqueness: none
	* caseExact: true
	* idcsSearchable: false
* `clock_skew_seconds` - The clock skew (in secs) that's allowed for the token issue and expiry time.

	**Added In:** 2308181911

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: integer
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
* `description` - The description of the Identity Propagation Trust.

	**SCIM++ Properties:**
	* type: string
	* multiValued: false
	* required: false
	* mutability: readWrite
	* returned: default
	* uniqueness: none
	* caseExact: false
	* idcsSearchable: false
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
	* `_ref` - The URI of the SCIM resource that represents the User or App who created this Resource

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
	* `_ref` - The URI of the SCIM resource that represents the User or App who modified this Resource

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
* `impersonation_service_users` - The Impersonating Principal.

	**SCIM++ Properties:**
	* idcsCompositeKey: [rule, value]
	* idcsSearchable: false
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: request
	* type: complex
	* uniqueness: none
	* `_ref` - The URI that corresponds to the Service User.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `ocid` - The OCID of the Service User.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `rule` - The rule expression to be used for matching the inbound token for impersonation.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - The ID of the Service User.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `issuer` - The issuer claim of the Identity provider.

	**SCIM++ Properties:**
	* type: string
	* multiValued: false
	* required: true
	* mutability: readWrite
	* returned: always
	* caseExact: true
	* idcsSearchable: true
	* uniqueness: server
* `keytab` - The keytab stored in the tenancy's Vault. This is required if the identity propagation type is 'SPNEGO'.

	**SCIM++ Properties:**
	* idcsCompositeKey: [secretOcid]
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `secret_ocid` - The OCID of the secret. The secret content corresponding to the OCID is expected to be in Base64 encoded content type.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `secret_version` - The version of the secret. When the version is not specified, then the latest secret version is used during runtime.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
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
* `name` - The name of the the Identity Propagation Trust.

	**SCIM++ Properties:**
	* type: string
	* caseExact: false
	* idcsSearchable: true
	* multiValued: false
	* required: true
	* mutability: immutable
	* returned: default
	* uniqueness: none
* `oauth_clients` - The value of all the authorized OAuth Clients.

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: true
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
* `public_certificate` - Store the public key if public key cert.

	**SCIM++ Properties:**
	* type: string
	* multiValued: false
	* required: false
	* mutability: readWrite
	* returned: default
	* uniqueness: none
	* idcsSearchable: false
* `public_key_endpoint` - The cloud provider's public key API of SAML and OIDC providers for signature validation.

	**SCIM++ Properties:**
	* type: string
	* multiValued: false
	* required: false
	* mutability: readWrite
	* returned: default
	* uniqueness: none
	* caseExact: false
	* idcsSearchable: false
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
* `subject_claim_name` - Used for locating the subject claim from the incoming token.

	**SCIM++ Properties:**
	* type: string
	* multiValued: false
	* required: false
	* mutability: readWrite
	* returned: default
	* uniqueness: none
	* caseExact: true
	* idcsSearchable: false
* `subject_mapping_attribute` - Subject Mapping Attribute to which the value from subject claim name value would be used for identity lookup.

	**SCIM++ Properties:**
	* type: string
	* multiValued: false
	* idcsSearchable: false
	* required: false
	* mutability: readWrite
	* returned: default
	* uniqueness: none
* `subject_type` - The type of the resource against which lookup will be made in the identity domain in IAM for the incoming subject claim value.

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
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
* `type` - The type of the inbound token from the Identity cloud provider.

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: false
	* required: true
	* mutability: readWrite
	* returned: default
	* type: string
	* multiValued: false
	* uniqueness: none

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Identity Propagation Trust
	* `update` - (Defaults to 20 minutes), when updating the Identity Propagation Trust
	* `delete` - (Defaults to 20 minutes), when destroying the Identity Propagation Trust


## Import

IdentityPropagationTrusts can be imported using the `id`, e.g.

```
$ terraform import oci_identity_domains_identity_propagation_trust.test_identity_propagation_trust "idcsEndpoint/{idcsEndpoint}/identityPropagationTrusts/{identityPropagationTrustId}" 
```

