---
subcategory: "Identity Domains"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_domains_grant"
sidebar_current: "docs-oci-resource-identity_domains-grant"
description: |-
  Provides the Grant resource in Oracle Cloud Infrastructure Identity Domains service
---

# oci_identity_domains_grant
This resource provides the Grant resource in Oracle Cloud Infrastructure Identity Domains service.

Add a Grantee to an AppRole

## Example Usage

```hcl
resource "oci_identity_domains_grant" "test_grant" {
	#Required
	grant_mechanism = var.grant_grant_mechanism
	grantee {
		#Required
		type = var.grant_grantee_type
		value = var.grant_grantee_value
	}
	idcs_endpoint = data.oci_identity_domain.test_domain.url
	schemas = ["urn:ietf:params:scim:schemas:oracle:idcs:Grant"]

	#Optional
	app {
		#Required
		value = var.grant_app_value
	}
	app_entitlement_collection {
		#Required
		value = var.grant_app_entitlement_collection_value
	}
	attribute_sets = ["all"]
	attributes = ""
	authorization = var.grant_authorization
	entitlement {
		#Required
		attribute_name = "appRoles"
		attribute_value = var.grant_entitlement_attribute_value
	}
	granted_attribute_values_json = var.grant_granted_attribute_values_json
	id = var.grant_id
	ocid = var.grant_ocid
	resource_type_schema_version = var.grant_resource_type_schema_version
	tags {
		#Required
		key = var.grant_tags_key
		value = var.grant_tags_value
	}
}
```

## Argument Reference

The following arguments are supported:

* `app` - (Optional) Application that is being granted. Each Grant must grant either an App or an App-Entitlement-Collection.

	**SCIM++ Properties:**
	* idcsCsvAttributeNameMappings: [[columnHeaderName:App Name, defaultValue:jobParam:appDisplayName, mapsTo:app.value]]
	* idcsSearchable: true
	* multiValued: false
	* mutability: immutable
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `display` - (Optional) (Updatable) Application display name

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `ref` - (Optional) (Updatable) Application URI

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - (Required) Application identifier

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `app_entitlement_collection` - (Optional) Application-Entitlement-Collection that is being granted. Each Grant must grant either an App or an App-Entitlement-Collection.

	**Added In:** 18.2.4

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: immutable
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `ref` - (Optional) (Updatable) Application Entitlement Collection URI

		**Added In:** 18.2.4

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - (Required) Application Entitlement Collection identifier

		**Added In:** 18.2.4

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
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
* `composite_key` - (Optional) (Updatable) Unique key of grant, composed by combining a subset of app, entitlement, grantee, grantor and grantMechanism.  Used to prevent duplicate Grants.

	**Added In:** 18.1.2

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: request
	* type: string
	* uniqueness: server
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
* `entitlement` - (Optional) The entitlement or privilege that is being granted

	**SCIM++ Properties:**
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Entitlement Value, csvColumnForResolvingResourceType:Entitlement Name, mapsTo:entitlement.attributeValue, referencedResourceTypeUniqueAttributeNameMappings:[[mapsFromColumnName:Entitlement Value, resourceTypeAttributeName:displayName], [mapsFromColumnName:App Name, resourceTypeAttributeName:app.display]], resolveValueUsingResourceType:[[resolveBy:AppRole, valueToBeResolved:appRoles]]], [columnHeaderName:Entitlement Name, defaultValue:appRoles, mapsTo:entitlement.attributeName]]
	* idcsSearchable: true
	* multiValued: false
	* mutability: immutable
	* required: false
	* returned: default
	* type: complex
	* `attribute_name` - (Required) The name of the attribute whose value (specified by attributeValue) confers privilege within the service-instance (specified by app).

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `attribute_value` - (Required) The value of the attribute (specified by attributeName) that confers privilege within the service-instance (specified by app).  If attributeName is 'appRoles', then attributeValue is the ID of the AppRole.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsCsvAttributeName: Display Name
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `grant_mechanism` - (Required) Each value of grantMechanism indicates how (or by what component) some App (or App-Entitlement) was granted. A customer or the UI should use only grantMechanism values that start with 'ADMINISTRATOR':
	* 'ADMINISTRATOR_TO_USER' is for a direct grant to a specific User.
	* 'ADMINISTRATOR_TO_GROUP' is for a grant to a specific Group, which results in indirect grants to Users who are members of that Group.
	* 'ADMINISTRATOR_TO_APP' is for a grant to a specific App.  The grantee (client) App gains access to the granted (server) App.

	**SCIM++ Properties:**
	* caseExact: true
	* idcsCsvAttributeNameMappings: [[defaultValue:IMPORT_GRANTS]]
	* idcsSearchable: true
	* multiValued: false
	* mutability: immutable
	* required: true
	* returned: default
	* type: string
	* uniqueness: none
