---
layout: "oci"
page_title: "OCI: oci_identity_api_keys"
sidebar_current: "docs-oci-datasource-identity-api_keys"
description: |-
  Provides a list of ApiKeys
---

# Data Source: oci_identity_api_keys
The ApiKeys data source allows access to the list of OCI api_keys

Lists the API signing keys for the specified user. A user can have a maximum of three keys.

Every user has permission to use this API call for *their own user ID*.  An administrator in your
organization does not need to write a policy to give users this ability.


## Example Usage

```hcl
data "oci_identity_api_keys" "test_api_keys" {
	#Required
	user_id = "${oci_identity_user.test_user.id}"
}
```

## Argument Reference

The following arguments are supported:

* `user_id` - (Required) The OCID of the user.


## Attributes Reference

The following attributes are exported:

* `api_keys` - The list of api_keys.

### ApiKey Reference

The following attributes are exported:

* `fingerprint` - The key's fingerprint (e.g., 12:34:56:78:90:ab:cd:ef:12:34:56:78:90:ab:cd:ef).
* `id` - An Oracle-assigned identifier for the key, in this format: TENANCY_OCID/USER_OCID/KEY_FINGERPRINT. 
* `inactive_status` - The detailed status of INACTIVE lifecycleState.
* `key_value` - The key's value.
* `state` - The API key's current state. After creating an `ApiKey` object, make sure its `lifecycleState` changes from CREATING to ACTIVE before using it. 
* `time_created` - Date and time the `ApiKey` object was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `user_id` - The OCID of the user the key belongs to.

