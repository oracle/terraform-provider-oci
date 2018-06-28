# oci_database_db_system

## DbSystem Resource

### DbSystem Reference

The following attributes are exported:

* `availability_domain` - The name of the Availability Domain that the DB System is located in.
* `backup_subnet_id` - The OCID of the backup network subnet the DB System is associated with. Applicable only to Exadata.  **Subnet Restriction:** See above subnetId's 'Subnet Restriction'. to malfunction. 
* `cluster_name` - Cluster name for Exadata and 2-node RAC DB Systems. The cluster name must begin with an an alphabetic character, and may contain hyphens (-). Underscores (_) are not permitted. The cluster name can be no longer than 11 characters and is not case sensitive. 
* `compartment_id` - The OCID of the compartment.
* `cpu_core_count` - The number of CPU cores enabled on the DB System.
* `data_storage_percentage` - The percentage assigned to DATA storage (user data and database files). The remaining percentage is assigned to RECO storage (database redo logs, archive logs, and recovery manager backups). Accepted values are 40 and 80. 
* `data_storage_size_in_gb` - Data storage size, in GBs, that is currently available to the DB system. This is applicable only for VM-based DBs. 
* `database_edition` - The Oracle Database Edition that applies to all the databases on the DB System. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `disk_redundancy` - The type of redundancy configured for the DB System. Normal is 2-way redundancy. High is 3-way redundancy. 
* `display_name` - The user-friendly name for the DB System. It does not have to be unique.
* `domain` - The domain name for the DB System.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `hostname` - The host name for the DB Node.
* `id` - The OCID of the DB System.
* `last_patch_history_entry_id` - The OCID of the last patch history. This is updated as soon as a patch operation is started.
* `license_model` - The Oracle license model that applies to all the databases on the DB System. The default is LICENSE_INCLUDED. 
* `lifecycle_details` - Additional information about the current lifecycleState.
* `listener_port` - The port number configured for the listener on the DB System.
* `node_count` - Number of nodes in this DB system. For RAC DBs, this will be greater than 1. 
* `reco_storage_size_in_gb` - RECO/REDO storage size, in GBs, that is currently allocated to the DB system. This is applicable only for VM-based DBs. 
* `scan_dns_record_id` - The OCID of the DNS record for the SCAN IP addresses that are associated with the DB System. 
* `scan_ip_ids` - The OCID of the Single Client Access Name (SCAN) IP addresses associated with the DB System. SCAN IP addresses are typically used for load balancing and are not assigned to any interface. Clusterware directs the requests to the appropriate nodes in the cluster.  - For a single-node DB System, this list is empty. 
* `shape` - The shape of the DB System. The shape determines resources to allocate to the DB system - CPU cores and memory for VM shapes; CPU cores, memory and storage for non-VM (or bare metal) shapes.
* `ssh_public_keys` - The public key portion of one or more key pairs used for SSH access to the DB System.
* `state` - The current state of the DB System.
* `subnet_id` - The OCID of the subnet the DB System is associated with.  **Subnet Restrictions:** - For single node and 2-node (RAC) DB Systems, do not use a subnet that overlaps with 192.168.16.16/28 - For Exadata and VM-based RAC DB Systems, do not use a subnet that overlaps with 192.168.128.0/20  These subnets are used by the Oracle Clusterware private interconnect on the database instance. Specifying an overlapping subnet will cause the private interconnect to malfunction. This restriction applies to both the client subnet and backup subnet. 
* `time_created` - The date and time the DB System was created.
* `version` - The version of the DB System.
* `vip_ids` - The OCID of the virtual IP (VIP) addresses associated with the DB System. The Cluster Ready Services (CRS) creates and maintains one VIP address for each node in the DB System to enable failover. If one node fails, the VIP is reassigned to another active node in the cluster.  - For a single-node DB System, this list is empty. 



