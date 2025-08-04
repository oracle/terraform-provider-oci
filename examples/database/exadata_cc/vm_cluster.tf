// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.2019 Oracle and/or its affiliates. All rights reserved.

resource "oci_database_vm_cluster_network" "test_vm_cluster_network" {
  compartment_id = var.compartment_ocid
  display_name   = "testVmClusterRecommendedNetwork"
  dns            = ["192.168.10.10"]
  ntp            = ["192.168.10.20"]

  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id

  scans {
    hostname = "myprefix1-nsubz-scan"

    ips = [
      "192.168.19.7",
      "192.168.19.6",
      "192.168.19.8",
    ]

    port = 1521
    scan_listener_port_tcp = 1521
    scan_listener_port_tcp_ssl = 2484
  }

  vm_networks {
    domain_name  = "oracle.com"
    gateway      = "192.169.20.1"
    netmask      = "255.255.0.0"
    network_type = "BACKUP"

    nodes {
      hostname = "myprefix2-cghdm1"
      ip       = "192.169.19.18"
      db_server_id = data.oci_database_db_servers.test_db_servers.db_servers.0.id
    }

    nodes {
      hostname = "myprefix2-cghdm2"
      ip       = "192.169.19.20"
      db_server_id = data.oci_database_db_servers.test_db_servers.db_servers.1.id
    }

    vlan_id = "11"
  }

  vm_networks {
    domain_name  = "oracle.com"
    gateway      = "192.168.20.1"
    netmask      = "255.255.0.0"
    network_type = "CLIENT"

    nodes {
      hostname     = "myprefix1-r64zc1"
      ip           = "192.168.19.10"
      vip          = "192.168.19.11"
      vip_hostname = "myprefix1-r64zc1-vip"
      db_server_id = data.oci_database_db_servers.test_db_servers.db_servers.0.id
    }

    nodes {
      hostname     = "myprefix1-r64zc2"
      ip           = "192.168.19.14"
      vip          = "192.168.19.15"
      vip_hostname = "myprefix1-r64zc2-vip"
      db_server_id = data.oci_database_db_servers.test_db_servers.db_servers.1.id
    }

    vlan_id = "10"
  }

  #Optional
  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedvalue"
  }

  freeform_tags = {
    "Department" = "Accounting"
  }

  validate_vm_cluster_network = true

  action = "ADD_DBSERVER_NETWORK"

  lifecycle {
    ignore_changes = [
      vm_networks,
    ]
  }
}

resource "oci_database_vm_cluster_network" "test_vm_cluster_network2" {
  compartment_id = var.compartment_ocid
  display_name   = "testVmClusterRecommendedNetwork2"
  dns            = ["192.178.10.10"]
  ntp            = ["192.178.10.20"]

  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id

  scans {
    hostname = "myprefix3-nsubz-scan"

    ips = [
      "192.178.19.7",
      "192.178.19.6",
      "192.178.19.8",
    ]

    port = 1521
    scan_listener_port_tcp = 1521
    scan_listener_port_tcp_ssl = 2484
  }

  vm_networks {
    domain_name  = "oracle.com"
    gateway      = "192.179.20.1"
    netmask      = "255.255.0.0"
    network_type = "BACKUP"

    nodes {
      hostname = "myprefix4-cghdm1"
      ip       = "192.179.19.18"
      db_server_id = data.oci_database_db_servers.test_db_servers.db_servers.0.id
    }

    nodes {
      hostname = "myprefix4-cghdm2"
      ip       = "192.179.19.20"
      db_server_id = data.oci_database_db_servers.test_db_servers.db_servers.1.id
    }

    vlan_id = "31"
  }

  vm_networks {
    domain_name  = "oracle.com"
    gateway      = "192.178.20.1"
    netmask      = "255.255.0.0"
    network_type = "CLIENT"

    nodes {
      hostname     = "myprefix5-r64zc1"
      ip           = "192.178.19.10"
      vip          = "192.178.19.11"
      vip_hostname = "myprefix5-r64zc1-vip"
      db_server_id = data.oci_database_db_servers.test_db_servers.db_servers.0.id
    }

    nodes {
      hostname     = "myprefix5-r64zc2"
      ip           = "192.178.19.14"
      vip          = "192.178.19.15"
      vip_hostname = "myprefix5-r64zc2-vip"
      db_server_id = data.oci_database_db_servers.test_db_servers.db_servers.1.id
    }

    vlan_id = "41"
  }

  #Optional
  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedvalue"
  }

  freeform_tags = {
    "Department" = "Accounting"
  }

  validate_vm_cluster_network = true

  action = "ADD_DBSERVER_NETWORK"

  lifecycle {
    ignore_changes = [
      vm_networks,
    ]
  }
}

