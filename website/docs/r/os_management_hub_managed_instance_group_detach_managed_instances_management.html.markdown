---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_managed_instance_group_detach_managed_instances_management"
sidebar_current: "docs-oci-resource-os_management_hub-managed_instance_group_detach_managed_instances_management"
description: |-
  Provides the Managed Instance Group Detach Managed Instances Management resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_managed_instance_group_detach_managed_instances_management
This resource provides the Managed Instance Group Detach Managed Instances Management resource in Oracle Cloud Infrastructure Os Management Hub service.

Removes a managed instance from the specified managed instance group.


## Example Usage

```hcl
resource "oci_os_management_hub_managed_instance_group_detach_managed_instances_management" "test_managed_instance_group_detach_managed_instances_management" {
	#Required
	managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group.id
	managed_instances = var.managed_instance_group_detach_managed_instances_management_managed_instances
}
```

## Argument Reference

The following arguments are supported:

* `managed_instance_group_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance group.
* `managed_instances` - (Required) List of managed instance [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) to detach from the group.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Managed Instance Group Detach Managed Instances Management
	* `update` - (Defaults to 20 minutes), when updating the Managed Instance Group Detach Managed Instances Management
	* `delete` - (Defaults to 20 minutes), when destroying the Managed Instance Group Detach Managed Instances Management


## Import

ManagedInstanceGroupDetachManagedInstancesManagement can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_managed_instance_group_detach_managed_instances_management.test_managed_instance_group_detach_managed_instances_management "id"
```

