---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_runbook"
sidebar_current: "docs-oci-resource-fleet_apps_management-runbook"
description: |-
  Provides the Runbook resource in Oracle Cloud Infrastructure Fleet Apps Management service
---

# oci_fleet_apps_management_runbook
This resource provides the Runbook resource in Oracle Cloud Infrastructure Fleet Apps Management service.

Creates a new Runbook.


## Example Usage

```hcl
resource "oci_fleet_apps_management_runbook" "test_runbook" {
	#Required
	associations {
		#Required
		execution_workflow_details {
			#Required
			workflow {
				#Required
				group_name = oci_identity_group.test_group.name
				steps {
					#Required
					type = var.runbook_associations_execution_workflow_details_workflow_steps_type

					#Optional
					group_name = oci_identity_group.test_group.name
					step_name = var.runbook_associations_execution_workflow_details_workflow_steps_step_name
					steps = var.runbook_associations_execution_workflow_details_workflow_steps_steps
				}
				type = var.runbook_associations_execution_workflow_details_workflow_type
			}
		}
		groups {
			#Required
			name = var.runbook_associations_groups_name
			type = var.runbook_associations_groups_type

			#Optional
			properties {
				#Required
				action_on_failure = var.runbook_associations_groups_properties_action_on_failure

				#Optional
				condition = var.runbook_associations_groups_properties_condition
				notification_preferences {

					#Optional
					should_notify_on_pause = var.runbook_associations_groups_properties_notification_preferences_should_notify_on_pause
					should_notify_on_task_failure = var.runbook_associations_groups_properties_notification_preferences_should_notify_on_task_failure
					should_notify_on_task_success = var.runbook_associations_groups_properties_notification_preferences_should_notify_on_task_success
				}
				pause_details {
					#Required
					kind = var.runbook_associations_groups_properties_pause_details_kind

					#Optional
					duration_in_minutes = var.runbook_associations_groups_properties_pause_details_duration_in_minutes
				}
				run_on = var.runbook_associations_groups_properties_run_on
			}
		}
		tasks {
			#Required
			association_type = var.runbook_associations_tasks_association_type
			step_name = var.runbook_associations_tasks_step_name
			task_record_details {
				#Required
				scope = var.runbook_associations_tasks_task_record_details_scope

				#Optional
				description = var.runbook_associations_tasks_task_record_details_description
				execution_details {
					#Required
					execution_type = var.runbook_associations_tasks_task_record_details_execution_details_execution_type

					#Optional
					command = var.runbook_associations_tasks_task_record_details_execution_details_command
					content {
						#Required
						bucket = var.runbook_associations_tasks_task_record_details_execution_details_content_bucket
						checksum = var.runbook_associations_tasks_task_record_details_execution_details_content_checksum
						namespace = var.runbook_associations_tasks_task_record_details_execution_details_content_namespace
						object = var.runbook_associations_tasks_task_record_details_execution_details_content_object
						source_type = var.runbook_associations_tasks_task_record_details_execution_details_content_source_type
					}
					credentials {

						#Optional
						display_name = var.runbook_associations_tasks_task_record_details_execution_details_credentials_display_name
						id = var.runbook_associations_tasks_task_record_details_execution_details_credentials_id
					}
					endpoint = var.runbook_associations_tasks_task_record_details_execution_details_endpoint
					variables {

						#Optional
						input_variables {

							#Optional
							description = var.runbook_associations_tasks_task_record_details_execution_details_variables_input_variables_description
							name = var.runbook_associations_tasks_task_record_details_execution_details_variables_input_variables_name
							type = var.runbook_associations_tasks_task_record_details_execution_details_variables_input_variables_type
						}
						output_variables = var.runbook_associations_tasks_task_record_details_execution_details_variables_output_variables
					}
				}
				is_apply_subject_task = var.runbook_associations_tasks_task_record_details_is_apply_subject_task
				is_copy_to_library_enabled = var.runbook_associations_tasks_task_record_details_is_copy_to_library_enabled
				is_discovery_output_task = var.runbook_associations_tasks_task_record_details_is_discovery_output_task
				name = var.runbook_associations_tasks_task_record_details_name
				os_type = var.runbook_associations_tasks_task_record_details_os_type
				platform = var.runbook_associations_tasks_task_record_details_platform
				properties {

					#Optional
					num_retries = var.runbook_associations_tasks_task_record_details_properties_num_retries
					timeout_in_seconds = var.runbook_associations_tasks_task_record_details_properties_timeout_in_seconds
				}
				task_record_id = oci_fleet_apps_management_task_record.test_task_record.id
			}

			#Optional
			output_variable_mappings {
				#Required
				name = var.runbook_associations_tasks_output_variable_mappings_name
				output_variable_details {
					#Required
					output_variable_name = var.runbook_associations_tasks_output_variable_mappings_output_variable_details_output_variable_name
					step_name = var.runbook_associations_tasks_output_variable_mappings_output_variable_details_step_name
				}
			}
			step_properties {
				#Required
				action_on_failure = var.runbook_associations_tasks_step_properties_action_on_failure

				#Optional
				condition = var.runbook_associations_tasks_step_properties_condition
				notification_preferences {

					#Optional
					should_notify_on_pause = var.runbook_associations_tasks_step_properties_notification_preferences_should_notify_on_pause
					should_notify_on_task_failure = var.runbook_associations_tasks_step_properties_notification_preferences_should_notify_on_task_failure
					should_notify_on_task_success = var.runbook_associations_tasks_step_properties_notification_preferences_should_notify_on_task_success
				}
				pause_details {
					#Required
					kind = var.runbook_associations_tasks_step_properties_pause_details_kind

					#Optional
					duration_in_minutes = var.runbook_associations_tasks_step_properties_pause_details_duration_in_minutes
				}
				run_on = var.runbook_associations_tasks_step_properties_run_on
			}
		}

		#Optional
		rollback_workflow_details {
			#Required
			scope = var.runbook_associations_rollback_workflow_details_scope
			workflow {
				#Required
				group_name = oci_identity_group.test_group.name
				steps {
					#Required
					type = var.runbook_associations_rollback_workflow_details_workflow_steps_type

					#Optional
					group_name = oci_identity_group.test_group.name
					step_name = var.runbook_associations_rollback_workflow_details_workflow_steps_step_name
					steps = var.runbook_associations_rollback_workflow_details_workflow_steps_steps
				}
				type = var.runbook_associations_rollback_workflow_details_workflow_type
			}
		}
		version = var.runbook_associations_version
	}
	compartment_id = var.compartment_id
	operation = var.runbook_operation
	os_type = var.runbook_os_type
	runbook_relevance = var.runbook_runbook_relevance

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.runbook_description
	display_name = var.runbook_display_name
	estimated_time = var.runbook_estimated_time
	freeform_tags = {"bar-key"= "value"}
	is_default = var.runbook_is_default
	platform = var.runbook_platform
}
```

