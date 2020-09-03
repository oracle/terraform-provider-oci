---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_user"
sidebar_current: "docs-oci-resource-identity-user_capabilities"
description: |-
  Provides the User Capabilities Management resource in Oracle Cloud Infrastructure Identity service
---

# oci_identity_user
This resource provides the User Capabilities Management resource in Oracle Cloud Infrastructure Identity service.

Manages the capabilities of the specified user.

**Important:** Deleting the User Capabilities Management leaves the User resource in its existing state (rather than returning to its defaults)


## Example Usage

```hcl
resource "oci_identity_user_capabilities_management" "test_user_capabilities_management" {
	#Required
	user_id = oci_identity_user.user1.id

	#Optional 
	can_use_api_keys             = "true"
	can_use_auth_tokens          = "true"
	can_use_console_password     = "false"
	can_use_customer_secret_keys = "true"
	can_use_smtp_credentials     = "true"
}
```

## Argument Reference

The following arguments are supported:

* `user_id` - (Required) The OCID of the user.
* `can_use_api_keys` - (Optional) (Updatable) Indicates if the user can use API keys.
* `can_use_auth_tokens` - (Optional) (Updatable) Indicates if the user can use SWIFT passwords / auth tokens.
* `can_use_console_password` - (Optional) (Updatable) Indicates if the user can log in to the console.
* `can_use_customer_secret_keys` - (Optional) (Updatable) Indicates if the user can use SigV4 symmetric keys.
* `can_use_smtp_credentials` - (Optional) (Updatable) Indicates if the user can use SMTP passwords.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `user_id` - The OCID of the user.
* `can_use_api_keys` - Indicates if the user can use API keys.
* `can_use_auth_tokens` - Indicates if the user can use SWIFT passwords / auth tokens.
* `can_use_console_password` - Indicates if the user can log in to the console.
* `can_use_customer_secret_keys` - Indicates if the user can use SigV4 symmetric keys.
* `can_use_smtp_credentials` - Indicates if the user can use SMTP passwords.

## Import

Users can be imported using the `id`, e.g.

```
$ terraform import oci_identity_user_capabilities_management.test_user_capabilities_management "capabilities/{userId}"
```

