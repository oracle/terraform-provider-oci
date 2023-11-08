---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_security_assessment"
sidebar_current: "docs-oci-resource-data_safe-security_assessment"
description: |-
  Provides the Security Assessment resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_security_assessment
This resource provides the Security Assessment resource in Oracle Cloud Infrastructure Data Safe service.

Creates a new saved security assessment for one or multiple targets in a compartment. When this operation is performed,
it will save the latest assessments in the specified compartment. If a schedule is passed, it will persist the latest assessments,
at the defined date and time, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).


## Example Usage

```hcl
resource "oci_data_safe_security_assessment" "test_security_assessment" {
	#Required
	compartment_id = var.compartment_id
	target_id = oci_cloud_guard_target.test_target.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.security_assessment_description
	display_name = var.security_assessment_display_name
	freeform_tags = {"Department"= "Finance"}
	schedule = var.security_assessment_schedule
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment that contains the security assessment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) Description of the security assessment.
* `display_name` - (Optional) (Updatable) The display name of the security assessment.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `schedule` - (Optional) (Updatable) To schedule the assessment for running periodically, specify the schedule in this attribute. Create or schedule one assessment per compartment. If not defined, the assessment runs immediately. Format - <version-string>;<version-specific-schedule>

	Allowed version strings - "v1" v1's version specific schedule -<ss> <mm> <hh> <day-of-week> <day-of-month> Each of the above fields potentially introduce constraints. A workrequest is created only when clock time satisfies all the constraints. Constraints introduced: 1. seconds = <ss> (So, the allowed range for <ss> is [0, 59]) 2. minutes = <mm> (So, the allowed range for <mm> is [0, 59]) 3. hours = <hh> (So, the allowed range for <hh> is [0, 23]) <day-of-week> can be either '*' (without quotes or a number between 1(Monday) and 7(Sunday)) 4. No constraint introduced when it is '*'. When not, day of week must equal the given value <day-of-month> can be either '*' (without quotes or a number between 1 and 28) 5. No constraint introduced when it is '*'. When not, day of month must equal the given value 
* `target_id` - (Required) The OCID of the target database on which security assessment is to be run.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the security assessment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the security assessment.
* `display_name` - The display name of the security assessment.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the security assessment.
* `ignored_assessment_ids` - List containing maps as values. Example: `{"Operations": [ {"CostCenter": "42"} ] }` 
* `ignored_targets` - List containing maps as values. Example: `{"Operations": [ {"CostCenter": "42"} ] }` 
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
* `target_ids` - Array of database target OCIDs.
* `target_version` - The version of the target database.
* `time_created` - The date and time when the security assessment was created. Conforms to the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_last_assessed` - The date and time when the security assessment was last run. Conforms to the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_updated` - The date and time when the security assessment was last updated. Conforms to the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `triggered_by` - Indicates whether the security assessment was created by system or by a user.
* `type` - The type of this security assessment. The possible types are:

	LATEST: The most up-to-date assessment that is running automatically for a target. It is system generated. SAVED: A saved security assessment. LATEST assessments are always saved in order to maintain the history of runs. A SAVED assessment is also generated by a 'refresh' action (triggered by the user). SAVE_SCHEDULE: The schedule for periodic saves of LATEST assessments. COMPARTMENT: An automatically managed assessment type that stores all details of targets in one compartment. This type keeps an up-to-date assessment of all database risks in one compartment. It is automatically updated when the latest assessment or refresh action is executed. It is also automatically updated when a target is deleted or move to a different compartment. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Security Assessment
	* `update` - (Defaults to 20 minutes), when updating the Security Assessment
	* `delete` - (Defaults to 20 minutes), when destroying the Security Assessment


## Import

SecurityAssessments can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_security_assessment.test_security_assessment "id"
```

