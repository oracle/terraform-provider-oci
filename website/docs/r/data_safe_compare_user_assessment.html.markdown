---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_compare_user_assessment"
sidebar_current: "docs-oci-resource-data_safe-compare_user_assessment"
description: |-
  Provides the Compare User Assessment resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_compare_user_assessment
This resource provides the Compare User Assessment resource in Oracle Cloud Infrastructure Data Safe service.

Compares two user assessments. For this comparison, a user assessment can be a saved, a latest assessment, or a baseline.
As an example, it can be used to compare a user assessment saved or a latest assessment with a baseline.


## Example Usage

```hcl
resource "oci_data_safe_compare_user_assessment" "test_compare_user_assessment" {
	#Required
	comparison_user_assessment_id = oci_data_safe_user_assessment.test_user_assessment.id
	user_assessment_id = oci_data_safe_user_assessment.test_user_assessment.id
}
```

## Argument Reference

The following arguments are supported:

* `comparison_user_assessment_id` - (Required) The OCID of the user assessment to be compared. You can compare with another user assessment, a latest assessment, or a baseline. 
* `user_assessment_id` - (Required) The OCID of the user assessment.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Compare User Assessment
	* `update` - (Defaults to 20 minutes), when updating the Compare User Assessment
	* `delete` - (Defaults to 20 minutes), when destroying the Compare User Assessment


## Import

CompareUserAssessment can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_compare_user_assessment.test_compare_user_assessment "id"
```

