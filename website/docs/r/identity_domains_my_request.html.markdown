---
subcategory: "Identity Domains"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_domains_my_request"
sidebar_current: "docs-oci-resource-identity_domains-my_request"
description: |-
  Provides the My Request resource in Oracle Cloud Infrastructure Identity Domains service
---

# oci_identity_domains_my_request
This resource provides the My Request resource in Oracle Cloud Infrastructure Identity Domains service.

Create a Request

** IMPORTANT **
In our latest release, the property `status` is changed to readonly. It will now be automatically handled by the system. Please remove any manual assignment to this property to use the latest version.

## Example Usage

```hcl
resource "oci_identity_domains_my_request" "test_my_request" {
	#Required
	idcs_endpoint = data.oci_identity_domain.test_domain.url
	justification = var.my_request_justification
	requesting {
		#Required
		type = var.my_request_requesting_type
		value = oci_identity_domains_group.group_to_request.id

		#Optional
		description = var.my_request_requesting_description
	}
	schemas = ["urn:ietf:params:scim:schemas:oracle:idcs:Request"]

	#Optional
	action = var.my_request_action
	approval_details {
	}
	attribute_sets = ["all"]
	attributes = ""
	authorization = var.my_request_authorization
	ocid = var.my_request_ocid
	requestor {
		#Required
		value = var.my_request_requestor_value
	}
	resource_type_schema_version = var.my_request_resource_type_schema_version
	tags {
		#Required
		key = var.my_request_tags_key
		value = var.my_request_tags_value
	}
}
```

## Argument Reference

The following arguments are supported:

* `action` - (Optional) Requestor can set action to CANCEL to cancel the request or to ESCALATE to escalate the request while the request status is IN_PROGRESS. Requestor can't escalate the request if canceling or escalation is in progress.

    **Added In:** 2307071836

    **SCIM++ Properties:**
    * caseExact: true
    * idcsSearchable: true
    * multiValued: false
    * mutability: readWrite
    * required: false
    * returned: default
    * type: string
    * uniqueness: none
* `approval_details` - (Optional) Approvals created for this request.

    **Added In:** 2307071836

    **SCIM++ Properties:**
    * idcsSearchable: false
    * multiValued: true
    * mutability: readOnly
    * returned: request
    * type: complex
    * uniqueness: none
    * `approval_type` - (Optional) (Updatable) Approval Type (Escalation or Regular)

        **Added In:** 2307071836

        **SCIM++ Properties:**
        * idcsSearchable: false
        * multiValued: false
        * returned: default
        * type: string
        * uniqueness: none
        * mutability: readOnly
    * `approver_display_name` - (Optional) (Updatable) Approver display name

        **Added In:** 2307071836

        **SCIM++ Properties:**
        * idcsSearchable: false
        * multiValued: false
        * returned: default
        * type: string
        * uniqueness: none
        * mutability: readOnly
    * `approver_id` - (Optional) (Updatable) Approver Id

        **Added In:** 2307071836

        **SCIM++ Properties:**
        * idcsSearchable: false
        * multiValued: false
        * returned: default
        * type: string
        * uniqueness: none
        * mutability: readOnly
    * `justification` - (Optional) (Updatable) Approval Justification

        **Added In:** 2307071836

        **SCIM++ Properties:**
        * multiValued: false
        * idcsSearchable: false
        * returned: default
        * type: string
        * uniqueness: none
        * mutability: readOnly
    * `order` - (Optional) (Updatable) Approval Order

        **Added In:** 2307071836

        **SCIM++ Properties:**
        * idcsSearchable: false
        * multiValued: false
        * returned: default
        * type: integer
        * uniqueness: none
        * mutability: readOnly
    * `status` - (Optional) (Updatable) Approval Status

        **Added In:** 2307071836

        **SCIM++ Properties:**
        * idcsSearchable: false
        * multiValued: false
        * returned: default
        * type: string
        * uniqueness: none
        * mutability: readOnly
    * `time_updated` - (Optional) (Updatable) Approval Update Time

        **Added In:** 2307071836

        **SCIM++ Properties:**
        * multiValued: false
        * idcsSearchable: false
        * returned: default
        * type: dateTime
        * uniqueness: none
        * mutability: readOnly
