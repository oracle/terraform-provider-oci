---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_advanced_cluster_file_systems"
sidebar_current: "docs-oci-datasource-database-advanced_cluster_file_systems"
description: |-
  Provides the list of Advanced Cluster File Systems in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_advanced_cluster_file_systems
This data source provides the list of Advanced Cluster File Systems in Oracle Cloud Infrastructure Database service.

Lists the advanced cluster file system resources in the specified compartment.


## Example Usage

```hcl
data "oci_database_advanced_cluster_file_systems" "test_advanced_cluster_file_systems" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	name = var.advanced_cluster_file_system_name
	resource_id = oci_cloud_guard_resource.test_resource.id
	state = var.advanced_cluster_file_system_state
	vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `name` - (Optional) A filter to return only resources that match the entire name given. The match is not case sensitive.
* `resource_id` - (Optional) If provided, filters the results for the specified resource Id.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.
* `vm_cluster_id` - (Optional) A filter to return only ACFS that match the given vm cluster id exactly.


## Attributes Reference

The following attributes are exported:

* `advanced_cluster_file_system_collection` - The list of advanced_cluster_file_system_collection.

### AdvancedClusterFileSystem Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `description` - Description of the advanced cluster file system.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the advanced cluster file system.
* `is_mounted` - True if the file system is mounted on all VMs within VM Cluster.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `mount_point` - The mount point of file system.
* `name` - The user-friendly name for the Advanced cluster file system. The file system name is unique for a cluster.
* `state` - The current state of the advanced cluster file system. Valid states are CREATING, AVAILABLE, UPDATING, FAILED, DELETED. 
* `storage_in_gbs` - The total storage allocated for advanced cluster file system in GBs.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `time_created` - The date and time the advanced cluster file system was created.
* `time_updated` - The last date and time that the advanced cluster file system was updated.
* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.
* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster.

