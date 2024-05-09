---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_managed_instance_attach_profile_management"
sidebar_current: "docs-oci-resource-os_management_hub-managed_instance_attach_profile_management"
description: |-
  Provides the Managed Instance Attach Profile Management resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_managed_instance_attach_profile_management
This resource provides the Managed Instance Attach Profile Management resource in Oracle Cloud Infrastructure Os Management Hub service.

Adds profile to a managed instance. After the profile has been added,
the instance can be registered as a managed instance.


## Example Usage

```hcl
resource "oci_os_management_hub_managed_instance_attach_profile_management" "test_managed_instance_attach_profile_management" {
	#Required
	managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id
	profile_id = oci_os_management_hub_profile.test_profile.id
}
```

## Argument Reference

The following arguments are supported:

* `managed_instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance.
* `profile_id` - (Required) The profile [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) to attach to the managed instance.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Managed Instance Attach Profile Management
	* `update` - (Defaults to 20 minutes), when updating the Managed Instance Attach Profile Management
	* `delete` - (Defaults to 20 minutes), when destroying the Managed Instance Attach Profile Management


## Import

ManagedInstanceAttachProfileManagement can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_managed_instance_attach_profile_management.test_managed_instance_attach_profile_management "id"
```

