---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_operations_insights_warehouse_users"
sidebar_current: "docs-oci-datasource-opsi-operations_insights_warehouse_users"
description: |-
  Provides the list of Operations Insights Warehouse Users in Oracle Cloud Infrastructure Opsi service
---

# Data Source: oci_opsi_operations_insights_warehouse_users
This data source provides the list of Operations Insights Warehouse Users in Oracle Cloud Infrastructure Opsi service.

Gets a list of Operations Insights Warehouse users. Either compartmentId or id must be specified. All these resources are expected to be in root compartment.


## Example Usage

```hcl
data "oci_opsi_operations_insights_warehouse_users" "test_operations_insights_warehouse_users" {
	#Required
	operations_insights_warehouse_id = oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id

	#Optional
	compartment_id = var.compartment_id
	display_name = var.operations_insights_warehouse_user_display_name
	id = var.operations_insights_warehouse_user_id
	state = var.operations_insights_warehouse_user_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the entire display name.
* `id` - (Optional) Unique Operations Insights Warehouse User identifier
* `operations_insights_warehouse_id` - (Required) Unique Operations Insights Warehouse identifier
* `state` - (Optional) Lifecycle states


## Attributes Reference

The following attributes are exported:

* `operations_insights_warehouse_user_summary_collection` - The list of operations_insights_warehouse_user_summary_collection.

### OperationsInsightsWarehouseUser Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `connection_password` - User provided connection password for the AWR Data,  Enterprise Manager Data and Operations Insights OPSI Hub.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Hub User OCID
* `is_awr_data_access` - Indicate whether user has access to AWR data.
* `is_em_data_access` - Indicate whether user has access to EM data.
* `is_opsi_data_access` - Indicate whether user has access to OPSI data.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `name` - Username for schema which would have access to AWR Data,  Enterprise Manager Data and Operations Insights OPSI Hub.
* `operations_insights_warehouse_id` - OPSI Warehouse OCID
* `state` - Possible lifecycle states
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time at which the resource was first created. An RFC3339 formatted datetime string
* `time_updated` - The time at which the resource was last updated. An RFC3339 formatted datetime string

