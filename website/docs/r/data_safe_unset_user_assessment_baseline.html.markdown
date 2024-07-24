---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_unset_user_assessment_baseline"
sidebar_current: "docs-oci-resource-data_safe-unset_user_assessment_baseline"
description: |-
  Provides the Unset User Assessment Baseline resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_unset_user_assessment_baseline
This resource provides the Unset User Assessment Baseline resource in Oracle Cloud Infrastructure Data Safe service.

Removes the baseline setting for the saved user assessment associated with the targetId passed via body.
If no body or empty body is passed then the baseline settings of all the saved user assessments pertaining to the baseline assessment OCID provided in the path will be removed.
Sets the if-match parameter to the value of the etag from a previous GET or POST response for that resource.


## Example Usage

```hcl
resource "oci_data_safe_unset_user_assessment_baseline" "test_unset_user_assessment_baseline" {
	#Required
	user_assessment_id = oci_data_safe_user_assessment.test_user_assessment.id

	#Optional
	target_ids = var.unset_user_assessment_baseline_target_ids
}
```

## Argument Reference

The following arguments are supported:

* `target_ids` - (Optional) The list of database target OCIDs for which the user intends to unset the baseline.
* `user_assessment_id` - (Required) The OCID of the user assessment.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Unset User Assessment Baseline
	* `update` - (Defaults to 20 minutes), when updating the Unset User Assessment Baseline
	* `delete` - (Defaults to 20 minutes), when destroying the Unset User Assessment Baseline


## Import

UnsetUserAssessmentBaseline can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_unset_user_assessment_baseline.test_unset_user_assessment_baseline "id"
```

