---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_data_guard_association"
sidebar_current: "docs-oci-resource-database-data_guard_association"
description: |-
  Provides the Data Guard Association resource in Oracle Cloud Infrastructure Database service
---

# oci_database_data_guard_association
This resource provides the Data Guard Association resource in Oracle Cloud Infrastructure Database service.

Creates a new Data Guard association.  A Data Guard association represents the replication relationship between the
specified database and a peer database. For more information, see [Using Oracle Data Guard](https://docs.cloud.oracle.com/iaas/Content/Database/Tasks/usingdataguard.htm).

All Oracle Cloud Infrastructure resources, including Data Guard associations, get an Oracle-assigned, unique ID
called an Oracle Cloud Identifier (OCID). When you create a resource, you can find its OCID in the response.
You can also retrieve a resource's OCID by using a List API operation on that resource type, or by viewing the
resource in the Console. For more information, see
[Resource Identifiers](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Example Usage

```hcl
resource "oci_database_data_guard_association" "test_data_guard_association" {
	#Required
	creation_type = var.data_guard_association_creation_type
	database_admin_password = var.data_guard_association_database_admin_password
	database_id = oci_database_database.test_database.id
	delete_standby_db_home_on_delete = var.data_guard_association_delete_standby_db_home_on_delete
	protection_mode = var.data_guard_association_protection_mode
	transport_type = var.data_guard_association_transport_type

	#Optional
	availability_domain = var.data_guard_association_availability_domain
	backup_network_nsg_ids = var.data_guard_association_backup_network_nsg_ids
	database_software_image_id = oci_database_database_software_image.test_database_software_image.id
	display_name = var.data_guard_association_display_name
	hostname = var.data_guard_association_hostname
	is_active_data_guard_enabled = var.data_guard_association_is_active_data_guard_enabled
	nsg_ids = var.data_guard_association_nsg_ids
	peer_db_home_id = oci_database_db_home.test_db_home.id
	peer_db_system_id = oci_database_db_system.test_db_system.id
	peer_db_unique_name = var.data_guard_association_peer_db_unique_name
	peer_sid_prefix = var.data_guard_association_peer_sid_prefix
	peer_vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id
	shape = var.data_guard_association_shape
	subnet_id = oci_core_subnet.test_subnet.id
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Applicable when creation_type=NewDbSystem) The name of the availability domain that the standby database DB system will be located in. For example- "Uocm:PHX-AD-1".
* `backup_network_nsg_ids` - (Applicable when creation_type=NewDbSystem) A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
* `creation_type` - (Required) Specifies whether to create the peer database in an existing DB system or in a new DB system. 
* `database_admin_password` - (Required) (Updatable) A strong password for the `SYS`, `SYSTEM`, and `PDB Admin` users to apply during standby creation.

    The password must contain no fewer than nine characters and include:
    * At least two uppercase characters.
    * At least two lowercase characters.
    * At least two numeric characters.
    * At least two special characters. Valid special characters include "_", "#", and "-" only.

    **The password MUST be the same as the primary admin password.** 
* `database_id` - (Required) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `database_software_image_id` - (Optional) The database software image [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Applicable only when creationType=`ExistingDbSystem` and when the existing database has Exadata shape.
* `delete_standby_db_home_on_delete` - (Required) (Updatable) if set to true the destroy operation will destroy the standby dbHome/dbSystem that is referenced in the Data Guard Association. The Data Guard Association gets destroyed when standby dbHome/dbSystem is terminated. Only `true` is supported at this time. If you change an argument that is used during the delete operation you must run `terraform apply` first so that that the change in the value is registered in the statefile before running `terraform destroy`. `terraform destroy` only looks at what is currently on the statefile and ignores the terraform configuration files. 
* `display_name` - (Applicable when creation_type=NewDbSystem) The user-friendly name of the DB system that will contain the the standby database. The display name does not have to be unique.
* `hostname` - (Applicable when creation_type=NewDbSystem) The hostname for the DB node.
* `is_active_data_guard_enabled` - (Optional) (Updatable) True if active Data Guard is enabled.
* `nsg_ids` - (Applicable when creation_type=NewDbSystem) A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that this resource belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). **NsgIds restrictions:**
    * Autonomous Databases with private access require at least 1 Network Security Group (NSG). The nsgIds array cannot be empty. 
* `peer_db_home_id` - (Applicable when creation_type=ExistingDbSystem | ExistingVmCluster) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB home in which to create the standby database. You must supply this value to create standby database with an existing DB home 
* `peer_db_system_id` - (Applicable when creation_type=ExistingDbSystem) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system in which to create the standby database. You must supply this value if creationType is `ExistingDbSystem`. 
* `peer_db_unique_name` - (Optional) Specifies the `DB_UNIQUE_NAME` of the peer database to be created. 
* `peer_sid_prefix` - (Optional) Specifies a prefix for the `Oracle SID` of the database to be created. 
* `peer_vm_cluster_id` - (Applicable when creation_type=ExistingVmCluster) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM Cluster in which to create the standby database. You must supply this value if creationType is `ExistingVmCluster`. 
* `protection_mode` - (Required) (Updatable) The protection mode to set up between the primary and standby databases. For more information, see [Oracle Data Guard Protection Modes](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000) in the Oracle Data Guard documentation.

    **IMPORTANT** - The only protection mode currently supported by the Database service is MAXIMUM_PERFORMANCE. 
* `shape` - (Applicable when creation_type=NewDbSystem) The virtual machine DB system shape to launch for the standby database in the Data Guard association. The shape determines the number of CPU cores and the amount of memory available for the DB system. Only virtual machine shapes are valid options. If you do not supply this parameter, the default shape is the shape of the primary DB system.

    To get a list of all shapes, use the [ListDbSystemShapes](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/DbSystemShapeSummary/ListDbSystemShapes) operation. 
* `subnet_id` - (Applicable when creation_type=NewDbSystem) The OCID of the subnet the DB system is associated with. **Subnet Restrictions:**
    * For 1- and 2-node RAC DB systems, do not use a subnet that overlaps with 192.168.16.16/28

    These subnets are used by the Oracle Clusterware private interconnect on the database instance. Specifying an overlapping subnet will cause the private interconnect to malfunction. This restriction applies to both the client subnet and backup subnet. 
* `transport_type` - (Required) (Updatable) The redo transport type to use for this Data Guard association.  Valid values depend on the specified `protectionMode`:
    * MAXIMUM_AVAILABILITY - SYNC or FASTSYNC
    * MAXIMUM_PERFORMANCE - ASYNC
    * MAXIMUM_PROTECTION - SYNC

    For more information, see [Redo Transport Services](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-redo-transport-services.htm#SBYDB00400) in the Oracle Data Guard documentation.

    **IMPORTANT** - The only transport type currently supported by the Database service is ASYNC. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `apply_lag` - The lag time between updates to the primary database and application of the redo data on the standby database, as computed by the reporting database.  Example: `9 seconds` 
* `apply_rate` - The rate at which redo logs are synced between the associated databases.  Example: `180 Mb per second` 
* `database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the reporting database.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Data Guard association.
* `is_active_data_guard_enabled` - True if active Data Guard is enabled.
* `lifecycle_details` - Additional information about the current lifecycleState, if available. 
* `peer_data_guard_association_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the peer database's Data Guard association.
* `peer_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated peer database.
* `peer_db_home_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home containing the associated peer database. 
* `peer_db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system containing the associated peer database. 
* `peer_role` - The role of the peer database in this Data Guard association.
* `protection_mode` - The protection mode of this Data Guard association. For more information, see [Oracle Data Guard Protection Modes](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000) in the Oracle Data Guard documentation. 
* `role` - The role of the reporting database in this Data Guard association.
* `state` - The current state of the Data Guard association.
* `time_created` - The date and time the Data Guard association was created.
* `transport_type` - The redo transport type used by this Data Guard association.  For more information, see [Redo Transport Services](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-redo-transport-services.htm#SBYDB00400) in the Oracle Data Guard documentation. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 2 hours), when creating the Data Guard Association
	* `update` - (Defaults to 2 hours), when updating the Data Guard Association
	* `delete` - (Defaults to 2 hours), when destroying the Data Guard Association


## Import

Import is not supported for this resource.

