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
	attached_shape_attributes = var.exascale_db_storage_vault_attached_shape_attributes
	attached_shape_attributes_not_equal_to = var.exascale_db_storage_vault_attached_shape_attributes_not_equal_to
	cluster_placement_group_id = oci_cluster_placement_groups_cluster_placement_group.test_cluster_placement_group.id
	display_name = var.exascale_db_storage_vault_display_name
	exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
	state = var.exascale_db_storage_vault_state
	vm_cluster_count_greater_than_or_equal_to = var.exascale_db_storage_vault_vm_cluster_count_greater_than_or_equal_to
	vm_cluster_count_less_than_or_equal_to = var.exascale_db_storage_vault_vm_cluster_count_less_than_or_equal_to
}
```

## Argument Reference

The following arguments are supported:

* `attached_shape_attributes` - (Optional) A filter to return only Exadata Database Storage Vaults which match the given attachedShapeAttributes or has null attachedShapeAttributes
* `attached_shape_attributes_not_equal_to` - (Optional) A filter to return only Exadata Database Storage Vaults which do not match the given attachedShapeAttributes
* `cluster_placement_group_id` - (Optional) A filter to return only resources that match the given cluster placement group ID exactly.
* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `exadata_infrastructure_id` - (Optional) A filter to return only list of Vaults that are linked to the exadata infrastructure Id.
* `state` - (Optional) A filter to return only Exadata Database Storage Vaults that match the given lifecycle state exactly.
* `vm_cluster_count_greater_than_or_equal_to` - (Optional) A filter to return only Exadata Database Storage Vaults with associated Exadata VM Clusters greater than or equal to the given count
* `vm_cluster_count_less_than_or_equal_to` - (Optional) A filter to return only Exadata Database Storage Vaults with associated Exadata VM Clusters less than or equal to the given count


## Attributes Reference

The following attributes are exported:

* `exascale_db_storage_vaults` - The list of exascale_db_storage_vaults.

### ExascaleDbStorageVault Reference

The following attributes are exported:

* `additional_flash_cache_in_percent` - The size of additional Flash Cache in percentage of High Capacity database storage.
* `attached_shape_attributes` - The shapeAttribute of the Exadata VM cluster(s) associated with the Exadata Database Storage Vault.
* `autoscale_limit_in_gbs` - Maximum limit storage size in gigabytes, that is applicable for the Database Storage Vault.
* `availability_domain` - The name of the availability domain in which the Exadata Database Storage Vault is located.
* `cluster_placement_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cluster placement group of the Exadata Infrastructure or Db System.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `description` - Exadata Database Storage Vault description.
* `display_name` - The user-friendly name for the Exadata Database Storage Vault. The name does not need to be unique.
* `exadata_infrastructure_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `high_capacity_database_storage` - Exadata Database Storage Details 
	* `available_size_in_gbs` - Available Capacity 
	* `total_size_in_gbs` - Total Capacity 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Database Storage Vault.
* `is_autoscale_enabled` - Indicates if autoscale feature is enabled for the Database Storage Vault. The default value is `FALSE`.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `state` - The current state of the Exadata Database Storage Vault.
* `subscription_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `time_created` - The date and time that the Exadata Database Storage Vault was created.
* `time_zone` - The time zone that you want to use for the Exadata Database Storage Vault. For details, see [Time Zones](https://docs.cloud.oracle.com/iaas/Content/Database/References/timezones.htm). 
* `vm_cluster_count` - The number of Exadata VM clusters used the Exadata Database Storage Vault. 
* `vm_cluster_ids` - The List of Exadata VM cluster on Exascale Infrastructure [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) **Note:** If Exadata Database Storage Vault is not used for any Exadata VM cluster on Exascale Infrastructure, this list is empty. 

