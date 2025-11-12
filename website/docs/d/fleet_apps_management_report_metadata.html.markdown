---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_report_metadata"
sidebar_current: "docs-oci-datasource-fleet_apps_management-report_metadata"
description: |-
  Provides the list of Report Metadata in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_report_metadata
This data source provides the list of Report Metadata in Oracle Cloud Infrastructure Fleet Apps Management service.

Returns a list of all the report metadata.


## Example Usage

```hcl
data "oci_fleet_apps_management_report_metadata" "test_report_metadata" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	report_name = var.report_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `report_name` - (Optional) A filter to return data for given report name.


## Attributes Reference

The following attributes are exported:

* `report_metadata_collection` - The list of report_metadata_collection.

### ReportMetadata Reference

The following attributes are exported:

* `items` - List of ReportMetadata.
	* `column_metadata` - Column Metadata.
		* `description` - Column description
		* `name` - Column name.
		* `type` - Column value type. 
	* `default_order_clause` - default order clause for reports.
		* `sort_by` - Column to sort by.
		* `sort_order` - Sort direction either ASC or DESC.
	* `description` - Description of report.
	* `filters` - metricMetadata.
		* `description` - Filter description.
		* `name` - Filter Name.
		* `value_source` - Filter value source.
	* `metric` - Metric Name.
	* `name` - Name of Report.

