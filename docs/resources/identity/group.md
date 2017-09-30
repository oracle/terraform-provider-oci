# oci\_identity\_group

[Group Reference][b6a4bdfa]

  [b6a4bdfa]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/Group/ "GroupReference"

Provides a group resource.

## Example Usage

```
resource "oci_identity_group" "t" {
    name = "name!"
    description = "desc!"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name you assign to the group during creation. The name must be unique across all groups in the tenancy and cannot be changed. Avoid entering confidential information.
* `description` - (Required) The description you assign to the group during creation. Does not have to be unique, and it's changeable. Avoid entering confidential information.

## Attributes Reference
* `id` - The OCID of the group.
* `compartment_id` - The OCID of the compartment containing the group.
* `name` - The name you assign to the group during creation. The name must be unique across all groups in the tenancy and cannot be changed. Avoid entering confidential information.
* `descriptions` - The description you assign to the group. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `time_created` - Date and time the group was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z`.
* `state` - The group's current state. Allowed values are: [CREATING, ACTIVE, INACTIVE, DELETING, DELETED]
* `inactive_status` - The detailed status of INACTIVE `lifecycleState`.
