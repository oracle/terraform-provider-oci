---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_data_warehouse_wallet"
sidebar_current: "docs-oci-datasource-database-autonomous_data_warehouse_wallet"
description: |-
  Provides details about a specific Autonomous Data Warehouse Wallet in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_data_warehouse_wallet
This data source provides details about a specific Autonomous Data Warehouse Wallet resource in Oracle Cloud Infrastructure Database service.



## Example Usage

```hcl
data "oci_database_autonomous_data_warehouse_wallet" "test_autonomous_data_warehouse_wallet" {
	#Required
	autonomous_data_warehouse_id = "${oci_database_autonomous_data_warehouse.test_autonomous_data_warehouse.id}"
	password = "${var.autonomous_data_warehouse_wallet_password}"
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_data_warehouse_id` - (Required) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `password` - (Required) The password to encrypt the keys inside the wallet.


## Attributes Reference

The following attributes are exported:

* `content` - content of the downloaded zipped wallet for the Autonomous Data Warehouse
