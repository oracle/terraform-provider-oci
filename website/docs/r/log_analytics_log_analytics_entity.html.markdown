---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_log_analytics_entity"
sidebar_current: "docs-oci-resource-log_analytics-log_analytics_entity"
description: |-
  Provides the Log Analytics Entity resource in Oracle Cloud Infrastructure Log Analytics service
---

# oci_log_analytics_log_analytics_entity
This resource provides the Log Analytics Entity resource in Oracle Cloud Infrastructure Log Analytics service.

Create a new log analytics entity.

## Example Usage

```hcl
resource "oci_log_analytics_log_analytics_entity" "test_log_analytics_entity" {
	#Required
	compartment_id = var.compartment_id
	entity_type_name = var.log_analytics_entity_entity_type_name
	name = var.log_analytics_entity_name
	namespace = var.log_analytics_entity_namespace

	#Optional
	cloud_resource_id = oci_log_analytics_cloud_resource.test_cloud_resource.id
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	hostname = var.log_analytics_entity_hostname
	management_agent_id = oci_management_agent_management_agent.test_management_agent.id
	metadata {

		#Optional
		items {

			#Optional
			name = var.log_analytics_entity_metadata_items_name
			type = var.log_analytics_entity_metadata_items_type
			value = var.log_analytics_entity_metadata_items_value
		}
	}
	properties = var.log_analytics_entity_properties
	source_id = oci_log_analytics_source.test_source.id
	time_last_discovered = var.log_analytics_entity_time_last_discovered
	timezone_region = var.log_analytics_entity_timezone_region
}
```

## Argument Reference

The following arguments are supported:

* `cloud_resource_id` - (Optional) The OCID of the Cloud resource which this entity is a representation of. This may be blank when the entity represents a non-cloud resource that the customer may have on their premises. 
* `compartment_id` - (Required) (Updatable) Compartment Identifier [OCID] (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `entity_type_name` - (Required) Log analytics entity type name. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `hostname` - (Optional) (Updatable) The hostname where the entity represented here is actually present. This would be the output one would get if they run `echo $HOSTNAME` on Linux or an equivalent OS command. This may be different from management agents host since logs may be collected remotely. 
* `management_agent_id` - (Optional) (Updatable) The OCID of the Management Agent. 
* `metadata` - (Optional) (Updatable) Details of Entity Metadata.
	* `items` - (Optional) (Updatable) An array of entity metadata details.
		* `name` - (Optional) (Updatable) The metadata name.
		* `type` - (Optional) (Updatable) The metadata type.
		* `value` - (Optional) (Updatable) The metadata value.
* `name` - (Required) (Updatable) Log analytics entity name. 
* `namespace` - (Required) The Logging Analytics namespace used for the request. 
* `properties` - (Optional) (Updatable) The name/value pairs for parameter values to be used in file patterns specified in log sources. 
* `source_id` - (Optional) This indicates the type of source. It is primarily for Enterprise Manager Repository ID. 
* `time_last_discovered` - (Optional) (Updatable) The date and time the resource was last discovered, in the format defined by RFC3339. 
* `timezone_region` - (Optional) (Updatable) The timezone region of the log analytics entity. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `are_logs_collected` - The Boolean flag to indicate if logs are collected for an entity for log analytics usage. 
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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Log Analytics Entity
	* `update` - (Defaults to 20 minutes), when updating the Log Analytics Entity
	* `delete` - (Defaults to 20 minutes), when destroying the Log Analytics Entity


## Import

LogAnalyticsEntities can be imported using the `id`, e.g.

```
$ terraform import oci_log_analytics_log_analytics_entity.test_log_analytics_entity "namespaces/{namespaceName}/logAnalyticsEntities/{logAnalyticsEntityId}" 
```

