---
subcategory: "Kms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_kms_create_replica"
sidebar_current: "docs-oci-resource-kms-create_replica"
description: |-
  Provides the Create Replica resource in Oracle Cloud Infrastructure Kms service
---

# oci_kms_create_replica
This resource provides the Create Replica resource in Oracle Cloud Infrastructure Kms service.

Creates a replica for the vault in another region in the same realm

The API is a no-op if called for same region that a vault is already replicated to.
409 if called on a vault that is already replicated to a different region. Users need to delete
existing replica first before calling it with a different region.

As a provisioning operation, this call is subject to a Key Management limit that applies to
the total number of requests across all provisioning write operations. Key Management might
throttle this call to reject an otherwise valid request when the total rate of provisioning
write operations exceeds 10 requests per second for a given tenancy.


## Example Usage

```hcl
resource "oci_kms_create_replica" "test_create_replica" {
	#Required
	replica_region = var.create_replica_replica_region
	vault_id = oci_kms_vault.test_vault.id
}
```

## Argument Reference

The following arguments are supported:

* `replica_region` - (Required) The region in the realm to which the vault need to be replicated to 
* `vault_id` - (Required) The OCID of the vault.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Import

CreateReplica can be imported using the `id`, e.g.

```
$ terraform import oci_kms_create_replica.test_create_replica "id"
```

