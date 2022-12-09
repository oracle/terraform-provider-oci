---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_pluggable_database_pluggabledatabasemanagements_management"
sidebar_current: "docs-oci-resource-database-pluggable_database_pluggabledatabasemanagements_management"
description: |-
  Provides the Pluggable Database Pluggabledatabasemanagements Management resource in Oracle Cloud Infrastructure Database service
---

# oci_database_pluggable_database_pluggabledatabasemanagements_management
This resource provides the Pluggable Database Pluggabledatabasemanagements Management resource in Oracle Cloud Infrastructure Database service.

Enables the Database Management service for an Oracle Pluggable Database located in Oracle Cloud Infrastructure. This service allows the pluggable database to access tools including Metrics and Performance hub. Database Management is enabled at the pluggable database (PDB) level.

## Example Usage

```hcl
resource "oci_database_pluggable_database_pluggabledatabasemanagements_management" "test_pluggable_database_pluggabledatabasemanagements_management" {
	#Required
	pluggable_database_id = oci_database_pluggable_database.test_pluggable_database.id
	enable_pluggabledatabasemanagement = var.enable_pluggabledatabasemanagement

	#Optional
	credential_details {

		#Optional
		password_secret_id = oci_vault_secret.test_secret.id
		user_name = oci_identity_user.test_user.name
	}
	private_end_point_id = oci_database_private_end_point.test_private_end_point.id
	service_name = oci_core_service.test_service.name
	port = var.pluggable_database_pluggabledatabasemanagements_management_port
	protocol = var.pluggable_database_pluggabledatabasemanagements_management_protocol
	role = var.pluggable_database_pluggabledatabasemanagements_management_role
	ssl_secret_id = oci_vault_secret.test_secret.id
}
```

## Argument Reference

The following arguments are supported:

* `credential_details` - (Optional) Data for the credential used to connect to the database. 
	* `password_secret_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [secret](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
	* `user_name` - (Optional) The name of the Oracle Database user that will be used to connect to the database.
* `pluggable_database_id` - (Required) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `port` - (Optional) The port used to connect to the pluggable database.
* `private_end_point_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint. 
* `protocol` - (Optional) Protocol used by the database connection.
* `role` - (Optional) The role of the user that will be connecting to the pluggable database.
* `service_name` - (Optional) The name of the Oracle Database service that will be used to connect to the database.
* `ssl_secret_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [secret](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
* `enable_pluggabledatabasemanagement` - (Required) (Updatable) A required field when set to `true` calls enable action and when set to `false` calls disable action.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `connection_strings` - Connection strings to connect to an Oracle Pluggable Database. 
	* `all_connection_strings` - All connection strings to use to connect to the pluggable database.
	* `pdb_default` - A host name-based PDB connection string.
	* `pdb_ip_default` - An IP-based PDB connection string.
* `container_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the CDB.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the pluggable database.
* `is_restricted` - The restricted mode of the pluggable database. If a pluggable database is opened in restricted mode, the user needs both create a session and have restricted session privileges to connect to it. 
* `lifecycle_details` - Detailed message for the lifecycle state.
* `open_mode` - The mode that pluggable database is in. Open mode can only be changed to READ_ONLY or MIGRATE directly from the backend (within the Oracle Database software). 
* `pdb_name` - The name for the pluggable database (PDB). The name is unique in the context of a [container database](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/Database/). The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. The pluggable database name should not be same as the container database name.
* `pluggable_database_management_config` - The configuration of the Pluggable Database Management service.
	* `management_status` - The status of the Pluggable Database Management service.
* `state` - The current state of the pluggable database.
* `time_created` - The date and time the pluggable database was created.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Pluggable Database Pluggabledatabasemanagements Management
	* `update` - (Defaults to 20 minutes), when updating the Pluggable Database Pluggabledatabasemanagements Management
	* `delete` - (Defaults to 20 minutes), when destroying the Pluggable Database Pluggabledatabasemanagements Management
