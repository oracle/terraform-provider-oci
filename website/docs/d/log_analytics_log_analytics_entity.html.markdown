---
subcategory: "Log Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_log_analytics_log_analytics_entity"
sidebar_current: "docs-oci-datasource-log_analytics-log_analytics_entity"
description: |-
  Provides details about a specific Log Analytics Entity in Oracle Cloud Infrastructure Log Analytics service
---

# Data Source: oci_log_analytics_log_analytics_entity
This data source provides details about a specific Log Analytics Entity resource in Oracle Cloud Infrastructure Log Analytics service.

Retrieve the log analytics entity with the given id.

## Example Usage

```hcl
data "oci_log_analytics_log_analytics_entity" "test_log_analytics_entity" {
	#Required
	log_analytics_entity_id = oci_log_analytics_log_analytics_entity.test_log_analytics_entity.id
	namespace = var.log_analytics_entity_namespace
}
```

## Argument Reference

The following arguments are supported:

* `log_analytics_entity_id` - (Required) The log analytics entity OCID. 
* `namespace` - (Required) The Logging Analytics namespace used for the request. 


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
* `name` - Log analytics entity name. 
* `properties` - The name/value pairs for parameter values to be used in file patterns specified in log sources. 
* `source_id` - This indicates the type of source. It is primarily for Enterprise Manager Repository ID. 
* `state` - The current state of the log analytics entity. 
* `time_created` - The date and time the resource was created, in the format defined by RFC3339. 
* `time_updated` - The date and time the resource was last updated, in the format defined by RFC3339. 
* `timezone_region` - The timezone region of the log analytics entity. 

