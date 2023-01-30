---
subcategory: "Disaster Recovery"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_disaster_recovery_dr_plan_execution"
sidebar_current: "docs-oci-resource-disaster_recovery-dr_plan_execution"
description: |-
  Provides the Dr Plan Execution resource in Oracle Cloud Infrastructure Disaster Recovery service
---

# oci_disaster_recovery_dr_plan_execution
This resource provides the Dr Plan Execution resource in Oracle Cloud Infrastructure Disaster Recovery service.

Execute a DR Plan for a DR Protection Group.

## Example Usage

```hcl
resource "oci_disaster_recovery_dr_plan_execution" "test_dr_plan_execution" {
	#Required
	execution_options {
		#Required
		plan_execution_type = var.dr_plan_execution_execution_options_plan_execution_type

		#Optional
		are_prechecks_enabled = var.dr_plan_execution_execution_options_are_prechecks_enabled
		are_warnings_ignored = var.dr_plan_execution_execution_options_are_warnings_ignored
	}
	plan_id = oci_disaster_recovery_plan.test_plan.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.dr_plan_execution_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) The display name of the DR Plan Execution.  Example: `Execution - EBS Switchover PHX to IAD` 
* `execution_options` - (Required) The options for a plan execution.
	* `are_prechecks_enabled` - (Applicable when plan_execution_type=FAILOVER | SWITCHOVER) A flag indicating whether a precheck should be executed before the plan.  Example: `false` 
	* `are_warnings_ignored` - (Optional) A flag indicating whether warnigs should be ignored during the switchover.  Example: `true` 
	* `plan_execution_type` - (Required) The type of the plan execution. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"Department": "Finance"}` 
* `plan_id` - (Required) The OCID of the DR Plan.  Example: `ocid1.drplan.oc1.iad.exampleocid2` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing this DR Plan Execution.  Example: `ocid1.compartment.oc1..exampleocid1` 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The display name of this DR Plan Execution.  Example: `Execution - EBS Switchover PHX to IAD` 
* `dr_protection_group_id` - The OCID of the DR Protection Group to which this DR Plan Execution belongs.  Example: `ocid1.drprotectiongroup.oc1.iad.exampleocid2` 
* `execution_duration_in_sec` - The total duration in seconds taken to complete the DR Plan Execution.  Example: `750` 
* `execution_options` - The options for a plan execution.
	* `are_prechecks_enabled` - A flag indicating whether a precheck was executed before the plan.  Example: `false` 
	* `are_warnings_ignored` - A flag indicating whether warnigs was ignored during the switchover.  Example: `true` 
	* `plan_execution_type` - The type of the plan execution. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"Department": "Finance"}` 
* `group_executions` - A list of groups executed in this DR Plan Execution. 
	* `display_name` - The display name of group that was executed.  Example: `DATABASE_SWITCHOVER` 
	* `execution_duration_in_sec` - The total duration in seconds taken to complete group execution.  Example: `120` 
	* `group_id` - The unique id of the group. Must not be modified by user.  Example: `sgid1.group..examplegroupsgid` 
	* `status` - The status of the group execution. 
	* `status_details` - Additional details about the group execution status.  Example: `A total of three steps failed in the group` 
	* `step_executions` - A list of details of each step executed in this group. 
		* `display_name` - The display name of the step.  Example: `DATABASE_SWITCHOVER` 
		* `execution_duration_in_sec` - The total duration in seconds taken to complete step execution.  Example: `35` 
		* `group_id` - The unique id of the group to which this step belongs. Must not be modified by user.  Example: `sgid1.group..examplegroupsgid` 
		* `log_location` - Information about an Object Storage log location for a DR Protection Group.
			* `bucket` - The bucket name inside the Object Storage namespace.  Example: `operation_logs` 
			* `namespace` - The namespace in Object Storage (Note - this is usually the tenancy name).  Example: `myocitenancy` 
			* `object` - The object name inside the Object Storage bucket.  Example: `switchover_plan_executions` 
		* `status` - The status of the step execution. 
		* `status_details` - Additional details about the step execution status.  Example: `This step failed to complete due to a timeout` 
		* `step_id` - The unique id of this step. Must not be modified by user.  Example: `sgid1.step..examplestepsgid` 
		* `time_ended` - The time at which step execution ended. An RFC3339 formatted datetime string.  Example: `2019-03-29T09:36:42Z` 
		* `time_started` - The time at which step execution began. An RFC3339 formatted datetime string.  Example: `2019-03-29T09:36:42Z` 
		* `type` - The plan step type. 
	* `time_ended` - The time at which group execution ended. An RFC3339 formatted datetime string.  Example: `2019-03-29T09:36:42Z` 
	* `time_started` - The time at which group execution began. An RFC3339 formatted datetime string.  Example: `2019-03-29T09:36:42Z` 
	* `type` - The plan group type. 
* `id` - The OCID of the DR Plan Execution.  Example: `ocid1.drplanexecution.oc1.iad.exampleocid2` 
* `life_cycle_details` - A message describing the DR Plan Execution's current state in more detail.  Example: `The DR Plan Execution [Execution - EBS Switchover PHX to IAD] is currently in progress` 
* `log_location` - Information about an Object Storage log location for a DR Protection Group.
	* `bucket` - The bucket name inside the Object Storage namespace.  Example: `operation_logs` 
	* `namespace` - The namespace in Object Storage (Note - this is usually the tenancy name).  Example: `myocitenancy` 
	* `object` - The object name inside the Object Storage bucket.  Example: `switchover_plan_executions` 
* `peer_dr_protection_group_id` - The OCID of peer (remote) DR Protection Group associated with this plan's DR Protection Group.  Example: `ocid1.drprotectiongroup.oc1.phx.exampleocid1` 
* `peer_region` - The region of the peer (remote) DR Protection Group.  Example: `us-ashburn-1` 
* `plan_execution_type` - The type of the DR Plan executed. 
* `plan_id` - The OCID of the DR Plan.  Example: `ocid1.drplan.oc1.iad.exampleocid2` 
* `state` - The current state of the DR Plan Execution. 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time at which DR Plan Execution was created. An RFC3339 formatted datetime string.  Example: `2019-03-29T09:36:42Z` 
* `time_ended` - The date and time at which DR Plan Execution succeeded, failed, was paused, or was canceled. An RFC3339 formatted datetime string.  Example: `2019-03-29T09:36:42Z` 
* `time_started` - The date and time at which DR Plan Execution began. An RFC3339 formatted datetime string.  Example: `2019-03-29T09:36:42Z` 
* `time_updated` - The time at which DR Plan Execution was last updated. An RFC3339 formatted datetime string.  Example: `2019-03-29T09:36:42Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Dr Plan Execution
	* `update` - (Defaults to 20 minutes), when updating the Dr Plan Execution
	* `delete` - (Defaults to 20 minutes), when destroying the Dr Plan Execution


## Import

DrPlanExecutions can be imported using the `id`, e.g.

```
$ terraform import oci_disaster_recovery_dr_plan_execution.test_dr_plan_execution "id"
```

