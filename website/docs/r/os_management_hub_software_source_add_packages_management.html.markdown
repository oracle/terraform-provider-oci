---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_software_source_add_packages_management"
sidebar_current: "docs-oci-resource-os_management_hub-software_source_add_packages_management"
description: |-
  Provides the Software Source Add Packages Management resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_software_source_add_packages_management
This resource provides the Software Source Add Packages Management resource in Oracle Cloud Infrastructure Os Management Hub service.

Adds packages to a software source. This operation can only be done for custom and versioned custom software sources that are not created using filters. 
For a versioned custom software source, you can only add packages when the source is created. Once content is added to a versioned custom software source, it is immutable.


## Example Usage

```hcl
resource "oci_os_management_hub_software_source_add_packages_management" "test_software_source_add_packages_management" {
	#Required
	packages = var.software_source_add_packages_management_packages
	software_source_id = oci_os_management_hub_software_source.test_software_source.id
}
```

## Argument Reference

The following arguments are supported:

* `packages` - (Required) List of packages specified by the full package name (NEVRA.rpm).
* `software_source_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Software Source Add Packages Management
	* `update` - (Defaults to 20 minutes), when updating the Software Source Add Packages Management
	* `delete` - (Defaults to 20 minutes), when destroying the Software Source Add Packages Management


## Import

SoftwareSourceAddPackagesManagement can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_software_source_add_packages_management.test_software_source_add_packages_management "id"
```

