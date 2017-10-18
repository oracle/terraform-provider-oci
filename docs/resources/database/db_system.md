# oci\_database\_db\_system

Provides an DBSystem resource.

## Example Usage

```
resource "oci_database_db_system" "TFDBNode" {
  availability_domain = "${var.AvailabilityDomain}"
  compartment_id = "${var.CompartmentOCID}"
  cpu_core_count = "${var.CPUCoreCount}"
  database_edition = "${var.DBEdition}"
  db_home {
    database {
      "admin_password" = "${var.DBAdminPassword}"
      "db_name" = "${var.DBName}"
      "character_set" = "${var.CharacterSet}"
      "ncharacter_set" = "${var.NCharacterSet}"
      "db_workload" = "${var.DBWorkload}"
      "pdb_name" = "${var.PDBName}"
    }
    db_version = "${var.DBVersion}"
    display_name = "${var.DBDisplayName}"
  }
  disk_redundancy = "${var.DBDiskRedundancy}"
  shape = "${var.DBNodeShape}"
  subnet_id = "${var.SubnetOCID}"
  backup_subnet_id = "${var.BackupSubnetOCID}"
  ssh_public_keys = ["${var.SSHPublicKey}"]
  cluster_name = "${var.ClusterName}"
  display_name = "${var.DBNodeDisplayName}"
  domain = "${var.DBNodeDomainName}"
  hostname = "${var.DBNodeHostName}"
  data_storage_percentage = "${var.DataStoragePercentage}"
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The name of the Availability Domain that the DB System is located in.
* `backup_subnet_id` - (Optional) The OCID of the backup network subnet the DB System is associated with. Applicable only to Exadata.
* `cluster_name` - (Optional) cluster name is used for Exadata DBSystems
* `compartment_id` - (Required) The OCID of the compartment.
* `cpu_core_count` - (Required) The number of CPU cores enabled on the DB System.
* `data_storage_percentage` - (Optional) The percentage assigned to DATA storage (user data and database files).
* `database_edition` - (Required) The Oracle Database Edition that applies to all the databases on the DB System.
* `db_home` - (Required) Create DBHome details. See [Create DBHome Details](#create-dbhome-details) below for detials.
* `disk_redundancy` - (Optional) The type of redundancy configured for the DB System.
* `display_name` - (Optional) The user-friendly name for the DB System. It does not have to be unique.
* `domain` - (Optional) A domain name to assign to the DB System.
* `hostname` - (Required) The host name to assign to the DB Node.
* `shape` - (Required) The shape of the DB System.
* `ssh_public_keys` - (Required) The public key portion of the key pair to use for SSH access to the DB System.
* `subnet_id` - (Required) The OCID of the subnet the DB System is associated with.

## Create DBHome Details

The following arguments are supported:

* `character_set` - (Optional) The character set for the database.
* `database` - (Required) Create Database details. See [Create Database Details](#create-database-details) below for details.
* `db_version` - (Required) A valid Oracle database version.
* `db_workload` - (Optional) Database workload type.
* `display_name` - (Optional) The user-provided name of the database home.
* `ncharacter_set` - (Optional) National character set for the database.
* `pdb_name` - (Optional) Pluggable database name. It must begin with an alphabetic character and can contain a maximum of eight alphanumeric characters. Special characters are not permitted. Pluggable database should not be same as database name.

## Create Database Details

The following arguments are supported:

* `admin_password` - (Required) A strong password for SYS, SYSTEM, and PDB Admin.
* `db_name` - (Required) The database name (alphanumeric only).

## Attributes Reference

The following attributes are exported:

* `id` - The OCID of the DB System.
* `lifecycle_details` - Additional information about the current lifecycleState.
* `listener_port` - The port number configured for the listener on the DB System.
* `scan_dns_record_id` - The OCID of the DNS record for the SCAN IP addresses that are associated with the DB System.
* `scan_ip_ids` - The OCID of the Single Client Access Name (SCAN) IP addresses associated with the DB System. SCAN IP addresses are typically used for load balancing and are not assigned to any interface.
* `state` - The current state of the DB System.
* `time_created` - The date and time the DB System was created.
* `version` - The version of the DB System.
* `vip_ids` - The OCID of the virtual IP (VIP) addresses associated with the DB System. The Cluster Ready Services (CRS) creates and maintains one VIP address for each node in the DB System to enable failover. If one node fails, the VIP is reassigned to another active node in the cluster.
