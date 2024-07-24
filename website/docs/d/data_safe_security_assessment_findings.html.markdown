---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_security_assessment_findings"
sidebar_current: "docs-oci-datasource-data_safe-security_assessment_findings"
description: |-
  Provides the list of Security Assessment Findings in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_security_assessment_findings
This data source provides the list of Security Assessment Findings in Oracle Cloud Infrastructure Data Safe service.

List all the findings from all the targets in the specified compartment.


## Example Usage

```hcl
data "oci_data_safe_security_assessment_findings" "test_security_assessment_findings" {
	#Required
	security_assessment_id = oci_data_safe_security_assessment.test_security_assessment.id

	#Optional
	access_level = var.security_assessment_finding_access_level
	compartment_id_in_subtree = var.security_assessment_finding_compartment_id_in_subtree
	finding_key = var.security_assessment_finding_finding_key
	is_top_finding = var.security_assessment_finding_is_top_finding
	references {
	}
	severity = var.security_assessment_finding_severity
	state = var.security_assessment_finding_state
	target_id = oci_cloud_guard_target.test_target.id
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `finding_key` - (Optional) Each finding in security assessment has an associated key (think of key as a finding's name). For a given finding, the key will be the same across targets. The user can use these keys to filter the findings. 
* `is_top_finding` - (Optional) A filter to return only the findings that are marked as top findings.
* `references` - (Optional) An optional filter to return only findings containing the specified reference.
* `security_assessment_id` - (Required) The OCID of the security assessment.
* `severity` - (Optional) A filter to return only findings of a particular risk level.
* `state` - (Optional) A filter to return only the findings that match the specified lifecycle states.
* `target_id` - (Optional) A filter to return only items related to a specific target OCID.


## Attributes Reference

The following attributes are exported:

* `findings` - The list of findings.

### SecurityAssessmentFinding Reference

The following attributes are exported:

* `assessment_id` - The OCID of the assessment that generated this finding.
* `details` - The details of the finding. Provides detailed information to explain the finding summary, typically results from the assessed database, followed by any recommendations for changes.
* `is_top_finding` - Indicates whether a given finding is marked as topFinding or not.
* `has_target_db_risk_level_changed` - Determines if this risk level has changed on the target database since the last time 'severity' was modified by user.
* `is_risk_modified` - Determines if this risk level was modified by user.
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

