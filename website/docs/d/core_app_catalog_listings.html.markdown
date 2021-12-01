---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_app_catalog_listings"
sidebar_current: "docs-oci-datasource-core-app_catalog_listings"
description: |-
  Provides the list of App Catalog Listings in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_app_catalog_listings
This data source provides the list of App Catalog Listings in Oracle Cloud Infrastructure Core service.

Lists the published listings.

## Example Usage

```hcl
data "oci_core_app_catalog_listings" "test_app_catalog_listings" {

	#Optional
	display_name = var.app_catalog_listing_display_name
	publisher_name = var.app_catalog_listing_publisher_name
	publisher_type = var.app_catalog_listing_publisher_type
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `publisher_name` - (Optional) A filter to return only the publisher that matches the given publisher name exactly. 
* `publisher_type` - (Optional) A filter to return only publishers that match the given publisher type exactly. Valid types are OCI, ORACLE, TRUSTED, STANDARD. 


## Attributes Reference

The following attributes are exported:

* `app_catalog_listings` - The list of app_catalog_listings.

### AppCatalogListing Reference

The following attributes are exported:

* `contact_url` - Listing's contact URL.
* `description` - Description of the listing.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `listing_id` - the region free ocid of the listing resource.
* `publisher_logo_url` - Publisher's logo URL.
* `publisher_name` - The name of the publisher who published this listing.
* `summary` - The short summary for the listing.
* `time_published` - Date and time the listing was published, in [RFC3339](https://tools.ietf.org/html/rfc3339) format. Example: `2018-03-20T12:32:53.532Z` 

