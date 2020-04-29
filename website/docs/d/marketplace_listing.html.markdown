---
subcategory: "Marketplace"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_marketplace_listing"
sidebar_current: "docs-oci-datasource-marketplace-listing"
description: |-
  Provides details about a specific Listing in Oracle Cloud Infrastructure Marketplace service
---

# Data Source: oci_marketplace_listing
This data source provides details about a specific Listing resource in Oracle Cloud Infrastructure Marketplace service.

Gets detailed information about a listing, including the listing's name, version, description, and
resources.


## Example Usage

```hcl
data "oci_marketplace_listing" "test_listing" {
	#Required
	listing_id = "${oci_marketplace_listing.test_listing.id}"

	#Optional
	compartment_id = "${var.compartment_id}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The unique identifier for the compartment.
* `listing_id` - (Required) The unique identifier for the listing.


## Attributes Reference

The following attributes are exported:

* `banner` - 
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
* `icon` - 
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
* `long_description` - A long description of the listing.
* `name` - The name of the listing.
* `package_type` - The listing's package type.
* `publisher` - 
	* `contact_email` - The email address of the publisher.
	* `contact_phone` - The phone number of the publisher.
	* `description` - A description of the publisher.
	* `hq_address` - The address of the publisher's headquarters.
	* `id` - Unique identifier for the publisher.
	* `links` - Reference links.
		* `href` - The anchor tag.
		* `rel` - Reference links to the previous page, next page, and other pages.
	* `logo` - 
		* `content_url` - The content URL of the upload data.
		* `file_extension` - The file extension of the upload data.
		* `mime_type` - The MIME type of the upload data.
		* `name` - The name used to refer to the upload data.
	* `name` - The name of the publisher.
	* `website_url` - The publisher's website.
	* `year_founded` - The year the publisher's company or organization was founded.
* `regions` - The regions where the listing is available.
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
* `system_requirements` - System requirements for the listing.
* `tagline` - The tagline of the listing.
* `time_released` - The release date of the listing.
* `usage_information` - Usage information for the listing.
* `version` - The version of the listing.
* `videos` - Videos of the listing.
	* `name` - Text that describes the resource.
	* `url` - The URL of the resource.

