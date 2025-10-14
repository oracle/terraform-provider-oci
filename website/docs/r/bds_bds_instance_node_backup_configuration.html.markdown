---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_node_backup_configuration"
sidebar_current: "docs-oci-resource-bds-bds_instance_node_backup_configuration"
description: |-
  Provides the Bds Instance Node Backup Configuration resource in Oracle Cloud Infrastructure Big Data Service service
---

# oci_bds_bds_instance_node_backup_configuration
This resource provides the Bds Instance Node Backup Configuration resource in Oracle Cloud Infrastructure Big Data Service service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/bigdata/latest/BdsInstanceNodeBackupConfiguration

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/big_data_service

Add a node volume backup configuration to the cluster for an indicated node type or node.


## Example Usage

```hcl
resource "oci_bds_bds_instance_node_backup_configuration" "test_bds_instance_node_backup_configuration" {
	#Required
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id
	level_type_details {
		#Required
		level_type = var.bds_instance_node_backup_configuration_level_type_details_level_type

		#Optional
		node_host_name = var.bds_instance_node_backup_configuration_level_type_details_node_host_name
		node_type = var.bds_instance_node_backup_configuration_level_type_details_node_type
	}
	schedule = var.bds_instance_node_backup_configuration_schedule

	#Optional
	backup_type = var.bds_instance_node_backup_configuration_backup_type
	display_name = var.bds_instance_node_backup_configuration_display_name
	number_of_backups_to_retain = var.bds_instance_node_backup_configuration_number_of_backups_to_retain
	timezone = var.bds_instance_node_backup_configuration_timezone
}
```

## Argument Reference

The following arguments are supported:

* `backup_type` - (Optional) (Updatable) Incremental backup type includes only the changes since the last backup. Full backup type includes all changes since the volume was created.
* `bds_instance_id` - (Required) The OCID of the cluster.
* `display_name` - (Optional) (Updatable) A user-friendly name. Only ASCII alphanumeric characters with no spaces allowed. The name does not have to be unique, and it may be changed. Avoid entering confidential information.
* `level_type_details` - (Required) (Updatable) Details of the type of level used to trigger the creation of a new node backup configuration or node replacement configuration.
	* `level_type` - (Required) (Updatable) Type of level used to trigger the creation of a new node backup configuration or node replacement configuration.
	* `node_host_name` - (Required when level_type=NODE_LEVEL) (Updatable) Host name of the node to create backup configuration.
	* `node_type` - (Required when level_type=NODE_TYPE_LEVEL) (Updatable) Type of the node or nodes of the node backup configuration or node replacement configuration which are going to be created. Accepted values are MASTER and UTILITY.
* `number_of_backups_to_retain` - (Optional) (Updatable) Number of backup copies to retain.
* `schedule` - (Required) (Updatable) Day/time recurrence (specified following RFC 5545) at which to trigger the backup process. Currently only DAILY, WEEKLY and MONTHLY frequency is supported. Days of the week are specified using BYDAY field. Time of the day is specified using BYHOUR. Other fields are not supported. 
* `timezone` - (Optional) (Updatable) The time zone of the execution schedule, in IANA time zone database name format


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `backup_type` - Incremental backup type includes only the changes since the last backup. Full backup type includes all changes since the volume was created.
* `bds_instance_id` - The OCID of the bdsInstance which is the parent resource id.
* `display_name` - A user-friendly name. Only ASCII alphanumeric characters with no spaces allowed. The name does not have to be unique, and it may be changed. Avoid entering confidential information.
* `id` - The unique identifier for the NodeBackupConfiguration.
* `level_type_details` - Details of the type of level used to trigger the creation of a new node backup configuration or node replacement configuration.
	* `level_type` - Type of level used to trigger the creation of a new node backup configuration or node replacement configuration.
	* `node_host_name` - Host name of the node to create backup configuration.
	* `node_type` - Type of the node or nodes of the node backup configuration or node replacement configuration which are going to be created. Accepted values are MASTER and UTILITY.
* `number_of_backups_to_retain` - Number of backup copies to retain.
* `schedule` - Day/time recurrence (specified following RFC 5545) at which to trigger the backup process. Currently only DAILY, WEEKLY and MONTHLY frequency is supported. Days of the week are specified using BYDAY field. Time of the day is specified using BYHOUR. Other fields are not supported. 
* `state` - The state of the NodeBackupConfiguration.
* `time_created` - The time the NodeBackupConfiguration was created, shown as an RFC 3339 formatted datetime string.
* `time_updated` - The time the NodeBackupConfiguration was updated, shown as an RFC 3339 formatted datetime string. 
* `timezone` - The time zone of the execution schedule, in IANA time zone database name format

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Bds Instance Node Backup Configuration
	* `update` - (Defaults to 20 minutes), when updating the Bds Instance Node Backup Configuration
	* `delete` - (Defaults to 20 minutes), when destroying the Bds Instance Node Backup Configuration


## Import

BdsInstanceNodeBackupConfigurations can be imported using the `id`, e.g.

```
$ terraform import oci_bds_bds_instance_node_backup_configuration.test_bds_instance_node_backup_configuration "bdsInstances/{bdsInstanceId}/nodeBackupConfigurations/{nodeBackupConfigurationId}" 
```

