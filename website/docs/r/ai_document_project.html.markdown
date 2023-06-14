---
subcategory: "Ai Document"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_document_project"
sidebar_current: "docs-oci-resource-ai_document-project"
description: |-
  Provides the Project resource in Oracle Cloud Infrastructure Ai Document service
---

# oci_ai_document_project
This resource provides the Project resource in Oracle Cloud Infrastructure Ai Document service.

Create a new project.


## Example Usage

```hcl
resource "oci_ai_document_project" "test_project" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	defined_tags = var.project_defined_tags
	description = var.project_description
	display_name = var.project_display_name
	freeform_tags = var.project_freeform_tags
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The compartment identifier.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For example: `{"foo-namespace": {"bar-key": "value"}}` 
* `description` - (Optional) (Updatable) An optional description of the project.
* `display_name` - (Optional) (Updatable) A human-friendly name for the project, that can be changed.
* `freeform_tags` - (Optional) (Updatable) A simple key-value pair that is applied without any predefined name, type, or scope. It exists for cross-compatibility only. For example: `{"bar-key": "value"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Project
	* `update` - (Defaults to 20 minutes), when updating the Project
	* `delete` - (Defaults to 20 minutes), when destroying the Project


## Import

Projects can be imported using the `id`, e.g.

```
$ terraform import oci_ai_document_project.test_project "id"
```

