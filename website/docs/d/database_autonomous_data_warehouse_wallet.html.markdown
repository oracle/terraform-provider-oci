---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_data_warehouse_wallet"
sidebar_current: "docs-oci-datasource-database-autonomous_data_warehouse_wallet"
description: |-
  Provides details about a specific Autonomous Data Warehouse Wallet in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_data_warehouse_wallet
This data source provides details about a specific Autonomous Data Warehouse Wallet resource in Oracle Cloud Infrastructure Database service.


**IMPORTANT:** This data source is being **deprecated**, use `oci_database_autonomous_database_wallet` instead.

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
* `password` - (Required) The password to encrypt the keys inside the wallet. The password must be at least 8 characters long and must include at least 1 letter and either 1 numeric character or 1 special character.
* `base64_encode_content` - (Optional) Encodes the downloaded zipped wallet in base64. It is recommended to set this to `true` to avoid corrupting the zip file in Terraform state. The default value is `false` to preserve backwards compatibility with Terraform v0.11 configurations.

## Attributes Reference

The following attributes are exported:

* `content` - content of the downloaded zipped wallet for the Autonomous Data Warehouse. If `base64_encode_content` is set to `true`, then this content will be base64 encoded.
If passing the base64 encoded content to a `local_file` resource, please use the `content_base64` attribute of the `local_file` resource.
See this [example](https://github.com/terraform-providers/terraform-provider-oci/blob/master/examples/database/adb/autonomous_database_wallet.tf) for more details.