---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_set_user_assessment_baseline"
sidebar_current: "docs-oci-resource-data_safe-set_user_assessment_baseline"
description: |-
  Provides the Set User Assessment Baseline resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_set_user_assessment_baseline
This resource provides the Set User Assessment Baseline resource in Oracle Cloud Infrastructure Data Safe service.

Sets the saved user assessment as the baseline in the compartment where the specified assessment resides. The user assessment needs to be of type 'SAVED'.

## Example Usage

```hcl
resource "oci_data_safe_set_user_assessment_baseline" "test_set_user_assessment_baseline" {
	#Required
	user_assessment_id = oci_data_safe_user_assessment.test_user_assessment.id

	#Optional
	assessment_ids = var.set_user_assessment_baseline_assessment_ids
}
```

## Argument Reference

The following arguments are supported:

* `assessment_ids` - (Optional) The list of OCIDs for the user assessments that need to be updated while setting the baseline.
* `user_assessment_id` - (Required) The OCID of the user assessment.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Set User Assessment Baseline
	* `update` - (Defaults to 20 minutes), when updating the Set User Assessment Baseline
	* `delete` - (Defaults to 20 minutes), when destroying the Set User Assessment Baseline


## Import

SetUserAssessmentBaseline can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_set_user_assessment_baseline.test_set_user_assessment_baseline "id"
```

