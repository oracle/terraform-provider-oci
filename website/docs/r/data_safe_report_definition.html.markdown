---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_report_definition"
sidebar_current: "docs-oci-resource-data_safe-report_definition"
description: |-
  Provides the Report Definition resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_report_definition
This resource provides the Report Definition resource in Oracle Cloud Infrastructure Data Safe service.

Creates a new report definition with parameters specified in the body. The report definition is stored in the specified compartment.


## Example Usage

```hcl
resource "oci_data_safe_report_definition" "test_report_definition" {
	#Required
	column_filters {
		#Required
		expressions = var.report_definition_column_filters_expressions
		field_name = var.report_definition_column_filters_field_name
		is_enabled = var.report_definition_column_filters_is_enabled
		is_hidden = var.report_definition_column_filters_is_hidden
		operator = var.report_definition_column_filters_operator
	}
	column_info {
		#Required
		display_name = var.report_definition_column_info_display_name
		display_order = var.report_definition_column_info_display_order
		field_name = var.report_definition_column_info_field_name
		is_hidden = var.report_definition_column_info_is_hidden

		#Optional
		data_type = var.report_definition_column_info_data_type
	}
	column_sortings {
		#Required
		field_name = var.report_definition_column_sortings_field_name
		is_ascending = var.report_definition_column_sortings_is_ascending
		sorting_order = var.report_definition_column_sortings_sorting_order
	}
	compartment_id = var.compartment_id
	display_name = var.report_definition_display_name
	parent_id = oci_data_safe_parent.test_parent.id
	summary {
		#Required
		display_order = var.report_definition_summary_display_order
		name = var.report_definition_summary_name

		#Optional
		count_of = var.report_definition_summary_count_of
		group_by_field_name = var.report_definition_summary_group_by_field_name
		is_hidden = var.report_definition_summary_is_hidden
		scim_filter = var.report_definition_summary_scim_filter
	}

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.report_definition_description
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `column_filters` - (Required) (Updatable) An array of column filter objects. A column Filter object stores all information about a column filter including field name, an operator, one or more expressions, if the filter is enabled, or if the filter is hidden.
	* `expressions` - (Required) (Updatable) An array of expressions based on the operator type. A filter may have one or more expressions.
	* `field_name` - (Required) (Updatable) Name of the column on which the filter must be applied.
	* `is_enabled` - (Required) (Updatable) Indicates whether the filter is enabled. Values can either be 'true' or 'false'.
	* `is_hidden` - (Required) (Updatable) Indicates whether the filter is hidden. Values can either be 'true' or 'false'.
	* `operator` - (Required) (Updatable) Specifies the type of operator that must be applied for example in, eq etc.
* `column_info` - (Required) (Updatable) An array of column objects in the order (left to right) displayed in the report. A column object stores all information about a column, including the name displayed on the UI, corresponding field name in the data source, data type of the column, and column visibility (if the column is visible to the user).
	* `data_type` - (Optional) (Updatable) Specifies the data type of the column.
	* `display_name` - (Required) (Updatable) Name of the column displayed on UI.
	* `display_order` - (Required) (Updatable) Specifies the display order of the column.
	* `field_name` - (Required) (Updatable) Specifies the corresponding field name in the data source.
	* `is_hidden` - (Required) (Updatable) Indicates if the column is hidden. Values can either be 'true' or 'false'.
* `column_sortings` - (Required) (Updatable) An array of column sorting objects. Each column sorting object stores the column name to be sorted and if the sorting is in ascending order; sorting is done by the first column in the array, then by the second column in the array, etc.
	* `field_name` - (Required) (Updatable) Name of the column that must be sorted.
	* `is_ascending` - (Required) (Updatable) Indicates if the column must be sorted in ascending order. Values can either be 'true' or 'false'.
	* `sorting_order` - (Required) (Updatable) Indicates the order at which column must be sorted.
* `compartment_id` - (Required) (Updatable) The OCID of the compartment containing the report definition.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) The description of the report definition.
* `display_name` - (Required) (Updatable) Specifies the name of the report definition.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `parent_id` - (Required) The OCID of the parent report definition.
* `summary` - (Required) (Updatable) An array of report summary objects in the order (left to right)  displayed in the report.  A  report summary object stores all information about summary of report to be displayed, including the name displayed on UI, the display order, corresponding group by and count of values, summary visibility (if the summary is visible to user).
	* `count_of` - (Optional) (Updatable) Name of the key or count of object.
	* `display_order` - (Required) (Updatable) Specifies the order in which the summary must be displayed.
	* `group_by_field_name` - (Optional) (Updatable) A comma-delimited string that specifies the names of the fields by which the records must be aggregated to get the summary.
	* `is_hidden` - (Optional) (Updatable) Indicates if the summary is hidden. Values can either be 'true' or 'false'.
	* `name` - (Required) (Updatable) Name of the report summary.
	* `scim_filter` - (Optional) (Updatable) Additional scim filters used to get the specific summary.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `category` - Specifies the name of the category that this report belongs to.
* `column_filters` - An array of columnFilter objects. A columnFilter object stores all information about a column filter including field name, an operator, one or more expressions, if the filter is enabled, or if the filter is hidden.
	* `expressions` - An array of expressions based on the operator type. A filter may have one or more expressions.
	* `field_name` - Name of the column on which the filter must be applied.
	* `is_enabled` - Indicates whether the filter is enabled. Values can either be 'true' or 'false'.
	* `is_hidden` - Indicates whether the filter is hidden. Values can either be 'true' or 'false'.
	* `operator` - Specifies the type of operator that must be applied for example in, eq etc.
* `column_info` - An array of column objects in the order (left to right) displayed in the report. A column object stores all information about a column, including the name displayed on the UI, corresponding field name in the data source, data type of the column, and column visibility (if the column is visible to the user).
	* `data_type` - Specifies the data type of the column.
	* `display_name` - Name of the column displayed on UI.
	* `display_order` - Specifies the display order of the column.
	* `field_name` - Specifies the corresponding field name in the data source.
	* `is_hidden` - Indicates if the column is hidden. Values can either be 'true' or 'false'.
* `column_sortings` - An array of column sorting objects. Each column sorting object stores the column name to be sorted and if the sorting is in ascending order; sorting is done by the first column in the array, then by the second column in the array, etc.
	* `field_name` - Name of the column that must be sorted.
	* `is_ascending` - Indicates if the column must be sorted in ascending order. Values can either be 'true' or 'false'.
	* `sorting_order` - Indicates the order at which column must be sorted.
* `compartment_id` - The OCID of the compartment containing the report definition.
* `compliance_standards` - The list of the data protection regulations/standards used in the report that will help demonstrate compliance.
* `data_source` - Specifies the name of a resource that provides data for the report. For example alerts, events.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - A description of the report definition.
* `display_name` - Name of the report definition.
* `display_order` - Specifies how the report definitions are ordered in the display.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the report definition.
* `is_seeded` - Signifies whether the definition is seeded or user defined. Values can either be 'true' or 'false'.
* `lifecycle_details` - Details about the current state of the report definition in Data Safe.
* `parent_id` - The OCID of the parent report definition. In the case of seeded report definition, this is same as definition OCID.
* `record_time_span` - The time span for the records in the report to be scheduled. <period-value><period> Allowed period strings - "H","D","M","Y" Each of the above fields potentially introduce constraints. A workRequest is created only when period-value satisfies all the constraints. Constraints introduced: 1. period = H (The allowed range for period-value is [1, 23]) 2. period = D (The allowed range for period-value is [1, 30]) 3. period = M (The allowed range for period-value is [1, 11]) 4. period = Y (The minimum period-value is 1) 
* `schedule` - The schedule to generate the report periodically in the specified format: <version-string>;<version-specific-schedule>

	Allowed version strings - "v1" v1's version specific schedule -<ss> <mm> <hh> <day-of-week> <day-of-month> Each of the above fields potentially introduce constraints. A workrequest is created only when clock time satisfies all the constraints. Constraints introduced: 1. seconds = <ss> (So, the allowed range for <ss> is [0, 59]) 2. minutes = <mm> (So, the allowed range for <mm> is [0, 59]) 3. hours = <hh> (So, the allowed range for <hh> is [0, 23]) 4. <day-of-week> can be either '*' (without quotes or a number between 1(Monday) and 7(Sunday)) No constraint introduced when it is '*'. When not, day of week must equal the given value 5. <day-of-month> can be either '*' (without quotes or a number between 1 and 28) No constraint introduced when it is '*'. When not, day of month must equal the given value 
* `scheduled_report_compartment_id` - The OCID of the compartment in which the scheduled resource will be created. 
* `scheduled_report_mime_type` - Specifies the format of the report ( either .xls or .pdf or .json)
* `scheduled_report_name` - The name of the report to be scheduled.
* `scheduled_report_row_limit` - Specifies the limit on the number of rows in the report.
* `scim_filter` - Additional SCIM filters used to define the report.
* `state` - The current state of the report.
* `summary` - An array of report summary objects in the order (left to right)  displayed in the report.  A  report summary object stores all information about summary of report to be displayed, including the name displayed on UI, the display order, corresponding group by and count of values, summary visibility (if the summary is visible to user).
	* `count_of` - Name of the key or count of object.
	* `display_order` - Specifies the order in which the summary must be displayed.
	* `group_by_field_name` - A comma-delimited string that specifies the names of the fields by which the records must be aggregated to get the summary.
	* `is_hidden` - Indicates if the summary is hidden. Values can either be 'true' or 'false'.
	* `name` - Name of the report summary.
	* `scim_filter` - Additional scim filters used to get the specific summary.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Specifies the date and time the report definition was created.
* `time_updated` - The date and time the report definition was updated.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Report Definition
	* `update` - (Defaults to 20 minutes), when updating the Report Definition
	* `delete` - (Defaults to 20 minutes), when destroying the Report Definition


## Import

ReportDefinitions can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_report_definition.test_report_definition "id"
```

