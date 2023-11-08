---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_user_assessments"
sidebar_current: "docs-oci-datasource-data_safe-user_assessments"
description: |-
  Provides the list of User Assessments in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_user_assessments
This data source provides the list of User Assessments in Oracle Cloud Infrastructure Data Safe service.

Gets a list of user assessments.

The ListUserAssessments operation returns only the assessments in the specified `compartmentId`.
The list does not include any subcompartments of the compartmentId passed.

The parameter `accessLevel` specifies whether to return only those compartments for which the
requestor has INSPECT permissions on at least one resource directly
or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
Principal doesn't have access to even one of the child compartments. This is valid only when
`compartmentIdInSubtree` is set to `true`.

The parameter `compartmentIdInSubtree` applies when you perform ListUserAssessments on the
`compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
To get a full list of all compartments and subcompartments in the tenancy (root compartment),
set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.


## Example Usage

```hcl
data "oci_data_safe_user_assessments" "test_user_assessments" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.user_assessment_access_level
	compartment_id_in_subtree = var.user_assessment_compartment_id_in_subtree
	display_name = var.user_assessment_display_name
	is_baseline = var.user_assessment_is_baseline
	is_schedule_assessment = var.user_assessment_is_schedule_assessment
	schedule_user_assessment_id = oci_data_safe_user_assessment.test_user_assessment.id
	state = var.user_assessment_state
	target_id = oci_cloud_guard_target.test_target.id
	time_created_greater_than_or_equal_to = var.user_assessment_time_created_greater_than_or_equal_to
	time_created_less_than = var.user_assessment_time_created_less_than
	triggered_by = var.user_assessment_triggered_by
	type = var.user_assessment_type
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `display_name` - (Optional) A filter to return only resources that match the specified display name. 
* `is_baseline` - (Optional) A filter to return only user assessments that are set as baseline.
* `is_schedule_assessment` - (Optional) A filter to return only user assessments of type SAVE_SCHEDULE. 
* `schedule_user_assessment_id` - (Optional) The OCID of the user assessment of type SAVE_SCHEDULE.
* `state` - (Optional) The current state of the user assessment.
* `target_id` - (Optional) A filter to return only items related to a specific target OCID.
* `time_created_greater_than_or_equal_to` - (Optional) A filter to return only user assessments that were created after the specified date and time, as defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Using timeCreatedGreaterThanOrEqualTo parameter retrieves all assessments created after that date.

	**Example:** 2016-12-19T16:39:57.600Z 
* `time_created_less_than` - (Optional) Search for resources that were created before a specific date. Specifying this parameter corresponding `timeCreatedLessThan` parameter will retrieve all resources created before the specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339.

	**Example:** 2016-12-19T16:39:57.600Z 
* `triggered_by` - (Optional) A filter to return user assessments that were created by either the system or by a user only.
* `type` - (Optional) A filter to return only items that match the specified assessment type.


## Attributes Reference

The following attributes are exported:

* `user_assessments` - The list of user_assessments.

### UserAssessment Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the user assessment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the user assessment.
* `display_name` - The display name of the user assessment.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the user assessment.
* `ignored_assessment_ids` - List containing maps as values. Example: `{"Operations": [ {"CostCenter": "42"} ] }` 
* `ignored_targets` - List containing maps as values. Example: `{"Operations": [ {"CostCenter": "42"} ] }` 
* `is_baseline` - Indicates if the user assessment is set as a baseline. This is applicable only to saved user assessments.
* `is_deviated_from_baseline` - Indicates if the user assessment deviates from the baseline.
* `last_compared_baseline_id` - The OCID of the last user assessment baseline against which the latest assessment was compared.
* `lifecycle_details` - Details about the current state of the user assessment.
* `schedule` - Schedule of the assessment that runs periodically in this specified format: <version-string>;<version-specific-schedule>

	Allowed version strings - "v1" v1's version specific schedule -<ss> <mm> <hh> <day-of-week> <day-of-month> Each of the above fields potentially introduce constraints. A workrequest is created only when clock time satisfies all the constraints. Constraints introduced: 1. seconds = <ss> (So, the allowed range for <ss> is [0, 59]) 2. minutes = <mm> (So, the allowed range for <mm> is [0, 59]) 3. hours = <hh> (So, the allowed range for <hh> is [0, 23]) <day-of-week> can be either '*' (without quotes or a number between 1(Monday) and 7(Sunday)) 4. No constraint introduced when it is '*'. When not, day of week must equal the given value <day-of-month> can be either '*' (without quotes or a number between 1 and 28) 5. No constraint introduced when it is '*'. When not, day of month must equal the given value 
* `schedule_assessment_id` - The OCID of the user assessment that is responsible for creating this scheduled save assessment.
* `state` - The current state of the user assessment.
* `statistics` - Map that contains maps of values. Example: `{"Operations": {"CostCenter": "42"}}` 
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_ids` - Array of database target OCIDs.
* `time_created` - The date and time when the user assessment was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_last_assessed` - The date and time the user assessment was last run, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_updated` - The last date and time when the user assessment was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `triggered_by` - Indicates whether the user assessment was created by system or user.
* `type` - Type of user assessment. Type can be:

	LATEST: The most up-to-date assessment that is running automatically for a target. It is system generated. SAVED: A saved user assessment. LATEST assessments will always be saved to maintain the history of runs. A SAVED assessment is also generated by a 'refresh' action (triggered by the user). SAVE_SCHEDULE: A schedule to periodically save LATEST assessments. COMPARTMENT: An automatic managed assessment type that stores all details of targets in one compartment. This will keep an up-to-date status of all potential risks identified in the compartment. It also keeps track of user count and target count for each profile available on the targets in a given compartment.  It is automatically updated once the latest assessment or refresh action is executed, as well as when a target is deleted or moved to a different compartment. 

