---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_ui_password"
sidebar_current: "docs-oci-datasource-identity-ui_password"
description: |-
  Provides details about a specific Ui Password in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_ui_password
This data source provides details about a specific Ui Password resource in Oracle Cloud Infrastructure Identity service.

Gets the specified user's console password information. The returned object contains the user's OCID,
but not the password itself. The actual password is returned only when created or reset.


## Example Usage

```hcl
data "oci_identity_ui_password" "test_ui_password" {
	#Required
	user_id = "${oci_identity_user.test_user.id}"
}
```

## Argument Reference

The following arguments are supported:

* `user_id` - (Required) The OCID of the user.


## Attributes Reference

The following attributes are exported:

* `state` - The password's current state.
* `time_created` - Date and time the password was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `user_id` - The OCID of the user.

