---
subcategory: "Fleet Software Update"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_software_update_fsu_collections"
sidebar_current: "docs-oci-datasource-fleet_software_update-fsu_collections"
description: |-
  Provides the list of Fsu Collections in Oracle Cloud Infrastructure Fleet Software Update service
---

# Data Source: oci_fleet_software_update_fsu_collections
This data source provides the list of Fsu Collections in Oracle Cloud Infrastructure Fleet Software Update service.

Gets a list of all Exadata Fleet Update Collections in a compartment.


## Example Usage

```hcl
data "oci_fleet_software_update_fsu_collections" "test_fsu_collections" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.fsu_collection_display_name
	state = var.fsu_collection_state
	type = var.fsu_collection_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources. 
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. 
* `state` - (Optional) A filter to return only resources whose lifecycleState matches the given lifecycleState. 
* `type` - (Optional) A filter to return only resources whose type matches the given type. 


## Attributes Reference

The following attributes are exported:

* `fsu_collection_summary_collection` - The list of fsu_collection_summary_collection.

### FsuCollection Reference

The following attributes are exported:

* `active_fsu_cycle` - Active Exadata Fleet Update Cycle resource for this Collection. Object would be null if there is no active Cycle. 
	* `display_name` - Display name of the active Exadata Fleet Update Cycle resource. 
	* `id` - OCID of the active Exadata Fleet Update Cycle resource. 
* `compartment_id` - Compartment Identifier 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Exadata Fleet Update Collection resource display name. 
* `fleet_discovery` - Supported fleet discovery strategies for DB Collections. If specified on an Update Collection request, this will re-discover the targets of the Collection. 
	* `filters` - Filters to perform the target discovery. 
		* `entity_type` - Type of resource to match in the discovery. 
		* `identifiers` - Related resource Ids to include in the discovery.  All must match the specified entityType. 
		* `mode` - INCLUDE or EXCLUDE the filter results in the discovery for DB targets. Supported for 'FSUCOLLECTION' RESOURCE_ID filter only. 
		* `names` - List of Database unique names to include in the discovery. 
		* `operator` - Type of join for each element in this filter. 
		* `tags` - Freeform tags to include in the discovery. 
			* `key` - Freeform tag key. 
			* `namespace` - Defined tag namespace. 
			* `value` - Freeform tag value. 
		* `type` - Type of filters supported for Database targets discovery. 
		* `versions` - List of Versions strings to include in the discovery. 
	* `fsu_discovery_id` - OCIDs of Fleet Software Update Discovery. 
	* `query` - Oracle Cloud Infrastructure Search Service query string. 
	* `strategy` - Possible fleet discovery strategies. 
	* `targets` - OCIDs of target resources to include. For EXACC service type Collections only VMClusters are allowed. For EXACS service type Collections only CloudVMClusters are allowed. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - OCID identifier for the Exadata Fleet Update Collection. 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state. 
* `service_type` - Exadata service type for the target resource members. 
* `source_major_version` - Database Major Version of targets to be included in the Exadata Fleet Update Collection. https://docs.oracle.com/en-us/iaas/api/#/en/database/20160918/DbVersionSummary/ListDbVersions Only Database targets that match the version specified in this value would be added to the Exadata Fleet Update Collection. 
* `state` - The current state of the Exadata Fleet Update Collection.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_count` - Number of targets that are members of this Collection. 
* `time_created` - The time the Exadata Fleet Update Collection was created. An RFC3339 formatted datetime string. 
* `time_updated` - The time the Exadata Fleet Update Collection was updated. An RFC3339 formatted datetime string. 
* `type` - Exadata Fleet Update Collection type. 

