---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_exascale_db_storage_vaults"
sidebar_current: "docs-oci-datasource-database-exascale_db_storage_vaults"
description: |-
  Provides the list of Exascale Db Storage Vaults in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_exascale_db_storage_vaults
This data source provides the list of Exascale Db Storage Vaults in Oracle Cloud Infrastructure Database service.

Gets a list of the Exadata Database Storage Vaults in the specified compartment.


## Example Usage

```hcl
data "oci_database_exascale_db_storage_vaults" "test_exascale_db_storage_vaults" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.exascale_db_storage_vault_display_name
	state = var.exascale_db_storage_vault_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `state` - (Optional) A filter to return only Exadata Database Storage Vaults that match the given lifecycle state exactly.


## Attributes Reference

The following attributes are exported:

* `exascale_db_storage_vaults` - The list of exascale_db_storage_vaults.

### ExascaleDbStorageVault Reference

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

