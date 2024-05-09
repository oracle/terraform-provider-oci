---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_management_station_refresh_management"
sidebar_current: "docs-oci-resource-os_management_hub-management_station_refresh_management"
description: |-
  Provides the Management Station Refresh Management resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_management_station_refresh_management
This resource provides the Management Station Refresh Management resource in Oracle Cloud Infrastructure Os Management Hub service.

Refreshes the list of software sources mirrored by the management station to support the associated instances.


## Example Usage

```hcl
resource "oci_os_management_hub_management_station_refresh_management" "test_management_station_refresh_management" {
	#Required
	management_station_id = oci_os_management_hub_management_station.test_management_station.id
}
```

## Argument Reference

The following arguments are supported:

* `management_station_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management station.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Management Station Refresh Management
	* `update` - (Defaults to 20 minutes), when updating the Management Station Refresh Management
	* `delete` - (Defaults to 20 minutes), when destroying the Management Station Refresh Management


## Import

ManagementStationRefreshManagement can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_management_station_refresh_management.test_management_station_refresh_management "id"
```

