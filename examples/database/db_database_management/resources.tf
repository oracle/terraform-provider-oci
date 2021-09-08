// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
resource "oci_database_db_system" "test_db_system" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  database_edition    = var.db_edition

  db_home {
    database {
      admin_password = var.db_admin_password
      db_name        = "aTFdbVm"
      character_set  = var.character_set
      ncharacter_set = var.n_character_set
      db_workload    = var.db_workload
      pdb_name       = var.pdb_name

      db_backup_config {
        auto_backup_enabled = false
      }
    }

    db_version   = var.db_version
    display_name = "MyTFDBHomeVm"
  }

  db_system_options {
    storage_management = "LVM"
  }

  disk_redundancy         = var.db_disk_redundancy
  shape                   = var.db_system_shape
  subnet_id               = "ocid1.subnet.oc1.ap-hyderabad-1.aaaaaaaaxp3p7plootrgchtk4s6olk7lmnjhp5xfxqwyh74jmr7dw2fxgjnq"
  ssh_public_keys         = [var.ssh_public_key]
  display_name            = "MyTFDBSystemVM"
  hostname                = var.hostname
  data_storage_size_in_gb = var.data_storage_size_in_gb
  license_model           = var.license_model
  node_count              = data.oci_database_db_system_shapes.test_db_system_shapes.db_system_shapes[0]["minimum_node_count"]
//  nsg_ids                 = [oci_core_network_security_group.test_network_security_group_backup.id, oci_core_network_security_group.test_network_security_group.id]

  #To use defined_tags, set the values below to an existing tag namespace, refer to the identity example on how to create tag namespaces
  #defined_tags  = {"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"}

  freeform_tags = {
    "Department" = "Finance"
  }
  //To ignore DbVersion after database upgrade
  lifecycle {
    ignore_changes = [
      db_home.0.db_version,
    ]
  }
}

resource "oci_database_cloud_database_management" "test" {
  database_id           = data.oci_database_databases.databases.databases.0.id
  management_type       = "BASIC"
  private_end_point_id  = "ocid1.dbmgmtprivateendpoint.oc1.ap-hyderabad-1.amaaaaaacsc5xjaamlmllhfxmxict6jf3irizwsydralyklninmwsrovggkq"
  service_name          = "DB0809_hyd17q.sub02231620340.dbmgmtcustomer.oraclevcn.com"
  credentialdetails {
    user_name           = "dbsnmp"
    password_secret_id  = "ocid1.vaultsecret.oc1.ap-hyderabad-1.amaaaaaacsc5xjaa2q7r6kfzdm44ylxqwomht6uinb5zyhezka7sl2t62ecq"
  }
  enable_management     = "true"
}