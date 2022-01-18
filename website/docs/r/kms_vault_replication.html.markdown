---
subcategory: "Kms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_kms_vault_replication"
sidebar_current: "docs-oci-resource-kms-vault-replication"
description: |-
  Create, update and delete replica for a vault in Oracle Cloud Infrastructure Kms service
---

# oci_kms_vault_replication
This source triggers action to create, update and delete replica for a vault in Oracle Cloud Infrastructure Kms service.

A vault replica is a mirror of that vault in a different region in the same realm. 
The vault replica and all the resources have same OCID with corresponding original ones.

This only supports virtual private vault for now. 
This supports only one replica in a region for a vault. Multiple replica will be supported in the future.


## Example Usage

```hcl
resource "oci_kms_vault_replication" "test_replication" {
	#Required
  vault_id = oci_kms_vault.test_vault.id
  replica_region = var.replica_region
}
```

## Argument Reference

The following arguments are supported:

* `vault_id` - (Required) The OCID of the primary vault to create replica from.
* `replica_region` - (Required) (Updatable) The region to be created replica to. When updated,
replica will be deleted from old region, and created to updated region. 


## Attributes Reference

No attributes are exported.

