---
subcategory: "Identity Domains"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_domains_oauth_client_certificates"
sidebar_current: "docs-oci-datasource-identity_domains-oauth_client_certificates"
description: |-
  Provides the list of O Auth Client Certificates in Oracle Cloud Infrastructure Identity Domains service
---

# Data Source: oci_identity_domains_oauth_client_certificates
This data source provides the list of O Auth Client Certificates in Oracle Cloud Infrastructure Identity Domains service.

Search OAuth Client Certificates

## Example Usage

```hcl
data "oci_identity_domains_oauth_client_certificates" "test_oauth_client_certificates" {
	#Required
	idcs_endpoint = data.oci_identity_domain.test_domain.url

	#Optional
	oauth_client_certificate_count = var.oauth_client_certificate_oauth_client_certificate_count
	oauth_client_certificate_filter = var.oauth_client_certificate_oauth_client_certificate_filter
	authorization = var.oauth_client_certificate_authorization
	resource_type_schema_version = var.oauth_client_certificate_resource_type_schema_version
	start_index = var.oauth_client_certificate_start_index
}
```

## Argument Reference

The following arguments are supported:

* `oauth_client_certificate_count` - (Optional) OPTIONAL. An integer that indicates the desired maximum number of query results per page. 1000 is the largest value that you can use. See the Pagination section of the System for Cross-Domain Identity Management Protocol specification for more information. (Section 3.4.2.4).
* `oauth_client_certificate_filter` - (Optional) OPTIONAL. The filter string that is used to request a subset of resources. The filter string MUST be a valid filter expression. See the Filtering section of the SCIM specification for more information (Section 3.4.2.2). The string should contain at least one condition that each item must match in order to be returned in the search results. Each condition specifies an attribute, an operator, and a value. Conditions within a filter can be connected by logical operators (such as AND and OR). Sets of conditions can be grouped together using parentheses.
* `authorization` - (Optional) The Authorization field value consists of credentials containing the authentication information of the user agent for the realm of the resource being requested.
* `idcs_endpoint` - (Required) The basic endpoint for the identity domain
* `resource_type_schema_version` - (Optional) An endpoint-specific schema version number to use in the Request. Allowed version values are Earliest Version or Latest Version as specified in each REST API endpoint description, or any sequential number inbetween. All schema attributes/body parameters are a part of version 1. After version 1, any attributes added or deprecated will be tagged with the version that they were added to or deprecated in. If no version is provided, the latest schema version is returned.
* `start_index` - (Optional) OPTIONAL. An integer that indicates the 1-based index of the first query result. See the Pagination section of the SCIM specification for more information. (Section 3.4.2.4). The number of results pages to return. The first page is 1. Specify 2 to access the second page of results, and so on.


## Attributes Reference

The following attributes are exported:

* `oauth_client_certificates` - The list of oauth_client_certificates.

### OAuthClientCertificate Reference

The following attributes are exported:

* `cert_end_date` - Certificate end date

	**SCIM++ Properties:**
	* type: dateTime
	* multiValued: false
	* required: false
	* caseExact: true
	* mutability: readWrite
	* returned: default
	* uniqueness: none
* `cert_start_date` - Certificate start date

	**SCIM++ Properties:**
	* type: dateTime
	* multiValued: false
	* required: false
	* caseExact: true
	* mutability: readWrite
	* returned: default
	* uniqueness: none
* `certificate_alias` - Certificate alias

	**SCIM++ Properties:**
	* type: string
	* multiValued: false
	* required: true
	* caseExact: false
	* mutability: readWrite
	* returned: always
	* uniqueness: none
	* idcsSearchable: true
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
* `external_id` - An identifier for the Resource as defined by the Service Consumer. The externalId may simplify identification of the Resource between Service Consumer and Service Provider by allowing the Consumer to refer to the Resource with its own identifier, obviating the need to store a local mapping between the local identifier of the Resource and the identifier used by the Service Provider. Each Resource MAY include a non-empty externalId value. The value of the externalId attribute is always issued by the Service Consumer and can never be specified by the Service Provider. The Service Provider MUST always interpret the externalId as scoped to the Service Consumer's tenant.

	**SCIM++ Properties:**
	* type: string
	* multiValued: false
	* required: false
	* caseExact: false
	* mutability: readWrite
	* returned: default
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
* `key_store_id` - Key store ID

	**SCIM++ Properties:**
	* type: string
	* multiValued: false
	* required: false
	* caseExact: true
	* returned: never
	* uniqueness: none
	* idcsSearchable: true
* `key_store_name` - Key store name

	**SCIM++ Properties:**
	* type: string
	* multiValued: false
	* required: false
	* caseExact: false
	* returned: always
	* uniqueness: none
* `key_store_password` - Key store password

	**SCIM++ Properties:**
	* type: string
	* multiValued: false
	* idcsSensitive: none
	* required: false
	* caseExact: true
	* returned: never
	* uniqueness: none
* `map` - Map

	**SCIM++ Properties:**
	* type: string
	* multiValued: false
	* required: false
	* caseExact: false
	* returned: always
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
* `sha1thumbprint` - SHA-1 Thumbprint

	**SCIM++ Properties:**
	* type: string
	* multiValued: false
	* idcsSearchable: true
	* required: false
	* caseExact: true
	* returned: default
	* uniqueness: none
* `sha256thumbprint` - SHA-256 Thumbprint

	**SCIM++ Properties:**
	* type: string
	* multiValued: false
	* idcsSearchable: true
	* required: false
	* caseExact: true
	* returned: default
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
* `x509base64certificate` - Base 64Key data attribute

	**SCIM++ Properties:**
	* caseExact: true
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: default
	* type: string
	* uniqueness: none

