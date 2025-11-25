---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_log_analytics_entity_associations_list"
sidebar_current: "docs-oci-datasource-log_analytics-log_analytics_entity_associations_list"
description: |-
  Provides the list of Log Analytics Entity Associations List in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_log_analytics_entity_associations_list
This data source provides the list of Log Analytics Entity Associations List in Oracle Cloud Infrastructure Log Analytics service.

Return a list of log analytics entities associated with input source log analytics entity.

## Example Usage

```hcl
data "oci_log_analytics_log_analytics_entity_associations_list" "test_log_analytics_entity_associations_list" {
	#Required
	log_analytics_entity_id = oci_log_analytics_log_analytics_entity.test_log_analytics_entity.id
	namespace = var.log_analytics_entity_associations_list_namespace

	#Optional
	direct_or_all_associations = var.log_analytics_entity_associations_list_direct_or_all_associations
}
```

## Argument Reference

The following arguments are supported:

* `direct_or_all_associations` - (Optional) Indicates whether to return direct associated entities or direct and inferred associated entities. 
* `log_analytics_entity_id` - (Required) The Log analytics entity OCID. 
* `namespace` - (Required) The Log Analytics namespace used for the request. 


## Attributes Reference

The following attributes are exported:

* `log_analytics_entity_collection` - The list of log_analytics_entity_collection.

### LogAnalyticsEntityAssociationsList Reference

The following attributes are exported:

* `items` - Array of log analytics entity summary.
	* `are_logs_collected` - The Boolean flag to indicate if logs are collected for an entity for log analytics usage. 
	* `associated_sources_count` - The count of associated log sources for a given log analytics entity. 
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

