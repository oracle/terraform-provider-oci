---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_customer_secret_key"
sidebar_current: "docs-oci-resource-identity-customer_secret_key"
description: |-
  Provides the Customer Secret Key resource in Oracle Cloud Infrastructure Identity service
---

# oci_identity_customer_secret_key
This resource provides the Customer Secret Key resource in Oracle Cloud Infrastructure Identity service.

Creates a new secret key for the specified user. Secret keys are used for authentication with the Object Storage Service's Amazon S3
compatible API. The secret key consists of an Access Key/Secret Key pair. For information, see
[Managing User Credentials](https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcredentials.htm).

You must specify a *description* for the secret key (although it can be an empty string). It does not
have to be unique, and you can change it anytime with
[UpdateCustomerSecretKey](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/CustomerSecretKeySummary/UpdateCustomerSecretKey).

Every user has permission to create a secret key for *their own user ID*. An administrator in your organization
does not need to write a policy to give users this ability. To compare, administrators who have permission to the
tenancy can use this operation to create a secret key for any user, including themselves.


## Example Usage

```hcl
resource "oci_identity_customer_secret_key" "test_customer_secret_key" {
	#Required
	display_name = var.customer_secret_key_display_name
	user_id = oci_identity_user.test_user.id
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) (Updatable) The name you assign to the secret key during creation. Does not have to be unique, and it's changeable. 
* `user_id` - (Required) The OCID of the user.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `display_name` - The display name you assign to the secret key. Does not have to be unique, and it's changeable.
* `id` - The access key portion of the key pair.
* `inactive_state` - The detailed status of INACTIVE lifecycleState.
* `key` - The secret key. 
* `state` - The secret key's current state.
* `time_created` - Date and time the `CustomerSecretKey` object was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `time_expires` - Date and time when this password will expire, in the format defined by RFC3339. Null if it never expires.  Example: `2016-08-25T21:10:29.600Z` 
* `user_id` - The OCID of the user the password belongs to.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Customer Secret Key
	* `update` - (Defaults to 20 minutes), when updating the Customer Secret Key
	* `delete` - (Defaults to 20 minutes), when destroying the Customer Secret Key


## Import

CustomerSecretKeys can be imported using the `id`, e.g.

```
$ terraform import oci_identity_customer_secret_key.test_customer_secret_key "users/{userId}/customerSecretKeys/{customerSecretKeyId}" 
```

