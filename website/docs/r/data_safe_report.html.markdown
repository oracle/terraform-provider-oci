---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_report"
sidebar_current: "docs-oci-resource-data_safe-report"
description: |-
  Provides the Report resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_report
This resource provides the Report resource in Oracle Cloud Infrastructure Data Safe service.

Updates the specified report. Only tags can be updated.

## Example Usage

```hcl
resource "oci_data_safe_report" "test_report" {
	#Required
	report_id = oci_data_safe_report.test_report.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) (Updatable) The OCID of the compartment containing the report.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `report_id` - (Required) Unique report identifier


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the report.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - Specifies a description of the report.
* `display_name` - Name of the report.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the report.
* `mime_type` - Specifies the format of report to be excel or pdf
* `report_definition_id` - The OCID of the report definition.
* `state` - The current state of the audit report.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_generated` - Specifies the date and time the report was generated.
* `type` - The type of the audit report.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Report
	* `update` - (Defaults to 20 minutes), when updating the Report
	* `delete` - (Defaults to 20 minutes), when destroying the Report


## Import

Reports can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_report.test_report "id"
```

