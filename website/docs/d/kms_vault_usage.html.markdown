---
subcategory: "Kms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_kms_vault_usage"
sidebar_current: "docs-oci-datasource-kms-vault_usage"
description: |-
  Provides details about a specific Vault Usage in Oracle Cloud Infrastructure Kms service
---

# Data Source: oci_kms_vault_usage
This data source provides details about a specific Vault Usage resource in Oracle Cloud Infrastructure Kms service.

Gets the count of keys and key versions in the specified vault to calculate usage against service limits.


## Example Usage

```hcl
data "oci_kms_vault_usage" "test_vault_usage" {
	#Required
	vault_id = oci_kms_vault.test_vault.id
}
```

## Argument Reference

The following arguments are supported:

* `vault_id` - (Required) The OCID of the vault.


## Attributes Reference

The following attributes are exported:

* `key_count` - The number of keys in this vault, across all compartments, excluding keys in a `DELETED` state.
* `key_version_count` - The number of key versions in this vault, across all compartments, excluding key versions in a `DELETED` state.

