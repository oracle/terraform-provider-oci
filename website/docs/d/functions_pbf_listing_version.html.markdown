---
subcategory: "Functions"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_functions_pbf_listing_version"
sidebar_current: "docs-oci-datasource-functions-pbf_listing_version"
description: |-
  Provides details about a specific Pbf Listing Version in Oracle Cloud Infrastructure Functions service
---

# Data Source: oci_functions_pbf_listing_version
This data source provides details about a specific Pbf Listing Version resource in Oracle Cloud Infrastructure Functions service.

Gets a PbfListingVersion by identifier for a PbfListing.

## Example Usage

```hcl
data "oci_functions_pbf_listing_version" "test_pbf_listing_version" {
	#Required
	pbf_listing_version_id = oci_functions_pbf_listing_version.test_pbf_listing_version.id
}
```

## Argument Reference

The following arguments are supported:

* `pbf_listing_version_id` - (Required) unique PbfListingVersion identifier


## Attributes Reference

The following attributes are exported:

* `change_summary` - Details changes are included in this version.
* `config` - Details about the required and optional Function configurations needed for proper performance of the PBF. 
	* `description` - Details about why this config is required and what it will be used for.
	* `is_optional` - Is this a required config or an optional one. Requests with required config params missing will be rejected.
	* `key` - The key name of the config param.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation
* `name` - Semantic version
* `pbf_listing_id` - The OCID of the PbfListing this resource version belongs to.
* `requirements` - Minimum memory required by this PBF. The user should use memory greater than or equal to this value  while configuring the Function. 
	* `min_memory_required_in_mbs` - Minimum memory required by this PBF. The user should use memory greater than or equal to  this value while configuring the Function. 
	* `policies` - List of policies required for this PBF execution.
		* `description` - Details about why this policy is required and what it will be used for.
		* `policy` - Policy required for PBF execution
* `state` - The current state of the PBF resource.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the PbfListingVersion was created. An RFC3339 formatted datetime string.
* `time_updated` - The last time the PbfListingVersion was updated. An RFC3339 formatted datetime string.
* `triggers` - An array of Trigger. A list of triggers that may activate the PBF.
	* `name` - A brief descriptive name for the PBF trigger.

