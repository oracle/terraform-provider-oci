---
subcategory: "Disaster Recovery"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_disaster_recovery_dr_plan_executions"
sidebar_current: "docs-oci-datasource-disaster_recovery-dr_plan_executions"
description: |-
  Provides the list of Dr Plan Executions in Oracle Cloud Infrastructure Disaster Recovery service
---

# Data Source: oci_disaster_recovery_dr_plan_executions
This data source provides the list of Dr Plan Executions in Oracle Cloud Infrastructure Disaster Recovery service.

Get a summary list of all DR plan executions for a DR protection group.

## Example Usage

```hcl
data "oci_disaster_recovery_dr_plan_executions" "test_dr_plan_executions" {
	#Required
	dr_protection_group_id = oci_disaster_recovery_dr_protection_group.test_dr_protection_group.id

	#Optional
	display_name = var.dr_plan_execution_display_name
	dr_plan_execution_id = oci_disaster_recovery_dr_plan_execution.test_dr_plan_execution.id
	state = var.dr_plan_execution_state
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A filter to return only resources that match the given display name.  Example: `MyResourceDisplayName` 
* `dr_plan_execution_id` - (Optional) The OCID of the DR plan execution.  Example: `ocid1.drplanexecution.oc1..uniqueID` 
* `dr_protection_group_id` - (Required) The OCID of the DR protection group. Mandatory query param.  Example: `ocid1.drprotectiongroup.oc1..uniqueID` 
* `state` - (Optional) A filter to return only DR plan executions that match the given lifecycle state. 


## Attributes Reference

The following attributes are exported:

* `dr_plan_execution_collection` - The list of dr_plan_execution_collection.

### DrPlanExecution Reference

The following attributes are exported:

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
	* `time_ended` - The time when group execution ended. An RFC3339 formatted datetime string.  Example: `2019-03-29T09:36:42Z` 
	* `time_started` - The time when group execution began. An RFC3339 formatted datetime string.  Example: `2019-03-29T09:36:42Z` 
	* `type` - The group type.  Example: `BUILT_IN` 
* `id` - The OCID of the DR plan execution.  Example: `ocid1.drplanexecution.oc1..uniqueID` 
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
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time at which DR plan execution was created. An RFC3339 formatted datetime string.  Example: `2019-03-29T09:36:42Z` 
* `time_ended` - The date and time at which DR plan execution succeeded, failed, was paused, or was canceled. An RFC3339 formatted datetime string.  Example: `2019-03-29T09:36:42Z` 
* `time_started` - The date and time at which DR plan execution began. An RFC3339 formatted datetime string.  Example: `2019-03-29T09:36:42Z` 
* `time_updated` - The time when DR plan execution was last updated. An RFC3339 formatted datetime string.  Example: `2019-03-29T09:36:42Z` 

