---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_exascale_db_storage_vault"
sidebar_current: "docs-oci-datasource-database-exascale_db_storage_vault"
description: |-
  Provides details about a specific Exascale Db Storage Vault in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_exascale_db_storage_vault
This data source provides details about a specific Exascale Db Storage Vault resource in Oracle Cloud Infrastructure Database service.

Gets information about the specified Exadata Database Storage Vaults in the specified compartment.


## Example Usage

```hcl
data "oci_database_exascale_db_storage_vault" "test_exascale_db_storage_vault" {
	#Required
	exascale_db_storage_vault_id = oci_database_exascale_db_storage_vault.test_exascale_db_storage_vault.id
}
```

## Argument Reference

The following arguments are supported:

* `exascale_db_storage_vault_id` - (Required) The Exadata Database Storage Vault [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


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

