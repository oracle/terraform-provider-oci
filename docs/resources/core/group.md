# baremetal\_identity\_groups

Provide a group resource

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

## Attribute Reference
* `id` - The OCID of the group.
* `compartment_id` - The OCID of the tenancy containing the group.
* `name` - The name you assign to the group during creation. The name must be unique across all groups in the tenancy and cannot be changed.
* `description` - The description you assign to the group. Does not have to be unique, and it's changeable.
* `time_created` - Date and time the ApiKey was created.
* `state` - The compartment's current state. [CREATING, ACTIVE, INACTIVE, DELETING, DELETED]
* `inactive_status` - The detailed status of INACTIVE lifecycleState.
