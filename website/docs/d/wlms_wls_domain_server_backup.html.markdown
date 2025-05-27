---
subcategory: "Wlms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_wlms_wls_domain_server_backup"
sidebar_current: "docs-oci-datasource-wlms-wls_domain_server_backup"
description: |-
  Provides details about a specific Wls Domain Server Backup in Oracle Cloud Infrastructure Wlms service
---

# Data Source: oci_wlms_wls_domain_server_backup
This data source provides details about a specific Wls Domain Server Backup resource in Oracle Cloud Infrastructure Wlms service.

Get details of specific backup for the WebLogic Domain.


## Example Usage

```hcl
data "oci_wlms_wls_domain_server_backup" "test_wls_domain_server_backup" {
	#Required
	backup_id = oci_database_backup.test_backup.id
	server_id = oci_wlms_server.test_server.id
	wls_domain_id = oci_wlms_wls_domain.test_wls_domain.id
}
```

## Argument Reference

The following arguments are supported:

* `backup_id` - (Required) The unique identifier of the backup.

	**Note:** Not an [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `server_id` - (Required) The unique identifier of a server.

	**Note:** Not an [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `wls_domain_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WebLogic domain.


## Attributes Reference

The following attributes are exported:

* `backup_location` - The location of the backup. For backups of type LOCAL_FILE this is the absolute path of the backup file.
* `content_type` - The type of content of the backup.
* `id` - The unique identifier of the backup.

	**Note:** Not an [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `managed_instance_id` - The managed instance ID of the server for which the backup was created.
* `time_created` - The date and time when the backup was created (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format).  Example: `2016-08-25T21:10:29.600Z` 
* `type` - The type of the backup.