## Argument Reference

The following arguments are supported:

* `associations` - (Required) (Updatable) Associations for the runbook.
	* `execution_workflow_details` - (Required) (Updatable) Execution Workflow details.
		* `workflow` - (Required) (Updatable) Execution Workflow for the runbook.
			* `group_name` - (Required) (Updatable) Name of the group.
			* `steps` - (Required) (Updatable) Steps within the Group.
				* `group_name` - (Required when type=PARALLEL_TASK_GROUP) (Updatable) Name of the group.
				* `step_name` - (Required when type=TASK) (Updatable) Provide StepName for the Task.
				* `steps` - (Applicable when type=PARALLEL_TASK_GROUP) (Updatable) Tasks within the Group. Provide the stepName for all applicable tasks. 
				* `type` - (Required) (Updatable) Content Source Details. 
			* `type` - (Required) (Updatable) Workflow Group  Details. 
	* `groups` - (Required) (Updatable) The groups of the runbook. 
		* `name` - (Required) (Updatable) The name of the group.
		* `properties` - (Optional) (Updatable) The properties of the component.
			* `action_on_failure` - (Required) (Updatable) The action to be taken in case of a failure.
			* `condition` - (Optional) (Updatable) Build control flow conditions that determine the relevance of the task execution. 
			* `notification_preferences` - (Optional) (Updatable) Preferences to send notifications on the task activities.
				* `should_notify_on_pause` - (Optional) (Updatable) Enables notification on pause.
				* `should_notify_on_task_failure` - (Optional) (Updatable) Enables or disables notification on Task Failures.
				* `should_notify_on_task_success` - (Optional) (Updatable) Enables or disables notification on Task Success.
			* `pause_details` - (Optional) (Updatable) Pause Details
				* `duration_in_minutes` - (Required when kind=TIME_BASED) (Updatable) Time in minutes to apply Pause.
				* `kind` - (Required) (Updatable) Pause based On. 
			* `run_on` - (Optional) (Updatable) The runOn condition for the task/group/container. Build task execution conditions if applicable to product and product-specific components. This condition is relevant when handling product stack workflows. Example: target.product.name = Oracle WebLogic Server OR target.product.name = Oracle HTTP Server 
		* `type` - (Required) (Updatable) The type of the group. PARALLEL_TASK_GROUP : Helps to execute tasks parallelly inside a resource. PARALLEL_RESOURCE_GROUP : Executes tasks across resources parallelly. ROLLING_RESOURCE_GROUP : Executes tasks across resources in a rolling order. 
	* `rollback_workflow_details` - (Optional) (Updatable) Rollback Workflow details.
		* `scope` - (Required) (Updatable) rollback Scope 
		* `workflow` - (Required) (Updatable) Rollback Workflow for the runbook.
			* `group_name` - (Required) (Updatable) Name of the group.
			* `steps` - (Required) (Updatable) Steps within the Group.
				* `group_name` - (Required when type=PARALLEL_TASK_GROUP) (Updatable) Name of the group.
				* `step_name` - (Required when type=TASK) (Updatable) Provide StepName for the Task.
				* `steps` - (Applicable when type=PARALLEL_TASK_GROUP) (Updatable) Tasks within the Group. Provide the stepName for all applicable tasks. 
				* `type` - (Required) (Updatable) Content Source Details. 
			* `type` - (Required) (Updatable) Workflow Group  Details. 
	* `tasks` - (Required) (Updatable) A set of tasks to execute in the runbook.
		* `association_type` - (Required) (Updatable) The association type of the task
		* `output_variable_mappings` - (Optional) (Updatable) Mapping output variables of previous tasks to the input variables of the current task.
			* `name` - (Required) (Updatable) The name of the input variable.
			* `output_variable_details` - (Required) (Updatable) The details of the output variable that will be used for mapping.
				* `output_variable_name` - (Required) (Updatable) The name of the output variable whose value has to be mapped.
				* `step_name` - (Required) (Updatable) The name of the task step the output variable belongs to.
		* `step_name` - (Required) (Updatable) The name of the task step.
		* `step_properties` - (Optional) (Updatable) The properties of the component.
			* `action_on_failure` - (Required) (Updatable) The action to be taken in case of a failure.
			* `condition` - (Optional) (Updatable) Build control flow conditions that determine the relevance of the task execution. 
			* `notification_preferences` - (Optional) (Updatable) Preferences to send notifications on the task activities.
				* `should_notify_on_pause` - (Optional) (Updatable) Enables notification on pause.
				* `should_notify_on_task_failure` - (Optional) (Updatable) Enables or disables notification on Task Failures.
				* `should_notify_on_task_success` - (Optional) (Updatable) Enables or disables notification on Task Success.
			* `pause_details` - (Optional) (Updatable) Pause Details
				* `duration_in_minutes` - (Required when kind=TIME_BASED) (Updatable) Time in minutes to apply Pause.
				* `kind` - (Required) (Updatable) Pause based On. 
			* `run_on` - (Optional) (Updatable) The runOn condition for the task/group/container. Build task execution conditions if applicable to product and product-specific components. This condition is relevant when handling product stack workflows. Example: target.product.name = Oracle WebLogic Server OR target.product.name = Oracle HTTP Server 
		* `task_record_details` - (Required) (Updatable) The details of the task.
			* `description` - (Applicable when scope=LOCAL) (Updatable) The description of the task.
			* `execution_details` - (Required when scope=LOCAL) (Updatable) Execution details.
				* `command` - (Applicable when execution_type=SCRIPT) (Updatable) Optional command to execute the content. You can provide any commands/arguments that can't be part of the script. 
				* `content` - (Applicable when execution_type=SCRIPT) (Updatable) Content Source details.
					* `bucket` - (Required) (Updatable) Bucket Name.
					* `checksum` - (Required) (Updatable) md5 checksum of the artifact.
					* `namespace` - (Required) (Updatable) Namespace.
					* `object` - (Required) (Updatable) Object Name.
					* `source_type` - (Required) (Updatable) Content Source type details. 
				* `credentials` - (Applicable when execution_type=SCRIPT) (Updatable) Credentials required for executing the task. 
					* `display_name` - (Applicable when execution_type=SCRIPT) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
					* `id` - (Required when execution_type=SCRIPT) (Updatable) The OCID of the resource.
				* `endpoint` - (Required when execution_type=API) (Updatable) Endpoint to be invoked.
				* `execution_type` - (Required) (Updatable) The action type of the task
				* `variables` - (Applicable when execution_type=SCRIPT) (Updatable) The variable of the task. At least one of the dynamicArguments or output needs to be provided. 
					* `input_variables` - (Applicable when execution_type=SCRIPT) (Updatable) The input variables for the task.
						* `description` - (Applicable when execution_type=SCRIPT) (Updatable) The description of the argument.
						* `name` - (Required when execution_type=SCRIPT) (Updatable) The name of the argument.
						* `type` - (Required when execution_type=SCRIPT) (Updatable) Input argument Type. 
					* `output_variables` - (Applicable when execution_type=SCRIPT) (Updatable) The list of output variables.
			* `is_apply_subject_task` - (Applicable when scope=LOCAL) (Updatable) Is this an Apply Subject Task? Ex. Patch Execution Task
			* `is_copy_to_library_enabled` - (Applicable when scope=LOCAL) (Updatable) Make a copy of this task in Library
			* `is_discovery_output_task` - (Applicable when scope=LOCAL) (Updatable) Is this a discovery output task?
			* `name` - (Applicable when scope=LOCAL) (Updatable) The name of the task
			* `os_type` - (Required when scope=LOCAL) (Updatable) The OS for the task.
			* `platform` - (Applicable when scope=LOCAL) (Updatable) The platform of the runbook.
			* `properties` - (Applicable when scope=LOCAL) (Updatable) The properties of the task.
				* `num_retries` - (Required when scope=LOCAL) (Updatable) The number of retries allowed.
				* `timeout_in_seconds` - (Required when scope=LOCAL) (Updatable) The timeout in seconds for the task.
			* `scope` - (Required) (Updatable) The scope of the task.
			* `task_record_id` - (Required when scope=SHARED) (Updatable) The ID of taskRecord.
	* `version` - (Optional) (Updatable) The version of the runbook.
