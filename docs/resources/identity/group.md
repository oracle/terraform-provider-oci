# baremetal\_identity\_group

Provides a group resource.

## Example Usage

```
resource "baremetal_identity_group" "t" {
    name = "name!"
    description = "desc!"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name you assign to the group during creation. The name must be unique across all groups in the tenancy and cannot be changed.
* `description` - (Required) The description you assign to the group during creation. Does not have to be unique, and it's changeable.

## Attributes Reference
* `id` - The OCID of the group.
* `compartment_id` - The OCID of the compartment containing the group.
* `name` - The name you assign to the group during creation. The name must be unique across all groups in the tenancy and cannot be changed.
* `descriptions` - The description you assign to the group. Does not have to be unique, and it's changeable.
* `time_created` - Date and time the group was created.
* `state` - The group's current state. [CREATING, ACTIVE, INACTIVE, DELETING, DELETED]
* `inactive_status` - The detailed status of INACTIVE lifecycleState.
