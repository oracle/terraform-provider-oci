---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_profile_attach_management_station_management"
sidebar_current: "docs-oci-resource-os_management_hub-profile_attach_management_station_management"
description: |-
  Provides the Profile Attach Management Station Management resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_profile_attach_management_station_management
This resource provides the Profile Attach Management Station Management resource in Oracle Cloud Infrastructure Os Management Hub service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/os-management/latest/ProfileAttachManagementStationManagement

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/os_management_hub

Attaches the specified management station to a profile.


## Example Usage

```hcl
resource "oci_os_management_hub_profile_attach_management_station_management" "test_profile_attach_management_station_management" {
	#Required
	management_station_id = oci_os_management_hub_management_station.test_management_station.id
	profile_id = oci_os_management_hub_profile.test_profile.id
}
```

## Argument Reference

The following arguments are supported:

* `management_station_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management station that the instance will be associated with.
* `profile_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the registration profile.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Profile Attach Management Station Management
	* `update` - (Defaults to 20 minutes), when updating the Profile Attach Management Station Management
	* `delete` - (Defaults to 20 minutes), when destroying the Profile Attach Management Station Management


## Import

ProfileAttachManagementStationManagement can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_profile_attach_management_station_management.test_profile_attach_management_station_management "id"
```

