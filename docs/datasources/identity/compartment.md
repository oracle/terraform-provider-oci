# baremetal\_identity\_compartments

Lists compartments

## Example Usage

```
data "baremetal_identity_compartments" "t" {
  compartment_id = "compartment"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the tenancy.

## Attribute Reference
* `compartments` - A list of compartments

## Group Reference
* `id` - The OCID of the compartment.
* `compartment_id` - The OCID of the tenancy containing the compartment.
* `name` - The name you assign to the compartment during creation. The name must be unique across all compartments in the tenancy and cannot be changed.
* `description` - The description you assign to the compartment. Does not have to be unique, and it's changeable.
* `time_created` - Date and time the compartment was created.
* `state` - The compartment's current state. [CREATING, ACTIVE, INACTIVE, DELETING, DELETED]
* `inactive_status` - The detailed status of INACTIVE lifecycleState.
