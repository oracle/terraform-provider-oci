---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_operations_insights_warehouse_user"
sidebar_current: "docs-oci-resource-opsi-operations_insights_warehouse_user"
description: |-
  Provides the Operations Insights Warehouse User resource in Oracle Cloud Infrastructure Opsi service
---

# oci_opsi_operations_insights_warehouse_user
This resource provides the Operations Insights Warehouse User resource in Oracle Cloud Infrastructure Opsi service.

Create a Operations Insights Warehouse user resource for the tenant in Operations Insights.
This resource will be created in root compartment.


## Example Usage

```hcl
resource "oci_opsi_operations_insights_warehouse_user" "test_operations_insights_warehouse_user" {
	#Required
	compartment_id = var.compartment_id
	connection_password = var.operations_insights_warehouse_user_connection_password
	is_awr_data_access = var.operations_insights_warehouse_user_is_awr_data_access
	name = var.operations_insights_warehouse_user_name
	operations_insights_warehouse_id = oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	is_em_data_access = var.operations_insights_warehouse_user_is_em_data_access
	is_opsi_data_access = var.operations_insights_warehouse_user_is_opsi_data_access
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `connection_password` - (Required) (Updatable) User provided connection password for the AWR Data,  Enterprise Manager Data and Ops Insights OPSI Hub.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `is_awr_data_access` - (Required) (Updatable) Indicate whether user has access to AWR data.
* `is_em_data_access` - (Optional) (Updatable) Indicate whether user has access to EM data.
* `is_opsi_data_access` - (Optional) (Updatable) Indicate whether user has access to OPSI data.
* `name` - (Required) Username for schema which would have access to AWR Data,  Enterprise Manager Data and Ops Insights OPSI Hub.
* `operations_insights_warehouse_id` - (Required) OPSI Warehouse OCID


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `connection_password` - User provided connection password for the AWR Data,  Enterprise Manager Data and Ops Insights OPSI Hub.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Hub User OCID
* `is_awr_data_access` - Indicate whether user has access to AWR data.
* `is_em_data_access` - Indicate whether user has access to EM data.
* `is_opsi_data_access` - Indicate whether user has access to OPSI data.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `name` - Username for schema which would have access to AWR Data,  Enterprise Manager Data and Ops Insights OPSI Hub.
* `operations_insights_warehouse_id` - OPSI Warehouse OCID
* `state` - Possible lifecycle states
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time at which the resource was first created. An RFC3339 formatted datetime string
* `time_updated` - The time at which the resource was last updated. An RFC3339 formatted datetime string

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Operations Insights Warehouse User
	* `update` - (Defaults to 20 minutes), when updating the Operations Insights Warehouse User
	* `delete` - (Defaults to 20 minutes), when destroying the Operations Insights Warehouse User


## Import

OperationsInsightsWarehouseUsers can be imported using the `id`, e.g.

```
$ terraform import oci_opsi_operations_insights_warehouse_user.test_operations_insights_warehouse_user "id"
```

