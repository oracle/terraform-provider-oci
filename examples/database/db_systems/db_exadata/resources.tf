// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

resource "oci_database_db_system" "test_db_system" {
  availability_domain = "${data.oci_identity_availability_domain.ad.name}"
  compartment_id      = "${var.compartment_ocid}"
  cpu_core_count      = "${var.cpu_core_count}"
  database_edition    = "${var.db_edition}"
  time_zone           = "${var.time_zone}"

  db_home {
    database {
      admin_password = "${var.db_admin_password}"
      db_name        = "TFdb1Exa"
      character_set  = "${var.character_set}"
      ncharacter_set = "${var.n_character_set}"
      db_workload    = "${var.db_workload}"
      pdb_name       = "${var.pdb_name}"

      db_backup_config {
        auto_backup_enabled = false
      }
    }

    db_version   = "${var.db_version}"
    display_name = "MyTFDBHome1Exa"
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
  license_model          = "${var.license_model}"
  node_count             = "${lookup(data.oci_database_db_system_shapes.test_db_system_shapes.db_system_shapes[0], "minimum_node_count")}"
  backup_network_nsg_ids = ["${oci_core_network_security_group.test_network_security_group.id}"]
  nsg_ids                = ["${oci_core_network_security_group.test_network_security_group_backup.id}", "${oci_core_network_security_group.test_network_security_group.id}"]

  #To use defined_tags, set the values below to an existing tag namespace, refer to the identity example on how to create tag namespaces
  #defined_tags = "${map("example-tag-namespace-all.example-tag", "originalValue")}"

  freeform_tags = {
    "Department" = "Finance"
  }
}

resource "oci_database_exadata_iorm_config" "test_exadata_iorm_config" {
  db_system_id = "${oci_database_db_system.test_db_system.id}"
  objective    = "AUTO"

  db_plans {
    db_name = "default"
    share   = 1
  }
}

resource "oci_database_database" "test_database" {
  #Required
  database {
    admin_password = "${var.db_admin_password}"
    db_name        = "TFdb2Exa"
    character_set  = "${var.character_set}"
    ncharacter_set = "${var.n_character_set}"
    db_workload    = "${var.db_workload}"

    db_backup_config {
      auto_backup_enabled = false
    }
  }

  db_home_id = "${data.oci_database_db_homes.db_homes.db_homes.0.db_home_id}"
  source     = "NONE"
}
