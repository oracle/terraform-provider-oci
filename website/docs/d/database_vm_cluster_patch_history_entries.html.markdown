---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_vm_cluster_patch_history_entries"
sidebar_current: "docs-oci-datasource-database-vm_cluster_patch_history_entries"
description: |-
  Provides the list of Vm Cluster Patch History Entries in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_vm_cluster_patch_history_entries
This data source provides the list of Vm Cluster Patch History Entries in Oracle Cloud Infrastructure Database service.

Gets the history of the patch actions performed on the specified Vm cluster.


## Example Usage

```hcl
data "oci_database_vm_cluster_patch_history_entries" "test_vm_cluster_patch_history_entries" {
	#Required
	vm_cluster_id = "${oci_database_vm_cluster.test_vm_cluster.id}"
}
```

## Argument Reference

The following arguments are supported:

* `vm_cluster_id` - (Required) The VM cluster [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `patch_history_entries` - The list of patch_history_entries.

### VmClusterPatchHistoryEntry Reference

The following attributes are exported:

* `action` - The action being performed or was completed.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the patch history entry.
* `lifecycle_details` - A descriptive text associated with the lifecycleState. Typically contains additional displayable text. 
* `patch_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the patch.
* `state` - The current state of the action.
* `time_ended` - The date and time when the patch action completed
* `time_started` - The date and time when the patch action started.

