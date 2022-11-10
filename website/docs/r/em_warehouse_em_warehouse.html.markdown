---
subcategory: "Em Warehouse"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_em_warehouse_em_warehouse"
sidebar_current: "docs-oci-resource-em_warehouse-em_warehouse"
description: |-
  Provides the Em Warehouse resource in Oracle Cloud Infrastructure Em Warehouse service
---

# oci_em_warehouse_em_warehouse
This resource provides the Em Warehouse resource in Oracle Cloud Infrastructure Em Warehouse service.

Creates a new EmWarehouse.


## Example Usage

```hcl
resource "oci_em_warehouse_em_warehouse" "test_em_warehouse" {
	#Required
	compartment_id = var.compartment_id
	em_bridge_id = oci_em_warehouse_em_bridge.test_em_bridge.id
	operations_insights_warehouse_id = oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	display_name = var.em_warehouse_display_name
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment Identifier
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Optional) EmWarehouse Identifier
* `em_bridge_id` - (Required) (Updatable) EMBridge Identifier
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `operations_insights_warehouse_id` - (Required) operations Insights Warehouse Identifier


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Em Warehouse
	* `update` - (Defaults to 20 minutes), when updating the Em Warehouse
	* `delete` - (Defaults to 20 minutes), when destroying the Em Warehouse


## Import

EmWarehouses can be imported using the `id`, e.g.

```
$ terraform import oci_em_warehouse_em_warehouse.test_em_warehouse "id"
```

