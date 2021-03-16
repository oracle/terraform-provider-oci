---
subcategory: "Kms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_kms_vault_replicas"
sidebar_current: "docs-oci-datasource-kms-vault_replicas"
description: |-
  Provides the list of Vault Replicas in Oracle Cloud Infrastructure Kms service
---

# Data Source: oci_kms_vault_replicas
This data source provides the list of Vault Replicas in Oracle Cloud Infrastructure Kms service.

Lists the replicas for a vault

As a provisioning operation, this call is subject to a Key Management limit that applies to
the total number of requests across all provisioning write operations. Key Management might
throttle this call to reject an otherwise valid request when the total rate of provisioning
write operations exceeds 10 requests per second for a given tenancy.


## Example Usage

```hcl
data "oci_kms_vault_replicas" "test_vault_replicas" {
	#Required
	vault_id = oci_kms_vault.test_vault.id
}
```

## Argument Reference

The following arguments are supported:

* `vault_id` - (Required) The OCID of the vault.


## Attributes Reference

The following attributes are exported:

* `vault_replicas` - The list of vault_replicas.

### VaultReplica Reference

The following attributes are exported:

* `crypto_endpoint` - The vault replica's crypto endpoint 
* `management_endpoint` - The vault replica's management endpoint 
* `region` - Region to which vault is replicated to 
* `status` - The vault replica's status

