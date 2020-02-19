---
subcategory: "Marketplace"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_marketplace_listing_packages"
sidebar_current: "docs-oci-datasource-marketplace-listing_packages"
description: |-
  Provides the list of Listing Packages in Oracle Cloud Infrastructure Marketplace service
---

# Data Source: oci_marketplace_listing_packages
This data source provides the list of Listing Packages in Oracle Cloud Infrastructure Marketplace service.

Gets the list of packages for a listing.


## Example Usage

```hcl
data "oci_marketplace_listing_packages" "test_listing_packages" {
	#Required
	listing_id = "${oci_marketplace_listing.test_listing.id}"

	#Optional
	package_type = "${var.listing_package_package_type}"
	package_version = "${var.listing_package_package_version}"
}
```

## Argument Reference

The following arguments are supported:

* `listing_id` - (Required) The unique identifier for the listing.
* `package_type` - (Optional) A filter to return only packages that match the given package type exactly. 
* `package_version` - (Optional) The version of the package. Package versions are unique within a listing.


## Attributes Reference

The following attributes are exported:

* `listing_packages` - The list of listing_packages.

### ListingPackage Reference

The following attributes are exported:

* `listing_id` - The id of the listing the specified package belongs to.
* `resource_id` - The unique identifier for the package resource.
* `time_created` - The date and time this listing package was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339)  timestamp format.  Example: `2016-08-25T21:10:29.600Z` 
* `package_type` - The specified package's type.
* `package_version` - The version of this package.
