---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_security_assessment_check"
sidebar_current: "docs-oci-resource-data_safe-security_assessment_check"
description: |-
  Provides the Security Assessment Check resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_security_assessment_check
This resource provides the Security Assessment Check resource in Oracle Cloud Infrastructure Data Safe service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/data-safe/latest/SecurityAssessmentCheck

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/datasafe


  Patches one or more checks in the specified template type security assessment. Use it to add or delete checks.
To add check, use CreateCheckDetails as the patch value.


## Example Usage

```hcl
resource "oci_data_safe_security_assessment_check" "test_security_assessment_check" {
	#Required
	security_assessment_id = oci_data_safe_security_assessment.test_security_assessment.id

	#Optional
	patch_operations {
		#Required
		operation = var.security_assessment_check_patch_operations_operation
		selection = var.security_assessment_check_patch_operations_selection

		#Optional
		value = var.security_assessment_check_patch_operations_value
	}
}
```

## Argument Reference

The following arguments are supported:

* `patch_operations` - (Optional) (Updatable) 
	* `operation` - (Required) (Updatable) The operation can be one of these values: `INSERT`, `MERGE`, `REMOVE`
	* `selection` - (Required) (Updatable) 
	* `value` - (Required when operation=INSERT | MERGE) (Updatable) 
* `security_assessment_id` - (Required) 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `category` - The category to which the check belongs to.
* `key` - A unique identifier for the check.
* `oneline` - Provides a recommended approach to take to remediate the check reported.
* `references` - Provides information on whether the check is related to a CIS Oracle Database Benchmark recommendation, STIG rule, GDPR Article/Recital or related to the Oracle Best Practice.
	* `cis` - Relevant section from CIS.
	* `gdpr` - Relevant section from GDPR.
	* `obp` - Relevant section from OBP.
	* `stig` - Relevant section from STIG.
* `remarks` - The explanation of the issue in this check. It explains the reason for the rule and, if a risk is reported, it may also explain the recommended actions for remediation.
* `suggested_severity` - The severity of the check as suggested by Data Safe security assessment. This will be the default severity in the template baseline security assessment.
* `title` - The short title for the check.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Security Assessment Check
	* `update` - (Defaults to 20 minutes), when updating the Security Assessment Check
	* `delete` - (Defaults to 20 minutes), when destroying the Security Assessment Check


## Import

SecurityAssessmentChecks can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_security_assessment_check.test_security_assessment_check "securityAssessments/{securityAssessmentId}/checks" 
```

