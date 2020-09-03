---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_compartments"
sidebar_current: "docs-oci-datasource-identity-compartments"
description: |-
  Provides the list of Compartments in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_compartments
This data source provides the list of Compartments in Oracle Cloud Infrastructure Identity service.

Lists the compartments in a specified compartment. The members of the list
returned depends on the values set for several parameters.

With the exception of the tenancy (root compartment), the ListCompartments operation
returns only the first-level child compartments in the parent compartment specified in
`compartmentId`. The list does not include any subcompartments of the child
compartments (grandchildren).

The parameter `accessLevel` specifies whether to return only those compartments for which the
requestor has INSPECT permissions on at least one resource directly
or indirectly (the resource can be in a subcompartment).

The parameter `compartmentIdInSubtree` applies only when you perform ListCompartments on the
tenancy (root compartment). When set to true, the entire hierarchy of compartments can be returned.
To get a full list of all compartments and subcompartments in the tenancy (root compartment),
set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ANY.

See [Where to Get the Tenancy's OCID and User's OCID](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/apisigningkey.htm#five).


## Example Usage

```hcl
data "oci_identity_compartments" "test_compartments" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.compartment_access_level
	compartment_id_in_subtree = var.compartment_compartment_id_in_subtree
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are `ANY` and `ACCESSIBLE`. Default is `ANY`. Setting this to `ACCESSIBLE` returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). For the compartments on which the user indirectly has INSPECT permissions, a restricted set of fields is returned.

	When set to `ANY` permissions are not checked. 
* `compartment_id` - (Required) The OCID of the compartment (remember that the tenancy is simply the root compartment). 
* `compartment_id_in_subtree` - (Optional) Default is false. Can only be set to true when performing ListCompartments on the tenancy (root compartment). When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned depending on the the setting of `accessLevel`. 


## Attributes Reference

The following attributes are exported:

* `compartments` - The list of compartments.

### Compartment Reference

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

