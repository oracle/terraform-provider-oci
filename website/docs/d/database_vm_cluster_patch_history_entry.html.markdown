---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_vm_cluster_patch_history_entry"
sidebar_current: "docs-oci-datasource-database-vm_cluster_patch_history_entry"
description: |-
  Provides details about a specific Vm Cluster Patch History Entry in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_vm_cluster_patch_history_entry
This data source provides details about a specific Vm Cluster Patch History Entry resource in Oracle Cloud Infrastructure Database service.

Gets the patch history details for the specified patchHistoryEntryId.


## Example Usage

```hcl
data "oci_database_vm_cluster_patch_history_entry" "test_vm_cluster_patch_history_entry" {
	#Required
	patch_history_entry_id = oci_database_patch_history_entry.test_patch_history_entry.id
	vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id
}
```

## Argument Reference

The following arguments are supported:

* `patch_history_entry_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the patch history entry.
* `vm_cluster_id` - (Required) The VM cluster [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `action` - The action being performed or was completed.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the patch history entry.
* `lifecycle_details` - A descriptive text associated with the lifecycleState. Typically contains additional displayable text. 
* `patch_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the patch.
* `state` - The current state of the action.
* `time_ended` - The date and time when the patch action completed
* `time_started` - The date and time when the patch action started.

