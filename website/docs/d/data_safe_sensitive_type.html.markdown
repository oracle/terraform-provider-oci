---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sensitive_type"
sidebar_current: "docs-oci-datasource-data_safe-sensitive_type"
description: |-
  Provides details about a specific Sensitive Type in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_sensitive_type
This data source provides details about a specific Sensitive Type resource in Oracle Cloud Infrastructure Data Safe service.

Gets the details of the specified sensitive type.

## Example Usage

```hcl
data "oci_data_safe_sensitive_type" "test_sensitive_type" {
	#Required
	sensitive_type_id = oci_data_safe_sensitive_type.test_sensitive_type.id
}
```

## Argument Reference

The following arguments are supported:

* `sensitive_type_id` - (Required) The OCID of the sensitive type.


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