* `attribute_sets` - (Optional) A multi-valued list of strings indicating the return type of attribute definition. The specified set of attributes can be fetched by the return type of the attribute. One or more values can be given together to fetch more than one group of attributes. If 'attributes' query parameter is also available, union of the two is fetched. Valid values - all, always, never, request, default. Values are case-insensitive.
* `attributes` - (Optional) A comma-delimited string that specifies the names of resource attributes that should be returned in the response. By default, a response that contains resource attributes contains only attributes that are defined in the schema for that resource type as returned=always or returned=default. An attribute that is defined as returned=request is returned in a response only if the request specifies its name in the value of this query parameter. If a request specifies this query parameter, the response contains the attributes that this query parameter specifies, as well as any attribute that is defined as returned=always.
* `authorization` - (Optional) The Authorization field value consists of credentials containing the authentication information of the user agent for the realm of the resource being requested.
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
* `expires` - (Optional) (Updatable) Time by when Request expires

    **Added In:** 2307071836

    **SCIM++ Properties:**
    * idcsSearchable: true
    * multiValued: false
    * mutability: readOnly
    * required: false
    * returned: default
    * type: dateTime
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
    * `display` - (Optional) The displayName of the User or App who created this Resource

        **SCIM++ Properties:**
        * caseExact: true
        * idcsSearchable: false
        * multiValued: false
        * mutability: readOnly
        * required: false
        * returned: default
        * type: string
        * uniqueness: none
    * `ocid` - (Optional) The OCID of the SCIM resource that represents the User or App who created this Resource

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
    * `type` - (Optional) The type of resource, User or App, that created this Resource

        **SCIM++ Properties:**
        * caseExact: false
        * idcsSearchable: false
        * multiValued: false
        * mutability: readOnly
        * required: false
        * returned: default
        * type: string
        * uniqueness: none
    * `value` - (Required) The ID of the SCIM resource that represents the User or App who created this Resource

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
    * `display` - (Optional) The displayName of the User or App who modified this Resource

        **SCIM++ Properties:**
        * caseExact: true
        * idcsSearchable: false
        * multiValued: false
        * mutability: readOnly
        * required: false
        * returned: default
        * type: string
        * uniqueness: none
    * `ocid` - (Optional) The OCID of the SCIM resource that represents the User or App who modified this Resource

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
    * `type` - (Optional) The type of resource, User or App, that modified this Resource

        **SCIM++ Properties:**
        * caseExact: false
        * idcsSearchable: false
        * multiValued: false
        * mutability: readOnly
        * required: false
        * returned: default
        * type: string
        * uniqueness: none
    * `value` - (Required) The ID of the SCIM resource that represents the User or App who modified this Resource

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
* `justification` - (Required) justification

    **SCIM++ Properties:**
    * caseExact: true
    * idcsSearchable: true
    * multiValued: false
    * mutability: immutable
    * required: true
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
    * `created` - (Optional) The DateTime the Resource was added to the Service Provider

        **SCIM++ Properties:**
        * caseExact: false
        * idcsSearchable: true
        * multiValued: false
        * mutability: readOnly
        * required: false
        * returned: default
        * type: dateTime
        * uniqueness: none
    * `last_modified` - (Optional) The most recent DateTime that the details of this Resource were updated at the Service Provider. If this Resource has never been modified since its initial creation, the value MUST be the same as the value of created. The attribute MUST be a DateTime.

        **SCIM++ Properties:**
        * caseExact: false
        * idcsSearchable: true
        * multiValued: false
        * mutability: readOnly
        * required: false
        * returned: default
        * type: dateTime
        * uniqueness: none
    * `location` - (Optional) The URI of the Resource being returned. This value MUST be the same as the Location HTTP response header.

        **SCIM++ Properties:**
        * caseExact: false
        * idcsSearchable: false
        * multiValued: false
        * mutability: readOnly
        * required: false
        * returned: default
        * type: string
        * uniqueness: none
    * `resource_type` - (Optional) Name of the resource type of the resource--for example, Users or Groups

        **SCIM++ Properties:**
        * caseExact: false
        * idcsSearchable: false
        * multiValued: false
        * mutability: readOnly
        * required: false
        * returned: default
        * type: string
        * uniqueness: none
    * `version` - (Optional) The version of the Resource being returned. This value must be the same as the ETag HTTP response header.

        **SCIM++ Properties:**
        * caseExact: false
        * idcsSearchable: false
        * multiValued: false
        * mutability: readOnly
        * required: false
        * returned: default
        * type: string
        * uniqueness: none
