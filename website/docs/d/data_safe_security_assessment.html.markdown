---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_security_assessment"
sidebar_current: "docs-oci-datasource-data_safe-security_assessment"
description: |-
  Provides details about a specific Security Assessment in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_security_assessment
This data source provides details about a specific Security Assessment resource in Oracle Cloud Infrastructure Data Safe service.

Gets the details of the specified security assessment.

## Example Usage

```hcl
data "oci_data_safe_security_assessment" "test_security_assessment" {
	#Required
	security_assessment_id = oci_data_safe_security_assessment.test_security_assessment.id
}
```

## Argument Reference

The following arguments are supported:

* `security_assessment_id` - (Required) The OCID of the security assessment.


## Attributes Reference

The following attributes are exported:

* `baseline_assessment_id` - The ocid of a security assessment which is of type TEMPLATE_BASELINE, this will be null or empty when type is TEMPLATE_BASELINE.
* `checks` - The security checks to be evaluated for type template.
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
* `compartment_id` - The OCID of the compartment that contains the security assessment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the security assessment.
* `display_name` - The display name of the security assessment.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the security assessment.
* `ignored_assessment_ids` - List containing maps as values. Example: `{"Operations": [ {"CostCenter": "42"} ] }` 
* `ignored_targets` - List containing maps as values. Example: `{"Operations": [ {"CostCenter": "42"} ] }` 
* `is_assessment_scheduled` - Indicates whether the assessment is scheduled to run.
* `is_baseline` - Indicates whether or not the security assessment is set as a baseline. This is applicable only for saved security assessments.
* `is_deviated_from_baseline` - Indicates whether or not the security assessment deviates from the baseline.
* `last_compared_baseline_id` - The OCID of the baseline against which the latest security assessment was compared.
* `lifecycle_details` - Details about the current state of the security assessment.
* `link` - The summary of findings for the security assessment. 
* `schedule` - Schedule of the assessment that runs periodically in the specified format: - <version-string>;<version-specific-schedule>

	Allowed version strings - "v1" v1's version specific schedule -<ss> <mm> <hh> <day-of-week> <day-of-month> Each of the above fields potentially introduce constraints. A workrequest is created only when clock time satisfies all the constraints. Constraints introduced: 1. seconds = <ss> (So, the allowed range for <ss> is [0, 59]) 2. minutes = <mm> (So, the allowed range for <mm> is [0, 59]) 3. hours = <hh> (So, the allowed range for <hh> is [0, 23]) <day-of-week> can be either '*' (without quotes or a number between 1(Monday) and 7(Sunday)) 4. No constraint introduced when it is '*'. When not, day of week must equal the given value <day-of-month> can be either '*' (without quotes or a number between 1 and 28) 5. No constraint introduced when it is '*'. When not, day of month must equal the given value 
