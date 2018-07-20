---
layout: "oci"
page_title: "OCI: oci_identity_api_keys"
sidebar_current: "docs-oci-datasource-api_keys"
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