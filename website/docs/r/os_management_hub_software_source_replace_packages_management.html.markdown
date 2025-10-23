---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_software_source_replace_packages_management"
sidebar_current: "docs-oci-resource-os_management_hub-software_source_replace_packages_management"
description: |-
  Provides the Software Source Replace Packages Management resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_software_source_replace_packages_management
This resource provides the Software Source Replace Packages Management resource in Oracle Cloud Infrastructure Os Management Hub service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/os-management/latest/SoftwareSourceReplacePackagesManagement

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/os_management_hub

Replaces packages in a software source with the provided list of packages. This operation can only be done for custom software sources that are not created using filters.
Packages can be of the format:
 * name (for example: git). If isLatestContentOnly is true, only the latest version of the package will be added, otherwise all versions of the package will be added.
 * name-version-release.architecture (for example: git-2.43.5-1.el8_10.x86_64)
 * name-epoch:version-release.architecture (for example: git-0:2.43.5-1.el8_10.x86_64)


## Example Usage

```hcl
resource "oci_os_management_hub_software_source_replace_packages_management" "test_software_source_replace_packages_management" {
	#Required
	packages = var.software_source_replace_packages_management_packages
	software_source_id = oci_os_management_hub_software_source.test_software_source.id
}
```

## Argument Reference

The following arguments are supported:

* `packages` - (Required) List of packages specified by the name of the package (N) or the full package name (NVRA or NEVRA).
* `software_source_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Software Source Replace Packages Management
	* `update` - (Defaults to 20 minutes), when updating the Software Source Replace Packages Management
	* `delete` - (Defaults to 20 minutes), when destroying the Software Source Replace Packages Management


## Import

SoftwareSourceReplacePackagesManagement can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_software_source_replace_packages_management.test_software_source_replace_packages_management "id"
```

