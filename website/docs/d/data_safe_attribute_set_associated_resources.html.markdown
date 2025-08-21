---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_attribute_set_associated_resources"
sidebar_current: "docs-oci-datasource-data_safe-attribute_set_associated_resources"
description: |-
  Provides the list of Attribute Set Associated Resources in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_attribute_set_associated_resources
This data source provides the list of Attribute Set Associated Resources in Oracle Cloud Infrastructure Data Safe service.

Returns list of all associated resources.

## Example Usage

```hcl
data "oci_data_safe_attribute_set_associated_resources" "test_attribute_set_associated_resources" {
	#Required
	attribute_set_id = oci_data_safe_attribute_set.test_attribute_set.id

	#Optional
	associated_resource_id = oci_cloud_guard_resource.test_resource.id
	associated_resource_type = var.attribute_set_associated_resource_associated_resource_type
}
```

## Argument Reference

The following arguments are supported:

* `associated_resource_id` - (Optional) A filter to return attribute set associated resource that matches the specified associated resource id query param.
* `associated_resource_type` - (Optional) A filter to return attribute set associated resources that matches the specified resource type query param.
* `attribute_set_id` - (Required) OCID of an attribute set.


## Attributes Reference

The following attributes are exported:

* `associated_resource_collection` - The list of associated_resource_collection.

### AttributeSetAssociatedResource Reference

The following attributes are exported:

* `items` - Array of associated resources.
	* `associated_resource_id` - The OCID of the resource that is associated with the attribute set.
	* `associated_resource_name` - The display name of the resource that is associated with the attribute set. The name does not have to be unique, and is changeable.
	* `associated_resource_type` - The resource type that is associated with the attribute set.
	* `time_created` - The date and time when associated started between resource and the attribute set, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
	* `time_updated` - The date and time when associated is removed between resources and the attribute set, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

