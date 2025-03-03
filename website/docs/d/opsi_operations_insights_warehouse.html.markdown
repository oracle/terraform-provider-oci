---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_operations_insights_warehouse"
sidebar_current: "docs-oci-datasource-opsi-operations_insights_warehouse"
description: |-
  Provides details about a specific Operations Insights Warehouse in Oracle Cloud Infrastructure Opsi service
---

# Data Source: oci_opsi_operations_insights_warehouse
This data source provides details about a specific Operations Insights Warehouse resource in Oracle Cloud Infrastructure Opsi service.

Gets details of an Ops Insights Warehouse.
There is only expected to be 1 warehouse per tenant. The warehouse is expected to be in the root compartment.


## Example Usage

```hcl
data "oci_opsi_operations_insights_warehouse" "test_operations_insights_warehouse" {
	#Required
	operations_insights_warehouse_id = oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id
}
```

## Argument Reference

The following arguments are supported:

* `operations_insights_warehouse_id` - (Required) Unique Ops Insights Warehouse identifier


## Attributes Reference

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

