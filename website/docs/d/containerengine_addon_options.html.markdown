---
subcategory: "Container Engine"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_addon_options"
sidebar_current: "docs-oci-datasource-containerengine-addon_options"
description: |-
  Provides the list of Addon Options in Oracle Cloud Infrastructure Container Engine service
---

# Data Source: oci_containerengine_addon_options
This data source provides the list of Addon Options in Oracle Cloud Infrastructure Container Engine service.

Get list of supported addons for a specific kubernetes version.

## Example Usage

```hcl
data "oci_containerengine_addon_options" "test_addon_options" {
	#Required
	kubernetes_version = var.addon_option_kubernetes_version

	#Optional
	addon_name = oci_containerengine_addon.test_addon.name
}
```

## Argument Reference

The following arguments are supported:

* `addon_name` - (Optional) The name of the addon.
* `kubernetes_version` - (Required) The kubernetes version to fetch the addons.


## Attributes Reference

The following attributes are exported:

* `addon_options` - The list of addon_options.

### AddonOption Reference

The following attributes are exported:

* `addon_group` - Addon group info, a namespace concept that groups addons with similar functionalities.
* `addon_schema_version` - Addon definition schema version to validate addon.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - Description on the addon.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `is_essential` - Is it an essential addon for cluster operation or not.
* `name` - Name of the addon and it would be unique.
* `state` - The life cycle state of the addon.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the work request was created.
* `versions` - The resources this work request affects.
	* `configurations` - Addon version configuration details.
		* `description` - Information about the addon version configuration.
		* `display_name` - Display name of addon version.
		* `is_required` - If the the configuration is required or not.
		* `key` - Addon configuration key
		* `value` - Addon configuration value
	* `description` - Information about the addon version.
	* `kubernetes_version_filters` - The range of kubernetes versions an addon can be configured.
		* `exact_kubernetes_versions` - The exact version of kubernetes that are compatible. 
		* `maximum_version` - The latest kubernetes version.
		* `minimal_version` - The earliest kubernetes version.
	* `status` - Current state of the addon, only active will be visible to customer, visibility of versions in other status will be filtered  based on limits property.
	* `version_number` - Version number, need be comparable within an addon.

