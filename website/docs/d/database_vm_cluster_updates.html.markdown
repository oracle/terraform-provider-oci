---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_vm_cluster_updates"
sidebar_current: "docs-oci-datasource-database-vm_cluster_updates"
description: |-
  Provides the list of Vm Cluster Updates in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_vm_cluster_updates
This data source provides the list of Vm Cluster Updates in Oracle Cloud Infrastructure Database service.

Lists the maintenance updates that can be applied to the specified VM cluster. Applies to Exadata Cloud@Customer instances only.


## Example Usage

```hcl
data "oci_database_vm_cluster_updates" "test_vm_cluster_updates" {
	#Required
	vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id

	#Optional
	state = var.vm_cluster_update_state
	update_type = var.vm_cluster_update_update_type
}
```

## Argument Reference

The following arguments are supported:

* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.
* `update_type` - (Optional) A filter to return only resources that match the given update type exactly.
* `vm_cluster_id` - (Required) The VM cluster [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `vm_cluster_updates` - The list of vm_cluster_updates.

### VmClusterUpdate Reference

The following attributes are exported:

* `available_actions` - The possible actions that can be performed using this maintenance update.
* `description` - Details of the maintenance update package.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance update.
* `last_action` - The update action performed most recently using this maintenance update.
* `lifecycle_details` - Descriptive text providing additional details about the lifecycle state. 
* `state` - The current state of the maintenance update. Dependent on value of `lastAction`.
* `time_released` - The date and time the maintenance update was released.
* `update_type` - The type of VM cluster maintenance update.
* `version` - The version of the maintenance update package.

