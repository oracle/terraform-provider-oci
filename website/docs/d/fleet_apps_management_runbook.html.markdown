---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_runbook"
sidebar_current: "docs-oci-datasource-fleet_apps_management-runbook"
description: |-
  Provides details about a specific Runbook in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_runbook
This data source provides details about a specific Runbook resource in Oracle Cloud Infrastructure Fleet Apps Management service.

Gets a Runbook by identifier

## Example Usage

```hcl
data "oci_fleet_apps_management_runbook" "test_runbook" {
	#Required
	runbook_id = oci_fleet_apps_management_runbook.test_runbook.id
}
```

## Argument Reference

The following arguments are supported:

* `runbook_id` - (Required) Unique Runbook identifier


## Attributes Reference

The following attributes are exported:

* `associations` - JSON content with required associations
	* `execution_workflow_details` - Execution Workflow.
		* `workflow` - Execution Workflow for the runbook.
			* `group_name` - Provide the name of the group.
			* `steps` - Steps within the Group.
				* `group_name` - Provide the name of the group.
				* `step_name` - Provide StepName for the Task.
				* `steps` - Tasks within the Group. Provide the stepName for all tasks that are applicable 
				* `type` - Content Source Details. 
			* `type` - Workflow Group  Details. 
	* `groups` - The groups of the runbook
		* `name` - The name of the group
		* `properties` - The properties of the task.
			* `action_on_failure` - The action to be taken in case of task failure.
			* `condition` - The condition in which the task is to be executed.
			* `run_on` - The hosts to execute on.
		* `type` - The type of the group
	* `tasks` - A set of tasks to execute in the runbook
		* `association_type` - The association type of the task
		* `output_variable_mappings` - Mapping output variables of previous tasks to the input variables of the current task.
			* `name` - The name of the input variable
			* `output_variable_details` - The details of the output variable that will be used for mapping.
				* `output_variable_name` - The name of the output variable whose value that has to be mapped.
				* `step_name` - The name of the task step the output variable belongs to.
		* `step_name` - The name of the task step.
		* `step_properties` - The properties of the task.
			* `action_on_failure` - The action to be taken in case of task failure.
			* `condition` - The condition in which the task is to be executed.
			* `run_on` - The hosts to execute on.
		* `task_record_details` - The details of the task.
			* `description` - The description of the task.
			* `execution_details` - Content Source Details
				* `command` - Optional Command to execute the content.
				* `content` - Content Source Details.
					* `bucket` - Bucket Name.
					* `checksum` - SHA256 checksum of the artifact.
					* `namespace` - Namespace.
					* `object` - Object Name.
					* `source_type` - Content Source Details. 
				* `endpoint` - Endpoint to be invoked.
				* `execution_type` - The action type of the task
				* `variables` - The variable of the task.Atleast one of dynamicArguments or output needs to be provided.
					* `input_variables` - The input variables for the task.
						* `description` - The description of the argument.
						* `name` - The name of the argument
						* `type` - Input argument Type. 
					* `output_variables` - The list of output variables.
			* `is_copy_to_library_enabled` - Make a copy of this task in Library
			* `name` - The name of the task
			* `os_type` - The OS for the task.
			* `platform` - The platform of the runbook.
			* `properties` - The properties of the task.
				* `num_retries` - The number of retries allowed.
				* `timeout_in_seconds` - The timeout in seconds for the task.
			* `scope` - The scope of the task.
			* `task_record_id` - The ID of taskRecord.
	* `version` - The version of the runbook.
* `compartment_id` - 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A user-friendly description. To provide some insight about the resource. Avoid entering confidential information. 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `estimated_time` - Estimated time to successfully complete the runbook execution
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the resource.
* `is_default` - Is the runbook default?
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `operation` - The lifecycle operation performed by the task.
* `os_type` - The OS type for the runbook.
* `platform` - The platform of the runbook.
* `resource_region` - Associated region
* `runbook_relevance` - Type of runbook structure.
* `state` - The current state of the Runbook.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.
* `type` - The type of the runbook.

