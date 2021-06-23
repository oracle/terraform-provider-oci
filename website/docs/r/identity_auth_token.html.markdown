---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_auth_token"
sidebar_current: "docs-oci-resource-identity-auth_token"
description: |-
  Provides the Auth Token resource in Oracle Cloud Infrastructure Identity service
---

# oci_identity_auth_token
This resource provides the Auth Token resource in Oracle Cloud Infrastructure Identity service.

Creates a new auth token for the specified user. For information about what auth tokens are for, see
[Managing User Credentials](https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcredentials.htm).

You must specify a *description* for the auth token (although it can be an empty string). It does not
have to be unique, and you can change it anytime with
[UpdateAuthToken](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/AuthToken/UpdateAuthToken).

Every user has permission to create an auth token for *their own user ID*. An administrator in your organization
does not need to write a policy to give users this ability. To compare, administrators who have permission to the
tenancy can use this operation to create an auth token for any user, including themselves.


## Example Usage

```hcl
resource "oci_identity_auth_token" "test_auth_token" {
	#Required
	description = var.auth_token_description
	user_id = oci_identity_user.test_user.id
}
```

## Argument Reference

The following arguments are supported:

* `description` - (Required) (Updatable) The description you assign to the auth token during creation. Does not have to be unique, and it's changeable. 
* `user_id` - (Required) The OCID of the user.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `description` - The description you assign to the auth token. Does not have to be unique, and it's changeable.
* `id` - The OCID of the auth token.
* `inactive_state` - The detailed status of INACTIVE lifecycleState.
* `state` - The token's current state.
* `time_created` - Date and time the `AuthToken` object was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `time_expires` - Date and time when this auth token will expire, in the format defined by RFC3339. Null if it never expires.  Example: `2016-08-25T21:10:29.600Z` 
* `token` - The auth token. The value is available only in the response for `CreateAuthToken`, and not for `ListAuthTokens` or `UpdateAuthToken`. 
* `user_id` - The OCID of the user the auth token belongs to.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Auth Token
	* `update` - (Defaults to 20 minutes), when updating the Auth Token
	* `delete` - (Defaults to 20 minutes), when destroying the Auth Token


## Import

AuthTokens can be imported using the `id`, e.g.

```
$ terraform import oci_identity_auth_token.test_auth_token "users/{userId}/authTokens/{authTokenId}" 
```

