---
subcategory: "Metering Computation"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_metering_computation_custom_table"
sidebar_current: "docs-oci-datasource-metering_computation-custom_table"
description: |-
  Provides details about a specific Custom Table in Oracle Cloud Infrastructure Metering Computation service
---

# Data Source: oci_metering_computation_custom_table
This data source provides details about a specific Custom Table resource in Oracle Cloud Infrastructure Metering Computation service.

Returns the saved custom table.


## Example Usage

```hcl
data "oci_metering_computation_custom_table" "test_custom_table" {
	#Required
	custom_table_id = oci_metering_computation_custom_table.test_custom_table.id
}
```

## Argument Reference

The following arguments are supported:

* `custom_table_id` - (Required) The custom table unique OCID.


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

