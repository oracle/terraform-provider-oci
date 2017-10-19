# oci\_identity\_swift\_password

[SwiftPassword Reference][a4047143]

  [a4047143]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/SwiftPassword/ "SwiftPasswordReference"

Provides a Swift password resource.

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
* `description` - (Required) The description you assign to the Swift password. Does not have to be unique, and it's changeable. Avoid entering confidential information.

## Attributes Reference
* `user_id` - The OCID of the user.
* `description` - The description you assign to the Swift password. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `password` - The Swift password. The value is available only in the response for `CreateSwiftPassword`, and not for `ListSwiftPasswords` or `UpdateSwiftPassword`.
* `expires_on` - Date and time when this password will expire, in the format defined by RFC3339. Null if it never expires.
* `time_created` - The date and time the password was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z`.
* `state` - The user's current state. Allowed values are: [CREATING, ACTIVE, INACTIVE, DELETING, DELETED]
* `inactive_status` - The detailed status of INACTIVE `lifecycleState`.
