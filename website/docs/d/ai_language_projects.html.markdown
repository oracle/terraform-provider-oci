---
subcategory: "Ai Language"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_language_projects"
sidebar_current: "docs-oci-datasource-ai_language-projects"
description: |-
  Provides the list of Projects in Oracle Cloud Infrastructure Ai Language service
---

# Data Source: oci_ai_language_projects
This data source provides the list of Projects in Oracle Cloud Infrastructure Ai Language service.

Returns a list of  Projects.


## Example Usage

```hcl
data "oci_ai_language_projects" "test_projects" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.project_display_name
	project_id = oci_ai_language_project.test_project.id
	state = var.project_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `project_id` - (Optional) The ID of the project for which to list the objects.
* `state` - (Optional) <b>Filter</b> results by the specified lifecycle state. Must be a valid state for the resource type. 


## Attributes Reference

The following attributes are exported:

* `project_collection` - The list of project_collection.

### Project Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)  for the project's compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A short description of a project.
* `display_name` - A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier OCID of the project
* `lifecycle_details` - A message describing the current state in more detail.
* `state` - The state of the project.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the resource was created in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_updated` - The date and time the resource was updated in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

