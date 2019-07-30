---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_app_catalog_listing_resource_version"
sidebar_current: "docs-oci-datasource-core-app_catalog_listing_resource_version"
description: |-
  Provides details about a specific App Catalog Listing Resource Version in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_app_catalog_listing_resource_version
This data source provides details about a specific App Catalog Listing Resource Version resource in Oracle Cloud Infrastructure Core service.

Gets the specified listing resource version.

## Example Usage

```hcl
data "oci_core_app_catalog_listing_resource_version" "test_app_catalog_listing_resource_version" {
	#Required
	listing_id = "${data.oci_core_app_catalog_listing.test_listing.id}"
	resource_version = "${var.app_catalog_listing_resource_version_resource_version}"
}
```

## Argument Reference

The following arguments are supported:

* `listing_id` - (Required) The OCID of the listing.
* `resource_version` - (Required) Listing Resource Version.


## Attributes Reference

The following attributes are exported:

* `accessible_ports` - List of accessible ports for instances launched with this listing resource version.
* `allowed_actions` - Allowed actions for the listing resource.
* `available_regions` - List of regions that this listing resource version is available.

	For information about Regions, see [Regions](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm).

	Example: `["us-ashburn-1", "us-phoenix-1"]` 
* `compatible_shapes` - Array of shapes compatible with this resource.

	You may enumerate all available shapes by calling [ListShapes] (https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Shape/ListShapes).

	Example: `["VM.Standard1.1", "VM.Standard1.2"]` 
* `listing_id` - The OCID of the listing this resource version belongs to.
* `listing_resource_id` - OCID of the listing resource.
* `listing_resource_version` - Resource Version.
* `time_published` - Date and time the listing resource version was published, in RFC3339 format. Example: `2018-03-20T12:32:53.532Z` 

