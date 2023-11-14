---
subcategory: "Identity Domains"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_domains_groups"
sidebar_current: "docs-oci-datasource-identity_domains-groups"
description: |-
  Provides the list of Groups in Oracle Cloud Infrastructure Identity Domains service
---

# Data Source: oci_identity_domains_groups
This data source provides the list of Groups in Oracle Cloud Infrastructure Identity Domains service.

Search for groups. <b>Important:</b> The Group SEARCH and GET operations on users and members will throw an exception if the response has more than 10,000 members. To avoid the exception, use the pagination filter to GET or SEARCH group members.

## Example Usage

```hcl
data "oci_identity_domains_groups" "test_groups" {
	#Required
	idcs_endpoint = data.oci_identity_domain.test_domain.url

	#Optional
	group_count = var.group_group_count
	group_filter = var.group_group_filter
	attribute_sets = []
	attributes = ""
	authorization = var.group_authorization
	resource_type_schema_version = var.group_resource_type_schema_version
	start_index = var.group_start_index
}
```

## Argument Reference

The following arguments are supported:

* `group_count` - (Optional) OPTIONAL. An integer that indicates the desired maximum number of query results per page. 1000 is the largest value that you can use. See the Pagination section of the System for Cross-Domain Identity Management Protocol specification for more information. (Section 3.4.2.4).
* `group_filter` - (Optional) OPTIONAL. The filter string that is used to request a subset of resources. The filter string MUST be a valid filter expression. See the Filtering section of the SCIM specification for more information (Section 3.4.2.2). The string should contain at least one condition that each item must match in order to be returned in the search results. Each condition specifies an attribute, an operator, and a value. Conditions within a filter can be connected by logical operators (such as AND and OR). Sets of conditions can be grouped together using parentheses.
* `attribute_sets` - (Optional) A multi-valued list of strings indicating the return type of attribute definition. The specified set of attributes can be fetched by the return type of the attribute. One or more values can be given together to fetch more than one group of attributes. If 'attributes' query parameter is also available, union of the two is fetched. Valid values - all, always, never, request, default. Values are case-insensitive.
* `attributes` - (Optional) A comma-delimited string that specifies the names of resource attributes that should be returned in the response. By default, a response that contains resource attributes contains only attributes that are defined in the schema for that resource type as returned=always or returned=default. An attribute that is defined as returned=request is returned in a response only if the request specifies its name in the value of this query parameter. If a request specifies this query parameter, the response contains the attributes that this query parameter specifies, as well as any attribute that is defined as returned=always.
* `authorization` - (Optional) The Authorization field value consists of credentials containing the authentication information of the user agent for the realm of the resource being requested.
* `idcs_endpoint` - (Required) The basic endpoint for the identity domain
* `resource_type_schema_version` - (Optional) An endpoint-specific schema version number to use in the Request. Allowed version values are Earliest Version or Latest Version as specified in each REST API endpoint description, or any sequential number inbetween. All schema attributes/body parameters are a part of version 1. After version 1, any attributes added or deprecated will be tagged with the version that they were added to or deprecated in. If no version is provided, the latest schema version is returned.
* `start_index` - (Optional) OPTIONAL. An integer that indicates the 1-based index of the first query result. See the Pagination section of the SCIM specification for more information. (Section 3.4.2.4). The number of results pages to return. The first page is 1. Specify 2 to access the second page of results, and so on.


## Attributes Reference

The following attributes are exported:

* `groups` - The list of groups.

