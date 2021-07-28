---
subcategory: "Metering Computation"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_metering_computation_custom_table"
sidebar_current: "docs-oci-resource-metering_computation-custom_table"
description: |-
  Provides the Custom Table resource in Oracle Cloud Infrastructure Metering Computation service
---

# oci_metering_computation_custom_table
This resource provides the Custom Table resource in Oracle Cloud Infrastructure Metering Computation service.

Returns the created custom table.


## Example Usage

```hcl
resource "oci_metering_computation_custom_table" "test_custom_table" {
	#Required
	compartment_id = var.compartment_id
	saved_custom_table {
		#Required
		display_name = var.custom_table_saved_custom_table_display_name

		#Optional
		column_group_by = var.custom_table_saved_custom_table_column_group_by
		compartment_depth = var.custom_table_saved_custom_table_compartment_depth
		group_by_tag {

			#Optional
			key = var.custom_table_saved_custom_table_group_by_tag_key
			namespace = var.custom_table_saved_custom_table_group_by_tag_namespace
			value = var.custom_table_saved_custom_table_group_by_tag_value
		}
		row_group_by = var.custom_table_saved_custom_table_row_group_by
		version = var.custom_table_saved_custom_table_version
	}
	saved_report_id = oci_metering_computation_saved_report.test_saved_report.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment OCID.
* `saved_custom_table` - (Required) (Updatable) The custom table for Cost Analysis UI rendering.
	* `column_group_by` - (Optional) (Updatable) The column groupBy key list. example: `["tagNamespace", "tagKey", "tagValue", "service", "skuName", "skuPartNumber", "unit", "compartmentName", "compartmentPath", "compartmentId", "platform", "region", "logicalAd", "resourceId", "tenantId", "tenantName"]` 
	* `compartment_depth` - (Optional) (Updatable) The compartment depth level.
	* `display_name` - (Required) (Updatable) The name of the custom table.
	* `group_by_tag` - (Optional) (Updatable) GroupBy a specific tagKey. Provide the tagNamespace and tagKey in the tag object. Only one tag in the list is supported. For example: `[{"namespace":"oracle", "key":"createdBy"]` 
		* `key` - (Optional) (Updatable) The tag key.
		* `namespace` - (Optional) (Updatable) The tag namespace.
		* `value` - (Optional) (Updatable) The tag value.
	* `row_group_by` - (Optional) (Updatable) The row groupBy key list. example: `["tagNamespace", "tagKey", "tagValue", "service", "skuName", "skuPartNumber", "unit", "compartmentName", "compartmentPath", "compartmentId", "platform", "region", "logicalAd", "resourceId", "tenantId", "tenantName"]` 
	* `version` - (Optional) (Updatable) The version of the custom table.
* `saved_report_id` - (Required) The associated saved report OCID.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The custom table compartment OCID.
* `id` - The custom table OCID.
* `saved_custom_table` - The custom table for Cost Analysis UI rendering.
	* `column_group_by` - The column groupBy key list. example: `["tagNamespace", "tagKey", "tagValue", "service", "skuName", "skuPartNumber", "unit", "compartmentName", "compartmentPath", "compartmentId", "platform", "region", "logicalAd", "resourceId", "tenantId", "tenantName"]` 
	* `compartment_depth` - The compartment depth level.
	* `display_name` - The name of the custom table.
	* `group_by_tag` - GroupBy a specific tagKey. Provide the tagNamespace and tagKey in the tag object. Only one tag in the list is supported. For example: `[{"namespace":"oracle", "key":"createdBy"]` 
		* `key` - The tag key.
		* `namespace` - The tag namespace.
		* `value` - The tag value.
	* `row_group_by` - The row groupBy key list. example: `["tagNamespace", "tagKey", "tagValue", "service", "skuName", "skuPartNumber", "unit", "compartmentName", "compartmentPath", "compartmentId", "platform", "region", "logicalAd", "resourceId", "tenantId", "tenantName"]` 
	* `version` - The version of the custom table.
* `saved_report_id` - The custom table associated saved report OCID.

## Import

CustomTables can be imported using the `id`, e.g.

```
$ terraform import oci_metering_computation_custom_table.test_custom_table "id"
```

