---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_replace_node_action"
sidebar_current: "docs-oci-resource-bds-bds_instance_replace_node_action"
description: |-
  Provides the Bds Instance Node Replace Action  resource in Oracle Cloud Infrastructure Big Data Service service
---

# oci_bds_bds_instance_replace_node_action
This resource replaces the node with the given hostname, in Oracle Cloud Infrastructure Big Data Service cluster.

Replace the node with the given host name in the cluster.


## Example Usage

```hcl
resource "oci_bds_bds_instance_replace_node_action" "test_bds_instance_replace_node_action" {
	#Required
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id
	node_host_name = var.bds_instance_replace_node_action.node_host_name
	node_backup_id = var.bds_instance_replace_node_action.node_backup_id
	cluster_admin_password = oci_bds_bds_instance.test_bds_instance.cluster_admin_password
	
	#Optional
	shape = var.shape
}
```

## Argument Reference

The following arguments are supported:

* `bds_instance_id` - (Required) The OCID of the cluster.
* `node_host_name`  - (Required) Host name of the node to replace. MASTER, UTILITY and EDGE node are only supported types
* `node_backup_id`  - (Required) The id of the nodeBackup to use for replacing the node.
* `cluster_admin_password` - (Required) Base-64 encoded password for the cluster admin user.
* `shape` - (Optional) Shape of the new vm when replacing the node. If not provided, BDS will attempt to replace the node with the shape of current node.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `bds_instance_id` - The OCID of the bdsInstance which is the parent resource id.
* `display_name` - A user-friendly name. Only ASCII alphanumeric characters with no spaces allowed. The name does not have to be unique, and it may be changed. Avoid entering confidential information.
* `id` - The unique identifier for the NodeBackupConfiguration.
* `level_type_details` - Details of the type of level used to trigger the creation of a new node backup configuration or node replacement configuration.
	* `level_type` - Type of level used to trigger the creation of a new node backup configuration or node replacement configuration.
	* `node_host_name` - Host name of the node to create backup configuration.
	* `node_type` - Type of the node or nodes of the node backup configuration or node replacement configuration which are going to be created.
* `number_of_backups_to_retain` - Number of backup copies to retain.
* `schedule` - Day/time recurrence (specified following RFC 5545) at which to trigger the backup process. Currently only DAILY, WEEKLY and MONTHLY frequency is supported. Days of the week are specified using BYDAY field. Time of the day is specified using BYHOUR. Other fields are not supported. 
* `state` - The state of the NodeBackupConfiguration.
* `time_created` - The time the NodeBackupConfiguration was created, shown as an RFC 3339 formatted datetime string.
* `time_updated` - The time the NodeBackupConfiguration was updated, shown as an RFC 3339 formatted datetime string. 
* `timezone` - The time zone of the execution schedule, in IANA time zone database name format

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Bds Instance Replace Node Action


