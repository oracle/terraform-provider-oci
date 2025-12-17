---
subcategory: "Disaster Recovery"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_disaster_recovery_dr_plan_execution"
sidebar_current: "docs-oci-datasource-disaster_recovery-dr_plan_execution"
description: |-
  Provides details about a specific Dr Plan Execution in Oracle Cloud Infrastructure Disaster Recovery service
---

# Data Source: oci_disaster_recovery_dr_plan_execution
This data source provides details about a specific Dr Plan Execution resource in Oracle Cloud Infrastructure Disaster Recovery service.

Get details for the DR plan execution identified by *drPlanExecutionId*.

## Example Usage

```hcl
data "oci_disaster_recovery_dr_plan_execution" "test_dr_plan_execution" {
	#Required
	dr_plan_execution_id = oci_disaster_recovery_dr_plan_execution.test_dr_plan_execution.id
}
```

## Argument Reference

The following arguments are supported:

* `dr_plan_execution_id` - (Required) The OCID of the DR plan execution.  Example: `ocid1.drplanexecution.oc1..uniqueID` 


## Attributes Reference

The following attributes are exported:

* `automatic_execution_details` - The details of the event that started the automatic DR plan execution.
	* `event_name` - The name of the Oracle Cloud Infrastructure event that started the automatic DR plan execution.  Example: `SwitchoverAutonomousDatabase` 
	* `member_id` - The OCID of the member that emitted the event that started the automatic DR plan execution.  Example: "ocid1.autonomousdatabase.oc1..uniqueID" 
* `compartment_id` - The OCID of the compartment containing this DR plan execution.  Example: `ocid1.compartment.oc1..uniqueID` 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The display name of the DR plan execution.  Example: `Execution - EBS Switchover PHX to IAD` 
* `dr_protection_group_id` - The OCID of the DR protection group to which this DR plan execution belongs.  Example: `ocid1.drprotectiongroup.oc1..uniqueID` 
* `execution_duration_in_sec` - The total duration in seconds taken to complete the DR plan execution.  Example: `750` 
* `execution_options` - The options for a plan execution.
	* `are_prechecks_enabled` - A flag indicating whether a precheck should be executed before the plan execution.  Example: `true` 
	* `are_warnings_ignored` - A flag indicating whether warnings should be ignored during the plan execution.  Example: `false` 
	* `plan_execution_type` - The type of the plan execution. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  Example: `{"Department": "Finance"}` 
* `group_executions` - A list of groups executed in this DR plan execution. 
	* `display_name` - The display name of the group execution.  Example: `DATABASE_SWITCHOVER` 
	* `execution_duration_in_sec` - The total duration in seconds taken to complete group execution.  Example: `120` 
	* `group_id` - The unique id of the group. Must not be modified by user.  Example: `sgid1.group..uniqueID` 
	* `status` - The status of the group execution. 
	* `status_details` - Additional details on the group execution status.  Example: `A total of [3] steps failed in the group` 
	* `step_executions` - A list of step executions in the group. 
		* `display_name` - The display name of the step execution.  Example: `DATABASE_SWITCHOVER` 
		* `execution_duration_in_sec` - The total duration in seconds taken to complete the step execution.  Example: `35` 
		* `group_id` - The unique id of the group to which this step belongs. Must not be modified by user.  Example: `sgid1.group..uniqueID` 
		* `log_location` - The details of an object storage log location for a DR protection group.
			* `bucket` - The bucket name inside the object storage namespace.  Example: `operation_logs` 
			* `namespace` - The namespace in object storage (Note - this is usually the tenancy name).  Example: `myocitenancy` 
			* `object` - The object name inside the object storage bucket.  Example: `switchover_plan_executions` 
		* `status` - The status of the step execution. 
		* `status_details` - Additional details on the step execution status.  Example: `This step failed to complete due to a timeout` 
		* `step_id` - The unique id of the step. Must not be modified by user.  Example: `sgid1.step..uniqueID` 
		* `time_ended` - The time when execution ended. An RFC3339 formatted datetime string.  Example: `2019-03-29T09:36:42Z` 
		* `time_started` - The time when step execution began. An RFC3339 formatted datetime string.  Example: `2019-03-29T09:36:42Z` 
		* `type` - The step type. 
		* `type_display_name` - The display name of the DR Plan step type.  Example: `Database Switchover` 
	* `time_ended` - The time when group execution ended. An RFC3339 formatted datetime string.  Example: `2019-03-29T09:36:42Z` 
	* `time_started` - The time when group execution began. An RFC3339 formatted datetime string.  Example: `2019-03-29T09:36:42Z` 
	* `type` - The group type.  Example: `BUILT_IN` 
