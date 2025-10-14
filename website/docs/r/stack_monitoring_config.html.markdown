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
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/stack-monitoring/latest/Config

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/stack_monitoring

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

* `additional_configurations` - (Applicable when config_type=ONBOARD) (Updatable) Property Details
	* `properties_map` - (Applicable when config_type=ONBOARD) (Updatable) Key/Value pair of Property
* `compartment_id` - (Required) (Updatable) Compartment in which the configuration is created.
* `config_type` - (Required) The type of configuration. The only valid value is `"AUTO_PROMOTE"`.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Optional) (Updatable) The display name of the configuration.
* `dynamic_groups` - (Applicable when config_type=ONBOARD) (Updatable) List of dynamic groups dedicated for Stack Monitoring.
	* `domain` - (Applicable when config_type=ONBOARD) (Updatable) Identity domain name 
	* `name` - (Required when config_type=ONBOARD) (Updatable) Name of dynamic Group 
	* `stack_monitoring_assignment` - (Required when config_type=ONBOARD) (Updatable) Assignment of dynamic group in context of Stack Monitoring service. It describes the purpose of dynamic groups in Stack Monitoring. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `is_enabled` - (Required when config_type=AUTO_PROMOTE | LICENSE_ENTERPRISE_EXTENSIBILITY) (Updatable) True if enterprise extensibility is enabled, false if it is not enabled.
* `is_manually_onboarded` - (Required when config_type=ONBOARD) (Updatable) True if customer decides marks configuration as manually configured.
* `license` - (Required when config_type=LICENSE_AUTO_ASSIGN) (Updatable) License edition.
* `policy_names` - (Applicable when config_type=ONBOARD) (Updatable) List of policy names assigned for onboarding
* `resource_type` - (Required when config_type=AUTO_PROMOTE) The type of resource to configure for automatic promotion.
* `user_groups` - (Applicable when config_type=ONBOARD) (Updatable) List of user groups dedicated for Stack Monitoring.
	* `domain` - (Applicable when config_type=ONBOARD) (Updatable) Identity domain name 
	* `name` - (Required when config_type=ONBOARD) (Updatable) Name of user Group 
	* `stack_monitoring_role` - (Required when config_type=ONBOARD) (Updatable) Role assigned to user group in context of Stack Monitoring service. Access role can be for example: ADMINISTRATOR, OPERATOR, VIEWER, any other access role 
* `version` - (Applicable when config_type=ONBOARD) (Updatable) Assigned version to given onboard configuration.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `additional_configurations` - Property Details
	* `properties_map` - Key/Value pair of Property
* `compartment_id` - The OCID of the compartment containing the configuration.
* `config_type` - The type of configuration.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `dynamic_groups` - List of dynamic groups dedicated for Stack Monitoring.
	* `domain` - Identity domain name 
	* `name` - Name of dynamic Group 
	* `stack_monitoring_assignment` - Assignment of dynamic group in context of Stack Monitoring service. It describes the purpose of dynamic groups in Stack Monitoring. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The Unique Oracle ID (OCID) that is immutable on creation.
* `is_manually_onboarded` - True if customer decides marks configuration as manually configured.
* `is_enabled` - True if automatic activation of the Management Agent plugin, automatic promotion or enterprise extensibility is enabled, false if it is not enabled.
* `license` - License edition.
* `policy_names` - List of policy names assigned for onboarding
* `resource_type` - The type of resource to configure for automatic promotion.
* `state` - The current state of the configuration.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the configuration was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the Config was updated.
* `user_groups` - List of user groups dedicated for Stack Monitoring.
	* `domain` - Identity domain name 
	* `name` - Name of user Group 
	* `stack_monitoring_role` - Role assigned to user group in context of Stack Monitoring service. Access role can be for example: ADMINISTRATOR, OPERATOR, VIEWER, any other access role 
* `version` - Assigned version to given onboard configuration.

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

