# oci\_identity\_user_group_memberships

Lists user_group_memberships

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
* `memberships` - A list of user_group_memberships

## User Group Membership Reference
* `id` - The OCID of the user.
* `compartment_id` - The OCID of the tenancy containing the user.
* `group_id` - The OCID of the group.
* `user_id` - The OCID of the user.
* `time_created` - Date and time the user was created, in the format defined by RFC3339.
* `state` - The membership's current state. [CREATING, ACTIVE, INACTIVE, DELETING, DELETED]
* `inactive_status` - The detailed status of INACTIVE lifecycleState.
