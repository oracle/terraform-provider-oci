---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sensitive_types_export"
sidebar_current: "docs-oci-datasource-data_safe-sensitive_types_export"
description: |-
  Provides details about a specific Sensitive Types Export in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_sensitive_types_export
This data source provides details about a specific Sensitive Types Export resource in Oracle Cloud Infrastructure Data Safe service.

Gets the details of the specified sensitive types export by identifier.

## Example Usage

```hcl
data "oci_data_safe_sensitive_types_export" "test_sensitive_types_export" {
	#Required
	sensitive_types_export_id = oci_data_safe_sensitive_types_export.test_sensitive_types_export.id
}
```

## Argument Reference

The following arguments are supported:

* `sensitive_types_export_id` - (Required) The OCID of the sensitive types export.


## Attributes Reference

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

