---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_exadb_vm_cluster_updates"
sidebar_current: "docs-oci-datasource-database-exadb_vm_cluster_updates"
description: |-
  Provides the list of Exadb Vm Cluster Updates in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_exadb_vm_cluster_updates
This data source provides the list of Exadb Vm Cluster Updates in Oracle Cloud Infrastructure Database service.

Lists the maintenance updates that can be applied to the specified Exadata VM cluster on Exascale Infrastructure.


## Example Usage

```hcl
data "oci_database_exadb_vm_cluster_updates" "test_exadb_vm_cluster_updates" {
	#Required
	exadb_vm_cluster_id = oci_database_exadb_vm_cluster.test_exadb_vm_cluster.id

	#Optional
	update_type = var.exadb_vm_cluster_update_update_type
	version = var.exadb_vm_cluster_update_version
}
```

## Argument Reference

The following arguments are supported:

* `exadb_vm_cluster_id` - (Required) The Exadata VM cluster [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) on Exascale Infrastructure.
* `update_type` - (Optional) A filter to return only resources that match the given update type exactly.
* `version` - (Optional) A filter to return only resources that match the given update version exactly.


## Attributes Reference

The following attributes are exported:

* `exadb_vm_cluster_updates` - The list of exadb_vm_cluster_updates.

### ExadbVmClusterUpdate Reference

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

