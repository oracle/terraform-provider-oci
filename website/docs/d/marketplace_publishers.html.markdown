---
subcategory: "Marketplace"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_marketplace_publishers"
sidebar_current: "docs-oci-datasource-marketplace-publishers"
description: |-
  Provides the list of Publishers in Oracle Cloud Infrastructure Marketplace service
---

# Data Source: oci_marketplace_publishers
This data source provides the list of Publishers in Oracle Cloud Infrastructure Marketplace service.

Gets the list of all the publishers of listings available in Oracle Cloud Infrastructure Marketplace.


## Example Usage

```hcl
data "oci_marketplace_publishers" "test_publishers" {

	#Optional
	compartment_id = var.compartment_id
	publisher_id = oci_marketplace_publisher.test_publisher.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The unique identifier for the compartment.
* `publisher_id` - (Optional) Limit results to just this publisher.


## Attributes Reference

The following attributes are exported:

* `publishers` - The list of publishers.

### Publisher Reference

The following attributes are exported:

* `description` - A description of the publisher.
* `id` - The unique identifier for the publisher.
* `name` - The name of the publisher.

