---
subcategory: "Identity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_tag_standard_tag_namespace_template"
sidebar_current: "docs-oci-datasource-identity-tag_standard_tag_namespace_template"
description: |-
  Provides details about a specific Tag Standard Tag Namespace Template in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_tag_standard_tag_namespace_template
This data source provides details about a specific Tag Standard Tag Namespace Template resource in Oracle Cloud Infrastructure Identity service.

Retrieve the standard tag namespace template given the standard tag namespace name.


## Example Usage

```hcl
data "oci_identity_tag_standard_tag_namespace_template" "test_tag_standard_tag_namespace_template" {
	#Required
	compartment_id = var.compartment_id
	standard_tag_namespace_name = oci_identity_tag_namespace.test_tag_namespace.name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment (remember that the tenancy is simply the root compartment). 
* `standard_tag_namespace_name` - (Required) The name of the standard tag namespace tempate that is requested 


## Attributes Reference

The following attributes are exported:

* `description` - The default description of the tag namespace that users can use to create the tag namespace
* `standard_tag_namespace_name` - The reserved name of this standard tag namespace
* `status` - The status of the standard tag namespace
* `tag_definition_templates` - The template of the tag definition. This object includes necessary details to create the provided standard tag definition.
	* `description` - The default description of the tag namespace that users can use to create the tag definition
	* `enum_mutability` - The mutability of the possible values list for enum tags. This will default to IMMUTABLE for string value tags
	* `is_cost_tracking` - Is the tag a cost tracking tag. Default will be false as cost tracking tags have been deprecated
	* `possible_values` - List of possible values. An optional parameter that will be present if the type of definition is enum.
	* `tag_definition_name` - The name of this standard tag definition
	* `type` - The type of tag definition. Enum or string.

