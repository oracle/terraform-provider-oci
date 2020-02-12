---
subcategory: "Marketplace"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_marketplace_listing_package"
sidebar_current: "docs-oci-datasource-marketplace-listing_package"
description: |-
  Provides details about a specific Listing Package in Oracle Cloud Infrastructure Marketplace service
---

# Data Source: oci_marketplace_listing_package
This data source provides details about a specific Listing Package resource in Oracle Cloud Infrastructure Marketplace service.

Get the details of the specified version of a package, including information needed to launch the package.


## Example Usage

```hcl
data "oci_marketplace_listing_package" "test_listing_package" {
	#Required
	listing_id = "${oci_marketplace_listing.test_listing.id}"
	package_version = "${var.listing_package_package_version}"
}
```

## Argument Reference

The following arguments are supported:

* `listing_id` - (Required) The unique identifier for the listing.
* `package_version` - (Required) The version of the package. Package versions are unique within a listing.


## Attributes Reference

The following attributes are exported:

* `app_catalog_listing_id` - The id of the AppCatalogListing associated with this ListingPackage.
* `app_catalog_listing_resource_version` - The resource version of the AppCatalogListing associated with this ListingPackage.
* `description` - Description of this package.
* `listing_id` - The id of the listing the specified package belongs to.
* `package_type` - The specified package's type.
* `pricing` - 
	* `currency` - The currency of the pricing model.
	* `pay_go_strategy` - The type of pricing for a PAYGO model, eg PER_OCPU_LINEAR, PER_OCPU_MIN_BILLING, PER_INSTANCE.  Null if type is not PAYGO.
	* `rate` - The pricing rate.
	* `type` - The type of the pricing model.
* `regions` - List of regions in which this ListingPackage is available.
	* `code` - The code of the region.
	* `countries` - Countries in the region.
		* `code` - A code assigned to the item.
		* `name` - The name of the item.
	* `name` - The name of the region.
* `resource_id` - The unique identifier for the package resource.
* `resource_link` - Link to the orchestration resource.
* `time_created` - The date and time this listing package was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339)  timestamp format.  Example: `2016-08-25T21:10:29.600Z` 
* `variables` - List of variables for the orchestration resource.
	* `data_type` - The data type of the variable.
	* `default_value` - The variable's default value.
	* `description` - A description of the variable.
	* `hint_message` - A brief textual description that helps to explain the variable.
	* `is_mandatory` - Whether the variable is mandatory.
	* `name` - The name of the variable.
* `version` - The version of this package.

