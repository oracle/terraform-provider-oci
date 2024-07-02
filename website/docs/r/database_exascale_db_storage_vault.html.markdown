---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_exascale_db_storage_vault"
sidebar_current: "docs-oci-resource-database-exascale_db_storage_vault"
description: |-
  Provides the Exascale Db Storage Vault resource in Oracle Cloud Infrastructure Database service
---

# oci_database_exascale_db_storage_vault
This resource provides the Exascale Db Storage Vault resource in Oracle Cloud Infrastructure Database service.

Creates an Exadata Database Storage Vault


## Example Usage

```hcl
resource "oci_database_exascale_db_storage_vault" "test_exascale_db_storage_vault" {
	#Required
	availability_domain = var.exascale_db_storage_vault_availability_domain
	compartment_id = var.compartment_id
	display_name = var.exascale_db_storage_vault_display_name
	high_capacity_database_storage {
		#Required
		total_size_in_gbs = var.exascale_db_storage_vault_high_capacity_database_storage_total_size_in_gbs
	}

	#Optional
	additional_flash_cache_in_percent = var.exascale_db_storage_vault_additional_flash_cache_in_percent
	defined_tags = var.exascale_db_storage_vault_defined_tags
	description = var.exascale_db_storage_vault_description
	freeform_tags = {"Department"= "Finance"}
	time_zone = var.exascale_db_storage_vault_time_zone
}
```

## Argument Reference

The following arguments are supported:

* `additional_flash_cache_in_percent` - (Optional) (Updatable) The size of additional Flash Cache in percentage of High Capacity database storage. 
* `availability_domain` - (Required) The name of the availability domain in which the Exadata Database Storage Vault is located.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `description` - (Optional) (Updatable) Exadata Database Storage Vault description.
* `display_name` - (Required) (Updatable) The user-friendly name for the Exadata Database Storage Vault. The name does not need to be unique.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `high_capacity_database_storage` - (Required) (Updatable) Create exadata Database Storage Details 
	* `total_size_in_gbs` - (Required) (Updatable) Total Capacity 
* `time_zone` - (Optional) The time zone that you want to use for the Exadata Database Storage Vault. For details, see [Time Zones](https://docs.cloud.oracle.com/iaas/Content/Database/References/timezones.htm). 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `additional_flash_cache_in_percent` - The size of additional Flash Cache in percentage of High Capacity database storage. 
* `availability_domain` - The name of the availability domain in which the Exadata Database Storage Vault is located.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `description` - Exadata Database Storage Vault description.
* `display_name` - The user-friendly name for the Exadata Database Storage Vault. The name does not need to be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `high_capacity_database_storage` - Exadata Database Storage Details 
	* `available_size_in_gbs` - Available Capacity 
	* `total_size_in_gbs` - Total Capacity 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Database Storage Vault.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `state` - The current state of the Exadata Database Storage Vault.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `time_created` - The date and time that the Exadata Database Storage Vault was created.
* `time_zone` - The time zone that you want to use for the Exadata Database Storage Vault. For details, see [Time Zones](https://docs.cloud.oracle.com/iaas/Content/Database/References/timezones.htm). 
* `vm_cluster_count` - The number of Exadata VM clusters used the Exadata Database Storage Vault. 
* `vm_cluster_ids` - The List of Exadata VM cluster on Exascale Infrastructure [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) **Note:** If Exadata Database Storage Vault is not used for any Exadata VM cluster on Exascale Infrastructure, this list is empty. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Exascale Db Storage Vault
	* `update` - (Defaults to 20 minutes), when updating the Exascale Db Storage Vault
	* `delete` - (Defaults to 20 minutes), when destroying the Exascale Db Storage Vault


## Import

ExascaleDbStorageVaults can be imported using the `id`, e.g.

```
$ terraform import oci_database_exascale_db_storage_vault.test_exascale_db_storage_vault "id"
```