* `compartment_id` - (Required) 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) A user-friendly description. To provide some insight about the resource. Avoid entering confidential information. 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `estimated_time` - (Optional) (Updatable) Estimated time to successfully complete the runbook execution
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `is_default` - (Optional) (Updatable) Is the runbook default?
* `operation` - (Required) (Updatable) The lifecycle operation performed by the task.
* `os_type` - (Required) (Updatable) The OS type for the runbook.
* `platform` - (Optional) (Updatable) The platform of the runbook.
* `runbook_relevance` - (Required) (Updatable) Type of runbook structure.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `associations` - Associations for the runbook.
	* `execution_workflow_details` - Execution Workflow details.
		* `workflow` - Execution Workflow for the runbook.
			* `group_name` - Name of the group.
			* `steps` - Steps within the Group.
				* `group_name` - Name of the group.
				* `step_name` - Provide StepName for the Task.
				* `steps` - Tasks within the Group. Provide the stepName for all applicable tasks. 
				* `type` - Content Source Details. 
			* `type` - Workflow Group  Details. 
	* `groups` - The groups of the runbook. 
		* `name` - The name of the group.
		* `properties` - The properties of the component.
			* `action_on_failure` - The action to be taken in case of a failure.
			* `condition` - Build control flow conditions that determine the relevance of the task execution. 
			* `notification_preferences` - Preferences to send notifications on the task activities.
				* `should_notify_on_pause` - Enables notification on pause.
				* `should_notify_on_task_failure` - Enables or disables notification on Task Failures.
				* `should_notify_on_task_success` - Enables or disables notification on Task Success.
			* `pause_details` - Pause Details
				* `duration_in_minutes` - Time in minutes to apply Pause.
				* `kind` - Pause based On. 
			* `run_on` - The runOn condition for the task/group/container. Build task execution conditions if applicable to product and product-specific components. This condition is relevant when handling product stack workflows. Example: target.product.name = Oracle WebLogic Server OR target.product.name = Oracle HTTP Server 
		* `type` - The type of the group. PARALLEL_TASK_GROUP : Helps to execute tasks parallelly inside a resource. PARALLEL_RESOURCE_GROUP : Executes tasks across resources parallelly. ROLLING_RESOURCE_GROUP : Executes tasks across resources in a rolling order. 
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
	* `tasks` - A set of tasks to execute in the runbook.
		* `association_type` - The association type of the task
		* `output_variable_mappings` - Mapping output variables of previous tasks to the input variables of the current task.
			* `name` - The name of the input variable.
			* `output_variable_details` - The details of the output variable that will be used for mapping.
				* `output_variable_name` - The name of the output variable whose value has to be mapped.
				* `step_name` - The name of the task step the output variable belongs to.
		* `step_name` - The name of the task step.
		* `step_properties` - The properties of the component.
			* `action_on_failure` - The action to be taken in case of a failure.
			* `condition` - Build control flow conditions that determine the relevance of the task execution. 
			* `notification_preferences` - Preferences to send notifications on the task activities.
				* `should_notify_on_pause` - Enables notification on pause.
				* `should_notify_on_task_failure` - Enables or disables notification on Task Failures.
				* `should_notify_on_task_success` - Enables or disables notification on Task Success.
			* `pause_details` - Pause Details
				* `duration_in_minutes` - Time in minutes to apply Pause.
				* `kind` - Pause based On. 
			* `run_on` - The runOn condition for the task/group/container. Build task execution conditions if applicable to product and product-specific components. This condition is relevant when handling product stack workflows. Example: target.product.name = Oracle WebLogic Server OR target.product.name = Oracle HTTP Server 
		* `task_record_details` - The details of the task.
			* `description` - The description of the task.
			* `execution_details` - Execution details.
				* `command` - Optional command to execute the content. You can provide any commands/arguments that can't be part of the script. 
				* `content` - Content Source details.
					* `bucket` - Bucket Name.
					* `checksum` - md5 checksum of the artifact.
					* `namespace` - Namespace.
					* `object` - Object Name.
					* `source_type` - Content Source type details. 
				* `credentials` - Credentials required for executing the task. 
					* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
					* `id` - The OCID of the resource.
				* `endpoint` - Endpoint to be invoked.
				* `execution_type` - The action type of the task
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
	* `version` - The version of the runbook.
* `compartment_id` - 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A user-friendly description. To provide some insight about the resource. Avoid entering confidential information. 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `estimated_time` - Estimated time to successfully complete the runbook execution.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the resource.
* `is_default` - Is the runbook default? Sets this runbook as the default for the chosen product/product stack for the specified lifecycle operation. 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `operation` - The lifecycle operation performed by the runbook.
* `os_type` - The OS type for the runbook.
* `platform` - The platform of the runbook.
* `resource_region` - Associated region
* `runbook_relevance` - Relevance of the runbook. 
* `state` - The current state of the Runbook.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.
* `type` - The type of the runbook.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Runbook
	* `update` - (Defaults to 20 minutes), when updating the Runbook
	* `delete` - (Defaults to 20 minutes), when destroying the Runbook


## Import

Runbooks can be imported using the `id`, e.g.

```
$ terraform import oci_fleet_apps_management_runbook.test_runbook "id"
```

