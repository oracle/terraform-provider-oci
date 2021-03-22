---
subcategory: "OS Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_osmanagement_software_source"
sidebar_current: "docs-oci-resource-osmanagement-software_source"
description: |-
  Provides the Software Source resource in Oracle Cloud Infrastructure OS Management service
---

# oci_osmanagement_software_source
This resource provides the Software Source resource in Oracle Cloud Infrastructure OS Management service.

Creates a new custom Software Source on the management system.
This will not contain any packages after it is first created,
and they must be added later.


## Example Usage

```hcl
resource "oci_osmanagement_software_source" "test_software_source" {
	#Required
	arch_type = var.software_source_arch_type
	compartment_id = var.compartment_id
	display_name = var.software_source_display_name

	#Optional
	checksum_type = var.software_source_checksum_type
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.software_source_description
	freeform_tags = {"bar-key"= "value"}
	maintainer_email = var.software_source_maintainer_email
	maintainer_name = var.software_source_maintainer_name
	maintainer_phone = var.software_source_maintainer_phone
	parent_id = oci_osmanagement_parent.test_parent.id
}
```

## Argument Reference

The following arguments are supported:

* `arch_type` - (Required) The architecture type supported by the Software Source
* `checksum_type` - (Optional) (Updatable) The yum repository checksum type used by this software source
* `compartment_id` - (Required) (Updatable) OCID for the Compartment
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Information specified by the user about the software source
* `display_name` - (Required) (Updatable) User friendly name for the software source
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `maintainer_email` - (Optional) (Updatable) Email address of the person maintaining this software source
* `maintainer_name` - (Optional) (Updatable) Name of the person maintaining this software source
* `maintainer_phone` - (Optional) (Updatable) Phone number of the person maintaining this software source
* `parent_id` - (Optional) OCID for the parent software source, if there is one


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `arch_type` - The architecture type supported by the Software Source
* `associated_managed_instances` - list of the Managed Instances associated with this Software Sources
	* `display_name` - User friendly name
	* `id` - unique identifier that is immutable on creation
* `checksum_type` - The yum repository checksum type used by this software source
* `compartment_id` - OCID for the Compartment
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Information specified by the user about the software source
* `display_name` - User friendly name for the software source
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `gpg_key_fingerprint` - Fingerprint of the GPG key for this software source
* `gpg_key_id` - ID of the GPG key for this software source
* `gpg_key_url` - URL of the GPG key for this software source
* `id` - OCID for the Software Source
* `maintainer_email` - Email address of the person maintaining this software source
* `maintainer_name` - Name of the person maintaining this software source
* `maintainer_phone` - Phone number of the person maintaining this software source
* `packages` - Number of packages
* `parent_id` - OCID for the parent software source, if there is one
* `parent_name` - Display name the parent software source, if there is one
* `repo_type` - Type of the Software Source
* `state` - The current state of the Software Source.
* `status` - status of the software source.
* `url` - URL for the repostiory

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Software Source
	* `update` - (Defaults to 20 minutes), when updating the Software Source
	* `delete` - (Defaults to 20 minutes), when destroying the Software Source


## Import

SoftwareSources can be imported using the `id`, e.g.

```
$ terraform import oci_osmanagement_software_source.test_software_source "id"
```

