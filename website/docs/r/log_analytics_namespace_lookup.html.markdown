---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_namespace_lookup"
sidebar_current: "docs-oci-resource-log_analytics-namespace_lookup"
description: |-
  Provides the Namespace Lookup resource in Oracle Cloud Infrastructure Log Analytics service
---

# oci_log_analytics_namespace_lookup
This resource provides the Namespace Lookup resource in Oracle Cloud Infrastructure Log Analytics service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/logan-api-spec/latest/NamespaceLookup

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/log_analytics

Creates a lookup with the specified name, type and description. The csv file containing the lookup content is passed in as binary data in the request.


## Example Usage

```hcl
resource "oci_log_analytics_namespace_lookup" "test_namespace_lookup" {
	#Required
	lookup_name = var.namespace_lookup_lookup_name
	namespace = var.namespace_lookup_namespace
	register_lookup_file = var.namespace_lookup_register_lookup_file
	type = var.namespace_lookup_type

	#Optional
	categories {

		#Optional
		description = var.namespace_lookup_categories_description
		display_name = var.namespace_lookup_categories_display_name
		is_system = var.namespace_lookup_categories_is_system
		name = var.namespace_lookup_categories_name
		type = var.namespace_lookup_categories_type
	}
	char_encoding = var.namespace_lookup_char_encoding
	compartment_id = var.namespace_lookup_compartment_id
	default_match_value = var.namespace_lookup_default_match_value
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.namespace_lookup_description
	fields {

		#Optional
		common_field_name = var.namespace_lookup_fields_common_field_name
		default_match_value = var.namespace_lookup_fields_default_match_value
		display_name = var.namespace_lookup_fields_display_name
		is_common_field = var.namespace_lookup_fields_is_common_field
		match_operator = var.namespace_lookup_fields_match_operator
		name = var.namespace_lookup_fields_name
		position = var.namespace_lookup_fields_position
	}
	freeform_tags = {"bar-key"= "value"}
	is_hidden = var.namespace_lookup_is_hidden
	max_matches = var.namespace_lookup_max_matches
}
```

## Argument Reference

The following arguments are supported:

* `categories` - (Optional) (Updatable) An array of categories to assign to the lookup. Specifying the name attribute for each category would suffice. Oracle-defined category assignments cannot be removed. 
    * `description` - (Optional) (Updatable) The category description.
    * `display_name` - (Optional) (Updatable) The category display name.
    * `is_system` - (Optional) (Updatable) The system flag. A value of false denotes a user-created category. A value of true denotes an Oracle-defined category. 
    * `name` - (Optional) (Updatable) The unique name that identifies the category.
    * `type` - (Optional) (Updatable) The category type. Values include "PRODUCT", "TIER", "VENDOR" and "GENERIC".
* `char_encoding` - (Optional) The character encoding of the uploaded file.
* `compartment_id` - (Optional) (Updatable) Compartment Identifier [OCID] (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `default_match_value` - (Optional) (Updatable) The default match value.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) The lookup description.
* `fields` - (Optional) (Updatable) The lookup fields.
    * `common_field_name` - (Optional) (Updatable) The common field name.
    * `default_match_value` - (Optional) (Updatable) The default match value.
    * `display_name` - (Optional) (Updatable) The display name.
    * `is_common_field` - (Optional) (Updatable) A flag indicating whether or not the field is a common field. 
    * `match_operator` - (Optional) (Updatable) The match operator.
    * `name` - (Optional) (Updatable) The field name.
    * `position` - (Optional) (Updatable) The position.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
* `is_hidden` - (Optional) A flag indicating whether or not the new lookup should be hidden.
* `lookup_name` - (Required) The name of the lookup to operate on.
* `max_matches` - (Optional) (Updatable) The maximum number of matches.
* `namespace` - (Required) The Logging Analytics namespace used for the request.
* `register_lookup_file` - (Required) Path to the file containing data for lookup creation.
* `type` - (Required) The lookup type. Valid values are Lookup, Dictionary or Module.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `active_edit_version` - The active edit version.
* `canonical_link` - The canonical link.
* `categories` - An array of categories assigned to this lookup. The isSystem flag denotes if each category assignment is user-created or Oracle-defined. 
    * `description` - The category description.
    * `display_name` - The category display name.
    * `is_system` - The system flag. A value of false denotes a user-created category. A value of true denotes an Oracle-defined category. 
    * `name` - The unique name that identifies the category.
    * `type` - The category type. Values include "PRODUCT", "TIER", "VENDOR" and "GENERIC".
* `compartment_id` - Compartment Identifier [OCID] (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - The lookup description.
* `edit_version` - The edit version.
* `fields` - The lookup fields.
    * `common_field_name` - The common field name.
    * `default_match_value` - The default match value.
    * `display_name` - The field display name.
    * `is_common_field` - A flag indicating whether or not the lookup field is a common field. 
    * `match_operator` - The match operator.
    * `name` - The field name.
    * `position` - THe field position.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
* `id` - A unique string that identifies this lookup in terraform.
* `is_built_in` - A flag indicating if the lookup is custom (user-defined) or built in. 
* `is_hidden` - A flag indicating if the lookup is hidden or not.  A hidden lookup will not be returned in list operations by default. 
* `lookup_display_name` - The lookup display name.
* `lookup_id` - The lookup OCID.
* `lookup_name` - The lookup name.
* `lookup_reference` - The lookup reference as an integer.
* `lookup_reference_string` - The lookup reference as a string.
* `referring_sources` - AutoLookups
    * `canonical_link` - The canonical link.
    * `total_count` - The total count.
* `status_summary` - StatusSummary
    * `chunks_processed` - The number of chunks processed.
    * `failure_details` - The failure details, if any.
    * `filename` - The filename.
    * `status` - The status.
    * `total_chunks` - The total number of chunks.
* `time_updated` - The last updated date.
* `type` - The lookup type. Valid values are Lookup, Dictionary or Module.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Namespace Lookup
	* `update` - (Defaults to 20 minutes), when updating the Namespace Lookup
	* `delete` - (Defaults to 20 minutes), when destroying the Namespace Lookup


## Import

NamespaceLookups can be imported using the `id`, e.g.

```
$ terraform import oci_log_analytics_namespace_lookup.test_namespace_lookup "namespaces/{namespaceName}/lookups/{lookupName}" 
```

