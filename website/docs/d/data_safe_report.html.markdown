---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_report"
sidebar_current: "docs-oci-datasource-data_safe-report"
description: |-
  Provides details about a specific Report in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_report
This data source provides details about a specific Report resource in Oracle Cloud Infrastructure Data Safe service.

Gets a report by identifier

## Example Usage

```hcl
data "oci_data_safe_report" "test_report" {
	#Required
	report_id = oci_data_safe_report.test_report.id
}
```

## Argument Reference

The following arguments are supported:

* `report_id` - (Required) Unique report identifier


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the report.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - Specifies a description of the report.
* `display_name` - Name of the report.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the report.
* `mime_type` - Specifies the format of report to be .xls or .pdf
* `report_definition_id` - The OCID of the report definition.
* `state` - The current state of the audit report.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_generated` - Specifies the date and time the report was generated.
* `type` - The type of the audit report.

