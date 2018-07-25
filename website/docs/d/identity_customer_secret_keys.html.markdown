---
layout: "oci"
page_title: "OCI: oci_identity_customer_secret_keys"
sidebar_current: "docs-oci-datasource-identity-customer_secret_keys"
description: |-
  Provides a list of CustomerSecretKeys
---

# Data Source: oci_identity_customer_secret_keys
The `oci_identity_customer_secret_keys` data source allows access to the list of OCI customer_secret_keys

Lists the secret keys for the specified user. The returned object contains the secret key's OCID, but not
the secret key itself. The actual secret key is returned only upon creation.


## Example Usage

```hcl
data "oci_identity_customer_secret_keys" "test_customer_secret_keys" {
	#Required
	user_id = "${oci_identity_user.test_user.id}"
}
```

## Argument Reference

The following arguments are supported:

* `user_id` - (Required) The OCID of the user.


## Attributes Reference

The following attributes are exported:

* `customer_secret_keys` - The list of customer_secret_keys.

### CustomerSecretKey Reference

The following attributes are exported:

* `display_name` - The displayName you assign to the secret key. Does not have to be unique, and it's changeable.
* `id` - The OCID of the secret key.
* `inactive_state` - The detailed status of INACTIVE lifecycleState.
* `state` - The secret key's current state. After creating a secret key, make sure its `lifecycleState` changes from CREATING to ACTIVE before using it. 
* `time_created` - Date and time the `CustomerSecretKey` object was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `time_expires` - Date and time when this password will expire, in the format defined by RFC3339. Null if it never expires.  Example: `2016-08-25T21:10:29.600Z` 
* `user_id` - The OCID of the user the password belongs to.

