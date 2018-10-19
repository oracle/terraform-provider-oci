---
layout: "oci"
page_title: "OCI: oci_database_autonomous_database_wallet"
sidebar_current: "docs-oci-datasource-database-autonomous_database_wallet"
description: |-
  Provides details about a specific AutonomousDatabaseWallet
---

# Data Source: oci_database_autonomous_database_wallet
The `oci_database_autonomous_database_wallet` data source provides details about a specific AutonomousDatabaseWallet



## Example Usage

```hcl
data "oci_database_autonomous_database_wallet" "test_autonomous_database_wallet" {
	#Required
	autonomous_database_id = "${oci_database_autonomous_database.test_autonomous_database.id}"
	password = "${var.autonomous_database_wallet_password}"
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_database_id` - (Required) The database [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
* `password` - (Required) The password to encrypt the keys inside the wallet.


## Attributes Reference

The following attributes are exported:

* `content` - content of the downloaded zipped wallet for the Autonomous Database
