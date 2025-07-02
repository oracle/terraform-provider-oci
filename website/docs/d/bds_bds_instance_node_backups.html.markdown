---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_node_backups"
sidebar_current: "docs-oci-datasource-bds-bds_instance_node_backups"
description: |-
  Provides the list of Bds Instance Node Backups in Oracle Cloud Infrastructure Big Data Service service
---

# Data Source: oci_bds_bds_instance_node_backups
This data source provides the list of Bds Instance Node Backups in Oracle Cloud Infrastructure Big Data Service service.

Returns information about the node Backups.


## Example Usage

```hcl
data "oci_bds_bds_instance_node_backups" "test_bds_instance_node_backups" {
	#Required
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id

	#Optional
	display_name = var.bds_instance_node_backup_display_name
	node_host_name = var.bds_instance_node_backup_node_host_name
	state = var.bds_instance_node_backup_state
}
```

## Argument Reference

The following arguments are supported:

* `bds_instance_id` - (Required) The OCID of the cluster.
* `display_name` - (Optional) The display name belonged to the node backup.
* `node_host_name` - (Optional) The node host name belonged to a node that has a node backup.
* `state` - (Optional) The state of the Node's Backup.


## Attributes Reference

The following attributes are exported:

* `node_backups` - The list of node_backups.

### BdsInstanceNodeBackup Reference

The following attributes are exported:

* `backup_trigger_type` - type based on how backup action was initiated.
* `backup_type` - Incremental backup type includes only the changes since the last backup. Full backup type includes all changes since the volume was created.
* `display_name` - BDS generated name for the backup. Format is nodeHostName_timeCreated.
* `id` - The id of the node backup.
* `node_backup_config_id` - The ID of the nodeBackupConfiguration if the NodeBackup is automatically created by applying the configuration.
* `node_host_name` - Host name of the node to which this backup belongs.
* `node_instance_id` - The instance OCID of the node, which is the resource from which the node backup was acquired.
* `state` - The state of the NodeBackup.
* `time_created` - The time the cluster was created, shown as an RFC 3339 formatted datetime string.

