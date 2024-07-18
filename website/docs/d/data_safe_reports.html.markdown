---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_reports"
sidebar_current: "docs-oci-datasource-data_safe-reports"
description: |-
  Provides the list of Reports in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_reports
This data source provides the list of Reports in Oracle Cloud Infrastructure Data Safe service.

Gets a list of all the reports in the compartment. It contains information such as report generation time.

## Example Usage

```hcl
data "oci_data_safe_reports" "test_reports" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.report_access_level
	compartment_id_in_subtree = var.report_compartment_id_in_subtree
	display_name = var.report_display_name
	report_definition_id = oci_data_safe_report_definition.test_report_definition.id
	state = var.report_state
	time_generated_greater_than_or_equal_to = var.report_time_generated_greater_than_or_equal_to
	time_generated_less_than = var.report_time_generated_less_than
	type = var.report_type
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `display_name` - (Optional) The name of the report definition to query.
* `report_definition_id` - (Optional) The ID of the report definition to filter the list of reports
* `state` - (Optional) An optional filter to return only resources that match the specified lifecycle state.
* `time_generated_greater_than_or_equal_to` - (Optional) A filter to return only the resources that were generated after the specified date and time, as defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Using TimeGeneratedGreaterThanOrEqualToQueryParam parameter retrieves all resources generated after that date.

	**Example:** 2016-12-19T16:39:57.600Z 
* `time_generated_less_than` - (Optional) Search for resources that were generated before a specific date. Specifying this parameter corresponding `timeGeneratedLessThan` parameter will retrieve all resources generated before the specified generated date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339.

	**Example:** 2016-12-19T16:39:57.600Z 
* `type` - (Optional) An optional filter to return only resources that match the specified type.


## Attributes Reference

The following attributes are exported:

* `report_collection` - The list of report_collection.

### Report Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the report.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - Specifies a description of the report.
* `display_name` - Name of the report.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the report.
* `mime_type` - Specifies the format of report to be .xls or .pdf or .json
* `report_definition_id` - The OCID of the report definition.
* `state` - The current state of the audit report.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_generated` - Specifies the date and time the report was generated.
* `type` - The type of the audit report.

