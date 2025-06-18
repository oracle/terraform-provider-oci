---
subcategory: "Wlms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_wlms_wls_domain_server_backup_content"
sidebar_current: "docs-oci-datasource-wlms-wls_domain_server_backup_content"
description: |-
  Provides details about a specific Wls Domain Server Backup Content in Oracle Cloud Infrastructure Wlms service
---

# Data Source: oci_wlms_wls_domain_server_backup_content
This data source provides details about a specific Wls Domain Server Backup Content resource in Oracle Cloud Infrastructure Wlms service.

Get details of specific backup for the WebLogic Domain.


## Example Usage

```hcl
data "oci_wlms_wls_domain_server_backup_content" "test_wls_domain_server_backup_content" {
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

* `content_type` - The type of content of the backup.
* `middleware` - The content of the middleware binaries included in a backup. 
	* `patches` - The list of patches installed in the middleware included in the backup.
		* `description` - The description of the WebLogic patch.
		* `display_name` - The display name of the WebLogic patch.
		* `id` - The ID of the WebLogic patch.

			**Note:** Not an [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
	* `version` - The version of the middleware binaries included in the backup.

