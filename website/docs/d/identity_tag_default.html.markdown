---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_tag_default"
sidebar_current: "docs-oci-datasource-identity-tag_default"
description: |-
  Provides details about a specific Tag Default in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_tag_default
This data source provides details about a specific Tag Default resource in Oracle Cloud Infrastructure Identity service.

Retrieves the specified tag default.


## Example Usage

```hcl
data "oci_identity_tag_default" "test_tag_default" {
	#Required
	tag_default_id = "${oci_identity_tag_default.test_tag_default.id}"
}
```

## Argument Reference

The following arguments are supported:

* `tag_default_id` - (Required) The OCID of the tag default.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment. The tag default applies to all new resources that get created in the compartment. Resources that existed before the tag default was created are not tagged. 
* `id` - The OCID of the tag default.
* `is_required` - If you specify that a value is required, a value is set during resource creation (either by the  user creating the resource or another tag defualt). If no value is set, resource creation is  blocked.
	* If the `isRequired` flag is set to "true", the value is set during resource creation.
	* If the `isRequired` flag is set to "false", the value you enter is set during resource creation.

	Example: `false` 
* `state` - The tag default's current state. After creating a `TagDefault`, make sure its `lifecycleState` is ACTIVE before using it. 
* `tag_definition_id` - The OCID of the tag definition. The tag default will always assign a default value for this tag definition. 
* `tag_definition_name` - The name used in the tag definition. This field is informational in the context of the tag default. 
* `tag_namespace_id` - The OCID of the tag namespace that contains the tag definition. 
* `time_created` - Date and time the `TagDefault` object was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `value` - The default value for the tag definition. This will be applied to all new resources created in the compartment. 

