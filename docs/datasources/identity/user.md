# oci\_identity\_users

[User Reference][bb05d25c]

  [bb05d25c]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/User/ "UserReference"

Lists users. A user is an individual or system that needs to manage OCI resources.

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
* `users` - A list of users.

## Group Reference
* `id` - The OCID of the user.
* `compartment_id` - The OCID of the tenancy containing the user.
* `name` - The name you assign to the user during creation. This is the user's login for the Console. The name must be unique across all users in the tenancy and cannot be changed. Avoid entering confidential information.
* `description` - The description you assign to the user. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `time_created` - Date and time the user was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z`.
* `state` - The user's current state. Allowed values are: [CREATING, ACTIVE, INACTIVE, DELETING, DELETED]
* `inactive_status` - The detailed status of INACTIVE lifecycleState.
