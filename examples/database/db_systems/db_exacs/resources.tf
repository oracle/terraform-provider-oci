// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
resource "oci_database_cloud_exadata_infrastructure" "test_cloud_exadata_infrastructure" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "MyTFDBExaCs"
  shape               = var.cloud_exadata_infrastructure_shape

  #Optional
  compute_count = var.cloud_exadata_infrastructure_compute_count
  storage_count = var.cloud_exadata_infrastructure_storage_count
}

resource "oci_database_cloud_vm_cluster" "test_cloud_vm_cluster" {
  #Required
  backup_subnet_id                = oci_core_subnet.subnet_backup.id
  cloud_exadata_infrastructure_id = oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure.id
  compartment_id                  = var.compartment_ocid
  cpu_core_count                  = var.cloud_vm_cluster_cpu_core_count
  display_name                    = "MyTFVmClusterExaCs"
  domain                          = oci_core_subnet.subnet.subnet_domain_name
  gi_version                      = var.cloud_vm_cluster_gi_version
  hostname                        = var.cloud_vm_cluster_hostname
  ssh_public_keys                 = [var.ssh_public_key]
  subnet_id                       = oci_core_subnet.subnet.id
}

resource "oci_database_db_home" "test_db_home_vm_cluster" {
  vm_cluster_id = oci_database_cloud_vm_cluster.test_cloud_vm_cluster.id

  database {
    admin_password      = "BEstrO0ng_#11"
    backup_id           = oci_database_backup.test_backup.id
    backup_tde_password = "BEstrO0ng_#11"
    db_name             = "dbVMClus"
    character_set       = "AL32UTF8"
    ncharacter_set      = "AL16UTF16"
    db_workload         = "OLTP"
    pdb_name            = "pdbName"

    freeform_tags = {
      "Department" = "Finance"
    }
  }

  source       = "VM_CLUSTER_BACKUP"
  display_name = "createdDbHome"
}

resource "oci_database_db_system" "test_db_system" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  cpu_core_count      = "22"
  database_edition    = var.db_edition

  db_home {
    database {
      admin_password = var.db_admin_password
      db_name        = "TFdb1Exa"
      character_set  = var.character_set
      ncharacter_set = var.n_character_set
      db_workload    = var.db_workload
      pdb_name       = var.pdb_name

      db_backup_config {
        auto_backup_enabled = false
      }
    }

    db_version   = "12.1.0.2"
    display_name = "MyTFDBHome1Exa"
  }

  maintenance_window_details {
    preference = "CUSTOM_PREFERENCE"

    days_of_week {
      name = "MONDAY"
    }

    hours_of_day       = ["4"]
    lead_time_in_weeks = 2

    months {
      name = "APRIL"
    }

    weeks_of_month = ["2"]
  }

  disk_redundancy  = var.db_disk_redundancy
  shape            = var.db_system_shape
  subnet_id        = oci_core_subnet.subnet.id
  backup_subnet_id = oci_core_subnet.subnet_backup.id
  ssh_public_keys  = [var.ssh_public_key]
  display_name     = var.db_system_display_name
  sparse_diskgroup = var.sparse_diskgroup

  hostname                = var.hostname
  data_storage_percentage = var.data_storage_percentage

  #data_storage_size_in_gb = var.data_storage_size_in_gb
  license_model          = var.license_model
  node_count             = lookup(data.oci_database_db_system_shapes.test_db_system_shapes.db_system_shapes[0], "minimum_node_count")
  backup_network_nsg_ids = [oci_core_network_security_group.test_network_security_group.id]
  nsg_ids                = [oci_core_network_security_group.test_network_security_group_backup.id, oci_core_network_security_group.test_network_security_group.id]

  #To use defined_tags, set the values below to an existing tag namespace, refer to the identity example on how to create tag namespaces
  #defined_tags = map("example-tag-namespace-all.example-tag", "originalValue")

  freeform_tags = {
    "Department" = "Finance"
  }
}

resource "oci_database_backup" "test_backup" {
  depends_on   = ["oci_database_db_system.test_db_system"]
  database_id  = oci_database_db_system.test_db_system.db_home.0.database.0.id
  display_name = "FirstBackup"
}
