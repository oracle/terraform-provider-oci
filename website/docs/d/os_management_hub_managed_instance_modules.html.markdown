---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_managed_instance_modules"
sidebar_current: "docs-oci-datasource-os_management_hub-managed_instance_modules"
description: |-
  Provides the list of Managed Instance Modules in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_managed_instance_modules
This data source provides the list of Managed Instance Modules in Oracle Cloud Infrastructure Os Management Hub service.

Retrieves a list of modules, along with streams of the modules, from a managed instance. Filters may be applied to select a subset of modules based on the filter criteria.


## Example Usage

```hcl
data "oci_os_management_hub_managed_instance_modules" "test_managed_instance_modules" {
	#Required
	managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id

	#Optional
	compartment_id = var.compartment_id
	name = var.managed_instance_module_name
	name_contains = var.managed_instance_module_name_contains
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
* `managed_instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance.
* `name` - (Optional) The resource name. 
* `name_contains` - (Optional) A filter to return resources that may partially match the name given.


## Attributes Reference

The following attributes are exported:

* `managed_instance_module_collection` - The list of managed_instance_module_collection.

### ManagedInstanceModule Reference

The following attributes are exported:

* `items` - List of module streams.
	* `active_streams` - List of streams that are active in the module.
	* `disabled_streams` - List of streams that are disabled in the module.
	* `enabled_stream` - The stream that is enabled in the module. 
	* `installed_profiles` - List of installed profiles in the enabled stream of the module.
	* `name` - The module name. 
	* `software_source_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source that provides this module and the associated streams. 

