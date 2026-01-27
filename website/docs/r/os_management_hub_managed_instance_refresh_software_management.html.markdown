---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_managed_instance_refresh_software_management"
sidebar_current: "docs-oci-resource-os_management_hub-managed_instance_refresh_software_management"
description: |-
  Provides the Managed Instance Refresh Software Management resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_managed_instance_refresh_software_management
This resource provides the Managed Instance Refresh Software Management resource in Oracle Cloud Infrastructure Os Management Hub service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/osmh/latest/ManagedInstance/RefreshSoftware

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/os_management_hub

Refreshes the package or Windows update information on a managed instance with the latest data from the software source. This does not update packages on the instance. It provides the service with the latest package data.


## Example Usage

```hcl
resource "oci_os_management_hub_managed_instance_refresh_software_management" "test_managed_instance_refresh_software_management" {
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
	* `create` - (Defaults to 20 minutes), when creating the Managed Instance Refresh Software Management
	* `update` - (Defaults to 20 minutes), when updating the Managed Instance Refresh Software Management
	* `delete` - (Defaults to 20 minutes), when destroying the Managed Instance Refresh Software Management


## Import

ManagedInstanceRefreshSoftwareManagement can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_managed_instance_refresh_software_management.test_managed_instance_refresh_software_management "id"
```

