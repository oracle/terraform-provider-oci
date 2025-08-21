---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_security_assessments"
sidebar_current: "docs-oci-datasource-data_safe-security_assessments"
description: |-
  Provides the list of Security Assessments in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_security_assessments
This data source provides the list of Security Assessments in Oracle Cloud Infrastructure Data Safe service.

Gets a list of security assessments.

The ListSecurityAssessments operation returns only the assessments in the specified `compartmentId`.
The list does not include any subcompartments of the compartmentId passed.

The parameter `accessLevel` specifies whether to return only those compartments for which the
requestor has INSPECT permissions on at least one resource directly
or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
Principal doesn't have access to even one of the child compartments. This is valid only when
`compartmentIdInSubtree` is set to `true`.

The parameter `compartmentIdInSubtree` applies when you perform ListSecurityAssessments on the
`compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
To get a full list of all compartments and subcompartments in the tenancy (root compartment),
set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.


## Example Usage

```hcl
data "oci_data_safe_security_assessments" "test_security_assessments" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.security_assessment_access_level
	compartment_id_in_subtree = var.security_assessment_compartment_id_in_subtree
	display_name = var.security_assessment_display_name
	is_baseline = var.security_assessment_is_baseline
	is_schedule_assessment = var.security_assessment_is_schedule_assessment
	schedule_assessment_id = oci_data_safe_schedule_assessment.test_schedule_assessment.id
	state = var.security_assessment_state
	target_database_group_id = oci_data_safe_target_database_group.test_target_database_group.id
	target_id = oci_cloud_guard_target.test_target.id
	target_type = var.security_assessment_target_type
	template_assessment_id = oci_data_safe_template_assessment.test_template_assessment.id
	time_created_greater_than_or_equal_to = var.security_assessment_time_created_greater_than_or_equal_to
	time_created_less_than = var.security_assessment_time_created_less_than
	triggered_by = var.security_assessment_triggered_by
	type = var.security_assessment_type
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `display_name` - (Optional) A filter to return only resources that match the specified display name. 
* `is_baseline` - (Optional) A filter to return only the security assessments that are set as a baseline.
* `is_schedule_assessment` - (Optional) A filter to return only security assessments of type save schedule. 
* `schedule_assessment_id` - (Optional) The OCID of the security assessment of type SAVE_SCHEDULE.
* `state` - (Optional) A filter to return only resources that match the specified lifecycle state.
* `target_database_group_id` - (Optional) A filter to return the target database group that matches the specified OCID.
* `target_id` - (Optional) A filter to return only items related to a specific target OCID.
* `target_type` - (Optional) A filter to return only only target database resources or target database group resources.
* `template_assessment_id` - (Optional) The OCID of the security assessment of type TEMPLATE.
* `time_created_greater_than_or_equal_to` - (Optional) A filter to return only the resources that were created after the specified date and time, as defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Using TimeCreatedGreaterThanOrEqualToQueryParam parameter retrieves all resources created after that date.

	**Example:** 2016-12-19T16:39:57.600Z 
* `time_created_less_than` - (Optional) Search for resources that were created before a specific date. Specifying this parameter corresponding `timeCreatedLessThan` parameter will retrieve all resources created before the specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339.

	**Example:** 2016-12-19T16:39:57.600Z 
* `triggered_by` - (Optional) A filter to return only security assessments that were created by either user or system.
* `type` - (Optional) A filter to return only items that match the specified security assessment type.


## Attributes Reference

The following attributes are exported:

* `security_assessments` - The list of security_assessments.

### SecurityAssessment Reference

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

