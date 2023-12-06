---
subcategory: "Identity Domains"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_domains_dynamic_resource_group"
sidebar_current: "docs-oci-resource-identity_domains-dynamic_resource_group"
description: |-
  Provides the Dynamic Resource Group resource in Oracle Cloud Infrastructure Identity Domains service
---

# oci_identity_domains_dynamic_resource_group
This resource provides the Dynamic Resource Group resource in Oracle Cloud Infrastructure Identity Domains service.

Create a Dynamic Resource Group.

## Example Usage

```hcl
resource "oci_identity_domains_dynamic_resource_group" "test_dynamic_resource_group" {
	#Required
	display_name = var.dynamic_resource_group_display_name
	idcs_endpoint = data.oci_identity_domain.test_domain.url
	matching_rule = var.dynamic_resource_group_matching_rule
	schemas = ["urn:ietf:params:scim:schemas:oracle:idcs:DynamicResourceGroup"]

	#Optional
	attribute_sets = []
	attributes = ""
	authorization = var.dynamic_resource_group_authorization
	description = var.dynamic_resource_group_description
	id = var.dynamic_resource_group_id
	ocid = var.dynamic_resource_group_ocid
	resource_type_schema_version = var.dynamic_resource_group_resource_type_schema_version
	tags {
		#Required
		key = var.dynamic_resource_group_tags_key
		value = var.dynamic_resource_group_tags_value
	}
	urnietfparamsscimschemasoracleidcsextension_oci_tags {

		#Optional
		defined_tags {
			#Required
			key = var.dynamic_resource_group_urnietfparamsscimschemasoracleidcsextension_oci_tags_defined_tags_key
			namespace = var.dynamic_resource_group_urnietfparamsscimschemasoracleidcsextension_oci_tags_defined_tags_namespace
			value = var.dynamic_resource_group_urnietfparamsscimschemasoracleidcsextension_oci_tags_defined_tags_value
		}
		freeform_tags {
			#Required
			key = var.dynamic_resource_group_urnietfparamsscimschemasoracleidcsextension_oci_tags_freeform_tags_key
			value = var.dynamic_resource_group_urnietfparamsscimschemasoracleidcsextension_oci_tags_freeform_tags_value
		}
	}
}
```

## Argument Reference

The following arguments are supported:

* `attribute_sets` - (Optional) (Updatable) A multi-valued list of strings indicating the return type of attribute definition. The specified set of attributes can be fetched by the return type of the attribute. One or more values can be given together to fetch more than one group of attributes. If 'attributes' query parameter is also available, union of the two is fetched. Valid values - all, always, never, request, default. Values are case-insensitive.
* `attributes` - (Optional) (Updatable) A comma-delimited string that specifies the names of resource attributes that should be returned in the response. By default, a response that contains resource attributes contains only attributes that are defined in the schema for that resource type as returned=always or returned=default. An attribute that is defined as returned=request is returned in a response only if the request specifies its name in the value of this query parameter. If a request specifies this query parameter, the response contains the attributes that this query parameter specifies, as well as any attribute that is defined as returned=always.
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
* `description` - (Optional) (Updatable) text that explains the purpose of this Dynamic Resource Group

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCsvAttributeName: Description
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Description]]
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: always
	* type: string
	* uniqueness: none
* `display_name` - (Required) (Updatable) User-friendly, mutable identifier

	**SCIM++ Properties:**
	* idcsCsvAttributeName: Display Name
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Name, deprecatedColumnHeaderName:Display Name]]
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: always
	* type: string
	* uniqueness: global
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
* `dynamic_group_app_roles` - (Optional) (Updatable) A list of appRoles that are currently granted to this Dynamic Resource Group.  The Identity service will assert these AppRoles for any resource that satisfies the matching-rule of this DynamicResourceGroup.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCompositeKey: [value]
	* idcsSearchable: true
	* multiValued: true
	* mutability: readOnly
	* required: false
	* returned: request
	* type: complex
	* uniqueness: none
	* `admin_role` - (Optional) (Updatable) If true, then the role provides administrative access privileges. READ-ONLY.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: boolean
		* uniqueness: none
	* `app_id` - (Optional) (Updatable) ID of parent App. READ-ONLY.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `app_name` - (Optional) (Updatable) Name of parent App. READ-ONLY.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `display` - (Optional) (Updatable) A human readable name, primarily used for display purposes. READ-ONLY.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `legacy_group_name` - (Optional) (Updatable) The name of the legacy group associated with this AppRole.

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `ref` - (Optional) (Updatable) The URI of the corresponding appRole resource to which the user belongs

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: reference
		* uniqueness: none
	* `value` - (Required) (Updatable) The identifier of the appRole

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: always
		* type: string
		* uniqueness: none
