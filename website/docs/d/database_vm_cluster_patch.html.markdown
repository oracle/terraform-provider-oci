---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_vm_cluster_patch"
sidebar_current: "docs-oci-datasource-database-vm_cluster_patch"
description: |-
  Provides details about a specific Vm Cluster Patch in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_vm_cluster_patch
This data source provides details about a specific Vm Cluster Patch resource in Oracle Cloud Infrastructure Database service.

Gets information about a specified patch package.


## Example Usage

```hcl
data "oci_database_vm_cluster_patch" "test_vm_cluster_patch" {
	#Required
	patch_id = "${oci_database_patch.test_patch.id}"
	vm_cluster_id = "${oci_database_vm_cluster.test_vm_cluster.id}"
}
```

## Argument Reference

The following arguments are supported:

* `patch_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the patch.
* `vm_cluster_id` - (Required) The VM cluster [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `available_actions` - Actions that can possibly be performed using this patch.
* `description` - The text describing this patch package.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the patch.
* `last_action` - Action that is currently being performed or was completed last.
* `lifecycle_details` - A descriptive text associated with the lifecycleState. Typically can contain additional displayable text. 
* `state` - The current state of the patch as a result of lastAction.
* `time_released` - The date and time that the patch was released.
* `version` - The version of this patch package.

