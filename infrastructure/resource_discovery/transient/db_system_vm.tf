// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_database_db_system_shapes" "test_db_system_shapes" {
  availability_domain = "${data.oci_identity_availability_domain.ad1.name}"
  compartment_id      = "${var.compartment_ocid}"

  filter {
    name   = "shape"
    values = ["${var.db_system_shape}"]
  }
}

resource "oci_database_db_system" "test_db_system" {
  availability_domain = "${data.oci_identity_availability_domain.ad1.name}"
  compartment_id      = "${var.compartment_ocid}"
  database_edition    = "${var.db_edition}"

  db_home {
    database {
      admin_password = "${var.db_admin_password}"
      db_name        = "aTFdbVm"
      character_set  = "${var.character_set}"
      ncharacter_set = "${var.n_character_set}"
      db_workload    = "${var.db_workload}"
      pdb_name       = "${var.pdb_name}"

      db_backup_config {
        auto_backup_enabled = false
      }
    }

    db_version   = "${var.db_version}"
    display_name = "MyTFDBHomeVm"
  }

  db_system_options {
    storage_management = "LVM"
  }

  disk_redundancy         = "${var.db_disk_redundancy}"
  shape                   = "${var.db_system_shape}"
  subnet_id               = "${oci_core_subnet.subnet1_lb_rd.id}"
  ssh_public_keys         = ["${var.ssh_public_key}"]
  display_name            = "MyTFDBSystemVM"
  hostname                = "${var.hostname}"
  data_storage_size_in_gb = "${var.data_storage_size_in_gb}"
  license_model           = "${var.license_model}"
  node_count              = "${lookup(data.oci_database_db_system_shapes.test_db_system_shapes.db_system_shapes[0], "minimum_node_count")}"
  nsg_ids                 = ["${oci_core_network_security_group.network_security_group_backup_rd.id}", "${oci_core_network_security_group.network_security_group_rd.id}"]

  #To use defined_tags, set the values below to an existing tag namespace, refer to the identity example on how to create tag namespaces
  #defined_tags = "${map("example-tag-namespace-all.example-tag", "originalValue")}"

  freeform_tags = {
    "Department" = "Finance"
  }
}
