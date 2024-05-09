---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_management_station_synchronize_mirrors_management"
sidebar_current: "docs-oci-resource-os_management_hub-management_station_synchronize_mirrors_management"
description: |-
  Provides the Management Station Synchronize Mirrors Management resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_management_station_synchronize_mirrors_management
This resource provides the Management Station Synchronize Mirrors Management resource in Oracle Cloud Infrastructure Os Management Hub service.

Synchronize the specified software sources mirrors on the management station.


## Example Usage

```hcl
resource "oci_os_management_hub_management_station_synchronize_mirrors_management" "test_management_station_synchronize_mirrors_management" {
	#Required
	management_station_id = oci_os_management_hub_management_station.test_management_station.id
	software_source_list = var.management_station_synchronize_mirrors_management_software_source_list
}
```

## Argument Reference

The following arguments are supported:

* `management_station_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management station.
* `software_source_list` - (Required) List of Software Source OCIDs to synchronize


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Management Station Synchronize Mirrors Management
	* `update` - (Defaults to 20 minutes), when updating the Management Station Synchronize Mirrors Management
	* `delete` - (Defaults to 20 minutes), when destroying the Management Station Synchronize Mirrors Management


## Import

ManagementStationSynchronizeMirrorsManagement can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_management_station_synchronize_mirrors_management.test_management_station_synchronize_mirrors_management "id"
```

