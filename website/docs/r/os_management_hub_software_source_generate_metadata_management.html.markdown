---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_software_source_generate_metadata_management"
sidebar_current: "docs-oci-resource-os_management_hub-software_source_generate_metadata_management"
description: |-
  Provides the Software Source Generate Metadata Management resource in Oracle Cloud Infrastructure Os Management Hub service
---

# oci_os_management_hub_software_source_generate_metadata_management
This resource provides the Software Source Generate Metadata Management resource in Oracle Cloud Infrastructure Os Management Hub service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/os-management/latest/SoftwareSourceGenerateMetadataManagement

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/os_management_hub

Regenerates metadata for the specified custom software source.

## Example Usage

```hcl
resource "oci_os_management_hub_software_source_generate_metadata_management" "test_software_source_generate_metadata_management" {
	#Required
	software_source_id = oci_os_management_hub_software_source.test_software_source.id
}
```

## Argument Reference

The following arguments are supported:

* `software_source_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Software Source Generate Metadata Management
	* `update` - (Defaults to 20 minutes), when updating the Software Source Generate Metadata Management
	* `delete` - (Defaults to 20 minutes), when destroying the Software Source Generate Metadata Management


## Import

SoftwareSourceGenerateMetadataManagement can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_software_source_generate_metadata_management.test_software_source_generate_metadata_management "id"
```

