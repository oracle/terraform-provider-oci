---
layout: "oci"
page_title: "OCI: oci_identity_compartment"
sidebar_current: "docs-oci-resource-identity-compartment"
description: |-
Creates and manages an OCI Compartment
---

# oci_identity_compartment
The `oci_identity_compartment` resource creates and manages an OCI Compartment

Creates a new compartment in your tenancy.

If a compartment with the given `name` already exists, then that compartment will be used instead of creating a new compartment.

**Important:** Compartments cannot be deleted.

You must specify your tenancy's OCID as the compartment ID in the request object. Remember that the tenancy
is simply the root compartment. For information about OCIDs, see
[Resource Identifiers](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).

You must also specify a *name* for the compartment, which must be unique across all compartments in
your tenancy. You can use this name or the OCID when writing policies that apply
to the compartment. For more information about policies, see
[How Policies Work](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/policies.htm).

You must also specify a *description* for the compartment (although it can be an empty string). It does
not have to be unique, and you can change it anytime with
[UpdateCompartment](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/Compartment/UpdateCompartment).

After you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using the
object, first make sure its `lifecycleState` has changed to ACTIVE.


## Example Usage

```hcl
resource "oci_identity_compartment" "test_compartment" {
	#Required
	compartment_id = "${var.tenancy_ocid}"
	description = "${var.compartment_description}"
	name = "${var.compartment_name}"

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Required) (Updatable) The description you assign to the compartment during creation. Does not have to be unique, and it's changeable. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `name` - (Required) (Updatable) The name you assign to the compartment during creation. The name must be unique across all compartments in the tenancy. Avoid entering confidential information. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the tenancy containing the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description you assign to the compartment. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the compartment.
* `inactive_state` - The detailed status of INACTIVE lifecycleState.
* `name` - The name you assign to the compartment during creation. The name must be unique across all compartments in the tenancy. Avoid entering confidential information. 
* `state` - The compartment's current state. After creating a compartment, make sure its `lifecycleState` changes from CREATING to ACTIVE before using it. 
* `time_created` - Date and time the compartment was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

## Import

Compartments can be imported using the `id`, e.g.

```
$ terraform import oci_identity_compartment.test_compartment "id"
```
