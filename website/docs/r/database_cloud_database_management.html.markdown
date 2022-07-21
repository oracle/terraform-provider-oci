---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_cloud_database_management"
sidebar_current: "docs-oci-resource-database-cloud_database_management"
description: |-
  Provides the Database Management resource in Oracle Cloud Infrastructure Database service
---

# oci_database_cloud_database_management
This resource provides the Database Management resource in Oracle Cloud Infrastructure Database service.

Enable / Update / Disable database management for the specified Oracle Database instance.

Database Management requires `USER_NAME`, `PASSWORD_SECRET_ID` and `PRIVATE_END_POINT_ID`.
`database.0.database_management_config` is updated to appropriate managementType and managementStatus for the specified Oracle Database instance.

## Example Usage

```hcl
resource "oci_database_cloud_database_management" "test" {
  database_id           = oci_database_database.test_database.id
  management_type       = var.database_cloud_database_management_details_management_type
  private_end_point_id  = var.database_cloud_database_management_details_private_end_point_id
  service_name          = var.database_cloud_database_management_details_service_name
  credentialdetails {
    user_name           = var.database_cloud_database_management_details_user_name
    password_secret_id  = var.database_cloud_database_management_details_password_secret_id
  }
  enable_management     = var.database_cloud_database_management_details_enable_management
  port = var.cloud_database_management_port
  protocol = var.cloud_database_management_protocol
  role = var.cloud_database_management_role
  ssl_secret_id = oci_vault_secret.test_secret.id
}
```

## Argument Reference

The following arguments are supported:

* `database_id` - (Required) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `port` - (Optional) The port used to connect to the database.
* `protocol` - (Optional) Protocol used by the database connection.
* `role` - (Optional) The role of the user that will be connecting to the database.
* `ssl_secret_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [secret](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
* `private_end_point_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint.
* `service_name` - (Required) The name of the Oracle Database service that will be used to connect to the database.
* `management_type` - (Required) (Updatable) Specifies database management type
  enum:
  - `BASIC`
  - `ADVANCED`
* `credentaildetails` - (Required) (Updatable) Credential details to connect to the database
    * `user_name` - Database username
    * `password_secret_id` - Specific database username's password [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `enable_management` - (Required) (Updatable) Use this flag to enable/disable database management

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `character_set` - The character set for the database.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `connection_strings` - The Connection strings used to connect to the Oracle Database.
	* `all_connection_strings` - All connection strings to use to connect to the Database.
	* `cdb_default` - Host name based CDB Connection String.
	* `cdb_ip_default` - IP based CDB Connection String.
* `database_management_config` - The configuration of the Database Management service.
	* `management_status` - The status of the Database Management service.
	* `management_type` - The Database Management type.
* `database_software_image_id` - The database software image [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
* `db_backup_config` - Backup Options To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see [Getting Started with Policies](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm). 
	* `auto_backup_enabled` - If set to true, configures automatic backups. If you previously used RMAN or dbcli to configure backups and then you switch to using the Console or the API for backups, a new backup configuration is created and associated with your database. This means that you can no longer rely on your previously configured unmanaged backups to work.
	* `auto_backup_window` - Time window selected for initiating automatic backup for the database system. There are twelve available two-hour time windows. If no option is selected, a start time between 12:00 AM to 7:00 AM in the region of the database is automatically chosen. For example, if the user selects SLOT_TWO from the enum list, the automatic backup job will start in between 2:00 AM (inclusive) to 4:00 AM (exclusive).  Example: `SLOT_TWO` 
	* `backup_destination_details` - Backup destination details.
		* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup destination.
		* `internet_proxy` - Proxy URL to connect to object store.
		* `type` - Type of the database backup destination.
		* `vpc_password` - For a RECOVERY_APPLIANCE backup destination, the password for the VPC user that is used to access the Recovery Appliance.
		* `vpc_user` - For a RECOVERY_APPLIANCE backup destination, the Virtual Private Catalog (VPC) user that is used to access the Recovery Appliance.
	* `recovery_window_in_days` - Number of days between the current and the earliest point of recoverability covered by automatic backups. This value applies to automatic backups only. After a new automatic backup has been created, Oracle removes old automatic backups that are created before the window. When the value is updated, it is applied to all existing automatic backups. 
* `db_home_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
* `db_name` - The database name.
* `db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system.
* `db_unique_name` - A system-generated name for the database to ensure uniqueness within an Oracle Data Guard group (a primary database and its standby databases). The unique name cannot be changed. 
* `db_workload` - The database workload type.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
* `is_cdb` - True if the database is a container database.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
* `create` - (Defaults to 20 minutes), when creating the Database Management
* `update` - (Defaults to 20 minutes), when updating the Database Management
* `delete` - (Defaults to 20 minutes), when destroying the Database Management


## Import

Import is not supported for this resource.
