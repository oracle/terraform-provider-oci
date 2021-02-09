// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
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

  # VM_CLUSTER_NONE can also be specified as a source for cloud VM clusters.
  source       = "VM_CLUSTER_BACKUP"
  display_name = "createdDbHome"
}

resource "oci_database_backup" "test_backup" {
  depends_on   = ["oci_database_db_home.test_db_home_vm_cluster"]
  database_id  = oci_database_db_home.test_db_home_vm_cluster.database.id
  display_name = "FirstBackup"
}
