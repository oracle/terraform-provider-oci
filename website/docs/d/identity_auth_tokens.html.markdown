---
layout: "oci"
page_title: "OCI: oci_identity_auth_tokens"
sidebar_current: "docs-oci-datasource-identity-auth_tokens"
description: |-
  Provides a list of AuthTokens
---

# Data Source: oci_identity_auth_tokens
The `oci_identity_auth_tokens` data source allows access to the list of OCI auth_tokens

Lists the auth tokens for the specified user. The returned object contains the token's OCID, but not
the token itself. The actual token is returned only upon creation.


## Example Usage

```hcl
data "oci_identity_auth_tokens" "test_auth_tokens" {
	#Required
	user_id = "${oci_identity_user.test_user.id}"
}
```

## Argument Reference

The following arguments are supported:

* `user_id` - (Required) The OCID of the user.


## Attributes Reference

The following attributes are exported:

* `tokens` - The list of tokens.

### AuthToken Reference

The following attributes are exported:

* `description` - The description you assign to the auth token. Does not have to be unique, and it's changeable.
* `id` - The OCID of the auth token.
* `inactive_state` - The detailed status of INACTIVE lifecycleState.
* `state` - The token's current state.
* `time_created` - Date and time the `AuthToken` object was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `time_expires` - Date and time when this auth token will expire, in the format defined by RFC3339. Null if it never expires.  Example: `2016-08-25T21:10:29.600Z` 
* `token` - The auth token. The value is available only in the response for `CreateAuthToken`, and not for `ListAuthTokens` or `UpdateAuthToken`. 
* `user_id` - The OCID of the user the auth token belongs to.

