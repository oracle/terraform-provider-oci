---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_vm_cluster_update_history_entry"
sidebar_current: "docs-oci-datasource-database-vm_cluster_update_history_entry"
description: |-
  Provides details about a specific Vm Cluster Update History Entry in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_vm_cluster_update_history_entry
This data source provides details about a specific Vm Cluster Update History Entry resource in Oracle Cloud Infrastructure Database service.

Gets the maintenance update history details for the specified update history entry. Applies to Exadata Cloud@Customer instances only.


## Example Usage

```hcl
data "oci_database_vm_cluster_update_history_entry" "test_vm_cluster_update_history_entry" {
	#Required
	update_history_entry_id = oci_database_update_history_entry.test_update_history_entry.id
	vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id
}
```

## Argument Reference

The following arguments are supported:

* `update_history_entry_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance update history entry.
* `vm_cluster_id` - (Required) The VM cluster [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance update history entry.
* `lifecycle_details` - Descriptive text providing additional details about the lifecycle state. 
* `state` - The current lifecycle state of the maintenance update operation.
* `time_completed` - The date and time when the maintenance update action completed.
* `time_started` - The date and time when the maintenance update action started.
* `update_action` - The update action performed using this maintenance update.
* `update_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance update.
* `update_type` - The type of VM cluster maintenance update.

