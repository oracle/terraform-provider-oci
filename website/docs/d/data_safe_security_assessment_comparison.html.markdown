---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_security_assessment_comparison"
sidebar_current: "docs-oci-datasource-data_safe-security_assessment_comparison"
description: |-
	Provides details about a specific Security Assessment Comparison in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_security_assessment_comparison
This data source provides details about a specific Security Assessment Comparison resource in Oracle Cloud Infrastructure Data Safe service.

Gets the details of the comparison report for the security assessments submitted for comparison.

## Example Usage

```hcl
data "oci_data_safe_security_assessment_comparison" "test_security_assessment_comparison" {
	#Required
	comparison_security_assessment_id = oci_data_safe_security_assessment.test_security_assessment.id
	security_assessment_id = oci_data_safe_security_assessment.test_security_assessment.id
}
```

## Argument Reference

The following arguments are supported:

* `comparison_security_assessment_id` - (Required) The OCID of the security assessment baseline.
* `security_assessment_id` - (Required) The OCID of the security assessment.


## Attributes Reference

The following attributes are exported:

* `baseline_id` - The OCID of the security assessment that is set as a baseline.
* `id` - The OCID of the security assessment that is being compared with a baseline security assessment.
* `state` - The current state of the security assessment comparison.
* `targets` - A target-based comparison between two security assessments.
	* `auditing` - A comparison between findings belonging to Auditing category.
		* `added_items` - This array identifies the items that are present in the current assessment, but are missing from the baseline.
		* `baseline` - The particular finding reported by the security assessment.
			* `assessment_id` - The OCID of the assessment that generated this finding.
			* `details` - The details of the finding. Provides detailed information to explain the finding summary, typically results from the assessed database, followed by any recommendations for changes.
			* `has_target_db_risk_level_changed` - Determines if this risk level has changed on the target database since the last time 'severity' was modified by user.
			* `is_risk_modified` - Determines if this risk level was modified by user.
			* `justification` - User provided reason for accepting or modifying this finding if they choose to do so.
			* `key` - A unique identifier for the finding. This is common for the finding across targets.
			* `lifecycle_details` - Details about the current state of the finding.
			* `oracle_defined_severity` - The severity of the finding as determined by security assessment. This cannot be modified by user.
			* `references` - Provides information on whether the finding is related to a CIS Oracle Database Benchmark recommendation, STIG rule, or related to a GDPR Article/Recital.
				* `cis` - Relevant section from CIS.
				* `gdpr` - Relevant section from GDPR.
				* `obp` - Relevant section from OBP.
				* `stig` - Relevant section from STIG.
			* `remarks` - The explanation of the issue in this finding. It explains the reason for the rule and, if a risk is reported, it may also explain the recommended actions for remediation.
			* `severity` - The severity of the finding.
			* `state` - The current state of the finding.
			* `summary` - The brief summary of the finding. When the finding is informational, the summary typically reports only the number of data elements that were examined.
			* `target_id` - The OCID of the target database.
			* `time_updated` - The date and time the risk level of finding was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
			* `time_valid_until` - The time until which the change in severity(deferred/modified) of this finding is valid.
			* `title` - The short title for the finding.
		* `current` - The particular finding reported by the security assessment.
			* `assessment_id` - The OCID of the assessment that generated this finding.
			* `details` - The details of the finding. Provides detailed information to explain the finding summary, typically results from the assessed database, followed by any recommendations for changes.
			* `has_target_db_risk_level_changed` - Determines if this risk level has changed on the target database since the last time 'severity' was modified by user.
			* `is_risk_modified` - Determines if this risk level was modified by user.
			* `justification` - User provided reason for accepting or modifying this finding if they choose to do so.
			* `key` - A unique identifier for the finding. This is common for the finding across targets.
			* `lifecycle_details` - Details about the current state of the finding.
			* `oracle_defined_severity` - The severity of the finding as determined by security assessment. This cannot be modified by user.
			* `references` - Provides information on whether the finding is related to a CIS Oracle Database Benchmark recommendation, STIG rule, or related to a GDPR Article/Recital.
				* `cis` - Relevant section from CIS.
				* `gdpr` - Relevant section from GDPR.
				* `obp` - Relevant section from OBP.
				* `stig` - Relevant section from STIG.
			* `remarks` - The explanation of the issue in this finding. It explains the reason for the rule and, if a risk is reported, it may also explain the recommended actions for remediation.
			* `severity` - The severity of the finding.
			* `state` - The current state of the finding.
			* `summary` - The brief summary of the finding. When the finding is informational, the summary typically reports only the number of data elements that were examined.
			* `target_id` - The OCID of the target database.
			* `time_updated` - The date and time the risk level of finding was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
			* `time_valid_until` - The time until which the change in severity(deferred/modified) of this finding is valid.
			* `title` - The short title for the finding.
		* `modified_items` - This array contains the items that are present in both the current assessment and the baseline, but are different in the two assessments.
		* `removed_items` - This array identifies the items that are present in the baseline, but are missing from the current assessment.
		* `severity` - The severity of this diff.
	* `authorization_control` - A comparison between findings belonging to Authorization Control category.
		* `added_items` - This array identifies the items that are present in the current assessment, but are missing from the baseline.
		* `baseline` - The particular finding reported by the security assessment.
			* `assessment_id` - The OCID of the assessment that generated this finding.
			* `details` - The details of the finding. Provides detailed information to explain the finding summary, typically results from the assessed database, followed by any recommendations for changes.
			* `has_target_db_risk_level_changed` - Determines if this risk level has changed on the target database since the last time 'severity' was modified by user.
			* `is_risk_modified` - Determines if this risk level was modified by user.
			* `justification` - User provided reason for accepting or modifying this finding if they choose to do so.
			* `key` - A unique identifier for the finding. This is common for the finding across targets.
			* `lifecycle_details` - Details about the current state of the finding.
			* `oracle_defined_severity` - The severity of the finding as determined by security assessment. This cannot be modified by user.
			* `references` - Provides information on whether the finding is related to a CIS Oracle Database Benchmark recommendation, STIG rule, or related to a GDPR Article/Recital.
				* `cis` - Relevant section from CIS.
				* `gdpr` - Relevant section from GDPR.
				* `obp` - Relevant section from OBP.
				* `stig` - Relevant section from STIG.
			* `remarks` - The explanation of the issue in this finding. It explains the reason for the rule and, if a risk is reported, it may also explain the recommended actions for remediation.
			* `severity` - The severity of the finding.
			* `state` - The current state of the finding.
			* `summary` - The brief summary of the finding. When the finding is informational, the summary typically reports only the number of data elements that were examined.
			* `target_id` - The OCID of the target database.
			* `time_updated` - The date and time the risk level of finding was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
			* `time_valid_until` - The time until which the change in severity(deferred/modified) of this finding is valid.
			* `title` - The short title for the finding.
		* `current` - The particular finding reported by the security assessment.
			* `assessment_id` - The OCID of the assessment that generated this finding.
			* `details` - The details of the finding. Provides detailed information to explain the finding summary, typically results from the assessed database, followed by any recommendations for changes.
			* `has_target_db_risk_level_changed` - Determines if this risk level has changed on the target database since the last time 'severity' was modified by user.
			* `is_risk_modified` - Determines if this risk level was modified by user.
			* `justification` - User provided reason for accepting or modifying this finding if they choose to do so.
			* `key` - A unique identifier for the finding. This is common for the finding across targets.
			* `lifecycle_details` - Details about the current state of the finding.
			* `oracle_defined_severity` - The severity of the finding as determined by security assessment. This cannot be modified by user.
			* `references` - Provides information on whether the finding is related to a CIS Oracle Database Benchmark recommendation, STIG rule, or related to a GDPR Article/Recital.
				* `cis` - Relevant section from CIS.
				* `gdpr` - Relevant section from GDPR.
				* `obp` - Relevant section from OBP.
				* `stig` - Relevant section from STIG.
			* `remarks` - The explanation of the issue in this finding. It explains the reason for the rule and, if a risk is reported, it may also explain the recommended actions for remediation.
			* `severity` - The severity of the finding.
			* `state` - The current state of the finding.
			* `summary` - The brief summary of the finding. When the finding is informational, the summary typically reports only the number of data elements that were examined.
			* `target_id` - The OCID of the target database.
			* `time_updated` - The date and time the risk level of finding was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
			* `time_valid_until` - The time until which the change in severity(deferred/modified) of this finding is valid.
			* `title` - The short title for the finding.
		* `modified_items` - This array contains the items that are present in both the current assessment and the baseline, but are different in the two assessments.
		* `removed_items` - This array identifies the items that are present in the baseline, but are missing from the current assessment.
		* `severity` - The severity of this diff.
	* `baseline_target_id` - The OCID of the target that is used as a baseline in this comparison.
	* `current_target_id` - The OCID of the target to be compared against the baseline target.
	* `data_encryption` - Comparison between findings belonging to Data Encryption category.
		* `added_items` - This array identifies the items that are present in the current assessment, but are missing from the baseline.
		* `baseline` - The particular finding reported by the security assessment.
			* `assessment_id` - The OCID of the assessment that generated this finding.
			* `details` - The details of the finding. Provides detailed information to explain the finding summary, typically results from the assessed database, followed by any recommendations for changes.
			* `has_target_db_risk_level_changed` - Determines if this risk level has changed on the target database since the last time 'severity' was modified by user.
			* `is_risk_modified` - Determines if this risk level was modified by user.
			* `justification` - User provided reason for accepting or modifying this finding if they choose to do so.
			* `key` - A unique identifier for the finding. This is common for the finding across targets.
			* `lifecycle_details` - Details about the current state of the finding.
			* `oracle_defined_severity` - The severity of the finding as determined by security assessment. This cannot be modified by user.
			* `references` - Provides information on whether the finding is related to a CIS Oracle Database Benchmark recommendation, STIG rule, or related to a GDPR Article/Recital.
				* `cis` - Relevant section from CIS.
				* `gdpr` - Relevant section from GDPR.
				* `obp` - Relevant section from OBP.
				* `stig` - Relevant section from STIG.
			* `remarks` - The explanation of the issue in this finding. It explains the reason for the rule and, if a risk is reported, it may also explain the recommended actions for remediation.
			* `severity` - The severity of the finding.
			* `state` - The current state of the finding.
			* `summary` - The brief summary of the finding. When the finding is informational, the summary typically reports only the number of data elements that were examined.
			* `target_id` - The OCID of the target database.
			* `time_updated` - The date and time the risk level of finding was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
			* `time_valid_until` - The time until which the change in severity(deferred/modified) of this finding is valid.
			* `title` - The short title for the finding.
		* `current` - The particular finding reported by the security assessment.
			* `assessment_id` - The OCID of the assessment that generated this finding.
			* `details` - The details of the finding. Provides detailed information to explain the finding summary, typically results from the assessed database, followed by any recommendations for changes.
			* `has_target_db_risk_level_changed` - Determines if this risk level has changed on the target database since the last time 'severity' was modified by user.
			* `is_risk_modified` - Determines if this risk level was modified by user.
			* `justification` - User provided reason for accepting or modifying this finding if they choose to do so.
			* `key` - A unique identifier for the finding. This is common for the finding across targets.
			* `lifecycle_details` - Details about the current state of the finding.
			* `oracle_defined_severity` - The severity of the finding as determined by security assessment. This cannot be modified by user.
			* `references` - Provides information on whether the finding is related to a CIS Oracle Database Benchmark recommendation, STIG rule, or related to a GDPR Article/Recital.
				* `cis` - Relevant section from CIS.
				* `gdpr` - Relevant section from GDPR.
				* `obp` - Relevant section from OBP.
				* `stig` - Relevant section from STIG.
			* `remarks` - The explanation of the issue in this finding. It explains the reason for the rule and, if a risk is reported, it may also explain the recommended actions for remediation.
			* `severity` - The severity of the finding.
			* `state` - The current state of the finding.
			* `summary` - The brief summary of the finding. When the finding is informational, the summary typically reports only the number of data elements that were examined.
			* `target_id` - The OCID of the target database.
			* `time_updated` - The date and time the risk level of finding was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
			* `time_valid_until` - The time until which the change in severity(deferred/modified) of this finding is valid.
			* `title` - The short title for the finding.
		* `modified_items` - This array contains the items that are present in both the current assessment and the baseline, but are different in the two assessments.
		* `removed_items` - This array identifies the items that are present in the baseline, but are missing from the current assessment.
		* `severity` - The severity of this diff.
	* `db_configuration` - Comparison between findings belonging to Database Configuration category.
		* `added_items` - This array identifies the items that are present in the current assessment, but are missing from the baseline.
		* `baseline` - The particular finding reported by the security assessment.
			* `assessment_id` - The OCID of the assessment that generated this finding.
			* `details` - The details of the finding. Provides detailed information to explain the finding summary, typically results from the assessed database, followed by any recommendations for changes.
			* `has_target_db_risk_level_changed` - Determines if this risk level has changed on the target database since the last time 'severity' was modified by user.
			* `is_risk_modified` - Determines if this risk level was modified by user.
			* `justification` - User provided reason for accepting or modifying this finding if they choose to do so.
			* `key` - A unique identifier for the finding. This is common for the finding across targets.
			* `lifecycle_details` - Details about the current state of the finding.
			* `oracle_defined_severity` - The severity of the finding as determined by security assessment. This cannot be modified by user.
			* `references` - Provides information on whether the finding is related to a CIS Oracle Database Benchmark recommendation, STIG rule, or related to a GDPR Article/Recital.
				* `cis` - Relevant section from CIS.
				* `gdpr` - Relevant section from GDPR.
				* `obp` - Relevant section from OBP.
				* `stig` - Relevant section from STIG.
			* `remarks` - The explanation of the issue in this finding. It explains the reason for the rule and, if a risk is reported, it may also explain the recommended actions for remediation.
			* `severity` - The severity of the finding.
			* `state` - The current state of the finding.
			* `summary` - The brief summary of the finding. When the finding is informational, the summary typically reports only the number of data elements that were examined.
			* `target_id` - The OCID of the target database.
			* `time_updated` - The date and time the risk level of finding was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
			* `time_valid_until` - The time until which the change in severity(deferred/modified) of this finding is valid.
			* `title` - The short title for the finding.
		* `current` - The particular finding reported by the security assessment.
			* `assessment_id` - The OCID of the assessment that generated this finding.
			* `details` - The details of the finding. Provides detailed information to explain the finding summary, typically results from the assessed database, followed by any recommendations for changes.
			* `has_target_db_risk_level_changed` - Determines if this risk level has changed on the target database since the last time 'severity' was modified by user.
			* `is_risk_modified` - Determines if this risk level was modified by user.
			* `justification` - User provided reason for accepting or modifying this finding if they choose to do so.
			* `key` - A unique identifier for the finding. This is common for the finding across targets.
			* `lifecycle_details` - Details about the current state of the finding.
			* `oracle_defined_severity` - The severity of the finding as determined by security assessment. This cannot be modified by user.
			* `references` - Provides information on whether the finding is related to a CIS Oracle Database Benchmark recommendation, STIG rule, or related to a GDPR Article/Recital.
				* `cis` - Relevant section from CIS.
				* `gdpr` - Relevant section from GDPR.
				* `obp` - Relevant section from OBP.
				* `stig` - Relevant section from STIG.
			* `remarks` - The explanation of the issue in this finding. It explains the reason for the rule and, if a risk is reported, it may also explain the recommended actions for remediation.
			* `severity` - The severity of the finding.
			* `state` - The current state of the finding.
			* `summary` - The brief summary of the finding. When the finding is informational, the summary typically reports only the number of data elements that were examined.
			* `target_id` - The OCID of the target database.
			* `time_updated` - The date and time the risk level of finding was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
			* `time_valid_until` - The time until which the change in severity(deferred/modified) of this finding is valid.
			* `title` - The short title for the finding.
		* `modified_items` - This array contains the items that are present in both the current assessment and the baseline, but are different in the two assessments.
		* `removed_items` - This array identifies the items that are present in the baseline, but are missing from the current assessment.
		* `severity` - The severity of this diff.
	* `fine_grained_access_control` - Comparison between findings belonging to Fine-Grained Access Control category.
		* `added_items` - This array identifies the items that are present in the current assessment, but are missing from the baseline.
		* `baseline` - The particular finding reported by the security assessment.
			* `assessment_id` - The OCID of the assessment that generated this finding.
			* `details` - The details of the finding. Provides detailed information to explain the finding summary, typically results from the assessed database, followed by any recommendations for changes.
			* `has_target_db_risk_level_changed` - Determines if this risk level has changed on the target database since the last time 'severity' was modified by user.
			* `is_risk_modified` - Determines if this risk level was modified by user.
			* `justification` - User provided reason for accepting or modifying this finding if they choose to do so.
			* `key` - A unique identifier for the finding. This is common for the finding across targets.
			* `lifecycle_details` - Details about the current state of the finding.
			* `oracle_defined_severity` - The severity of the finding as determined by security assessment. This cannot be modified by user.
			* `references` - Provides information on whether the finding is related to a CIS Oracle Database Benchmark recommendation, STIG rule, or related to a GDPR Article/Recital.
				* `cis` - Relevant section from CIS.
				* `gdpr` - Relevant section from GDPR.
				* `obp` - Relevant section from OBP.
				* `stig` - Relevant section from STIG.
			* `remarks` - The explanation of the issue in this finding. It explains the reason for the rule and, if a risk is reported, it may also explain the recommended actions for remediation.
			* `severity` - The severity of the finding.
			* `state` - The current state of the finding.
			* `summary` - The brief summary of the finding. When the finding is informational, the summary typically reports only the number of data elements that were examined.
			* `target_id` - The OCID of the target database.
			* `time_updated` - The date and time the risk level of finding was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
			* `time_valid_until` - The time until which the change in severity(deferred/modified) of this finding is valid.
			* `title` - The short title for the finding.
		* `current` - The particular finding reported by the security assessment.
			* `assessment_id` - The OCID of the assessment that generated this finding.
			* `details` - The details of the finding. Provides detailed information to explain the finding summary, typically results from the assessed database, followed by any recommendations for changes.
			* `has_target_db_risk_level_changed` - Determines if this risk level has changed on the target database since the last time 'severity' was modified by user.
			* `is_risk_modified` - Determines if this risk level was modified by user.
			* `justification` - User provided reason for accepting or modifying this finding if they choose to do so.
			* `key` - A unique identifier for the finding. This is common for the finding across targets.
			* `lifecycle_details` - Details about the current state of the finding.
			* `oracle_defined_severity` - The severity of the finding as determined by security assessment. This cannot be modified by user.
			* `references` - Provides information on whether the finding is related to a CIS Oracle Database Benchmark recommendation, STIG rule, or related to a GDPR Article/Recital.
				* `cis` - Relevant section from CIS.
				* `gdpr` - Relevant section from GDPR.
				* `obp` - Relevant section from OBP.
				* `stig` - Relevant section from STIG.
			* `remarks` - The explanation of the issue in this finding. It explains the reason for the rule and, if a risk is reported, it may also explain the recommended actions for remediation.
			* `severity` - The severity of the finding.
			* `state` - The current state of the finding.
			* `summary` - The brief summary of the finding. When the finding is informational, the summary typically reports only the number of data elements that were examined.
			* `target_id` - The OCID of the target database.
			* `time_updated` - The date and time the risk level of finding was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
			* `time_valid_until` - The time until which the change in severity(deferred/modified) of this finding is valid.
			* `title` - The short title for the finding.
		* `modified_items` - This array contains the items that are present in both the current assessment and the baseline, but are different in the two assessments.
		* `removed_items` - This array identifies the items that are present in the baseline, but are missing from the current assessment.
		* `severity` - The severity of this diff.
	* `privileges_and_roles` - Comparison between findings belonging to Privileges and Roles category.
		* `added_items` - This array identifies the items that are present in the current assessment, but are missing from the baseline.
		* `baseline` - The particular finding reported by the security assessment.
			* `assessment_id` - The OCID of the assessment that generated this finding.
			* `details` - The details of the finding. Provides detailed information to explain the finding summary, typically results from the assessed database, followed by any recommendations for changes.
			* `has_target_db_risk_level_changed` - Determines if this risk level has changed on the target database since the last time 'severity' was modified by user.
			* `is_risk_modified` - Determines if this risk level was modified by user.
			* `justification` - User provided reason for accepting or modifying this finding if they choose to do so.
			* `key` - A unique identifier for the finding. This is common for the finding across targets.
			* `lifecycle_details` - Details about the current state of the finding.
			* `oracle_defined_severity` - The severity of the finding as determined by security assessment. This cannot be modified by user.
			* `references` - Provides information on whether the finding is related to a CIS Oracle Database Benchmark recommendation, STIG rule, or related to a GDPR Article/Recital.
				* `cis` - Relevant section from CIS.
				* `gdpr` - Relevant section from GDPR.
				* `obp` - Relevant section from OBP.
				* `stig` - Relevant section from STIG.
			* `remarks` - The explanation of the issue in this finding. It explains the reason for the rule and, if a risk is reported, it may also explain the recommended actions for remediation.
			* `severity` - The severity of the finding.
			* `state` - The current state of the finding.
			* `summary` - The brief summary of the finding. When the finding is informational, the summary typically reports only the number of data elements that were examined.
			* `target_id` - The OCID of the target database.
			* `time_updated` - The date and time the risk level of finding was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
			* `time_valid_until` - The time until which the change in severity(deferred/modified) of this finding is valid.
			* `title` - The short title for the finding.
		* `current` - The particular finding reported by the security assessment.
			* `assessment_id` - The OCID of the assessment that generated this finding.
			* `details` - The details of the finding. Provides detailed information to explain the finding summary, typically results from the assessed database, followed by any recommendations for changes.
			* `has_target_db_risk_level_changed` - Determines if this risk level has changed on the target database since the last time 'severity' was modified by user.
			* `is_risk_modified` - Determines if this risk level was modified by user.
			* `justification` - User provided reason for accepting or modifying this finding if they choose to do so.
			* `key` - A unique identifier for the finding. This is common for the finding across targets.
			* `lifecycle_details` - Details about the current state of the finding.
			* `oracle_defined_severity` - The severity of the finding as determined by security assessment. This cannot be modified by user.
			* `references` - Provides information on whether the finding is related to a CIS Oracle Database Benchmark recommendation, STIG rule, or related to a GDPR Article/Recital.
				* `cis` - Relevant section from CIS.
				* `gdpr` - Relevant section from GDPR.
				* `obp` - Relevant section from OBP.
				* `stig` - Relevant section from STIG.
			* `remarks` - The explanation of the issue in this finding. It explains the reason for the rule and, if a risk is reported, it may also explain the recommended actions for remediation.
			* `severity` - The severity of the finding.
			* `state` - The current state of the finding.
			* `summary` - The brief summary of the finding. When the finding is informational, the summary typically reports only the number of data elements that were examined.
			* `target_id` - The OCID of the target database.
			* `time_updated` - The date and time the risk level of finding was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
			* `time_valid_until` - The time until which the change in severity(deferred/modified) of this finding is valid.
			* `title` - The short title for the finding.
		* `modified_items` - This array contains the items that are present in both the current assessment and the baseline, but are different in the two assessments.
		* `removed_items` - This array identifies the items that are present in the baseline, but are missing from the current assessment.
		* `severity` - The severity of this diff.
	* `user_accounts` - Comparison between findings belonging to User Accounts category.
		* `added_items` - This array identifies the items that are present in the current assessment, but are missing from the baseline.
		* `baseline` - The particular finding reported by the security assessment.
			* `assessment_id` - The OCID of the assessment that generated this finding.
			* `details` - The details of the finding. Provides detailed information to explain the finding summary, typically results from the assessed database, followed by any recommendations for changes.
			* `has_target_db_risk_level_changed` - Determines if this risk level has changed on the target database since the last time 'severity' was modified by user.
			* `is_risk_modified` - Determines if this risk level was modified by user.
			* `justification` - User provided reason for accepting or modifying this finding if they choose to do so.
			* `key` - A unique identifier for the finding. This is common for the finding across targets.
			* `lifecycle_details` - Details about the current state of the finding.
			* `oracle_defined_severity` - The severity of the finding as determined by security assessment. This cannot be modified by user.
			* `references` - Provides information on whether the finding is related to a CIS Oracle Database Benchmark recommendation, STIG rule, or related to a GDPR Article/Recital.
				* `cis` - Relevant section from CIS.
				* `gdpr` - Relevant section from GDPR.
				* `obp` - Relevant section from OBP.
				* `stig` - Relevant section from STIG.
			* `remarks` - The explanation of the issue in this finding. It explains the reason for the rule and, if a risk is reported, it may also explain the recommended actions for remediation.
			* `severity` - The severity of the finding.
			* `state` - The current state of the finding.
			* `summary` - The brief summary of the finding. When the finding is informational, the summary typically reports only the number of data elements that were examined.
			* `target_id` - The OCID of the target database.
			* `time_updated` - The date and time the risk level of finding was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
			* `time_valid_until` - The time until which the change in severity(deferred/modified) of this finding is valid.
			* `title` - The short title for the finding.
		* `current` - The particular finding reported by the security assessment.
			* `assessment_id` - The OCID of the assessment that generated this finding.
			* `details` - The details of the finding. Provides detailed information to explain the finding summary, typically results from the assessed database, followed by any recommendations for changes.
			* `has_target_db_risk_level_changed` - Determines if this risk level has changed on the target database since the last time 'severity' was modified by user.
			* `is_risk_modified` - Determines if this risk level was modified by user.
			* `justification` - User provided reason for accepting or modifying this finding if they choose to do so.
			* `key` - A unique identifier for the finding. This is common for the finding across targets.
			* `lifecycle_details` - Details about the current state of the finding.
			* `oracle_defined_severity` - The severity of the finding as determined by security assessment. This cannot be modified by user.
			* `references` - Provides information on whether the finding is related to a CIS Oracle Database Benchmark recommendation, STIG rule, or related to a GDPR Article/Recital.
				* `cis` - Relevant section from CIS.
				* `gdpr` - Relevant section from GDPR.
				* `obp` - Relevant section from OBP.
				* `stig` - Relevant section from STIG.
			* `remarks` - The explanation of the issue in this finding. It explains the reason for the rule and, if a risk is reported, it may also explain the recommended actions for remediation.
			* `severity` - The severity of the finding.
			* `state` - The current state of the finding.
			* `summary` - The brief summary of the finding. When the finding is informational, the summary typically reports only the number of data elements that were examined.
			* `target_id` - The OCID of the target database.
			* `time_updated` - The date and time the risk level of finding was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
			* `time_valid_until` - The time until which the change in severity(deferred/modified) of this finding is valid.
			* `title` - The short title for the finding.
		* `modified_items` - This array contains the items that are present in both the current assessment and the baseline, but are different in the two assessments.
		* `removed_items` - This array identifies the items that are present in the baseline, but are missing from the current assessment.
		* `severity` - The severity of this diff.
* `time_created` - The date and time when the security assessment comparison was created. Conforms to the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
