---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_user_assessment"
sidebar_current: "docs-oci-resource-data_safe-user_assessment"
description: |-
  Provides the User Assessment resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_user_assessment
This resource provides the User Assessment resource in Oracle Cloud Infrastructure Data Safe service.

Creates a new saved user assessment for one or multiple targets in a compartment. It saves the latest assessments in the
specified compartment. If a scheduled is passed in, this operation persists the latest assessments that exist at the defined
date and time, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).


## Example Usage

```hcl
resource "oci_data_safe_user_assessment" "test_user_assessment" {
	#Required
	compartment_id = var.compartment_id
	target_id = oci_cloud_guard_target.test_target.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.user_assessment_description
	display_name = var.user_assessment_display_name
	freeform_tags = {"Department"= "Finance"}
	schedule = var.user_assessment_schedule
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment that contains the user assessment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) The description of the user assessment.
* `display_name` - (Optional) (Updatable) The display name of the user assessment.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `schedule` - (Optional) (Updatable) To schedule the assessment for saving periodically, specify the schedule in this attribute. Create or schedule one assessment per compartment. If not defined, the assessment runs immediately. Format - <version-string>;<version-specific-schedule>

	Allowed version strings - "v1" v1's version specific schedule -<ss> <mm> <hh> <day-of-week> <day-of-month> Each of the above fields potentially introduce constraints. A workrequest is created only when clock time satisfies all the constraints. Constraints introduced: 1. seconds = <ss> (So, the allowed range for <ss> is [0, 59]) 2. minutes = <mm> (So, the allowed range for <mm> is [0, 59]) 3. hours = <hh> (So, the allowed range for <hh> is [0, 23]) <day-of-week> can be either '*' (without quotes or a number between 1(Monday) and 7(Sunday)) 4. No constraint introduced when it is '*'. When not, day of week must equal the given value <day-of-month> can be either '*' (without quotes or a number between 1 and 28) 5. No constraint introduced when it is '*'. When not, day of month must equal the given value 
* `target_id` - (Required) The OCID of the target database on which the user assessment is to be run.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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
* `time_created` - The date and time the user assessment was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_last_assessed` - The date and time the user assessment was last executed, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_updated` - The date and time the user assessment was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `triggered_by` - Indicates whether the user assessment was created by the system or the user.
* `type` - The type of the user assessment. The possible types are:

	LATEST: The latest assessment that was executed for a target. It can either be system generated as part of the scheduled assessments or user driven by refreshing the latest assessment. SAVED: A saved user assessment. All user assessments are saved in the user assessment history. SAVE_SCHEDULE: The schedule to periodically save the LATEST assessment of a target database. COMPARTMENT: An automatic managed assessment type that stores all details of the targets in one compartment. This will keep an up-to-date status of all potential risks identified in the compartment. It also keeps track of user count and target count for each profile available on the targets in a given compartment.  It is automatically updated once the latest assessment or refresh action is executed, as well as when a target is deleted or moved to a different compartment. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the User Assessment
	* `update` - (Defaults to 20 minutes), when updating the User Assessment
	* `delete` - (Defaults to 20 minutes), when destroying the User Assessment


## Import

UserAssessments can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_user_assessment.test_user_assessment "id"
```

