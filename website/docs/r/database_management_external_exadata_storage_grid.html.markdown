---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_exadata_storage_grid"
sidebar_current: "docs-oci-resource-database_management-external_exadata_storage_grid"
description: |-
  Provides the External Exadata Storage Grid resource in Oracle Cloud Infrastructure Database Management service
---

# oci_database_management_external_exadata_storage_grid
This resource provides the External Exadata Storage Grid resource in Oracle Cloud Infrastructure Database Management service.

Updates the Exadata storage server grid specified by exadataStorageGridId.


## Example Usage

```hcl
resource "oci_database_management_external_exadata_storage_grid" "test_external_exadata_storage_grid" {
	#Required
	external_exadata_storage_grid_id = oci_database_management_external_exadata_storage_grid.test_external_exadata_storage_grid.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `external_exadata_storage_grid_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata storage grid.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `additional_details` - The additional details of the resource defined in `{"key": "value"}` format. Example: `{"bar-key": "value"}` 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The name of the Exadata resource. English letters, numbers, "-", "_" and "." only.
* `exadata_infrastructure_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata resource.
* `internal_id` - The internal ID of the Exadata resource.
* `lifecycle_details` - The details of the lifecycle state of the Exadata resource.
* `resource_type` - The type of Exadata resource.
* `server_count` - The number of Exadata storage servers in the Exadata infrastructure.
* `state` - The current lifecycle state of the database resource.
* `status` - The status of the Exadata resource.
* `storage_servers` - A list of monitored Exadata storage servers.
	* `additional_details` - The additional details of the resource defined in `{"key": "value"}` format. Example: `{"bar-key": "value"}` 
	* `connector_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the connector.
	* `cpu_count` - The CPU count of the Exadata storage server.
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
	* `display_name` - The name of the Exadata resource. English letters, numbers, "-", "_" and "." only.
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata resource.
	* `internal_id` - The internal ID of the Exadata resource.
	* `ip_address` - The IP address of the Exadata storage server.
	* `lifecycle_details` - The details of the lifecycle state of the Exadata resource.
	* `make_model` - The make model of the Exadata storage server.
	* `max_flash_disk_iops` - The maximum flash disk IO operations per second of the Exadata storage server.
	* `max_flash_disk_throughput` - The maximum flash disk IO throughput in MB/s of the Exadata storage server.
	* `max_hard_disk_iops` - The maximum hard disk IO operations per second of the Exadata storage server.
	* `max_hard_disk_throughput` - The maximum hard disk IO throughput in MB/s of the Exadata storage server.
	* `memory_gb` - The Exadata storage server memory size in GB.
	* `resource_type` - The type of Exadata resource.
	* `state` - The current lifecycle state of the database resource.
	* `status` - The status of the Exadata resource.
	* `time_created` - The timestamp of the creation of the Exadata resource.
	* `time_updated` - The timestamp of the last update of the Exadata resource.
	* `version` - The version of the Exadata resource.
* `time_created` - The timestamp of the creation of the Exadata resource.
* `time_updated` - The timestamp of the last update of the Exadata resource.
* `version` - The version of the Exadata resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the External Exadata Storage Grid
	* `update` - (Defaults to 20 minutes), when updating the External Exadata Storage Grid
	* `delete` - (Defaults to 20 minutes), when destroying the External Exadata Storage Grid


## Import

ExternalExadataStorageGrids can be imported using the `id`, e.g.

```
$ terraform import oci_database_management_external_exadata_storage_grid.test_external_exadata_storage_grid "id"
```

