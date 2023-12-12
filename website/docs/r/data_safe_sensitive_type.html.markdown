---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sensitive_type"
sidebar_current: "docs-oci-resource-data_safe-sensitive_type"
description: |-
  Provides the Sensitive Type resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_sensitive_type
This resource provides the Sensitive Type resource in Oracle Cloud Infrastructure Data Safe service.

Creates a new sensitive type, which can be a basic sensitive type with regular expressions or a sensitive category.
While sensitive types are used for data discovery, sensitive categories are used for logically grouping the related
or similar sensitive types.


## Example Usage

```hcl
resource "oci_data_safe_sensitive_type" "test_sensitive_type" {
	#Required
	compartment_id = var.compartment_id
	entity_type = var.sensitive_type_entity_type

	#Optional
	comment_pattern = var.sensitive_type_comment_pattern
	data_pattern = var.sensitive_type_data_pattern
	default_masking_format_id = oci_data_safe_default_masking_format.test_default_masking_format.id
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.sensitive_type_description
	display_name = var.sensitive_type_display_name
	freeform_tags = {"Department"= "Finance"}
	name_pattern = var.sensitive_type_name_pattern
	parent_category_id = oci_marketplace_category.test_category.id
	search_type = var.sensitive_type_search_type
	short_name = var.sensitive_type_short_name
}
```

## Argument Reference

The following arguments are supported:

* `comment_pattern` - (Applicable when entity_type=SENSITIVE_TYPE) (Updatable) A regular expression to be used by data discovery for matching column comments.
* `compartment_id` - (Required) (Updatable) The OCID of the compartment where the sensitive type should be created.
* `data_pattern` - (Applicable when entity_type=SENSITIVE_TYPE) (Updatable) A regular expression to be used by data discovery for matching column data values.
* `default_masking_format_id` - (Applicable when entity_type=SENSITIVE_TYPE) (Updatable) The OCID of the library masking format that should be used to mask the sensitive columns associated with the sensitive type.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) The description of the sensitive type.
* `display_name` - (Optional) (Updatable) The display name of the sensitive type. The name does not have to be unique, and it's changeable.
* `entity_type` - (Required) (Updatable) The entity type. It can be either a sensitive type with regular expressions or a sensitive category used for grouping similar sensitive types. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `name_pattern` - (Applicable when entity_type=SENSITIVE_TYPE) (Updatable) A regular expression to be used by data discovery for matching column names.
* `parent_category_id` - (Optional) (Updatable) The OCID of the parent sensitive category.
* `search_type` - (Applicable when entity_type=SENSITIVE_TYPE) (Updatable) The search type indicating how the column name, comment and data patterns should be used by data discovery. [Learn more](https://docs.oracle.com/en/cloud/paas/data-safe/udscs/sensitive-types.html#GUID-1D1AD98E-B93F-4FF2-80AE-CB7D8A14F6CC). 
* `short_name` - (Optional) (Updatable) The short name of the sensitive type.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `comment_pattern` - A regular expression to be used by data discovery for matching column comments.
* `compartment_id` - The OCID of the compartment that contains the sensitive type.
* `data_pattern` - A regular expression to be used by data discovery for matching column data values.
* `default_masking_format_id` - The OCID of the library masking format that should be used to mask the sensitive columns associated with the sensitive type.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the sensitive type.
* `display_name` - The display name of the sensitive type.
* `entity_type` - The entity type. It can be either a sensitive type with regular expressions or a sensitive category used for grouping similar sensitive types. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the sensitive type.
* `is_common` - Specifies whether the sensitive type is common. Common sensitive types belong to  library sensitive types which are frequently used to perform sensitive data discovery. 
* `name_pattern` - A regular expression to be used by data discovery for matching column names.
* `parent_category_id` - The OCID of the parent sensitive category.
* `search_type` - The search type indicating how the column name, comment and data patterns should be used by data discovery. [Learn more](https://docs.oracle.com/en/cloud/paas/data-safe/udscs/sensitive-types.html#GUID-1D1AD98E-B93F-4FF2-80AE-CB7D8A14F6CC). 
* `short_name` - The short name of the sensitive type.
* `source` - Specifies whether the sensitive type is user-defined or predefined.
* `state` - The current state of the sensitive type.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the sensitive type was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_updated` - The date and time the sensitive type was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Sensitive Type
	* `update` - (Defaults to 20 minutes), when updating the Sensitive Type
	* `delete` - (Defaults to 20 minutes), when destroying the Sensitive Type


## Import

SensitiveTypes can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_sensitive_type.test_sensitive_type "id"
```

