---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_masking_policy_health_report"
sidebar_current: "docs-oci-datasource-data_safe-masking_policy_health_report"
description: |-
  Provides details about a specific Masking Policy Health Report in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_masking_policy_health_report
This data source provides details about a specific Masking Policy Health Report resource in Oracle Cloud Infrastructure Data Safe service.

Gets the details of the specified masking policy health report.

## Example Usage

```hcl
data "oci_data_safe_masking_policy_health_report" "test_masking_policy_health_report" {
	#Required
	masking_policy_health_report_id = oci_data_safe_masking_policy_health_report.test_masking_policy_health_report.id
}
```

## Argument Reference

The following arguments are supported:

* `masking_policy_health_report_id` - (Required) The OCID of the masking health report.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the health report.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the masking health report.
* `display_name` - The display name of the health report.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the health report.
* `masking_policy_id` - The OCID of the masking policy.
* `state` - The current state of the health report.
* `target_id` - The OCID of the target database for which this report was created.
* `time_created` - The date and time the report was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `time_updated` - The date and time the report was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339)  

