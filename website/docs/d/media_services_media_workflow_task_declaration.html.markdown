---
subcategory: "Media Services"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_media_services_media_workflow_task_declaration"
sidebar_current: "docs-oci-datasource-media_services-media_workflow_task_declaration"
description: |-
  Provides details about a specific Media Workflow Task Declaration in Oracle Cloud Infrastructure Media Services service
---

# Data Source: oci_media_services_media_workflow_task_declaration
This data source provides details about a specific Media Workflow Task Declaration resource in Oracle Cloud Infrastructure Media Services service.

Returns a list of MediaWorkflowTaskDeclarations.


## Example Usage

```hcl
data "oci_media_services_media_workflow_task_declaration" "test_media_workflow_task_declaration" {

	#Optional
	compartment_id = var.compartment_id
	is_current = var.media_workflow_task_declaration_is_current
	name = var.media_workflow_task_declaration_name
	version = var.media_workflow_task_declaration_version
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `is_current` - (Optional) A filter to only select the newest version for each MediaWorkflowTaskDeclaration name.
* `name` - (Optional) A filter to return only the resources with their system defined, unique name matching the given name.
* `version` - (Optional) A filter to select MediaWorkflowTaskDeclaration by version.


## Attributes Reference

The following attributes are exported:

* `items` - List of MediaWorkflowTaskDeclaration objects.
	* `name` - MediaWorkflowTaskDeclaration identifier. The name and version should be unique among MediaWorkflowTaskDeclarations. 
	* `parameters_schema` - JSON schema specifying the parameters supported by this type of task. This is used to validate tasks' parameters when jobs are created. 
	* `parameters_schema_allowing_references` - JSON schema similar to the parameterSchema, but permits parameter values to refer to other parameters using the ${/path/to/another/parmeter} syntax.  This is used to validate task parameters when workflows are created. 
	* `version` - The version of MediaWorkflowTaskDeclaration, incremented whenever the team implementing the task processor modifies the JSON schema of this declaration's definitions, parameters or list of required parameters. 

