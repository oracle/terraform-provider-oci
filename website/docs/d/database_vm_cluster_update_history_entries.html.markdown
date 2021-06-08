---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_vm_cluster_update_history_entries"
sidebar_current: "docs-oci-datasource-database-vm_cluster_update_history_entries"
description: |-
  Provides the list of Vm Cluster Update History Entries in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_vm_cluster_update_history_entries
This data source provides the list of Vm Cluster Update History Entries in Oracle Cloud Infrastructure Database service.

Gets the history of the maintenance update actions performed on the specified VM cluster. Applies to Exadata Cloud@Customer instances only.


## Example Usage

```hcl
data "oci_database_vm_cluster_update_history_entries" "test_vm_cluster_update_history_entries" {
	#Required
	vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id

	#Optional
	state = var.vm_cluster_update_history_entry_state
	update_type = var.vm_cluster_update_history_entry_update_type
}
```

## Argument Reference

The following arguments are supported:

* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.
* `update_type` - (Optional) A filter to return only resources that match the given update type exactly.
* `vm_cluster_id` - (Required) The VM cluster [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `vm_cluster_update_history_entries` - The list of vm_cluster_update_history_entries.

### VmClusterUpdateHistoryEntry Reference

The following attributes are exported:

* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance update history entry.
* `lifecycle_details` - Descriptive text providing additional details about the lifecycle state. 
* `state` - The current lifecycle state of the maintenance update operation.
* `time_completed` - The date and time when the maintenance update action completed.
* `time_started` - The date and time when the maintenance update action started.
* `update_action` - The update action performed using this maintenance update.
* `update_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance update.
* `update_type` - The type of VM cluster maintenance update.

