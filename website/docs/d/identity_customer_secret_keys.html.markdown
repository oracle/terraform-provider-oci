---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_customer_secret_keys"
sidebar_current: "docs-oci-datasource-identity-customer_secret_keys"
description: |-
  Provides the list of Customer Secret Keys in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_customer_secret_keys
This data source provides the list of Customer Secret Keys in Oracle Cloud Infrastructure Identity service.

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

* `display_name` - The display name you assign to the secret key. Does not have to be unique, and it's changeable.
* `id` - The OCID of the secret key.
* `inactive_state` - The detailed status of INACTIVE lifecycleState.
* `state` - The secret key's current state.
* `time_created` - Date and time the `CustomerSecretKey` object was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `time_expires` - Date and time when this password will expire, in the format defined by RFC3339. Null if it never expires.  Example: `2016-08-25T21:10:29.600Z` 
* `user_id` - The OCID of the user the password belongs to.

