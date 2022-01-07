---
subcategory: "Marketplace"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_marketplace_publication_packages"
sidebar_current: "docs-oci-datasource-marketplace-publication_packages"
description: |-
  Provides the list of Publication Packages in Oracle Cloud Infrastructure Marketplace service
---

# Data Source: oci_marketplace_publication_packages
This data source provides the list of Publication Packages in Oracle Cloud Infrastructure Marketplace service.

Lists the packages in the specified publication.

## Example Usage

```hcl
data "oci_marketplace_publication_packages" "test_publication_packages" {
	#Required
	publication_id = oci_marketplace_publication.test_publication.id

	#Optional
	package_type = var.publication_package_package_type
	package_version = var.publication_package_package_version
}
```

## Argument Reference

The following arguments are supported:

* `package_type` - (Optional) A filter to return only packages that match the given package type exactly. 
* `package_version` - (Optional) The version of the package. Package versions are unique within a listing.
* `publication_id` - (Required) The unique identifier for the publication.


## Attributes Reference

The following attributes are exported:

* `publication_packages` - The list of publication_packages.

### PublicationPackage Reference

The following attributes are exported:

* `app_catalog_listing_id` - The ID of the listing resource associated with this publication package. For more information, see [AppCatalogListing](https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/latest/AppCatalogListing/) in the Core Services API. 
* `app_catalog_listing_resource_version` - The resource version of the listing resource associated with this publication package.
* `description` - A description of the package.
* `image_id` - The ID of the image that corresponds to the package.
* `listing_id` - The ID of the listing that the specified package belongs to.
* `operating_system` - The operating system used by the listing.
	* `name` - The name of the operating system.
* `package_type` - The specified package's type.
* `resource_id` - The unique identifier for the package resource.
* `resource_link` - A link to the stack resource.
* `time_created` - The date and time the publication package was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 
* `variables` - A list of variables for the stack resource.
	* `data_type` - The data type of the variable.
	* `default_value` - The variable's default value.
	* `description` - A description of the variable.
	* `hint_message` - A brief textual description that helps to explain the variable.
	* `is_mandatory` - Whether the variable is mandatory.
	* `name` - The name of the variable.
* `version` - The package version.

