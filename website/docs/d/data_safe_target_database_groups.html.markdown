---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_target_database_groups"
sidebar_current: "docs-oci-datasource-data_safe-target_database_groups"
description: |-
  Provides the list of Target Database Groups in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_target_database_groups
This data source provides the list of Target Database Groups in Oracle Cloud Infrastructure Data Safe service.

Retrieves a list of target database groups according to the specified query parameters.


## Example Usage

```hcl
data "oci_data_safe_target_database_groups" "test_target_database_groups" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.target_database_group_access_level
	compartment_id_in_subtree = var.target_database_group_compartment_id_in_subtree
	display_name = var.target_database_group_display_name
	state = var.target_database_group_state
	target_database_group_id = oci_data_safe_target_database_group.test_target_database_group.id
	target_database_group_filter = var.target_database_group_target_database_group_filter
	time_created_greater_than_or_equal_to = var.target_database_group_time_created_greater_than_or_equal_to
	time_created_less_than = var.target_database_group_time_created_less_than
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `display_name` - (Optional) A filter to return only resources that match the specified display name. 
* `state` - (Optional) A filter to retrieve resources that exclusively align with the designated lifecycle state.
* `target_database_group_id` - (Optional) A filter to return the target database group that matches the specified OCID.
* `target_database_group_filter` - (Optional) The scim query filter parameter accepts filter expressions that use the syntax described in Section 3.2.2.2 of the System for Cross-Domain Identity Management (SCIM) specification, which is available at [RFC3339](https://tools.ietf.org/html/draft-ietf-scim-api-12). In SCIM filtering expressions, text, date, and time values must be enclosed in quotation marks, with date and time values using ISO-8601 format. (Numeric and boolean values should not be quoted.) Ex:** filter=(targetDatabaseId eq 'ocid1.datasafetargetdatabase.oc1.iad.abuwcljr3u2va4ba5wek53idpe5qq5kkbigzclscc6mysfecxzjt5dgmxqza') 
* `time_created_greater_than_or_equal_to` - (Optional) A filter to return only the resources that were created after the specified date and time, as defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Using TimeCreatedGreaterThanOrEqualToQueryParam parameter retrieves all resources created after that date.

	**Example:** 2016-12-19T16:39:57.600Z 
* `time_created_less_than` - (Optional) Search for resources that were created before a specific date. Specifying this parameter corresponding `timeCreatedLessThan` parameter will retrieve all resources created before the specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339.

	**Example:** 2016-12-19T16:39:57.600Z 


## Attributes Reference

The following attributes are exported:

* `target_database_group_collection` - The list of target_database_group_collection.

### TargetDatabaseGroup Reference

The following attributes are exported:

* `compartment_id` - The OCID for the compartment containing the target database group.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
* `description` - Description of the target database group.
* `display_name` - The name of the target database group.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the specified target database group.
* `lifecycle_details` - Details for the lifecycle status of the target database group.
* `matching_criteria` - Criteria to either include or exclude target databases from the target database group. These criteria can be based on compartments or tags or a list of target databases. See examples below for more details. Include: Target databases will be added to the target database group if they match at least one of the include criteria. Exclude: Target databases that will be excluded from the target database group (even if they match any of the include criteria). 
	* `exclude` - Criteria to exclude certain target databases from the target database group.
		* `target_database_ids` - The list of target database OCIDS, that should be excluded from the target database group (even if they match some of the other criteria).
	* `include` - Criteria to determine whether a target database should be included in the target database group. If the database satisfies any of compartments, targetDatabaseIds, freeformTags, or definedTags criteria, it qualifies for inclusion in the target database group 
		* `compartments` - List of compartment objects, each containing the OCID of the compartment and a boolean value that indicates whether the target databases in the compartments and sub-compartments should also be included in the target database group.
			* `id` - The OCID of the compartment for including target databases to the target database group. All target databases in the compartment will be members of the target database group.
			* `is_include_subtree` - This indicates whether the target databases of sub-compartments should also be included in the target database group. By default, this parameter is set to false.
		* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
		* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
		* `target_database_ids` - The list of target database OCIDs to be included in the target database group.
* `membership_count` - The number of target databases in the specified target database group.
* `membership_update_time` - Time when the members of the target database group were last changed, i.e. the list was refreshed, a target database was added or removed.
* `state` - The lifecycle status of the target database group.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Time when the target database group was created.
* `time_updated` - Time when the target database group was last updated.

