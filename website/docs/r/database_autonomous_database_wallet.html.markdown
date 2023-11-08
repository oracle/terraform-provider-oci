---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_database_wallet"
sidebar_current: "docs-oci-resource-database-autonomous_database_wallet"
description: |-
  Provides the Autonomous Database Wallet resource in Oracle Cloud Infrastructure Database service
---

# oci_database_autonomous_database_wallet
This resource provides the Autonomous Database Wallet resource in Oracle Cloud Infrastructure Database service.

Creates and downloads a wallet for the specified Autonomous Database.

If passing the base64 encoded content to a `local_file` resource, please use the `content_base64` attribute of the `local_file` resource.
See this [example](https://github.com/oracle/terraform-provider-oci/blob/master/examples/database/adb/autonomous_data_warehouse_wallet.tf) for more details.

Recreate the resource to create and download a new wallet.

## Example Usage

```hcl
resource "oci_database_autonomous_database_wallet" "test_autonomous_database_wallet" {
	#Required
	autonomous_database_id = oci_database_autonomous_database.test_autonomous_database.id
	password = var.autonomous_database_wallet_password

	#Optional
	base64_encode_content = "false"
	generate_type = var.autonomous_database_wallet_generate_type
	is_regional = var.autonomous_database_wallet_is_regional
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_database_id` - (Required) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `base64_encode_content` - (Optional) Encodes the downloaded zipped wallet in base64. It is recommended to set this to `true` to avoid corrupting the zip file in Terraform state. The default value is `false` to preserve backwards compatibility with Terraform v0.11 configurations.
* `generate_type` - (Optional) The type of wallet to generate.

	**Serverless instance usage:**
	* `SINGLE` - used to generate a wallet for a single database
	* `ALL` - used to generate wallet for all databases in the region

	**Dedicated Exadata infrastructure usage:** Value must be `NULL` if attribute is used. 
* `is_regional` - (Optional) True when requesting regional connection strings in PDB connect info, applicable to cross-region DG only.
* `password` - (Required) The password to encrypt the keys inside the wallet. The password must be at least 8 characters long and must include at least 1 letter and either 1 numeric character or 1 special character.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `content` - content of the downloaded zipped wallet for the Autonomous Database. If `base64_encode_content` is set to `true`, then this content will be base64 encoded.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Autonomous Database Wallet
	* `update` - (Defaults to 20 minutes), when updating the Autonomous Database Wallet
	* `delete` - (Defaults to 20 minutes), when destroying the Autonomous Database Wallet


## Import

Import is not supported for this resource.

