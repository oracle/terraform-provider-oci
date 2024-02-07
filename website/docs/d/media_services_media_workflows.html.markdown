---
subcategory: "Media Services"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_media_services_media_workflows"
sidebar_current: "docs-oci-datasource-media_services-media_workflows"
description: |-
  Provides the list of Media Workflows in Oracle Cloud Infrastructure Media Services service
---

# Data Source: oci_media_services_media_workflows
This data source provides the list of Media Workflows in Oracle Cloud Infrastructure Media Services service.

Lists the MediaWorkflows.

## Example Usage

```hcl
data "oci_media_services_media_workflows" "test_media_workflows" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.media_workflow_display_name
	id = var.media_workflow_id
	state = var.media_workflow_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only the resources that match the entire display name given.
* `id` - (Optional) Unique MediaWorkflow identifier.
* `state` - (Optional) A filter to return only the resources with lifecycleState matching the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `media_workflow_collection` - The list of media_workflow_collection.

### MediaWorkflow Reference

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