### Group Reference

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
* `display_name` - The Group display name.

	**SCIM++ Properties:**
	* caseExact: false
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
* `external_id` - An identifier for the Resource as defined by the Service Consumer. The externalId may simplify identification of the Resource between Service Consumer and Service Provider by allowing the Consumer to refer to the Resource with its own identifier, obviating the need to store a local mapping between the local identifier of the Resource and the identifier used by the Service Provider. Each Resource MAY include a non-empty externalId value. The value of the externalId attribute is always issued by the Service Consumer and can never be specified by the Service Provider. The Service Provider MUST always interpret the externalId as scoped to the Service Consumer's tenant.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsSearchable: true
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
* `members` - The group members. <b>Important:</b> When requesting group members, a maximum of 10,000 members can be returned in a single request. If the response contains more than 10,000 members, the request will fail. Use 'startIndex' and 'count' to return members in pages instead of in a single response, for example: #attributes=members[startIndex=1%26count=10]. This REST API is SCIM compliant.

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCompositeKey: [value]
	* idcsCsvAttributeNameMappings: [[columnHeaderName:User Members, mapsTo:members[User].value, multiValueDelimiter:;]]
	* idcsSearchable: true
	* multiValued: true
	* mutability: readWrite
	* required: false
	* returned: request
	* idcsPaginateResponse: true
	* type: complex
	* uniqueness: none
	* `date_added` - The date and time that the member was added to the group.

		**Added In:** 2012271618

		**SCIM++ Properties:**
		* caseExact: false
		* multiValued: false
		* mutability: readOnly
		* returned: default
		* type: dateTime
		* uniqueness: none
	* `display` - The member's display name.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `membership_ocid` - The membership OCID.

		**Added In:** 2102181953

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `name` - The member's name.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ocid` - The OCID of the member of this group.

		**Added In:** 2012271618

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: always
		* type: string
		* uniqueness: none
	* `ref` - The URI that corresponds to the member Resource of this group.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `type` - Indicates the type of resource, for example, User or Group.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* idcsDefaultValue: User
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - The ID of the member of this Group

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: true
		* returned: always
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
* `non_unique_display_name` - A human readable name for the group as defined by the Service Consumer.

	**Added In:** 2011192329

	**SCIM++ Properties:**
	* caseExact: false
	* idcsCsvAttributeName: Non-Unique Display Name
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: always
	* type: string
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
* `urnietfparamsscimschemasoracleidcsextensiondbcs_group` - Schema for Database Service  Resource
	* `domain_level_schema` - DBCS Domain-level schema-name.  This attribute refers implicitly to a value of 'domainLevelSchemaNames' for a particular DB Domain.

		**Added In:** 18.2.4

		**SCIM++ Properties:**
		* idcsSearchable: false
		* idcsSensitive: none
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `domain_level_schema_names` - DBCS Domain-level schema-names. Each value is specific to a DB Domain.

		**Added In:** 18.2.4

		**SCIM++ Properties:**
		* idcsCompositeKey: [domainName, schemaName]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* `domain_name` - DBCS Domain Name

			**Added In:** 18.2.4

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `schema_name` - The DBCS schema-name granted to this group in the DB domain that 'domainName' specifies.

			**Added In:** 18.2.4

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
	* `instance_level_schema` - DBCS instance-level schema-name. This attribute refers implicitly to a value of 'instanceLevelSchemaNames' for a particular DB Instance.

		**Added In:** 18.2.4

		**SCIM++ Properties:**
		* idcsSearchable: false
		* idcsSensitive: none
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `instance_level_schema_names` - DBCS instance-level schema-names. Each schema-name is specific to a DB Instance.

		**Added In:** 18.2.4

		**SCIM++ Properties:**
		* idcsCompositeKey: [dbInstanceId, schemaName]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* `db_instance_id` - App Id of DBCS App instance

			**Added In:** 18.2.4

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `schema_name` - The DBCS schema-name granted to this Group for the DB instance that 'dbInstanceId' specifies.

			**Added In:** 18.2.4

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensiondynamic_group` - Dynamic Group
	* `membership_rule` - Membership rule

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `membership_type` - Membership type

		**Added In:** 19.2.1

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: always
		* type: string
		* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensiongroup_group` - Oracle Identity Cloud Service Group
	* `app_roles` - A list of appRoles that the user belongs to, either thorough direct membership, nested groups, or dynamically calculated

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
			* returned: default
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
		* `type` - A label indicating the attribute's function; e.g., 'direct' or 'indirect'.

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: request
			* type: string
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
	* `creation_mechanism` - Source from which this group got created.

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCsvAttributeNameMappings: [[defaultValue:import]]
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `description` - Group description

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCsvAttributeName: Description
		* idcsCsvAttributeNameMappings: [[columnHeaderName:Description]]
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `grants` - Grants assigned to group

		**SCIM++ Properties:**
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
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - Grant URI

			**SCIM++ Properties:**
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
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
	* `owners` - Group owners

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCompositeKey: [value, type]
		* idcsSearchable: true
		* multiValued: true
		* mutability: readWrite
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `display` - Owner display name

			**SCIM++ Properties:**
			* caseExact: false
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - The URI that corresponds to the owning Resource of this Group

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `type` - Indicates the type of resource--for example, User or Group

			**SCIM++ Properties:**
			* caseExact: true
			* idcsDefaultValue: User
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `value` - ID of the owner of this Group

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readWrite
			* required: true
			* returned: always
			* type: string
			* uniqueness: none
	* `password_policy` - Password Policy associated with this Group.

		**Added In:** 20.1.3

		**SCIM++ Properties:**
		* caseExact: false
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `name` - PasswordPolicy Name

			**Added In:** 20.1.3

			**SCIM++ Properties:**
			* caseExact: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `priority` - PasswordPolicy priority

			**Added In:** 20.1.3

			**SCIM++ Properties:**
			* caseExact: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: integer
			* uniqueness: none
		* `ref` - PasswordPolicy URI

			**Added In:** 20.1.3

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `value` - The ID of the PasswordPolicy.

			**Added In:** 20.1.3

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: true
			* returned: always
			* type: string
			* uniqueness: none
	* `synced_from_app` - The entity that created this Group.

		**Added In:** 18.4.2

		**SCIM++ Properties:**
		* idcsCompositeKey: [value]
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: complex
		* uniqueness: none
		* `display` - App Display Name

			**Added In:** 18.4.2

			**SCIM++ Properties:**
			* caseExact: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - App URI

			**Added In:** 18.4.2

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `type` - The type of the entity that created this Group.

			**Added In:** 18.4.2

			**SCIM++ Properties:**
			* idcsDefaultValue: App
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `value` - The ID of the App.

			**Added In:** 18.4.2

			**SCIM++ Properties:**
			* caseExact: true
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
* `urnietfparamsscimschemasoracleidcsextensionposix_group` - POSIX Group extension
	* `gid_number` - Integer uniquely identifying a group in a POSIX administrative domain

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: integer
		* uniqueness: server
* `urnietfparamsscimschemasoracleidcsextensionrequestable_group` - Requestable Group
	* `requestable` - Flag controlling whether group membership can be request by user through self service console.

		**Added In:** 17.3.4

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* idcsCsvAttributeNameMappings: [[columnHeaderName:Requestable, mapsTo:requestable]]
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: request
		* type: boolean
		* uniqueness: none

