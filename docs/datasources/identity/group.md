# oci\_identity\_groups

[Group Reference][852a87b4]

  [852a87b4]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/Group/ "GroupReference"

Lists groups. A group is a collection of users who all need the same type of access to a particular set of resources or compartment.

## Example Usage

```
data "oci_identity_groups" "t" {
  compartment_id = "compartment"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.

## Attribute Reference
* `groups` - A list of groups

## Group Reference
* `id` - The OCID of the group.
* `compartment_id` - The OCID of the tenancy containing the group.
* `name` - The name you assign to the group during creation. The name must be unique across all groups in the tenancy and cannot be changed. Avoid entering confidential information.
* `description` - The description you assign to the group. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `time_created` - Date and time the group was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`.
* `state` - The group's current state. [CREATING, ACTIVE, INACTIVE, DELETING, DELETED]
* `inactive_status` - The detailed status of INACTIVE lifecycleState.
