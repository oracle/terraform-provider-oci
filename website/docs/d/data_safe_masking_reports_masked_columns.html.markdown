---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_masking_reports_masked_columns"
sidebar_current: "docs-oci-datasource-data_safe-masking_reports_masked_columns"
description: |-
  Provides the list of Masking Reports Masked Columns in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_masking_reports_masked_columns
This data source provides the list of Masking Reports Masked Columns in Oracle Cloud Infrastructure Data Safe service.

Gets a list of masked columns present in the specified masking report and based on the specified query parameters.


## Example Usage

```hcl
data "oci_data_safe_masking_reports_masked_columns" "test_masking_reports_masked_columns" {
	#Required
	masking_report_id = oci_data_safe_masking_report.test_masking_report.id

	#Optional
	column_name = var.masking_reports_masked_column_column_name
	masking_column_group = var.masking_reports_masked_column_masking_column_group
	object = var.masking_reports_masked_column_object
	object_type = var.masking_reports_masked_column_object_type
	schema_name = var.masking_reports_masked_column_schema_name
	sensitive_type_id = oci_data_safe_sensitive_type.test_sensitive_type.id
}
```

## Argument Reference

The following arguments are supported:

* `column_name` - (Optional) A filter to return only a specific column based on column name.
* `masking_column_group` - (Optional) A filter to return only the resources that match the specified masking column group.
* `masking_report_id` - (Required) The OCID of the masking report.
* `object` - (Optional) A filter to return only items related to a specific object name.
* `object_type` - (Optional) A filter to return only items related to a specific object type.
* `schema_name` - (Optional) A filter to return only items related to specific schema name.
* `sensitive_type_id` - (Optional) A filter to return only items related to a specific sensitive type OCID.


## Attributes Reference

The following attributes are exported:

* `masked_column_collection` - The list of masked_column_collection.

### MaskingReportsMaskedColumn Reference

The following attributes are exported:

* `items` - An array of masking column summary objects.
	* `column_name` - The name of the masked column.
	* `key` - The unique key that identifies the masked column. It's numeric and unique within a masking policy.
	* `masking_column_group` - The masking group of the masked column.
	* `masking_format_used` - The masking format used for masking the column.
	* `object` - The name of the object (table or editioning view) that contains the masked column.
	* `object_type` - The type of the object (table or editioning view) that contains the masked column.
	* `parent_column_key` - The unique key that identifies the parent column of the masked column.
	* `schema_name` - The name of the schema that contains the masked column.
	* `sensitive_type_id` - The OCID of the sensitive type associated with the masked column.
	* `total_masked_values` - The total number of values masked in the column.

