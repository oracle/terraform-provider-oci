---
subcategory: "Fleet Software Update"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_software_update_fsu_collection"
sidebar_current: "docs-oci-resource-fleet_software_update-fsu_collection"
description: |-
  Provides the Fsu Collection resource in Oracle Cloud Infrastructure Fleet Software Update service
---

# oci_fleet_software_update_fsu_collection
This resource provides the Fsu Collection resource in Oracle Cloud Infrastructure Fleet Software Update service.

Creates a new Exadata Fleet Update Collection.


## Example Usage

```hcl
resource "oci_fleet_software_update_fsu_collection" "test_fsu_collection" {
	#Required
	compartment_id = var.compartment_id
	service_type = var.fsu_collection_service_type
	source_major_version = var.fsu_collection_source_major_version
	type = var.fsu_collection_type

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	display_name = var.fsu_collection_display_name
	fleet_discovery {
		#Required
		strategy = var.fsu_collection_fleet_discovery_strategy

		#Optional
		filters {
			#Required
			type = var.fsu_collection_fleet_discovery_filters_type

			#Optional
			entity_type = var.fsu_collection_fleet_discovery_filters_entity_type
			identifiers = var.fsu_collection_fleet_discovery_filters_identifiers
			mode = var.fsu_collection_fleet_discovery_filters_mode
			names = var.fsu_collection_fleet_discovery_filters_names
			operator = var.fsu_collection_fleet_discovery_filters_operator
			tags {
				#Required
				key = var.fsu_collection_fleet_discovery_filters_tags_key
				value = var.fsu_collection_fleet_discovery_filters_tags_value

				#Optional
				namespace = var.fsu_collection_fleet_discovery_filters_tags_namespace
			}
			versions = var.fsu_collection_fleet_discovery_filters_versions
		}
		fsu_discovery_id = oci_fleet_software_update_fsu_discovery.test_fsu_discovery.id
		query = var.fsu_collection_fleet_discovery_query
		targets = var.fsu_collection_fleet_discovery_targets
	}
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment Identifier
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Optional) (Updatable) Exadata Fleet Update Collection Identifier. 
* `fleet_discovery` - (Optional) Supported fleet discovery strategies for DB Collections. If specified on an Update Collection request, this will re-discover the targets of the Collection. 
	* `filters` - (Required when strategy=FILTERS) Filters to perform the target discovery. 
		* `entity_type` - (Required when type=RESOURCE_ID) Type of resource to match in the discovery. 
		* `identifiers` - (Required when type=COMPARTMENT_ID | RESOURCE_ID) Related resource Ids to include in the discovery.  All must match the specified entityType. 
		* `mode` - (Applicable when strategy=FILTERS) INCLUDE or EXCLUDE the filter results in the discovery for DB targets. Supported for 'FSUCOLLECTION' RESOURCE_ID filter only. 
		* `names` - (Required when type=DB_HOME_NAME | DB_NAME | DB_UNIQUE_NAME) List of Database unique names to include in the discovery. 
		* `operator` - (Applicable when type=DEFINED_TAG | FREEFORM_TAG | RESOURCE_ID) Type of join for each element in this filter. 
		* `tags` - (Required when type=DEFINED_TAG | FREEFORM_TAG) Freeform tags to include in the discovery. 
			* `key` - (Required) Freeform tag key. 
			* `namespace` - (Required when type=DEFINED_TAG) Defined tag namespace. 
			* `value` - (Required) Freeform tag value. 
		* `type` - (Required) Type of filters supported for Database targets discovery. 
		* `versions` - (Required when type=VERSION) List of Versions strings to include in the discovery. 
	* `fsu_discovery_id` - (Required when strategy=DISCOVERY_RESULTS) OCIDs of Fleet Software Update Discovery. 
	* `query` - (Required when strategy=SEARCH_QUERY) Oracle Cloud Infrastructure Search Service query string. 
	* `strategy` - (Required) Possible fleet discovery strategies. 
	* `targets` - (Required when strategy=TARGET_LIST) OCIDs of target resources to include. For EXACC service type Collections only VMClusters are allowed. For EXACS service type Collections only CloudVMClusters are allowed. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `service_type` - (Required) Exadata service type for the target resource members. 
* `source_major_version` - (Required) Database Major Version of targets to be included in the Exadata Fleet Update Collection. https://docs.oracle.com/en-us/iaas/api/#/en/database/20160918/DbVersionSummary/ListDbVersions Only Database targets that match the version specified in this value would be added to the Exadata Fleet Update Collection. 
* `type` - (Required) Collection type. DB: Only Database entity type resources allowed. GI: CloudVMCluster and VMCluster entity type resources allowed. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Fsu Collection
	* `update` - (Defaults to 20 minutes), when updating the Fsu Collection
	* `delete` - (Defaults to 20 minutes), when destroying the Fsu Collection


## Import

FsuCollections can be imported using the `id`, e.g.

```
$ terraform import oci_fleet_software_update_fsu_collection.test_fsu_collection "id"
```

