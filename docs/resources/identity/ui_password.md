# oci\_identity\_ui\_password

[UIPassword Reference][b574e93d]

  [b574e93d]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/UIPassword/ "UIPasswordReference"

Provides a UI password resource.

## Example Usage

```
resource "oci_identity_ui_password" "t" {
    user_id = "user_id"
}
```

## Argument Reference

The following arguments are supported:

* `user_id` - (Required) The OCID of the user.

## Attributes Reference
* `user_id` - The OCID of the user.
* `password` - The user's password for the Console.
* `time_created` - The date and time the password was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z`.
* `state` - The user's current state. Allowed values are: [CREATING, ACTIVE, INACTIVE, DELETING, DELETED]
* `inactive_status` - The detailed status of INACTIVE `lifecycleState`.
