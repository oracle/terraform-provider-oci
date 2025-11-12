---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_catalog_item_variables_definition"
sidebar_current: "docs-oci-datasource-fleet_apps_management-catalog_item_variables_definition"
description: |-
  Provides details about a specific Catalog Item Variables Definition in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_catalog_item_variables_definition
This data source provides details about a specific Catalog Item Variables Definition resource in Oracle Cloud Infrastructure Fleet Apps Management service.

Gets information about a CatalogItem Variables.

## Example Usage

```hcl
data "oci_fleet_apps_management_catalog_item_variables_definition" "test_catalog_item_variables_definition" {
	#Required
	catalog_item_id = oci_fleet_apps_management_catalog_item.test_catalog_item.id
}
```

## Argument Reference

The following arguments are supported:

* `catalog_item_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the CatalogItem.


## Attributes Reference

The following attributes are exported:

* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `schema_document` - Schema Document representing Schema.yaml (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/terraformconfigresourcemanager_topic-schema.htm)
	* `can_allow_view_state` - Indicates if the stack allows users to view state information.
	* `description` - A detailed description of the stack or schema.
	* `groupings` - variable groups object.
		* `array` - Map of group names to their respective VariableGroup objects.
			* `title` - Display title for the group of variables.
			* `variables` - Array of variable references assigned to this group.
			* `visible` - Hint controlling the group's visibility.
	* `informational_text` - Informational text or notes relevant to the stack or its use.
	* `instructions` - Setup or usage instructions for this stack.
	* `locale` - The locale/language for the schema user interface (default is EN).
	* `logo_url` - logo url.
	* `output_groups` - Array of output group objects to group outputs for display or logical purposes.
		* `outputs` - Array of output strings included in group.
		* `title` - Display name for group of outputs.
	* `outputs` - A mapping of output variable names to their definitions.
		* `description` - Extended help or summary for understanding output.
		* `display_text` - Display label abel for the URL.
		* `format` - Hint about formatting or rendering the output value.
		* `is_sensitive` - If true, marks this output as sensitive.
		* `title` - Output label shown to the user.
		* `type` - Data type of the output value (such as STRING, NUMBER, OCID, etc).
		* `value` - Value of string that user can easily copy.
		* `visible` - Expression to show/hide this output.
	* `package_version` - The version of the package associated with this schema.
	* `primary_output_button` - primary output button value.
	* `schema_version` - The version of the schema definition format in use for this document.
	* `source` - Object representing the source information for the stack, indicating origin type and a reference string.
		* `reference` - Reference string providing a pointer or identifier for the source.
		* `type` - The source type of the stack (e.g. MARKETPLACE, QUICKSTART, or WEB).
	* `stack_description` - Additional details describing the stack's purpose or use-case.
	* `title` - The display name or title for this schema document.
	* `troubleshooting` - Troubleshooting tips, guidance, or steps for stack usage.
	* `variable_groups` - An array of variable group definitions for organizing variables together.
		* `title` - Display title for the group of variables.
		* `variables` - Array of variable references assigned to this group.
		* `visible` - Hint controlling the group's visibility.
	* `variables` - Key-value map of input variables defined for use by the stack.
	* `version` - The version identifier for this schema document.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 

