---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_task_record"
sidebar_current: "docs-oci-datasource-fleet_apps_management-task_record"
description: |-
  Provides details about a specific Task Record in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_task_record
This data source provides details about a specific Task Record resource in Oracle Cloud Infrastructure Fleet Apps Management service.

Gets a Task by identifier

## Example Usage

```hcl
data "oci_fleet_apps_management_task_record" "test_task_record" {
	#Required
	task_record_id = oci_fleet_apps_management_task_record.test_task_record.id
}
```

## Argument Reference

The following arguments are supported:

* `task_record_id` - (Required) unique TaskDetail identifier


## Attributes Reference

The following attributes are exported:

* `compartment_id` - 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A user-friendly description. To provide some insight about the resource. Avoid entering confidential information. 
* `details` - The details of the task.
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
	* `is_apply_subject_task` - Is this an Apply Subject Task?  Set this to true for a Patch Execution Task which applies patches(subjects) on a target. 
	* `is_discovery_output_task` - Is this a discovery output task?
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
* `state` - The current state of the TaskRecord.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.
* `type` - Task type.
* `version` - The version of the task

