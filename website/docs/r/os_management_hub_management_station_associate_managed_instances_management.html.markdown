---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_management_station_associate_managed_instances_management"
sidebar_current: "docs-oci-resource-os_management_hub-management_station_associate_managed_instances_management"
description: |-
  Provides the Management Station Associate Managed Instances Management resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_management_station_associate_managed_instances_management
This resource provides the Management Station Associate Managed Instances Management resource in Oracle Cloud Infrastructure Os Management Hub service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/os-management/latest/ManagementStationAssociateManagedInstancesManagement

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/os_management_hub

Associates managed instances to the specified management station


## Example Usage

```hcl
resource "oci_os_management_hub_management_station_associate_managed_instances_management" "test_management_station_associate_managed_instances_management" {
	#Required
	managed_instances = var.management_station_associate_managed_instances_management_managed_instances
	management_station_id = oci_os_management_hub_management_station.test_management_station.id

	#Optional
	work_request_details {

		#Optional
		description = var.management_station_associate_managed_instances_management_work_request_details_description
		display_name = var.management_station_associate_managed_instances_management_work_request_details_display_name
	}
}
```

## Argument Reference

The following arguments are supported:

* `managed_instances` - (Required) List of managed instance [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) to associate to the management station.
* `management_station_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management station.
* `work_request_details` - (Optional) Provides the name and description of the job.
	* `description` - (Optional) User-specified information about the job. Avoid entering confidential information.
	* `display_name` - (Optional) A user-friendly name for the job. The name does not have to be unique. Avoid entering confidential information.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Management Station Associate Managed Instances Management
	* `update` - (Defaults to 20 minutes), when updating the Management Station Associate Managed Instances Management
	* `delete` - (Defaults to 20 minutes), when destroying the Management Station Associate Managed Instances Management


## Import

ManagementStationAssociateManagedInstancesManagement can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_management_station_associate_managed_instances_management.test_management_station_associate_managed_instances_management "id"
```

