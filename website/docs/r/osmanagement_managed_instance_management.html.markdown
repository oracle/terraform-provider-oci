---
subcategory: "OS Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_osmanagement_managed_instance_management"
sidebar_current: "docs-oci-resource-osmanagement-managed_instance_management"
description: |-
  Provides the Managed Instance Management resource in Oracle Cloud Infrastructure OS Management service
---

# oci_osmanagement_managed_instance_management
This resource provides the Managed Instance Management in Oracle Cloud Infrastructure OS Management service.
The resource can be used to attach/detach parent software source, child software sources and managed instance groups from managed instances.

Adds a parent software source to a managed instance. After the software source has been added, then packages from that software source can be installed on the managed instance. Software sources that have this software source as a parent will be able to be added to this managed instance.
Removes a software source from a managed instance. All child software sources will also be removed from the managed instance. Packages will no longer be able to be installed from these software sources.
        
Adds a child software source to a managed instance. After the software source has been added, then packages from that software source can be installed on the managed instance.   
Removes a child software source from a managed instance. Packages will no longer be able to be installed from these software sources.

Adds a Managed Instance to a Managed Instance Group. After the Managed Instance has been added, then operations can be performed on the Managed Instance Group which will then apply to all Managed Instances in the group. 
Removes a Managed Instance from a Managed Instance Group.
        
**NOTE** The resource on CREATE will detach any already attached parent software source, child software sources, managed instance groups to the managed instance. 
Destroying this resource will not delete any associations.

## Example Usage

```hcl
resource "oci_osmanagement_managed_instance_management" "test_managed_instance_management" {
	#Required
	managed_instance_id = oci_osmanagement_managed_instance.test_managed_instance.id
	
	#optional
	parent_software_source {
        id   = oci_osmanagement_software_source.test_parent_software_source.id
        name = oci_osmanagement_software_source.test_parent_software_source.display_name
    }
    
    managed_instance_groups {
       id           = oci_osmanagement_managed_instance_group.test_managed_instance_group.id
       display_name = var.managed_instance_group_display_name
    }
    
    child_software_sources {
       id   = oci_osmanagement_software_source.test_software_source_child.id
       name = oci_osmanagement_software_source.test_software_source_child.display_name
    }
}
```

## Argument Reference

The following arguments are supported:

* `managed_instance_id` - (Required) OCID for the managed instance
* `child_software_sources` - (Optional) (Updatable) list of child Software Sources attached to the Managed Instance
	* `id` - software source identifier
	* `name` - software source name
* `managed_instance_groups` - (Optional) (Updatable) The ids of the managed instance groups of which this instance is a member. 
	* `display_name` - User friendly name
	* `id` - unique identifier that is immutable on creation
* `parent_software_source` - (Optional) (Updatable) the parent (base) Software Source attached to the Managed Instance
	* `id` - software source identifier
	* `name` - software source name

	
** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values
	
## Attributes Reference

The following attributes are exported:

* `child_software_sources` - list of child Software Sources attached to the Managed Instance
	* `id` - software source identifier
	* `name` - software source name
* `compartment_id` - OCID for the Compartment
* `description` - Information specified by the user about the managed instance
* `display_name` - Managed Instance identifier
* `id` - OCID for the managed instance
* `last_boot` - Time at which the instance last booted
* `last_checkin` - Time at which the instance last checked in
* `managed_instance_groups` - The ids of the managed instance groups of which this instance is a member. 
	* `display_name` - User friendly name
	* `id` - unique identifier that is immutable on creation
* `os_kernel_version` - Operating System Kernel Version
* `os_name` - Operating System Name
* `os_version` - Operating System Version
* `parent_software_source` - the parent (base) Software Source attached to the Managed Instance
	* `id` - software source identifier
	* `name` - software source name
* `status` - status of the managed instance.
* `updates_available` - Number of updates available to be installed

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Managed Instance Management
	* `update` - (Defaults to 20 minutes), when updating the Managed Instance Management
	* `delete` - (Defaults to 20 minutes), when destroying the Managed Instance Management


## Import

Import is not supported for this resource.

