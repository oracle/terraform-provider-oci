---
layout: "oci"
page_title: "OCI: oci_kms_vaults"
sidebar_current: "docs-oci-datasource-kms-vaults"
description: |-
  Provides a list of Vaults
---

# Data Source: oci_kms_vaults
The `oci_kms_vaults` data source allows access to the list of OCI vaults

Lists vaults in the specified compartment.


## Example Usage

```hcl
data "oci_kms_vaults" "test_vaults" {
	#Required
	compartment_id = "${var.compartment_id}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.


## Attributes Reference

The following attributes are exported:

* `vaults` - The list of vaults.

### Vault Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains this vault.
* `crypto_endpoint` - The service endpoint to perform cryptographic operations against. Cryptographic operations include 'Encrypt,' 'Decrypt,' and 'GenerateDataEncryptionKey' operations. 
* `display_name` - A user-friendly name for the vault. It does not have to be unique, and it is changeable. Avoid entering confidential information. 
* `id` - The OCID of the vault.
* `management_endpoint` - The service endpoint to perform management operations against. Management operations include 'Create,' 'Update,' 'List,' 'Get,' and 'Delete' operations. 
* `state` - The vault's current state.  Example: `DELETED` 
* `time_created` - The date and time this vault was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2018-04-03T21:10:29.600Z` 
* `time_of_deletion` - An optional property for the deletion time of the Vault expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z` 
* `vault_type` - The type of vault. Each type of vault stores the key with different degrees of isolation and has different options and pricing.

