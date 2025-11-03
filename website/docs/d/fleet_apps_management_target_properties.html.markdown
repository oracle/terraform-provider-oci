---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_target_properties"
sidebar_current: "docs-oci-datasource-fleet_apps_management-target_properties"
description: |-
  Provides the list of Target Properties in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_target_properties
This data source provides the list of Target Properties in Oracle Cloud Infrastructure Fleet Apps Management service.

Gets a list of target properties for the specified target.


## Example Usage

```hcl
data "oci_fleet_apps_management_target_properties" "test_target_properties" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	target_id = var.fleet_target_id
	target_name = var.fleet_target_name
	severity = var.target_property_severity
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment in which to list resources.
* `severity` - (Optional) Patch severity.
* `target_id` - (Optional) Target identifier.
* `target_name` - (Optional) Target name.


## Attributes Reference

The following attributes are exported:

* `target_property_collection` - The list of target_property_collection.

### TargetProperty Reference

The following attributes are exported:

* `items` - List of target properties.
	* `name` - Name of the property.
	* `value` - Value of the property.

