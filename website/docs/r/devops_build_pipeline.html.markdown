---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_build_pipeline"
sidebar_current: "docs-oci-resource-devops-build_pipeline"
description: |-
  Provides the Build Pipeline resource in Oracle Cloud Infrastructure Devops service
---

# oci_devops_build_pipeline
This resource provides the Build Pipeline resource in Oracle Cloud Infrastructure Devops service.

Creates a new build pipeline.


## Example Usage

```hcl
resource "oci_devops_build_pipeline" "test_build_pipeline" {
	#Required
	project_id = oci_devops_project.test_project.id

	#Optional
	build_pipeline_parameters {
		#Required
		items {
			#Required
			name = var.build_pipeline_build_pipeline_parameters_items_name

			#Optional
			default_value = var.build_pipeline_build_pipeline_parameters_items_default_value
			description = var.build_pipeline_build_pipeline_parameters_items_description
		}
	}
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.build_pipeline_description
	display_name = var.build_pipeline_display_name
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `build_pipeline_parameters` - (Optional) (Updatable) Specifies list of parameters present in a build pipeline. An UPDATE operation replaces the existing parameters list entirely. 
	* `items` - (Required) (Updatable) List of parameters defined for a build pipeline.
		* `default_value` - (Optional) (Updatable) Default value of the parameter.
		* `description` - (Optional) (Updatable) Description of the parameter.
		* `name` - (Required) (Updatable) Name of the parameter (case-sensitive). Parameter name must be ^[a-zA-Z][a-zA-Z_0-9]*$. Example: 'Build_Pipeline_param' is not same as 'build_pipeline_Param' 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `description` - (Optional) (Updatable) Optional description about the build pipeline.
* `display_name` - (Optional) (Updatable) Build pipeline display name. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `project_id` - (Required) The OCID of the DevOps project.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `build_pipeline_parameters` - Specifies list of parameters present in a build pipeline. An UPDATE operation replaces the existing parameters list entirely. 
	* `items` - List of parameters defined for a build pipeline.
		* `default_value` - Default value of the parameter.
		* `description` - Description of the parameter.
		* `name` - Name of the parameter (case-sensitive). Parameter name must be ^[a-zA-Z][a-zA-Z_0-9]*$. Example: 'Build_Pipeline_param' is not same as 'build_pipeline_Param' 
* `compartment_id` - The OCID of the compartment where the build pipeline is created.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `description` - Optional description about the build pipeline.
* `display_name` - Build pipeline display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `id` - Unique identifier that is immutable on creation.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `project_id` - The OCID of the DevOps project.
* `state` - The current state of the build pipeline.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - The time the build pipeline was created. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
* `time_updated` - The time the build pipeline was updated. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Build Pipeline
	* `update` - (Defaults to 20 minutes), when updating the Build Pipeline
	* `delete` - (Defaults to 20 minutes), when destroying the Build Pipeline


## Import

BuildPipelines can be imported using the `id`, e.g.

```
$ terraform import oci_devops_build_pipeline.test_build_pipeline "id"
```

