# baremetal\_identity\_compartment

Provides a compartment resource.

## Example Usage

```
resource "baremetal_identity_compartment" "t" {
    name = "name!"
    description = "desc!"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name you assign to the compartment during creation. The name must be unique across all compartments in the tenancy and cannot be changed.
* `description` - (Required) The description you assign to the compartment during creation. Does not have to be unique, and it's changeable.

## Attributes Reference
* `id` - The OCID of the compartment.
* `compartment_id` - The OCID of the tenancy containing the compartment.
* `name` - The name you assign to the compartment during creation. The name must be unique across all compartments in the tenancy and cannot be changed.
* `descriptions` - The description you assign to the compartment. Does not have to be unique, and it's changeable.
* `time_created` - Date and time the compartment was created.
* `state` - The compartment's current state. [CREATING, ACTIVE, INACTIVE, DELETING, DELETED]
* `inactive_status` - The detailed status of INACTIVE lifecycleState.
