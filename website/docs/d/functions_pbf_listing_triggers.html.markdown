---
subcategory: "Functions"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_functions_pbf_listing_triggers"
sidebar_current: "docs-oci-datasource-functions-pbf_listing_triggers"
description: |-
  Provides the list of Pbf Listing Triggers in Oracle Cloud Infrastructure Functions service
---

# Data Source: oci_functions_pbf_listing_triggers
This data source provides the list of Pbf Listing Triggers in Oracle Cloud Infrastructure Functions service.

Returns a list of Triggers.


## Example Usage

```hcl
data "oci_functions_pbf_listing_triggers" "test_pbf_listing_triggers" {

	#Optional
	name = var.pbf_listing_trigger_name
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) A filter to return only resources that match the service trigger source of a PBF.


## Attributes Reference

The following attributes are exported:

* `triggers_collection` - The list of triggers_collection.

### PbfListingTrigger Reference

The following attributes are exported:

* `items` - List of TriggerSummary.
	* `name` - A brief descriptive name for the PBF trigger.

