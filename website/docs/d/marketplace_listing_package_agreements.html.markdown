---
subcategory: "Marketplace"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_marketplace_listing_package_agreements"
sidebar_current: "docs-oci-datasource-marketplace-listing_package_agreements"
description: |-
  Provides the list of Listing Package Agreements in Oracle Cloud Infrastructure Marketplace service
---

# Data Source: oci_marketplace_listing_package_agreements
This data source provides the list of Listing Package Agreements in Oracle Cloud Infrastructure Marketplace service.

Returns the terms of use agreements that must be accepted before you can deploy the specified version of a package.


## Example Usage

```hcl
data "oci_marketplace_listing_package_agreements" "test_listing_package_agreements" {
	#Required
	listing_id = oci_marketplace_listing.test_listing.id
	package_version = var.listing_package_agreement_package_version

	#Optional
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The unique identifier for the compartment.
* `listing_id` - (Required) The unique identifier for the listing.
* `package_version` - (Required) The version of the package. Package versions are unique within a listing.


## Attributes Reference

The following attributes are exported:

* `agreements` - The list of agreements.

### ListingPackageAgreement Reference

The following attributes are exported:

* `author` - Who authored the agreement.
* `compartment_id` - The unique identifier for the compartment.
* `content_url` - The content URL of the agreement.
* `id` - The unique identifier for the agreement.
* `prompt` - Textual prompt to read and accept the agreement.
* `signature` - A time-based signature that can be used to accept an agreement or remove a previously accepted agreement from the list that Marketplace checks before a deployment. 

