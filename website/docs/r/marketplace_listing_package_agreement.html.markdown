---
subcategory: "Marketplace"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_marketplace_listing_package_agreement"
sidebar_current: "docs-oci-datasource-marketplace-listing_package_agreement"
description: |-
  Provides details about a specific Listing Package Agreement in Oracle Cloud Infrastructure Marketplace service
---

# Data Source: oci_marketplace_listing_package_agreement_management
This resource provides details about a specific Listing Package Agreement resource in Oracle Cloud Infrastructure Marketplace service.

This resource can be used to retrieve the time-based signature of terms of use agreement for a package that can be used to
accept the agreement.


## Example Usage

```hcl
resource "oci_marketplace_listing_package_agreement" "test_listing_package_agreement" {
	#Required
	agreement_id = oci_marketplace_agreement.test_agreement.id
	listing_id = oci_marketplace_listing.test_listing.id
	package_version = var.listing_package_agreement_package_version

	#Optional
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `agreement_id` - (Required) The unique identifier for the agreement.
* `listing_id` - (Required) The unique identifier for the listing.
* `package_version` - (Required) The version of the package. Package versions are unique within a listing.
* `compartment_id` - (Optional) The unique identifier for the compartment, required in gov regions.

## Attributes Reference

The following attributes are exported:

* `agreement_id` - The unique identifier for the agreement.
* `listing_id` - The unique identifier for the listing.
* `package_version` - The version of the package. Package versions are unique within a listing.
* `author` - Who authored the agreement.
* `compartment_id` - The unique identifier for the compartment.
* `content_url` - The content URL of the agreement.
* `id` - The unique identifier for the agreement.
* `prompt` - Textual prompt to read and accept the agreement.
* `signature` - A time-based signature that can be used to accept an agreement or remove a previously accepted agreement from the list that Marketplace checks before a deployment. 

## Import

Import is not supported for this resource.