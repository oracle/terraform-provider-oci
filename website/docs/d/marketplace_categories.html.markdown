---
subcategory: "Marketplace"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_marketplace_categories"
sidebar_current: "docs-oci-datasource-marketplace-categories"
description: |-
  Provides the list of Categories in Oracle Cloud Infrastructure Marketplace service
---

# Data Source: oci_marketplace_categories
This data source provides the list of Categories in Oracle Cloud Infrastructure Marketplace service.

Gets the list of all the categories for listings published to Oracle Cloud Infrastructure Marketplace. Categories apply
to the software product provided by the listing.


## Example Usage

```hcl
data "oci_marketplace_categories" "test_categories" {

	#Optional
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The unique identifier for the compartment.


## Attributes Reference

The following attributes are exported:

* `categories` - The list of categories.

### Category Reference

The following attributes are exported:

* `name` - Name of the product category.

