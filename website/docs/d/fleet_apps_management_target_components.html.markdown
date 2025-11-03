---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_target_components"
sidebar_current: "docs-oci-datasource-fleet_apps_management-target_components"
description: |-
  Provides the list of Target Components in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_target_components
This data source provides the list of Target Components in Oracle Cloud Infrastructure Fleet Apps Management service.

Gets a list of target component for the specified target.


## Example Usage

```hcl
data "oci_fleet_apps_management_target_components" "test_target_components" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	target_id = var.fleet_target_id
	target_name = var.fleet_target_name
	name = var.target_component_name
	severity = var.target_component_severity
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment in which to list resources.
* `name` - (Optional) Target Component Name.
* `severity` - (Optional) Patch severity.
* `target_id` - (Optional) Target identifier.
* `target_name` - (Optional) Target name.


## Attributes Reference

The following attributes are exported:

* `target_component_collection` - The list of target_component_collection.

### TargetComponent Reference

The following attributes are exported:

* `items` - List of target components.
	* `name` - Name of the component.
	* `path` - Path of the component.
	* `properties` - List of properties.
		* `name` - Name of the property.
		* `value` - Value of the property.
	* `version` - Version of the component.

