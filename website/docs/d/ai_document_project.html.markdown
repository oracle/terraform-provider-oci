---
subcategory: "Ai Document"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_document_project"
sidebar_current: "docs-oci-datasource-ai_document-project"
description: |-
  Provides details about a specific Project in Oracle Cloud Infrastructure Ai Document service
---

# Data Source: oci_ai_document_project
This data source provides details about a specific Project resource in Oracle Cloud Infrastructure Ai Document service.

Get a project by identifier.

## Example Usage

```hcl
data "oci_ai_document_project" "test_project" {
	#Required
	project_id = oci_ai_document_project.test_project.id
}
```

## Argument Reference

The following arguments are supported:

* `project_id` - (Required) A unique project identifier.


## Attributes Reference

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

