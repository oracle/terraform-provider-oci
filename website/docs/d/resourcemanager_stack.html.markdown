---
subcategory: "Resourcemanager"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_resourcemanager_stack"
sidebar_current: "docs-oci-datasource-resourcemanager-stack"
description: |-
  Provides details about a specific Stack in Oracle Cloud Infrastructure Resourcemanager service
---

# Data Source: oci_resourcemanager_stack
This data source provides details about a specific Stack resource in Oracle Cloud Infrastructure Resourcemanager service.

Gets a stack using the stack ID.

## Example Usage

```hcl
data "oci_resourcemanager_stack" "test_stack" {
	#Required
	stack_id = "${oci_resourcemanager_stack.test_stack.id}"
}
```

## Argument Reference

The following arguments are supported:

* `stack_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stack.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - Unique identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) for the compartment where the stack is located.
* `config_source` - 
	* `config_source_type` - Specifies the `configSourceType` for uploading the Terraform configuration. Presently, the .zip file type (`ZIP_UPLOAD`) is the only supported `configSourceType`. 
	* `working_directory` - File path to the directory from which Terraform runs. If not specified, we use the root directory. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - General description of the stack.
* `display_name` - Human-readable display name for the stack.
* `freeform_tags` - Free-form tags associated with this resource. Each tag is a key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - Unique identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) for the stack.
* `state` - The current lifecycle state of the stack.
* `time_created` - The date and time at which the stack was created.
* `variables` - Terraform variables associated with this resource. Maximum number of variables supported is 100. The maximum size of each variable, including both name and value, is 4096 bytes. Example: `{"CompartmentId": "compartment-id-value"}` 

