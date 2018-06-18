# oci_identity_user_group_membership

## UserGroupMembership Resource

### UserGroupMembership Reference

The following attributes are exported:

* `compartment_id` - The OCID of the tenancy containing the user, group, and membership object.
* `group_id` - The OCID of the group.
* `id` - The OCID of the membership.
* `inactive_state` - The detailed status of INACTIVE lifecycleState.
* `state` - The membership's current state.  After creating a membership object, make sure its `lifecycleState` changes from CREATING to ACTIVE before using it. 
* `time_created` - Date and time the membership was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `user_id` - The OCID of the user.



### Create Operation
Adds the specified user to the specified group and returns a `UserGroupMembership` object with its own OCID.

After you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using the
object, first make sure its `lifecycleState` has changed to ACTIVE.


The following arguments are supported:

* `group_id` - (Required) The OCID of the group.
* `user_id` - (Required) The OCID of the user.


### Update Operation


The following arguments support updates:
* NO arguments in this resource support updates

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_identity_user_group_membership" "test_user_group_membership" {
	#Required
	group_id = "${oci_identity_group.test_group.id}"
	user_id = "${oci_identity_user.test_user.id}"
}
```

# oci_identity_user_group_memberships

## UserGroupMembership DataSource

Gets a list of user_group_memberships.

### List Operation
Lists the `UserGroupMembership` objects in your tenancy. You must specify your tenancy's OCID
as the value for the compartment ID
(see [Where to Get the Tenancy's OCID and User's OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/apisigningkey.htm#five)).
You must also then filter the list in one of these ways:

- You can limit the results to just the memberships for a given user by specifying a `userId`.
- Similarly, you can limit the results to just the memberships for a given group by specifying a `groupId`.
- You can set both the `userId` and `groupId` to determine if the specified user is in the specified group.
If the answer is no, the response is an empty list.
- Although`userId` and `groupId` are not indvidually required, you must set one of them.

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment (remember that the tenancy is simply the root compartment). 
* `group_id` - (Optional) The OCID of the group.
* `user_id` - (Optional) The OCID of the user.


The following attributes are exported:

* `memberships` - The list of memberships.

### Example Usage

```hcl
data "oci_identity_user_group_memberships" "test_user_group_memberships" {
	#Required
	compartment_id = "${var.tenancy_ocid}"

	#Optional
	group_id = "${oci_identity_group.test_group.id}"
	user_id = "${oci_identity_user.test_user.id}"
}
```