---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_security_assessment_finding"
sidebar_current: "docs-oci-resource-data_safe-security_assessment_finding"
description: |-
  Provides the Security Assessment Finding resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_security_assessment_finding
This resource provides the Security Assessment Finding resource in Oracle Cloud Infrastructure Data Safe service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/data-safe/latest/SecurityAssessmentFinding

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/datasafe


  Patches one or more findings in the specified template baseline type security assessment. Use it to modify max allowed risk level in template baseline.


## Example Usage

```hcl
resource "oci_data_safe_security_assessment_finding" "test_security_assessment_finding" {
	#Required
	security_assessment_id = oci_data_safe_security_assessment.test_security_assessment.id

	#Optional
	patch_operations {
		#Required
		operation = var.security_assessment_finding_patch_operations_operation
		selection = var.security_assessment_finding_patch_operations_selection

		#Optional
		value = var.security_assessment_finding_patch_operations_value
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

* `assessment_id` - The OCID of the assessment that generated this finding.
* `category` - The category to which the finding belongs to.
* `details` - The details of the finding. Provides detailed information to explain the finding summary, typically results from the assessed database, followed by any recommendations for changes.
* `has_target_db_risk_level_changed` - Determines if this risk level has changed on the target database since the last time 'severity' was modified by user.
* `is_risk_modified` - Determines if this risk level was modified by user.
* `is_top_finding` - Indicates whether a given finding is marked as topFinding or not.
* `justification` - User provided reason for accepting or modifying this finding if they choose to do so.
* `key` - The unique finding key. This is a system-generated identifier. To get the finding key for a finding, use ListFindings.
* `lifecycle_details` - Details about the current state of the finding.
* `oneline` - Provides a recommended approach to take to remediate the finding reported.
* `oracle_defined_severity` - The severity of the finding as determined by security assessment. This cannot be modified by user.
* `references` - Provides information on whether the finding is related to a CIS Oracle Database Benchmark recommendation, a STIG rule, or a GDPR Article/Recital.
	* `cis` - Relevant section from CIS.
	* `gdpr` - Relevant section from GDPR.
	* `obp` - Relevant section from OBP.
	* `stig` - Relevant section from STIG.
* `remarks` - The explanation of the issue in this finding. It explains the reason for the rule and, if a risk is reported, it may also explain the recommended actions for remediation.
* `severity` - The severity of the finding as determined by security assessment and is same as oracleDefinedSeverity, unless modified by user.
* `state` - The current state of the finding.
* `summary` - The brief summary of the finding. When the finding is informational, the summary typically reports only the number of data elements that were examined.
* `target_id` - The OCID of the target database.
* `time_updated` - The date and time the risk level of finding was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `time_valid_until` - The time until which the change in severity(deferred / modified) of this finding is valid.
* `title` - The short title for the finding.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Security Assessment Finding
	* `update` - (Defaults to 20 minutes), when updating the Security Assessment Finding
	* `delete` - (Defaults to 20 minutes), when destroying the Security Assessment Finding


## Import

SecurityAssessmentFindings can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_security_assessment_finding.test_security_assessment_finding "securityAssessments/{securityAssessmentId}/findings" 
```

