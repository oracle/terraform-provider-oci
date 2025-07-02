---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_node_backup"
sidebar_current: "docs-oci-datasource-bds-bds_instance_node_backup"
description: |-
  Provides details about a specific Bds Instance Node Backup in Oracle Cloud Infrastructure Big Data Service service
---

# Data Source: oci_bds_bds_instance_node_backup
This data source provides details about a specific Bds Instance Node Backup resource in Oracle Cloud Infrastructure Big Data Service service.

Returns details of NodeBackup identified by the given ID.


## Example Usage

```hcl
data "oci_bds_bds_instance_node_backup" "test_bds_instance_node_backup" {
	#Required
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id
	node_backup_id = oci_database_backup.test_backup.id
}
```

## Argument Reference

The following arguments are supported:

* `bds_instance_id` - (Required) The OCID of the cluster.
* `node_backup_id` - (Required) Unique assigned identifier of the nodeBackupId.


## Attributes Reference

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

