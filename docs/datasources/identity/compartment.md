# oci\_identity\_compartments

[Compartment Reference][1897664e]

  [1897664e]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/Compartment/ "CompartmentReference"

Lists compartments. Each compartment is a collection of related resources.

## Example Usage

```
data "oci_identity_compartments" "t" {
  compartment_id = "compartment"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the tenancy.

## Attribute Reference
* `compartments` - A list of compartments.

## Compartment Reference
* `id` - The OCID of the compartment.
* `compartment_id` - The OCID of the tenancy containing the compartment.
* `name` - The name you assign to the compartment during creation. The name must be unique across all compartments in the tenancy and cannot be changed. Avoid entering confidential information.
* `description` - The description you assign to the compartment. Does not have to be unique, and it's changeable.  Avoid entering confidential information.
* `time_created` - Date and time the compartment was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z`.
* `state` - The compartment's current state. [CREATING, ACTIVE, INACTIVE, DELETING, DELETED]
* `inactive_status` - The detailed status of INACTIVE lifecycleState.
