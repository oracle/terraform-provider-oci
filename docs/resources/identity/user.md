# oci\_identity\_user

[User Reference][9d58a624]

  [9d58a624]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/User/ "UserReference"

Provides a user resource, which can be an individual or system.

## Example Usage

```
resource "oci_identity_user" "t" {
			name = "name!"
			description = "desc!"
		}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `name` - (Required) The name you assign to the user during creation. This is the user's login for the Console. The name must be unique across all users in the tenancy and cannot be changed. Avoid entering confidential information.
* `description` - (Required) The description you assign to the user during creation. Does not have to be unique, and it's changeable. Avoid entering confidential information.

## Attributes Reference
* `id` - The internet gateway's Oracle Cloud ID (OCID).
* `compartment_id` - The OCID of the compartment containing the user.
* `name` - The name you assign to the user during creation. This is the user's login for the Console. The name must be unique across all users in the tenancy and cannot be changed. Avoid entering confidential information.
* `description` - The description you assign to the user. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `time_created` - The date and time the user was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z`.
* `state` - The user's current state. Allowed values are: [CREATING, ACTIVE, INACTIVE, DELETING, DELETED]
* `inactive_status` - Returned only if the user's `lifecycleState` is INACTIVE. A 16-bit value showing the reason why the user is inactive: [bit 0: SUSPENDED, bit 1: DISABLED, bit 2: BLOCKED]
