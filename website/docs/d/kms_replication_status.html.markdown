---
subcategory: "Kms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_kms_replication_status"
sidebar_current: "docs-oci-datasource-kms-replication_status"
description: |-
  Provides details about a specific Replication Status in Oracle Cloud Infrastructure Kms service
---

# Data Source: oci_kms_replication_status
This data source provides details about a specific Replication Status resource in Oracle Cloud Infrastructure Kms service.

When a vault has a replica, each operation on the vault or its resources, such as
keys, is replicated and has an associated replicationId. Replication status provides
details about whether the operation associated with the given replicationId has been
successfully applied across replicas.


## Example Usage

```hcl
data "oci_kms_replication_status" "test_replication_status" {
	#Required
	replication_id = oci_kms_replication.test_replication.id
	management_endpoint = var.replication_status_management_endpoint
}
```

## Argument Reference

The following arguments are supported:

* `replication_id` - (Required) replicationId associated with an operation on a resource 
* `management_endpoint` - (Required) The service endpoint to perform management operations against. See Vault Management endpoint. 

## Attributes Reference

The following attributes are exported:

* `replica_details` - 
	* `region` - The replica region 
	* `status` - Replication status associated with a replicationId

