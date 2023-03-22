---
subcategory: "Functions"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_functions_pbf_listing"
sidebar_current: "docs-oci-datasource-functions-pbf_listing"
description: |-
  Provides details about a specific Pbf Listing in Oracle Cloud Infrastructure Functions service
---

# Data Source: oci_functions_pbf_listing
This data source provides details about a specific Pbf Listing resource in Oracle Cloud Infrastructure Functions service.

Fetches a Pre-built Function(PBF) Listing. Returns a PbfListing response model.


## Example Usage

```hcl
data "oci_functions_pbf_listing" "test_pbf_listing" {
	#Required
	pbf_listing_id = oci_functions_pbf_listing.test_pbf_listing.id
}
```

## Argument Reference

The following arguments are supported:

* `pbf_listing_id` - (Required) unique PbfListing identifier


## Attributes Reference

The following attributes are exported:

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

