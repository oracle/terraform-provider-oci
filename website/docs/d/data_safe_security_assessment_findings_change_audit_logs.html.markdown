---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_security_assessment_findings_change_audit_logs"
sidebar_current: "docs-oci-datasource-data_safe-security_assessment_findings_change_audit_logs"
description: |-
  Provides the list of Security Assessment Findings Change Audit Logs in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_security_assessment_findings_change_audit_logs
This data source provides the list of Security Assessment Findings Change Audit Logs in Oracle Cloud Infrastructure Data Safe service.

List all changes made by user to risk level of findings of the specified assessment.


## Example Usage

```hcl
data "oci_data_safe_security_assessment_findings_change_audit_logs" "test_security_assessment_findings_change_audit_logs" {
	#Required
	security_assessment_id = oci_data_safe_security_assessment.test_security_assessment.id

	#Optional
	finding_key = var.security_assessment_findings_change_audit_log_finding_key
	finding_title = var.security_assessment_findings_change_audit_log_finding_title
	is_risk_deferred = var.security_assessment_findings_change_audit_log_is_risk_deferred
	modified_by = var.security_assessment_findings_change_audit_log_modified_by
	severity = var.security_assessment_findings_change_audit_log_severity
	time_updated_greater_than_or_equal_to = var.security_assessment_findings_change_audit_log_time_updated_greater_than_or_equal_to
	time_updated_less_than = var.security_assessment_findings_change_audit_log_time_updated_less_than
	time_valid_until_greater_than_or_equal_to = var.security_assessment_findings_change_audit_log_time_valid_until_greater_than_or_equal_to
	time_valid_until_less_than = var.security_assessment_findings_change_audit_log_time_valid_until_less_than
}
```

## Argument Reference

The following arguments are supported:

* `finding_key` - (Optional) The unique key that identifies the finding. It is a string and unique within a security assessment.
* `finding_title` - (Optional) The unique title that identifies the finding. It is a string and unique within a security assessment.
* `is_risk_deferred` - (Optional) A filter to check findings whose risks were deferred by the user.
* `modified_by` - (Optional) A filter to check which user modified the risk level of the finding.
* `security_assessment_id` - (Required) The OCID of the security assessment.
* `severity` - (Optional) A filter to return only findings of a particular risk level.
* `time_updated_greater_than_or_equal_to` - (Optional) Search for resources that were updated after a specific date. Specifying this parameter corresponding `timeUpdatedGreaterThanOrEqualTo` parameter will retrieve all resources updated after the specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339. 
* `time_updated_less_than` - (Optional) Search for resources that were updated before a specific date. Specifying this parameter corresponding `timeUpdatedLessThan` parameter will retrieve all resources updated before the specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339. 
* `time_valid_until_greater_than_or_equal_to` - (Optional) Specifying `TimeValidUntilGreaterThanOrEqualToQueryParam` parameter  will retrieve all items for which the risk level modification by user will  no longer be valid greater than the date and time specified, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

	**Example:** 2016-12-19T00:00:00.000Z 
* `time_valid_until_less_than` - (Optional) Specifying `TimeValidUntilLessThanQueryParam` parameter will retrieve all items for which the risk level modification by user will  be valid until less than the date and time specified, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

	**Example:** 2016-12-19T00:00:00.000Z 


## Attributes Reference

The following attributes are exported:

* `findings_change_audit_log_collection` - The list of findings_change_audit_log_collection.

### SecurityAssessmentFindingsChangeAuditLog Reference

The following attributes are exported:

* `items` - An array of finding risk change audit log summary objects.
	* `assessment_id` - The OCID of the latest security assessment.
	* `finding_key` - The unique key that identifies the finding.
	* `finding_title` - The short title for the finding whose risk is being modified.
	* `is_risk_deferred` - Determines if the user has deferred the risk level of this finding when he is ok with it  and does not plan to do anything about it. 
	* `justification` - The justification given by the user for accepting or modifying the risk level.
	* `key` - The unique key that identifies the finding risk change.
	* `modified_by` - The user who initiated change of risk level of the finding
	* `oracle_defined_severity` - The severity of the finding as determined by security assessment by Oracle.
	* `previous_severity` - If the risk level is changed more than once, the previous modified value.
	* `severity` - The original severity / risk level of the finding as determined by security assessment.
	* `target_id` - The OCID of the target database.
	* `time_updated` - The date and time the risk level of finding was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
	* `time_valid_until` - The date and time, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339), the risk level change as updated by user is valid until. After this date passes, the risk level will be that of what is determined by the latest security assessment. 

