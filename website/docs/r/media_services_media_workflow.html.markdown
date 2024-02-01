---
subcategory: "Media Services"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_media_services_media_workflow"
sidebar_current: "docs-oci-resource-media_services-media_workflow"
description: |-
  Provides the Media Workflow resource in Oracle Cloud Infrastructure Media Services service
---

# oci_media_services_media_workflow
This resource provides the Media Workflow resource in Oracle Cloud Infrastructure Media Services service.

Creates a new MediaWorkflow.


## Example Usage

```hcl
resource "oci_media_services_media_workflow" "test_media_workflow" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.media_workflow_display_name

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	locks {
		#Required
		compartment_id = var.compartment_id
		type = var.media_workflow_locks_type

		#Optional
		message = var.media_workflow_locks_message
		related_resource_id = oci_usage_proxy_resource.test_resource.id
		time_created = var.media_workflow_locks_time_created
	}
	media_workflow_configuration_ids = var.media_workflow_media_workflow_configuration_ids
	parameters = var.media_workflow_parameters
	tasks {
		#Required
		key = var.media_workflow_tasks_key
		parameters = var.media_workflow_tasks_parameters
		type = var.media_workflow_tasks_type
		version = var.media_workflow_tasks_version

		#Optional
		enable_parameter_reference = var.media_workflow_tasks_enable_parameter_reference
		enable_when_referenced_parameter_equals = var.media_workflow_tasks_enable_when_referenced_parameter_equals
		prerequisites = var.media_workflow_tasks_prerequisites
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment Identifier.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) Name for the MediaWorkflow. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `locks` - (Optional) Locks associated with this resource.
	* `compartment_id` - (Required) (Updatable) The compartment ID of the lock.
	* `message` - (Optional) A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - (Optional) The ID of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - (Optional) When the lock was created.
	* `type` - (Required) Type of the lock.
* `media_workflow_configuration_ids` - (Optional) (Updatable) Configurations to be applied to all the jobs for this workflow. Parameters in these configurations are overridden by parameters in the MediaWorkflowConfigurations of the MediaWorkflowJob and the parameters of the MediaWorkflowJob. 
* `parameters` - (Optional) (Updatable) JSON object representing named parameters and their default values that can be referenced throughout this workflow. The values declared here can be overridden by the MediaWorkflowConfigurations or parameters supplied when creating MediaWorkflowJobs from this MediaWorkflow. 
* `tasks` - (Optional) (Updatable) The processing to be done in this workflow. Each key of the MediaWorkflowTasks in this array must be unique within the array. The order of tasks given here will be preserved. 
	* `enable_parameter_reference` - (Optional) (Updatable) Allows this task to be conditionally enabled.  If no value or a blank value is given, the task is unconditionally enbled.  Otherwise the given string specifies a parameter of the job created for this task's workflow using the JSON pointer syntax. The JSON pointer is validated when a job is created from the workflow of this task. 
	* `enable_when_referenced_parameter_equals` - (Optional) (Updatable) Used in conjunction with enableParameterReference to conditionally enable a task.  When a job is created from the workflow of this task, the task will only be enabled if the value of the parameter specified by enableParameterReference is equal to the value of this property. This property must be prenset if and only if a enableParameterReference is given. The value is a JSON node. 
	* `key` - (Required) (Updatable) A unique identifier for this task within its workflow. Keys are used to reference a task within workflows and MediaWorkflowJobs. Tasks are referenced as prerequisites and to track output and state. 
	* `parameters` - (Required) (Updatable) Data specifiying how this task is to be run. The data is a JSON object that must conform to the JSON Schema specified by the parameters of the MediaWorkflowTaskDeclaration this task references. The parameters may contain values or references to other parameters. 
	* `prerequisites` - (Optional) (Updatable) Keys to the other tasks in this workflow that must be completed before execution of this task can begin. 
	* `type` - (Required) (Updatable) The type of process to run at this task. Refers to the name of a MediaWorkflowTaskDeclaration. 
	* `version` - (Required) (Updatable) The version of the MediaWorkflowTaskDeclaration.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Name of the Media Workflow. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation.
* `lifecyle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `locks` - Locks associated with this resource.
	* `compartment_id` - The compartment ID of the lock.
	* `message` - A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - The ID of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - When the lock was created.
	* `type` - Type of the lock.
* `media_workflow_configuration_ids` - Configurations to be applied to all the runs of this workflow. Parameters in these configurations are overridden by parameters in the MediaWorkflowConfigurations of the MediaWorkflowJob and the parameters of the MediaWorkflowJob. If the same parameter appears in multiple configurations, the values that appear in the configuration at the highest index will be used. 
* `parameters` - JSON object representing named parameters and their default values that can be referenced throughout this workflow. The values declared here can be overridden by the MediaWorkflowConfigurations or parameters supplied when creating MediaWorkflowJobs from this MediaWorkflow. 
* `state` - The current state of the MediaWorkflow.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `tasks` - The processing to be done in this workflow. Each key of the MediaWorkflowTasks in this array is unique within the array.  The order of the items is preserved from the order of the tasks array in CreateMediaWorkflowDetails or UpdateMediaWorkflowDetails. 
	* `enable_parameter_reference` - Allows this task to be conditionally enabled.  If no value or a blank value is given, the task is unconditionally enbled.  Otherwise the given string specifies a parameter of the job created for this task's workflow using the JSON pointer syntax. The JSON pointer is validated when a job is created from the workflow of this task. 
	* `enable_when_referenced_parameter_equals` - Used in conjunction with enableParameterReference to conditionally enable a task.  When a job is created from the workflow of this task, the task will only be enabled if the value of the parameter specified by enableParameterReference is equal to the value of this property. This property must be prenset if and only if a enableParameterReference is given. The value is a JSON node. 
	* `key` - A unique identifier for this task within its workflow. Keys are used to reference a task within workflows and MediaWorkflowJobs. Tasks are referenced as prerequisites and to track output and state. 
	* `parameters` - Data specifiying how this task is to be run. The data is a JSON object that must conform to the JSON Schema specified by the parameters of the MediaWorkflowTaskDeclaration this task references. The parameters may contain values or references to other parameters. 
	* `prerequisites` - Keys to the other tasks in this workflow that must be completed before execution of this task can begin. 
	* `type` - The type of process to run at this task. Refers to the name of a MediaWorkflowTaskDeclaration. 
	* `version` - The version of the MediaWorkflowTaskDeclaration.
* `time_created` - The time when the MediaWorkflow was created. An RFC3339 formatted datetime string.
* `time_updated` - The time when the MediaWorkflow was updated. An RFC3339 formatted datetime string.
* `version` - The version of the MediaWorkflow.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Media Workflow
	* `update` - (Defaults to 20 minutes), when updating the Media Workflow
	* `delete` - (Defaults to 20 minutes), when destroying the Media Workflow


## Import

MediaWorkflows can be imported using the `id`, e.g.

```
$ terraform import oci_media_services_media_workflow.test_media_workflow "id"
```

