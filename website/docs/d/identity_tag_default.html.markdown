---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_tag_default"
sidebar_current: "docs-oci-datasource-identity-tag_default"
description: |-
  Provides details about a specific Tag Default in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_tag_default
This data source provides details about a specific Tag Default resource in Oracle Cloud Infrastructure Identity service.

Retrieves the specified Tag Default.


## Example Usage

```hcl
data "oci_identity_tag_default" "test_tag_default" {
	#Required
	tag_default_id = "${oci_identity_tag_default.test_tag_default.id}"
}
```

## Argument Reference

The following arguments are supported:

* `tag_default_id` - (Required) The OCID of the Tag Default.


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

