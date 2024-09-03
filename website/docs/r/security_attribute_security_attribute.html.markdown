---
subcategory: "Security Attribute"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_security_attribute_security_attribute"
sidebar_current: "docs-oci-resource-security_attribute-security_attribute"
description: |-
  Provides the Security Attribute resource in Oracle Cloud Infrastructure Security Attribute service
---

# oci_security_attribute_security_attribute
This resource provides the Security Attribute resource in Oracle Cloud Infrastructure Security Attribute service.

Creates a new security attribute in the specified security attribute namespace.

The security attribute requires either the OCID or the name of the security attribute namespace that will contain this
security attribute.

You must specify a *name* for the attribute, which must be unique across all attributes in the security attribute namespace
and cannot be changed. The only valid characters for security attribute names are: 0-9, A-Z, a-z, -, _ characters.
Names are case insensitive. That means, for example, "mySecurityAttribute" and "mysecurityattribute" are not allowed in the same namespace.
If you specify a name that's already in use in the security attribute namespace, a 409 error is returned.

The security attribute must have a *description*. It does not have to be unique, and you can change it with
[UpdateSecurityAttribute](https://docs.cloud.oracle.com/iaas/api/#/en/securityattribute/latest/Tag/UpdateSecurityAttribute).

When a validator is specified, The security attribute must have a value type. Security attribute can use either a static value or a list of possible values. Static values are entered by a user when applying the security attribute to a resource. Lists are created by the user and the user must apply a value from the list. Lists are validated.


## Example Usage

```hcl
resource "oci_security_attribute_security_attribute" "test_security_attribute" {
	#Required
	description = var.security_attribute_description
	name = var.security_attribute_name
	security_attribute_namespace_id = oci_security_attribute_security_attribute_namespace.test_security_attribute_namespace.id

	#Optional
	validator {
		#Required
		validator_type = var.security_attribute_validator_validator_type

		#Optional
		values = var.security_attribute_validator_values
	}
}
```

## Argument Reference

The following arguments are supported:

* `description` - (Required) (Updatable) The description you assign to the security attribute during creation.
* `name` - (Required) The name you assign to the security attribute during creation. This is the security attribute key. The name must be unique within the namespace and cannot be changed. 
* `security_attribute_namespace_id` - (Required) The OCID of the security attribute namespace. 
* `validator` - (Optional) (Updatable) Validates a security attribute value. Each validator performs validation steps in addition to the standard validation for security attribute values. For more information, see [Limits on Security Attributes](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/overview.htm).

	If you define a validator after a value has been set for a security attribute, then any updates that attempt to change the value must pass the additional validation defined by the current rule. Previously set values (even those that would fail the current validation) are not updated. You can still update other attributes to resources that contain a non-valid security attribute.

	To clear the validator call UpdateSecurityAttribute with [DefaultSecuirtyAttributeValidator](https://docs.cloud.oracle.com/iaas/api/#/en/securityattribute/latest/datatypes/DefaultTagDefinitionValidator). 
	* `validator_type` - (Required) (Updatable) Specifies the type of validation: a static value (no validation) or a list. 
	* `values` - (Applicable when validator_type=ENUM) (Updatable) The list of allowed values for a security attribute value. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the security attribute definition.
* `description` - The description you assign to the security attribute.
* `id` - The OCID of the security attribute definition.
* `is_retired` - Indicates whether the security attribute is retired. See [Managing Security Attribute Namespaces](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/managing-security-attribute-namespaces.htm). 
* `name` - The name assigned to the security attribute during creation. This is the security attribute key. The name must be unique within the security attribute namespace and cannot be changed. 
* `security_attribute_namespace_id` - The OCID of the security attribute namespace that contains the security attribute definition.
* `security_attribute_namespace_name` - The name of the security attribute namespace that contains the security attribute. 
* `state` - The security attribute's current state. After creating a security attribute, make sure its `lifecycleState` is ACTIVE before using it. After retiring a security attribute, make sure its `lifecycleState` is INACTIVE before using it. If you delete a security attribute, you cannot delete another security attribute until the deleted security attribute's `lifecycleState` changes from DELETING to DELETED.
* `time_created` - Date and time the security attribute was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `type` - The data type of the security attribute.
* `validator` - Validates a security attribute value. Each validator performs validation steps in addition to the standard validation for security attribute values. For more information, see [Limits on Security Attributes](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/overview.htm).

	If you define a validator after a value has been set for a security attribute, then any updates that attempt to change the value must pass the additional validation defined by the current rule. Previously set values (even those that would fail the current validation) are not updated. You can still update other attributes to resources that contain a non-valid security attribute.

	To clear the validator call UpdateSecurityAttribute with [DefaultSecuirtyAttributeValidator](https://docs.cloud.oracle.com/iaas/api/#/en/securityattribute/latest/datatypes/DefaultTagDefinitionValidator). 
	* `validator_type` - Specifies the type of validation: a static value (no validation) or a list. 
	* `values` - The list of allowed values for a security attribute value. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 15 minutes), when creating the Security Attribute
	* `update` - (Defaults to 15 minutes), when updating the Security Attribute
	* `delete` - (Defaults to 12 hours), when destroying the Security Attribute


## Import

SecurityAttributes can be imported using the `id`, e.g.

```
$ terraform import oci_security_attribute_security_attribute.test_security_attribute "securityAttributeNamespaces/{securityAttributeNamespaceId}/securityAttributes/{securityAttributeName}" 
```

