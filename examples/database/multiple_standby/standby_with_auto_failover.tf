resource "oci_database_cloud_exadata_infrastructure" "test_cloud_exadata_infrastructure_second_standby" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_id
  display_name        = "TstExaInfra3"
  shape               = "Exadata.X8M"

  #Optional
  compute_count = 2
  storage_count = 3
}

resource "oci_database_cloud_vm_cluster" "test_cloud_vm_cluster3" {
  #Required
  backup_subnet_id                = oci_core_subnet.test_subnet_backup4.id
  cloud_exadata_infrastructure_id = oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure_second_standby.id
  compartment_id                  = var.compartment_id
  cpu_core_count                  = "22"
  display_name                    = "StandbyTwoCluster"
  gi_version                      = "19.0.0.0"
  hostname                        = "myOracleDB"
  ssh_public_keys                 = [var.ssh_public_key]
  subnet_id                       = oci_core_subnet.t3.id

  #Optional
  scan_listener_port_tcp          = "1521"
  scan_listener_port_tcp_ssl      = "2484"

  data_collection_options {
    #Optional
    is_diagnostics_events_enabled = "true"
    is_health_monitoring_enabled = "true"
    is_incident_logs_enabled = "true"
  }
}

resource "oci_database_db_home" "test_db_home3" {
  db_system_id = oci_database_cloud_vm_cluster.test_cloud_vm_cluster3.id
  db_version   = "19.27.0.0"
  source = "NONE"
  display_name = "createdDbHomeNone"
}

resource "oci_database_database" "second_standby_database" {
  #Required
  database {
    admin_password = "BEstrO0ng_#11"
    database_admin_password  = "BEstrO0ng_#11" #required for add standby
    character_set  = "AL32UTF8"
    ncharacter_set = "AL16UTF16"
    db_workload    = "OLTP"
    source_database_id = oci_database_database.primary_database.id
    protection_mode = "MAXIMUM_PERFORMANCE"
    transport_type = "ASYNC"
    source_tde_wallet_password = "BEstrO0ng_#11"
    defined_tags     = map("example-tag-namespace-all.example-tag", "databaseDefinedTags1")
    freeform_tags    = {"databaseFreeformTagsK" = "databaseFreeformTagsV"}
    db_unique_name = "StbDb2"

    auto_failover_configuration {
      managed_auto_failover = "ENABLE"
      failover_targets = [oci_database_database.primary_database.db_unique_name, oci_database_database.standby_database.db_unique_name]
    }
  }

  db_home_id = oci_database_db_home.test_db_home3.id
  source     = "DATAGUARD"
}