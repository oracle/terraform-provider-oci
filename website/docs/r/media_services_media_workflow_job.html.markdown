---
subcategory: "Media Services"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_media_services_media_workflow_job"
sidebar_current: "docs-oci-resource-media_services-media_workflow_job"
description: |-
  Provides the Media Workflow Job resource in Oracle Cloud Infrastructure Media Services service
---

# oci_media_services_media_workflow_job
This resource provides the Media Workflow Job resource in Oracle Cloud Infrastructure Media Services service.

Run the MediaWorkflow according to the given mediaWorkflow definition and configuration.

## Example Usage

```hcl
resource "oci_media_services_media_workflow_job" "test_media_workflow_job" {
	#Required
	compartment_id = var.compartment_id
	workflow_identifier_type = var.media_workflow_job_workflow_identifier_type

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	display_name = var.media_workflow_job_display_name
	freeform_tags = {"bar-key"= "value"}
	locks {
		#Required
		compartment_id = var.compartment_id
		type = var.media_workflow_job_locks_type

		#Optional
		message = var.media_workflow_job_locks_message
		related_resource_id = oci_usage_proxy_resource.test_resource.id
		time_created = var.media_workflow_job_locks_time_created
	}
	media_workflow_configuration_ids = var.media_workflow_job_media_workflow_configuration_ids
	media_workflow_id = oci_media_services_media_workflow.test_media_workflow.id
	media_workflow_name = oci_media_services_media_workflow.test_media_workflow.name
	parameters = var.media_workflow_job_parameters
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) ID of the compartment in which the job should be created.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Optional) (Updatable) Name of the Media Workflow Job. Does not have to be unique. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `locks` - (Optional) Locks associated with this resource.
	* `compartment_id` - (Required) (Updatable) The compartment ID of the lock.
	* `message` - (Optional) A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - (Optional) The ID of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - (Optional) When the lock was created.
	* `type` - (Required) Type of the lock.
* `media_workflow_configuration_ids` - (Optional) Configurations to be applied to this run of the workflow.
* `media_workflow_id` - (Applicable when workflow_identifier_type=ID) OCID of the MediaWorkflow that should be run.
* `media_workflow_name` - (Applicable when workflow_identifier_type=NAME) Name of the system MediaWorkflow that should be run.
* `parameters` - (Optional) Parameters that override parameters specified in MediaWorkflowTaskDeclarations, the MediaWorkflow, the MediaWorkflow's MediaWorkflowConfigurations and the MediaWorkflowConfigurations of this MediaWorkflowJob. The parameters are given as JSON. The top level and 2nd level elements must be JSON objects (vs arrays, scalars, etc). The top level keys refer to a task's key and the 2nd level keys refer to a parameter's name. 
* `workflow_identifier_type` - (Required) Discriminate identification of a workflow by name versus a workflow by ID.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Name of the Media Workflow Job. Does not have to be unique. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier for this run of the workflow.
* `lifecycle_details` - The lifecyle details.
* `locks` - Locks associated with this resource.
	* `compartment_id` - The compartment ID of the lock.
	* `message` - A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - The ID of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - When the lock was created.
	* `type` - Type of the lock.
* `media_workflow_configuration_ids` - Configurations to be applied to this run of the workflow.
* `media_workflow_id` - The workflow to execute.
* `outputs` - A list of JobOutput for the workflowJob.
	* `asset_type` - Type of job output.
	* `bucket` - The bucket name of the job output.
	* `id` - The ID associated with the job output.
	* `namespace` - The namespace name of the job output.
	* `object` - The object name of the job output.
* `parameters` - Parameters that override parameters specified in MediaWorkflowTaskDeclarations, the MediaWorkflow, the MediaWorkflow's MediaWorkflowConfigurations and the MediaWorkflowConfigurations of this MediaWorkflowJob. The parameters are given as JSON.  The top level and 2nd level elements must be JSON objects (vs arrays, scalars, etc). The top level keys refer to a task's key and the 2nd level keys refer to a parameter's name. 
* `runnable` - A JSON representation of the job as it will be run by the system. All the task declarations, configurations and parameters are merged. Parameter values are all fully resolved. 
* `state` - The current state of the MediaWorkflowJob.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `task_lifecycle_state` - Status of each task.
	* `key` - Unique key within a MediaWorkflowJob for the task.
	* `lifecycle_details` - The lifecycle details of MediaWorkflowJob task.
	* `state` - The current state of the MediaWorkflowJob task.
* `time_created` - Creation time of the job. An RFC3339 formatted datetime string.
* `time_ended` - Time when the job finished. An RFC3339 formatted datetime string.
* `time_started` - Time when the job started to execute. An RFC3339 formatted datetime string.
* `time_updated` - Updated time of the job. An RFC3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Media Workflow Job
	* `update` - (Defaults to 20 minutes), when updating the Media Workflow Job
	* `delete` - (Defaults to 20 minutes), when destroying the Media Workflow Job


## Import

MediaWorkflowJobs can be imported using the `id`, e.g.

```
$ terraform import oci_media_services_media_workflow_job.test_media_workflow_job "id"
```

