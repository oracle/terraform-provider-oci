---
subcategory: "Fleet Software Update"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_software_update_fsu_collection"
sidebar_current: "docs-oci-datasource-fleet_software_update-fsu_collection"
description: |-
  Provides details about a specific Fsu Collection in Oracle Cloud Infrastructure Fleet Software Update service
---

# Data Source: oci_fleet_software_update_fsu_collection
This data source provides details about a specific Fsu Collection resource in Oracle Cloud Infrastructure Fleet Software Update service.

Gets a Exadata Fleet Update Collection by identifier.


## Example Usage

```hcl
data "oci_fleet_software_update_fsu_collection" "test_fsu_collection" {
	#Required
	fsu_collection_id = oci_fleet_software_update_fsu_collection.test_fsu_collection.id
}
```

## Argument Reference

The following arguments are supported:

* `fsu_collection_id` - (Required) Unique Exadata Fleet Update Collection identifier. 


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

