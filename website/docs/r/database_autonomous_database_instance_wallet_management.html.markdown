---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_database_instance_wallet_management"
sidebar_current: "docs-oci-resource-database-autonomous_database_instance_wallet_management"
description: |-
  Provides the Autonomous Database Instance Wallet Management resource in Oracle Cloud Infrastructure Database service
---

# oci_database_autonomous_database_instance_wallet_management
This resource provides the Autonomous Database Instance Wallet Management resource in Oracle Cloud Infrastructure Database service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/database/latest/AutonomousDatabaseInstanceWalletManagement

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/database

Updates the wallet for the specified Autonomous AI Database.


## Example Usage

```hcl
resource "oci_database_autonomous_database_instance_wallet_management" "test_autonomous_database_instance_wallet_management" {
	#Required
	autonomous_database_id = oci_database_autonomous_database.test_autonomous_database.id

	#Optional
	grace_period = var.autonomous_database_instance_wallet_management_grace_period
	should_rotate = var.autonomous_database_instance_wallet_management_should_rotate
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_database_id` - (Required) (Updatable) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `grace_period` - (Optional) (Updatable) The number of hours that the old wallet can be used after it has been rotated. The old wallet will no longer be valid after the number of hours in the wallet rotation grace period has passed. During the grace period, both the old wallet and the current wallet can be used.
* `should_rotate` - (Optional) (Updatable) Indicates whether to rotate the wallet or not. If `false`, the wallet will not be rotated. The default is `false`.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `autonomous_database_id` - The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `should_rotate` - Indicates whether to rotate the wallet or not. If `false`, the wallet will not be rotated. The default is `false`.
* `state` - The current lifecycle state of the Autonomous AI Database wallet.
* `time_rotated` - The date and time the wallet was last rotated.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Autonomous Database Instance Wallet Management
	* `update` - (Defaults to 20 minutes), when updating the Autonomous Database Instance Wallet Management
	* `delete` - (Defaults to 20 minutes), when destroying the Autonomous Database Instance Wallet Management


## Import

Import is not supported for this resource.

