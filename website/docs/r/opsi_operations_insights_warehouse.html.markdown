---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_operations_insights_warehouse"
sidebar_current: "docs-oci-resource-opsi-operations_insights_warehouse"
description: |-
  Provides the Operations Insights Warehouse resource in Oracle Cloud Infrastructure Opsi service
---

# oci_opsi_operations_insights_warehouse
This resource provides the Operations Insights Warehouse resource in Oracle Cloud Infrastructure Opsi service.

Create a Ops Insights Warehouse resource for the tenant in Ops Insights. New ADW will be provisioned for this tenant.
There is only expected to be 1 warehouse per tenant. The warehouse is expected to be in the root compartment. If the 'opsi-warehouse-type'
header is passed to the API, a warehouse resource without ADW or Schema provisioning is created.


## Example Usage

```hcl
resource "oci_opsi_operations_insights_warehouse" "test_operations_insights_warehouse" {
	#Required
	compartment_id = var.compartment_id
	cpu_allocated = var.operations_insights_warehouse_cpu_allocated
	display_name = var.operations_insights_warehouse_display_name

	#Optional
	compute_model = var.operations_insights_warehouse_compute_model
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	storage_allocated_in_gbs = var.operations_insights_warehouse_storage_allocated_in_gbs
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compute_model` - (Optional) (Updatable) The compute model for the OPSI warehouse ADW (OCPU or ECPU)
* `cpu_allocated` - (Required) (Updatable) Number of CPUs allocated to OPSI Warehouse ADW. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) User-friedly name of Ops Insights Warehouse that does not have to be unique.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `storage_allocated_in_gbs` - (Optional) (Updatable) Storage allocated to OPSI Warehouse ADW. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Operations Insights Warehouse
	* `update` - (Defaults to 20 minutes), when updating the Operations Insights Warehouse
	* `delete` - (Defaults to 20 minutes), when destroying the Operations Insights Warehouse


## Import

OperationsInsightsWarehouses can be imported using the `id`, e.g.

```
$ terraform import oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse "id"
```

