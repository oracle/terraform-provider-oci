---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_exadb_vm_cluster_update_history_entries"
sidebar_current: "docs-oci-datasource-database-exadb_vm_cluster_update_history_entries"
description: |-
  Provides the list of Exadb Vm Cluster Update History Entries in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_exadb_vm_cluster_update_history_entries
This data source provides the list of Exadb Vm Cluster Update History Entries in Oracle Cloud Infrastructure Database service.

Gets the history of the maintenance update actions performed on the specified Exadata VM cluster on Exascale Infrastructure.


## Example Usage

```hcl
data "oci_database_exadb_vm_cluster_update_history_entries" "test_exadb_vm_cluster_update_history_entries" {
	#Required
	exadb_vm_cluster_id = oci_database_exadb_vm_cluster.test_exadb_vm_cluster.id

	#Optional
	update_type = var.exadb_vm_cluster_update_history_entry_update_type
}
```

## Argument Reference

The following arguments are supported:

* `exadb_vm_cluster_id` - (Required) The Exadata VM cluster [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) on Exascale Infrastructure.
* `update_type` - (Optional) A filter to return only resources that match the given update type exactly.


## Attributes Reference

The following attributes are exported:

* `exadb_vm_cluster_update_history_entries` - The list of exadb_vm_cluster_update_history_entries.

### ExadbVmClusterUpdateHistoryEntry Reference

The following attributes are exported:

* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance update history entry.
* `lifecycle_details` - Descriptive text providing additional details about the lifecycle state. 
* `state` - The current lifecycle state of the maintenance update operation.
* `time_completed` - The date and time when the maintenance update action completed.
* `time_started` - The date and time when the maintenance update action started.
* `update_action` - The update action.
* `update_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance update.
* `update_type` - The type of cloud VM cluster maintenance update.
* `version` - The version of the maintenance update package.

