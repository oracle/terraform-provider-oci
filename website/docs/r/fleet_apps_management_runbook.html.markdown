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
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/fleet-management/latest/Runbook

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/fleet_apps_management

Creates a runbook.


## Example Usage

```hcl
resource "oci_fleet_apps_management_runbook" "test_runbook" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.runbook_display_name
	operation = var.runbook_operation
	runbook_version {
		#Required
		execution_workflow_details {
			#Required
			workflow {
				#Required
				group_name = oci_identity_group.test_group.name
				steps {
					#Required
					type = var.runbook_runbook_version_execution_workflow_details_workflow_steps_type

					#Optional
					group_name = oci_identity_group.test_group.name
					step_name = var.runbook_runbook_version_execution_workflow_details_workflow_steps_step_name
					steps = var.runbook_runbook_version_execution_workflow_details_workflow_steps_steps
				}
				type = var.runbook_runbook_version_execution_workflow_details_workflow_type
			}
		}
		groups {
			#Required
			name = var.runbook_runbook_version_groups_name
			type = var.runbook_runbook_version_groups_type

			#Optional
			properties {
				#Required
				action_on_failure = var.runbook_runbook_version_groups_properties_action_on_failure

				#Optional
				notification_preferences {

					#Optional
					should_notify_on_pause = var.runbook_runbook_version_groups_properties_notification_preferences_should_notify_on_pause
					should_notify_on_task_failure = var.runbook_runbook_version_groups_properties_notification_preferences_should_notify_on_task_failure
					should_notify_on_task_success = var.runbook_runbook_version_groups_properties_notification_preferences_should_notify_on_task_success
				}
				pause_details {
					#Required
					kind = var.runbook_runbook_version_groups_properties_pause_details_kind

					#Optional
					duration_in_minutes = var.runbook_runbook_version_groups_properties_pause_details_duration_in_minutes
				}
				pre_condition = var.runbook_runbook_version_groups_properties_pre_condition
				run_on {
					#Required
					kind = var.runbook_runbook_version_groups_properties_run_on_kind

					#Optional
					condition = var.runbook_runbook_version_groups_properties_run_on_condition
					host = var.runbook_runbook_version_groups_properties_run_on_host
					previous_task_instance_details {

						#Optional
						output_variable_details {

							#Optional
							output_variable_name = var.runbook_runbook_version_groups_properties_run_on_previous_task_instance_details_output_variable_details_output_variable_name
							step_name = var.runbook_runbook_version_groups_properties_run_on_previous_task_instance_details_output_variable_details_step_name
						}
						resource_id = oci_cloud_guard_resource.test_resource.id
						resource_type = var.runbook_runbook_version_groups_properties_run_on_previous_task_instance_details_resource_type
					}
				}
			}
		}
		tasks {
			#Required
			step_name = var.runbook_runbook_version_tasks_step_name
			task_record_details {
				#Required
				scope = var.runbook_runbook_version_tasks_task_record_details_scope

				#Optional
				description = var.runbook_runbook_version_tasks_task_record_details_description
				execution_details {
					#Required
					execution_type = var.runbook_runbook_version_tasks_task_record_details_execution_details_execution_type

					#Optional
					catalog_id = oci_datacatalog_catalog.test_catalog.id
					command = var.runbook_runbook_version_tasks_task_record_details_execution_details_command
					config_file = var.runbook_runbook_version_tasks_task_record_details_execution_details_config_file
					content {
						#Required
						source_type = var.runbook_runbook_version_tasks_task_record_details_execution_details_content_source_type

						#Optional
						bucket = var.runbook_runbook_version_tasks_task_record_details_execution_details_content_bucket
						catalog_id = oci_datacatalog_catalog.test_catalog.id
						checksum = var.runbook_runbook_version_tasks_task_record_details_execution_details_content_checksum
						namespace = var.runbook_runbook_version_tasks_task_record_details_execution_details_content_namespace
						object = var.runbook_runbook_version_tasks_task_record_details_execution_details_content_object
					}
					credentials {

						#Optional
						display_name = var.runbook_runbook_version_tasks_task_record_details_execution_details_credentials_display_name
						id = var.runbook_runbook_version_tasks_task_record_details_execution_details_credentials_id
					}
					endpoint = var.runbook_runbook_version_tasks_task_record_details_execution_details_endpoint
					is_executable_content = var.runbook_runbook_version_tasks_task_record_details_execution_details_is_executable_content
					is_locked = var.runbook_runbook_version_tasks_task_record_details_execution_details_is_locked
					is_read_output_variable_enabled = var.runbook_runbook_version_tasks_task_record_details_execution_details_is_read_output_variable_enabled
					target_compartment_id = oci_identity_compartment.test_compartment.id
					variables {

						#Optional
						input_variables {

							#Optional
							description = var.runbook_runbook_version_tasks_task_record_details_execution_details_variables_input_variables_description
							name = var.runbook_runbook_version_tasks_task_record_details_execution_details_variables_input_variables_name
							type = var.runbook_runbook_version_tasks_task_record_details_execution_details_variables_input_variables_type
						}
						output_variables = var.runbook_runbook_version_tasks_task_record_details_execution_details_variables_output_variables
					}
				}
				is_apply_subject_task = var.runbook_runbook_version_tasks_task_record_details_is_apply_subject_task
				is_copy_to_library_enabled = var.runbook_runbook_version_tasks_task_record_details_is_copy_to_library_enabled
				is_discovery_output_task = var.runbook_runbook_version_tasks_task_record_details_is_discovery_output_task
				name = var.runbook_runbook_version_tasks_task_record_details_name
				os_type = var.runbook_runbook_version_tasks_task_record_details_os_type
				platform = var.runbook_runbook_version_tasks_task_record_details_platform
				properties {

					#Optional
					num_retries = var.runbook_runbook_version_tasks_task_record_details_properties_num_retries
					timeout_in_seconds = var.runbook_runbook_version_tasks_task_record_details_properties_timeout_in_seconds
				}
				task_record_id = oci_fleet_apps_management_task_record.test_task_record.id
			}

			#Optional
			output_variable_mappings {
				#Required
				name = var.runbook_runbook_version_tasks_output_variable_mappings_name
				output_variable_details {
					#Required
					output_variable_name = var.runbook_runbook_version_tasks_output_variable_mappings_output_variable_details_output_variable_name
					step_name = var.runbook_runbook_version_tasks_output_variable_mappings_output_variable_details_step_name
				}
			}
			step_properties {
				#Required
				action_on_failure = var.runbook_runbook_version_tasks_step_properties_action_on_failure

				#Optional
				notification_preferences {

					#Optional
					should_notify_on_pause = var.runbook_runbook_version_tasks_step_properties_notification_preferences_should_notify_on_pause
					should_notify_on_task_failure = var.runbook_runbook_version_tasks_step_properties_notification_preferences_should_notify_on_task_failure
					should_notify_on_task_success = var.runbook_runbook_version_tasks_step_properties_notification_preferences_should_notify_on_task_success
				}
				pause_details {
					#Required
					kind = var.runbook_runbook_version_tasks_step_properties_pause_details_kind

					#Optional
					duration_in_minutes = var.runbook_runbook_version_tasks_step_properties_pause_details_duration_in_minutes
				}
				pre_condition = var.runbook_runbook_version_tasks_step_properties_pre_condition
				run_on {
					#Required
					kind = var.runbook_runbook_version_tasks_step_properties_run_on_kind

					#Optional
					condition = var.runbook_runbook_version_tasks_step_properties_run_on_condition
					host = var.runbook_runbook_version_tasks_step_properties_run_on_host
					previous_task_instance_details {

						#Optional
						output_variable_details {

							#Optional
							output_variable_name = var.runbook_runbook_version_tasks_step_properties_run_on_previous_task_instance_details_output_variable_details_output_variable_name
							step_name = var.runbook_runbook_version_tasks_step_properties_run_on_previous_task_instance_details_output_variable_details_step_name
						}
						resource_id = oci_cloud_guard_resource.test_resource.id
						resource_type = var.runbook_runbook_version_tasks_step_properties_run_on_previous_task_instance_details_resource_type
					}
				}
			}
		}

		#Optional
		is_latest = var.runbook_runbook_version_is_latest
		rollback_workflow_details {
			#Required
			scope = var.runbook_runbook_version_rollback_workflow_details_scope
			workflow {
				#Required
				group_name = oci_identity_group.test_group.name
				steps {
					#Required
					type = var.runbook_runbook_version_rollback_workflow_details_workflow_steps_type

					#Optional
					group_name = oci_identity_group.test_group.name
					step_name = var.runbook_runbook_version_rollback_workflow_details_workflow_steps_step_name
					steps = var.runbook_runbook_version_rollback_workflow_details_workflow_steps_steps
				}
				type = var.runbook_runbook_version_rollback_workflow_details_workflow_type
			}
		}
		version = var.runbook_runbook_version_version
	}

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.runbook_description
	estimated_time = var.runbook_estimated_time
	freeform_tags = {"bar-key"= "value"}
	is_default = var.runbook_is_default
	is_sudo_access_needed = var.runbook_is_sudo_access_needed
	os_type = var.runbook_os_type
	platform = var.runbook_platform
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) A user-friendly description. To provide some insight about the resource. Avoid entering confidential information. 
* `display_name` - (Required) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `estimated_time` - (Optional) (Updatable) Estimated time to successfully complete the runbook execution.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `is_default` - (Optional) (Updatable) Is the runbook default?
* `is_sudo_access_needed` - (Optional) (Updatable) Does this runbook need SUDO access to execute?
* `operation` - (Required) (Updatable) The lifecycle operation performed by the task.
* `os_type` - (Optional) (Updatable) The OS type for the runbook.
* `platform` - (Optional) (Updatable) The platform of the runbook.
* `runbook_version` - (Required) Version for the runbook.
	* `execution_workflow_details` - (Required) Execution Workflow details.
		* `workflow` - (Required) Execution Workflow for the runbook.
			* `group_name` - (Required) Name of the group.
			* `steps` - (Required) Steps within the Group.
				* `group_name` - (Required when type=PARALLEL_TASK_GROUP) Name of the group.
				* `step_name` - (Required when type=TASK) Provide StepName for the Task.
				* `steps` - (Applicable when type=PARALLEL_TASK_GROUP) Tasks within the Group. Provide the stepName for all applicable tasks. 
				* `type` - (Required) Content Source Details. 
			* `type` - (Required) Workflow Group  Details. 
	* `groups` - (Required) The groups of the runbook. 
		* `name` - (Required) The name of the group.
		* `properties` - (Optional) The properties of the component.
			* `action_on_failure` - (Required) The action to be taken in case of a failure.
			* `notification_preferences` - (Optional) Preferences to send notifications on the task activities.
				* `should_notify_on_pause` - (Optional) Enables notification on pause.
				* `should_notify_on_task_failure` - (Optional) Enables or disables notification on Task Failures.
				* `should_notify_on_task_success` - (Optional) Enables or disables notification on Task Success.
			* `pause_details` - (Optional) Pause Details
				* `duration_in_minutes` - (Required when kind=TIME_BASED) Time in minutes to apply Pause.
				* `kind` - (Required) Pause based On. 
			* `pre_condition` - (Optional) Build control flow conditions that determine the relevance of the task execution. 
			* `run_on` - (Optional) The runon conditions
				* `condition` - (Required when kind=SCHEDULE_INSTANCES) Build control flow conditions that determine the relevance of the task execution. 
				* `host` - (Required when kind=SELF_HOSTED_INSTANCES) OCID of the self hosted instance.
				* `kind` - (Required) Run on based On. 
				* `previous_task_instance_details` - (Required when kind=PREVIOUS_TASK_INSTANCES) Previous Task Instance Details 
					* `output_variable_details` - (Required when kind=PREVIOUS_TASK_INSTANCES) The details of the output variable that will be used for mapping.
						* `output_variable_name` - (Required when kind=PREVIOUS_TASK_INSTANCES) The name of the output variable whose value has to be mapped.
						* `step_name` - (Required when kind=PREVIOUS_TASK_INSTANCES) The name of the task step the output variable belongs to.
					* `resource_id` - (Required when kind=PREVIOUS_TASK_INSTANCES) Resource Ocid.
					* `resource_type` - (Applicable when kind=PREVIOUS_TASK_INSTANCES) Resource Type.
		* `type` - (Required) The type of the group. PARALLEL_TASK_GROUP : Helps to execute tasks parallelly inside a resource. PARALLEL_RESOURCE_GROUP : Executes tasks across resources parallelly. ROLLING_RESOURCE_GROUP : Executes tasks across resources in a rolling order. 
	* `is_latest` - (Optional) Is this version the latest? 
	* `rollback_workflow_details` - (Optional) Rollback Workflow details.
		* `scope` - (Required) rollback Scope 
		* `workflow` - (Required) Rollback Workflow for the runbook.
			* `group_name` - (Required) Name of the group.
			* `steps` - (Required) Steps within the Group.
				* `group_name` - (Required when type=PARALLEL_TASK_GROUP) Name of the group.
				* `step_name` - (Required when type=TASK) Provide StepName for the Task.
				* `steps` - (Applicable when type=PARALLEL_TASK_GROUP) Tasks within the Group. Provide the stepName for all applicable tasks. 
				* `type` - (Required) Content Source Details. 
			* `type` - (Required) Workflow Group  Details. 
	* `tasks` - (Required) A set of tasks to execute in the runbook.
		* `output_variable_mappings` - (Optional) Mapping output variables of previous tasks to the input variables of the current task.
			* `name` - (Required) The name of the input variable.
			* `output_variable_details` - (Required) The details of the output variable that will be used for mapping.
				* `output_variable_name` - (Required) The name of the output variable whose value has to be mapped.
				* `step_name` - (Required) The name of the task step the output variable belongs to.
		* `step_name` - (Required) The name of the task step.
		* `step_properties` - (Optional) The properties of the component.
			* `action_on_failure` - (Required) The action to be taken in case of a failure.
			* `notification_preferences` - (Optional) Preferences to send notifications on the task activities.
				* `should_notify_on_pause` - (Optional) Enables notification on pause.
				* `should_notify_on_task_failure` - (Optional) Enables or disables notification on Task Failures.
				* `should_notify_on_task_success` - (Optional) Enables or disables notification on Task Success.
			* `pause_details` - (Optional) Pause Details
				* `duration_in_minutes` - (Required when kind=TIME_BASED) Time in minutes to apply Pause.
				* `kind` - (Required) Pause based On. 
			* `pre_condition` - (Optional) Build control flow conditions that determine the relevance of the task execution. 
			* `run_on` - (Optional) The runon conditions
				* `condition` - (Required when kind=SCHEDULE_INSTANCES) Build control flow conditions that determine the relevance of the task execution. 
				* `host` - (Required when kind=SELF_HOSTED_INSTANCES) OCID of the self hosted instance.
				* `kind` - (Required) Run on based On. 
				* `previous_task_instance_details` - (Required when kind=PREVIOUS_TASK_INSTANCES) Previous Task Instance Details 
					* `output_variable_details` - (Required when kind=PREVIOUS_TASK_INSTANCES) The details of the output variable that will be used for mapping.
						* `output_variable_name` - (Required when kind=PREVIOUS_TASK_INSTANCES) The name of the output variable whose value has to be mapped.
						* `step_name` - (Required when kind=PREVIOUS_TASK_INSTANCES) The name of the task step the output variable belongs to.
					* `resource_id` - (Required when kind=PREVIOUS_TASK_INSTANCES) Resource Ocid.
					* `resource_type` - (Applicable when kind=PREVIOUS_TASK_INSTANCES) Resource Type.
		* `task_record_details` - (Required) The details of the task.
			* `description` - (Applicable when scope=LOCAL) The description of the task.
			* `execution_details` - (Required when scope=LOCAL) Execution details.
				* `catalog_id` - (Required when execution_type=TERRAFORM) Catalog Id having terraform package.
				* `command` - (Applicable when execution_type=SCRIPT) Optional command to execute the content. You can provide any commands/arguments that can't be part of the script. 
				* `config_file` - (Applicable when execution_type=TERRAFORM) Catalog Id having config file.
				* `content` - (Applicable when execution_type=SCRIPT) Content Source details.
					* `bucket` - (Required when source_type=OBJECT_STORAGE_BUCKET) Bucket Name.
					* `catalog_id` - (Required when source_type=CATALOG) Catalog Id having terraform package.
					* `checksum` - (Required when source_type=OBJECT_STORAGE_BUCKET) md5 checksum of the artifact.
					* `namespace` - (Required when source_type=OBJECT_STORAGE_BUCKET) Namespace.
					* `object` - (Required when source_type=OBJECT_STORAGE_BUCKET) Object Name.
					* `source_type` - (Required) Content Source type details. 
				* `credentials` - (Applicable when execution_type=SCRIPT) Credentials required for executing the task. 
					* `display_name` - (Applicable when execution_type=SCRIPT) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
					* `id` - (Required when execution_type=SCRIPT) The OCID of the resource.
				* `endpoint` - (Required when execution_type=API) Endpoint to be invoked.
				* `execution_type` - (Required) The action type of the task
				* `is_executable_content` - (Applicable when execution_type=SCRIPT) Is the Content an executable file?
				* `is_locked` - (Applicable when execution_type=SCRIPT) Is the script locked to prevent changes directly in Object Storage?
				* `is_read_output_variable_enabled` - (Applicable when execution_type=TERRAFORM) Is read output variable enabled
				* `target_compartment_id` - (Required when execution_type=TERRAFORM) OCID of the compartment to which the resource belongs to.
				* `variables` - (Applicable when execution_type=SCRIPT) The variable of the task. At least one of the dynamicArguments or output needs to be provided. 
					* `input_variables` - (Applicable when execution_type=SCRIPT) The input variables for the task.
						* `description` - (Applicable when execution_type=SCRIPT) The description of the argument.
						* `name` - (Required when execution_type=SCRIPT) The name of the argument.
						* `type` - (Required when execution_type=SCRIPT) Input argument Type. 
					* `output_variables` - (Applicable when execution_type=SCRIPT) The list of output variables.
			* `is_apply_subject_task` - (Applicable when scope=LOCAL) Is this an Apply Subject Task? Ex. Patch Execution Task
			* `is_copy_to_library_enabled` - (Applicable when scope=LOCAL) Make a copy of this task in Library
			* `is_discovery_output_task` - (Applicable when scope=LOCAL) Is this a discovery output task?
			* `name` - (Applicable when scope=LOCAL) The name of the task
			* `os_type` - (Applicable when scope=LOCAL) The OS for the task.
			* `platform` - (Applicable when scope=LOCAL) The platform of the runbook.
			* `properties` - (Applicable when scope=LOCAL) The properties of the task.
				* `num_retries` - (Required when scope=LOCAL) The number of retries allowed.
				* `timeout_in_seconds` - (Required when scope=LOCAL) The timeout in seconds for the task.
			* `scope` - (Required) The scope of the task.
			* `task_record_id` - (Required when scope=SHARED) The ID of taskRecord.
	* `version` - (Optional) The version of the runbook.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A user-friendly description. To provide some insight about the resource. Avoid entering confidential information. 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `estimated_time` - Estimated time to successfully complete the runbook execution.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `has_draft_version` - Does this runbook has draft versions?
* `id` - The OCID of the resource.
* `is_default` - Is the runbook default? Sets this runbook as the default for the chosen product/product stack for the specified lifecycle operation. 
* `is_sudo_access_needed` - Does this runbook need SUDO access to execute?
* `latest_version` - Latest runbook version
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `operation` - The lifecycle operation performed by the runbook.
* `os_type` - The OS type for the runbook.
* `platform` - The platform of the runbook.
* `resource_region` - Associated region
* `runbook_version` - Version for the runbook.
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
			* `notification_preferences` - Preferences to send notifications on the task activities.
				* `should_notify_on_pause` - Enables notification on pause.
				* `should_notify_on_task_failure` - Enables or disables notification on Task Failures.
				* `should_notify_on_task_success` - Enables or disables notification on Task Success.
			* `pause_details` - Pause Details
				* `duration_in_minutes` - Time in minutes to apply Pause.
				* `kind` - Pause based On. 
			* `pre_condition` - Build control flow conditions that determine the relevance of the task execution. 
			* `run_on` - The runon conditions
				* `condition` - Build control flow conditions that determine the relevance of the task execution. 
				* `host` - OCID of the self hosted instance.
				* `kind` - Run on based On. 
				* `previous_task_instance_details` - Previous Task Instance Details 
					* `output_variable_details` - The details of the output variable that will be used for mapping.
						* `output_variable_name` - The name of the output variable whose value has to be mapped.
						* `step_name` - The name of the task step the output variable belongs to.
					* `resource_id` - Resource Ocid.
					* `resource_type` - Resource Type.
		* `type` - The type of the group. PARALLEL_TASK_GROUP : Helps to execute tasks parallelly inside a resource. PARALLEL_RESOURCE_GROUP : Executes tasks across resources parallelly. ROLLING_RESOURCE_GROUP : Executes tasks across resources in a rolling order. 
	* `is_latest` - Is this version the latest? 
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
				* `condition` - Build control flow conditions that determine the relevance of the task execution. 
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
	* `version` - The version of the runbook.
* `state` - The current state of the runbook.
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