data "oci_database_gi_versions" "gi_version" {
  compartment_id = var.compartment_ocid
  shape = "ExadataCC.Quarter3.100"
}

resource "oci_database_vm_cluster" "test_vm_cluster" {
  #Required
  compartment_id            = var.compartment_ocid
  cpu_core_count            = "16"
  display_name              = "testVmCluster"
  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
  gi_version                = data.oci_database_gi_versions.gi_version.gi_versions.0.version
  ssh_public_keys           = [var.ssh_public_key]
  vm_cluster_network_id     = oci_database_vm_cluster_network.test_vm_cluster_network.id
  db_servers                = [data.oci_database_db_servers.test_db_servers.db_servers.0.id, data.oci_database_db_servers.test_db_servers.db_servers.1.id]
  #Optional
  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedvalue"
  }

  freeform_tags = {
    "Department" = "Accounting"
  }

  is_local_backup_enabled     = "false"
  is_sparse_diskgroup_enabled = "false"
  license_model               = "LICENSE_INCLUDED"
  data_storage_size_in_tbs    = "84"
  db_node_storage_size_in_gbs = "120"
  memory_size_in_gbs          = "60"
  vm_cluster_type             = "REGULAR"
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

data "oci_database_db_servers" "test_db_servers" {
  #Required
  compartment_id            = var.compartment_ocid
  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
}

data "oci_database_vm_cluster_recommended_network" "test_vm_cluster_recommended_network" {
  #Required
  compartment_id            = var.compartment_ocid
  display_name              = "testVmClusterRecommendedNetwork"
  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id

  networks {
    #Required
    cidr         = "192.168.19.2/16"
    domain       = "oracle.com"
    gateway      = "192.168.20.1"
    netmask      = "255.255.0.0"
    network_type = "CLIENT"
    prefix       = "myprefix1"
    vlan_id      = "10"
  }

  networks {
    #Required
    cidr         = "192.169.19.1/16"
    domain       = "oracle.com"
    gateway      = "192.169.20.1"
    netmask      = "255.255.0.0"
    network_type = "BACKUP"
    prefix       = "myprefix2"
    vlan_id      = "11"
  }

  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedvalue"
  }

  freeform_tags = {
    "Department" = "Accounting"
  }
}

resource "oci_database_db_home" "test_db_home_vm_cluster" {
  vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id
  source       = "VM_CLUSTER_NEW"
  db_version   = "19.0.0.0"
  display_name = "createdDbHome"
}

resource "oci_database_database" "test_exacc_database"{
  database {
    admin_password = "BEstrO0ng_#11"
    db_name        = "dbVMClus"
    character_set  = "AL32UTF8"
    ncharacter_set = "AL16UTF16"
    db_workload    = "OLTP"
    pdb_name       = "pdbName"

    freeform_tags = {
      "Department" = "Finance"
    }

    db_backup_config {
      auto_backup_enabled = true
      auto_backup_window  = "SLOT_TWO"

      backup_destination_details {
        id   = oci_database_backup_destination.test_backup_destination_nfs.id
        type = "NFS"
      }
    }
    encryption_key_location_details {
        #Required
        hsm_password  = "hsmPassword"
        provider_type = "EXTERNAL"
    }
  }
  db_home_id = oci_database_db_home.test_db_home_vm_cluster.id
  source     = "NONE"
}

resource "oci_database_backup_destination" "test_backup_destination_nfs" {
  #Required
  compartment_id = var.compartment_ocid
  display_name   = "testBackupDestination"
  type           = "NFS"

  #Optional

  freeform_tags = {
    "Department" = "Finance"
  }
  mount_type_details {
    local_mount_point_path = "localMountPointPath"
    mount_type             = "SELF_MOUNT"
  }
}

data "oci_database_vm_cluster_network_download_config_file" "test_vm_cluster_network_download_config_file" {
  #Required
  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
  vm_cluster_network_id     = oci_database_vm_cluster_network.test_vm_cluster_network.id

  #Optional
  base64_encode_content = true
}

data "oci_database_vm_cluster_networks" "test_vm_cluster_networks" {
  #Required
  compartment_id            = var.compartment_ocid
  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
}

data "oci_database_vm_clusters" "test_vm_clusters" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
}

