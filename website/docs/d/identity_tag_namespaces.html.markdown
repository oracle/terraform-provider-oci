---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_tag_namespaces"
sidebar_current: "docs-oci-datasource-identity-tag_namespaces"
description: |-
  Provides the list of Tag Namespaces in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_tag_namespaces
This data source provides the list of Tag Namespaces in Oracle Cloud Infrastructure Identity service.

Lists the tag namespaces in the specified compartment.


## Example Usage

```hcl
data "oci_identity_tag_namespaces" "test_tag_namespaces" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	include_subcompartments = "${var.tag_namespace_include_subcompartments}"
	state = "${var.tag_namespace_state}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment (remember that the tenancy is simply the root compartment). 
* `include_subcompartments` - (Optional) An optional boolean parameter indicating whether to retrieve all tag namespaces in subcompartments. If this parameter is not specified, only the tag namespaces defined in the specified compartment are retrieved. 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `tag_namespaces` - The list of tag_namespaces.

### TagNamespace Reference

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