* `granted_attribute_values_json` - (Optional) Store granted attribute-values as a string in Javascript Object Notation (JSON) format.

	**Added In:** 18.3.4

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `grantee` - (Required) Grantee beneficiary. The grantee may be a User, Group, App or DynamicResourceGroup.

	**SCIM++ Properties:**
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Grantee Name, csvColumnForResolvingResourceType:Grantee Type, mapsTo:grantee.value], [columnHeaderName:Grantee Type, mapsTo:grantee.type]]
	* idcsSearchable: true
	* multiValued: false
	* mutability: immutable
	* required: true
	* returned: default
	* type: complex
	* uniqueness: none
	* `display` - (Optional) (Updatable) Grantee display name

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `ref` - (Optional) (Updatable) Grantee URI

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `type` - (Required) Grantee resource type. Allowed values are User, Group, App and DynamicResourceGroup.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsCsvAttributeName: Member Type
		* idcsDefaultValue: User
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - (Required) Grantee identifier

		**SCIM++ Properties:**
		* caseExact: true
		* idcsCsvAttributeName: Member
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `grantor` - (Optional) (Updatable) User conferring the grant to the beneficiary

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* idcsIgnoreReadOnlyAndImmutableRefAttrsDuringForceDelete: true
	* required: false
	* returned: default
	* type: complex
	* `display` - (Optional) (Updatable) Grantor display name

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `ref` - (Optional) (Updatable) Grantor URI

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `type` - (Optional) (Updatable) Resource type of the grantor. Allowed values are User and App.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsDefaultValue: User
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - (Optional) (Updatable) Grantor user identifier

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: string
		* uniqueness: none
* `id` - (Optional) Unique identifier for the SCIM Resource as defined by the Service Provider. Each representation of the Resource MUST include a non-empty id value. This identifier MUST be unique across the Service Provider's entire set of Resources. It MUST be a stable, non-reassignable identifier that does not change when the same Resource is returned in subsequent requests. The value of the id attribute is always issued by the Service Provider and MUST never be specified by the Service Consumer. bulkId: is a reserved keyword and MUST NOT be used in the unique identifier.

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
* `is_fulfilled` - (Optional) (Updatable) If true, this Grant has been fulfilled successfully.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
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

* `app` - Application that is being granted. Each Grant must grant either an App or an App-Entitlement-Collection.

	**SCIM++ Properties:**
	* idcsCsvAttributeNameMappings: [[columnHeaderName:App Name, defaultValue:jobParam:appDisplayName, mapsTo:app.value]]
	* idcsSearchable: true
	* multiValued: false
	* mutability: immutable
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `display` - Application display name

		**SCIM++ Properties:**
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
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
	* `value` - Application identifier

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `app_entitlement_collection` - Application-Entitlement-Collection that is being granted. Each Grant must grant either an App or an App-Entitlement-Collection.

	**Added In:** 18.2.4

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: immutable
	* required: false
	* returned: default
	* type: complex
	* uniqueness: none
	* `ref` - Application Entitlement Collection URI

		**Added In:** 18.2.4

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `value` - Application Entitlement Collection identifier

		**Added In:** 18.2.4

		**SCIM++ Properties:**
		* caseExact: true
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
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
* `composite_key` - Unique key of grant, composed by combining a subset of app, entitlement, grantee, grantor and grantMechanism.  Used to prevent duplicate Grants.

	**Added In:** 18.1.2

	**SCIM++ Properties:**
	* caseExact: true
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* required: false
	* returned: request
	* type: string
	* uniqueness: server
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
* `entitlement` - The entitlement or privilege that is being granted

	**SCIM++ Properties:**
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Entitlement Value, csvColumnForResolvingResourceType:Entitlement Name, mapsTo:entitlement.attributeValue, referencedResourceTypeUniqueAttributeNameMappings:[[mapsFromColumnName:Entitlement Value, resourceTypeAttributeName:displayName], [mapsFromColumnName:App Name, resourceTypeAttributeName:app.display]], resolveValueUsingResourceType:[[resolveBy:AppRole, valueToBeResolved:appRoles]]], [columnHeaderName:Entitlement Name, defaultValue:appRoles, mapsTo:entitlement.attributeName]]
	* idcsSearchable: true
	* multiValued: false
	* mutability: immutable
	* required: false
	* returned: default
	* type: complex
	* `attribute_name` - The name of the attribute whose value (specified by attributeValue) confers privilege within the service-instance (specified by app).

		**SCIM++ Properties:**
		* caseExact: false
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `attribute_value` - The value of the attribute (specified by attributeName) that confers privilege within the service-instance (specified by app).  If attributeName is 'appRoles', then attributeValue is the ID of the AppRole.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsCsvAttributeName: Display Name
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `grant_mechanism` - Each value of grantMechanism indicates how (or by what component) some App (or App-Entitlement) was granted. A customer or the UI should use only grantMechanism values that start with 'ADMINISTRATOR':
	* 'ADMINISTRATOR_TO_USER' is for a direct grant to a specific User.
	* 'ADMINISTRATOR_TO_GROUP' is for a grant to a specific Group, which results in indirect grants to Users who are members of that Group.
	* 'ADMINISTRATOR_TO_APP' is for a grant to a specific App.  The grantee (client) App gains access to the granted (server) App.

	**SCIM++ Properties:**
	* caseExact: true
	* idcsCsvAttributeNameMappings: [[defaultValue:IMPORT_GRANTS]]
	* idcsSearchable: true
	* multiValued: false
	* mutability: immutable
	* required: true
	* returned: default
	* type: string
	* uniqueness: none
