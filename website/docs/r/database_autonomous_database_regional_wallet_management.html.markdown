---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_database_regional_wallet_management"
sidebar_current: "docs-oci-resource-database-autonomous_database_regional_wallet_management"
description: |-
  Provides the Autonomous Database Regional Wallet Management resource in Oracle Cloud Infrastructure Database service
---

# oci_database_autonomous_database_regional_wallet_management
This resource provides the Autonomous Database Regional Wallet Management resource in Oracle Cloud Infrastructure Database service.

Updates the Autonomous Database regional wallet.


## Example Usage

```hcl
resource "oci_database_autonomous_database_regional_wallet_management" "test_autonomous_database_regional_wallet_management" {

	#Optional
	should_rotate = var.autonomous_database_regional_wallet_management_should_rotate
}
```

## Argument Reference

The following arguments are supported:

* `should_rotate` - (Optional) (Updatable) Indicates whether to rotate the wallet or not. If `false`, the wallet will not be rotated. The default is `false`.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `state` - The current lifecycle state of the Autonomous Database wallet.
* `time_rotated` - The date and time the wallet was last rotated.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 2 hours), when creating the Autonomous Database Regional Wallet Management
	* `update` - (Defaults to 2 hours), when updating the Autonomous Database Regional Wallet Management
	* `delete` - (Defaults to 2 hours), when destroying the Autonomous Database Regional Wallet Management


## Import

Import is not supported for this resource.

