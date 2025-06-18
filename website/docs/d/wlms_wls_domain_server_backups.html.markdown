---
subcategory: "Wlms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_wlms_wls_domain_server_backups"
sidebar_current: "docs-oci-datasource-wlms-wls_domain_server_backups"
description: |-
  Provides the list of Wls Domain Server Backups in Oracle Cloud Infrastructure Wlms service
---

# Data Source: oci_wlms_wls_domain_server_backups
This data source provides the list of Wls Domain Server Backups in Oracle Cloud Infrastructure Wlms service.

Gets a list of backups for the server of a specific WebLogic Domain.


## Example Usage

```hcl
data "oci_wlms_wls_domain_server_backups" "test_wls_domain_server_backups" {
	#Required
	server_id = oci_wlms_server.test_server.id
	wls_domain_id = oci_wlms_wls_domain.test_wls_domain.id
}
```

## Argument Reference

The following arguments are supported:

* `server_id` - (Required) The unique identifier of a server.

	**Note:** Not an [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `wls_domain_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WebLogic domain.


## Attributes Reference

The following attributes are exported:

* `backup_collection` - The list of backup_collection.

### WlsDomainServerBackup Reference

The following attributes are exported:

* `backup_location` - The location of the backup. For backups of type LOCAL_FILE this is the absolute path of the backup file.
* `content_type` - The type of content of the backup.
* `id` - The unique identifier of the backup.

	**Note:** Not an [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `managed_instance_id` - The managed instance ID of the server for which the backup was created.
* `time_created` - The date and time when the backup was created (in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) format).  Example: `2016-08-25T21:10:29.600Z` 
* `type` - The type of the backup.

