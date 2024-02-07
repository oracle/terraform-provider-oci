---
subcategory: "Media Services"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_media_services_media_workflow_configuration"
sidebar_current: "docs-oci-datasource-media_services-media_workflow_configuration"
description: |-
  Provides details about a specific Media Workflow Configuration in Oracle Cloud Infrastructure Media Services service
---

# Data Source: oci_media_services_media_workflow_configuration
This data source provides details about a specific Media Workflow Configuration resource in Oracle Cloud Infrastructure Media Services service.

Gets a MediaWorkflowConfiguration by identifier

## Example Usage

```hcl
data "oci_media_services_media_workflow_configuration" "test_media_workflow_configuration" {
	#Required
	media_workflow_configuration_id = oci_media_services_media_workflow_configuration.test_media_workflow_configuration.id
}
```

## Argument Reference

The following arguments are supported:

* `media_workflow_configuration_id` - (Required) Unique MediaWorkflowConfiguration identifier.


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

