---
subcategory: "Em Warehouse"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_em_warehouse_em_warehouses"
sidebar_current: "docs-oci-datasource-em_warehouse-em_warehouses"
description: |-
  Provides the list of Em Warehouses in Oracle Cloud Infrastructure Em Warehouse service
---

# Data Source: oci_em_warehouse_em_warehouses
This data source provides the list of Em Warehouses in Oracle Cloud Infrastructure Em Warehouse service.

Returns a list of EmWarehouses.


## Example Usage

```hcl
data "oci_em_warehouse_em_warehouses" "test_em_warehouses" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.em_warehouse_display_name
	id = var.em_warehouse_id
	operations_insights_warehouse_id = oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id
	state = var.em_warehouse_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) unique EmWarehouse identifier
* `operations_insights_warehouse_id` - (Optional) unique operationsInsightsWarehouseId identifier
* `state` - (Optional) A filter to return only resources their lifecycleState matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `em_warehouse_collection` - The list of em_warehouse_collection.

### EmWarehouse Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - EmWarehouse Identifier, can be renamed
* `em_bridge_id` - EMBridge Identifier
* `em_warehouse_type` - Type of the EmWarehouse.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation
* `latest_etl_run_message` - Data Flow Run Status Message
* `latest_etl_run_status` - Data Flow Run Status
* `latest_etl_run_time` - Data Flow Run Total Time
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `operations_insights_warehouse_id` - operations Insights Warehouse Identifier
* `state` - The current state of the EmWarehouse.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the the EmWarehouse was created. An RFC3339 formatted datetime string
* `time_updated` - The time the EmWarehouse was updated. An RFC3339 formatted datetime string

