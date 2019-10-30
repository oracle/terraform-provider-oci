---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_database_instance_wallet_management"
sidebar_current: "docs-oci-datasource-database-autonomous_database_instance_wallet_management"
description: |-
  Provides details about a specific Autonomous Database Instance Wallet Management in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_database_instance_wallet_management
This data source provides details about a specific Autonomous Database Instance Wallet Management resource in Oracle Cloud Infrastructure Database service.

Gets the wallet details for the specified Autonomous Database.


## Example Usage

```hcl
data "oci_database_autonomous_database_instance_wallet_management" "test_autonomous_database_instance_wallet_management" {
	#Required
	autonomous_database_id = "${oci_database_autonomous_database.test_autonomous_database.id}"
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_database_id` - (Required) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `state` - The current lifecycle state of the Autonomous Database wallet.
* `time_rotated` - The date and time the wallet was last rotated.

