---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_runbook_versions"
sidebar_current: "docs-oci-datasource-fleet_apps_management-runbook_versions"
description: |-
  Provides the list of Runbook Versions in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_runbook_versions
This data source provides the list of Runbook Versions in Oracle Cloud Infrastructure Fleet Apps Management service.

List versions for a runbook in Fleet Application Management.


## Example Usage

```hcl
data "oci_fleet_apps_management_runbook_versions" "test_runbook_versions" {

	#Optional
	compartment_id = var.compartment_id
	id = var.runbook_version_id
	name = var.runbook_version_name
	runbook_id = oci_fleet_apps_management_runbook.test_runbook.id
	state = var.runbook_version_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources. Empty only if the resource OCID query param is not specified. 
* `id` - (Optional) A filter to return runbook versions whose identifier matches the given identifier.
* `name` - (Optional) A filter to return only resources that match the entire name given.
* `runbook_id` - (Optional) A filter to return only schedule definitions whose associated runbookId matches the given runbookId.
* `state` - (Optional) A filter to return only resources whose lifecycleState matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `runbook_version_collection` - The list of runbook_version_collection.

### RunbookVersion Reference

The following attributes are exported:

* `compartment_id` - 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `execution_workflow_details` - Execution Workflow details.
	* `workflow` - Execution Workflow for the runbook.
		* `group_name` - Name of the group.
		* `steps` - Steps within the Group.
			* `group_name` - Name of the group.
			* `step_name` - Provide StepName for the Task.
			* `steps` - Tasks within the Group. Provide the stepName for all applicable tasks. 
			* `type` - Content Source Details. 
		* `type` - Workflow Group  Details. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `groups` - The groups of the runbook. 
	* `name` - The name of the group.
	* `properties` - The properties of the component.
		* `action_on_failure` - The action to be taken in case of a failure.
		* `notification_preferences` - Preferences to send notifications on the task activities.
			* `should_notify_on_pause` - Enables notification on pause.
			* `should_notify_on_task_failure` - Enables or disables notification on Task Failures.
			* `should_notify_on_task_success` - Enables or disables notification on Task Success.
		* `pause_details` - Pause Details
			* `duration_in_minutes` - Time in minutes to apply Pause.
			* `kind` - Pause based On. 
		* `pre_condition` - Build control flow conditions that determine the relevance of the task execution. 
		* `run_on` - The runon conditions
			* `condition` - Build control flow conditions that determine the relevance of the task execution against targets. 
			* `host` - OCID of the self hosted instance.
			* `kind` - Run on based On. 
			* `previous_task_instance_details` - Previous Task Instance Details 
				* `output_variable_details` - The details of the output variable that will be used for mapping.
					* `output_variable_name` - The name of the output variable whose value has to be mapped.
					* `step_name` - The name of the task step the output variable belongs to.
				* `resource_id` - Resource Ocid.
				* `resource_type` - Resource Type.
	* `type` - The type of the group. PARALLEL_TASK_GROUP : Helps to execute tasks parallelly inside a resource. PARALLEL_RESOURCE_GROUP : Executes tasks across resources parallelly. ROLLING_RESOURCE_GROUP : Executes tasks across resources in a rolling order. 
* `id` - The OCID of the resource.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `name` - The version of the runbook.
* `rollback_workflow_details` - Rollback Workflow details.
	* `scope` - rollback Scope 
	* `workflow` - Rollback Workflow for the runbook.
		* `group_name` - Name of the group.
		* `steps` - Steps within the Group.
			* `group_name` - Name of the group.
			* `step_name` - Provide StepName for the Task.
			* `steps` - Tasks within the Group. Provide the stepName for all applicable tasks. 
			* `type` - Content Source Details. 
		* `type` - Workflow Group  Details. 
* `runbook_id` - The OCID of the resource.
* `state` - The current state of the runbook version.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `tasks` - A set of tasks to execute in the runbook.
	* `output_variable_mappings` - Mapping output variables of previous tasks to the input variables of the current task.
		* `name` - The name of the input variable.
		* `output_variable_details` - The details of the output variable that will be used for mapping.
			* `output_variable_name` - The name of the output variable whose value has to be mapped.
			* `step_name` - The name of the task step the output variable belongs to.
	* `step_name` - The name of the task step.
	* `step_properties` - The properties of the component.
		* `action_on_failure` - The action to be taken in case of a failure.
		* `notification_preferences` - Preferences to send notifications on the task activities.
			* `should_notify_on_pause` - Enables notification on pause.
			* `should_notify_on_task_failure` - Enables or disables notification on Task Failures.
			* `should_notify_on_task_success` - Enables or disables notification on Task Success.
		* `pause_details` - Pause Details
			* `duration_in_minutes` - Time in minutes to apply Pause.
			* `kind` - Pause based On. 
		* `pre_condition` - Build control flow conditions that determine the relevance of the task execution. 
		* `run_on` - The runon conditions
			* `condition` - Build control flow conditions that determine the relevance of the task execution against targets. 
			* `host` - OCID of the self hosted instance.
			* `kind` - Run on based On. 
			* `previous_task_instance_details` - Previous Task Instance Details 
				* `output_variable_details` - The details of the output variable that will be used for mapping.
					* `output_variable_name` - The name of the output variable whose value has to be mapped.
					* `step_name` - The name of the task step the output variable belongs to.
				* `resource_id` - Resource Ocid.
				* `resource_type` - Resource Type.
	* `task_record_details` - The details of the task.
		* `description` - The description of the task.
		* `execution_details` - Execution details.
			* `catalog_id` - Catalog Id having terraform package.
			* `command` - Optional command to execute the content. You can provide any commands/arguments that can't be part of the script. 
			* `config_file` - Catalog Id having config file.
			* `content` - Content Source details.
				* `bucket` - Bucket Name.
				* `catalog_id` - Catalog Id having terraform package.
				* `checksum` - md5 checksum of the artifact.
				* `namespace` - Namespace.
				* `object` - Object Name.
				* `source_type` - Content Source type details. 
			* `credentials` - Credentials required for executing the task. 
				* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
				* `id` - The OCID of the resource.
			* `endpoint` - Endpoint to be invoked.
			* `execution_type` - The action type of the task
			* `is_executable_content` - Is the Content an executable file?
			* `is_locked` - Is the script locked to prevent changes directly in Object Storage?
			* `is_read_output_variable_enabled` - Is read output variable enabled
			* `system_variables` - The list of system variables.
			* `target_compartment_id` - OCID of the compartment to which the resource belongs to.
			* `variables` - The variable of the task. At least one of the dynamicArguments or output needs to be provided. 
				* `input_variables` - The input variables for the task.
					* `description` - The description of the argument.
					* `name` - The name of the argument.
					* `type` - Input argument Type. 
				* `output_variables` - The list of output variables.
		* `is_apply_subject_task` - Is this an Apply Subject Task? Ex. Patch Execution Task
		* `is_copy_to_library_enabled` - Make a copy of this task in Library
		* `is_discovery_output_task` - Is this a discovery output task?
		* `name` - The name of the task
		* `os_type` - The OS for the task.
		* `platform` - The platform of the runbook.
		* `properties` - The properties of the task.
			* `num_retries` - The number of retries allowed.
			* `timeout_in_seconds` - The timeout in seconds for the task.
		* `scope` - The scope of the task.
		* `task_record_id` - The ID of taskRecord.
* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.

