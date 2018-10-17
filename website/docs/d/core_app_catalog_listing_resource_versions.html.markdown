---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_app_catalog_listing_resource_versions"
sidebar_current: "docs-oci-datasource-core-app_catalog_listing_resource_versions"
description: |-
  Provides the list of App Catalog Listing Resource Versions in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_app_catalog_listing_resource_versions
This data source provides the list of App Catalog Listing Resource Versions in Oracle Cloud Infrastructure Core service.

Gets all resource versions for a particular listing.

## Example Usage

```hcl
data "oci_core_app_catalog_listing_resource_versions" "test_app_catalog_listing_resource_versions" {
	#Required
	listing_id = "${oci_core_listing.test_listing.id}"
}
```

## Argument Reference

The following arguments are supported:

* `listing_id` - (Required) The OCID of the listing.


## Attributes Reference

The following attributes are exported:

* `app_catalog_listing_resource_versions` - The list of app_catalog_listing_resource_versions.

### AppCatalogListingResourceVersion Reference

The following attributes are exported:

* `accessible_ports` - List of accessible ports for instances launched with this listing resource version.
* `allowed_actions` - Allowed actions for the listing resource.
* `available_regions` - List of regions that this listing resource version is available.

	For information about Regions, see [Regions](../../../#General/Concepts/regions.htm).

	Example: `["us-ashburn-1", "us-phoenix-1"]` 
* `compatible_shapes` - Array of shapes compatible with this resource.

	You may enumerate all available shapes by calling [ListShapes] (#listShapes).

	Example: `["VM.Standard1.1", "VM.Standard1.2"]` 
* `listing_id` - The OCID of the listing this resource version belongs to.
* `listing_resource_id` - OCID of the listing resource.
* `listing_resource_version` - Resource Version.
* `time_published` - Date and time the listing resource version was published, in RFC3339 format. Example: `2018-03-20T12:32:53.532Z` 

