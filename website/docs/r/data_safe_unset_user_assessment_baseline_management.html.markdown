---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_unset_user_assessment_baseline_management"
sidebar_current: "docs-oci-resource-data_safe-unset_user_assessment_baseline_management"
description: |-
  Provides the Unset User Assessment Baseline Management resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_set_user_assessment_baseline_management
This resource provides the Unset User Assessment Baseline management resource in Oracle Cloud Infrastructure Data Safe service.

Removes the baseline setting for the saved user assessment as the baseline in the compartment where the the specified assessment resides.


## Example Usage

```hcl
resource "oci_data_safe_unset_user_assessment_baseline_management" "test_unset_user_assessment_baseline_management" {
	#Required
	user_assessment_id = oci_data_safe_set_user_assessment_baseline_management.test_set_user_assessment_baseline_management.user_assessment_id
	compartment_id = var.comaprtment_id
}
```

## Argument Reference

The following arguments are supported:

* `user_assessment_id` - (Required) The OCID of the user assessment.
* `compartment_id` - (Required) The compartment OCID where the assessment resides.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Unset User Assessment Baseline Management
	* `update` - (Defaults to 20 minutes), when updating the Unset User Assessment Baseline Management
	* `delete` - (Defaults to 20 minutes), when destroying the Unset User Assessment Baseline Management


## Import

UnsetUserAssessmentBaselineManagement can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_unset_user_assessment_baseline_management.test_unset_user_assessment_baseline_management "id"
```

