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
	cpu_core_count = var.data_guard_association_cpu_core_count
	database_defined_tags = var.data_guard_association_database_defined_tags
	database_freeform_tags = var.data_guard_association_database_freeform_tags
	data_collection_options {

		#Optional
		is_diagnostics_events_enabled = var.data_guard_association_data_collection_options_is_diagnostics_events_enabled
		is_health_monitoring_enabled = var.data_guard_association_data_collection_options_is_health_monitoring_enabled
		is_incident_logs_enabled = var.data_guard_association_data_collection_options_is_incident_logs_enabled
	}
	database_software_image_id = oci_database_database_software_image.test_database_software_image.id
	db_system_defined_tags = var.data_guard_association_db_system_defined_tags
	db_system_freeform_tags = var.data_guard_association_db_system_freeform_tags
	display_name = var.data_guard_association_display_name
	domain = var.data_guard_association_domain
	fault_domains = var.data_guard_association_fault_domains
	hostname = var.data_guard_association_hostname
	is_active_data_guard_enabled = var.data_guard_association_is_active_data_guard_enabled
	license_model = var.data_guard_association_license_model
	node_count = var.data_guard_association_node_count
	nsg_ids = var.data_guard_association_nsg_ids
	peer_db_home_id = oci_database_db_home.test_db_home.id
	peer_db_system_id = oci_database_db_system.test_db_system.id
	peer_db_unique_name = var.data_guard_association_peer_db_unique_name
	peer_sid_prefix = var.data_guard_association_peer_sid_prefix
	peer_vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id
	private_ip = var.data_guard_association_private_ip
	shape = var.data_guard_association_shape
	storage_volume_performance_mode = var.data_guard_association_storage_volume_performance_mode
	subnet_id = oci_core_subnet.test_subnet.id
	time_zone = var.data_guard_association_time_zone
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Applicable when creation_type=NewDbSystem) The name of the availability domain that the standby database DB system will be located in. For example- "Uocm:PHX-AD-1".
* `backup_network_nsg_ids` - (Applicable when creation_type=NewDbSystem) A list of the [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems. 
* `cpu_core_count` - (Applicable when creation_type=NewDbSystem) The number of CPU cores available for AMD-based virtual machine DB systems.
* `creation_type` - (Required) Specifies whether to create the peer database in an existing DB system or in a new DB system. 
* `data_collection_options` - (Applicable when creation_type=NewDbSystem) Indicates user preferences for the various diagnostic collection options for the VM cluster/Cloud VM cluster/VMBM DBCS. 
	* `is_diagnostics_events_enabled` - (Applicable when creation_type=NewDbSystem) Indicates whether diagnostic collection is enabled for the VM cluster/Cloud VM cluster/VMBM DBCS. Enabling diagnostic collection allows you to receive Events service notifications for guest VM issues. Diagnostic collection also allows Oracle to provide enhanced service and proactive support for your Exadata system. You can enable diagnostic collection during VM cluster/Cloud VM cluster provisioning. You can also disable or enable it at any time using the `UpdateVmCluster` or `updateCloudVmCluster` API. 
	* `is_health_monitoring_enabled` - (Applicable when creation_type=NewDbSystem) Indicates whether health monitoring is enabled for the VM cluster / Cloud VM cluster / VMBM DBCS. Enabling health monitoring allows Oracle to collect diagnostic data and share it with its operations and support personnel. You may also receive notifications for some events. Collecting health diagnostics enables Oracle to provide proactive support and enhanced service for your system. Optionally enable health monitoring while provisioning a system. You can also disable or enable health monitoring anytime using the `UpdateVmCluster`, `UpdateCloudVmCluster` or `updateDbsystem` API. 
	* `is_incident_logs_enabled` - (Applicable when creation_type=NewDbSystem) Indicates whether incident logs and trace collection are enabled for the VM cluster / Cloud VM cluster / VMBM DBCS. Enabling incident logs collection allows Oracle to receive Events service notifications for guest VM issues, collect incident logs and traces, and use them to diagnose issues and resolve them. Optionally enable incident logs collection while provisioning a system. You can also disable or enable incident logs collection anytime using the `UpdateVmCluster`, `updateCloudVmCluster` or `updateDbsystem` API. 
* `database_admin_password` - (Required) (Updatable) A strong password for the `SYS`, `SYSTEM`, and `PDB Admin` users to apply during standby creation.

    The password must contain no fewer than nine characters and include:
    * At least two uppercase characters.
    * At least two lowercase characters.
    * At least two numeric characters.
    * At least two special characters. Valid special characters include "_", "#", and "-" only.

	**The password MUST be the same as the primary admin password.** 
* `database_defined_tags` - (Applicable when creation_type=NewDbSystem) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `database_freeform_tags` - (Applicable when creation_type=NewDbSystem) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `database_id` - (Required) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `database_software_image_id` - (Optional) The database software image [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Applicable only when creationType=`ExistingDbSystem` and when the existing database has Exadata shape.
* `db_system_defined_tags` - (Applicable when creation_type=NewDbSystem) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `db_system_freeform_tags` - (Applicable when creation_type=NewDbSystem) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `delete_standby_db_home_on_delete` - (Required) (Updatable) if set to true the destroy operation will destroy the standby dbHome/dbSystem that is referenced in the Data Guard Association. The Data Guard Association gets destroyed when standby dbHome/dbSystem is terminated. Only `true` is supported at this time. If you change an argument that is used during the delete operation you must run `terraform apply` first so that that the change in the value is registered in the statefile before running `terraform destroy`. `terraform destroy` only looks at what is currently on the statefile and ignores the terraform configuration files. 
* `display_name` - (Applicable when creation_type=NewDbSystem) The user-friendly name of the DB system that will contain the the standby database. The display name does not have to be unique.
* `domain` - (Applicable when creation_type=NewDbSystem) A domain name used for the DB system. If the Oracle-provided Internet and VCN Resolver is enabled for the specified subnet, the domain name for the subnet is used (do not provide one). Otherwise, provide a valid DNS domain name. Hyphens (-) are not permitted. 
* `fault_domains` - (Applicable when creation_type=NewDbSystem) A Fault Domain is a grouping of hardware and infrastructure within an availability domain. Fault Domains let you distribute your instances so that they are not on the same physical hardware within a single availability domain. A hardware failure or maintenance that affects one Fault Domain does not affect DB systems in other Fault Domains.

	If you do not specify the Fault Domain, the system selects one for you. To change the Fault Domain for a DB system, terminate it and launch a new DB system in the preferred Fault Domain.

	If the node count is greater than 1, you can specify which Fault Domains these nodes will be distributed into. The system assigns your nodes automatically to the Fault Domains you specify so that no Fault Domain contains more than one node.

	To get a list of Fault Domains, use the [ListFaultDomains](https://docs.cloud.oracle.com/iaas/api/#/en/identity/latest/FaultDomain/ListFaultDomains) operation in the Identity and Access Management Service API.

	Example: `FAULT-DOMAIN-1` 
* `hostname` - (Applicable when creation_type=NewDbSystem) The hostname for the DB node.
* `is_active_data_guard_enabled` - (Optional) (Updatable) True if active Data Guard is enabled.
* `license_model` - (Applicable when creation_type=NewDbSystem) The Oracle license model that applies to all the databases on the dataguard standby DB system. The default is LICENSE_INCLUDED. Bring your own license (BYOL) allows you to select the DB edition using the optional parameter, for Autonomous Database Serverless. 
* `node_count` - (Applicable when creation_type=NewDbSystem) The number of nodes to launch for the DB system of the standby in the Data Guard association. For a 2-node RAC virtual machine DB system, specify either 1 or 2. If you do not supply this parameter, the default is the node count of the primary DB system. 
* `nsg_ids` - (Applicable when creation_type=NewDbSystem) The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs. Setting this to an empty list removes all resources from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). **NsgIds restrictions:**
	* A network security group (NSG) is optional for Autonomous Databases with private access. The nsgIds list can be empty.
* `peer_db_home_id` - (Applicable when creation_type=ExistingDbSystem | ExistingVmCluster) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB home in which to create the standby database. You must supply this value to create standby database with an existing DB home 
* `peer_db_system_id` - (Applicable when creation_type=ExistingDbSystem) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system in which to create the standby database. You must supply this value if creationType is `ExistingDbSystem`. 
* `peer_db_unique_name` - (Optional) Specifies the `DB_UNIQUE_NAME` of the peer database to be created. 
* `peer_sid_prefix` - (Optional) Specifies a prefix for the `Oracle SID` of the database to be created. 
* `peer_vm_cluster_id` - (Applicable when creation_type=ExistingVmCluster) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM Cluster in which to create the standby database. You must supply this value if creationType is `ExistingVmCluster`. 
* `private_ip` - (Applicable when creation_type=NewDbSystem) The IPv4 address from the provided Oracle Cloud Infrastructure subnet which needs to be assigned to the VNIC. If not provided, it will be auto-assigned with an available IPv4 address from the subnet. 
* `protection_mode` - (Required) (Updatable) The protection mode to set up between the primary and standby databases. For more information, see [Oracle Data Guard Protection Modes](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000) in the Oracle Data Guard documentation.

    **IMPORTANT** - The only protection mode currently supported by the Database service is MAXIMUM_PERFORMANCE. 
* `shape` - (Applicable when creation_type=NewDbSystem) The virtual machine DB system shape to launch for the standby database in the Data Guard association. The shape determines the number of CPU cores and the amount of memory available for the DB system. Only virtual machine shapes are valid options. If you do not supply this parameter, the default shape is the shape of the primary DB system.

	To get a list of all shapes, use the [ListDbSystemShapes](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/DbSystemShapeSummary/ListDbSystemShapes) operation. 
* `storage_volume_performance_mode` - (Applicable when creation_type=NewDbSystem) The block storage volume performance level. Valid values are `BALANCED` and `HIGH_PERFORMANCE`. See [Block Volume Performance](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeperformance.htm) for more information.
* `subnet_id` - (Applicable when creation_type=NewDbSystem) The OCID of the subnet the DB system is associated with. **Subnet Restrictions:**
    * For 1- and 2-node RAC DB systems, do not use a subnet that overlaps with 192.168.16.16/28

	These subnets are used by the Oracle Clusterware private interconnect on the database instance. Specifying an overlapping subnet will cause the private interconnect to malfunction. This restriction applies to both the client subnet and backup subnet. 
* `time_zone` - (Applicable when creation_type=NewDbSystem) The time zone of the dataguard standby DB system. For details, see [DB System Time Zones](https://docs.cloud.oracle.com/iaas/Content/Database/References/timezones.htm).
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

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 2 hours), when creating the Data Guard Association
	* `update` - (Defaults to 2 hours), when updating the Data Guard Association
	* `delete` - (Defaults to 2 hours), when destroying the Data Guard Association


## Import

Import is not supported for this resource.

