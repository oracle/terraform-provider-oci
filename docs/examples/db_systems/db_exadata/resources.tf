resource "oci_database_db_system" "test_db_system" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain - 1],"name")}"
  compartment_id      = "${var.compartment_ocid}"
  cpu_core_count      = "${var.cpu_core_count}"
  database_edition    = "${var.db_edition}"

  db_home {
    database {
      admin_password = "${var.db_admin_password}"
      db_name        = "${var.db_name}"
      character_set  = "${var.character_set}"
      ncharacter_set = "${var.n_character_set}"
      db_workload    = "${var.db_workload}"
      pdb_name       = "${var.pdb_name}"

      db_backup_config {
        auto_backup_enabled = false
      }
    }

    db_version   = "${var.db_version}"
    display_name = "${var.db_home_display_name}"
  }

  disk_redundancy  = "${var.db_disk_redundancy}"
  shape            = "${var.db_system_shape}"
  subnet_id        = "${oci_core_subnet.subnet.id}"
  backup_subnet_id = "${oci_core_subnet.subnet_backup.id}"
  ssh_public_keys  = ["${var.ssh_public_key}"]
  display_name     = "${var.db_system_display_name}"
  sparse_diskgroup = "${var.sparse_diskgroup}"

  hostname                = "${var.hostname}"
  data_storage_percentage = "${var.data_storage_percentage}"

  #data_storage_size_in_gb = "${var.data_storage_size_in_gb}"
  license_model = "${var.license_model}"
  node_count    = "${lookup(data.oci_database_db_system_shapes.test_db_system_shapes.db_system_shapes[0], "minimum_node_count")}"

  #To use defined_tags, set the values below to an existing tag namespace, refer to the identity example on how to create tag namespaces
  #defined_tags = "${map("example-tag-namespace-all.example-tag", "originalValue")}"

  freeform_tags = {
    "Department" = "Finance"
  }
}
