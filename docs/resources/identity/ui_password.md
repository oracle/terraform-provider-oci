# baremetal\_identity\_ui\_password

Provides a ui password resource.

## Example Usage

```
resource "baremetal_identity_ui_password" "t" {
    user_id = "user_id"
}
```

## Argument Reference

The following arguments are supported:

* `user_id` - (Required) The OCID of the user.

## Attributes Reference
* `user_id` - The OCID of the user.
* `password` - The user's password for the Console.
* `time_created` - The date and time the password was created.
* `state` - The user's current state. [CREATING, ACTIVE, INACTIVE, DELETING, DELETED]
* `inactive_status` - The detailed status of INACTIVE lifecycleState.