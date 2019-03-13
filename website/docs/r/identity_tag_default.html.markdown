---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_tag_default"
sidebar_current: "docs-oci-resource-identity-tag_default"
description: |-
  Provides the Tag Default resource in Oracle Cloud Infrastructure Identity service
---

# oci_identity_tag_default
This resource provides the Tag Default resource in Oracle Cloud Infrastructure Identity service.

Creates a new Tag Default in the specified Compartment for the specified Tag Definition.


## Example Usage

```hcl
resource "oci_identity_tag_default" "test_tag_default" {
	#Required
	compartment_id = "${var.compartment_id}"
	tag_definition_id = "${oci_identity_tag_definition.test_tag_definition.id}"
	value = "${var.tag_default_value}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the Compartment. The Tag Default will apply to any resource contained in this Compartment. 
* `tag_definition_id` - (Required) The OCID of the Tag Definition. The Tag Default will always assign a default value for this Tag Definition. 
* `value` - (Required) (Updatable) The default value for the Tag Definition. This will be applied to all resources created in the Compartment. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the Compartment. The Tag Default will apply to any resource contained in this Compartment. 
* `id` - The OCID of the Tag Default.
* `state` - The tag default's current state. After creating a tagdefault, make sure its `lifecycleState` is ACTIVE before using it. 
* `tag_definition_id` - The OCID of the Tag Definition. The Tag Default will always assign a default value for this Tag Definition. 
* `tag_definition_name` - The name used in the Tag Definition. This field is informational in the context of the Tag Default. 
* `tag_namespace_id` - The OCID of the Tag Namespace that contains the Tag Definition. 
* `time_created` - Date and time the `TagDefault` object was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `value` - The default value for the Tag Definition. This will be applied to all resources created in the Compartment. 

## Import

TagDefaults can be imported using the `id`, e.g.

```
$ terraform import oci_identity_tag_default.test_tag_default "id"
```

