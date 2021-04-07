---
subcategory: "Kms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_kms_delete_replica"
sidebar_current: "docs-oci-resource-kms-delete_replica"
description: |-
  Provides the Delete Replica resource in Oracle Cloud Infrastructure Kms service
---

# oci_kms_delete_replica
This resource provides the Delete Replica resource in Oracle Cloud Infrastructure Kms service.

Deletes a vault replica

As a provisioning operation, this call is subject to a Key Management limit that applies to
the total number of requests across all provisioning write operations. Key Management might
throttle this call to reject an otherwise valid request when the total rate of provisioning
write operations exceeds 10 requests per second for a given tenancy.


## Example Usage

```hcl
resource "oci_kms_delete_replica" "test_delete_replica" {
	#Required
	replica_region = var.delete_replica_replica_region
	vault_id = oci_kms_vault.test_vault.id
}
```

## Argument Reference

The following arguments are supported:

* `replica_region` - (Required) The region in the realm on which the replica should be deleted 
* `vault_id` - (Required) The OCID of the vault.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Import

DeleteReplica can be imported using the `id`, e.g.

```
$ terraform import oci_kms_delete_replica.test_delete_replica "id"
```

