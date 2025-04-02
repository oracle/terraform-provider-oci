---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_monitored_resource_type"
sidebar_current: "docs-oci-resource-stack_monitoring-monitored_resource_type"
description: |-
  Provides the Monitored Resource Type resource in Oracle Cloud Infrastructure Stack Monitoring service
---

# oci_stack_monitoring_monitored_resource_type
This resource provides the Monitored Resource Type resource in Oracle Cloud Infrastructure Stack Monitoring service.

Creates a new monitored resource type.

## Example Usage

```hcl
resource "oci_stack_monitoring_monitored_resource_type" "test_monitored_resource_type" {
	#Required
	compartment_id = var.compartment_id
	name = var.monitored_resource_type_name

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.monitored_resource_type_description
	display_name = var.monitored_resource_type_display_name
	freeform_tags = {"bar-key"= "value"}
	metadata {
		#Required
		format = var.monitored_resource_type_metadata_format

		#Optional
		agent_properties = var.monitored_resource_type_metadata_agent_properties
		required_properties = var.monitored_resource_type_metadata_required_properties
		unique_property_sets {
			#Required
			properties = var.monitored_resource_type_metadata_unique_property_sets_properties
		}
		valid_properties_for_create = var.monitored_resource_type_metadata_valid_properties_for_create
		valid_properties_for_update = var.monitored_resource_type_metadata_valid_properties_for_update
		valid_property_values = var.monitored_resource_type_metadata_valid_property_values
		valid_sub_resource_types = var.monitored_resource_type_metadata_valid_sub_resource_types
	}
	metric_namespace = var.monitored_resource_type_metric_namespace
	resource_category = var.monitored_resource_type_resource_category
	source_type = var.monitored_resource_type_source_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) Compartment Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) A friendly description.
* `display_name` - (Optional) (Updatable) Monitored resource type display name.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `metadata` - (Optional) (Updatable) The metadata details for resource type.
	* `agent_properties` - (Optional) (Updatable) List of properties needed by the agent for monitoring the resource.  Valid only if resource type is Oracle Cloud Infrastructure management agent based. When specified,  these properties are passed to the management agent during resource create or update. 
	* `format` - (Required) (Updatable) ResourceType metadata format to be used. Currently supports only one format. Possible values - SYSTEM_FORMAT.
		* SYSTEM_FORMAT - The resource type metadata is defined in machine friendly format. 
	* `required_properties` - (Optional) (Updatable) List of required properties for resource type.
	* `unique_property_sets` - (Optional) (Updatable) List of property sets used to uniquely identify the resources.  This check is made during create or update of stack monitoring resource.  The resource has to pass unique check for each set in the list.  For example, database can have user, password and SID as one unique set.  Another unique set would be user, password and service name. 
		* `properties` - (Required) (Updatable) List of properties.
	* `valid_properties_for_create` - (Optional) (Updatable) List of valid properties for resource type while creating the monitored resource.  If resources of this type specifies any other properties during create operation,  the operation will fail. 
	* `valid_properties_for_update` - (Optional) (Updatable) List of valid properties for resource type while updating the monitored resource.  If resources of this type specifies any other properties during update operation,  the operation will fail. 
	* `valid_property_values` - (Optional) (Updatable) List of valid values for the properties. This is useful when resource type wants to restrict only certain values for some properties. For instance for 'osType' property,  supported values can be restricted to be either Linux or Windows. Example: `{ "osType": "Linux,Windows,Solaris"}` 
	* `valid_sub_resource_types` - (Optional) (Updatable) List of valid sub-resource types for a composite resource type. The sub-resource types will be obtained from the valid association pairs corresponding to the composite resource types. It will be empty for non composite resource types 
* `metric_namespace` - (Optional) (Updatable) Metric namespace for resource type.
* `name` - (Required) A unique monitored resource type name. The name must be unique across tenancy.  Name can not be changed. 
* `resource_category` - (Optional) (Updatable) Resource Category to indicate the kind of resource type. 
* `source_type` - (Optional) (Updatable) Source type to indicate if the resource is stack monitoring discovered, Oracle Cloud Infrastructure native resource, etc. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `additional_namespace_map` - Key/Value pair for additional namespaces used by stack monitoring services for SYSTEM (SMB) resource types.
* `availability_metrics_config` - Availability metrics details.
	* `collection_interval_in_seconds` - Availability metric collection internal in seconds.
	* `metrics` - List of metrics used for availability calculation for the resource.
* `compartment_id` - Compartment Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A friendly description.
* `display_name` - Monitored resource type display name.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `handler_config` - Specific resource mapping configurations for Agent Extension Handlers.
	* `collectd_resource_name_config` - Resource name generation overriding configurations for collectd resource types. 
		* `exclude_properties` - List of property names to be excluded.
		* `include_properties` - List of property names to be included.
		* `suffix` - String to be suffixed to the resource name.
	* `collector_types` - List of collector/plugin names.
	* `handler_properties` - List of handler configuration properties
		* `name` - Property name.
		* `value` - Property value.
	* `metric_mappings` - List of AgentExtensionHandlerMetricMappingDetails.
		* `collector_metric_name` - Metric name as defined by the collector.
		* `is_skip_upload` - Is ignoring this metric.
		* `metric_upload_interval_in_seconds` - Metric upload interval in seconds. Any metric sent by telegraf/collectd before the  configured interval expires will be dropped. 
		* `telemetry_metric_name` - Metric name to be upload to telemetry.
	* `metric_name_config` - Metric name generation overriding configurations.
		* `exclude_pattern_on_prefix` - String pattern to be removed from the prefix of the metric name.
		* `is_prefix_with_collector_type` - is prefixing the metric with collector type.
	* `metric_upload_interval_in_seconds` - Metric upload interval in seconds. Any metric sent by telegraf/collectd before the  configured interval expires will be dropped. 
	* `telegraf_resource_name_config` - Resource name generation overriding configurations for telegraf resource types. 
		* `exclude_tags` - List of tag names to be excluded.
		* `include_tags` - List of tag names to be included.
		* `is_use_tags_only` - Flag to indicate if only tags will be used for resource name generation.
	* `telemetry_resource_group` - Resource group string; if not specified, the resource group string will be generated by the handler.
* `id` - Monitored resource type identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `is_system_defined` - If boolean flag is true, then the resource type cannot be modified or deleted.
* `metadata` - The metadata details for resource type.
	* `agent_properties` - List of properties needed by the agent for monitoring the resource.  Valid only if resource type is Oracle Cloud Infrastructure management agent based. When specified,  these properties are passed to the management agent during resource create or update. 
	* `format` - ResourceType metadata format to be used. Currently supports only one format. Possible values - SYSTEM_FORMAT.
		* SYSTEM_FORMAT - The resource type metadata is defined in machine friendly format. 
	* `required_properties` - List of required properties for resource type.
	* `unique_property_sets` - List of property sets used to uniquely identify the resources.  This check is made during create or update of stack monitoring resource.  The resource has to pass unique check for each set in the list.  For example, database can have user, password and SID as one unique set.  Another unique set would be user, password and service name. 
		* `properties` - List of properties.
	* `valid_properties_for_create` - List of valid properties for resource type while creating the monitored resource.  If resources of this type specifies any other properties during create operation,  the operation will fail. 
	* `valid_properties_for_update` - List of valid properties for resource type while updating the monitored resource.  If resources of this type specifies any other properties during update operation,  the operation will fail. 
	* `valid_property_values` - List of valid values for the properties. This is useful when resource type wants to restrict only certain values for some properties. For instance for 'osType' property,  supported values can be restricted to be either Linux or Windows. Example: `{ "osType": ["Linux","Windows","Solaris"]}` 
	* `valid_sub_resource_types` - List of valid sub-resource types for a composite resource type. The sub-resource types will be obtained from the valid association pairs corresponding to the composite resource types. It will be empty for non composite resource types 
* `metric_namespace` - Metric namespace for resource type.
* `name` - A unique monitored resource type name. The name must be unique across tenancy.  Name can not be changed. 
* `resource_category` - Resource Category to indicate the kind of resource type. 
* `source_type` - Source type to indicate if the resource is stack monitoring discovered, Oracle Cloud Infrastructure native resource, etc. 
* `state` - Lifecycle state of the monitored resource type.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `tenancy_id` - Tenancy Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `time_created` - The date and time when the monitored resource type was created, expressed in  [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. 
* `time_updated` - The date and time when the monitored resource was updated, expressed in  [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Monitored Resource Type
	* `update` - (Defaults to 20 minutes), when updating the Monitored Resource Type
	* `delete` - (Defaults to 20 minutes), when destroying the Monitored Resource Type


## Import

MonitoredResourceTypes can be imported using the `id`, e.g.

```
$ terraform import oci_stack_monitoring_monitored_resource_type.test_monitored_resource_type "id"
```

