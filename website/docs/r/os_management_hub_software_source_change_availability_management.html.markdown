---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_software_source_change_availability_management"
sidebar_current: "docs-oci-resource-os_management_hub-software_source_change_availability_management"
description: |-
  Provides the Software Source Change Availability Management resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_software_source_change_availability_management
This resource provides the Software Source Change Availability Management resource in Oracle Cloud Infrastructure Os Management Hub service.

Updates the availability for a list of specified software sources.


## Example Usage

```hcl
resource "oci_os_management_hub_software_source_change_availability_management" "test_software_source_change_availability_management" {
	#Required
	software_source_availabilities {
		#Required
		software_source_id = oci_os_management_hub_software_source.test_software_source.id

		#Optional
		availability = var.software_source_change_availability_management_software_source_availabilities_availability
		availability_at_oci = var.software_source_change_availability_management_software_source_availabilities_availability_at_oci
	}
}
```

## Argument Reference

The following arguments are supported:

* `software_source_availabilities` - (Required) List of vendor software sources and their availability statuses.
	* `availability` - (Optional) Availability of the software source to instances in private data centers or third-party clouds.
	* `availability_at_oci` - (Optional) Availability of the software source to Oracle Cloud Infrastructure instances.
	* `software_source_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the vendor software source.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Software Source Change Availability Management
	* `update` - (Defaults to 20 minutes), when updating the Software Source Change Availability Management
	* `delete` - (Defaults to 20 minutes), when destroying the Software Source Change Availability Management


## Import

SoftwareSourceChangeAvailabilityManagement can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_software_source_change_availability_management.test_software_source_change_availability_management "id"
```

