---
subcategory: "Security Attribute"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_security_attribute_security_attribute_namespace"
sidebar_current: "docs-oci-resource-security_attribute-security_attribute_namespace"
description: |-
  Provides the Security Attribute Namespace resource in Oracle Cloud Infrastructure Security Attribute service
---

# oci_security_attribute_security_attribute_namespace
This resource provides the Security Attribute Namespace resource in Oracle Cloud Infrastructure Security Attribute service.

Creates a new security attribute namespace in the specified compartment.

You must specify the compartment ID in the request object (remember that the tenancy is simply the root
compartment).

You must also specify a *name* for the namespace, which must be unique across all namespaces in your tenancy
and cannot be changed. The only valid characters for security attribute names are:  0-9, A-Z, a-z, -, _ characters.
Names are case insensitive. That means, for example, "myNamespace" and "mynamespace" are not allowed
in the same tenancy. Once you created a namespace, you cannot change the name.
If you specify a name that's already in use in the tenancy, a 409 error is returned.

You must also specify a *description* for the namespace.
It does not have to be unique, and you can change it with
[UpdateSecurityAttributeNamespace](https://docs.cloud.oracle.com/iaas/api/#/en/securityattribute/latest/SecurityAttribute/SecurityAttributeNamespace).


## Example Usage

```hcl
resource "oci_security_attribute_security_attribute_namespace" "test_security_attribute_namespace" {
	#Required
	compartment_id = var.compartment_id
	description = var.security_attribute_namespace_description
	name = var.security_attribute_namespace_name

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the tenancy containing the security attribute namespace.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Required) (Updatable) The description you assign to the security attribute namespace during creation.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `name` - (Required) The name you assign to the security attribute namespace during creation. The name must be unique across all namespaces in the tenancy and cannot be changed. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the security attribute namespace.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - A description you create for the security attribute namespace to help you identify it.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the security attribute namespace.
* `is_retired` - Indicates whether the security attribute namespace is retired. See [Managing Security Attribute Namespaces](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/managing-security-attribute-namespaces.htm). 
* `mode` - Indicates possible modes the security attributes in this namespace can be set to. This is not accepted from the user. Currently the supported values are enforce and audit. 
* `name` - The name of the security attribute namespace. It must be unique across all security attribute namespaces in the tenancy and cannot be changed. 
* `state` - The security attribute namespace's current state. After creating a security attribute namespace, make sure its `lifecycleState` is ACTIVE before using it. After retiring a security attribute namespace, make sure its `lifecycleState` is INACTIVE.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Date and time the security attribute namespace was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Security Attribute Namespace
	* `update` - (Defaults to 20 minutes), when updating the Security Attribute Namespace
	* `delete` - (Defaults to 20 minutes), when destroying the Security Attribute Namespace


## Import

SecurityAttributeNamespaces can be imported using the `id`, e.g.

```
$ terraform import oci_security_attribute_security_attribute_namespace.test_security_attribute_namespace "id"
```

