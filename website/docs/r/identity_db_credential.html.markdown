---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_db_credential"
sidebar_current: "docs-oci-resource-identity-db_credential"
description: |-
  Provides the Db Credential resource in Oracle Cloud Infrastructure Identity service
---

# oci_identity_db_credential
This resource provides the Db Credential resource in Oracle Cloud Infrastructure Identity service.

Creates a new DB credential for the specified user.


## Example Usage

```hcl
resource "oci_identity_db_credential" "test_db_credential" {
	#Required
	description = var.db_credential_description
	password = var.db_credential_password
	user_id = oci_identity_user.test_user.id
}
```

## Argument Reference

The following arguments are supported:

* `description` - (Required) The description you assign to the DB credentials during creation. 
* `password` - (Required) The password for the DB credentials during creation. 
* `user_id` - (Required) The OCID of the user.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `description` - The description you assign to the DB credential. Does not have to be unique, and it's changeable.
* `id` - The OCID of the DB credential.
* `lifecycle_details` - The detailed status of INACTIVE lifecycleState.
* `state` - The credential's current state. After creating a DB credential, make sure its `lifecycleState` changes from CREATING to ACTIVE before using it. 
* `time_created` - Date and time the `DbCredential` object was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `time_expires` - Date and time when this credential will expire, in the format defined by RFC3339. Null if it never expires.  Example: `2016-08-25T21:10:29.600Z` 
* `user_id` - The OCID of the user the DB credential belongs to.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Db Credential
	* `update` - (Defaults to 20 minutes), when updating the Db Credential
	* `delete` - (Defaults to 20 minutes), when destroying the Db Credential


## Import

Import is not supported for this resource.

