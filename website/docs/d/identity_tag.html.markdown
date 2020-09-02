---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_tag"
sidebar_current: "docs-oci-datasource-identity-tag"
description: |-
  Provides details about a specific Tag in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_tag
This data source provides details about a specific Tag resource in Oracle Cloud Infrastructure Identity service.

Gets the specified tag's information.

## Example Usage

```hcl
data "oci_identity_tag" "test_tag" {
	#Required
	tag_name = "${oci_identity_tag.test_tag.name}"
	tag_namespace_id = "${oci_identity_tag_namespace.test_tag_namespace.id}"
}
```

## Argument Reference

The following arguments are supported:

* `tag_name` - (Required) The name of the tag. 
* `tag_namespace_id` - (Required) The OCID of the tag namespace. 


## Attributes Reference

The following attributes are exported:

* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description you assign to the tag.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the tag definition.
* `is_cost_tracking` - Indicates whether the tag is enabled for cost tracking. 
* `is_retired` - Indicates whether the tag is retired. See [Retiring Key Definitions and Namespace Definitions](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/taggingoverview.htm#Retiring). 
* `name` - The name assigned to the tag during creation. This is the tag key definition. The name must be unique within the tag namespace and cannot be changed. 
* `state` - The tag's current state. After creating a tag, make sure its `lifecycleState` is ACTIVE before using it. After retiring a tag, make sure its `lifecycleState` is INACTIVE before using it. If you delete a tag, you cannot delete another tag until the deleted tag's `lifecycleState` changes from DELETING to DELETED.
* `tag_namespace_id` - The OCID of the namespace that contains the tag definition.
* `time_created` - Date and time the tag was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `validator` - 
	* `validator_type` - Specifies the type of validation: a static value (no validation) or a list.  
	* `values` - The list of allowed values for a definedTag value. 

