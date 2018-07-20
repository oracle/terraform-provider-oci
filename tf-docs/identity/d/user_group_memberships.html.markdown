---
layout: "oci"
page_title: "OCI: oci_identity_user_group_memberships"
sidebar_current: "docs-oci-datasource-user_group_memberships"
description: |-
Provides a list of UserGroupMemberships
---
# Data Source: oci_identity_user_group_memberships
The UserGroupMemberships data source allows access to the list of OCI user_group_memberships

Lists the `UserGroupMembership` objects in your tenancy. You must specify your tenancy's OCID
as the value for the compartment ID
(see [Where to Get the Tenancy's OCID and User's OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/apisigningkey.htm#five)).
You must also then filter the list in one of these ways:

- You can limit the results to just the memberships for a given user by specifying a `userId`.
- Similarly, you can limit the results to just the memberships for a given group by specifying a `groupId`.
- You can set both the `userId` and `groupId` to determine if the specified user is in the specified group.
If the answer is no, the response is an empty list.
- Although`userId` and `groupId` are not indvidually required, you must set one of them.


## Example Usage

```hcl
data "oci_identity_user_group_memberships" "test_user_group_memberships" {
	#Required
	compartment_id = "${var.tenancy_ocid}"

	#Optional
	group_id = "${oci_identity_group.test_group.id}"
	user_id = "${oci_identity_user.test_user.id}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment (remember that the tenancy is simply the root compartment). 
* `group_id` - (Optional) The OCID of the group.
* `user_id` - (Optional) The OCID of the user.


## Attributes Reference

The following attributes are exported:

* `memberships` - The list of memberships.

### UserGroupMembership Reference

The following attributes are exported:

* `compartment_id` - The OCID of the tenancy containing the user, group, and membership object.
* `group_id` - The OCID of the group.
* `id` - The OCID of the membership.
* `inactive_state` - The detailed status of INACTIVE lifecycleState.
* `state` - The membership's current state.  After creating a membership object, make sure its `lifecycleState` changes from CREATING to ACTIVE before using it. 
* `time_created` - Date and time the membership was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `user_id` - The OCID of the user.

