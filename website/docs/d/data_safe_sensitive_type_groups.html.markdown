---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sensitive_type_groups"
sidebar_current: "docs-oci-datasource-data_safe-sensitive_type_groups"
description: |-
  Provides the list of Sensitive Type Groups in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_sensitive_type_groups
This data source provides the list of Sensitive Type Groups in Oracle Cloud Infrastructure Data Safe service.

Gets a list of sensitive type groups based on the specified query parameters.


## Example Usage

```hcl
data "oci_data_safe_sensitive_type_groups" "test_sensitive_type_groups" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.sensitive_type_group_access_level
	compartment_id_in_subtree = var.sensitive_type_group_compartment_id_in_subtree
	display_name = var.sensitive_type_group_display_name
	sensitive_type_group_id = oci_data_safe_sensitive_type_group.test_sensitive_type_group.id
	state = var.sensitive_type_group_state
	time_created_greater_than_or_equal_to = var.sensitive_type_group_time_created_greater_than_or_equal_to
	time_created_less_than = var.sensitive_type_group_time_created_less_than
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `display_name` - (Optional) A filter to return only resources that match the specified display name. 
* `sensitive_type_group_id` - (Optional) An optional filter to return only resources that match the specified OCID of the sensitive type group resource.
* `state` - (Optional) A filter to return only the resources that match the specified lifecycle state.
* `time_created_greater_than_or_equal_to` - (Optional) A filter to return only the resources that were created after the specified date and time, as defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Using TimeCreatedGreaterThanOrEqualToQueryParam parameter retrieves all resources created after that date.

	**Example:** 2016-12-19T16:39:57.600Z 
* `time_created_less_than` - (Optional) Search for resources that were created before a specific date. Specifying this parameter corresponding `timeCreatedLessThan` parameter will retrieve all resources created before the specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339.

	**Example:** 2016-12-19T16:39:57.600Z 


## Attributes Reference

The following attributes are exported:

* `sensitive_type_group_collection` - The list of sensitive_type_group_collection.

### SensitiveTypeGroup Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the sensitive type group.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the sensitive type group.
* `display_name` - The display name of the sensitive type group.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the sensitive type group.
* `sensitive_type_count` - The number of sensitive types in the specified sensitive type group.
* `state` - The current state of the sensitive type group.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the sensitive type group was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_updated` - The date and time the sensitive type group was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

