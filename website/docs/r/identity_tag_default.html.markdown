---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_tag_default"
sidebar_current: "docs-oci-resource-identity-tag_default"
description: |-
  Provides the Tag Default resource in Oracle Cloud Infrastructure Identity service
---

# oci_identity_tag_default
This resource provides the Tag Default resource in Oracle Cloud Infrastructure Identity service.

Creates a new tag default in the specified compartment for the specified tag definition.


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

* `compartment_id` - (Required) The OCID of the compartment. The tag default will be applied to all new resources created in this compartment. 
* `tag_definition_id` - (Required) The OCID of the tag definition. The tag default will always assign a default value for this tag definition. 
* `value` - (Required) (Updatable) The default value for the tag definition. This will be applied to all new resources created in the compartment. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment. The tag default applies to all new resources that get created in the compartment. Resources that existed before the tag default was created are not tagged. 
* `id` - The OCID of the tag default.
* `state` - The tag default's current state. After creating a `TagDefault`, make sure its `lifecycleState` is ACTIVE before using it. 
* `tag_definition_id` - The OCID of the tag definition. The tag default will always assign a default value for this tag definition. 
* `tag_definition_name` - The name used in the tag definition. This field is informational in the context of the tag default. 
* `tag_namespace_id` - The OCID of the tag namespace that contains the tag definition. 
* `time_created` - Date and time the `TagDefault` object was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `value` - The default value for the tag definition. This will be applied to all new resources created in the compartment. 

## Import

TagDefaults can be imported using the `id`, e.g.

```
$ terraform import oci_identity_tag_default.test_tag_default "id"
```

