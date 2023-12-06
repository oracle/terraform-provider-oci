---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_opsi_configurations"
sidebar_current: "docs-oci-datasource-opsi-opsi_configurations"
description: |-
  Provides the list of Opsi Configurations in Oracle Cloud Infrastructure Opsi service
---

# Data Source: oci_opsi_opsi_configurations
This data source provides the list of Opsi Configurations in Oracle Cloud Infrastructure Opsi service.

Gets a list of OPSI configuration resources based on the query parameters specified.


## Example Usage

```hcl
data "oci_opsi_opsi_configurations" "test_opsi_configurations" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.opsi_configuration_display_name
	opsi_config_type = var.opsi_configuration_opsi_config_type
	state = var.opsi_configuration_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) Filter to return based on resources that match the entire display name.
* `opsi_config_type` - (Optional) Filter to return based on configuration type of OPSI configuration.
* `state` - (Optional) Filter to return based on Lifecycle state of OPSI configuration.


## Attributes Reference

The following attributes are exported:

* `opsi_configurations_collection` - The list of opsi_configurations_collection.

### OpsiConfiguration Reference

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

