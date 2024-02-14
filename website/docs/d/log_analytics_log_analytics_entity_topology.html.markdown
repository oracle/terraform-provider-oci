---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_log_analytics_entity_topology"
sidebar_current: "docs-oci-datasource-log_analytics-log_analytics_entity_topology"
description: |-
  Provides details about a specific Log Analytics Entity Topology in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_log_analytics_entity_topology
This data source provides details about a specific Log Analytics Entity Topology resource in Oracle Cloud Infrastructure Log Analytics service.

Return a log analytics entity topology collection that contains a set of log analytics entities and a set of relationships between those, for the input source entity.

## Example Usage

```hcl
data "oci_log_analytics_log_analytics_entity_topology" "test_log_analytics_entity_topology" {
	#Required
	log_analytics_entity_id = oci_log_analytics_log_analytics_entity.test_log_analytics_entity.id
	namespace = var.log_analytics_entity_topology_namespace

	#Optional
	metadata_equals = var.log_analytics_entity_topology_metadata_equals
	state = var.log_analytics_entity_topology_state
}
```

## Argument Reference

The following arguments are supported:

* `log_analytics_entity_id` - (Required) The log analytics entity OCID. 
* `metadata_equals` - (Optional) A filter to return only log analytics entities whose metadata name, value and type matches the specified string. Each item in the array has the format "{name}:{value}:{type}".  All inputs are case-insensitive. 
* `namespace` - (Required) The Logging Analytics namespace used for the request. 
* `state` - (Optional) A filter to return only those log analytics entities with the specified lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `items` - Array of log analytics entity topologies.
	* `links` - Collection of log analytics entity relationship links. 
		* `items` - Array of log analytics entity relationship links.
			* `destination_entity_id` - The log analytics entity OCID. This ID is a reference used by log analytics features and it represents a resource that is provisioned and managed by the customer on their premises or on the cloud. 
			* `source_entity_id` - The log analytics entity OCID. This ID is a reference used by log analytics features and it represents a resource that is provisioned and managed by the customer on their premises or on the cloud. 
	* `nodes` - Collection of log analytics entities. 
		* `items` - Array of log analytics entity summary.
			* `are_logs_collected` - The Boolean flag to indicate if logs are collected for an entity for log analytics usage. 
			* `cloud_resource_id` - The OCID of the Cloud resource which this entity is a representation of. This may be blank when the entity represents a non-cloud resource that the customer may have on their premises. 
			* `compartment_id` - Compartment Identifier [OCID] (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
			* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
			* `entity_type_internal_name` - Internal name for the log analytics entity type. 
			* `entity_type_name` - Log analytics entity type name. 
			* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
			* `id` - The log analytics entity OCID. This ID is a reference used by log analytics features and it represents a resource that is provisioned and managed by the customer on their premises or on the cloud. 
			* `lifecycle_details` - lifecycleDetails has additional information regarding substeps such as management agent plugin deployment. 
			* `management_agent_id` - The OCID of the Management Agent. 
			* `metadata` - A collection of entity metadata information.
				* `items` - An array of entity metadata.
					* `name` - The metadata name.
					* `type` - The metadata type.
					* `value` - The metadata value.
			* `name` - Log analytics entity name. 
			* `source_id` - This indicates the type of source. It is primarily for Enterprise Manager Repository ID. 
			* `state` - The current state of the log analytics entity. 
			* `time_created` - The date and time the resource was created, in the format defined by RFC3339. 
			* `time_last_discovered` - The date and time the resource was last discovered, in the format defined by RFC3339. 
			* `time_updated` - The date and time the resource was last updated, in the format defined by RFC3339. 
			* `timezone_region` - The timezone region of the log analytics entity. 

