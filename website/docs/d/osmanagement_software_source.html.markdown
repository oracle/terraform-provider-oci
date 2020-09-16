---
subcategory: "Osmanagement"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_osmanagement_software_source"
sidebar_current: "docs-oci-datasource-osmanagement-software_source"
description: |-
  Provides details about a specific Software Source in Oracle Cloud Infrastructure Osmanagement service
---

# Data Source: oci_osmanagement_software_source
This data source provides details about a specific Software Source resource in Oracle Cloud Infrastructure Osmanagement service.

Returns a specific Software Source.


## Example Usage

```hcl
data "oci_osmanagement_software_source" "test_software_source" {
	#Required
	software_source_id = oci_osmanagement_software_source.test_software_source.id
}
```

## Argument Reference

The following arguments are supported:

* `software_source_id` - (Required) The OCID of the software source.


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

