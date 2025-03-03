---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_operations_insights_warehouses"
sidebar_current: "docs-oci-datasource-opsi-operations_insights_warehouses"
description: |-
  Provides the list of Operations Insights Warehouses in Oracle Cloud Infrastructure Opsi service
---

# Data Source: oci_opsi_operations_insights_warehouses
This data source provides the list of Operations Insights Warehouses in Oracle Cloud Infrastructure Opsi service.

Gets a list of Ops Insights warehouses. Either compartmentId or id must be specified.
There is only expected to be 1 warehouse per tenant. The warehouse is expected to be in the root compartment.


## Example Usage

```hcl
data "oci_opsi_operations_insights_warehouses" "test_operations_insights_warehouses" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.operations_insights_warehouse_display_name
	id = var.operations_insights_warehouse_id
	state = var.operations_insights_warehouse_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the entire display name.
* `id` - (Optional) Unique Ops Insights Warehouse identifier
* `state` - (Optional) Lifecycle states


## Attributes Reference

The following attributes are exported:

* `operations_insights_warehouse_summary_collection` - The list of operations_insights_warehouse_summary_collection.

### OperationsInsightsWarehouse Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compute_model` - The compute model for the OPSI warehouse ADW (OCPU or ECPU)
* `cpu_allocated` - Number of CPUs allocated to OPSI Warehouse ADW. 
* `cpu_used` - Number of OCPUs used by OPSI Warehouse ADW. Can be fractional. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - User-friedly name of Ops Insights Warehouse that does not have to be unique.
* `dynamic_group_id` - OCID of the dynamic group created for the warehouse
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - OPSI Warehouse OCID
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `operations_insights_tenancy_id` - Tenancy Identifier of Ops Insights service
* `state` - Possible lifecycle states
* `storage_allocated_in_gbs` - Storage allocated to OPSI Warehouse ADW. 
* `storage_used_in_gbs` - Storage by OPSI Warehouse ADW in GB. 
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time at which the resource was first created. An RFC3339 formatted datetime string
* `time_last_wallet_rotated` - The time at which the ADW wallet was last rotated for the Ops Insights Warehouse. An RFC3339 formatted datetime string
* `time_updated` - The time at which the resource was last updated. An RFC3339 formatted datetime string

