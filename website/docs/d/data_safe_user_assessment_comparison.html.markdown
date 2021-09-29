---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_user_assessment_comparison"
sidebar_current: "docs-oci-datasource-data_safe-user_assessment_comparison"
description: |-
  Provides details about a specific User Assessment Comparison in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_user_assessment_comparison
This data source provides details about a specific User Assessment Comparison resource in Oracle Cloud Infrastructure Data Safe service.

Gets the details of the comparison report for the user assessments provided.

## Example Usage

```hcl
data "oci_data_safe_user_assessment_comparison" "test_user_assessment_comparison" {
	#Required
	comparison_user_assessment_id = oci_data_safe_user_assessment.test_user_assessment.id
	user_assessment_id = oci_data_safe_user_assessment.test_user_assessment.id
}
```

## Argument Reference

The following arguments are supported:

* `comparison_user_assessment_id` - (Required) The OCID of the baseline user assessment.
* `user_assessment_id` - (Required) The OCID of the user assessment.


## Attributes Reference

The following attributes are exported:

* `state` - The current state of the user assessment comparison.
* `summary` - List containing maps as values. Example: `{"Operations": [ {"CostCenter": "42"} ] }` 
* `time_created` - The date and time the user assessment comparison was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

