# Gets a list of Availability Domains
data "oci_identity_availability_domains" "ADs" {
  compartment_id = "${var.tenancy_ocid}"
}

# Get DB node list
data "oci_database_db_nodes" "db_nodes" {
  compartment_id = "${var.compartment_ocid}"
  db_system_id   = "${oci_database_db_system.test_db_system.id}"
}

# Get DB node details
data "oci_database_db_node" "db_node_details" {
  db_node_id = "${lookup(data.oci_database_db_nodes.db_nodes.db_nodes[0], "id")}"
}

# Gets the OCID of the first (default) vNIC
#data "oci_core_vnic" "db_node_vnic" {
#    vnic_id = "${data.oci_database_db_node.db_node_details.vnic_id}"
#}

data "oci_database_db_homes" "db_homes" {
  compartment_id = "${var.compartment_ocid}"
  db_system_id   = "${oci_database_db_system.test_db_system.id}"
}

data "oci_database_databases" "databases" {
  compartment_id = "${var.compartment_ocid}"
  db_home_id     = "${data.oci_database_db_homes.db_homes.db_homes.0.db_home_id}"
}

data "oci_database_backups" "test_backups" {
  database_id = "${data.oci_database_databases.databases.databases.0.id}"

  filter {
    name   = "id"
    values = ["${oci_database_backup.test_backup.id}"]
  }

  filter {
    name   = "state"
    values = ["AVAILABLE"]
  }
}

data "oci_database_db_system_patches" "patches" {
  db_system_id = "${oci_database_db_system.test_db_system.id}"
}

data "oci_database_db_system_patch_history_entries" "patches_history" {
  db_system_id = "${oci_database_db_system.test_db_system.id}"
}

data "oci_database_db_home_patches" "patches" {
  db_home_id = "${data.oci_database_db_homes.db_homes.db_homes.0.db_home_id}"
}

data "oci_database_db_home_patch_history_entries" "patches_history" {
  db_home_id = "${data.oci_database_db_homes.db_homes.db_homes.0.db_home_id}"
}

data "oci_database_db_systems" "db_systems" {
  compartment_id = "${var.compartment_ocid}"
  backup_id      = "${oci_database_backup.test_backup.id}"
}

data "oci_database_db_versions" "test_db_versions_by_db_system_id" {
  compartment_id = "${var.compartment_ocid}"
  db_system_id   = "${oci_database_db_system.test_db_system.id}"
}

data "oci_database_db_system_shapes" "test_db_system_shapes" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain - 1],"name")}"
  compartment_id      = "${var.compartment_ocid}"

  filter {
    name   = "shape"
    values = ["${var.db_system_shape}"]
  }
}