# resource "local_file" "test_vm_cluster_network_downloaded_config_file" {
#   content  = data.oci_database_vm_cluster_network_download_config_file.test_vm_cluster_network_download_config_file.content
#   filename = "${path.module}/vm_cluster_config.txt"
# }


resource "oci_database_exadata_infrastructure_configure_exascale_management" "test_exadata_infrastructure_configure_exascale_management" {
  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
  total_storage_in_gbs            = var.exadata_infrastructure_configure_exascale_management_total_storage_in_gbs
}

resource "oci_database_exascale_db_storage_vault" "test_exascale_db_storage_exacc_vault" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "ExampleExascaleDbStorageVault"
  high_capacity_database_storage {
    total_size_in_gbs = 2048
  }
  exadata_infrastructure_id = oci_database_exadata_infrastructure_configure_exascale_management.test_exadata_infrastructure_configure_exascale_management.id
}

resource "oci_database_vm_cluster_network" "test_vm_cluster_network3" {
  compartment_id = var.compartment_ocid
  display_name   = "testVmClusterRecommendedNetwork3"
  dns            = ["192.178.10.10"]
  ntp            = ["192.178.10.20"]

  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id

  scans {
    hostname = "myprefix3-nsubzabc-scan"

    ips = [
      "192.178.20.7",
      "192.178.20.6",
      "192.178.20.8",
    ]

    port = 1521
    scan_listener_port_tcp = 1521
    scan_listener_port_tcp_ssl = 2484
  }

  vm_networks {
    domain_name  = "oracle.com"
    gateway      = "192.179.21.1"
    netmask      = "255.255.0.0"
    network_type = "BACKUP"

    nodes {
      hostname = "myprefix8-cghdm1"
      ip       = "192.179.20.18"
      db_server_id = data.oci_database_db_servers.test_db_servers.db_servers.0.id
    }

    nodes {
      hostname = "myprefix8-cghdm2"
      ip       = "192.179.20.20"
      db_server_id = data.oci_database_db_servers.test_db_servers.db_servers.1.id
    }

    vlan_id = "31"
  }

  vm_networks {
    domain_name  = "oracle.com"
    gateway      = "192.178.21.1"
    netmask      = "255.255.0.0"
    network_type = "CLIENT"

    nodes {
      hostname     = "myprefix9-r64zc1"
      ip           = "192.178.21.10"
      vip          = "192.178.21.11"
      vip_hostname = "myprefix8-r64zc1-vip"
      db_server_id = data.oci_database_db_servers.test_db_servers.db_servers.0.id
    }

    nodes {
      hostname     = "myprefix8-r64zc2"
      ip           = "192.178.21.14"
      vip          = "192.178.21.15"
      vip_hostname = "myprefix8-r64zc2-vip"
      db_server_id = data.oci_database_db_servers.test_db_servers.db_servers.1.id
    }

    vlan_id = "45"
  }

  #Optional
  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedvalue"
  }

  freeform_tags = {
    "Department" = "Accounting"
  }

  validate_vm_cluster_network = true

  action = "ADD_DBSERVER_NETWORK"

  lifecycle {
    ignore_changes = [
      vm_networks,
    ]
  }
}

resource "oci_database_vm_cluster" "test_exascale_vm_cluster" {
  #Required
  compartment_id            = var.compartment_ocid
  cpu_core_count            = "16"
  display_name              = "testVmCluster2"
  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
  gi_version                = "23.0.0.0.0"
  ssh_public_keys           = [var.ssh_public_key]
  vm_cluster_network_id     = oci_database_vm_cluster_network.test_vm_cluster_network3.id
  db_servers                = [data.oci_database_db_servers.test_db_servers.db_servers.0.id, data.oci_database_db_servers.test_db_servers.db_servers.1.id]
  exascale_db_storage_vault_id    = oci_database_exascale_db_storage_vault.test_exascale_db_storage_exacc_vault.id
  #Optional
  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedvalue"
  }

  #Optional
  is_local_backup_enabled     = "false"
  is_sparse_diskgroup_enabled = "false"
  license_model               = "LICENSE_INCLUDED"
  data_storage_size_in_tbs    = "26"
  db_node_storage_size_in_gbs = "120"
  memory_size_in_gbs          = "60"
}

variable "exadata_infrastructure_configure_exascale_management_total_storage_in_gbs" {
  default = 4096
}

data "oci_database_backup_destinations" "test_database_backup_destinations" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  type = "NFS"
}

data "oci_database_backup_destination" "test_database_backup_destination" {
  #Required
  backup_destination_id = oci_database_backup_destination.test_backup_destination_nfs.id
}