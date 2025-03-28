---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_profile_detach_software_sources_management"
sidebar_current: "docs-oci-resource-os_management_hub-profile_detach_software_sources_management"
description: |-
  Provides the Profile Detach Software Sources Management resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_profile_detach_software_sources_management
This resource provides the Profile Detach Software Sources Management resource in Oracle Cloud Infrastructure Os Management Hub service.

Detaches the specified software sources from a profile.


## Example Usage

```hcl
resource "oci_os_management_hub_profile_detach_software_sources_management" "test_profile_detach_software_sources_management" {
	#Required
	profile_id = oci_os_management_hub_profile.test_profile.id
	software_sources = var.profile_detach_software_sources_management_software_sources
}
```

## Argument Reference

The following arguments are supported:

* `profile_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the registration profile.
* `software_sources` - (Required) List of software source [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) to detach from the profile.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Profile Detach Software Sources Management
	* `update` - (Defaults to 20 minutes), when updating the Profile Detach Software Sources Management
	* `delete` - (Defaults to 20 minutes), when destroying the Profile Detach Software Sources Management


## Import

ProfileDetachSoftwareSourcesManagement can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_profile_detach_software_sources_management.test_profile_detach_software_sources_management "id"
```

