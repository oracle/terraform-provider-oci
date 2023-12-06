---
subcategory: "Em Warehouse"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_em_warehouse_em_warehouse_etl_run"
sidebar_current: "docs-oci-datasource-em_warehouse-em_warehouse_etl_run"
description: |-
  Provides details about a specific Em Warehouse Etl Run in Oracle Cloud Infrastructure Em Warehouse service
---

# Data Source: oci_em_warehouse_em_warehouse_etl_run
This data source provides details about a specific Em Warehouse Etl Run resource in Oracle Cloud Infrastructure Em Warehouse service.

Gets a list of runs of an EmWarehouseResource by identifier

## Example Usage

```hcl
data "oci_em_warehouse_em_warehouse_etl_run" "test_em_warehouse_etl_run" {
	#Required
	em_warehouse_id = oci_em_warehouse_em_warehouse.test_em_warehouse.id

	#Optional
	compartment_id = var.compartment_id
	display_name = var.em_warehouse_etl_run_display_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `em_warehouse_id` - (Required) unique EmWarehouse identifier


## Attributes Reference

The following attributes are exported:

* `items` - List of runs
	* `compartment_id` - Compartment Identifier
	* `data_read_in_bytes` - Data read by the dataflow run
	* `data_written` - Data written by the dataflow run
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
	* `display_name` - The name of the ETLRun.
	* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
	* `lifecycle_details` - Details of the lifecycle state
	* `run_duration_in_milliseconds` - Dataflow run duration
	* `state` - The current state of the etlRun.
	* `time_created` - Time when the dataflow run was created
	* `time_updated` - Time when the dataflow run was updated

