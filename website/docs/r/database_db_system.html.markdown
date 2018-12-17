---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_db_system"
sidebar_current: "docs-oci-resource-database-db_system"
description: |-
  Provides the Db System resource in Oracle Cloud Infrastructure Database service
---

# oci_database_db_system
This resource provides the Db System resource in Oracle Cloud Infrastructure Database service.

Launches a new DB system in the specified compartment and availability domain. The Oracle
Database edition that you specify applies to all the databases on that DB system. The selected edition cannot be changed.

An initial database is created on the DB system based on the request parameters you provide and some default
options. For more information,
see [Default Options for the Initial Database](https://docs.cloud.oracle.com/iaas/Content/Database/Tasks/launchingDB.htm#DefaultOptionsfortheInitialDatabase).

The DB System will include a command line interface (CLI) that you can use to create additional databases and
manage existing databases. For more information, see the
[Oracle Database CLI Reference](https://docs.cloud.oracle.com/iaas/Content/Database/References/dbacli.htm).


## Example Usage

```hcl
resource "oci_database_db_system" "test_db_system" {
	#Required
	availability_domain = "${var.db_system_availability_domain}"
	compartment_id = "${var.compartment_id}"
	database_edition = "${var.db_system_database_edition}"
	db_home {
		#Required
		database {
			#Required
			admin_password = "${var.db_system_db_home_database_admin_password}"

			#Optional
			backup_id = "${oci_database_backup.test_backup.id}"
			backup_tde_password = "${var.db_system_db_home_database_backup_tde_password}"
			character_set = "${var.db_system_db_home_database_character_set}"
			db_backup_config {

				#Optional
				auto_backup_enabled = "${var.db_system_db_home_database_db_backup_config_auto_backup_enabled}"
			}
			db_name = "${var.db_system_db_home_database_db_name}"
			db_workload = "${var.db_system_db_home_database_db_workload}"
			defined_tags = "${var.db_system_db_home_database_defined_tags}"
			freeform_tags = "${var.db_system_db_home_database_freeform_tags}"
			ncharacter_set = "${var.db_system_db_home_database_ncharacter_set}"
			pdb_name = "${var.db_system_db_home_database_pdb_name}"
		}

		#Optional
		db_version = "${var.db_system_db_home_db_version}"
		display_name = "${var.db_system_db_home_display_name}"
	}
	hostname = "${var.db_system_hostname}"
	shape = "${var.db_system_shape}"
	ssh_public_keys = "${var.db_system_ssh_public_keys}"
	subnet_id = "${oci_database_subnet.test_subnet.id}"

	#Optional
	backup_subnet_id = "${oci_database_backup_subnet.test_backup_subnet.id}"
	cluster_name = "${var.db_system_cluster_name}"
	cpu_core_count = "${var.db_system_cpu_core_count}"
	data_storage_percentage = "${var.db_system_data_storage_percentage}"
	data_storage_size_in_gb = "${var.db_system_data_storage_size_in_gb}"
	defined_tags = {"Operations.CostCenter"= "42"}
	disk_redundancy = "${var.db_system_disk_redundancy}"
	display_name = "${var.db_system_display_name}"
	domain = "${var.db_system_domain}"
	freeform_tags = {"Department"= "Finance"}
	license_model = "${var.db_system_license_model}"
	node_count = "${var.db_system_node_count}"
	source = "${var.db_system_source}"
	sparse_diskgroup = "${var.db_system_sparse_diskgroup}"
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The availability domain where the DB system is located.
* `backup_subnet_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet the DB system is associated with. Applicable only to Exadata DB systems.

	**Subnet Restrictions:** See the subnet restrictions information for **subnetId**. 
* `cluster_name` - (Optional) The cluster name for Exadata and 2-node RAC virtual machine DB systems. The cluster name must begin with an an alphabetic character, and may contain hyphens (-). Underscores (_) are not permitted. The cluster name can be no longer than 11 characters and is not case sensitive. 
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment the DB system  belongs in.
* `cpu_core_count` - (Required) (Updatable) The number of CPU cores to enable for a bare metal or Exadata DB system. The valid values depend on the specified shape:
	* BM.DenseIO1.36 - Specify a multiple of 2, from 2 to 36.
	* BM.DenseIO2.52 - Specify a multiple of 2, from 2 to 52.
	* Exadata.Quarter1.84 - Specify a multiple of 2, from 22 to 84.
	* Exadata.Half1.168 - Specify a multiple of 4, from 44 to 168.
	* Exadata.Full1.336 - Specify a multiple of 8, from 88 to 336.
	* Exadata.Quarter2.92 - Specify a multiple of 2, from 0 to 92.
	* Exadata.Half2.184 - Specify a multiple of 4, from 0 to 184.
	* Exadata.Full2.368 - Specify a multiple of 8, from 0 to 368.

	This parameter is not used for virtual machine DB systems because virtual machine DB systems have a set number of cores for each shape. For information about the number of cores for a virtual machine DB system shape, see [Virtual Machine DB Systems](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/overview.htm#virtualmachine) 
* `data_storage_percentage` - (Optional) The percentage assigned to DATA storage (user data and database files). The remaining percentage is assigned to RECO storage (database redo logs, archive logs, and recovery manager backups). Specify 80 or 40. The default is 80 percent assigned to DATA storage. Not applicable for virtual machine DB systems. 
* `data_storage_size_in_gb` - (Optional) (Updatable) Size (in GB) of the initial data volume that will be created and attached to a virtual machine DB system. You can scale up storage after provisioning, as needed. Note that the total storage size attached will be more than the amount you specify to allow for REDO/RECO space and software volume. 
* `database_edition` - (Required) The Oracle Database Edition that applies to all the databases on the DB system. Exadata DB systems and 2-node RAC DB systems require ENTERPRISE_EDITION_EXTREME_PERFORMANCE. 
* `db_home` - (Required) 
	* `database` - (Required) 
		* `admin_password` - (Required) A strong password for SYS, SYSTEM, PDB Admin and TDE Wallet. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numbers, and two special characters. The special characters must be _, \#, or -.
		* `backup_id` - (Required when source=DB_BACKUP) The backup [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
		* `backup_tde_password` - (Required when source=DB_BACKUP) The password to open the TDE wallet.
		* `character_set` - (Applicable when source=NONE) The character set for the database.  The default is AL32UTF8. Allowed values are:

			AL32UTF8, AR8ADOS710, AR8ADOS720, AR8APTEC715, AR8ARABICMACS, AR8ASMO8X, AR8ISO8859P6, AR8MSWIN1256, AR8MUSSAD768, AR8NAFITHA711, AR8NAFITHA721, AR8SAKHR706, AR8SAKHR707, AZ8ISO8859P9E, BG8MSWIN, BG8PC437S, BLT8CP921, BLT8ISO8859P13, BLT8MSWIN1257, BLT8PC775, BN8BSCII, CDN8PC863, CEL8ISO8859P14, CL8ISO8859P5, CL8ISOIR111, CL8KOI8R, CL8KOI8U, CL8MACCYRILLICS, CL8MSWIN1251, EE8ISO8859P2, EE8MACCES, EE8MACCROATIANS, EE8MSWIN1250, EE8PC852, EL8DEC, EL8ISO8859P7, EL8MACGREEKS, EL8MSWIN1253, EL8PC437S, EL8PC851, EL8PC869, ET8MSWIN923, HU8ABMOD, HU8CWI2, IN8ISCII, IS8PC861, IW8ISO8859P8, IW8MACHEBREWS, IW8MSWIN1255, IW8PC1507, JA16EUC, JA16EUCTILDE, JA16SJIS, JA16SJISTILDE, JA16VMS, KO16KSC5601, KO16KSCCS, KO16MSWIN949, LA8ISO6937, LA8PASSPORT, LT8MSWIN921, LT8PC772, LT8PC774, LV8PC1117, LV8PC8LR, LV8RST104090, N8PC865, NE8ISO8859P10, NEE8ISO8859P4, RU8BESTA, RU8PC855, RU8PC866, SE8ISO8859P3, TH8MACTHAIS, TH8TISASCII, TR8DEC, TR8MACTURKISHS, TR8MSWIN1254, TR8PC857, US7ASCII, US8PC437, UTF8, VN8MSWIN1258, VN8VN3, WE8DEC, WE8DG, WE8ISO8859P1, WE8ISO8859P15, WE8ISO8859P9, WE8MACROMAN8S, WE8MSWIN1252, WE8NCR4970, WE8NEXTSTEP, WE8PC850, WE8PC858, WE8PC860, WE8ROMAN8, ZHS16CGB231280, ZHS16GBK, ZHT16BIG5, ZHT16CCDC, ZHT16DBT, ZHT16HKSCS, ZHT16MSWIN950, ZHT32EUC, ZHT32SOPS, ZHT32TRIS 
		* `db_backup_config` - (Applicable when source=NONE) 
			* `auto_backup_enabled` - (Applicable when source=NONE) If set to true, configures automatic backups. If you previously used RMAN or dbcli to configure backups and then you switch to using the Console or the API for backups, a new backup configuration is created and associated with your database. This means that you can no longer rely on your previously configured unmanaged backups to work.
		* `db_name` - (Required when source=NONE) The database name. The name must begin with an alphabetic character and can contain a maximum of eight alphanumeric characters. Special characters are not permitted.
		* `db_workload` - (Applicable when source=NONE) The database workload type.
		* `defined_tags` - (Applicable when source=NONE) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
		* `freeform_tags` - (Applicable when source=NONE) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
		* `ncharacter_set` - (Applicable when source=NONE) The national character set for the database.  The default is AL16UTF16. Allowed values are: AL16UTF16 or UTF8. 
		* `pdb_name` - (Applicable when source=NONE) The name of the pluggable database. The name must begin with an alphabetic character and can contain a maximum of eight alphanumeric characters. Special characters are not permitted. Pluggable database should not be same as database name.
	* `db_version` - (Required when source=NONE) A valid Oracle Database version. To get a list of supported versions, use the [ListDbVersions](https://docs.cloud.oracle.com/iaas/api/#/en/database/20160918/DbVersion/ListDbVersions) operation.
	* `display_name` - (Optional) The user-provided name of the database home.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `disk_redundancy` - (Optional) The type of redundancy configured for the DB system. Normal is 2-way redundancy, recommended for test and development systems. High is 3-way redundancy, recommended for production systems. 
* `display_name` - (Optional) The user-friendly name for the DB system. The name does not have to be unique.
* `domain` - (Optional) A domain name used for the DB system. If the Oracle-provided Internet and VCN Resolver is enabled for the specified subnet, the domain name for the subnet is used (do not provide one). Otherwise, provide a valid DNS domain name. Hyphens (-) are not permitted. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `hostname` - (Required) The hostname for the DB system. The hostname must begin with an alphabetic character and can contain a maximum of 30 alphanumeric characters, including hyphens (-).

	The maximum length of the combined hostname and domain is 63 characters.

	**Note:** The hostname must be unique within the subnet. If it is not unique, the DB system will fail to provision. 
* `license_model` - (Optional) The Oracle license model that applies to all the databases on the DB system. The default is LICENSE_INCLUDED. Allowed values are: LICENSE_INCLUDED, BRING_YOUR_OWN_LICENSE.
* `node_count` - (Optional) The number of nodes to launch for a 2-node RAC virtual machine DB system. 
* `shape` - (Required) The shape of the DB system. The shape determines resources allocated to the DB system.
	* For virtual machine shapes, the number of CPU cores and memory
	* For bare metal and Exadata shapes, the number of CPU cores, memory, and storage

	To get a list of shapes, use the [ListDbSystemShapes](https://docs.cloud.oracle.com/iaas/api/#/en/database/20160918/DbSystemShapeSummary/ListDbSystemShapes) operation. 
* `source` - (Optional) The source of the database: NONE for creating a new database. DB_BACKUP for creating a new database by restoring from a backup. The default is NONE. 
* `sparse_diskgroup` - (Optional) If true, Sparse Diskgroup is configured for Exadata dbsystem. If False, Sparse diskgroup is not configured. 
* `ssh_public_keys` - (Required) (Updatable) The public key portion of the key pair to use for SSH access to the DB system. Multiple public keys can be provided. The length of the combined keys cannot exceed 10,000 characters.
* `subnet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the DB system is associated with.

	**Subnet Restrictions:**
	* For bare metal DB systems and for single node virtual machine DB systems, do not use a subnet that overlaps with 192.168.16.16/28.
	* For Exadata and virtual machine 2-node RAC DB systems, do not use a subnet that overlaps with 192.168.128.0/20.

	These subnets are used by the Oracle Clusterware private interconnect on the database instance. Specifying an overlapping subnet will cause the private interconnect to malfunction. This restriction applies to both the client subnet and the backup subnet. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The name of the availability domain that the DB system is located in.
* `backup_subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup network subnet the DB system is associated with. Applicable only to Exadata DB systems.

	**Subnet Restriction:** See the subnet restrictions information for **subnetId**. 
* `cluster_name` - The cluster name for Exadata and 2-node RAC virtual machine DB systems. The cluster name must begin with an an alphabetic character, and may contain hyphens (-). Underscores (_) are not permitted. The cluster name can be no longer than 11 characters and is not case sensitive. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `cpu_core_count` - The number of CPU cores enabled on the DB system.
* `data_storage_percentage` - The percentage assigned to DATA storage (user data and database files). The remaining percentage is assigned to RECO storage (database redo logs, archive logs, and recovery manager backups). Accepted values are 40 and 80. The default is 80 percent assigned to DATA storage. Not applicable for virtual machine DB systems. 
* `data_storage_size_in_gb` - The data storage size, in gigabytes, that is currently available to the DB system. Applies only for virtual machine DB systems. 
* `database_edition` - The Oracle Database edition that applies to all the databases on the DB system. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `disk_redundancy` - The type of redundancy configured for the DB system. NORMAL is 2-way redundancy. HIGH is 3-way redundancy. 
* `display_name` - The user-friendly name for the DB system. The name does not have to be unique.
* `domain` - The domain name for the DB system.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `hostname` - The hostname for the DB system.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system.
* `last_patch_history_entry_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last patch history. This value is updated as soon as a patch operation starts.
* `license_model` - The Oracle license model that applies to all the databases on the DB system. The default is LICENSE_INCLUDED. 
* `lifecycle_details` - Additional information about the current lifecycleState.
* `listener_port` - The port number configured for the listener on the DB system.
* `node_count` - The number of nodes in the DB system. For RAC DB systems, the value is greater than 1. 
* `reco_storage_size_in_gb` - The RECO/REDO storage size, in gigabytes, that is currently allocated to the DB system. Applies only for virtual machine DB systems. 
* `scan_dns_record_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DNS record for the SCAN IP addresses that are associated with the DB system. 
* `scan_ip_ids` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Single Client Access Name (SCAN) IP addresses associated with the DB system. SCAN IP addresses are typically used for load balancing and are not assigned to any interface. Oracle Clusterware directs the requests to the appropriate nodes in the cluster.

	**Note:** For a single-node DB system, this list is empty. 
* `shape` - The shape of the DB system. The shape determines resources to allocate to the DB system.
	* For virtual machine shapes, the number of CPU cores and memory
	* For bare metal and Exadata shapes, the number of CPU cores, storage, and memory 
* `sparse_diskgroup` - True, if Sparse Diskgroup is configured for Exadata dbsystem, False, if Sparse diskgroup was not configured. 
* `ssh_public_keys` - The public key portion of one or more key pairs used for SSH access to the DB system.
* `state` - The current state of the DB system.
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the DB system is associated with.

	**Subnet Restrictions:**
	* For bare metal DB systems and for single node virtual machine DB systems, do not use a subnet that overlaps with 192.168.16.16/28.
	* For Exadata and virtual machine 2-node RAC DB systems, do not use a subnet that overlaps with 192.168.128.0/20.

	These subnets are used by the Oracle Clusterware private interconnect on the database instance. Specifying an overlapping subnet will cause the private interconnect to malfunction. This restriction applies to both the client subnet and backup subnet. 
* `time_created` - The date and time the DB system was created.
* `version` - The Oracle Database version of the DB system.
* `vip_ids` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the virtual IP (VIP) addresses associated with the DB system. The Cluster Ready Services (CRS) creates and maintains one VIP address for each node in the DB system to enable failover. If one node fails, the VIP is reassigned to another active node in the cluster.

	**Note:** For a single-node DB system, this list is empty. 

## Import

DbSystems can be imported using the `id`, e.g.

```
$ terraform import oci_database_db_system.test_db_system "id"
```

