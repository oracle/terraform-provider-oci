---
layout: "oci"
page_title: "OCI: oci_identity_customer_secret_key"
sidebar_current: "docs-oci-resource-identity-customer_secret_key"
description: |-
Creates and manages an OCI CustomerSecretKey
---

# oci_identity_customer_secret_key
The `oci_identity_customer_secret_key` resource creates and manages an OCI CustomerSecretKey

Creates a new secret key for the specified user. Secret keys are used for authentication with the Object Storage Service's Amazon S3
compatible API. For information, see
[Managing User Credentials](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Tasks/managingcredentials.htm).

You must specify a *description* for the secret key (although it can be an empty string). It does not
have to be unique, and you can change it anytime with
[UpdateCustomerSecretKey](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/CustomerSecretKeySummary/UpdateCustomerSecretKey).

Every user has permission to create a secret key for *their own user ID*. An administrator in your organization
does not need to write a policy to give users this ability. To compare, administrators who have permission to the
tenancy can use this operation to create a secret key for any user, including themselves.


## Example Usage

```hcl
resource "oci_identity_customer_secret_key" "test_customer_secret_key" {
	#Required
	display_name = "${var.customer_secret_key_display_name}"
	user_id = "${oci_identity_user.test_user.id}"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) (Updatable) The name you assign to the secret key during creation. Does not have to be unique, and it's changeable. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `display_name` - The displayName you assign to the secret key. Does not have to be unique, and it's changeable.
* `id` - The OCID of the secret key.
* `inactive_state` - The detailed status of INACTIVE lifecycleState.
* `state` - The secret key's current state. After creating a secret key, make sure its `lifecycleState` changes from CREATING to ACTIVE before using it. 
* `time_created` - Date and time the `CustomerSecretKey` object was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `time_expires` - Date and time when this password will expire, in the format defined by RFC3339. Null if it never expires.  Example: `2016-08-25T21:10:29.600Z` 
* `user_id` - The OCID of the user the password belongs to.
