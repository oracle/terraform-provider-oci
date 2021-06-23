---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_tag_namespace"
sidebar_current: "docs-oci-resource-identity-tag_namespace"
description: |-
  Provides the Tag Namespace resource in Oracle Cloud Infrastructure Identity service
---

# oci_identity_tag_namespace
This resource provides the Tag Namespace resource in Oracle Cloud Infrastructure Identity service.

Creates a new tag namespace in the specified compartment.

You must specify the compartment ID in the request object (remember that the tenancy is simply the root
compartment).

You must also specify a *name* for the namespace, which must be unique across all namespaces in your tenancy
and cannot be changed. The name can contain any ASCII character except the space (_) or period (.).
Names are case insensitive. That means, for example, "myNamespace" and "mynamespace" are not allowed
in the same tenancy. Once you created a namespace, you cannot change the name.
If you specify a name that's already in use in the tenancy, a 409 error is returned.

You must also specify a *description* for the namespace.
It does not have to be unique, and you can change it with
[UpdateTagNamespace](https://docs.cloud.oracle.com/iaas/api/#/en/identity/latest/TagNamespace/UpdateTagNamespace).

## Example Usage

```hcl
resource "oci_identity_tag_namespace" "test_tag_namespace" {
	#Required
	compartment_id = var.compartment_id
	description = var.tag_namespace_description
	name = var.tag_namespace_name

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	is_retired = false
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the tenancy containing the tag namespace.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Required) (Updatable) The description you assign to the tag namespace during creation.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `name` - (Required) The name you assign to the tag namespace during creation. It must be unique across all tag namespaces in the tenancy and cannot be changed. 
* `is_retired` - (Optional) (Updatable) Whether the tag namespace is retired. For more information, see [Retiring Key Definitions and Namespace Definitions](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/taggingoverview.htm#Retiring). 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the tag namespace.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description you assign to the tag namespace.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the tag namespace.
* `is_retired` - Whether the tag namespace is retired. For more information, see [Retiring Key Definitions and Namespace Definitions](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/taggingoverview.htm#Retiring). 
* `name` - The name of the tag namespace. It must be unique across all tag namespaces in the tenancy and cannot be changed. 
* `state` - The tagnamespace's current state. After creating a tagnamespace, make sure its `lifecycleState` is ACTIVE before using it. After retiring a tagnamespace, make sure its `lifecycleState` is INACTIVE before using it.
* `time_created` - Date and time the tag namespace was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Tag Namespace
	* `update` - (Defaults to 20 minutes), when updating the Tag Namespace
	* `delete` - (Defaults to 20 minutes), when destroying the Tag Namespace


## Import

TagNamespaces can be imported using the `id`, e.g.

```
$ terraform import oci_identity_tag_namespace.test_tag_namespace "id"
```

