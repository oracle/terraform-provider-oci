---
layout: "oci"
page_title: "OCI: oci_identity_tag_namespace"
sidebar_current: "docs-oci-resource-identity-tag_namespace"
description: |-
  Creates and manages an OCI TagNamespace
---

# oci_identity_tag_namespace
The `oci_identity_tag_namespace` resource creates and manages an OCI TagNamespace

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
[UpdateTagNamespace](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/tagging/20170101/TagNamespace/UpdateTagNamespace).

Tag namespaces cannot be deleted, but they can be retired.
See [Retiring Key Definitions and Namespace Definitions](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/taggingoverview.htm#Retiring) for more information.


## Example Usage

```hcl
resource "oci_identity_tag_namespace" "test_tag_namespace" {
	#Required
	compartment_id = "${var.compartment_id}"
	description = "${var.tag_namespace_description}"
	name = "${var.tag_namespace_name}"

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	is_retired = false
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the tenancy containing the tag namespace.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Required) (Updatable) The description you assign to the tag namespace during creation.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `name` - (Required) The name you assign to the tag namespace during creation. It must be unique across all tag namespaces in the tenancy and cannot be changed. 
* `is_retired` - (Optional) (Updatable) Whether the tag namespace is retired. For more information, see [Retiring Key Definitions and Namespace Definitions](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/taggingoverview.htm#Retiring). 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the tag namespace.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description you assign to the tag namespace.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the tag namespace.
* `is_retired` - Whether the tag namespace is retired. For more information, see [Retiring Key Definitions and Namespace Definitions](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/taggingoverview.htm#Retiring). 
* `name` - The name of the tag namespace. It must be unique across all tag namespaces in the tenancy and cannot be changed. 
* `time_created` - Date and time the tagNamespace was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z` 

## Import

TagNamespaces can be imported using the `id`, e.g.

```
$ terraform import oci_identity_tag_namespace.test_tag_namespace "id"
```
