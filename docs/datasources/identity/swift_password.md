# oci\_identity\_swift\_passwords

Lists swift passwords

## Example Usage

```
data "oci_identity_swift_passwords" "p" {
    compartment_id = "compartment ocid"
    user_id = "user ocid"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment container the user.
* `group_id` - (Optional) The OCID of the group. At least one of group_id or user_id is required.
* `user_id` - (Optional) The OCID of the user. At least one of group_id or user_id is required.

## Attribute Reference
* `passwords` - A list of swift passwords

## Swift Password Reference
* `id` - The OCID of the Swift password.
* `user_id` - The OCID of the user the password belongs to.
* `password` - The Swift password. The value is available only in the response for CreateSwiftPassword, and not for ListSwiftPasswords or UpdateSwiftPassword.
* `description` - The description you assign to the Swift password. Does not have to be unique, and it's changeable.
* `expires_on` - Date and time the password will expire, in the format defined by RFC3339.
* `time_created` - Date and time the password was created, in the format defined by RFC3339.
* `state` - The membership's current state. [CREATING, ACTIVE, INACTIVE, DELETING, DELETED]
* `inactive_status` - The detailed status of INACTIVE lifecycleState.
