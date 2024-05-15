---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_management_station_mirror_synchronize_management"
sidebar_current: "docs-oci-resource-os_management_hub-management_station_mirror_synchronize_management"
description: |-
  Provides the Management Station Mirror Synchronize Management resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_management_station_mirror_synchronize_management
This resource provides the Management Station Mirror Synchronize Management resource in Oracle Cloud Infrastructure Os Management Hub service.

Synchronize the specified software source mirrors on the management station.


## Example Usage

```hcl
resource "oci_os_management_hub_management_station_mirror_synchronize_management" "test_management_station_mirror_synchronize_management" {
	#Required
	management_station_id = oci_os_management_hub_management_station.test_management_station.id
	mirror_id = oci_os_management_hub_mirror.test_mirror.id
}
```

## Argument Reference

The following arguments are supported:

* `management_station_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management station.
* `mirror_id` - (Required) Unique Software Source identifier


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Management Station Mirror Synchronize Management
	* `update` - (Defaults to 20 minutes), when updating the Management Station Mirror Synchronize Management
	* `delete` - (Defaults to 20 minutes), when destroying the Management Station Mirror Synchronize Management


## Import

ManagementStationMirrorSynchronizeManagement can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_management_station_mirror_synchronize_management.test_management_station_mirror_synchronize_management "id"
```

