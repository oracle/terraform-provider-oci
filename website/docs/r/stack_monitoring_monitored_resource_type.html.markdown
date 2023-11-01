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
	}
	metric_namespace = var.monitored_resource_type_metric_namespace
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tenancy containing the resource type. 
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
* `metric_namespace` - (Optional) (Updatable) Metric namespace for resource type.
* `name` - (Required) A unique monitored resource type name. The name must be unique across tenancy.  Name can not be changed. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tenancy containing the resource type. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A friendly description.
* `display_name` - Monitored resource type display name.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Monitored resource type identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `metadata` - The metadata details for resource type.
	* `agent_properties` - List of properties needed by the agent for monitoring the resource.  Valid only if resource type is Oracle Cloud Infrastructure management agent based. When specified,  these properties are passed to the management agent during resource create or update. 
	* `format` - ResourceType metadata format to be used. Currently supports only one format. Possible values - SYSTEM_FORMAT.
		* SYSTEM_FORMAT - The resource type metadata is defined in machine friendly format. 
	* `required_properties` - List of required properties for resource type.
	* `unique_property_sets` - List of property sets used to uniquely identify the resources.  This check is made during create or update of stack monitoring resource.  The resource has to pass unique check for each set in the list.  For example, database can have user, password and SID as one unique set.  Another unique set would be user, password and service name. 
		* `properties` - List of properties.
	* `valid_properties_for_create` - List of valid properties for resource type while creating the monitored resource.  If resources of this type specifies any other properties during create operation,  the operation will fail. 
	* `valid_properties_for_update` - List of valid properties for resource type while updating the monitored resource.  If resources of this type specifies any other properties during update operation,  the operation will fail. 
	* `valid_property_values` - List of valid values for the properties. This is useful when resource type wants to restrict only certain values for some properties. For instance for 'osType' property,  supported values can be restricted to be either Linux or Windows. Example: `{ "osType": "Linux,Windows,Solaris"}` 
* `metric_namespace` - Metric namespace for resource type.
* `name` - A unique monitored resource type name. The name must be unique across tenancy.  Name can not be changed. 
* `state` - Lifecycle state of the monitored resource type.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
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

