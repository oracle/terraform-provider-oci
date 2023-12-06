---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_operations_insights_warehouse_user"
sidebar_current: "docs-oci-datasource-opsi-operations_insights_warehouse_user"
description: |-
  Provides details about a specific Operations Insights Warehouse User in Oracle Cloud Infrastructure Opsi service
---

# Data Source: oci_opsi_operations_insights_warehouse_user
This data source provides details about a specific Operations Insights Warehouse User resource in Oracle Cloud Infrastructure Opsi service.

Gets details of an Operations Insights Warehouse User.

## Example Usage

```hcl
data "oci_opsi_operations_insights_warehouse_user" "test_operations_insights_warehouse_user" {
	#Required
	operations_insights_warehouse_user_id = oci_opsi_operations_insights_warehouse_user.test_operations_insights_warehouse_user.id
}
```

## Argument Reference

The following arguments are supported:

* `operations_insights_warehouse_user_id` - (Required) Unique Operations Insights Warehouse User identifier


## Attributes Reference

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

