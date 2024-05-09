---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_managed_instance_detach_profile_management"
sidebar_current: "docs-oci-resource-os_management_hub-managed_instance_detach_profile_management"
description: |-
  Provides the Managed Instance Detach Profile Management resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_managed_instance_detach_profile_management
This resource provides the Managed Instance Detach Profile Management resource in Oracle Cloud Infrastructure Os Management Hub service.

Detaches profile from a managed instance. After the profile has been removed,
the instance cannot be registered as a managed instance.


## Example Usage

```hcl
resource "oci_os_management_hub_managed_instance_detach_profile_management" "test_managed_instance_detach_profile_management" {
	#Required
	managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id
}
```

## Argument Reference

The following arguments are supported:

* `managed_instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Managed Instance Detach Profile Management
	* `update` - (Defaults to 20 minutes), when updating the Managed Instance Detach Profile Management
	* `delete` - (Defaults to 20 minutes), when destroying the Managed Instance Detach Profile Management


## Import

ManagedInstanceDetachProfileManagement can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_managed_instance_detach_profile_management.test_managed_instance_detach_profile_management "id"
```

