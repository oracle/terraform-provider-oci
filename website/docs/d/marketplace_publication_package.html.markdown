---
subcategory: "Marketplace"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_marketplace_publication_package"
sidebar_current: "docs-oci-datasource-marketplace-publication_package"
description: |-
  Provides details about a specific Publication Package in Oracle Cloud Infrastructure Marketplace service
---

# Data Source: oci_marketplace_publication_package
This data source provides details about a specific Publication Package resource in Oracle Cloud Infrastructure Marketplace service.

Gets the details of a specific package within a given Publication

## Example Usage

```hcl
data "oci_marketplace_publication_package" "test_publication_package" {
	#Required
	package_version = var.publication_package_package_version
	publication_id = oci_marketplace_publication.test_publication.id
}
```

## Argument Reference

The following arguments are supported:

* `package_version` - (Required) The version of the package. Package versions are unique within a listing.
* `publication_id` - (Required) The unique identifier for the listing.


## Attributes Reference

The following attributes are exported:

* `app_catalog_listing_id` - The ID of the listing resource associated with this publication package. For more information, see [AppCatalogListing](https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/latest/AppCatalogListing/) in the Core Services API. 
* `app_catalog_listing_resource_version` - The resource version of the listing resource associated with this listing package.
* `description` - Description of this package.
* `image_id` - The ID of the image corresponding to the package.
* `listing_id` - The ID of the listing that the specified package belongs to.
* `operating_system` - OS used by the listing.
	* `name` - name of the operating system
* `package_type` - The specified package's type.
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

