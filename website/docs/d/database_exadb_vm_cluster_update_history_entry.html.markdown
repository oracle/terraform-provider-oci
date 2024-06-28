---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_exadb_vm_cluster_update_history_entry"
sidebar_current: "docs-oci-datasource-database-exadb_vm_cluster_update_history_entry"
description: |-
  Provides details about a specific Exadb Vm Cluster Update History Entry in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_exadb_vm_cluster_update_history_entry
This data source provides details about a specific Exadb Vm Cluster Update History Entry resource in Oracle Cloud Infrastructure Database service.

Gets the maintenance update history details for the specified update history entry.


## Example Usage

```hcl
data "oci_database_exadb_vm_cluster_update_history_entry" "test_exadb_vm_cluster_update_history_entry" {
	#Required
	exadb_vm_cluster_id = oci_database_exadb_vm_cluster.test_exadb_vm_cluster.id
	update_history_entry_id = oci_database_update_history_entry.test_update_history_entry.id
}
```

## Argument Reference

The following arguments are supported:

* `exadb_vm_cluster_id` - (Required) The Exadata VM cluster [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) on Exascale Infrastructure.
* `update_history_entry_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance update history entry.


## Attributes Reference

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

