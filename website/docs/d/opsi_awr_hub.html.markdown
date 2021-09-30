---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_awr_hub"
sidebar_current: "docs-oci-datasource-opsi-awr_hub"
description: |-
  Provides details about a specific Awr Hub in Oracle Cloud Infrastructure Opsi service
---

# Data Source: oci_opsi_awr_hub
This data source provides details about a specific Awr Hub resource in Oracle Cloud Infrastructure Opsi service.

Gets details of an AWR hub.

## Example Usage

```hcl
data "oci_opsi_awr_hub" "test_awr_hub" {
	#Required
	awr_hub_id = oci_opsi_awr_hub.test_awr_hub.id
}
```

## Argument Reference

The following arguments are supported:

* `awr_hub_id` - (Required) Unique Awr Hub identifier


## Attributes Reference

The following attributes are exported:

* `awr_mailbox_url` - Mailbox URL required for AWR hub and AWR source setup.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - User-friedly name of AWR Hub that does not have to be unique.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - AWR Hub OCID
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `object_storage_bucket_name` - Object Storage Bucket Name
* `operations_insights_warehouse_id` - OPSI Warehouse OCID
* `state` - Possible lifecycle states
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time at which the resource was first created. An RFC3339 formatted datetime string
* `time_updated` - The time at which the resource was last updated. An RFC3339 formatted datetime string

