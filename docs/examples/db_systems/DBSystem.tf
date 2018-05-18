resource "oci_database_db_system" "TFDBNode" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
  compartment_id = "${var.compartment_ocid}"
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
  ssh_public_keys = ["${var.ssh_public_key}"]
  display_name = "${var.DBNodeDisplayName}"

  # Set this to specify the domain name for this DB System unless the Oracle-provided Internet and
  # VCN Resolver is enabled for the specified subnet above.
  #domain = "${var.DBNodeDomainName}"
  hostname = "${var.DBNodeHostName}"
  data_storage_percentage = "40"
  data_storage_size_in_gb = "${var.DataStorageSizeInGB}"
  license_model = "${var.LicenseModel}"
  node_count = "${var.NodeCount}"
}
