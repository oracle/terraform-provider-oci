---
subcategory: "Marketplace"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_marketplace_listing_taxes"
sidebar_current: "docs-oci-datasource-marketplace-listing_taxes"
description: |-
  Provides the list of Listing Taxes in Oracle Cloud Infrastructure Marketplace service
---

# Data Source: oci_marketplace_listing_taxes
This data source provides the list of Listing Taxes in Oracle Cloud Infrastructure Marketplace service.

Returns list of all tax implications that current tenant may be liable to once they launch the listing.

## Example Usage

```hcl
data "oci_marketplace_listing_taxes" "test_listing_taxes" {
	#Required
	listing_id = oci_marketplace_listing.test_listing.id

	#Optional
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The unique identifier for the compartment.
* `listing_id` - (Required) The unique identifier for the listing.


## Attributes Reference

The following attributes are exported:

* `taxes` - The list of taxes.

### ListingTax Reference

The following attributes are exported:

* `code` - Unique code for the tax.
* `country` - Country, which imposes the tax.
* `name` - Name of the tax code.
* `url` - The URL with more details about this tax.

