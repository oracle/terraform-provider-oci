---
subcategory: "Resource Manager"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_resourcemanager_stacks"
sidebar_current: "docs-oci-datasource-resourcemanager-stacks"
description: |-
  Provides the list of Stacks in Oracle Cloud Infrastructure Resource Manager service
---

# Data Source: oci_resourcemanager_stacks
This data source provides the list of Stacks in Oracle Cloud Infrastructure Resource Manager service.

Lists stacks according to the specified filter.
For more information, see
[Listing Stacks](https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Tasks/list-stacks.htm).
- If called using the compartment ID, returns all stacks in the specified compartment.
- If called using the stack ID, returns the specified stack. (See also [GetStack](https://docs.cloud.oracle.com/iaas/api/#/en/resourcemanager/latest/Stack/GetStack).)


## Example Usage

```hcl
data "oci_resourcemanager_stacks" "test_stacks" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.stack_display_name
	id = var.stack_id
	state = var.stack_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) on which to filter.
* `display_name` - (Optional) Display name on which to query.
* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) on which to query for a stack.
* `state` - (Optional) A filter that returns only those resources that match the specified lifecycle state. The state value is case-insensitive. For more information about stack lifecycle states, see [Key Concepts](https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm#concepts__StackStates).


## Attributes Reference

The following attributes are exported:

* `stacks` - The list of stacks.

### Stack Reference

The following attributes are exported:

* `compartment_id` - Unique identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) for the compartment where the stack is located.
* `config_source` - Location of the Terraform configuration. 
	* `config_source_type` - Specifies the `configSourceType` for uploading the Terraform configuration. Presently, the .zip file type (`ZIP_UPLOAD`) is the only supported `configSourceType`. 
	* `working_directory` - File path to the directory from which Terraform runs. If not specified, we use the root directory. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - General description of the stack.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags associated with the resource. Each tag is a key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - Unique identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) for the stack.
* `stack_drift_status` - Drift status of the stack. Drift refers to differences between the actual (current) state of the stack and the expected (defined) state of the stack. 
* `state` - The current lifecycle state of the stack. For more information about stack lifecycle states in Resource Manager, see [Key Concepts](https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm#concepts__StackStates). 
* `system_tags` - The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{orcl-cloud: {free-tier-retain: true}}` 
* `terraform_version` - The version of Terraform specified for the stack. Example: `1.5.x` 
* `time_created` - The date and time at which the stack was created. Format is defined by RFC3339. Example: `2020-01-25T21:10:29.600Z` 
* `time_drift_last_checked` - The date and time when the drift detection was last executed. Format is defined by RFC3339. Example: `2020-01-25T21:10:29.600Z` 
* `variables` - Terraform variables associated with this resource. Maximum number of variables supported is 250. The maximum size of each variable, including both name and value, is 8192 bytes. Example: `{"CompartmentId": "compartment-id-value"}`

