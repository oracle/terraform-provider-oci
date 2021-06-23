---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_dynamic_group"
sidebar_current: "docs-oci-resource-identity-dynamic_group"
description: |-
  Provides the Dynamic Group resource in Oracle Cloud Infrastructure Identity service
---

# oci_identity_dynamic_group
This resource provides the Dynamic Group resource in Oracle Cloud Infrastructure Identity service.

Creates a new dynamic group in your tenancy.

You must specify your tenancy's OCID as the compartment ID in the request object (remember that the tenancy
is simply the root compartment). Notice that IAM resources (users, groups, compartments, and some policies)
reside within the tenancy itself, unlike cloud resources such as compute instances, which typically
reside within compartments inside the tenancy. For information about OCIDs, see
[Resource Identifiers](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).

You must also specify a *name* for the dynamic group, which must be unique across all dynamic groups in your
tenancy, and cannot be changed. Note that this name has to be also unique across all groups in your tenancy.
You can use this name or the OCID when writing policies that apply to the dynamic group. For more information
about policies, see [How Policies Work](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/policies.htm).

You must also specify a *description* for the dynamic group (although it can be an empty string). It does not
have to be unique, and you can change it anytime with [UpdateDynamicGroup](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/DynamicGroup/UpdateDynamicGroup).


## Example Usage

```hcl
resource "oci_identity_dynamic_group" "test_dynamic_group" {
	#Required
	compartment_id = var.tenancy_ocid
	description = var.dynamic_group_description
	matching_rule = var.dynamic_group_matching_rule
	name = var.dynamic_group_name

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the tenancy containing the group.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Required) (Updatable) The description you assign to the group during creation. Does not have to be unique, and it's changeable.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `matching_rule` - (Required) (Updatable) The matching rule to dynamically match an instance certificate to this dynamic group. For rule syntax, see [Managing Dynamic Groups](https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingdynamicgroups.htm). 
* `name` - (Required) The name you assign to the group during creation. The name must be unique across all groups in the tenancy and cannot be changed. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the tenancy containing the group.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description you assign to the group. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the group.
* `inactive_state` - The detailed status of INACTIVE lifecycleState.
* `matching_rule` - A rule string that defines which instance certificates will be matched. For syntax, see [Managing Dynamic Groups](https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingdynamicgroups.htm). 
* `name` - The name you assign to the group during creation. The name must be unique across all groups in the tenancy and cannot be changed. 
* `state` - The group's current state.
* `time_created` - Date and time the group was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Dynamic Group
	* `update` - (Defaults to 20 minutes), when updating the Dynamic Group
	* `delete` - (Defaults to 20 minutes), when destroying the Dynamic Group


## Import

DynamicGroups can be imported using the `id`, e.g.

```
$ terraform import oci_identity_dynamic_group.test_dynamic_group "id"
```

