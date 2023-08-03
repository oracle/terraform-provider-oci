---
subcategory: "Identity Domains"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_domains_my_apps"
sidebar_current: "docs-oci-datasource-identity_domains-my_apps"
description: |-
  Provides the list of My Apps in Oracle Cloud Infrastructure Identity Domains service
---

# Data Source: oci_identity_domains_my_apps
This data source provides the list of My Apps in Oracle Cloud Infrastructure Identity Domains service.

Search My Apps

## Example Usage

```hcl
data "oci_identity_domains_my_apps" "test_my_apps" {
	#Required
	idcs_endpoint = data.oci_identity_domain.test_domain.url

	#Optional
	my_app_count = var.my_app_my_app_count
	my_app_filter = var.my_app_my_app_filter
	authorization = var.my_app_authorization
	resource_type_schema_version = var.my_app_resource_type_schema_version
	start_index = var.my_app_start_index
}
```

## Argument Reference

The following arguments are supported:

* `my_app_count` - (Optional) OPTIONAL. An integer that indicates the desired maximum number of query results per page. 1000 is the largest value that you can use. See the Pagination section of the System for Cross-Domain Identity Management Protocol specification for more information. (Section 3.4.2.4).
* `my_app_filter` - (Optional) OPTIONAL. The filter string that is used to request a subset of resources. The filter string MUST be a valid filter expression. See the Filtering section of the SCIM specification for more information (Section 3.4.2.2). The string should contain at least one condition that each item must match in order to be returned in the search results. Each condition specifies an attribute, an operator, and a value. Conditions within a filter can be connected by logical operators (such as AND and OR). Sets of conditions can be grouped together using parentheses.
* `authorization` - (Optional) The Authorization field value consists of credentials containing the authentication information of the user agent for the realm of the resource being requested.
* `idcs_endpoint` - (Required) The basic endpoint for the identity domain
* `resource_type_schema_version` - (Optional) An endpoint-specific schema version number to use in the Request. Allowed version values are Earliest Version or Latest Version as specified in each REST API endpoint description, or any sequential number inbetween. All schema attributes/body parameters are a part of version 1. After version 1, any attributes added or deprecated will be tagged with the version that they were added to or deprecated in. If no version is provided, the latest schema version is returned.
* `start_index` - (Optional) OPTIONAL. An integer that indicates the 1-based index of the first query result. See the Pagination section of the SCIM specification for more information. (Section 3.4.2.4). The number of results pages to return. The first page is 1. Specify 2 to access the second page of results, and so on.


## Attributes Reference

The following attributes are exported:

* `my_apps` - The list of my_apps.

### MyApp Reference

The following attributes are exported:

* `resources` - A multi-valued list of complex objects containing the requested resources. This MAY be a subset of the full set of resources if pagination is requested. REQUIRED if "totalResults" is non-zero.
	* `account_type` - Type of the Acccount

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `active` - If true, the account is activated.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `app` - Application on which the account is based

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: complex
		* uniqueness: none
		* `active` - If true, this App is able to participate in runtime services, such as automatic-login, OAuth, and SAML. If false, all runtime services are disabled for this App, and only administrative operations can be performed.

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: boolean
			* uniqueness: none
		* `app_icon` - Application icon.

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `app_thumbnail` - Application thumbnail.

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: request
			* type: string
			* uniqueness: none
		* `description` - Application description

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `display` - Application display name

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `is_alias_app` - If true, this App is an AliasApp and it cannot be granted to an end user directly

			**SCIM++ Properties:**
			* idcsSearchable: true
			* multiValued: false
			* mutability: writeOnly
			* required: false
			* returned: never
			* type: boolean
			* uniqueness: none
		* `is_login_target` - If true, this App allows runtime services to log end users into this App automatically.

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: boolean
			* uniqueness: none
		* `is_opc_service` - If true, this application is an Oracle Public Cloud service-instance.

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: boolean
			* uniqueness: none
		* `login_mechanism` - The protocol that runtime services will use to log end users in to this App automatically. If 'OIDC', then runtime services use the OpenID Connect protocol. If 'SAML', then runtime services use Security Assertion Markup Language protocol.

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - Application URI

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `service_type_urn` - This Uniform Resource Name (URN) value identifies the type of Oracle Public Cloud service of which this app is an instance.

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `show_in_my_apps` - If true, this App will be displayed in the MyApps page of each end-user who has access to this App.

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: boolean
			* uniqueness: none
		* `value` - Application identifier

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: true
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
	* `favorite` - If true, this account has been marked as a favorite of the User who owns it.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
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
	* `is_account` - If true, indicates that this managed object is an account, which is an identity that represents a user in the context of a specific application.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: boolean
		* uniqueness: none
	* `last_accessed` - Last accessed timestamp of an application

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: dateTime
		* uniqueness: none
	* `launch_url` - The URL that will be used to launch the application.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
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
	* `name` - Name of the Account

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: always
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
	* `owner` - Owning user of the account

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `display` - User display name

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - User URI

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - User Identifier

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: true
			* returned: default
			* type: string
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
	* `uid` - Unique identifier of the Account.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `user_wallet_artifact` - The UserWalletArtifact that contains the credentials that the system will use in performing Secure Form-Fill to log the User into this application.

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: complex
		* uniqueness: none
		* `ref` - UserWalletArtifact URI

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - UserWalletArtifact identifier

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: true
			* returned: always
			* type: string
			* uniqueness: none
* `items_per_page` - The number of resources returned in a list response page. REQUIRED when partial results returned due to pagination.
* `schemas` - The schemas attribute is an array of Strings which allows introspection of the supported schema version for a SCIM representation as well any schema extensions supported by that representation. Each String value must be a unique URI. All representations of SCIM schema MUST include a non-zero value array with value(s) of the URIs supported by that representation. Duplicate values MUST NOT be included. Value order is not specified and MUST not impact behavior. REQUIRED.
* `start_index` - The 1-based index of the first result in the current set of list results.  REQUIRED when partial results returned due to pagination.
* `total_results` - The total number of results returned by the list or query operation.  The value may be larger than the number of resources returned such as when returning a single page of results where multiple pages are available. REQUIRED.

