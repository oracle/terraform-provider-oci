---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sensitive_types"
sidebar_current: "docs-oci-datasource-data_safe-sensitive_types"
description: |-
  Provides the list of Sensitive Types in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_sensitive_types
This data source provides the list of Sensitive Types in Oracle Cloud Infrastructure Data Safe service.

Gets a list of sensitive types based on the specified query parameters.


## Example Usage

```hcl
data "oci_data_safe_sensitive_types" "test_sensitive_types" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.sensitive_type_access_level
	compartment_id_in_subtree = var.sensitive_type_compartment_id_in_subtree
	default_masking_format_id = oci_data_safe_default_masking_format.test_default_masking_format.id
	display_name = var.sensitive_type_display_name
	entity_type = var.sensitive_type_entity_type
	is_common = var.sensitive_type_is_common
	parent_category_id = oci_marketplace_category.test_category.id
	sensitive_type_id = oci_data_safe_sensitive_type.test_sensitive_type.id
	sensitive_type_source = var.sensitive_type_sensitive_type_source
	state = var.sensitive_type_state
	time_created_greater_than_or_equal_to = var.sensitive_type_time_created_greater_than_or_equal_to
	time_created_less_than = var.sensitive_type_time_created_less_than
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `default_masking_format_id` - (Applicable when entity_type=SENSITIVE_TYPE) A filter to return only the sensitive types that have the default masking format identified by the specified OCID.
* `display_name` - (Optional) A filter to return only resources that match the specified display name. 
* `entity_type` - (Optional) A filter to return the sensitive type resources based on the value of their entityType attribute.
* `is_common` - (Optional) A filter to return only the common sensitive type resources. Common sensitive types belong to  library sensitive types which are frequently used to perform sensitive data discovery. 
* `parent_category_id` - (Optional) A filter to return only the sensitive types that are children of the sensitive category identified by the specified OCID.
* `sensitive_type_id` - (Optional) A filter to return only items related to a specific sensitive type OCID.
* `sensitive_type_source` - (Optional) A filter to return the sensitive type resources based on the value of their source attribute.
* `state` - (Optional) A filter to return only the resources that match the specified lifecycle state.
* `time_created_greater_than_or_equal_to` - (Optional) A filter to return only the resources that were created after the specified date and time, as defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Using TimeCreatedGreaterThanOrEqualToQueryParam parameter retrieves all resources created after that date.

	**Example:** 2016-12-19T16:39:57.600Z 
* `time_created_less_than` - (Optional) Search for resources that were created before a specific date. Specifying this parameter corresponding `timeCreatedLessThan` parameter will retrieve all resources created before the specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339.

	**Example:** 2016-12-19T16:39:57.600Z 


## Attributes Reference

The following attributes are exported:

* `sensitive_type_collection` - The list of sensitive_type_collection.

### SensitiveType Reference

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

