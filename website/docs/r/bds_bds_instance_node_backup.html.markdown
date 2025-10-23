---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_node_backup"
sidebar_current: "docs-oci-resource-bds-bds_instance_node_backup"
description: |-
  Provides the Bds Instance Node Backup resource in Oracle Cloud Infrastructure Big Data Service service
---

# oci_bds_bds_instance_node_backup
This resource provides the Bds Instance Node Backup resource in Oracle Cloud Infrastructure Big Data Service service.

Add a node volume backup to the cluster for an indicated node type or node.

Api doc link for the resource: https://docs.oracle.com/en-us/iaas/api/#/en/bigdata/20190531/NodeBackup/

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/blob/master/examples/big_data_service/NodeBackup/main.tf


## Example Usage

```hcl
resource "oci_bds_bds_instance_node_backup" "test_bds_instance_node_backup" {
	#Required
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id
	level_type_details {
		#Required
		level_type = var.bds_instance_node_backup_level_type_details_level_type

		#Optional
		node_host_name = var.bds_instance_node_backup_level_type_details_node_host_name
		node_type = var.bds_instance_node_backup_level_type_details_node_type
	}

	#Optional
	backup_type = var.bds_instance_node_backup_backup_type
}
```

## Argument Reference

The following arguments are supported:

* `backup_type` - (Optional) Incremental backup type includes only the changes since the last backup. Full backup type includes all changes since the volume was created.
* `bds_instance_id` - (Required) The OCID of the cluster.
* `level_type_details` - (Required) Details of the type of level used to trigger the creation of a new node backup.
	* `level_type` - (Required) Type of level used to trigger the creation of a new node backup.
	* `node_host_name` - (Required when level_type=NODE_LEVEL) (Updatable) Host name of the node to create backup.
	* `node_type` - (Required when level_type=NODE_TYPE_LEVEL) (Updatable) Type of the node or nodes of the node backup which are going to be created.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `backup_type` - Incremental backup type includes only the changes since the last backup. Full backup type includes all changes since the volume was created.
* `bds_instance_id` - The OCID of the bdsInstance which is the parent resource id.
* `id` - The unique identifier for the NodeBackup.
* `level_type_details` - Details of the type of level used to trigger the creation of a new node backup.
	* `level_type` - Type of level used to trigger the creation of a new node backup.
	* `node_host_name` - Host name of the node to create backup.
	* `node_type` - Type of the node or nodes of the node backup which are going to be created.
* `state` - The state of the NodeBackup.
* `time_created` - The time the NodeBackup was created, shown as an RFC 3339 formatted datetime string.
* `timezone` - The time zone of the execution schedule, in IANA time zone database name format

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Bds Instance Node Backup
	* `delete` - (Defaults to 20 minutes), when destroying the Bds Instance Node Backup


## Import

BdsInstanceNodeBackupConfigurations can be imported using the `id`, e.g.

```
$ terraform import oci_bds_bds_instance_node_backup.test_bds_instance_node_backup "bdsInstances/{bdsInstanceId}/nodeBackup/{nodeBackupId}" 
```

