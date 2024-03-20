---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_set_security_assessment_baseline_management"
sidebar_current: "docs-oci-resource-data_safe-set_security_assessment_baseline_management"
description: |-
  Provides the Set Security Assessment Baseline Management resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_set_security_assessment_baseline_management
This resource provides the Set Security Assessment Baseline management resource in Oracle Cloud Infrastructure Data Safe service.

Sets the saved security assessment as the baseline in the compartment where the the specified assessment resides. The security assessment needs to be of type 'SAVED'.

## Example Usage

```hcl
resource "oci_data_safe_set_security_assessment_baseline_management" "test_set_security_assessment_baseline_management" {
	#Required
	target_id = oci_data_safe_target_database.test_target_database.id
	compartment_id = var.comaprtment_id
}
```

## Argument Reference

The following arguments are supported:

* `target_id` - (Required) The target OCID for which SA needs to be set as baseline.
* `compartment_id` - (Required) The compartment OCID of the target.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Set Security Assessment Baseline Management
	* `update` - (Defaults to 20 minutes), when updating the Set Security Assessment Baseline Management
	* `delete` - (Defaults to 20 minutes), when destroying the Set Security Assessment Baseline Management


## Import

SetSecurityAssessmentBaselineManagement can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_set_security_assessment_baseline_management.test_set_security_assessment_baseline_management "id"
```

