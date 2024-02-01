---
subcategory: "Media Services"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_media_services_media_workflow_job"
sidebar_current: "docs-oci-datasource-media_services-media_workflow_job"
description: |-
  Provides details about a specific Media Workflow Job in Oracle Cloud Infrastructure Media Services service
---

# Data Source: oci_media_services_media_workflow_job
This data source provides details about a specific Media Workflow Job resource in Oracle Cloud Infrastructure Media Services service.

Gets the MediaWorkflowJob.

## Example Usage

```hcl
data "oci_media_services_media_workflow_job" "test_media_workflow_job" {
	#Required
	media_workflow_job_id = oci_media_services_media_workflow_job.test_media_workflow_job.id
}
```

## Argument Reference

The following arguments are supported:

* `media_workflow_job_id` - (Required) Unique MediaWorkflowJob identifier.


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

