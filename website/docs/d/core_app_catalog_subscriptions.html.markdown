---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_app_catalog_subscriptions"
sidebar_current: "docs-oci-datasource-core-app_catalog_subscriptions"
description: |-
  Provides the list of App Catalog Subscriptions in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_app_catalog_subscriptions
This data source provides the list of App Catalog Subscriptions in Oracle Cloud Infrastructure Core service.

Lists subscriptions for a compartment.

## Example Usage

```hcl
data "oci_core_app_catalog_subscriptions" "test_app_catalog_subscriptions" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	listing_id = data.oci_core_app_catalog_listing.test_listing.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `listing_id` - (Optional) A filter to return only the listings that matches the given listing id. 


## Attributes Reference

The following attributes are exported:

* `app_catalog_subscriptions` - The list of app_catalog_subscriptions.

### AppCatalogSubscription Reference

The following attributes are exported:

* `compartment_id` - The compartmentID of the subscription.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `listing_id` - The ocid of the listing resource.
* `listing_resource_id` - Listing resource id.
* `listing_resource_version` - Listing resource version.
* `publisher_name` - Name of the publisher who published this listing.
* `summary` - The short summary to the listing.
* `time_created` - Date and time at which the subscription was created, in [RFC3339](https://tools.ietf.org/html/rfc3339) format. Example: `2018-03-20T12:32:53.532Z` 

