---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_awr_hubs"
sidebar_current: "docs-oci-datasource-opsi-awr_hubs"
description: |-
  Provides the list of Awr Hubs in Oracle Cloud Infrastructure Opsi service
---

# Data Source: oci_opsi_awr_hubs
This data source provides the list of Awr Hubs in Oracle Cloud Infrastructure Opsi service.

Gets a list of AWR hubs. Either compartmentId or id must be specified. All these resources are expected to be in root compartment. 


## Example Usage

```hcl
data "oci_opsi_awr_hubs" "test_awr_hubs" {
	#Required
	operations_insights_warehouse_id = oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id

	#Optional
	compartment_id = var.compartment_id
	display_name = var.awr_hub_display_name
	id = var.awr_hub_id
	state = var.awr_hub_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the entire display name.
* `id` - (Optional) Unique Awr Hub identifier
* `operations_insights_warehouse_id` - (Required) Unique Operations Insights Warehouse identifier
* `state` - (Optional) Lifecycle states


## Attributes Reference

The following attributes are exported:

* `awr_hub_summary_collection` - The list of awr_hub_summary_collection.

### AwrHub Reference

The following attributes are exported:

* `awr_mailbox_url` - Mailbox URL required for AWR hub and AWR source setup.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - User-friedly name of AWR Hub that does not have to be unique.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `hub_dst_timezone_version` - Dst Time Zone Version of the AWR Hub
* `id` - AWR Hub OCID
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `object_storage_bucket_name` - Object Storage Bucket Name
* `operations_insights_warehouse_id` - OPSI Warehouse OCID
* `state` - Possible lifecycle states
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time at which the resource was first created. An RFC3339 formatted datetime string
* `time_updated` - The time at which the resource was last updated. An RFC3339 formatted datetime string

