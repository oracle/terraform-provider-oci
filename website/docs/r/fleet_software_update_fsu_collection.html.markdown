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
	type = var.fsu_collection_type

	#Optional
	components {
		#Required
		component_type = var.fsu_collection_components_component_type
		source_major_version = var.fsu_collection_components_source_major_version

		#Optional
		fleet_discovery {
			#Required
			strategy = var.fsu_collection_components_fleet_discovery_strategy

			#Optional
			filters {
				#Required
				type = var.fsu_collection_components_fleet_discovery_filters_type

				#Optional
				entity_type = var.fsu_collection_components_fleet_discovery_filters_entity_type
				exadata_releases = var.fsu_collection_components_fleet_discovery_filters_exadata_releases
				identifiers = var.fsu_collection_components_fleet_discovery_filters_identifiers
				mode = var.fsu_collection_components_fleet_discovery_filters_mode
				operator = var.fsu_collection_components_fleet_discovery_filters_operator
				tags {

					#Optional
					key = var.fsu_collection_components_fleet_discovery_filters_tags_key
					namespace = var.fsu_collection_components_fleet_discovery_filters_tags_namespace
					value = var.fsu_collection_components_fleet_discovery_filters_tags_value
				}
				versions = var.fsu_collection_components_fleet_discovery_filters_versions
			}
			fsu_discovery_id = oci_fleet_software_update_fsu_discovery.test_fsu_discovery.id
			query = var.fsu_collection_components_fleet_discovery_query
			targets = var.fsu_collection_components_fleet_discovery_targets
		}
	}
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
			exadata_releases = var.fsu_collection_fleet_discovery_filters_exadata_releases
			identifiers = var.fsu_collection_fleet_discovery_filters_identifiers
			mode = var.fsu_collection_fleet_discovery_filters_mode
			names = var.fsu_collection_fleet_discovery_filters_names
			operator = var.fsu_collection_fleet_discovery_filters_operator
			tags {

				#Optional
				key = var.fsu_collection_fleet_discovery_filters_tags_key
				namespace = var.fsu_collection_fleet_discovery_filters_tags_namespace
				value = var.fsu_collection_fleet_discovery_filters_tags_value
			}
			versions = var.fsu_collection_fleet_discovery_filters_versions
		}
		fsu_discovery_id = oci_fleet_software_update_fsu_discovery.test_fsu_discovery.id
		query = var.fsu_collection_fleet_discovery_query
		targets = var.fsu_collection_fleet_discovery_targets
	}
	freeform_tags = {"bar-key"= "value"}
	source_major_version = var.fsu_collection_source_major_version
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Compartment.
* `components` - (Required when type=EXADB_STACK) Details of components in an Exadata software stack. 
	* `component_type` - (Required) Type of component in an Exadata software stack. 
	* `fleet_discovery` - (Optional) Fleet discovery strategies for a 'GUEST_OS' collection of Exadata VM Clusters. If specified for an UpdateCollection request, discovery for Exadata VM Clusters will be rerun. 
		* `filters` - (Required when strategy=FILTERS) Filters to perform the target discovery. 
			* `entity_type` - (Required when type=RESOURCE_ID) Type of resource to match in the discovery. 
			* `exadata_releases` - (Required when type=EXADATA_RELEASE_VERSION) List of Exadata Release versions to include when discovering Exadata VM Cluster targets for a 'GUEST_OS' collection. 
			* `identifiers` - (Required when type=COMPARTMENT_ID | RESOURCE_ID) The [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of associated resources to include when discovering Exadata VM Cluster targets for a 'GUEST_OS' collection.  Specified resources must match the specified 'entityType'. FsuCollection of type 'GI' or 'GUEST_OS' can be specified. 
			* `mode` - (Applicable when strategy=FILTERS) INCLUDE or EXCLUDE the filter results when discovering Exadata VM Cluster targets for a 'GUEST_OS' collection. Supported only for RESOURCE_ID filter. 
			* `operator` - (Applicable when type=DEFINED_TAG | FREEFORM_TAG | RESOURCE_ID) Type of join for each element in this filter. 
			* `tags` - (Required when type=DEFINED_TAG | FREEFORM_TAG) [Free-form tags](https://docs.cloud.oracle.com/iaas/Content/Tagging/Concepts/understandingfreeformtags.htm) to include when discovering Exadata VM Cluster targets for a 'GUEST_OS' collection. 
				* `key` - (Required when type=DEFINED_TAG | FREEFORM_TAG) Freeform tag key. 
				* `namespace` - (Required when type=DEFINED_TAG) Defined tag namespace. 
				* `value` - (Required when type=DEFINED_TAG | FREEFORM_TAG) Freeform tag value. 
			* `type` - (Required) Filters supported for searching Exadata VM Cluster targets for a 'GUEST_OS' collection. 
			* `versions` - (Required when type=VERSION) List of Exadata Image (Guest OS) version strings to include when discovering Exadata VM Cluster targets for a 'GUEST_OS' collection. 
		* `fsu_discovery_id` - (Required when strategy=DISCOVERY_RESULTS) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Fleet Update Discovery. 
		* `query` - (Required when strategy=SEARCH_QUERY) [OCI Search Service](https://docs.cloud.oracle.com/iaas/Content/Search/Concepts/queryoverview.htm) query string. 
		* `strategy` - (Required) Supported fleet discovery strategies. 
		* `targets` - (Required when strategy=TARGET_LIST) The [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of Exadata VM Cluster targets. Only Exadata VM Cluster targets associated with the specified 'serviceType' are allowed. 
	* `source_major_version` - (Required) Major version of Exadata Image (Guest OS) release for Exadata VM Cluster targets to be included in an Exadata Fleet Update Collection. Major Versions of Exadata Software are demarcated by the underlying Oracle Linux OS version. For more details, refer to [Oracle document 2075007.1](https://support.oracle.com/knowledge/Oracle%20Database%20Products/2075007_1.html) 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Optional) (Updatable) The user-friendly name for the Exadata Fleet Update Collection. 
* `fleet_discovery` - (Applicable when type=DB | GI | GUEST_OS) Fleet discovery strategies for a 'GUEST_OS' collection of Exadata VM Clusters. If specified for an UpdateCollection request, discovery for Exadata VM Clusters will be rerun. 
	* `filters` - (Required when strategy=FILTERS) Filters to perform the target discovery. 
		* `entity_type` - (Required when type=RESOURCE_ID) Type of resource to match in the discovery. 
		* `exadata_releases` - (Required when type=EXADATA_RELEASE_VERSION) List of Exadata Release versions to include when discovering Exadata VM Cluster targets for a 'GUEST_OS' collection. 
		* `identifiers` - (Required when type=COMPARTMENT_ID | RESOURCE_ID) The [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of associated resources to include when discovering Exadata VM Cluster targets for a 'GUEST_OS' collection.  Specified resources must match the specified 'entityType'. FsuCollection of type 'GI' or 'GUEST_OS' can be specified. 
		* `mode` - (Applicable when strategy=FILTERS) INCLUDE or EXCLUDE the filter results when discovering Exadata VM Cluster targets for a 'GUEST_OS' collection. Supported only for RESOURCE_ID filter. 
		* `names` - (Required when type=DB_HOME_NAME | DB_NAME | DB_UNIQUE_NAME) List of Database unique names to include in the discovery. 
		* `operator` - (Applicable when type=DEFINED_TAG | FREEFORM_TAG | RESOURCE_ID) Type of join for each element in this filter. 
		* `tags` - (Required when type=DEFINED_TAG | FREEFORM_TAG) [Free-form tags](https://docs.cloud.oracle.com/iaas/Content/Tagging/Concepts/understandingfreeformtags.htm) to include when discovering Exadata VM Cluster targets for a 'GUEST_OS' collection. 
			* `key` - (Required when type=DEFINED_TAG | FREEFORM_TAG) Freeform tag key. 
			* `namespace` - (Required when type=DEFINED_TAG) Defined tag namespace. 
			* `value` - (Required when type=DEFINED_TAG | FREEFORM_TAG) Freeform tag value. 
		* `type` - (Required) Filters supported for searching Exadata VM Cluster targets for a 'GUEST_OS' collection. 
		* `versions` - (Required when type=VERSION) List of Exadata Image (Guest OS) version strings to include when discovering Exadata VM Cluster targets for a 'GUEST_OS' collection. 
	* `fsu_discovery_id` - (Required when strategy=DISCOVERY_RESULTS) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Fleet Update Discovery. 
	* `query` - (Required when strategy=SEARCH_QUERY) [OCI Search Service](https://docs.cloud.oracle.com/iaas/Content/Search/Concepts/queryoverview.htm) query string. 
	* `strategy` - (Required) Supported fleet discovery strategies. 
	* `targets` - (Required when strategy=TARGET_LIST) The [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of Exadata VM Cluster targets. Only Exadata VM Cluster targets associated with the specified 'serviceType' are allowed. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `service_type` - (Required) Exadata service type for the target resource members. 
* `source_major_version` - (Required when type=DB | GI | GUEST_OS) Major version of Exadata Image (Guest OS) release for Exadata VM Cluster targets to be included in the Exadata Fleet Update Collection. Only Exadata VM Clusters whose 'systemVersion' is related to the major version will be added to the Exadata Fleet Update Collection. For more details, refer to [Oracle document 2075007.1](https://support.oracle.com/knowledge/Oracle%20Database%20Products/2075007_1.html) 
* `type` - (Required) Collection type. DB: Only Database entity type resources allowed. GI: CloudVMCluster and VMCluster entity type resources allowed. GUEST_OS: CloudVmCluster and VmCluster entity type resources are allowed. EXADB_STACK: CloudVmCluster and VmCluster entity type resources are allowed. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `active_fsu_cycle` - Active Exadata Fleet Update Cycle resource for this Collection. Object would be null if there is no active Cycle. 
	* `display_name` - Display name of the active Exadata Fleet Update Cycle resource. 
	* `id` - OCID of the active Exadata Fleet Update Cycle resource. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Compartment. 
* `components` - Details of components in an Exadata software stack. 
	* `component_type` - Type of component in an Exadata software stack. 
	* `fleet_discovery` - Fleet discovery strategies for a 'GUEST_OS' collection of Exadata VM Clusters. If specified for an UpdateCollection request, discovery for Exadata VM Clusters will be rerun. 
		* `filters` - Filters to perform the target discovery. 
			* `entity_type` - Type of resource to match in the discovery. 
			* `exadata_releases` - List of Exadata Release versions to include when discovering Exadata VM Cluster targets for a 'GUEST_OS' collection. 
			* `identifiers` - The [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of associated resources to include when discovering Exadata VM Cluster targets for a 'GUEST_OS' collection.  Specified resources must match the specified 'entityType'. FsuCollection of type 'GI' or 'GUEST_OS' can be specified. 
			* `mode` - INCLUDE or EXCLUDE the filter results when discovering Exadata VM Cluster targets for a 'GUEST_OS' collection. Supported only for RESOURCE_ID filter. 
			* `operator` - Type of join for each element in this filter. 
			* `tags` - [Free-form tags](https://docs.cloud.oracle.com/iaas/Content/Tagging/Concepts/understandingfreeformtags.htm) to include when discovering Exadata VM Cluster targets for a 'GUEST_OS' collection. 
				* `key` - Freeform tag key. 
				* `namespace` - Defined tag namespace. 
				* `value` - Freeform tag value. 
			* `type` - Filters supported for searching Exadata VM Cluster targets for a 'GUEST_OS' collection. 
			* `versions` - List of Exadata Image (Guest OS) version strings to include when discovering Exadata VM Cluster targets for a 'GUEST_OS' collection. 
		* `fsu_discovery_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Fleet Update Discovery. 
		* `query` - [OCI Search Service](https://docs.cloud.oracle.com/iaas/Content/Search/Concepts/queryoverview.htm) query string. 
		* `strategy` - Supported fleet discovery strategies. 
		* `targets` - The [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of Exadata VM Cluster targets. Only Exadata VM Cluster targets associated with the specified 'serviceType' are allowed. 
	* `source_major_version` - Major version of Exadata Image (Guest OS) release for Exadata VM Cluster targets to be included in an Exadata Fleet Update Collection. Major Versions of Exadata Software are demarcated by the underlying Oracle Linux OS version. For more details, refer to [Oracle document 2075007.1](https://support.oracle.com/knowledge/Oracle%20Database%20Products/2075007_1.html) 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - The user-friendly name for the Exadata Fleet Update Collection. 
* `fleet_discovery` - Fleet discovery strategies for a 'GUEST_OS' collection of Exadata VM Clusters. If specified for an UpdateCollection request, discovery for Exadata VM Clusters will be rerun. 
	* `filters` - Filters to perform the target discovery. 
		* `entity_type` - Type of resource to match in the discovery. 
		* `exadata_releases` - List of Exadata Release versions to include when discovering Exadata VM Cluster targets for a 'GUEST_OS' collection. 
		* `identifiers` - The [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of associated resources to include when discovering Exadata VM Cluster targets for a 'GUEST_OS' collection.  Specified resources must match the specified 'entityType'. FsuCollection of type 'GI' or 'GUEST_OS' can be specified. 
		* `mode` - INCLUDE or EXCLUDE the filter results when discovering Exadata VM Cluster targets for a 'GUEST_OS' collection. Supported only for RESOURCE_ID filter. 
		* `names` - List of Database unique names to include in the discovery. 
		* `operator` - Type of join for each element in this filter. 
		* `tags` - [Free-form tags](https://docs.cloud.oracle.com/iaas/Content/Tagging/Concepts/understandingfreeformtags.htm) to include when discovering Exadata VM Cluster targets for a 'GUEST_OS' collection. 
			* `key` - Freeform tag key. 
			* `namespace` - Defined tag namespace. 
			* `value` - Freeform tag value. 
		* `type` - Filters supported for searching Exadata VM Cluster targets for a 'GUEST_OS' collection. 
		* `versions` - List of Exadata Image (Guest OS) version strings to include when discovering Exadata VM Cluster targets for a 'GUEST_OS' collection. 
	* `fsu_discovery_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Fleet Update Discovery. 
	* `query` - [OCI Search Service](https://docs.cloud.oracle.com/iaas/Content/Search/Concepts/queryoverview.htm) query string. 
	* `strategy` - Supported fleet discovery strategies. 
	* `targets` - The [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of Exadata VM Cluster targets. Only Exadata VM Cluster targets associated with the specified 'serviceType' are allowed. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Fleet Update Collection. 
* `last_completed_fsu_cycle_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of last completed FSU Cycle. 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state. 
* `service_type` - Exadata service type for the target resource members. 
* `source_major_version` - Major version of Exadata Image (Guest OS) release for Exadata VM Cluster targets to be included in the Exadata Fleet Update Collection. Only Exadata VM Clusters whose 'systemVersion' is related to the major version will be added to the Exadata Fleet Update Collection. For more details, refer to [Oracle document 2075007.1](https://support.oracle.com/knowledge/Oracle%20Database%20Products/2075007_1.html) 
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

