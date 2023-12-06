---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_build_pipelines"
sidebar_current: "docs-oci-datasource-devops-build_pipelines"
description: |-
  Provides the list of Build Pipelines in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_build_pipelines
This data source provides the list of Build Pipelines in Oracle Cloud Infrastructure Devops service.

Returns a list of build pipelines.


## Example Usage

```hcl
data "oci_devops_build_pipelines" "test_build_pipelines" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.build_pipeline_display_name
	id = var.build_pipeline_id
	project_id = oci_devops_project.test_project.id
	state = var.build_pipeline_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) Unique identifier or OCID for listing a single resource by ID.
* `project_id` - (Optional) unique project identifier
* `state` - (Optional) A filter to return only build pipelines that matches the given lifecycle state.


## Attributes Reference

The following attributes are exported:

* `build_pipeline_collection` - The list of build_pipeline_collection.

### BuildPipeline Reference

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

