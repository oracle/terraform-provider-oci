---
subcategory: "Em Warehouse"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_em_warehouse_em_warehouse"
sidebar_current: "docs-oci-datasource-em_warehouse-em_warehouse"
description: |-
  Provides details about a specific Em Warehouse in Oracle Cloud Infrastructure Em Warehouse service
---

# Data Source: oci_em_warehouse_em_warehouse
This data source provides details about a specific Em Warehouse resource in Oracle Cloud Infrastructure Em Warehouse service.

Gets a EmWarehouse by identifier

## Example Usage

```hcl
data "oci_em_warehouse_em_warehouse" "test_em_warehouse" {
	#Required
	em_warehouse_id = oci_em_warehouse_em_warehouse.test_em_warehouse.id
}
```

## Argument Reference

The following arguments are supported:

* `em_warehouse_id` - (Required) unique EmWarehouse identifier


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

