---
subcategory: "Media Services"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_media_services_media_workflow_configuration"
sidebar_current: "docs-oci-resource-media_services-media_workflow_configuration"
description: |-
  Provides the Media Workflow Configuration resource in Oracle Cloud Infrastructure Media Services service
---

# oci_media_services_media_workflow_configuration
This resource provides the Media Workflow Configuration resource in Oracle Cloud Infrastructure Media Services service.

Creates a new MediaWorkflowConfiguration.


## Example Usage

```hcl
resource "oci_media_services_media_workflow_configuration" "test_media_workflow_configuration" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.media_workflow_configuration_display_name
	parameters = var.media_workflow_configuration_parameters

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	locks {
		#Required
		compartment_id = var.compartment_id
		type = var.media_workflow_configuration_locks_type

		#Optional
		message = var.media_workflow_configuration_locks_message
		related_resource_id = oci_usage_proxy_resource.test_resource.id
		time_created = var.media_workflow_configuration_locks_time_created
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment Identifier.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) MediaWorkflowConfiguration identifier. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `locks` - (Optional) Locks associated with this resource.
	* `compartment_id` - (Required) (Updatable) The compartment ID of the lock.
	* `message` - (Optional) A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - (Optional) The ID of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - (Optional) When the lock was created.
	* `type` - (Required) Type of the lock.
* `parameters` - (Required) (Updatable) Reuseable parameter values encoded as a JSON; the top and second level JSON elements are objects. Each key of the top level object refers to a task key that is unqiue to the workflow, each of the second level objects' keys refer to the name of a parameter that is unique to the task. taskKey -> parameterName -> parameterValue 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Display name for the MediaWorkflowConfiguration. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation.
* `lifecyle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `locks` - Locks associated with this resource.
	* `compartment_id` - The compartment ID of the lock.
	* `message` - A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - The ID of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - When the lock was created.
	* `type` - Type of the lock.
* `parameters` - Reuseable parameter values encoded as a JSON; the top and second level JSON elements are objects. Each key of the top level object refer to a task key that is unqiue to the workflow, each of the second level objects' keys refer to the name of a parameter that is unique to the task. taskKey -> parameterName -> parameterValue 
* `state` - The current state of the MediaWorkflowConfiguration.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when the the MediaWorkflowConfiguration was created. An RFC3339 formatted datetime string.
* `time_updated` - The time when the MediaWorkflowConfiguration was updated. An RFC3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Media Workflow Configuration
	* `update` - (Defaults to 20 minutes), when updating the Media Workflow Configuration
	* `delete` - (Defaults to 20 minutes), when destroying the Media Workflow Configuration


## Import

MediaWorkflowConfigurations can be imported using the `id`, e.g.

```
$ terraform import oci_media_services_media_workflow_configuration.test_media_workflow_configuration "id"
```

