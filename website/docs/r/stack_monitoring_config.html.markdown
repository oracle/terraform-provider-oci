---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_config"
sidebar_current: "docs-oci-resource-stack_monitoring-config"
description: |-
  Provides the Config resource in Oracle Cloud Infrastructure Stack Monitoring service
---

# oci_stack_monitoring_config
This resource provides the Config resource in Oracle Cloud Infrastructure Stack Monitoring service.

Creates a configuration item, for example to define 
whether resources of a specific type should be discovered automatically. 

For example, when a new Management Agent gets registered in a certain compartment, 
this Management Agent can potentially get promoted to a HOST resource. 
The configuration item will determine if HOST resources in the selected compartment will be
discovered automatically.


## Example Usage

```hcl
resource "oci_stack_monitoring_config" "test_config" {
	#Required
	compartment_id = var.compartment_id
	config_type = var.config_config_type

	#Optional
	is_enabled = var.config_is_enabled
	resource_type = var.config_resource_type
	license = var.config_license
	defined_tags = {"foo-namespace.bar-key"= "value"}
	display_name = var.config_display_name
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment in which the configuration is created.
* `config_type` - (Required) The type of configuration. The only valid value is `"AUTO_PROMOTE"`.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Optional) (Updatable) The display name of the configuration.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `is_enabled` - (Required when config_type=AUTO_PROMOTE | LICENSE_ENTERPRISE_EXTENSIBILITY) (Updatable) True if enterprise extensibility is enabled, false if it is not enabled.
* `license` - (Required when config_type=LICENSE_AUTO_ASSIGN) (Updatable) License edition.
* `resource_type` - (Required when config_type=AUTO_PROMOTE) The type of resource to configure for automatic promotion.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the configuration.
* `config_type` - The type of configuration.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The Unique Oracle ID (OCID) that is immutable on creation.
* `is_enabled` - True if automatic promotion or enterprise extensibility is enabled, false if it is not enabled.
* `license` - License edition.
* `resource_type` - The type of resource to configure for automatic promotion.
* `state` - The current state of the configuration.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the configuration was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the Config was updated.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Config
	* `update` - (Defaults to 20 minutes), when updating the Config
	* `delete` - (Defaults to 20 minutes), when destroying the Config


## Import

Configs can be imported using the `id`, e.g.

```
$ terraform import oci_stack_monitoring_config.test_config "id"
```

