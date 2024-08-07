---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_jms_plugins"
sidebar_current: "docs-oci-datasource-jms-jms_plugins"
description: |-
  Provides the list of Jms Plugins in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_jms_plugins
This data source provides the list of Jms Plugins in Oracle Cloud Infrastructure Jms service.

Lists the JmsPlugins.

## Example Usage

```hcl
data "oci_jms_jms_plugins" "test_jms_plugins" {

	#Optional
	agent_id = var.jms_plugin_agent_id
	availability_status = var.jms_plugin_availability_status
	compartment_id = var.compartment_id
	compartment_id_in_subtree = var.jms_plugin_compartment_id_in_subtree
	fleet_id = oci_jms_fleet.test_fleet.id
	hostname_contains = var.jms_plugin_hostname_contains
	id = var.jms_plugin_id
	state = var.jms_plugin_state
	time_last_seen_less_than_or_equal_to = var.jms_plugin_time_last_seen_less_than_or_equal_to
	time_registered_less_than_or_equal_to = var.jms_plugin_time_registered_less_than_or_equal_to
}
```

## Argument Reference

The following arguments are supported:

* `agent_id` - (Optional) The ManagementAgent (OMA) or Instance (OCA) [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that identifies the Agent.
* `availability_status` - (Optional) Filter JmsPlugin with its availability status.
* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources. 
* `compartment_id_in_subtree` - (Optional) Flag to determine whether the info should be gathered only in the compartment or in the compartment and its subcompartments. 
* `fleet_id` - (Optional) The ID of the Fleet.
* `hostname_contains` - (Optional) Filter the list with hostname contains the given value. 
* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the JmsPlugin.
* `state` - (Optional) Filter JmsPlugin with its lifecycle state.
* `time_last_seen_less_than_or_equal_to` - (Optional) If present, only plugins with a last seen time before this parameter are searched (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).
* `time_registered_less_than_or_equal_to` - (Optional) If present, only plugins with a registration time before this parameter are searched (formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)).


## Attributes Reference

The following attributes are exported:

* `jms_plugin_collection` - The list of jms_plugin_collection.

### JmsPlugin Reference

The following attributes are exported:

* `agent_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Management Agent (OMA) or the Oracle Cloud Agent (OCA) instance where the JMS plugin is deployed.
* `agent_type` - The agent type.
* `availability_status` - The availability status.
* `compartment_id` - The OMA/OCA agent's compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`. (See [Understanding Free-form Tags](https://docs.cloud.oracle.com/iaas/Content/Tagging/Tasks/managingtagsandtagnamespaces.htm)). 
* `fleet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the fleet. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`. (See [Managing Tags and Tag Namespaces](https://docs.cloud.oracle.com/iaas/Content/Tagging/Concepts/understandingfreeformtags.htm).) 
* `hostname` - The hostname of the agent. 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) to identify this JmsPlugin.
* `os_architecture` - The architecture of the operating system of the plugin.
* `os_distribution` - The distribution of the operating system of the plugin.
* `os_family` - The operating system family for the plugin.
* `plugin_version` - The version of the plugin.
* `state` - The lifecycle state.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_last_seen` - The date and time the resource was _last_ reported to JMS. This is potentially _after_ the specified time period provided by the filters. For example, a resource can be last reported to JMS before the start of a specified time period, if it is also reported during the time period. 
* `time_registered` - The date and time the plugin was registered. 

