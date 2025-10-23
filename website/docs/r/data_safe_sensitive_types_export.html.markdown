---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sensitive_types_export"
sidebar_current: "docs-oci-resource-data_safe-sensitive_types_export"
description: |-
  Provides the Sensitive Types Export resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_sensitive_types_export
This resource provides the Sensitive Types Export resource in Oracle Cloud Infrastructure Data Safe service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/data-safe/latest/SensitiveTypesExport

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/datasafe

Generates a downloadable file corresponding to the specified list of sensitive types. It's a prerequisite for the
DownloadSensitiveTypesExport operation. Use this endpoint to generate a sensitive Types Export file and then use 
DownloadSensitiveTypesExport to download the generated file.


## Example Usage

```hcl
resource "oci_data_safe_sensitive_types_export" "test_sensitive_types_export" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.sensitive_types_export_description
	display_name = var.sensitive_types_export_display_name
	freeform_tags = {"Department"= "Finance"}
	is_include_all_sensitive_types = var.sensitive_types_export_is_include_all_sensitive_types
	sensitive_type_ids_for_export = var.sensitive_types_export_sensitive_type_ids_for_export
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment where the sensitive types export should be created.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) The description of the sensitive types export.
* `display_name` - (Optional) (Updatable) The display name of the sensitive types export. The name does not have to be unique, and it's changeable.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `is_include_all_sensitive_types` - (Optional) Indicates if all the existing user-defined sensitive types are used for export. If it's set to true, the sensitiveTypeIdsForExport attribute is ignored and all user-defined sensitive types are used. 
* `sensitive_type_ids_for_export` - (Optional) The OCIDs of the sensitive types used to create sensitive types export. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Sensitive Types Export
	* `update` - (Defaults to 20 minutes), when updating the Sensitive Types Export
	* `delete` - (Defaults to 20 minutes), when destroying the Sensitive Types Export


## Import

SensitiveTypesExports can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_sensitive_types_export.test_sensitive_types_export "id"
```