* `granted_attribute_values_json` - Store granted attribute-values as a string in Javascript Object Notation (JSON) format.

	**Added In:** 18.3.4

	**SCIM++ Properties:**
	* idcsSearchable: false
	* multiValued: false
	* mutability: readWrite
	* required: false
	* returned: default
	* type: string
	* uniqueness: none
* `grantee` - Grantee beneficiary. The grantee may be a User, Group, App or DynamicResourceGroup.

	**SCIM++ Properties:**
	* idcsCsvAttributeNameMappings: [[columnHeaderName:Grantee Name, csvColumnForResolvingResourceType:Grantee Type, mapsTo:grantee.value], [columnHeaderName:Grantee Type, mapsTo:grantee.type]]
	* idcsSearchable: true
	* multiValued: false
	* mutability: immutable
	* required: true
	* returned: default
	* type: complex
	* uniqueness: none
	* `display` - Grantee display name

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `ref` - Grantee URI

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `type` - Grantee resource type. Allowed values are User, Group, App and DynamicResourceGroup.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsCsvAttributeName: Member Type
		* idcsDefaultValue: User
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - Grantee identifier

		**SCIM++ Properties:**
		* caseExact: true
		* idcsCsvAttributeName: Member
		* idcsSearchable: true
		* multiValued: false
		* mutability: immutable
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
* `grantor` - User conferring the grant to the beneficiary

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
	* idcsIgnoreReadOnlyAndImmutableRefAttrsDuringForceDelete: true
	* required: false
	* returned: default
	* type: complex
	* `display` - Grantor display name

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: request
		* type: string
		* uniqueness: none
	* `ref` - Grantor URI

		**SCIM++ Properties:**
		* idcsSearchable: false
		* multiValued: false
		* mutability: readOnly
		* required: false
		* returned: default
		* type: reference
		* uniqueness: none
	* `type` - Resource type of the grantor. Allowed values are User and App.

		**SCIM++ Properties:**
		* caseExact: true
		* idcsDefaultValue: User
		* idcsSearchable: true
		* multiValued: false
		* mutability: readOnly
		* required: true
		* returned: default
		* type: string
		* uniqueness: none
	* `value` - Grantor user identifier

		**SCIM++ Properties:**
		* caseExact: true
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
* `is_fulfilled` - If true, this Grant has been fulfilled successfully.

	**SCIM++ Properties:**
	* idcsSearchable: true
	* multiValued: false
	* mutability: readOnly
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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Grant
	* `update` - (Defaults to 20 minutes), when updating the Grant
	* `delete` - (Defaults to 20 minutes), when destroying the Grant


## Import

Grants can be imported using the `id`, e.g.

```
$ terraform import oci_identity_domains_grant.test_grant "idcsEndpoint/{idcsEndpoint}/grants/{grantId}" 
```

