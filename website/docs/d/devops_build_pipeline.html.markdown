---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_build_pipeline"
sidebar_current: "docs-oci-datasource-devops-build_pipeline"
description: |-
  Provides details about a specific Build Pipeline in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_build_pipeline
This data source provides details about a specific Build Pipeline resource in Oracle Cloud Infrastructure Devops service.

Retrieves a build pipeline by identifier.

## Example Usage

```hcl
data "oci_devops_build_pipeline" "test_build_pipeline" {
	#Required
	build_pipeline_id = oci_devops_build_pipeline.test_build_pipeline.id
}
```

## Argument Reference

The following arguments are supported:

* `build_pipeline_id` - (Required) Unique build pipeline identifier.


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

