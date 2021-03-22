---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_tag"
sidebar_current: "docs-oci-resource-identity-tag"
description: |-
  Provides the Tag resource in Oracle Cloud Infrastructure Identity service
---

# oci_identity_tag
This resource provides the Tag resource in Oracle Cloud Infrastructure Identity service.

Creates a new tag in the specified tag namespace.

The tag requires either the OCID or the name of the tag namespace that will contain this
tag definition.

You must specify a *name* for the tag, which must be unique across all tags in the tag namespace
and cannot be changed. The name can contain any ASCII character except the space (_) or period (.) characters.
Names are case insensitive. That means, for example, "myTag" and "mytag" are not allowed in the same namespace.
If you specify a name that's already in use in the tag namespace, a 409 error is returned.

The tag must have a *description*. It does not have to be unique, and you can change it with
[UpdateTag](https://docs.cloud.oracle.com/iaas/api/#/en/identity/latest/Tag/UpdateTag).

The tag must have a value type, which is specified with a validator. Tags can use either a
static value or a list of possible values. Static values are entered by a user applying the tag
to a resource. Lists are created by you and the user must apply a value from the list. Lists
are validiated.

* If no `validator` is set, the user applying the tag to a resource can type in a static
value or leave the tag value empty.
* If a `validator` is set, the user applying the tag to a resource must select from a list
of values that you supply with [EnumTagDefinitionValidator](https://docs.cloud.oracle.com/iaas/api/#/en/identity/latest/datatypes/EnumTagDefinitionValidator).


## Example Usage

```hcl
resource "oci_identity_tag" "test_tag" {
	#Required
	description = var.tag_description
	name = var.tag_name
	tag_namespace_id = oci_identity_tag_namespace.test_tag_namespace.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	is_cost_tracking = var.tag_is_cost_tracking
	validator {
		#Required
		validator_type = var.tag_validator_validator_type
		values = var.tag_validator_values
	}
	is_retired = false
}
```

## Argument Reference

The following arguments are supported:

* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Required) (Updatable) The description you assign to the tag during creation.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `is_cost_tracking` - (Optional) (Updatable) Indicates whether the tag is enabled for cost tracking. 
* `name` - (Required) The name you assign to the tag during creation. This is the tag key definition. The name must be unique within the tag namespace and cannot be changed. 
* `tag_namespace_id` - (Required) The OCID of the tag namespace. 
* `validator` - (Optional) (Updatable) Validates a definedTag value. Each validator performs validation steps in addition to the standard validation for definedTag values. For more information, see [Limits on Tags](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/taggingoverview.htm#Limits).

	If you define a validator after a value has been set for a defined tag, then any updates that attempt to change the value must pass the additional validation defined by the current rule. Previously set values (even those that would fail the current validation) are not updated. You can still update other attributes to resources that contain a non-valid defined tag.

	To clear the validator call UpdateTag with [DefaultTagDefinitionValidator](https://docs.cloud.oracle.com/iaas/api/#/en/identity/latest/datatypes/DefaultTagDefinitionValidator). 
	* `validator_type` - (Required) (Updatable) Specifies the type of validation: a static value (no validation) or a list. 
	* `values` - (Applicable when validator_type=ENUM) (Updatable) The list of allowed values for a definedTag value. 
* `is_retired` - (Optional) (Updatable) Indicates whether the tag is retired. See [Retiring Key Definitions and Namespace Definitions](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/taggingoverview.htm#Retiring).


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description you assign to the tag.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the tag definition.
* `is_cost_tracking` - Indicates whether the tag is enabled for cost tracking. 
* `is_retired` - Indicates whether the tag is retired. See [Retiring Key Definitions and Namespace Definitions](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/taggingoverview.htm#Retiring). 
* `name` - The name assigned to the tag during creation. This is the tag key definition. The name must be unique within the tag namespace and cannot be changed. 
* `state` - The tag's current state. After creating a tag, make sure its `lifecycleState` is ACTIVE before using it. After retiring a tag, make sure its `lifecycleState` is INACTIVE before using it. If you delete a tag, you cannot delete another tag until the deleted tag's `lifecycleState` changes from DELETING to DELETED.
* `tag_namespace_id` - The OCID of the namespace that contains the tag definition.
* `time_created` - Date and time the tag was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `validator` - Validates a definedTag value. Each validator performs validation steps in addition to the standard validation for definedTag values. For more information, see [Limits on Tags](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/taggingoverview.htm#Limits).

	If you define a validator after a value has been set for a defined tag, then any updates that attempt to change the value must pass the additional validation defined by the current rule. Previously set values (even those that would fail the current validation) are not updated. You can still update other attributes to resources that contain a non-valid defined tag.

	To clear the validator call UpdateTag with [DefaultTagDefinitionValidator](https://docs.cloud.oracle.com/iaas/api/#/en/identity/latest/datatypes/DefaultTagDefinitionValidator). 
	* `validator_type` - Specifies the type of validation: a static value (no validation) or a list. 
	* `values` - The list of allowed values for a definedTag value. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 15 minutes), when creating the Tag
	* `update` - (Defaults to 15 minutes), when updating the Tag
	* `delete` - (Defaults to 12 hours), when destroying the Tag


## Import

Tags can be imported using the `tagNamespaceId` and `tagName`, e.g.

```
$ terraform import oci_identity_tag.test_tag "tagNamespaces/{tagNamespaceId}/tags/{tagName}" 
```

