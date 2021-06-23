---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_user_group_membership"
sidebar_current: "docs-oci-resource-identity-user_group_membership"
description: |-
  Provides the User Group Membership resource in Oracle Cloud Infrastructure Identity service
---

# oci_identity_user_group_membership
This resource provides the User Group Membership resource in Oracle Cloud Infrastructure Identity service.

Adds the specified user to the specified group and returns a `UserGroupMembership` object with its own OCID.


## Example Usage

```hcl
resource "oci_identity_user_group_membership" "test_user_group_membership" {
	#Required
	group_id = oci_identity_group.test_group.id
	user_id = oci_identity_user.test_user.id
}
```

## Argument Reference

The following arguments are supported:

* `group_id` - (Required) The OCID of the group.
* `user_id` - (Required) The OCID of the user.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the tenancy containing the user, group, and membership object.
* `group_id` - The OCID of the group.
* `id` - The OCID of the membership.
* `inactive_state` - The detailed status of INACTIVE lifecycleState.
* `state` - The membership's current state.
* `time_created` - Date and time the membership was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `user_id` - The OCID of the user.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the User Group Membership
	* `update` - (Defaults to 20 minutes), when updating the User Group Membership
	* `delete` - (Defaults to 20 minutes), when destroying the User Group Membership


## Import

UserGroupMemberships can be imported using the `id`, e.g.

```
$ terraform import oci_identity_user_group_membership.test_user_group_membership "id"
```

