---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_compare_security_assessment"
sidebar_current: "docs-oci-resource-data_safe-compare_security_assessment"
description: |-
  Provides the Compare Security Assessment resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_compare_security_assessment
This resource provides the Compare Security Assessment resource in Oracle Cloud Infrastructure Data Safe service.

Compares two security assessments. For this comparison, a security assessment can be a saved assessment, a latest assessment, or a baseline assessment.
For example, you can compare saved assessment or a latest assessment against a baseline.


## Example Usage

```hcl
resource "oci_data_safe_compare_security_assessment" "test_compare_security_assessment" {
	#Required
	comparison_security_assessment_id = oci_data_safe_security_assessment.test_security_assessment.id
	security_assessment_id = oci_data_safe_security_assessment.test_security_assessment.id
}
```

## Argument Reference

The following arguments are supported:

* `comparison_security_assessment_id` - (Required) The OCID of the security assessment. In this case a security assessment can be another security assessment, a latest assessment or a baseline. 
* `security_assessment_id` - (Required) The OCID of the security assessment.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Compare Security Assessment
	* `update` - (Defaults to 20 minutes), when updating the Compare Security Assessment
	* `delete` - (Defaults to 20 minutes), when destroying the Compare Security Assessment


## Import

CompareSecurityAssessment can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_compare_security_assessment.test_compare_security_assessment "id"
```

