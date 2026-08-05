# This example clones an Autonomous Container Database from a backup timestamp.
# It creates a third VM cluster network and a third AVM cluster so this partial
# timestamp clone can run in parallel with the full clone-from-backup example.
#
# Execution order:
# 1. The base example creates the Exadata infrastructure, source VM cluster
#    network, source AVM cluster, source ACD, source ADB, and NFS backup
#    destination.
# 2. This file creates a third VM cluster network on the same Exadata
#    infrastructure.
# 3. This file creates a third AVM cluster using that third VM cluster network.
# 4. This data source lists active local ACD backups for the source ACD after
#    the source ADB exists.
# 5. The clone resource uses the first returned backup end timestamp to create
#    a partial ACD clone on the third AVM cluster.
# 6. The final data sources list cloned ADBs inside the timestamp-cloned ACD.

resource "oci_database_vm_cluster_network" "test_vm_cluster_network3" {
  compartment_id = var.compartment_ocid
  display_name   = "testVmClusterRecommendedNetwork3"
  dns            = ["192.168.10.10"]
  ntp            = ["192.168.10.20"]

  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id

  scans {
    hostname = "myprefix3-nsubz-scan"

    ips = [
      "192.168.19.67",
      "192.168.19.66",
      "192.168.19.68",
    ]

    port                       = 1521
    scan_listener_port_tcp     = 1521
    scan_listener_port_tcp_ssl = 2484
  }

  vm_networks {
    domain_name  = "oracle.com"
    gateway      = "192.169.22.1"
    netmask      = "255.255.0.0"
    network_type = "BACKUP"

    nodes {
      hostname     = "myprefix5-cghdm1"
      ip           = "192.169.19.78"
      db_server_id = data.oci_database_db_servers.test_db_servers.db_servers.0.id
    }

    nodes {
      hostname     = "myprefix5-cghdm2"
      ip           = "192.169.19.80"
      db_server_id = data.oci_database_db_servers.test_db_servers.db_servers.1.id
    }

    vlan_id = "31"
  }

  vm_networks {
    domain_name  = "oracle.com"
    gateway      = "192.168.22.1"
    netmask      = "255.255.0.0"
    network_type = "CLIENT"

    nodes {
      hostname     = "myprefix6-r64zc1"
      ip           = "192.168.19.80"
      vip          = "192.168.19.81"
      vip_hostname = "myprefix6-r64zc1-vip"
      db_server_id = data.oci_database_db_servers.test_db_servers.db_servers.0.id
    }

    nodes {
      hostname     = "myprefix6-r64zc2"
      ip           = "192.168.19.84"
      vip          = "192.168.19.85"
      vip_hostname = "myprefix6-r64zc2-vip"
      db_server_id = data.oci_database_db_servers.test_db_servers.db_servers.1.id
    }

    vlan_id = "50"
  }

  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedvalue"
  }

  freeform_tags = {
    "Department" = "Accounting"
  }

  validate_vm_cluster_network = true
  action                      = "ADD_DBSERVER_NETWORK"

  lifecycle {
    ignore_changes = [
      vm_networks,
    ]
  }
}

resource "oci_database_autonomous_vm_cluster" "autonomous_vm_cluster_3" {
  compartment_id                        = var.compartment_ocid
  display_name                          = "peerAutonomousVmClusterTimestamp"
  exadata_infrastructure_id             = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
  vm_cluster_network_id                 = oci_database_vm_cluster_network.test_vm_cluster_network3.id
  cpu_core_count_per_node               = "40"
  autonomous_data_storage_size_in_tbs   = "5.0"
  memory_per_oracle_compute_unit_in_gbs = "5"
  total_container_databases             = "2"
  is_local_backup_enabled               = "false"
  license_model                         = "LICENSE_INCLUDED"
  time_zone                             = "US/Pacific"

  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"
  }

  freeform_tags = {
    "Department" = "Finance"
  }
}

data "oci_database_autonomous_container_database_backups" "acd_clone_from_backup_timestamp_backups" {
  autonomous_container_database_id = oci_database_autonomous_container_database.autonomous_container_database.id
  infrastructure_type              = "CLOUD_AT_CUSTOMER"
  is_remote                        = false
  state                            = "ACTIVE"

  depends_on = [
    oci_database_autonomous_database.test_autonomous_database
  ]
}

resource "oci_database_autonomous_container_database" "acd_clone_from_backup_timestamp" {
  source                                  = "BACKUP_FROM_TIMESTAMP"
  source_autonomous_container_database_id = oci_database_autonomous_container_database.autonomous_container_database.id
  time_stamp_to_use_for_cloning           = data.oci_database_autonomous_container_database_backups.acd_clone_from_backup_timestamp_backups.autonomous_container_database_backup_collection.0.items.0.time_ended
  autonomous_databases_to_clone           = [oci_database_autonomous_database.test_autonomous_database.display_name]
  autonomous_vm_cluster_id                = oci_database_autonomous_vm_cluster.autonomous_vm_cluster_3.id
  clone_type                              = "PARTIAL"
  compartment_id                          = var.compartment_ocid
  display_name                            = "acdCloneFromBackupTimestamp"
  patch_model                             = "RELEASE_UPDATES"
  service_level_agreement_type            = "STANDARD"

  backup_config {
    backup_destination_details {
      type                                 = "NFS"
      id                                   = oci_database_backup_destination.test_backup_destination.id
      backup_retention_policy_on_terminate = "RETAIN_FOR_72_HOURS"
      is_retention_lock_enabled            = false
    }

    recovery_window_in_days = "10"
  }

  lifecycle {
    ignore_changes = [
      source,
      clone_type,
      source_autonomous_container_database_id,
      time_stamp_to_use_for_cloning,
      autonomous_databases_to_clone
    ]
  }
}

data "oci_database_autonomous_databases" "acd_clone_from_backup_timestamp_databases" {
  autonomous_container_database_id = oci_database_autonomous_container_database.acd_clone_from_backup_timestamp.id
  compartment_id                   = var.compartment_ocid
  display_name                     = oci_database_autonomous_database.test_autonomous_database.display_name

  depends_on = [
    oci_database_autonomous_container_database.acd_clone_from_backup_timestamp
  ]
}

data "oci_database_autonomous_databases" "all_acd_clone_from_backup_timestamp_databases" {
  autonomous_container_database_id = oci_database_autonomous_container_database.acd_clone_from_backup_timestamp.id
  compartment_id                   = var.compartment_ocid

  depends_on = [
    oci_database_autonomous_container_database.acd_clone_from_backup_timestamp
  ]
}
