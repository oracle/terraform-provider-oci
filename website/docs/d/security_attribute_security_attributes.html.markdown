---
subcategory: "Security Attribute"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_security_attribute_security_attributes"
sidebar_current: "docs-oci-datasource-security_attribute-security_attributes"
description: |-
  Provides the list of Security Attributes in Oracle Cloud Infrastructure Security Attribute service
---

# Data Source: oci_security_attribute_security_attributes
This data source provides the list of Security Attributes in Oracle Cloud Infrastructure Security Attribute service.

Lists the security attributes in the specified namespace.


## Example Usage

```hcl
data "oci_security_attribute_security_attributes" "test_security_attributes" {
	#Required
	security_attribute_namespace_id = oci_security_attribute_security_attribute_namespace.test_security_attribute_namespace.id

	#Optional
	state = var.security_attribute_state
}
```

## Argument Reference

The following arguments are supported:

* `security_attribute_namespace_id` - (Required) The OCID of the security attribute namespace. 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `security_attributes` - The list of security_attributes.

### SecurityAttribute Reference

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

