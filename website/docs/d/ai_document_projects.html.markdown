---
subcategory: "Ai Document"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_document_projects"
sidebar_current: "docs-oci-datasource-ai_document-projects"
description: |-
  Provides the list of Projects in Oracle Cloud Infrastructure Ai Document service
---

# Data Source: oci_ai_document_projects
This data source provides the list of Projects in Oracle Cloud Infrastructure Ai Document service.

Returns a list of projects.


## Example Usage

```hcl
data "oci_ai_document_projects" "test_projects" {

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
* `id` - (Optional) The filter to find the project with the given identifier.
* `state` - (Optional) The filter to match projects with the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `project_collection` - The list of project_collection.

### Project Reference

The following attributes are exported:

* `compartment_id` - The compartment identifier.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For example: `{"foo-namespace": {"bar-key": "value"}}` 
* `description` - An optional description of the project.
* `display_name` - A human-friendly name for the project, which can be changed.
* `freeform_tags` - A simple key-value pair that is applied without any predefined name, type, or scope. It exists for cross-compatibility only. For example: `{"bar-key": "value"}` 
* `id` - A unique identifier that is immutable after creation.
* `lifecycle_details` - A message describing the current state in more detail, that can provide actionable information if creation failed.
* `state` - The current state of the project.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. For example: `{"orcl-cloud": {"free-tier-retained": "true"}}` 
* `time_created` - When the project was created, as an RFC3339 datetime string.
* `time_updated` - When the project was updated, as an RFC3339 datetime string.

