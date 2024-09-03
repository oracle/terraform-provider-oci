---
subcategory: "Security Attribute"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_security_attribute_security_attribute_namespaces"
sidebar_current: "docs-oci-datasource-security_attribute-security_attribute_namespaces"
description: |-
  Provides the list of Security Attribute Namespaces in Oracle Cloud Infrastructure Security Attribute service
---

# Data Source: oci_security_attribute_security_attribute_namespaces
This data source provides the list of Security Attribute Namespaces in Oracle Cloud Infrastructure Security Attribute service.

Lists the security attribute namespaces in the specified compartment.


## Example Usage

```hcl
data "oci_security_attribute_security_attribute_namespaces" "test_security_attribute_namespaces" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	compartment_id_in_subtree = var.security_attribute_namespace_compartment_id_in_subtree
	name = var.security_attribute_namespace_name
	state = var.security_attribute_namespace_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `compartment_id_in_subtree` - (Optional) An optional boolean parameter indicating whether to retrieve all security attribute namespaces in subcompartments. If this parameter is not specified, only the namespaces defined in the specified compartment are retrieved. 
* `name` - (Optional) A filter to return only resources that match the entire display name given.
* `state` - (Optional) A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `security_attribute_namespaces` - The list of security_attribute_namespaces.

### SecurityAttributeNamespace Reference

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

