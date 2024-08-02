---
subcategory: "Psql"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_psql_backup"
sidebar_current: "docs-oci-datasource-psql-backup"
description: |-
  Provides details about a specific Backup in Oracle Cloud Infrastructure Psql service
---

# Data Source: oci_psql_backup
This data source provides details about a specific Backup resource in Oracle Cloud Infrastructure Psql service.

Gets a backup by identifier.

## Example Usage

```hcl
data "oci_psql_backup" "test_backup" {
	#Required
	backup_id = oci_psql_backup.test_backup.id
}
```

## Argument Reference

The following arguments are supported:

* `backup_id` - (Required) A unique identifier for the backup.


## Attributes Reference

The following attributes are exported:

* `backup_size` - The size of the backup, in gigabytes.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the backup.
* `db_system_details` - Information about the database system associated with a backup.
	* `db_version` - The major and minor versions of the database system software.
	* `system_type` - Type of the database system.
* `db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup's source database system.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A description for the backup.
* `display_name` - A user-friendly display name for the backup. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup.
* `last_accepted_request_token` - lastAcceptedRequestToken from MP.
* `last_completed_request_token` - lastCompletedRequestToken from MP.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `retention_period` - Backup retention period in days.
* `source_type` - Specifies whether the backup was created manually, or by a management policy.
* `state` - The current state of the backup.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the backup request was received, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the backup was updated, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 

