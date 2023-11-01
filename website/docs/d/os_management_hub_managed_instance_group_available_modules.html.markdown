---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_managed_instance_group_available_modules"
sidebar_current: "docs-oci-datasource-os_management_hub-managed_instance_group_available_modules"
description: |-
  Provides the list of Managed Instance Group Available Modules in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_managed_instance_group_available_modules
This data source provides the list of Managed Instance Group Available Modules in Oracle Cloud Infrastructure Os Management Hub service.

Lists available modules that for the specified managed instance group. Filter the list against a variety of 
criteria including but not limited to its name.


## Example Usage

```hcl
data "oci_os_management_hub_managed_instance_group_available_modules" "test_managed_instance_group_available_modules" {
	#Required
	managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group.id

	#Optional
	compartment_id = var.compartment_id
	name = var.managed_instance_group_available_module_name
	name_contains = var.managed_instance_group_available_module_name_contains
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment that contains the resources to list.
* `managed_instance_group_id` - (Required) The managed instance group OCID.
* `name` - (Optional) The resource name. 
* `name_contains` - (Optional) A filter to return resources that may partially match the name given.


## Attributes Reference

The following attributes are exported:

* `managed_instance_group_available_module_collection` - The list of managed_instance_group_available_module_collection.

### ManagedInstanceGroupAvailableModule Reference

The following attributes are exported:

* `items` - List of module stream profile.
	* `name` - The name of the module that is available to be enabled on the managed instance group.
	* `software_source_id` - The OCID of the software source that provides this module.