* `grants` - (Optional) (Updatable) Grants assigned to group

	**SCIM++ Properties:**
	* idcsAddedSinceVersion: 3
	* idcsSearchable: true
	* multiValued: true
	* mutability: readOnly
	* required: false
	* returned: request
	* type: complex
	* uniqueness: none
	* `app_id` - (Optional) (Updatable) App identifier

		**SCIM++ Properties:**
		* caseExact: true
		* idcsAddedSinceVersion: 3
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `grant_mechanism` - (Optional) (Updatable) Each value of grantMechanism indicates how (or by what component) some App (or App-Entitlement) was granted. A customer or the UI should use only grantMechanism values that start with 'ADMINISTRATOR':
		* 'ADMINISTRATOR_TO_USER' is for a direct grant to a specific User.
		* 'ADMINISTRATOR_TO_GROUP' is for a grant to a specific Group, which results in indirect grants to Users who are members of that Group.
		* 'ADMINISTRATOR_TO_APP' is for a grant to a specific App.  The grantee (client) App gains access to the granted (server) App.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsAddedSinceVersion: 3
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ref` - (Optional) (Updatable) Grant URI

		**SCIM++ Properties:**
		* idcsAddedSinceVersion: 3
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - (Optional) (Updatable) Grant identifier

		**SCIM++ Properties:**
		* caseExact: true
		* idcsAddedSinceVersion: 3
		* idcsSearchable: true
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
* `matching_rule` - (Required) (Updatable) Store as a string the matching-rule for this Dynamic Resource Group. This may match any number of Apps in this Domain, as well as matching any number of Oracle Cloud Infrastructure resources that are not in any Domain but that are in the Oracle Cloud Infrastructure Compartment that contains this Domain.

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: request
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
* `urnietfparamsscimschemasoracleidcsextension_oci_tags` - (Optional) (Updatable) Oracle Cloud Infrastructure Tags.
	* `defined_tags` - (Optional) (Updatable) Oracle Cloud Infrastructure Defined Tags

		**Added In:** 2011192329

		**SCIM++ Properties:**
		* idcsCompositeKey: [namespace, key, value]
		* type: complex
		* idcsSearchable: true
		* required: false
		* mutability: readWrite
		* multiValued: true
		* returned: default
		* `key` - (Required) (Updatable) Oracle Cloud Infrastructure Tag key

			**Added In:** 2011192329

			**SCIM++ Properties:**
			* caseExact: false
			* type: string
			* required: true
			* mutability: readWrite
			* returned: default
			* idcsSearchable: true
			* uniqueness: none
		* `namespace` - (Required) (Updatable) Oracle Cloud Infrastructure Tag namespace

			**Added In:** 2011192329

			**SCIM++ Properties:**
			* caseExact: false
			* type: string
			* required: true
			* mutability: readWrite
			* returned: default
			* idcsSearchable: true
			* uniqueness: none
		* `value` - (Required) (Updatable) Oracle Cloud Infrastructure Tag value

			**Added In:** 2011192329

			**SCIM++ Properties:**
			* caseExact: false
			* required: true
			* idcsReturnEmptyWhenNull: true
			* mutability: readWrite
			* returned: default
			* type: string
			* idcsSearchable: true
			* uniqueness: none
	* `freeform_tags` - (Optional) (Updatable) Oracle Cloud Infrastructure Freeform Tags

		**Added In:** 2011192329

		**SCIM++ Properties:**
		* idcsCompositeKey: [key, value]
		* idcsSearchable: true
		* type: complex
		* required: false
		* mutability: readWrite
		* returned: default
		* multiValued: true
		* `key` - (Required) (Updatable) Oracle Cloud Infrastructure Tag key

			**Added In:** 2011192329

			**SCIM++ Properties:**
			* caseExact: false
			* type: string
			* required: true
			* mutability: readWrite
			* returned: default
			* idcsSearchable: true
			* uniqueness: none
		* `value` - (Required) (Updatable) Oracle Cloud Infrastructure Tag value

			**Added In:** 2011192329

			**SCIM++ Properties:**
			* caseExact: false
			* required: true
			* idcsReturnEmptyWhenNull: true
			* mutability: readWrite
			* returned: default
			* type: string
			* idcsSearchable: true
			* uniqueness: none
	* `tag_slug` - (Optional) (Updatable) Oracle Cloud Infrastructure Tag slug

		**Added In:** 2011192329

		**SCIM++ Properties:**
		* type: binary
		* mutability: readOnly
		* returned: request


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

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
* `description` - text that explains the purpose of this Dynamic Resource Group

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCsvAttributeName: Description
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Description]]
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: always
	* type: string
	* uniqueness: none
* `display_name` - User-friendly, mutable identifier

	**SCIM++ Properties:**
	* idcsCsvAttributeName: Display Name
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Name, deprecatedColumnHeaderName:Display Name]]
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: always
	* type: string
	* uniqueness: global
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
* `dynamic_group_app_roles` - A list of appRoles that are currently granted to this Dynamic Resource Group.  The Identity service will assert these AppRoles for any resource that satisfies the matching-rule of this DynamicResourceGroup.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCompositeKey: [value]
	* idcsSearchable: true
	* multiValued: true
	* mutability: readOnly
	* required: false
	* returned: request
	* type: complex
	* uniqueness: none
	* `admin_role` - If true, then the role provides administrative access privileges. READ-ONLY.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: boolean
		* uniqueness: none
	* `app_id` - ID of parent App. READ-ONLY.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `app_name` - Name of parent App. READ-ONLY.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `display` - A human readable name, primarily used for display purposes. READ-ONLY.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `legacy_group_name` - The name of the legacy group associated with this AppRole.

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `ref` - The URI of the corresponding appRole resource to which the user belongs

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: reference
		* uniqueness: none
	* `value` - The identifier of the appRole

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: always
		* type: string
		* uniqueness: none
* `grants` - Grants assigned to group

	**SCIM++ Properties:**
	* idcsAddedSinceVersion: 3
	* idcsSearchable: true
	* multiValued: true
	* mutability: readOnly
	* required: false
	* returned: request
	* type: complex
	* uniqueness: none
	* `app_id` - App identifier

		**SCIM++ Properties:**
		* caseExact: true
		* idcsAddedSinceVersion: 3
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `grant_mechanism` - Each value of grantMechanism indicates how (or by what component) some App (or App-Entitlement) was granted. A customer or the UI should use only grantMechanism values that start with 'ADMINISTRATOR':
		* 'ADMINISTRATOR_TO_USER' is for a direct grant to a specific User.
		* 'ADMINISTRATOR_TO_GROUP' is for a grant to a specific Group, which results in indirect grants to Users who are members of that Group.
		* 'ADMINISTRATOR_TO_APP' is for a grant to a specific App.  The grantee (client) App gains access to the granted (server) App.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsAddedSinceVersion: 3
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ref` - Grant URI

		**SCIM++ Properties:**
		* idcsAddedSinceVersion: 3
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - Grant identifier

		**SCIM++ Properties:**
		* caseExact: true
		* idcsAddedSinceVersion: 3
		* idcsSearchable: true
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
* `matching_rule` - Store as a string the matching-rule for this Dynamic Resource Group. This may match any number of Apps in this Domain, as well as matching any number of Oracle Cloud Infrastructure resources that are not in any Domain but that are in the Oracle Cloud Infrastructure Compartment that contains this Domain.

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: true
	* returned: request
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
* `urnietfparamsscimschemasoracleidcsextension_oci_tags` - Oracle Cloud Infrastructure Tags.
	* `defined_tags` - Oracle Cloud Infrastructure Defined Tags

		**Added In:** 2011192329

		**SCIM++ Properties:**
		* idcsCompositeKey: [namespace, key, value]
		* type: complex
		* idcsSearchable: true
		* required: false
		* mutability: readWrite
		* multiValued: true
		* returned: default
		* `key` - Oracle Cloud Infrastructure Tag key

			**Added In:** 2011192329

			**SCIM++ Properties:**
			* caseExact: false
			* type: string
			* required: true
			* mutability: readWrite
			* returned: default
			* idcsSearchable: true
			* uniqueness: none
		* `namespace` - Oracle Cloud Infrastructure Tag namespace

			**Added In:** 2011192329

			**SCIM++ Properties:**
			* caseExact: false
			* type: string
			* required: true
			* mutability: readWrite
			* returned: default
			* idcsSearchable: true
			* uniqueness: none
		* `value` - Oracle Cloud Infrastructure Tag value

			**Added In:** 2011192329

			**SCIM++ Properties:**
			* caseExact: false
			* required: true
			* idcsReturnEmptyWhenNull: true
			* mutability: readWrite
			* returned: default
			* type: string
			* idcsSearchable: true
			* uniqueness: none
	* `freeform_tags` - Oracle Cloud Infrastructure Freeform Tags

		**Added In:** 2011192329

		**SCIM++ Properties:**
		* idcsCompositeKey: [key, value]
		* idcsSearchable: true
		* type: complex
		* required: false
		* mutability: readWrite
		* returned: default
		* multiValued: true
		* `key` - Oracle Cloud Infrastructure Tag key

			**Added In:** 2011192329

			**SCIM++ Properties:**
			* caseExact: false
			* type: string
			* required: true
			* mutability: readWrite
			* returned: default
			* idcsSearchable: true
			* uniqueness: none
		* `value` - Oracle Cloud Infrastructure Tag value

			**Added In:** 2011192329

			**SCIM++ Properties:**
			* caseExact: false
			* required: true
			* idcsReturnEmptyWhenNull: true
			* mutability: readWrite
			* returned: default
			* type: string
			* idcsSearchable: true
			* uniqueness: none
	* `tag_slug` - Oracle Cloud Infrastructure Tag slug

		**Added In:** 2011192329

		**SCIM++ Properties:**
		* type: binary
		* mutability: readOnly
		* returned: request

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Dynamic Resource Group
	* `update` - (Defaults to 20 minutes), when updating the Dynamic Resource Group
	* `delete` - (Defaults to 20 minutes), when destroying the Dynamic Resource Group


## Import

DynamicResourceGroups can be imported using the `id`, e.g.

```
$ terraform import oci_identity_domains_dynamic_resource_group.test_dynamic_resource_group "idcsEndpoint/{idcsEndpoint}/dynamicResourceGroups/{dynamicResourceGroupId}" 
```

