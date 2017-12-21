# oci\_identity\_compartment

[Compartment Reference][84ff5b4e]

  [84ff5b4e]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/Compartment/ "CompartmentReference"

Provides a compartment resource. If a compartment with the given `name` already exists, then that compartment will be used instead of creating a new compartment. Also, note that compartments may not be deleted.

## Example Usage

```
resource "oci_identity_compartment" "t" {
    name = "name!"
    description = "desc!"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name you assign to the compartment during creation. The name must be unique across all compartments in the tenancy, and it's changeable. Avoid entering confidential information.
* `description` - (Required) The description you assign to the compartment during creation. Does not have to be unique, and it's changeable. Avoid entering confidential information.

## Attributes Reference
* `id` - The OCID of the compartment.
* `compartment_id` - The OCID of the tenancy containing the compartment.
* `name` - The name you assign to the compartment during creation. The name must be unique across all compartments in the tenancy, and it's changeable. Avoid entering confidential information.
* `descriptions` - The description you assign to the compartment. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `time_created` - Date and time the compartment was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z`.
* `state` - The compartment's current state. Allowed values are: [CREATING, ACTIVE, INACTIVE, DELETING, DELETED]
* `inactive_status` - The detailed status of INACTIVE `lifecycleState`.
