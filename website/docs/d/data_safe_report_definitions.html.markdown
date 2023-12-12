---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_report_definitions"
sidebar_current: "docs-oci-datasource-data_safe-report_definitions"
description: |-
  Provides the list of Report Definitions in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_report_definitions
This data source provides the list of Report Definitions in Oracle Cloud Infrastructure Data Safe service.

Gets a list of report definitions.
The ListReportDefinitions operation returns only the report definitions in the specified `compartmentId`.
It also returns the seeded report definitions which are available to all the compartments.


## Example Usage

```hcl
data "oci_data_safe_report_definitions" "test_report_definitions" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.report_definition_access_level
	category = var.report_definition_category
	compartment_id_in_subtree = var.report_definition_compartment_id_in_subtree
	data_source = var.report_definition_data_source
	display_name = var.report_definition_display_name
	is_seeded = var.report_definition_is_seeded
	state = var.report_definition_state
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `category` - (Optional) An optional filter to return only resources that match the specified category.
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `data_source` - (Optional) Specifies the name of a resource that provides data for the report. For example  alerts, events.
* `display_name` - (Optional) The name of the report definition to query.
* `is_seeded` - (Optional) A boolean flag indicating to list seeded report definitions. Set this parameter to get list of seeded report definitions.
* `state` - (Optional) An optional filter to return only resources that match the specified lifecycle state.


## Attributes Reference

The following attributes are exported:

* `report_definition_collection` - The list of report_definition_collection.

### ReportDefinition Reference

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
* `parent_id` - The OCID of the parent report definition. In the case of seeded report definition, this is same as definition OCID.
* `record_time_span` - The time span for the records in the report to be scheduled. <period-value><period> Allowed period strings - "H","D","M","Y" Each of the above fields potentially introduce constraints. A workRequest is created only when period-value satisfies all the constraints. Constraints introduced: 1. period = H (The allowed range for period-value is [1, 23]) 2. period = D (The allowed range for period-value is [1, 30]) 3. period = M (The allowed range for period-value is [1, 11]) 4. period = Y (The minimum period-value is 1) 
* `schedule` - The schedule to generate the report periodically in the specified format: <version-string>;<version-specific-schedule>

	Allowed version strings - "v1" v1's version specific schedule -<ss> <mm> <hh> <day-of-week> <day-of-month> Each of the above fields potentially introduce constraints. A workrequest is created only when clock time satisfies all the constraints. Constraints introduced: 1. seconds = <ss> (So, the allowed range for <ss> is [0, 59]) 2. minutes = <mm> (So, the allowed range for <mm> is [0, 59]) 3. hours = <hh> (So, the allowed range for <hh> is [0, 23]) 4. <day-of-week> can be either '*' (without quotes or a number between 1(Monday) and 7(Sunday)) No constraint introduced when it is '*'. When not, day of week must equal the given value 5. <day-of-month> can be either '*' (without quotes or a number between 1 and 28) No constraint introduced when it is '*'. When not, day of month must equal the given value 
* `scheduled_report_compartment_id` - The OCID of the compartment in which the scheduled resource will be created. 
* `scheduled_report_mime_type` - Specifies the format of the report ( either .xls or .pdf )
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

