---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_log_analytics_entities"
sidebar_current: "docs-oci-datasource-log_analytics-log_analytics_entities"
description: |-
  Provides the list of Log Analytics Entities in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_log_analytics_entities
This data source provides the list of Log Analytics Entities in Oracle Cloud Infrastructure Log Analytics service.

Return a list of log analytics entities.

## Example Usage

```hcl
data "oci_log_analytics_log_analytics_entities" "test_log_analytics_entities" {
	#Required
	compartment_id = var.compartment_id
	namespace = var.log_analytics_entity_namespace

	#Optional
	cloud_resource_id = var.log_analytics_entity_cloud_resource_id
	defined_tag_equals = var.log_analytics_entity_defined_tag_equals
	defined_tag_exists = var.log_analytics_entity_defined_tag_exists
	entity_type_name = var.log_analytics_entity_entity_type_name
	freeform_tag_equals = var.log_analytics_entity_freeform_tag_equals
	freeform_tag_exists = var.log_analytics_entity_freeform_tag_exists
	hostname = var.log_analytics_entity_hostname
	hostname_contains = var.log_analytics_entity_hostname_contains
	is_management_agent_id_null = var.log_analytics_entity_is_management_agent_id_null
	is_show_associated_sources_count = var.log_analytics_entity_is_show_associated_sources_count
	lifecycle_details_contains = var.log_analytics_entity_lifecycle_details_contains
	metadata_equals = var.log_analytics_entity_metadata_equals
	name = var.log_analytics_entity_name
	name_contains = var.log_analytics_entity_name_contains
	source_id = oci_log_analytics_source.test_source.id
	state = var.log_analytics_entity_state
}
```

## Argument Reference

The following arguments are supported:

* `cloud_resource_id` - (Optional) A filter to return only log analytics entities whose cloudResourceId matches the cloudResourceId given. 
* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `defined_tag_equals` - (Optional) A list of tag filters to apply.  Only entities with a defined tag matching the value will be returned. Each item in the list has the format "{namespace}.{tagName}.{value}".  All inputs are case-insensitive. Multiple values for the same key (i.e. same namespace and tag name) are interpreted as "OR". Values for different keys (i.e. different namespaces, different tag names, or both) are interpreted as "AND". 
* `defined_tag_exists` - (Optional) A list of tag existence filters to apply.  Only entities for which the specified defined tags exist will be returned. Each item in the list has the format "{namespace}.{tagName}.true" (for checking existence of a defined tag) or "{namespace}.true".  All inputs are case-insensitive. Currently, only existence ("true" at the end) is supported. Absence ("false" at the end) is not supported. Multiple values for the same key (i.e. same namespace and tag name) are interpreted as "OR". Values for different keys (i.e. different namespaces, different tag names, or both) are interpreted as "AND". 
* `entity_type_name` - (Optional) A filter to return only log analytics entities whose entityTypeName matches the entire log analytics entity type name of one of the entityTypeNames given in the list. The match is case-insensitive. 
* `freeform_tag_equals` - (Optional) A list of tag filters to apply.  Only entities with a freeform tag matching the value will be returned. The key for each tag is "{tagName}.{value}".  All inputs are case-insensitive. Multiple values for the same tag name are interpreted as "OR".  Values for different tag names are interpreted as "AND". 
* `freeform_tag_exists` - (Optional) A list of tag existence filters to apply.  Only entities for which the specified freeform tags exist the value will be returned. The key for each tag is "{tagName}.true".  All inputs are case-insensitive. Currently, only existence ("true" at the end) is supported. Absence ("false" at the end) is not supported. Multiple values for different tag names are interpreted as "AND". 
* `hostname` - (Optional) A filter to return only log analytics entities whose hostname matches the entire hostname given. 
* `hostname_contains` - (Optional) A filter to return only log analytics entities whose hostname contains the substring given. The match is case-insensitive. 
* `is_management_agent_id_null` - (Optional) A filter to return only those log analytics entities whose managementAgentId is null or is not null. 
* `is_show_associated_sources_count` - (Optional) Option to return count of associated log sources for log analytics entity(s).
* `lifecycle_details_contains` - (Optional) A filter to return only log analytics entities whose lifecycleDetails contains the specified string. 
* `metadata_equals` - (Optional) A filter to return only log analytics entities whose metadata name, value and type matches the specified string. Each item in the array has the format "{name}:{value}:{type}".  All inputs are case-insensitive. 
* `name` - (Optional) A filter to return only log analytics entities whose name matches the entire name given. The match is case-insensitive. 
* `name_contains` - (Optional) A filter to return only log analytics entities whose name contains the name given. The match is case-insensitive. 
* `namespace` - (Required) The Logging Analytics namespace used for the request. 
* `source_id` - (Optional) A filter to return only log analytics entities whose sourceId matches the sourceId given. 
* `state` - (Optional) A filter to return only those log analytics entities with the specified lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `log_analytics_entity_collection` - The list of log_analytics_entity_collection.

### LogAnalyticsEntity Reference

The following attributes are exported:

* `are_logs_collected` - The Boolean flag to indicate if logs are collected for an entity for log analytics usage. 
* `associated_sources_count` - The count of associated log sources for a given log analytics entity. 
* `cloud_resource_id` - The OCID of the Cloud resource which this entity is a representation of. This may be blank when the entity represents a non-cloud resource that the customer may have on their premises. 
* `compartment_id` - Compartment Identifier [OCID] (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `entity_type_internal_name` - Internal name for the log analytics entity type. 
* `entity_type_name` - Log analytics entity type name. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `hostname` - The hostname where the entity represented here is actually present. This would be the output one would get if they run `echo $HOSTNAME` on Linux or an equivalent OS command. This may be different from management agents host since logs may be collected remotely. 
* `id` - The log analytics entity OCID. This ID is a reference used by log analytics features and it represents a resource that is provisioned and managed by the customer on their premises or on the cloud. 
* `lifecycle_details` - lifecycleDetails has additional information regarding substeps such as management agent plugin deployment. 
* `management_agent_compartment_id` - Management agent (management-agents resource kind) compartment OCID 
* `management_agent_display_name` - Management agent (management-agents resource kind) display name 
* `management_agent_id` - The OCID of the Management Agent. 
* `metadata` - Details of entity metadata information.
	* `items` - An array of entity metadata.
		* `name` - The metadata name.
		* `type` - The metadata type.
		* `value` - The metadata value.
* `name` - Log analytics entity name. 
* `properties` - The name/value pairs for parameter values to be used in file patterns specified in log sources. 
* `source_id` - This indicates the type of source. It is primarily for Enterprise Manager Repository ID. 
* `state` - The current state of the log analytics entity. 
* `time_created` - The date and time the resource was created, in the format defined by RFC3339. 
* `time_last_discovered` - The date and time the resource was last discovered, in the format defined by RFC3339. 
* `time_updated` - The date and time the resource was last updated, in the format defined by RFC3339. 
* `timezone_region` - The timezone region of the log analytics entity. 

