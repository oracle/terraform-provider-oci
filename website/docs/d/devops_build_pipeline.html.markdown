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

Gets a BuildPipeline by identifier

## Example Usage

```hcl
data "oci_devops_build_pipeline" "test_build_pipeline" {
	#Required
	build_pipeline_id = oci_devops_build_pipeline.test_build_pipeline.id
}
```

## Argument Reference

The following arguments are supported:

* `build_pipeline_id` - (Required) unique BuildPipeline identifier


## Attributes Reference

The following attributes are exported:

* `build_pipeline_parameters` - Specifies list of parameters present in BuildPipeline. In case of UPDATE operation, replaces existing parameters list. Merging with existing parameters is not supported. 
	* `items` - List of Parameters defined for a BuildPipeline.
		* `default_value` - Default value of the parameter
		* `description` - Description of the parameter
		* `name` - Name of the parameter (Case-sensitive). Example: 'Pipeline_param' is not same as 'pipeline_Param' 
* `compartment_id` - Compartment Identifier
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `description` - Optional description about the BuildPipeline
* `display_name` - BuildPipeline identifier which can be renamed and is not necessarily unique
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `id` - Unique identifier that is immutable on creation
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `project_id` - Project Identifier
* `state` - The current state of the BuildPipeline.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - The time the BuildPipeline was created. An RFC3339 formatted datetime string
* `time_updated` - The time at which the BuildPipeline was updated. An RFC3339 formatted datetime string

