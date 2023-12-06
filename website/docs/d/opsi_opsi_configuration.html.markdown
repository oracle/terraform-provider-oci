---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_opsi_configuration"
sidebar_current: "docs-oci-datasource-opsi-opsi_configuration"
description: |-
  Provides details about a specific Opsi Configuration in Oracle Cloud Infrastructure Opsi service
---

# Data Source: oci_opsi_opsi_configuration
This data source provides details about a specific Opsi Configuration resource in Oracle Cloud Infrastructure Opsi service.

Gets details of an OPSI configuration resource.
Values specified in configItemField and configItemCustomStatus query params will be considered, only if configItems field is requested as part of opsiConfigField query param.
Values specified in configItemCustomStatus will determine whether only customized configuration items or only non-customized configuration items or both have to be returned.


## Example Usage

```hcl
data "oci_opsi_opsi_configuration" "test_opsi_configuration" {
	#Required
	opsi_configuration_id = oci_opsi_opsi_configuration.test_opsi_configuration.id

	#Optional
	config_item_custom_status = var.opsi_configuration_config_item_custom_status
	config_item_field = var.opsi_configuration_config_item_field
	config_items_applicable_context = var.opsi_configuration_config_items_applicable_context
	opsi_config_field = var.opsi_configuration_opsi_config_field
}
```

## Argument Reference

The following arguments are supported:

* `config_item_custom_status` - (Optional) Specifies whether only customized configuration items or only non-customized configuration items or both have to be returned. By default only customized configuration items are returned. 
* `config_item_field` - (Optional) Specifies the fields to return in a config item summary.
* `config_items_applicable_context` - (Optional) Returns the configuration items filtered by applicable contexts sent in this param. By default configuration items of all applicable contexts are returned. 
* `opsi_config_field` - (Optional) Optional fields to return as part of OpsiConfiguration object. Unless requested, these fields will not be returned by default. 
* `opsi_configuration_id` - (Required) [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of OPSI configuration resource. 


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `config_items` - Array of configuration item summary objects.
	* `applicable_contexts` - List of contexts in Operations Insights where this configuration item is applicable.
	* `config_item_type` - Type of configuration item.
	* `default_value` - Value of configuration item.
	* `metadata` - Configuration item metadata.
		* `config_item_type` - Type of configuration item.
		* `data_type` - Data type of configuration item. Examples: STRING, BOOLEAN, NUMBER 
		* `description` - Description of configuration item .
		* `display_name` - User-friendly display name for the configuration item.
		* `unit_details` - Unit details of configuration item.
			* `display_name` - User-friendly display name for the configuration item unit.
			* `unit` - Unit of configuration item.
		* `value_input_details` - Allowed value details of configuration item, to validate what value can be assigned to a configuration item.
			* `allowed_value_type` - Allowed value type of configuration item.
			* `max_value` - Maximum value limit for the configuration item.
			* `min_value` - Minimum value limit for the configuration item.
			* `possible_values` - Allowed values to pick for the configuration item.
	* `name` - Name of configuration item.
	* `value` - Value of configuration item.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Description of OPSI configuration.
* `display_name` - User-friendly display name for the OPSI configuration. The name does not have to be unique.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of OPSI configuration resource. 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `opsi_config_type` - OPSI configuration type.
* `state` - OPSI configuration resource lifecycle state.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time at which the resource was first created. An RFC3339 formatted datetime string
* `time_updated` - The time at which the resource was last updated. An RFC3339 formatted datetime string