* `ocid` - (Optional) Unique Oracle Cloud Infrastructure identifier for the SCIM Resource.

    **SCIM++ Properties:**
    * caseExact: true
    * idcsSearchable: true
    * multiValued: false
    * mutability: immutable
    * required: false
    * returned: default
    * type: string
    * uniqueness: global
* `requesting` - (Required) Requestable resource reference.

    **SCIM++ Properties:**
    * idcsSearchable: true
    * multiValued: false
    * mutability: immutable
    * required: true
    * returned: default
    * type: complex
    * uniqueness: none
    * `display` - (Optional) Resource display name

        **SCIM++ Properties:**
        * idcsSearchable: true
        * multiValued: false
        * mutability: readOnly
        * required: false
        * returned: default
        * type: string
        * uniqueness: none
    * `description` - (Optional) Resource description

        **Added In:** 2307071836

        **SCIM++ Properties:**
        * idcsSearchable: true
        * multiValued: false
        * mutability: readOnly
        * required: false
        * returned: default
        * type: string
        * uniqueness: none
    * `ref` - (Optional) (Updatable) Resource URI

        **SCIM++ Properties:**
        * idcsSearchable: false
        * multiValued: false
        * mutability: readOnly
        * required: false
        * returned: default
        * type: reference
        * uniqueness: none
    * `type` - (Required) Requestable type. Allowed values are Group and App.

        **SCIM++ Properties:**
        * caseExact: true
        * idcsCsvAttributeName: Requestable Type
        * idcsDefaultValue: Group
        * idcsSearchable: true
        * multiValued: false
        * mutability: immutable
        * required: true
        * returned: default
        * type: string
        * uniqueness: none
    * `value` - (Required) Resource identifier

        **SCIM++ Properties:**
        * caseExact: true
        * idcsCsvAttributeName: requesting_id
        * idcsSearchable: true
        * multiValued: false
        * mutability: immutable
        * required: true
        * returned: default
        * type: string
        * uniqueness: none
* `requestor` - (Optional) Requesting User

    **SCIM++ Properties:**
    * idcsSearchable: true
    * multiValued: false
    * mutability: immutable
    * required: false
    * returned: request
    * type: complex
    * uniqueness: none
    * `display` - (Optional) User display name

        **SCIM++ Properties:**
        * idcsSearchable: false
        * multiValued: false
        * mutability: readOnly
        * required: false
        * returned: default
        * type: string
        * uniqueness: none
    * `ref` - (Optional) (Updatable) User URI

        **SCIM++ Properties:**
        * idcsSearchable: false
        * multiValued: false
        * mutability: readOnly
        * required: false
        * returned: default
        * type: reference
        * uniqueness: none
    * `value` - (Required) User identifier

        **SCIM++ Properties:**
        * caseExact: true
        * idcsCsvAttributeName: requestor_id
        * idcsSearchable: true
        * multiValued: false
        * mutability: immutable
        * required: true
        * returned: default
        * type: string
        * uniqueness: none
