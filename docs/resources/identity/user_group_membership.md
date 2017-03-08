# baremetal\_identity\_user\_group\_membership

Provides a user group membership resource.

## Example Usage

```
resource "baremetal_identity_user_group_membership" "t" {
			compartment_id = "cid"
            	user_id = "${baremetal_identity_user.u.id}"
            	group_id = "${baremetal_identity_group.g.id}"
		}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the tenancy containing the user, group, and membership object.
* `user_id` - (Required) The OCID of the group.
* `group_id` - (Required) The OCID of the user.

## Attributes Reference
* `id` - The internet gateway's Oracle Cloud ID (OCID).
* `compartment_id` - The OCID of the tenancy containing the user, group, and membership object.
* `user_id` - The OCID of the user.
* `group_id` - The OCID of the group.
* `time_created` - The date and time the security list was created.
* `state` - The user's current state. [CREATING, ACTIVE, INACTIVE, DELETING, DELETED]
* `inactive_status` - Returned only if the user's lifecycleState is INACTIVE. A 16-bit value showing the reason why the user is inactive: [bit 0: SUSPENDED, bit 1: DISABLED, bit 2: BLOCKED]