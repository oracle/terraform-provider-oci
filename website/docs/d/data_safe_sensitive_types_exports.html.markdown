---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sensitive_types_exports"
sidebar_current: "docs-oci-datasource-data_safe-sensitive_types_exports"
description: |-
  Provides the list of Sensitive Types Exports in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_sensitive_types_exports
This data source provides the list of Sensitive Types Exports in Oracle Cloud Infrastructure Data Safe service.

Retrieves a list of all sensitive types export in Data Safe based on the specified query parameters.
The ListSensitiveTypesExports operation returns only the sensitive types export in the specified `compartmentId`.


## Example Usage

```hcl
data "oci_data_safe_sensitive_types_exports" "test_sensitive_types_exports" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.sensitive_types_export_access_level
	compartment_id_in_subtree = var.sensitive_types_export_compartment_id_in_subtree
	display_name = var.sensitive_types_export_display_name
	sensitive_types_export_id = oci_data_safe_sensitive_types_export.test_sensitive_types_export.id
	state = var.sensitive_types_export_state
	time_created_greater_than_or_equal_to = var.sensitive_types_export_time_created_greater_than_or_equal_to
	time_created_less_than = var.sensitive_types_export_time_created_less_than
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `display_name` - (Optional) A filter to return only resources that match the specified display name. 
* `sensitive_types_export_id` - (Optional) An optional filter to return only resources that match the specified OCID of the sensitive types export resource.
* `state` - (Optional) A filter to return only the resources that match the specified lifecycle state.
* `time_created_greater_than_or_equal_to` - (Optional) A filter to return only the resources that were created after the specified date and time, as defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Using TimeCreatedGreaterThanOrEqualToQueryParam parameter retrieves all resources created after that date.

	**Example:** 2016-12-19T16:39:57.600Z 
* `time_created_less_than` - (Optional) Search for resources that were created before a specific date. Specifying this parameter corresponding `timeCreatedLessThan` parameter will retrieve all resources created before the specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339.

	**Example:** 2016-12-19T16:39:57.600Z 


## Attributes Reference

The following attributes are exported:

* `sensitive_types_export_collection` - The list of sensitive_types_export_collection.

### SensitiveTypesExport Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the sensitive types export.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the sensitive types export.
* `display_name` - The display name of the sensitive types export.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the sensitive types export.
* `is_include_all_sensitive_types` - Indicates if all the existing user-defined sensitive types are used for export. If it's set to true, the sensitiveTypeIdsForExport attribute is ignored and all user-defined sensitive types are exported. 
* `sensitive_type_ids_for_export` - The OCIDs of the sensitive types used to create sensitive types export. 
* `state` - The current state of the sensitive types export.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the sensitive types export was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_updated` - The date and time the sensitive types export was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