* `resource_type_schema_version` - (Optional) An endpoint-specific schema version number to use in the Request. Allowed version values are Earliest Version or Latest Version as specified in each REST API endpoint description, or any sequential number inbetween. All schema attributes/body parameters are a part of version 1. After version 1, any attributes added or deprecated will be tagged with the version that they were added to or deprecated in. If no version is provided, the latest schema version is returned.
* `schemas` - (Required) REQUIRED. The schemas attribute is an array of Strings which allows introspection of the supported schema version for a SCIM representation as well any schema extensions supported by that representation. Each String value must be a unique URI. This specification defines URIs for User, Group, and a standard \"enterprise\" extension. All representations of SCIM schema MUST include a non-zero value array with value(s) of the URIs supported by that representation. Duplicate values MUST NOT be included. Value order is not specified and MUST not impact behavior.

    **SCIM++ Properties:**
    * caseExact: false
    * idcsSearchable: false
    * multiValued: true
    * mutability: readWrite
    * required: true
    * returned: default
    * type: string
    * uniqueness: none
* `status` - (Optional) (Updatable) status.

    **SCIM++ Properties:**
    * caseExact: true
    * idcsSearchable: true
    * multiValued: false
    * mutability: readOnly
    * required: false
    * returned: default
    * type: string
    * uniqueness: none
