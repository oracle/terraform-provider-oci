---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_vm_cluster_update"
sidebar_current: "docs-oci-datasource-database-vm_cluster_update"
description: |-
  Provides details about a specific Vm Cluster Update in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_vm_cluster_update
This data source provides details about a specific Vm Cluster Update resource in Oracle Cloud Infrastructure Database service.

Gets information about a specified maintenance update package for a VM cluster. Applies to Exadata Cloud@Customer instances only.


## Example Usage

```hcl
data "oci_database_vm_cluster_update" "test_vm_cluster_update" {
	#Required
	update_id = oci_database_update.test_update.id
	vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id
}
```

## Argument Reference

The following arguments are supported:

* `update_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance update.
* `vm_cluster_id` - (Required) The VM cluster [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

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

