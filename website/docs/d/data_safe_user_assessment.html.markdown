---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_user_assessment"
sidebar_current: "docs-oci-datasource-data_safe-user_assessment"
description: |-
  Provides details about a specific User Assessment in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_user_assessment
This data source provides details about a specific User Assessment resource in Oracle Cloud Infrastructure Data Safe service.

Gets a user assessment by identifier.

## Example Usage

```hcl
data "oci_data_safe_user_assessment" "test_user_assessment" {
	#Required
	user_assessment_id = oci_data_safe_user_assessment.test_user_assessment.id
}
```

## Argument Reference

The following arguments are supported:

* `user_assessment_id` - (Required) The OCID of the user assessment.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the user assessment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the user assessment.
* `display_name` - The display name of the user assessment.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the user assessment.
* `ignored_assessment_ids` - List containing maps as values. Example: `{"Operations": [ {"CostCenter": "42"} ] }` 
* `ignored_targets` - List containing maps as values. Example: `{"Operations": [ {"CostCenter": "42"} ] }` 
* `is_assessment_scheduled` - Indicates whether the assessment is scheduled to run.
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