* `tags` - (Optional) A list of tags on this resource.

    **SCIM++ Properties:**
    * idcsCompositeKey: [key, value]
    * idcsSearchable: true
    * multiValued: true
    * mutability: readWrite
    * required: false
    * returned: request
    * type: complex
    * uniqueness: none
    * `key` - (Required) Key or name of the tag.

        **SCIM++ Properties:**
        * caseExact: false
        * idcsSearchable: true
        * multiValued: false
        * mutability: readWrite
        * required: true
        * returned: default
        * type: string
        * uniqueness: none
    * `value` - (Required) Value of the tag.

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


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `resources` - A multi-valued list of complex objects containing the requested resources. This MAY be a subset of the full set of resources if pagination is requested. REQUIRED if "totalResults" is non-zero.
	* `action` - Requestor can set action to CANCEL to cancel the request or to ESCALATE to escalate the request while the request status is IN_PROGRESS. Requestor can't escalate the request if canceling or escalation is in progress.

		**Added In:** 2307071836

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readWrite
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `approval_details` - Approvals created for this request.

		**Added In:** 2307071836

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: true
		* mutability: readOnly
		* returned: request
		* type: complex
		* uniqueness: none
		* `approval_type` - Approval Type (Escalation or Regular)

			**Added In:** 2307071836

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* returned: default
			* type: string
			* uniqueness: none
			* mutability: readOnly
		* `approver_display_name` - Approver display name

			**Added In:** 2307071836

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* returned: default
			* type: string
			* uniqueness: none
			* mutability: readOnly
		* `approver_id` - Approver Id

			**Added In:** 2307071836

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* returned: default
			* type: string
			* uniqueness: none
			* mutability: readOnly
		* `justification` - Approval Justification

			**Added In:** 2307071836

			**SCIM++ Properties:**
			* multiValued: false
			* idcsSearchable: false
			* returned: default
			* type: string
			* uniqueness: none
			* mutability: readOnly
		* `order` - Approval Order

			**Added In:** 2307071836

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* returned: default
			* type: integer
			* uniqueness: none
			* mutability: readOnly
		* `status` - Approval Status

			**Added In:** 2307071836

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* returned: default
			* type: string
			* uniqueness: none
			* mutability: readOnly
		* `time_updated` - Approval Update Time

			**Added In:** 2307071836

			**SCIM++ Properties:**
			* multiValued: false
			* idcsSearchable: false
			* returned: default
			* type: dateTime
			* uniqueness: none
			* mutability: readOnly
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
	* `expires` - Time by when Request expires

		**Added In:** 2307071836

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: dateTime
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
	* `justification` - justification

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: true
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
	* `requesting` - Requestable resource reference.

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: true
		* returned: default
		* type: complex
		* uniqueness: none
		* `display` - Resource display name

			**SCIM++ Properties:**
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `description` - Resource description

			**Added In:** 2307071836

			**SCIM++ Properties:**
			* idcsSearchable: true
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: string
			* uniqueness: none
		* `ref` - Resource URI

			**SCIM++ Properties:**
			* idcsSearchable: false
			* multiValued: false
			* mutability: readOnly
			* required: false
			* returned: default
			* type: reference
			* uniqueness: none
		* `type` - Requestable type. Allowed values are Group and App.

			**SCIM++ Properties:**
			* caseExact: true
			* idcsCsvAttributeName: Requestable Type
			* idcsDefaultValue: Group
			* idcsSearchable: true
			* multiValued: false
			* mutability: immutable
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
		* `value` - Resource identifier

			**SCIM++ Properties:**
			* caseExact: true
			* idcsCsvAttributeName: requesting_id
			* idcsSearchable: true
			* multiValued: false
			* mutability: immutable
			* required: true
			* returned: default
			* type: string
			* uniqueness: none
	* `requestor` - Requesting User

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: false
		* returned: request
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
		* `value` - User identifier

			**SCIM++ Properties:**
			* caseExact: true
			* idcsCsvAttributeName: requestor_id
			* idcsSearchable: true
			* multiValued: false
			* mutability: immutable
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
	* `status` - status

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
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
* `action` - Requestor can set action to CANCEL to cancel the request or to ESCALATE to escalate the request while the request status is IN_PROGRESS. Requestor can't escalate the request if canceling or escalation is in progress.

	**Added In:** 2307071836

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `approval_details` - Approvals created for this request.

	**Added In:** 2307071836

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: true
	* mutability: readOnly
	* returned: request
	* type: complex
	* uniqueness: none
	* `approval_type` - Approval Type (Escalation or Regular)

		**Added In:** 2307071836

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* returned: default
		* type: string
		* uniqueness: none
		* mutability: readOnly
	* `approver_display_name` - Approver display name

		**Added In:** 2307071836

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* returned: default
		* type: string
		* uniqueness: none
		* mutability: readOnly
	* `approver_id` - Approver Id

		**Added In:** 2307071836

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* returned: default
		* type: string
		* uniqueness: none
		* mutability: readOnly
	* `justification` - Approval Justification

		**Added In:** 2307071836

		**SCIM++ Properties:**
		* multiValued: false
		* idcsSearchable: false
		* returned: default
		* type: string
		* uniqueness: none
		* mutability: readOnly
	* `order` - Approval Order

		**Added In:** 2307071836

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* returned: default
		* type: integer
		* uniqueness: none
		* mutability: readOnly
	* `status` - Approval Status

		**Added In:** 2307071836

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* returned: default
		* type: string
		* uniqueness: none
		* mutability: readOnly
	* `time_updated` - Approval Update Time

		**Added In:** 2307071836

		**SCIM++ Properties:**
		* multiValued: false
		* idcsSearchable: false
		* returned: default
		* type: dateTime
		* uniqueness: none
		* mutability: readOnly
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
* `expires` - Time by when Request expires

	**Added In:** 2307071836

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: default
	* type: dateTime
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
* `items_per_page` - The number of resources returned in a list response page. REQUIRED when partial results returned due to pagination.
* `justification` - justification

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: immutable
	* required: true
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
* `requesting` - Requestable resource reference.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: immutable
	* required: true
	* returned: default
	* type: complex
	* uniqueness: none
	* `display` - Resource display name

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `description` - Resource description

		**Added In:** 2307071836

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
	* `ref` - Resource URI

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `type` - Requestable type. Allowed values are Group and App.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsCsvAttributeName: Requestable Type
		* idcsDefaultValue: Group
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - Resource identifier

		**SCIM++ Properties:**
		* caseExact: true
		* idcsCsvAttributeName: requesting_id
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `requestor` - Requesting User

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: immutable
	* required: false
	* returned: request
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
	* `value` - User identifier

		**SCIM++ Properties:**
		* caseExact: true
		* idcsCsvAttributeName: requestor_id
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
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
* `start_index` - The 1-based index of the first result in the current set of list results.  REQUIRED when partial results returned due to pagination.
* `status` - status

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
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
* `total_results` - The total number of results returned by the list or query operation.  The value may be larger than the number of resources returned such as when returning a single page of results where multiple pages are available. REQUIRED.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the My Request
	* `update` - (Defaults to 20 minutes), when updating the My Request
	* `delete` - (Defaults to 20 minutes), when destroying the My Request


## Import

Import is not supported for this resource.

