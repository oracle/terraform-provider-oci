---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_flex_components"
sidebar_current: "docs-oci-datasource-database-flex_components"
description: |-
  Provides the list of Flex Components in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_flex_components
This data source provides the list of Flex Components in Oracle Cloud Infrastructure Database service.

Gets a list of the flex components that can be used to launch a new DB system. The flex component determines resources to allocate to the DB system - Database Servers and Storage Servers.

## Example Usage

```hcl
data "oci_database_flex_components" "test_flex_components" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	name = var.flex_component_name
	shape = var.flex_component_shape
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `name` - (Optional) A filter to return only resources that match the entire name given. The match is not case sensitive.
* `shape` - (Optional) A filter to return only resources that belong to the entire shape name given. The match is not case sensitive.


## Attributes Reference

The following attributes are exported:

* `flex_component_collection` - The list of flex_component_collection.

### FlexComponent Reference

The following attributes are exported:

* `items` - 
	* `available_core_count` - The maximum number of CPU cores that can ben enabled on the DB Server for this Flex Component.
	* `available_db_storage_in_gbs` - The maximum  storage that can be enabled on the Storage Server for this Flex Component.
	* `available_local_storage_in_gbs` - The maximum local storage that can be enabled on the DB Server for this Flex Component.
	* `available_memory_in_gbs` - The maximum memory size that can be enabled on the DB Server for this Flex Component.
	* `compute_model` - The compute model of the DB Server for this Flex Component.
	* `description_summary` - The description summary for this Flex Component.
	* `hardware_type` - The hardware type of the DB (Compute) or Storage (Cell) Server for this Flex Component.
	* `minimum_core_count` - The minimum number of CPU cores that can be enabled on the DB Server for this Flex Component.
	* `name` - The name of the Flex Component used for the DB system.
	* `runtime_minimum_core_count` - The runtime minimum number of CPU cores that can be enabled for this Flex Component.
	* `shape` - The name of the DB system shape for this Flex Component.

