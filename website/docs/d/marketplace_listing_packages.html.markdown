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

If you plan to launch an instance from an image listing, you must first subscribe to the listing. When
you launch the instance, you also need to provide the image ID of the listing resource version that you want.

Subscribing to the listing requires you to first get a signature from the terms of use agreement for the
listing resource version. To get the signature, issue a [GetAppCatalogListingAgreements](https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/latest/AppCatalogListingResourceVersionAgreements/GetAppCatalogListingAgreements) API call.
The [AppCatalogListingResourceVersionAgreements](https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/latest/AppCatalogListingResourceVersionAgreements) object, including
its signature, is returned in the response. With the signature for the terms of use agreement for the desired
listing resource version, create a subscription by issuing a
[CreateAppCatalogSubscription](https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/latest/AppCatalogSubscription/CreateAppCatalogSubscription) API call.

To get the image ID to launch an instance, issue a [GetAppCatalogListingResourceVersion](https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/latest/AppCatalogListingResourceVersion/GetAppCatalogListingResourceVersion) API call.
Lastly, to launch the instance, use the image ID of the listing resource version to issue a [LaunchInstance](https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/latest/Instance/LaunchInstance) API call.


## Example Usage

```hcl
data "oci_marketplace_listing_packages" "test_listing_packages" {
	#Required
	listing_id = oci_marketplace_listing.test_listing.id

	#Optional
	compartment_id = var.compartment_id
	package_type = var.listing_package_package_type
	package_version = var.listing_package_package_version
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The unique identifier for the compartment.
* `listing_id` - (Required) The unique identifier for the listing.
* `package_type` - (Optional) A filter to return only packages that match the given package type exactly. 
* `package_version` - (Optional) The version of the package. Package versions are unique within a listing.


## Attributes Reference

The following attributes are exported:

* `listing_packages` - The list of listing_packages.

### ListingPackage Reference

The following attributes are exported:

* `app_catalog_listing_id` - The ID of the listing resource associated with this listing package. For more information, see [AppCatalogListing](https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/latest/AppCatalogListing/) in the Core Services API. 
* `app_catalog_listing_resource_version` - The resource version of the listing resource associated with this listing package.
* `description` - Description of this package.
* `image_id` - The ID of the image corresponding to the package.
* `listing_id` - The ID of the listing that the specified package belongs to.
* `operating_system` - The operating system used by the listing.
	* `name` - The name of the operating system.
* `package_type` - The specified package's type.
* `pricing` - The model for pricing.
	* `currency` - The currency of the pricing model.
	* `pay_go_strategy` - The type of pricing for a PAYGO model, eg PER_OCPU_LINEAR, PER_OCPU_MIN_BILLING, PER_INSTANCE.  Null if type is not PAYGO.
	* `rate` - The pricing rate.
	* `type` - The type of the pricing model.
* `regions` - The regions where the listing is available.
	* `code` - The code of the region.
	* `countries` - Countries in the region.
		* `code` - A code assigned to the item.
		* `name` - The name of the item.
	* `name` - The name of the region.
* `resource_id` - The unique identifier for the package resource.
* `resource_link` - Link to the orchestration resource.
* `time_created` - The date and time this listing package was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 
* `variables` - List of variables for the orchestration resource.
	* `data_type` - The data type of the variable.
	* `default_value` - The variable's default value.
	* `description` - A description of the variable.
	* `hint_message` - A brief textual description that helps to explain the variable.
	* `is_mandatory` - Whether the variable is mandatory.
	* `name` - The name of the variable.
* `version` - The package version.

