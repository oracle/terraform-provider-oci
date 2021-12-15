---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_awr_hub_awr_sources_summary"
sidebar_current: "docs-oci-datasource-opsi-awr_hub_awr_sources_summary"
description: |-
  Provides details about a specific Awr Hub Awr Sources Summary in Oracle Cloud Infrastructure Opsi service
---

# Data Source: oci_opsi_awr_hub_awr_sources_summary
This data source provides details about a specific Awr Hub Awr Sources Summary resource in Oracle Cloud Infrastructure Opsi service.

Gets a list of summary of AWR Sources.        


## Example Usage

```hcl
data "oci_opsi_awr_hub_awr_sources_summary" "test_awr_hub_awr_sources_summary" {
	#Required
	awr_hub_id = oci_opsi_awr_hub.test_awr_hub.id

	#Optional
	compartment_id = var.compartment_id
	name = var.awr_hub_awr_sources_summary_name
}
```

## Argument Reference

The following arguments are supported:

* `awr_hub_id` - (Required) Unique Awr Hub identifier
* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `name` - (Optional) Name for an Awr source database 


## Attributes Reference

The following attributes are exported:

* `items` - Array of AwrSource summary objects.
	* `awr_hub_id` - AWR Hub OCID
	* `awr_source_database_id` - DatabaseId of the Source database for which AWR Data will be uploaded to AWR Hub.
	* `hours_since_last_import` - Number of hours since last AWR snapshots import happened from the Source database.
	* `max_snapshot_identifier` - The maximum snapshot identifier of the source database for which AWR data is uploaded to AWR Hub.
	* `min_snapshot_identifier` - The minimum snapshot identifier of the source database for which AWR data is uploaded to AWR Hub.
	* `name` - Database name of the Source database for which AWR Data will be uploaded to AWR Hub.
	* `snapshots_uploaded` - Number of AWR snapshots uploaded from the Source database.
	* `time_first_snapshot_generated` - The time at which the earliest snapshot was generated in the source database for which data is uploaded to AWR Hub. An RFC3339 formatted datetime string
	* `time_last_snapshot_generated` - The time at which the latest snapshot was generated in the source database for which data is uploaded to AWR Hub. An RFC3339 formatted datetime string

