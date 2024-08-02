---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_jms_plugin"
sidebar_current: "docs-oci-resource-jms-jms_plugin"
description: |-
  Provides the Jms Plugin resource in Oracle Cloud Infrastructure Jms service
---

# oci_jms_jms_plugin
This resource provides the Jms Plugin resource in Oracle Cloud Infrastructure Jms service.

Registers an agent's JmsPlugin, optionally attaching to an existing fleet of the tenancy.
JmsPlugins registered fleet-less are created with lifecycle state INACTIVE.
For the operation to be authorized, the agent must exist, and the authorized user requires JMS_PLUGIN_CREATE permission for the agent's compartment.


## Example Usage

```hcl
resource "oci_jms_jms_plugin" "test_jms_plugin" {
	#Required
	agent_id = var.jms_plugin_agent_id
	compartment_id = var.compartment_id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	fleet_id = oci_jms_fleet.test_fleet.id
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `agent_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Management Agent (OMA) or the Oracle Cloud Agent (OCA) instance where the JMS plugin is deployed.
* `compartment_id` - (Required) The OMA/OCA agent's compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`. (See [Understanding Free-form Tags](https://docs.cloud.oracle.com/iaas/Content/Tagging/Tasks/managingtagsandtagnamespaces.htm)). 
* `fleet_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the fleet. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`. (See [Managing Tags and Tag Namespaces](https://docs.cloud.oracle.com/iaas/Content/Tagging/Concepts/understandingfreeformtags.htm).) 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Jms Plugin
	* `update` - (Defaults to 20 minutes), when updating the Jms Plugin
	* `delete` - (Defaults to 20 minutes), when destroying the Jms Plugin


## Import

JmsPlugins can be imported using the `id`, e.g.

```
$ terraform import oci_jms_jms_plugin.test_jms_plugin "id"
```

