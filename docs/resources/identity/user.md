# baremetal\_identity\_user

Provides a user resource.

## Example Usage

```
resource "baremetal_identity_user" "t" {
			name = "name!"
			description = "desc!"
		}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `name` - (Required) The name you assign to the user during creation. This is the user's login for the Console. The name must be unique across all users in the tenancy and cannot be changed.
* `description` - (Required) The description you assign to the user during creation. Does not have to be unique, and it's changeable.

## Attributes Reference
* `id` - The internet gateway's Oracle Cloud ID (OCID).
* `compartment_id` - The OCID of the compartment containing the user.
* `name` - The name you assign to the user during creation. This is the user's login for the Console. The name must be unique across all users in the tenancy and cannot be changed.
* `description` - The description you assign to the user. Does not have to be unique, and it's changeable.
* `time_created` - The date and time the security list was created.
* `state` - The user's current state. [CREATING, ACTIVE, INACTIVE, DELETING, DELETED]
* `inactive_status` - Returned only if the user's lifecycleState is INACTIVE. A 16-bit value showing the reason why the user is inactive: [bit 0: SUSPENDED, bit 1: DISABLED, bit 2: BLOCKED]