---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_awr_hub_sources"
sidebar_current: "docs-oci-datasource-opsi-awr_hub_sources"
description: |-
  Provides the list of Awr Hub Sources in Oracle Cloud Infrastructure Opsi service
---

# Data Source: oci_opsi_awr_hub_sources
This data source provides the list of Awr Hub Sources in Oracle Cloud Infrastructure Opsi service.

Gets a list of Awr Hub source objects.

## Example Usage

```hcl
data "oci_opsi_awr_hub_sources" "test_awr_hub_sources" {
	#Required
	awr_hub_id = oci_opsi_awr_hub.test_awr_hub.id

	#Optional
	awr_hub_source_id = oci_opsi_awr_hub_source.test_awr_hub_source.id
	compartment_id = var.compartment_id
	name = var.awr_hub_source_name
	source_type = var.awr_hub_source_source_type
	state = var.awr_hub_source_state
	status = var.awr_hub_source_status
}
```

## Argument Reference

The following arguments are supported:

* `awr_hub_id` - (Required) Unique Awr Hub identifier
* `awr_hub_source_id` - (Optional) Awr Hub source identifier
* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `name` - (Optional) Awr Hub source database name
* `source_type` - (Optional) Filter by one or more database type. Possible values are ADW-S, ATP-S, ADW-D, ATP-D, EXTERNAL-PDB, EXTERNAL-NONCDB. 
* `state` - (Optional) Lifecycle states
* `status` - (Optional) Resource Status


## Attributes Reference

The following attributes are exported:

* `awr_hub_source_summary_collection` - The list of awr_hub_source_summary_collection.

### AwrHubSource Reference

The following attributes are exported:

* `associated_opsi_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database id.
* `associated_resource_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database id.
* `awr_hub_id` - AWR Hub OCID
* `awr_hub_opsi_source_id` - The shorted string of the Awr Hub source database identifier.
* `awr_source_database_id` - DatabaseId of the Source database for which AWR Data will be uploaded to AWR Hub.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `hours_since_last_import` - Number of hours since last AWR snapshots import happened from the Source database.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Awr Hub source database.
* `is_registered_with_awr_hub` - This is `true` if the source databse is registered with a Awr Hub, otherwise `false`
* `max_snapshot_identifier` - The maximum snapshot identifier of the source database for which AWR data is uploaded to AWR Hub.
* `min_snapshot_identifier` - The minimum snapshot identifier of the source database for which AWR data is uploaded to AWR Hub.
* `name` - The name of the Awr Hub source database.
* `source_mail_box_url` - Opsi Mailbox URL based on the Awr Hub and Awr Hub source.
* `state` - the current state of the source database
* `status` - Indicates the status of a source database in Operations Insights
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time at which the resource was first created. An RFC3339 formatted datetime string
* `time_first_snapshot_generated` - The time at which the earliest snapshot was generated in the source database for which data is uploaded to AWR Hub. An RFC3339 formatted datetime string
* `time_last_snapshot_generated` - The time at which the latest snapshot was generated in the source database for which data is uploaded to AWR Hub. An RFC3339 formatted datetime string
* `time_updated` - The time at which the resource was last updated. An RFC3339 formatted datetime string
* `type` - source type of the database