* `id` - The OCID of the DR plan execution.  Example: `ocid1.drplanexecution.oc1..uniqueID` 
* `is_automatic` - A flag indicating whether execution was submitted automatically by Automatic DR Configuration.  Example: `false` 
* `life_cycle_details` - A message describing the DR plan execution's current state in more detail. 
* `log_location` - The details of an object storage log location for a DR protection group.
	* `bucket` - The bucket name inside the object storage namespace.  Example: `operation_logs` 
	* `namespace` - The namespace in object storage (Note - this is usually the tenancy name).  Example: `myocitenancy` 
	* `object` - The object name inside the object storage bucket.  Example: `switchover_plan_executions` 
* `peer_dr_protection_group_id` - The OCID of peer DR protection group associated with this plan's DR protection group.  Example: `ocid1.drprotectiongroup.oc1..uniqueID` 
* `peer_region` - The region of the peer DR protection group associated with this plan's DR protection group.  Example: `us-ashburn-1` 
* `plan_execution_type` - The type of the DR plan executed. 
* `plan_id` - The OCID of the DR plan.  Example: `ocid1.drplan.oc1..uniqueID` 
* `state` - The current state of the DR plan execution. 
* `step_status_counts` - A categorized summary of step execution statuses and their corresponding counts. 
	* `failed_steps` - A summary of steps that failed during a DR plan execution, including failed and timed out steps. 
		* `failed` - The total number of failed steps in a DR plan execution. 
		* `timed_out` - The total number of steps that timed out during a DR plan execution. 
		* `total_failed` - The total number of steps that failed during a DR plan execution. 
	* `remaining_steps` - A summary of remaining steps in a DR plan execution, including queued, paused, and in-progress steps. 
		* `in_progress` - The total number of steps in progress during a DR plan execution. 
		* `paused` - The total number of paused steps in a DR plan execution. 
		* `queued` - The total number of queued steps in a DR plan execution. 
		* `total_remaining` - The total number of remaining steps in a DR plan execution. 
	* `skipped_steps` - A summary of steps that were skipped during a DR plan execution, including disabled, failed but ignored, timed out but ignored, and canceled steps. 
		* `canceled` - The total number of canceled steps in a DR plan execution. 
		* `disabled` - The total number of disabled steps in a DR plan execution. 
		* `failed_ignored` - The total number of steps that failed but were ignored during a DR plan execution. 
		* `timed_out_ignored` - The total number of steps that timed out but were ignored during a DR plan execution. 
		* `total_skipped` - The total number of steps that were skipped during a DR plan execution. 
	* `successful_steps` - A summary of steps that completed successfully during a DR plan execution. 
		* `succeeded` - The total number of steps that succeeded during a DR plan execution. 
		* `total_successful` - The total number of successful steps in a DR plan execution. 
	* `total_steps` - The total number of steps in a DR plan execution. 
	* `warning_steps` - A summary of steps that encountered warnings during a DR plan execution. 
		* `total_warnings` - The total number of steps that encountered warnings in a DR plan execution. 
		* `warnings_ignored` - The total number of steps with warnings that were ignored during a DR plan execution. 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time at which DR plan execution was created. An RFC3339 formatted datetime string.  Example: `2019-03-29T09:36:42Z` 
* `time_ended` - The date and time at which DR plan execution succeeded, failed, was paused, or was canceled. An RFC3339 formatted datetime string.  Example: `2019-03-29T09:36:42Z` 
* `time_started` - The date and time at which DR plan execution began. An RFC3339 formatted datetime string.  Example: `2019-03-29T09:36:42Z` 
* `time_updated` - The time when DR plan execution was last updated. An RFC3339 formatted datetime string.  Example: `2019-03-29T09:36:42Z` 

