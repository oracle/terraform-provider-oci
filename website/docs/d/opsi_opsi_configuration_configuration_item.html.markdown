---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_opsi_configuration_configuration_item"
sidebar_current: "docs-oci-datasource-opsi-opsi_configuration_configuration_item"
description: |-
  Provides details about a specific Opsi Configuration Configuration Item in Oracle Cloud Infrastructure Opsi service
---

# Data Source: oci_opsi_opsi_configuration_configuration_item
This data source provides details about a specific Opsi Configuration Configuration Item resource in Oracle Cloud Infrastructure Opsi service.

Gets the applicable configuration items based on the query parameters specified. Configuration items for an opsiConfigType with respect to a compartmentId can be fetched.
Values specified in configItemField param will determine what fields for each configuration items have to be returned.


## Example Usage

```hcl
data "oci_opsi_opsi_configuration_configuration_item" "test_opsi_configuration_configuration_item" {

	#Optional
	compartment_id = var.compartment_id
	config_item_field = var.opsi_configuration_configuration_item_config_item_field
	config_items_applicable_context = var.opsi_configuration_configuration_item_config_items_applicable_context
	name = var.opsi_configuration_configuration_item_name
	opsi_config_type = var.opsi_configuration_configuration_item_opsi_config_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `config_item_field` - (Optional) Specifies the fields to return in a config item summary. 
* `config_items_applicable_context` - (Optional) Returns the configuration items filtered by applicable contexts sent in this param. By default configuration items of all applicable contexts are returned. 
* `name` - (Optional) A filter to return only configuration items that match the entire name.
* `opsi_config_type` - (Optional) Filter to return configuration items based on configuration type of OPSI configuration.


## Attributes Reference

The following attributes are exported:

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
	* `value_source_config` - Source configuration from where the value is taken for a configuration item.
* `opsi_config_type` - OPSI configuration type.

