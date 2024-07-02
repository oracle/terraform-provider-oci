---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_exadb_vm_cluster_update"
sidebar_current: "docs-oci-datasource-database-exadb_vm_cluster_update"
description: |-
  Provides details about a specific Exadb Vm Cluster Update in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_exadb_vm_cluster_update
This data source provides details about a specific Exadb Vm Cluster Update resource in Oracle Cloud Infrastructure Database service.

Gets information about a specified maintenance update package for a Exadata VM cluster on Exascale Infrastructure.


## Example Usage

```hcl
data "oci_database_exadb_vm_cluster_update" "test_exadb_vm_cluster_update" {
	#Required
	exadb_vm_cluster_id = oci_database_exadb_vm_cluster.test_exadb_vm_cluster.id
	update_id = oci_database_update.test_update.id
}
```

## Argument Reference

The following arguments are supported:

* `exadb_vm_cluster_id` - (Required) The Exadata VM cluster [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) on Exascale Infrastructure.
* `update_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance update.


## Attributes Reference

The following attributes are exported:

* `available_actions` - The possible actions performed by the update operation on the infrastructure components.
* `description` - Details of the maintenance update package.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance update.
* `last_action` - The previous update action performed.
* `lifecycle_details` - Descriptive text providing additional details about the lifecycle state. 
* `state` - The current state of the maintenance update. Dependent on value of `lastAction`.
* `time_released` - The date and time the maintenance update was released.
* `update_type` - The type of cloud VM cluster maintenance update.
* `version` - The version of the maintenance update package.