* `schedule_security_assessment_id` - The OCID of the security assessment that is responsible for creating this scheduled save assessment.
* `state` - The current state of the security assessment.
* `statistics` - Statistics showing the number of findings for each category grouped by risk levels for all the targets in the specified security assessment.

	The categories include Auditing, Authorization Control, Data Encryption, Database Configuration, Fine-Grained Access Control, Privileges and Roles, and User Accounts. The risk levels include High Risk, Medium Risk, Low Risk, Advisory, Evaluate, and Pass. 
	* `advisory` - Statistics showing the number of findings with a particular risk level for each category.
		* `auditing_findings_count` - The number of findings in the Auditing category.
		* `authorization_control_findings_count` - The number of findings in the Authorization Control category.
		* `data_encryption_findings_count` - The number of findings in the Data Encryption category.
		* `db_configuration_findings_count` - The number of findings in the Database Configuration category.
		* `fine_grained_access_control_findings_count` - The number of findings in the Fine-Grained Access Control category.
		* `privileges_and_roles_findings_count` - The number of findings in the Privileges and Roles category.
		* `targets_count` - The number of targets that contributed to the counts at this risk level.
		* `user_accounts_findings_count` - The number of findings in the User Accounts category.
	* `deferred` - Statistics showing the number of findings with a particular risk level for each category.
		* `auditing_findings_count` - The number of findings in the Auditing category.
		* `authorization_control_findings_count` - The number of findings in the Authorization Control category.
		* `data_encryption_findings_count` - The number of findings in the Data Encryption category.
		* `db_configuration_findings_count` - The number of findings in the Database Configuration category.
		* `fine_grained_access_control_findings_count` - The number of findings in the Fine-Grained Access Control category.
		* `privileges_and_roles_findings_count` - The number of findings in the Privileges and Roles category.
		* `targets_count` - The number of targets that contributed to the counts at this risk level.
		* `user_accounts_findings_count` - The number of findings in the User Accounts category.
	* `evaluate` - Statistics showing the number of findings with a particular risk level for each category.
		* `auditing_findings_count` - The number of findings in the Auditing category.
		* `authorization_control_findings_count` - The number of findings in the Authorization Control category.
		* `data_encryption_findings_count` - The number of findings in the Data Encryption category.
		* `db_configuration_findings_count` - The number of findings in the Database Configuration category.
		* `fine_grained_access_control_findings_count` - The number of findings in the Fine-Grained Access Control category.
		* `privileges_and_roles_findings_count` - The number of findings in the Privileges and Roles category.
		* `targets_count` - The number of targets that contributed to the counts at this risk level.
		* `user_accounts_findings_count` - The number of findings in the User Accounts category.
	* `high_risk` - Statistics showing the number of findings with a particular risk level for each category.
		* `auditing_findings_count` - The number of findings in the Auditing category.
		* `authorization_control_findings_count` - The number of findings in the Authorization Control category.
		* `data_encryption_findings_count` - The number of findings in the Data Encryption category.
		* `db_configuration_findings_count` - The number of findings in the Database Configuration category.
		* `fine_grained_access_control_findings_count` - The number of findings in the Fine-Grained Access Control category.
		* `privileges_and_roles_findings_count` - The number of findings in the Privileges and Roles category.
		* `targets_count` - The number of targets that contributed to the counts at this risk level.
		* `user_accounts_findings_count` - The number of findings in the User Accounts category.
	* `low_risk` - Statistics showing the number of findings with a particular risk level for each category.
		* `auditing_findings_count` - The number of findings in the Auditing category.
		* `authorization_control_findings_count` - The number of findings in the Authorization Control category.
		* `data_encryption_findings_count` - The number of findings in the Data Encryption category.
		* `db_configuration_findings_count` - The number of findings in the Database Configuration category.
		* `fine_grained_access_control_findings_count` - The number of findings in the Fine-Grained Access Control category.
		* `privileges_and_roles_findings_count` - The number of findings in the Privileges and Roles category.
		* `targets_count` - The number of targets that contributed to the counts at this risk level.
		* `user_accounts_findings_count` - The number of findings in the User Accounts category.
	* `medium_risk` - Statistics showing the number of findings with a particular risk level for each category.
		* `auditing_findings_count` - The number of findings in the Auditing category.
		* `authorization_control_findings_count` - The number of findings in the Authorization Control category.
		* `data_encryption_findings_count` - The number of findings in the Data Encryption category.
		* `db_configuration_findings_count` - The number of findings in the Database Configuration category.
		* `fine_grained_access_control_findings_count` - The number of findings in the Fine-Grained Access Control category.
		* `privileges_and_roles_findings_count` - The number of findings in the Privileges and Roles category.
		* `targets_count` - The number of targets that contributed to the counts at this risk level.
		* `user_accounts_findings_count` - The number of findings in the User Accounts category.
	* `pass` - Statistics showing the number of findings with a particular risk level for each category.
		* `auditing_findings_count` - The number of findings in the Auditing category.
		* `authorization_control_findings_count` - The number of findings in the Authorization Control category.
		* `data_encryption_findings_count` - The number of findings in the Data Encryption category.
		* `db_configuration_findings_count` - The number of findings in the Database Configuration category.
		* `fine_grained_access_control_findings_count` - The number of findings in the Fine-Grained Access Control category.
		* `privileges_and_roles_findings_count` - The number of findings in the Privileges and Roles category.
		* `targets_count` - The number of targets that contributed to the counts at this risk level.
		* `user_accounts_findings_count` - The number of findings in the User Accounts category.
	* `targets_count` - The total number of targets in this security assessment.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_database_group_id` - The OCID of the target database group that the group assessment is created for.
* `target_ids` - Array of database target OCIDs.
* `target_type` - Indicates whether the security assessment is for a target database or a target database group.
* `target_version` - The version of the target database.
* `template_assessment_id` - The ocid of a security assessment which is of type TEMPLATE, this will be null or empty when type is TEMPLATE.
* `time_created` - The date and time the security assessment was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_last_assessed` - The date and time the security assessment was last executed, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_updated` - The date and time the security assessment was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `triggered_by` - Indicates whether the security assessment was created by system or by a user.
* `type` - The type of the security assessment. Possible values are:

	LATEST: The most up-to-date assessment that is running automatically for a target. It is system generated. SAVED: A saved security assessment. LATEST assessments are always saved in order to maintain the history of runs. A SAVED assessment is also generated by a 'refresh' action (triggered by the user). SAVE_SCHEDULE: The schedule for periodic saves of LATEST assessments. TEMPLATE: The security assessment contains the checks that the user would like to run. It is user defined. TEMPLATE_BASELINE: The security assessment contains the checks that the user would like to run, together with the max allowed severity. The max allowed severity can be defined by the user. COMPARTMENT: An automatically managed assessment type that stores all details of targets in one compartment. This type keeps an up-to-date assessment of all database risks in one compartment. It is automatically updated when the latest assessment or refresh action is executed. It is also automatically updated when a target is deleted or move to a different compartment. 

