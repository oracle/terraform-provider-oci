---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_swift_password"
sidebar_current: "docs-oci-resource-identity-swift_password"
description: |-
  Provides the Swift Password resource in Oracle Cloud Infrastructure Identity service
---

# oci_identity_swift_password
This resource provides the Swift Password resource in Oracle Cloud Infrastructure Identity service.

**Deprecated. Use [CreateAuthToken](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/AuthToken/CreateAuthToken) instead.**

Creates a new Swift password for the specified user. For information about what Swift passwords are for, see
[Managing User Credentials](https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcredentials.htm).

You must specify a *description* for the Swift password (although it can be an empty string). It does not
have to be unique, and you can change it anytime with
[UpdateSwiftPassword](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/SwiftPassword/UpdateSwiftPassword).

Every user has permission to create a Swift password for *their own user ID*. An administrator in your organization
does not need to write a policy to give users this ability. To compare, administrators who have permission to the
tenancy can use this operation to create a Swift password for any user, including themselves.


## Example Usage

```hcl
resource "oci_identity_swift_password" "test_swift_password" {
	#Required
	description = var.swift_password_description
	user_id = oci_identity_user.test_user.id
}
```

## Argument Reference

The following arguments are supported:

* `description` - (Required) (Updatable) The description you assign to the Swift password during creation. Does not have to be unique, and it's changeable. 
* `user_id` - (Required) The OCID of the user.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `description` - The description you assign to the Swift password. Does not have to be unique, and it's changeable.
* `expires_on` - Date and time when this password will expire, in the format defined by RFC3339. Null if it never expires.  Example: `2016-08-25T21:10:29.600Z` 
* `id` - The OCID of the Swift password.
* `inactive_state` - The detailed status of INACTIVE lifecycleState.
* `password` - The Swift password. The value is available only in the response for `CreateSwiftPassword`, and not for `ListSwiftPasswords` or `UpdateSwiftPassword`. 
* `state` - The password's current state.
* `time_created` - Date and time the `SwiftPassword` object was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `user_id` - The OCID of the user the password belongs to.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Swift Password
	* `update` - (Defaults to 20 minutes), when updating the Swift Password
	* `delete` - (Defaults to 20 minutes), when destroying the Swift Password


## Import

SwiftPasswords can be imported using the `id`, e.g.

```
$ terraform import oci_identity_swift_password.test_swift_password "users/{userId}/swiftPasswords/{swiftPasswordId}" 
```

