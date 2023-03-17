---
subcategory: "Functions"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_functions_pbf_listings"
sidebar_current: "docs-oci-datasource-functions-pbf_listings"
description: |-
  Provides the list of Pbf Listings in Oracle Cloud Infrastructure Functions service
---

# Data Source: oci_functions_pbf_listings
This data source provides the list of Pbf Listings in Oracle Cloud Infrastructure Functions service.

Fetches a wrapped list of all Pre-built Function(PBF) Listings. Returns a PbfListingCollection containing 
an array of PbfListingSummary response models.


## Example Usage

```hcl
data "oci_functions_pbf_listings" "test_pbf_listings" {

	#Optional
	name = var.pbf_listing_name
	name_contains = var.pbf_listing_name_contains
	name_starts_with = var.pbf_listing_name_starts_with
	pbf_listing_id = oci_functions_pbf_listing.test_pbf_listing.id
	state = var.pbf_listing_state
	trigger = var.pbf_listing_trigger
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) A filter to return only resources that match the entire PBF name given.
* `name_contains` - (Optional) A filter to return only resources that contain the supplied filter text in the PBF name given.
* `name_starts_with` - (Optional) A filter to return only resources that start with the supplied filter text in the PBF name given.
* `pbf_listing_id` - (Optional) unique PbfListing identifier
* `state` - (Optional) A filter to return only resources their lifecycleState matches the given lifecycleState.
* `trigger` - (Optional) A filter to return only resources that match the service trigger sources of a PBF.


## Attributes Reference

The following attributes are exported:

* `pbf_listings_collection` - The list of pbf_listings_collection.

### PbfListing Reference

The following attributes are exported:

* `items` - List of PbfListingSummary.
    * `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
    * `description` - A short overview of the PBF Listing: the purpose of the PBF and and associated information.
    * `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
    * `id` - Unique identifier that is immutable on creation.
    * `name` - A brief descriptive name for the PBF listing. The PBF listing name must be unique, and not match and existing PBF. 
    * `publisher_details` - Contains details about the publisher of this PBF Listing.
	    * `name` - Name of the Publisher
    * `state` - The current state of the PBF resource.
    * `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
    * `time_created` - The time the PbfListing was created. An RFC3339 formatted datetime string.
    * `time_updated` - The last time the PbfListing was updated. An RFC3339 formatted datetime string.
    * `triggers` - An array of Trigger. A list of triggers that may activate the PBF.
	    * `name` - A brief descriptive name for the PBF trigger.

