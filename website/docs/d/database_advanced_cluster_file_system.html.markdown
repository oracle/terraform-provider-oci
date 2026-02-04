---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_advanced_cluster_file_system"
sidebar_current: "docs-oci-datasource-database-advanced_cluster_file_system"
description: |-
  Provides details about a specific Advanced Cluster File System in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_advanced_cluster_file_system
This data source provides details about a specific Advanced Cluster File System resource in Oracle Cloud Infrastructure Database service.

Gets information about the specified advanced cluster file systems.


## Example Usage

```hcl
data "oci_database_advanced_cluster_file_system" "test_advanced_cluster_file_system" {
	#Required
	advanced_cluster_file_system_id = oci_database_advanced_cluster_file_system.test_advanced_cluster_file_system.id
}
```

## Argument Reference

The following arguments are supported:

* `advanced_cluster_file_system_id` - (Required) The advanced cluster file system Id [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

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

