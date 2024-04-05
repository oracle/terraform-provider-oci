---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_monitored_resource_type"
sidebar_current: "docs-oci-datasource-stack_monitoring-monitored_resource_type"
description: |-
  Provides details about a specific Monitored Resource Type in Oracle Cloud Infrastructure Stack Monitoring service
---

# Data Source: oci_stack_monitoring_monitored_resource_type
This data source provides details about a specific Monitored Resource Type resource in Oracle Cloud Infrastructure Stack Monitoring service.

Gets a monitored resource type by identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).

## Example Usage

```hcl
data "oci_stack_monitoring_monitored_resource_type" "test_monitored_resource_type" {
	#Required
	monitored_resource_type_id = oci_stack_monitoring_monitored_resource_type.test_monitored_resource_type.id
}
```

## Argument Reference

The following arguments are supported:

* `monitored_resource_type_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of monitored resource type.


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
	* `valid_property_values` - List of valid values for the properties. This is useful when resource type wants to restrict only certain values for some properties. For instance for 'osType' property,  supported values can be restricted to be either Linux or Windows. Example: `{"osType": "Linux,Windows,Solaris", "osVersion": "v6.0,v7.0"}` 
* `metric_namespace` - Metric namespace for resource type.
* `name` - A unique monitored resource type name. The name must be unique across tenancy.  Name can not be changed. 
* `resource_category` - Resource Category to indicate the kind of resource type. 
* `source_type` - Source type to indicate if the resource is stack monitoring discovered, Oracle Cloud Infrastructure native resource, etc. 
* `state` - Lifecycle state of the monitored resource type.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time when the monitored resource type was created, expressed in  [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. 
* `time_updated` - The date and time when the monitored resource was updated, expressed in  [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. 

