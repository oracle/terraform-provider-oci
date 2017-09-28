# oci\_identity\_users

Lists users

## Example Usage

```
data "oci_identity_users" "t" {
  compartment_id = "compartment"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.

## Attribute Reference
* `users` - A list of users

## User Reference
* `id` - The OCID of the user.
* `compartment_id` - The OCID of the tenancy containing the user.
* `name` - The name you assign to the user during creation. This is the user's login for the Console. The name must be unique across all users in the tenancy and cannot be changed.
* `description` - The description you assign to the user. Does not have to be unique, and it's changeable.
* `time_created` - Date and time the user was created, in the format defined by RFC3339.
* `state` - The user's current state. [CREATING, ACTIVE, INACTIVE, DELETING, DELETED]
* `inactive_status` - The detailed status of INACTIVE lifecycleState.
