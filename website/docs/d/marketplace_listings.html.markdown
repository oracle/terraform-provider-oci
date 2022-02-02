---
subcategory: "Marketplace"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_marketplace_listings"
sidebar_current: "docs-oci-datasource-marketplace-listings"
description: |-
  Provides the list of Listings in Oracle Cloud Infrastructure Marketplace service
---

# Data Source: oci_marketplace_listings
This data source provides the list of Listings in Oracle Cloud Infrastructure Marketplace service.

Gets a list of listings from Oracle Cloud Infrastructure Marketplace by searching keywords and
filtering according to listing attributes.

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
data "oci_marketplace_listings" "test_listings" {

	#Optional
	category = var.listing_category
	compartment_id = var.compartment_id
	image_id = oci_core_image.test_image.id
	is_featured = var.listing_is_featured
	listing_id = oci_marketplace_listing.test_listing.id
	listing_types = var.listing_listing_types
	name = var.listing_name
	operating_systems = var.listing_operating_systems
	package_type = var.listing_package_type
	pricing = var.listing_pricing
	publisher_id = oci_marketplace_publisher.test_publisher.id
}
```

## Argument Reference

The following arguments are supported:

* `category` - (Optional) Name of the product category or categories. If you specify multiple categories, then Marketplace returns any listing with one or more matching categories. 
* `compartment_id` - (Optional) The unique identifier for the compartment.
* `image_id` - (Optional) Image ID of the listing
* `is_featured` - (Optional) Indicates whether to show only featured listings. If this is set to `false` or is omitted, then all listings will be returned. 
* `listing_id` - (Optional) The unique identifier for the listing.
* `listing_types` - (Optional) The type of the listing.
* `name` - (Optional) The name of the listing.
* `operating_systems` - (Optional) The operating system of the listing.
* `package_type` - (Optional) A filter to return only packages that match the given package type exactly. 
* `pricing` - (Optional) Name of the pricing type. If multiple pricing types are provided, then any listing with one or more matching pricing models will be returned. 
* `publisher_id` - (Optional) Limit results to just this publisher.


## Attributes Reference

The following attributes are exported:

* `listings` - The list of listings.

### Listing Reference

The following attributes are exported:

* `banner` - The model for upload data for images and icons.
	* `content_url` - The content URL of the upload data.
	* `file_extension` - The file extension of the upload data.
	* `mime_type` - The MIME type of the upload data.
	* `name` - The name used to refer to the upload data.
* `categories` - Product categories that the listing belongs to.
* `default_package_version` - The default package version.
* `documentation_links` - Links to additional documentation provided by the publisher specifically for the listing.
	* `document_category` - The category that the document belongs to.
	* `name` - Text that describes the resource.
	* `url` - The URL of the resource.
* `icon` - The model for upload data for images and icons.
	* `content_url` - The content URL of the upload data.
	* `file_extension` - The file extension of the upload data.
	* `mime_type` - The MIME type of the upload data.
	* `name` - The name used to refer to the upload data.
* `id` - The unique identifier for the listing in Marketplace.
* `is_featured` - Indicates whether the listing is included in Featured Listings.
* `keywords` - Keywords associated with the listing.
* `languages` - Languages supported by the listing.
	* `code` - A code assigned to the item.
	* `name` - The name of the item.
* `license_model_description` - A description of the publisher's licensing model for the listing.
* `links` - Links to reference material.
	* `href` - The anchor tag.
	* `rel` - Reference links to the previous page, next page, and other pages.
* `listing_type` - The publisher category to which the listing belongs. The publisher category informs where the listing appears for use.
* `long_description` - A long description of the listing.
* `name` - The name of the listing.
* `package_type` - The listing's package type.
* `publisher` - Summary details about the publisher of the listing.
	* `contact_email` - The email address of the publisher.
	* `contact_phone` - The phone number of the publisher.
	* `description` - A description of the publisher.
	* `hq_address` - The address of the publisher's headquarters.
	* `id` - The unique identifier for the publisher.
	* `links` - Reference links.
		* `href` - The anchor tag.
		* `rel` - Reference links to the previous page, next page, and other pages.
	* `logo` - The model for upload data for images and icons.
		* `content_url` - The content URL of the upload data.
		* `file_extension` - The file extension of the upload data.
		* `mime_type` - The MIME type of the upload data.
		* `name` - The name used to refer to the upload data.
	* `name` - The name of the publisher.
	* `website_url` - The publisher's website.
	* `year_founded` - The year the publisher's company or organization was founded.
* `regions` - The regions where the listing is eligible to be deployed.
	* `code` - The code of the region.
	* `countries` - Countries in the region.
		* `code` - A code assigned to the item.
		* `name` - The name of the item.
	* `name` - The name of the region.
* `release_notes` - Release notes for the listing.
* `screenshots` - Screenshots of the listing.
	* `content_url` - The content URL of the screenshot.
	* `description` - A description of the screenshot.
	* `file_extension` - The file extension of the screenshot.
	* `mime_type` - The MIME type of the screenshot.
	* `name` - The name of the screenshot.
* `short_description` - A short description of the listing.
* `support_contacts` - Contact information to use to get support from the publisher for the listing.
	* `email` - The email of the contact.
	* `name` - The name of the contact.
	* `phone` - The phone number of the contact.
	* `subject` - The email subject line to use when contacting support.
* `support_links` - Links to support resources for the listing.
	* `name` - Text that describes the resource.
	* `url` - The URL of the resource.
* `supported_operating_systems` - The list of operating systems supported by the listing.
	* `name` - The name of the operating system.
* `system_requirements` - System requirements for the listing.
* `tagline` - The tagline of the listing.
* `time_released` - The release date of the listing.
* `usage_information` - Usage information for the listing.
* `version` - The version of the listing.
* `videos` - Videos of the listing.
	* `name` - Text that describes the resource.
	* `url` - The URL of the resource.

