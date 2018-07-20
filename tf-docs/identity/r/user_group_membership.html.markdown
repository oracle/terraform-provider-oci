---
layout: "oci"
page_title: "OCI: oci_identity_user_group_membership"
sidebar_current: "docs-oci-resource-user_group_membership"
description: |-
Creates and manages an OCI UserGroupMembership
---

# oci_identity_user_group_membership
The `oci_identity_user_group_membership` resource creates and manages an OCI UserGroupMembership

Adds the specified user to the specified group and returns a `UserGroupMembership` object with its own OCID.

After you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using the
object, first make sure its `lifecycleState` has changed to ACTIVE.


## Example Usage

```hcl
resource "oci_identity_user_group_membership" "test_user_group_membership" {
	#Required
	group_id = "${oci_identity_group.test_group.id}"
	user_id = "${oci_identity_user.test_user.id}"
}
```

## Argument Reference

The following arguments are supported:

* `group_id` - (Required) The OCID of the group.
* `user_id` - (Required) The OCID of the user.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the tenancy containing the user, group, and membership object.
* `group_id` - The OCID of the group.
* `id` - The OCID of the membership.
* `inactive_state` - The detailed status of INACTIVE lifecycleState.
* `state` - The membership's current state.  After creating a membership object, make sure its `lifecycleState` changes from CREATING to ACTIVE before using it. 
* `time_created` - Date and time the membership was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `user_id` - The OCID of the user.

## Import

UserGroupMemberships can be imported using the `id`, e.g.

```
$ terraform import oci_identity_user_group_membership.test_user_group_membership "id"
```
