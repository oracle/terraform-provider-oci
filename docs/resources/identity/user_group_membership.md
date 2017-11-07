# oci\_identity\_user\_group\_membership

[UserGroupMembership Reference][c5525919]

  [c5525919]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/UserGroupMembership/ "UserGroupMembershipReference"

Provides a user group membership resource.

## Example Usage

```
resource "oci_identity_user_group_membership" "t" {
			compartment_id = "cid"
            	user_id = "${oci_identity_user.u.id}"
            	group_id = "${oci_identity_group.g.id}"
		}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the tenancy containing the user, group, and membership object.
* `user_id` - (Required) The OCID of the group.
* `group_id` - (Required) The OCID of the user.

## Attributes Reference
* `id` - The internet gateway's Oracle Cloud ID (OCID).
* `compartment_id` - The OCID of the tenancy containing the user, group, and membership object.
* `user_id` - The OCID of the user.
* `group_id` - The OCID of the group.
* `time_created` - The date and time the membership was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z`.
* `state` - The user's current state. Allowed values are: [CREATING, ACTIVE, INACTIVE, DELETING, DELETED]
* `inactive_status` - Returned only if the user's `lifecycleState` is INACTIVE. A 16-bit value showing the reason why the user is inactive: [bit 0: SUSPENDED, bit 1: DISABLED, bit 2: BLOCKED]
