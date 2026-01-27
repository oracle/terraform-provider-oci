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
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/osmh/latest/Manifest

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/os_management_hub

Updates the package list document for the software source.


## Example Usage

```hcl
resource "oci_os_management_hub_software_source_manifest" "test_software_source_manifest" {
	#Required
	update_software_source_manifest_details = var.software_source_manifest_update_software_source_manifest_details
	software_source_id = oci_os_management_hub_software_source.test_software_source.id
}
```

## Argument Reference

The following arguments are supported:

* `update_software_source_manifest_details` - (Required) (Updatable) Provides the document used to update the package list of the software source.
* `software_source_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `arch_type` - The architecture type supported by the software source.
* `availability` - Availability of the software source (for non-OCI environments).
* `availability_at_oci` - Availability of the software source (for Oracle Cloud Infrastructure environments).
* `checksum_type` - The yum repository checksum type used by this software source.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the software source.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - User-specified description for the software source.
* `display_name` - User-friendly name for the software source.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `gpg_key_fingerprint` - Fingerprint of the GPG key for this software source.
* `gpg_key_id` - ID of the GPG key for this software source.
* `gpg_key_url` - URI of the GPG key for this software source.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.
* `os_family` - The OS family of the software source.
* `package_count` - Number of packages the software source contains.
* `repo_id` - The repository ID for the software source.
* `size` - The size of the software source in bytes (B).
* `software_source_type` - Type of software source.
* `state` - The current state of the software source.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the software source was created (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format). 
* `url` - URL for the repository. For vendor software sources, this is the URL to the regional yum server. For custom software sources, this is 'custom/<repoId>'.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Software Source Manifest
	* `update` - (Defaults to 20 minutes), when updating the Software Source Manifest
	* `delete` - (Defaults to 20 minutes), when destroying the Software Source Manifest


## Import

SoftwareSourceManifests can be imported using the `id`, e.g.

```
$ terraform import oci_os_management_hub_software_source_manifest.test_software_source_manifest "softwareSources/{softwareSourceId}/manifest" 
```

