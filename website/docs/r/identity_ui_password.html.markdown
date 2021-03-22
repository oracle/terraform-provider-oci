---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_ui_password"
sidebar_current: "docs-oci-resource-identity-ui_password"
description: |-
  Provides the Ui Password resource in Oracle Cloud Infrastructure Identity service
---

# oci_identity_ui_password
This resource provides the Ui Password resource in Oracle Cloud Infrastructure Identity service.

Creates a new Console one-time password for the specified user. For more information about user
credentials, see [User Credentials](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/usercredentials.htm).

Use this operation after creating a new user, or if a user forgets their password. The new one-time
password is returned to you in the response, and you must securely deliver it to the user. They'll
be prompted to change this password the next time they sign in to the Console. If they don't change
it within 7 days, the password will expire and you'll need to create a new one-time password for the
user.

**Note:** The user's Console login is the unique name you specified when you created the user
(see [CreateUser](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/User/CreateUser)).


## Example Usage

```hcl
resource "oci_identity_ui_password" "test_ui_password" {
	#Required
	user_id = oci_identity_user.test_user.id
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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Ui Password
	* `update` - (Defaults to 20 minutes), when updating the Ui Password
	* `delete` - (Defaults to 20 minutes), when destroying the Ui Password


## Import

Import is not supported for this resource.

