---
subcategory: "Osmanagement"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_osmanagement_managed_instances"
sidebar_current: "docs-oci-datasource-osmanagement-managed_instances"
description: |-
  Provides the list of Managed Instances in Oracle Cloud Infrastructure Osmanagement service
---

# Data Source: oci_osmanagement_managed_instances
This data source provides the list of Managed Instances in Oracle Cloud Infrastructure Osmanagement service.

Returns a list of all Managed Instances.


## Example Usage

```hcl
data "oci_osmanagement_managed_instances" "test_managed_instances" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.managed_instance_display_name
	os_family = var.managed_instance_os_family
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable.  Example: `My new resource` 
* `os_family` - (Optional) The OS family for which to list resources.


## Attributes Reference

The following attributes are exported:

* `managed_instances` - The list of managed_instances.

### ManagedInstance Reference

The following attributes are exported:

* `child_software_sources` - list of child Software Sources attached to the Managed Instance
	* `id` - software source identifier
	* `name` - software source name
* `compartment_id` - OCID for the Compartment
* `description` - Information specified by the user about the managed instance
* `display_name` - Managed Instance identifier
* `id` - OCID for the managed instance
* `is_reboot_required` - Indicates whether a reboot is required to complete installation of updates.
* `last_boot` - Time at which the instance last booted
* `last_checkin` - Time at which the instance last checked in
* `managed_instance_groups` - The ids of the managed instance groups of which this instance is a member. 
	* `display_name` - User friendly name
	* `id` - unique identifier that is immutable on creation
* `os_family` - The Operating System type of the managed instance.
* `os_kernel_version` - Operating System Kernel Version
* `os_name` - Operating System Name
* `os_version` - Operating System Version
* `parent_software_source` - the parent (base) Software Source attached to the Managed Instance
	* `id` - software source identifier
	* `name` - software source name
* `status` - status of the managed instance.
* `updates_available` - Number of updates available to be installed

