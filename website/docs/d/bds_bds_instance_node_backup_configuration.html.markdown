---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_node_backup_configuration"
sidebar_current: "docs-oci-datasource-bds-bds_instance_node_backup_configuration"
description: |-
  Provides details about a specific Bds Instance Node Backup Configuration in Oracle Cloud Infrastructure Big Data Service service
---

# Data Source: oci_bds_bds_instance_node_backup_configuration
This data source provides details about a specific Bds Instance Node Backup Configuration resource in Oracle Cloud Infrastructure Big Data Service service.

Returns details of the NodeBackupConfiguration identified by the given ID.


## Example Usage

```hcl
data "oci_bds_bds_instance_node_backup_configuration" "test_bds_instance_node_backup_configuration" {
	#Required
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id
	node_backup_configuration_id = oci_audit_configuration.test_configuration.id
}
```

## Argument Reference

The following arguments are supported:

* `bds_instance_id` - (Required) The OCID of the cluster.
* `node_backup_configuration_id` - (Required) Unique Oracle-assigned identifier of the NodeBackupConfiguration.


## Attributes Reference

The following attributes are exported:

* `backup_type` - Incremental backup type includes only the changes since the last backup. Full backup type includes all changes since the volume was created.
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

