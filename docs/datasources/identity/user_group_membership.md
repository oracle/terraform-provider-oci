# oci\_identity\_user_group_memberships

[UserGroupMembership Reference][f5d3fcd7]

  [f5d3fcd7]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/UserGroupMembership/ "UserGroupMembershipReference"

Lists user_group_memberships, the membership of a user in a group.

## Example Usage

```
data "oci_identity_user_group_memberships" "g_memberships" {
    compartment_id = "cid"
    group_id = "${oci_identity_group.g.id}"
}`
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the tenancy containing the user, group, and membership object.
* `group_id` - (Optional) The OCID of the group. At least one of group_id or user_id is required.
* `user_id` - (Optional) The OCID of the user. At least one of group_id or user_id is required.

## Attribute Reference
* `memberships` - A list of user_group_memberships.

## User Group Membership Reference
* `id` - The OCID of the user.
* `compartment_id` - The OCID of the tenancy containing the user.
* `group_id` - The OCID of the group.
* `user_id` - The OCID of the user.
* `time_created` - Date and time the membership was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z`.
* `state` - The membership's current state. Allowed values are: [CREATING, ACTIVE, INACTIVE, DELETING, DELETED]
* `inactive_status` - The detailed status of INACTIVE lifecycleState.
