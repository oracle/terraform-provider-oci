---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_tag_defaults"
sidebar_current: "docs-oci-datasource-identity-tag_defaults"
description: |-
  Provides the list of Tag Defaults in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_tag_defaults
This data source provides the list of Tag Defaults in Oracle Cloud Infrastructure Identity service.

Lists the Tag Defaults for Tag Definitions in the specified Compartment.


## Example Usage

```hcl
data "oci_identity_tag_defaults" "test_tag_defaults" {

	#Optional
	compartment_id = "${var.compartment_id}"
	id = "${var.tag_default_id}"
	state = "${var.tag_default_state}"
	tag_definition_id = "${oci_identity_tag_definition.test_tag_definition.id}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment (remember that the tenancy is simply the root compartment). 
* `id` - (Optional) A filter to only return resources that match the specified OCID exactly. 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive. 
* `tag_definition_id` - (Optional) The OCID of the Tag Definition. 


## Attributes Reference

The following attributes are exported:

* `tag_defaults` - The list of tag_defaults.

### TagDefault Reference

The following attributes are exported:

* `compartment_id` - The OCID of the Compartment. The Tag Default will apply to any resource contained in this Compartment. 
* `id` - The OCID of the Tag Default.
* `state` - The tag default's current state. After creating a tagdefault, make sure its `lifecycleState` is ACTIVE before using it. 
* `tag_definition_id` - The OCID of the Tag Definition. The Tag Default will always assign a default value for this Tag Definition. 
* `tag_definition_name` - The name used in the Tag Definition. This field is informational in the context of the Tag Default. 
* `tag_namespace_id` - The OCID of the Tag Namespace that contains the Tag Definition. 
* `time_created` - Date and time the `TagDefault` object was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `value` - The default value for the Tag Definition. This will be applied to all resources created in the Compartment. 

