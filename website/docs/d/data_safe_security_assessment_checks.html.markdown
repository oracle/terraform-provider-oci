---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_security_assessment_checks"
sidebar_current: "docs-oci-datasource-data_safe-security_assessment_checks"
description: |-
  Provides the list of Security Assessment Checks in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_security_assessment_checks
This data source provides the list of Security Assessment Checks in Oracle Cloud Infrastructure Data Safe service.

Lists all the security checks in the specified compartment for security assessment of type TEMPLATE.


## Example Usage

```hcl
data "oci_data_safe_security_assessment_checks" "test_security_assessment_checks" {
	#Required
	security_assessment_id = oci_data_safe_security_assessment.test_security_assessment.id

	#Optional
	access_level = var.security_assessment_check_access_level
	compartment_id_in_subtree = var.security_assessment_check_compartment_id_in_subtree
	contains_references = var.security_assessment_check_contains_references
	contains_severity = var.security_assessment_check_contains_severity
	key = var.security_assessment_check_key
	suggested_severity = var.security_assessment_check_suggested_severity
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `contains_references` - (Optional) An optional filter to return only findings that match the specified references. Use containsReferences param if need to filter by multiple references.
* `contains_severity` - (Optional) A filter to return only findings that match the specified risk level(s). Use containsSeverity parameter if need to filter by multiple risk levels.
* `key` - (Optional) Each check in security assessment has an associated key (think of key as a check's name). For a given check, the key will be the same across targets. The user can use these keys to filter the checks. 
* `security_assessment_id` - (Required) The OCID of the security assessment.
* `suggested_severity` - (Optional) A filter to return only checks of a particular risk level.


## Attributes Reference

The following attributes are exported:

* `checks` - The list of checks.

### SecurityAssessmentCheck Reference

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

