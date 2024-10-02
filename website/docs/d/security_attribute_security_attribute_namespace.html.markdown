---
subcategory: "Security Attribute"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_security_attribute_security_attribute_namespace"
sidebar_current: "docs-oci-datasource-security_attribute-security_attribute_namespace"
description: |-
  Provides details about a specific Security Attribute Namespace in Oracle Cloud Infrastructure Security Attribute service
---

# Data Source: oci_security_attribute_security_attribute_namespace
This data source provides details about a specific Security Attribute Namespace resource in Oracle Cloud Infrastructure Security Attribute service.

Gets the specified security attribute namespace's information.


## Example Usage

```hcl
data "oci_security_attribute_security_attribute_namespace" "test_security_attribute_namespace" {
	#Required
	security_attribute_namespace_id = oci_security_attribute_security_attribute_namespace.test_security_attribute_namespace.id
}
```

## Argument Reference

The following arguments are supported:

* `security_attribute_namespace_id` - (Required) The OCID of the security attribute namespace. 


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the security attribute namespace.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - Description of the Security Attribute Namespace.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the security attribute namespace.
* `is_retired` - Indicates whether the security attribute namespace is retired. See [Managing Security Attribute Namespaces](https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/managing-security-attribute-namespaces.htm). 
* `mode` - Indicates possible modes the security attributes in this namespace is set to. Supported values are `enforce` and `audit`. Currently mode cannot be controlled by the user 
* `name` - The name of the security attribute namespace. It must be unique across all security attribute namespaces in the tenancy and cannot be changed. 
* `state` - The security attribute namespace's current state. After creating a security attribute namespace, `lifecycleState` is in ACTIVE state. After retiring a security attribute namespace, its `lifecycleState` becomes INACTIVE. Security Attributes from a retired namespace cannot be attached to more resources.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Date and time the security attribute namespace was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z` 

