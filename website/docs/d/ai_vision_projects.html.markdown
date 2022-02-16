---
subcategory: "Ai Vision"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_vision_projects"
sidebar_current: "docs-oci-datasource-ai_vision-projects"
description: |-
  Provides the list of Projects in Oracle Cloud Infrastructure Ai Vision service
---

# Data Source: oci_ai_vision_projects
This data source provides the list of Projects in Oracle Cloud Infrastructure Ai Vision service.

Returns a list of Projects.


## Example Usage

```hcl
data "oci_ai_vision_projects" "test_projects" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.project_display_name
	id = var.project_id
	state = var.project_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) unique Project identifier
* `state` - (Optional) A filter to return only resources their lifecycleState matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `project_collection` - The list of project_collection.

### Project Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A short description of the project.
* `display_name` - Project Identifier, can be renamed
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `state` - The current state of the Project.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the Project was created. An RFC3339 formatted datetime string
* `time_updated` - The time the Project was updated. An RFC3339 formatted datetime string

