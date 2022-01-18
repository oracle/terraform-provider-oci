---
subcategory: "Ai Vision"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_vision_project"
sidebar_current: "docs-oci-datasource-ai_vision-project"
description: |-
  Provides details about a specific Project in Oracle Cloud Infrastructure Ai Vision service
---

# Data Source: oci_ai_vision_project
This data source provides details about a specific Project resource in Oracle Cloud Infrastructure Ai Vision service.

Gets a Project by identifier

## Example Usage

```hcl
data "oci_ai_vision_project" "test_project" {
	#Required
	project_id = oci_ai_vision_project.test_project.id
}
```

## Argument Reference

The following arguments are supported:

* `project_id` - (Required) unique Project identifier


## Attributes Reference

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

