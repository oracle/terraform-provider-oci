---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_compartment"
sidebar_current: "docs-oci-datasource-identity-compartment"
description: |-
  Provides details about a specific Compartment in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_compartment
This data source provides details about a specific Compartment resource in Oracle Cloud Infrastructure Identity service.

Gets the specified compartment's information.

This operation does not return a list of all the resources inside the compartment. There is no single
API operation that does that. Compartments can contain multiple types of resources (instances, block
storage volumes, etc.). To find out what's in a compartment, you must call the "List" operation for
each resource type and specify the compartment's OCID as a query parameter in the request. For example,
call the [ListInstances](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Instance/ListInstances) operation in the Cloud Compute
Service or the [ListVolumes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Volume/ListVolumes) operation in Cloud Block Storage.


## Example Usage

```hcl
data "oci_identity_compartment" "test_compartment" {
	#Required
	id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Required) The OCID of the compartment.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the parent compartment containing the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description you assign to the compartment. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the compartment.
* `inactive_state` - The detailed status of INACTIVE lifecycleState.
* `is_accessible` - Indicates whether or not the compartment is accessible for the user making the request. Returns true when the user has INSPECT permissions directly on a resource in the compartment or indirectly (permissions can be on a resource in a subcompartment). 
* `name` - The name you assign to the compartment during creation. The name must be unique across all compartments in the parent. Avoid entering confidential information. 
* `state` - The compartment's current state.
* `time_created` - Date and time the compartment was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

