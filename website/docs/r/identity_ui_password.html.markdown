---
layout: "oci"
page_title: "OCI: oci_identity_ui_password"
sidebar_current: "docs-oci-resource-identity-ui_password"
description: |-
  Creates and manages an OCI UiPassword
---

# oci_identity_ui_password
The `oci_identity_ui_password` resource creates and manages an OCI UiPassword

Creates a new Console one-time password for the specified user. For more information about user
credentials, see [User Credentials](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/usercredentials.htm).

Use this operation after creating a new user, or if a user forgets their password. The new one-time
password is returned to you in the response, and you must securely deliver it to the user. They'll
be prompted to change this password the next time they sign in to the Console. If they don't change
it within 7 days, the password will expire and you'll need to create a new one-time password for the
user.

**Note:** The user's Console login is the unique name you specified when you created the user
(see [CreateUser](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/User/CreateUser)).


## Example Usage

```hcl
resource "oci_identity_ui_password" "test_ui_password" {
	#Required
	user_id = "${oci_identity_user.test_user.id}"
}
```

## Argument Reference

The following arguments are supported:

* `user_id` - (Required) The OCID of the user.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `inactive_status` - The detailed status of INACTIVE lifecycleState.
* `password` - The user's password for the Console.
* `state` - The password's current state.
* `time_created` - Date and time the password was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `user_id` - The OCID of the user.

## Import

UiPasswords can be imported using the `id`, e.g.

```
$ terraform import oci_identity_ui_password.test_ui_password "id"
```
