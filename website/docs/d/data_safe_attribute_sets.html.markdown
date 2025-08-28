---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_attribute_sets"
sidebar_current: "docs-oci-datasource-data_safe-attribute_sets"
description: |-
  Provides the list of Attribute Sets in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_attribute_sets
This data source provides the list of Attribute Sets in Oracle Cloud Infrastructure Data Safe service.

Retrieves the list of attribute sets.

The ListAttributeSets operation returns only the attribute sets in the specified `compartmentId`.

The parameter `accessLevel` specifies whether to return only those compartments for which the
requester has INSPECT permissions on at least one resource directly
or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
Principal doesn't have access to even one of the child compartments. This is valid only when
`compartmentIdInSubtree` is set to `true`.

The parameter `compartmentIdInSubtree` applies when you perform ListAttributeSet on the
`compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
To get a full list of all compartments and subcompartments in the tenancy (root compartment),
set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.


## Example Usage

```hcl
data "oci_data_safe_attribute_sets" "test_attribute_sets" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.attribute_set_access_level
	attribute_set_id = oci_data_safe_attribute_set.test_attribute_set.id
	attribute_set_type = var.attribute_set_attribute_set_type
	compartment_id_in_subtree = var.attribute_set_compartment_id_in_subtree
	display_name = var.attribute_set_display_name
	in_use = var.attribute_set_in_use
	is_user_defined = var.attribute_set_is_user_defined
	state = var.attribute_set_state
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `attribute_set_id` - (Optional) A filter to return only attribute set resources that matches the specified attribute set OCID query param.
* `attribute_set_type` - (Optional) A filter to return only attribute set resources that matches the specified attribute set type query param.
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `display_name` - (Optional) A filter to return only resources that match the specified display name. 
* `in_use` - (Optional) A filter to return attribute set resources that are in use by other associated resources.
* `is_user_defined` - (Optional) A filter to return user defined or seeded attribute set resources that matches the specified is user defined query param. A true value indicates user defined attribute set.
* `state` - (Optional) The current state of an attribute set.


## Attributes Reference

The following attributes are exported:

* `attribute_set_collection` - The list of attribute_set_collection.

### AttributeSet Reference

The following attributes are exported:

* `attribute_set_type` - The type of attribute set.
* `attribute_set_values` - The list of values in an attribute set
* `compartment_id` - The OCID of the compartment where the attribute set is stored.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
* `description` - Description of an attribute set.
* `display_name` - The display name of an attribute set. The name does not have to be unique, and is changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of an attribute set.
* `in_use` - Indicates whether the attribute set is in use by other resource.
* `is_user_defined` - A boolean flag indicating to list user defined or seeded attribute sets.
* `state` - The current state of an attribute set.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time an attribute set was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_updated` - The date and time an attribute set was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