### Create Operation
Launches a new DB System in the specified compartment and Availability Domain. You'll specify a single Oracle
Database Edition that applies to all the databases on that DB System. The selected edition cannot be changed.

An initial database is created on the DB System based on the request parameters you provide and some default
options. For more information,
see [Default Options for the Initial Database](https://docs.us-phoenix-1.oraclecloud.com/Content/Database/Tasks/launchingDB.htm#Default_Options_for_the_Initial_Database).

The DB System will include a command line interface (CLI) that you can use to create additional databases and
manage existing databases. For more information, see the
[Oracle Database CLI Reference](https://docs.us-phoenix-1.oraclecloud.com/Content/Database/References/odacli.htm#Oracle_Database_CLI_Reference).


The following arguments are supported:

* `availability_domain` - (Required) The Availability Domain where the DB System is located.
* `backup_subnet_id` - (Optional) The OCID of the backup network subnet the DB System is associated with. Applicable only to Exadata.  **Subnet Restrictions:** See above subnetId's **Subnet Restriction**. 
* `cluster_name` - (Optional) Cluster name for Exadata and 2-node RAC DB Systems. The cluster name must begin with an an alphabetic character, and may contain hyphens (-). Underscores (_) are not permitted. The cluster name can be no longer than 11 characters and is not case sensitive. 
* `compartment_id` - (Required) The Oracle Cloud ID (OCID) of the compartment the DB System  belongs in.
* `cpu_core_count` - (Optional) The number of CPU cores to enable. The valid values depend on the specified shape:  - BM.DenseIO1.36 and BM.HighIO1.36 - Specify a multiple of 2, from 2 to 36. - BM.RACLocalStorage1.72 - Specify a multiple of 4, from 4 to 72. - Exadata.Quarter1.84 - Specify a multiple of 2, from 22 to 84. - Exadata.Half1.168 - Specify a multiple of 4, from 44 to 168. - Exadata.Full1.336 - Specify a multiple of 8, from 88 to 336.  For VM DB systems, the core count is inferred from the specific VM shape chosen, so this parameter is not used. 
* `data_storage_percentage` - (Optional) The percentage assigned to DATA storage (user data and database files). The remaining percentage is assigned to RECO storage (database redo logs, archive logs, and recovery manager backups). Specify 80 or 40. The default is 80 percent assigned to DATA storage. This is not applicable for VM based DB systems. 
* `data_storage_size_in_gb` - (Optional) Size, in GBs, of the initial data volume that will be created and attached to VM-shape based DB system. This storage can later be scaled up if needed. Note that the total storage size attached will be more than what is requested, to account for REDO/RECO space and software volume. 
* `database_edition` - (Required) The Oracle Database Edition that applies to all the databases on the DB System. Exadata DB Systems and 2-node RAC DB Systems require ENTERPRISE_EDITION_EXTREME_PERFORMANCE. 
* `db_home` - (Required) 
	* `database` - (Required) 
		* `admin_password` - (Required) A strong password for SYS, SYSTEM, PDB Admin and TDE Wallet. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numbers, and two special characters. The special characters must be _, \#, or -.
		* `backup_id` - (Required when source=DB_BACKUP) The backup OCID.
		* `backup_tde_password` - (Required when source=DB_BACKUP) The password to open the TDE wallet.
		* `character_set` - (Optional) The character set for the database.  The default is AL32UTF8. Allowed values are:  AL32UTF8, AR8ADOS710, AR8ADOS720, AR8APTEC715, AR8ARABICMACS, AR8ASMO8X, AR8ISO8859P6, AR8MSWIN1256, AR8MUSSAD768, AR8NAFITHA711, AR8NAFITHA721, AR8SAKHR706, AR8SAKHR707, AZ8ISO8859P9E, BG8MSWIN, BG8PC437S, BLT8CP921, BLT8ISO8859P13, BLT8MSWIN1257, BLT8PC775, BN8BSCII, CDN8PC863, CEL8ISO8859P14, CL8ISO8859P5, CL8ISOIR111, CL8KOI8R, CL8KOI8U, CL8MACCYRILLICS, CL8MSWIN1251, EE8ISO8859P2, EE8MACCES, EE8MACCROATIANS, EE8MSWIN1250, EE8PC852, EL8DEC, EL8ISO8859P7, EL8MACGREEKS, EL8MSWIN1253, EL8PC437S, EL8PC851, EL8PC869, ET8MSWIN923, HU8ABMOD, HU8CWI2, IN8ISCII, IS8PC861, IW8ISO8859P8, IW8MACHEBREWS, IW8MSWIN1255, IW8PC1507, JA16EUC, JA16EUCTILDE, JA16SJIS, JA16SJISTILDE, JA16VMS, KO16KSCCS, KO16MSWIN949, LA8ISO6937, LA8PASSPORT, LT8MSWIN921, LT8PC772, LT8PC774, LV8PC1117, LV8PC8LR, LV8RST104090, N8PC865, NE8ISO8859P10, NEE8ISO8859P4, RU8BESTA, RU8PC855, RU8PC866, SE8ISO8859P3, TH8MACTHAIS, TH8TISASCII, TR8DEC, TR8MACTURKISHS, TR8MSWIN1254, TR8PC857, US7ASCII, US8PC437, UTF8, VN8MSWIN1258, VN8VN3, WE8DEC, WE8DG, WE8ISO8859P15, WE8ISO8859P9, WE8MACROMAN8S, WE8MSWIN1252, WE8NCR4970, WE8NEXTSTEP, WE8PC850, WE8PC858, WE8PC860, WE8ROMAN8, ZHS16CGB231280, ZHS16GBK, ZHT16BIG5, ZHT16CCDC, ZHT16DBT, ZHT16HKSCS, ZHT16MSWIN950, ZHT32EUC, ZHT32SOPS, ZHT32TRIS 
		* `db_backup_config` - (Optional) 
			* `auto_backup_enabled` - (Optional) If set to true, configures automatic backups. If you previously used RMAN or dbcli to configure backups and then you switch to using the Console or the API for backups, a new backup configuration is created and associated with your database. This means that you can no longer rely on your previously configured unmanaged backups to work.
		* `db_name` - (Required) The database name. It must begin with an alphabetic character and can contain a maximum of eight alphanumeric characters. Special characters are not permitted.
		* `db_workload` - (Optional) Database workload type.
		* `ncharacter_set` - (Optional) National character set for the database.  The default is AL16UTF16. Allowed values are: AL16UTF16 or UTF8. 
		* `pdb_name` - (Optional) Pluggable database name. It must begin with an alphabetic character and can contain a maximum of eight alphanumeric characters. Special characters are not permitted. Pluggable database should not be same as database name.
	* `db_version` - (Required) A valid Oracle database version. To get a list of supported versions, use the [ListDbVersions](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/database/20160918/DbVersion/ListDbVersions) operation.
	* `display_name` - (Optional) The user-provided name of the database home.
* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `disk_redundancy` - (Optional) The type of redundancy configured for the DB System. Normal is 2-way redundancy, recommended for test and development systems. High is 3-way redundancy, recommended for production systems. 
* `display_name` - (Optional) The user-friendly name for the DB System. It does not have to be unique.
* `domain` - (Optional) A domain name used for the DB System. If the Oracle-provided Internet and VCN Resolver is enabled for the specified subnet, the domain name for the subnet is used (don't provide one). Otherwise, provide a valid DNS domain name. Hyphens (-) are not permitted. 
* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `hostname` - (Required) The host name for the DB System. The host name must begin with an alphabetic character and can contain a maximum of 30 alphanumeric characters, including hyphens (-).  The maximum length of the combined hostname and domain is 63 characters.  **Note:** The hostname must be unique within the subnet. If it is not unique, the DB System will fail to provision. 
* `license_model` - (Optional) The Oracle license model that applies to all the databases on the DB System. The default is LICENSE_INCLUDED. 
* `node_count` - (Optional) Number of nodes to launch for a VM-shape based RAC DB system. 
* `shape` - (Required) The shape of the DB System. The shape determines resources allocated to the DB System - CPU cores and memory for VM shapes; CPU cores, memory and storage for non-VM (or bare metal) shapes. To get a list of shapes, use the [ListDbSystemShapes](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/database/20160918/DbSystemShapeSummary/ListDbSystemShapes) operation.
* `source` - (Optional) Source of database:   NONE for creating a new database   DB_BACKUP for creating a new database by restoring a backup 
* `ssh_public_keys` - (Required) The public key portion of the key pair to use for SSH access to the DB System. Multiple public keys can be provided. The length of the combined keys cannot exceed 10,000 characters.
* `subnet_id` - (Required) The OCID of the subnet the DB System is associated with.  **Subnet Restrictions:** - For single node and 2-node (RAC) DB Systems, do not use a subnet that overlaps with 192.168.16.16/28 - For Exadata and VM-based RAC DB Systems, do not use a subnet that overlaps with 192.168.128.0/20  These subnets are used by the Oracle Clusterware private interconnect on the database instance. Specifying an overlapping subnet will cause the private interconnect to malfunction. This restriction applies to both the client subnet and backup subnet. 


### Update Operation
Updates the properties of a DB System, such as the CPU core count.

The following arguments support updates:
* `cpu_core_count` - The number of CPU cores to enable. The valid values depend on the specified shape:  - BM.DenseIO1.36 and BM.HighIO1.36 - Specify a multiple of 2, from 2 to 36. - BM.RACLocalStorage1.72 - Specify a multiple of 4, from 4 to 72. - Exadata.Quarter1.84 - Specify a multiple of 2, from 22 to 84. - Exadata.Half1.168 - Specify a multiple of 4, from 44 to 168. - Exadata.Full1.336 - Specify a multiple of 8, from 88 to 336.  For VM DB systems, the core count is inferred from the specific VM shape chosen, so this parameter is not used. 
* `data_storage_size_in_gb` - Size, in GBs, of the initial data volume that will be created and attached to VM-shape based DB system. This storage can later be scaled up if needed. Note that the total storage size attached will be more than what is requested, to account for REDO/RECO space and software volume. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `ssh_public_keys` - The public key portion of the key pair to use for SSH access to the DB System. Multiple public keys can be provided. The length of the combined keys cannot exceed 10,000 characters.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

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
			backup_id = "${oci_database_backup.test_backup.id}"
			backup_tde_password = "${var.db_system_db_home_database_backup_tde_password}"
			db_name = "${var.db_system_db_home_database_db_name}"

			#Optional
			character_set = "${var.db_system_db_home_database_character_set}"
			db_backup_config {

				#Optional
				auto_backup_enabled = "${var.db_system_db_home_database_db_backup_config_auto_backup_enabled}"
			}
			db_workload = "${var.db_system_db_home_database_db_workload}"
			ncharacter_set = "${var.db_system_db_home_database_ncharacter_set}"
			pdb_name = "${var.db_system_db_home_database_pdb_name}"
		}
		db_version = "${var.db_system_db_home_db_version}"

		#Optional
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
}
```

# oci_database_db_systems

## DbSystem DataSource

Gets a list of db_systems.

### List Operation
Gets a list of the DB Systems in the specified compartment. You can specify a backupId to list only the DB Systems that support creating a database using this backup in this compartment.
    

The following arguments are supported:

* `backup_id` - (Optional) The OCID of the backup. Specify a backupId to list only the DB Systems that support creating a database using this backup in this compartment.
* `compartment_id` - (Required) The compartment [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).


The following attributes are exported:

* `db_systems` - The list of db_systems.

### Example Usage

```hcl
data "oci_database_db_systems" "test_db_systems" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	backup_id = "${oci_database_backup.test_backup.id}"
}
```