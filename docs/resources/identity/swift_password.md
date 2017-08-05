# oci\_identity\_swift\_password

Provides a swift password resource.

## Example Usage

```
resource "oci_identity_swift_password" "t" {
    user_id = "user_id"
    description = "nah nah nah"
}
```

## Argument Reference

The following arguments are supported:

* `user_id` - (Required) The OCID of the user.
* `description` - (Required) The description you assign to the Swift password. Does not have to be unique, and it's changeable.

## Attributes Reference
* `user_id` - The OCID of the user.
* `description` - The description you assign to the Swift password. Does not have to be unique, and it's changeable.
* `password` - The Swift password. The value is available only in the response for CreateSwiftPassword, and not for ListSwiftPasswords or UpdateSwiftPassword.
* `expires_on` - Date and time when this password will expire, in the format defined by RFC3339. Null if it never expires.
* `time_created` - The date and time the password was created.
* `state` - The user's current state. [CREATING, ACTIVE, INACTIVE, DELETING, DELETED]
* `inactive_status` - The detailed status of INACTIVE lifecycleState.