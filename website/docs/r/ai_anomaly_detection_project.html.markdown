---
subcategory: "Ai Anomaly Detection"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_anomaly_detection_project"
sidebar_current: "docs-oci-resource-ai_anomaly_detection-project"
description: |-
  Provides the Project resource in Oracle Cloud Infrastructure Ai Anomaly Detection service
---

# oci_ai_anomaly_detection_project
This resource provides the Project resource in Oracle Cloud Infrastructure Ai Anomaly Detection service.

Creates a new Project.


## Example Usage

```hcl
resource "oci_ai_anomaly_detection_project" "test_project" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.project_description
	display_name = var.project_display_name
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID for the project's compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) A short description of the project.
* `display_name` - (Optional) (Updatable) A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID for the project's compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A short description of the project.
* `display_name` - A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the project that is immutable on creation.
* `state` - The lifecycle state of the Project.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the resource was created in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_updated` - The date and time the resource was updated in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Project
	* `update` - (Defaults to 20 minutes), when updating the Project
	* `delete` - (Defaults to 20 minutes), when destroying the Project


## Import

Projects can be imported using the `id`, e.g.

```
$ terraform import oci_ai_anomaly_detection_project.test_project "id"
```

