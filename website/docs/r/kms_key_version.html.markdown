---
layout: "oci"
page_title: "OCI: oci_kms_key_version"
sidebar_current: "docs-oci-resource-kms-key_version"
description: |-
  Creates and manages an OCI KeyVersion
---

# oci_kms_key_version
The `oci_kms_key_version` resource creates and manages an OCI KeyVersion

Generates new cryptographic material for a key. Key must be in an `ENABLED` state to be
rotated.


## Example Usage

```hcl
resource "oci_kms_key_version" "test_key_version" {
	#Required
	key_id = "${oci_kms_key.test_key.id}"
	management_endpoint = "${var.key_version_management_endpoint}"
}
```

## Argument Reference

The following arguments are supported:

* `key_id` - (Required) The OCID of the key.
* `management_endpoint` - (Required) The service endpoint to perform management operations against. Management operations include 'Create,' 'Update,' 'List,' 'Get,' and 'Delete' operations. See Vault Management endpoint.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains this key version.
* `id` - The OCID of the key version.
* `key_id` - The OCID of the key associated with this key version.
* `time_created` - The date and time this key version was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2018-04-03T21:10:29.600Z` 
* `vault_id` - The OCID of the vault that contains this key version.

## Import

KeyVersions can be imported using the `id`, e.g.

```
$ terraform import oci_kms_key_version.test_key_version "managementEndpoint/{managementEndpoint}/keys/{keyId}/keyVersions/{keyVersionId}" 
```
