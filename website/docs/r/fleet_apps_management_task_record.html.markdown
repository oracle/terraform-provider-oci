---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_task_record"
sidebar_current: "docs-oci-resource-fleet_apps_management-task_record"
description: |-
  Provides the Task Record resource in Oracle Cloud Infrastructure Fleet Apps Management service
---

# oci_fleet_apps_management_task_record
This resource provides the Task Record resource in Oracle Cloud Infrastructure Fleet Apps Management service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/fleet-management/latest/TaskRecord

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/fleet_apps_management

Creates a new task record.


## Example Usage

```hcl
resource "oci_fleet_apps_management_task_record" "test_task_record" {
	#Required
	compartment_id = var.compartment_id
	details {
		#Required
		execution_details {
			#Required
			execution_type = var.task_record_details_execution_details_execution_type

			#Optional
			catalog_id = oci_datacatalog_catalog.test_catalog.id
			command = var.task_record_details_execution_details_command
			config_file = var.task_record_details_execution_details_config_file
			content {
				#Required
				source_type = var.task_record_details_execution_details_content_source_type

				#Optional
				bucket = var.task_record_details_execution_details_content_bucket
				catalog_id = oci_datacatalog_catalog.test_catalog.id
				checksum = var.task_record_details_execution_details_content_checksum
				namespace = var.task_record_details_execution_details_content_namespace
				object = var.task_record_details_execution_details_content_object
			}
			credentials {

				#Optional
				display_name = var.task_record_details_execution_details_credentials_display_name
				id = var.task_record_details_execution_details_credentials_id
			}
			endpoint = var.task_record_details_execution_details_endpoint
			is_executable_content = var.task_record_details_execution_details_is_executable_content
			is_locked = var.task_record_details_execution_details_is_locked
			is_read_output_variable_enabled = var.task_record_details_execution_details_is_read_output_variable_enabled
			system_variables = var.task_record_details_execution_details_system_variables
			target_compartment_id = oci_identity_compartment.test_compartment.id
			variables {

				#Optional
				input_variables {

					#Optional
					description = var.task_record_details_execution_details_variables_input_variables_description
					name = var.task_record_details_execution_details_variables_input_variables_name
					type = var.task_record_details_execution_details_variables_input_variables_type
				}
				output_variables = var.task_record_details_execution_details_variables_output_variables
			}
		}
		scope = var.task_record_details_scope

		#Optional
		is_apply_subject_task = var.task_record_details_is_apply_subject_task
		is_discovery_output_task = var.task_record_details_is_discovery_output_task
		operation = var.task_record_details_operation
		os_type = var.task_record_details_os_type
		platform = var.task_record_details_platform
		properties {
			#Required
			num_retries = var.task_record_details_properties_num_retries
			timeout_in_seconds = var.task_record_details_properties_timeout_in_seconds
		}
	}
	display_name = var.task_record_display_name

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.task_record_description
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) A user-friendly description. To provide some insight about the resource. Avoid entering confidential information. 
* `details` - (Required) (Updatable) The details of the task.
	* `execution_details` - (Required) (Updatable) Execution details.
		* `catalog_id` - (Required when execution_type=TERRAFORM) (Updatable) Catalog Id having terraform package.
		* `command` - (Applicable when execution_type=SCRIPT) (Updatable) Optional command to execute the content. You can provide any commands/arguments that can't be part of the script. 
		* `config_file` - (Applicable when execution_type=TERRAFORM) (Updatable) Catalog Id having config file.
		* `content` - (Applicable when execution_type=SCRIPT) (Updatable) Content Source details.
			* `bucket` - (Required when source_type=OBJECT_STORAGE_BUCKET) (Updatable) Bucket Name.
			* `catalog_id` - (Required when source_type=CATALOG) (Updatable) Catalog Id having terraform package.
			* `checksum` - (Required when source_type=OBJECT_STORAGE_BUCKET) (Updatable) md5 checksum of the artifact.
			* `namespace` - (Required when source_type=OBJECT_STORAGE_BUCKET) (Updatable) Namespace.
			* `object` - (Required when source_type=OBJECT_STORAGE_BUCKET) (Updatable) Object Name.
			* `source_type` - (Required) (Updatable) Content Source type details. 
		* `credentials` - (Applicable when execution_type=SCRIPT) (Updatable) Credentials required for executing the task. 
			* `display_name` - (Applicable when execution_type=SCRIPT) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
			* `id` - (Required when execution_type=SCRIPT) (Updatable) The OCID of the resource.
		* `endpoint` - (Required when execution_type=API) (Updatable) Endpoint to be invoked.
		* `execution_type` - (Required) (Updatable) The action type of the task
		* `is_executable_content` - (Applicable when execution_type=SCRIPT) (Updatable) Is the Content an executable file?
		* `is_locked` - (Applicable when execution_type=SCRIPT) (Updatable) Is the script locked to prevent changes directly in Object Storage?
		* `is_read_output_variable_enabled` - (Applicable when execution_type=TERRAFORM) (Updatable) Is read output variable enabled
		* `system_variables` - (Optional) (Updatable) The list of system variables.
		* `target_compartment_id` - (Required when execution_type=TERRAFORM) (Updatable) OCID of the compartment to which the resource belongs to.
		* `variables` - (Applicable when execution_type=SCRIPT) (Updatable) The variable of the task. At least one of the dynamicArguments or output needs to be provided. 
			* `input_variables` - (Applicable when execution_type=SCRIPT) (Updatable) The input variables for the task.
				* `description` - (Applicable when execution_type=SCRIPT) (Updatable) The description of the argument.
				* `name` - (Required when execution_type=SCRIPT) (Updatable) The name of the argument.
				* `type` - (Required when execution_type=SCRIPT) (Updatable) Input argument Type. 
			* `output_variables` - (Applicable when execution_type=SCRIPT) (Updatable) The list of output variables.
	* `is_apply_subject_task` - (Optional) (Updatable) Is this an Apply Subject Task?  Set this to true for a Patch Execution Task which applies patches(subjects) on a target. 
	* `is_discovery_output_task` - (Optional) (Updatable) Is this a discovery output task?
	* `operation` - (Optional) (Updatable) The lifecycle operation performed by the runbook.
	* `os_type` - (Optional) (Updatable) The OS for the task
	* `platform` - (Optional) (Updatable) The platform of the runbook.
	* `properties` - (Optional) (Updatable) The properties of the task.
		* `num_retries` - (Required) (Updatable) The number of retries allowed.
		* `timeout_in_seconds` - (Required) (Updatable) The timeout in seconds for the task.
	* `scope` - (Required) (Updatable) The scope of the task
* `display_name` - (Required) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A user-friendly description. To provide some insight about the resource. Avoid entering confidential information. 
* `details` - The details of the task.
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
	* `is_apply_subject_task` - Is this an Apply Subject Task?  Set this to true for a Patch Execution Task which applies patches(subjects) on a target. 
	* `is_discovery_output_task` - Is this a discovery output task?
	* `operation` - The lifecycle operation performed by the runbook.
	* `os_type` - The OS for the task
	* `platform` - The platform of the runbook.
	* `properties` - The properties of the task.
		* `num_retries` - The number of retries allowed.
		* `timeout_in_seconds` - The timeout in seconds for the task.
	* `scope` - The scope of the task
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the resource.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `resource_region` - Associated region
* `state` - The current state of the task record.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.
* `type` - Task type.
* `version` - The version of the task record.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Task Record
	* `update` - (Defaults to 20 minutes), when updating the Task Record
	* `delete` - (Defaults to 20 minutes), when destroying the Task Record


## Import

TaskRecords can be imported using the `id`, e.g.

```
$ terraform import oci_fleet_apps_management_task_record.test_task_record "id"
```

