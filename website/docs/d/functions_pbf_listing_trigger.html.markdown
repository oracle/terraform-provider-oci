---
subcategory: "Functions"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_functions_pbf_listing_trigger"
sidebar_current: "docs-oci-datasource-functions-pbf_listing_trigger"
description: |-
  Provides details about a specific Pbf Listing Trigger in Oracle Cloud Infrastructure Functions service
---

# Data Source: oci_functions_pbf_listing_trigger
This data source provides details about a specific Pbf Listing Trigger resource in Oracle Cloud Infrastructure Functions service.

Returns a list of Triggers.


## Example Usage

```hcl
data "oci_functions_pbf_listing_trigger" "test_pbf_listing_trigger" {

	#Optional
	name = var.pbf_listing_trigger_name
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) A filter to return only resources that match the service trigger source of a PBF.


## Attributes Reference

The following attributes are exported:

* `items` - List of TriggerSummary.
	* `name` - A brief descriptive name for the PBF trigger.

