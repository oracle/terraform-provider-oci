---
subcategory: "Cloud Bridge"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_bridge_supported_cloud_regions"
sidebar_current: "docs-oci-datasource-cloud_bridge-supported_cloud_regions"
description: |-
  Provides the list of Supported Cloud Regions in Oracle Cloud Infrastructure Cloud Bridge service
---

# Data Source: oci_cloud_bridge_supported_cloud_regions
This data source provides the list of Supported Cloud Regions in Oracle Cloud Infrastructure Cloud Bridge service.

Returns a list of supported cloud regions related to AssetSourceTypeParam.


## Example Usage

```hcl
data "oci_cloud_bridge_supported_cloud_regions" "test_supported_cloud_regions" {

	#Optional
	asset_source_type = var.supported_cloud_region_asset_source_type
	name_contains = var.supported_cloud_region_name_contains
}
```

## Argument Reference

The following arguments are supported:

* `asset_source_type` - (Optional) The asset source type.
* `name_contains` - (Optional) A filter to return only supported cloud regions which name contains given nameContains as sub-string.


## Attributes Reference

The following attributes are exported:

* `supported_cloud_region_collection` - The list of supported_cloud_region_collection.

### SupportedCloudRegion Reference

The following attributes are exported:

* `items` - List of supported cloud regions.
	* `asset_source_type` - The asset source type associated with the supported cloud region.
	* `defined_tags` - The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
	* `freeform_tags` - The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no predefined name, type, or namespace/scope. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
	* `name` - The supported cloud region name.
	* `state` - The current state of the supported cloud region.

