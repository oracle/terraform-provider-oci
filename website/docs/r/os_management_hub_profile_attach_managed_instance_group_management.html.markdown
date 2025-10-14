---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_profile_attach_managed_instance_group_management"
sidebar_current: "docs-oci-resource-os_management_hub-profile_attach_managed_instance_group_management"
description: |-
  Provides the Profile Attach Managed Instance Group Management resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_profile_attach_managed_instance_group_management
This resource provides the Profile Attach Managed Instance Group Management resource in Oracle Cloud Infrastructure Os Management Hub service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/os-management/latest/ProfileAttachManagedInstanceGroupManagement

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/os_management_hub

Attaches the specified managed instance group to a profile.


## Example Usage

```hcl
resource "oci_os_management_hub_profile_attach_managed_instance_group_management" "test_profile_attach_managed_instance_group_management" {
	#Required
	managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group.id
	profile_id = oci_os_management_hub_profile.test_profile.id
}
```

## Argument Reference

The following arguments are supported:

* `managed_instance_group_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance group that the instance will be associated with.
* `profile_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the registration profile.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Profile Attach Managed Instance Group Management
	* `update` - (Defaults to 20 minutes), when updating the Profile Attach Managed Instance Group Management
	* `delete` - (Defaults to 20 minutes), when destroying the Profile Attach Managed Instance Group Management


## Import

ProfileAttachManagedInstanceGroupManagement can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_profile_attach_managed_instance_group_management.test_profile_attach_managed_instance_group_management "id"
```

