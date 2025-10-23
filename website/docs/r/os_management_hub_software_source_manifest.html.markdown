---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_software_source_manifest"
sidebar_current: "docs-oci-resource-os_management_hub-software_source_manifest"
description: |-
  Provides the Software Source Manifest resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_software_source_manifest
This resource provides the Software Source Manifest resource in Oracle Cloud Infrastructure Os Management Hub service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/os-management/latest/SoftwareSourceManifest

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/os_management_hub

Updates the package list document for the software source.


## Example Usage

```hcl
resource "oci_os_management_hub_software_source_manifest" "test_software_source_manifest" {
	#Required
	software_source_id = oci_os_management_hub_software_source.test_software_source.id
	content = var.content
}
```

## Argument Reference

The following arguments are supported:

* `content` - (Required) (Updatable) Provides the manifest content used to update the package list of the software source.
* `software_source_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Software Source Manifest
	* `update` - (Defaults to 20 minutes), when updating the Software Source Manifest
	* `delete` - (Defaults to 20 minutes), when destroying the Software Source Manifest


## Import

SoftwareSourceManifests can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_software_source_manifest.test_software_source_manifest "id" 
```

