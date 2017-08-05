# oci\_identity\_groups

Lists groups

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
* `name` - The name you assign to the group during creation. The name must be unique across all groups in the tenancy and cannot be changed.
* `description` - The description you assign to the group. Does not have to be unique, and it's changeable.
* `time_created` - Date and time the group was created.
* `state` - The group's current state. [CREATING, ACTIVE, INACTIVE, DELETING, DELETED]
* `inactive_status` - The detailed status of INACTIVE lifecycleState.
