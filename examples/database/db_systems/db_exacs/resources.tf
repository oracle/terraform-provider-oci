// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
resource "oci_database_cloud_exadata_infrastructure" "test_cloud_exadata_infrastructure" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "MyTFDBExaCs"
  shape               = var.cloud_exadata_infrastructure_shape

  #Optional
  cluster_placement_group_id = var.cloud_exadata_infrastructure_cluster_placement_group_id
  subscription_id = var.tenant_subscription_id
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
  domain                          = oci_dns_zone.test_zone.name
  gi_version                      = var.cloud_vm_cluster_gi_version
  hostname                        = var.cloud_vm_cluster_hostname
  ssh_public_keys                 = [var.ssh_public_key]
  subnet_id                       = oci_core_subnet.subnet.id

  #Optional
  data_storage_size_in_tbs        = var.cloud_vm_cluster_data_storage_size_in_tbs
  db_node_storage_size_in_gbs     = var.cloud_vm_cluster_db_node_storage_size_in_gbs
  db_servers                      = [data.oci_database_db_servers.test_cloud_db_servers.db_servers.0.id, data.oci_database_db_servers.test_cloud_db_servers.db_servers.1.id]
  memory_size_in_gbs              = var.cloud_vm_cluster_memory_size_in_gbs
  ocpu_count                      = var.cloud_vm_cluster_ocpu_count
  scan_listener_port_tcp          = var.cloud_vm_cluster_scan_listener_port_tcp
  scan_listener_port_tcp_ssl      = var.cloud_vm_cluster_scan_listener_port_tcp_ssl
  private_zone_id                 = oci_dns_zone.test_zone.id
  subscription_id                 = var.tenant_subscription_id

  data_collection_options {
    #Optional
    is_diagnostics_events_enabled = "true"
    is_health_monitoring_enabled = "true"
    is_incident_logs_enabled = "true"
  }

  cloud_automation_update_details{
    is_early_adoption_enabled = "true"
    apply_update_time_preference  {
      apply_update_preferred_start_time = "02:00"
      apply_update_preferred_end_time = "08:00"
    }
  }
}

resource "oci_database_db_home" "test_db_home_vm_cluster" {
  vm_cluster_id = oci_database_cloud_vm_cluster.test_cloud_vm_cluster.id

  database {
    admin_password      = "BEstrO0ng_#11"
    db_name             = "dbVMClus"
    character_set       = "AL32UTF8"
    ncharacter_set      = "AL16UTF16"
    db_workload         = "OLTP"
    pdb_name            = "pdbName"
    db_backup_config {
      auto_backup_enabled = "true"
      auto_backup_window     = "SLOT_TWO"
      auto_full_backup_day    = "SUNDAY"
      auto_full_backup_window = "SLOT_ONE"
      recovery_window_in_days = 10
      run_immediate_full_backup = false
    }
    freeform_tags = {
      "Department" = "Finance"
    }
  }

  # VM_CLUSTER_BACKUP can also be specified as a source for cloud VM clusters.
  source       = "VM_CLUSTER_NEW"
  db_version   = "19.0.0.0"
  display_name = "createdDbHome"
  is_unified_auditing_enabled = "true"
}

resource "oci_database_db_home" "test_dbrs_db_home_vm_cluster" {
  vm_cluster_id = oci_database_cloud_vm_cluster.test_cloud_vm_cluster.id

  database {
    admin_password = "BEstrO0ng_#11"
    db_name        = "dbrsDb"
    character_set  = "AL32UTF8"
    ncharacter_set = "AL16UTF16"
    db_workload    = "OLTP"
    pdb_name       = "pdbName"
    db_backup_config {
      auto_backup_enabled    = "true"
      auto_backup_window     = "SLOT_TWO"
      backup_deletion_policy = "DELETE_IMMEDIATELY"
      backup_destination_details {
        dbrs_policy_id = "dbrsPolicyOcid"
        type           = "DBRS"
      }
    }
    freeform_tags = {
      "Department" = "Finance"
    }
  }

  # VM_CLUSTER_BACKUP can also be specified as a source for cloud VM clusters.
  source       = "VM_CLUSTER_NEW"
  db_version   = "19.0.0.0"
  display_name = "createdDbrsDbHome"
}


resource "oci_database_db_home" "test_db_home_vm_cluster_no_db" {
  vm_cluster_id = oci_database_cloud_vm_cluster.test_cloud_vm_cluster.id

  # VM_CLUSTER_BACKUP can also be specified as a source for cloud VM clusters.
  source       = "VM_CLUSTER_NEW"
  db_version   = "19.0.0.0"
  display_name = "createdDbHomeNoDb"
}

resource "oci_database_backup" "test_backup" {
  depends_on   = [oci_database_db_home.test_db_home_vm_cluster]
  database_id  = oci_database_db_home.test_db_home_vm_cluster.database.0.id
  display_name = "FirstBackup"
}


resource "oci_database_cloud_vm_cluster_iorm_config" "test_cloud_vm_cluster_iorm_config" {
  cloud_vm_cluster_id = oci_database_cloud_vm_cluster.test_cloud_vm_cluster.id
  objective    = "AUTO"

  db_plans {
    db_name = "default"
    share   = 1
  }
}

resource "oci_database_application_vip" "test_application_vip" {
    #Required
    cloud_vm_cluster_id = oci_database_cloud_vm_cluster.test_cloud_vm_cluster.id
    hostname_label      = "hostnameLabel"
    subnet_id           = oci_core_subnet.subnet.id

    #Optional
    db_node_id          = data.oci_database_db_nodes.db_nodes.db_nodes[0]["id"]
}

resource "oci_database_database" "test_database" {
  #Required
  database {
    admin_password = "BEstrO0ng_#11"
    db_name        = "TFdb2Exa"
    character_set  = "AL32UTF8"
    ncharacter_set = "AL16UTF16"
    db_workload    = "OLTP"

    db_backup_config {
      auto_backup_enabled = false
    }
  }

  db_home_id = oci_database_db_home.test_db_home_vm_cluster_no_db.id
  source     = "NONE"
}

resource "oci_database_pluggable_database" "test_pluggable_database" {
        container_database_id = oci_database_database.test_database.id
        pdb_admin_password = "BEstrO0ng_#11"
        pdb_name = "SalesPdb"
        tde_wallet_password = "BEstrO0ng_#11"
}