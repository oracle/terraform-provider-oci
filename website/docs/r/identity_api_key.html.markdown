---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_api_key"
sidebar_current: "docs-oci-resource-identity-api_key"
description: |-
  Provides the Api Key resource in Oracle Cloud Infrastructure Identity service
---

# oci_identity_api_key
This resource provides the Api Key resource in Oracle Cloud Infrastructure Identity service.

Uploads an API signing key for the specified user.

Every user has permission to use this operation to upload a key for *their own user ID*. An
administrator in your organization does not need to write a policy to give users this ability.
To compare, administrators who have permission to the tenancy can use this operation to upload a
key for any user, including themselves.

**Important:** Even though you have permission to upload an API key, you might not yet
have permission to do much else. If you try calling an operation unrelated to your own credential
management (e.g., `ListUsers`, `LaunchInstance`) and receive an "unauthorized" error,
check with an administrator to confirm which IAM Service group(s) you're in and what access
you have. Also confirm you're working in the correct compartment.


## Example Usage

```hcl
resource "oci_identity_api_key" "test_api_key" {
	#Required
	key_value = var.api_key_key_value
	user_id = oci_identity_user.test_user.id
}
```

## Argument Reference

The following arguments are supported:

* `key_value` - (Required) The public key.  Must be an RSA key in PEM format.
* `user_id` - (Required) The OCID of the user.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `fingerprint` - The key's fingerprint (e.g., 12:34:56:78:90:ab:cd:ef:12:34:56:78:90:ab:cd:ef).
* `id` - An Oracle-assigned identifier for the key, in this format: TENANCY_OCID/USER_OCID/KEY_FINGERPRINT. 
* `inactive_status` - The detailed status of INACTIVE lifecycleState.
* `key_value` - The key's value.
* `state` - The API key's current state.
* `time_created` - Date and time the `ApiKey` object was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `user_id` - The OCID of the user the key belongs to.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Api Key
	* `update` - (Defaults to 20 minutes), when updating the Api Key
	* `delete` - (Defaults to 20 minutes), when destroying the Api Key


## Import

ApiKeys can be imported using the `id`, e.g.

```
$ terraform import oci_identity_api_key.test_api_key "users/{userId}/apiKeys/{fingerprint}" 
```

