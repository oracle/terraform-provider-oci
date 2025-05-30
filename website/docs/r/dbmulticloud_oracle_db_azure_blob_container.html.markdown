---
subcategory: "Dbmulticloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dbmulticloud_oracle_db_azure_blob_container"
sidebar_current: "docs-oci-resource-dbmulticloud-oracle_db_azure_blob_container"
description: |-
  Provides the Oracle Db Azure Blob Container resource in Oracle Cloud Infrastructure Dbmulticloud service
---

# oci_dbmulticloud_oracle_db_azure_blob_container
This resource provides the Oracle Db Azure Blob Container resource in Oracle Cloud Infrastructure Dbmulticloud service.

Capture Azure Container details for mounting Azure Container on multiple Oracle Cloud Infrastructure Database Resource.


## Example Usage

```hcl
resource "oci_dbmulticloud_oracle_db_azure_blob_container" "test_oracle_db_azure_blob_container" {
	#Required
	azure_storage_account_name = var.oracle_db_azure_blob_container_azure_storage_account_name
	azure_storage_container_name = var.oracle_db_azure_blob_container_azure_storage_container_name
	compartment_id = var.compartment_id
	display_name = var.oracle_db_azure_blob_container_display_name

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	private_endpoint_dns_alias = var.oracle_db_azure_blob_container_private_endpoint_dns_alias
	private_endpoint_ip_address = var.oracle_db_azure_blob_container_private_endpoint_ip_address
}
```

## Argument Reference

The following arguments are supported:

* `azure_storage_account_name` - (Required) (Updatable) Azure Storage Account Name.
* `azure_storage_container_name` - (Required) (Updatable) Azure Storage Container Name.
* `compartment_id` - (Required) (Updatable) The OCID of the compartment that contains Oracle DB Azure Blob Container Resource.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Required) (Updatable) Display name of Oracle DB Azure Blob Container.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `private_endpoint_dns_alias` - (Optional) (Updatable) Private endpoint DNS Alias.
* `private_endpoint_ip_address` - (Optional) (Updatable) Private endpoint IP.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `azure_storage_account_name` - Azure Storage Account Name.
* `azure_storage_container_name` - Azure Storage Container Name.
* `compartment_id` - The ID of the compartment that contains Oracle DB Azure Blob Container Resource.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - Display name of Oracle DB Azure Blob Container.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The ID of the compartment that contains Oracle DB Azure Blob Container Resource.
* `last_modification` - Description of the latest modification of the Oracle DB Azure Blob Container Resource.
* `lifecycle_state_details` - Description of the current lifecycle state in more detail.
* `private_endpoint_dns_alias` - Private endpoint DNS Alias.
* `private_endpoint_ip_address` - Private endpoint IP.
* `state` - The current lifecycle state of the Oracle DB Azure Blob Container Resource.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Time when the Oracle DB Azure Blob Container was created in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z' 
* `time_updated` - Time when the Oracle DB Azure Blob Container was last modified, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z' 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Oracle Db Azure Blob Container
	* `update` - (Defaults to 20 minutes), when updating the Oracle Db Azure Blob Container
	* `delete` - (Defaults to 20 minutes), when destroying the Oracle Db Azure Blob Container


## Import

OracleDbAzureBlobContainers can be imported using the `id`, e.g.

```
$ terraform import oci_dbmulticloud_oracle_db_azure_blob_container.test_oracle_db_azure_blob_container "id"
```

