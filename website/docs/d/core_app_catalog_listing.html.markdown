---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_app_catalog_listing"
sidebar_current: "docs-oci-datasource-core-app_catalog_listing"
description: |-
  Provides details about a specific App Catalog Listing in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_app_catalog_listing
This data source provides details about a specific App Catalog Listing resource in Oracle Cloud Infrastructure Core service.

Gets the specified listing.

## Example Usage

```hcl
data "oci_core_app_catalog_listing" "test_app_catalog_listing" {
	#Required
	listing_id = data.oci_core_app_catalog_listing.test_listing.id
}
```

## Argument Reference

The following arguments are supported:

* `listing_id` - (Required) The OCID of the listing.


## Attributes Reference

The following attributes are exported:

* `contact_url` - Listing's contact URL.
* `description` - Description of the listing.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `listing_id` - the region free ocid of the listing resource.
* `publisher_logo_url` - Publisher's logo URL.
* `publisher_name` - The name of the publisher who published this listing.
* `summary` - The short summary for the listing.
* `time_published` - Date and time the listing was published, in [RFC3339](https://tools.ietf.org/html/rfc3339) format. Example: `2018-03-20T12:32:53.532Z` 

