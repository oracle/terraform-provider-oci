---
subcategory: "Media Services"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_media_services_system_media_workflow"
sidebar_current: "docs-oci-datasource-media_services-system_media_workflow"
description: |-
  Provides details about a specific System Media Workflow in Oracle Cloud Infrastructure Media Services service
---

# Data Source: oci_media_services_system_media_workflow
This data source provides details about a specific System Media Workflow resource in Oracle Cloud Infrastructure Media Services service.

Lists the SystemMediaWorkflows that can be used to run a job by name or as a template to create a MediaWorkflow.


## Example Usage

```hcl
data "oci_media_services_system_media_workflow" "test_system_media_workflow" {

	#Optional
	compartment_id = var.compartment_id
	name = var.system_media_workflow_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `name` - (Optional) A filter to return only the resources with their system defined, unique name matching the given name.


## Attributes Reference

The following attributes are exported:

* `items` - List of SytemMediaWorkflow items.
	* `description` - Description of this workflow's processing and how that processing can be customized by specifying parameter values. 
	* `name` - System provided unique identifier for this static media workflow.
	* `parameters` - JSON object representing named parameters and their default values that can be referenced throughout this workflow. The values declared here can be overridden by the MediaWorkflowConfigurations or parameters supplied when creating MediaWorkflowJobs from this MediaWorkflow. 
	* `tasks` - The processing to be done in this workflow. Each key of the MediaWorkflowTasks in this array is unique within the array. The order of the items is preserved from the order of the tasks array in CreateMediaWorkflowDetails or UpdateMediaWorkflowDetails. 
		* `enable_parameter_reference` - Allows this task to be conditionally enabled.  If no value or a blank value is given, the task is unconditionally enbled.  Otherwise the given string specifies a parameter of the job created for this task's workflow using the JSON pointer syntax. The JSON pointer is validated when a job is created from the workflow of this task. 
		* `enable_when_referenced_parameter_equals` - Used in conjunction with enableParameterReference to conditionally enable a task.  When a job is created from the workflow of this task, the task will only be enabled if the value of the parameter specified by enableParameterReference is equal to the value of this property. This property must be prenset if and only if a enableParameterReference is given. The value is a JSON node. 
		* `key` - A unique identifier for this task within its workflow. Keys are used to reference a task within workflows and MediaWorkflowJobs. Tasks are referenced as prerequisites and to track output and state. 
		* `parameters` - Data specifiying how this task is to be run. The data is a JSON object that must conform to the JSON Schema specified by the parameters of the MediaWorkflowTaskDeclaration this task references. The parameters may contain values or references to other parameters. 
		* `prerequisites` - Keys to the other tasks in this workflow that must be completed before execution of this task can begin. 
		* `type` - The type of process to run at this task. Refers to the name of a MediaWorkflowTaskDeclaration. 
		* `version` - The version of the MediaWorkflowTaskDeclaration.

