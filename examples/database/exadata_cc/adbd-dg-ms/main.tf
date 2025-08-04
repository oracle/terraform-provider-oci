provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

#### Begin Resources ####
resource "oci_identity_tag_namespace" "tag-namespace1" {
  #Required
  compartment_id = var.tenancy_ocid
  description    = "tagNamespace1"
  name           = "testexamples-tag-namespace1"
}

resource "oci_identity_tag" "tag1" {
  #Required
  description      = "tf example tag"
  name             = "tf-example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag-namespace1.id
}

resource "local_file" "activation_file" {
  filename = "/tmp/activation.zip"
  content  = ""
}

resource "random_string" "db_unique_name" {
  length      = 8
  special     = false
  min_numeric = 0
  min_special = 0
}

resource "oci_database_exadata_infrastructure" "primary_exadata_infrastructure" {
  activation_file             = local_file.activation_file.filename
  admin_network_cidr          = "192.168.0.0/16"
  cloud_control_plane_server1 = "10.32.88.1"
  cloud_control_plane_server2 = "10.32.88.3"
  compartment_id              = var.compartment_ocid
  storage_count               = 3
  compute_count               = 2

  contacts {
    email        = "shravan.thatikonda@oracle.com"
    is_primary   = "true"
    name         = "Shravan Thatikonda"
    phone_number = "1234567891"
  }

  display_name = "PrimaryExaDataInfrastructure"
  dns_server = [
  "10.231.225.65"]
  gateway                  = "10.32.88.5"
  infini_band_network_cidr = "10.31.8.0/21"
  netmask                  = "255.255.255.0"
  ntp_server = [
  "10.231.225.76"]
  shape     = "ExadataCC.X7"
  time_zone = "US/Pacific"
}

resource "oci_database_exadata_infrastructure" "standby_exadata_infrastructure" {
  activation_file             = local_file.activation_file.filename
  admin_network_cidr          = "192.168.0.0/16"
  cloud_control_plane_server1 = "10.32.88.1"
  cloud_control_plane_server2 = "10.32.88.3"
  compartment_id              = var.compartment_ocid
  storage_count               = 3
  compute_count               = 2

  contacts {
    email        = "johndoe@acme.com"
    is_primary   = "true"
    name         = "John Doe"
    phone_number = "1234567891"
  }

  display_name = "StandbyExaDataInfrastructure"
  dns_server = [
  "10.231.225.65"]
  gateway                  = "10.32.88.5"
  infini_band_network_cidr = "10.31.8.0/21"
  netmask                  = "255.255.255.0"
  ntp_server = [
  "10.231.225.76"]
  shape     = "ExadataCC.X7"
  time_zone = "US/Pacific"
}

resource "oci_database_vm_cluster_network" "primary_vm_cluster_network" {
  compartment_id = var.compartment_ocid
  display_name   = "primaryVmClusterRecommendedNetwork"
  dns = [
  "192.168.10.10"]
  ntp = [
  "192.168.10.20"]

  exadata_infrastructure_id = oci_database_exadata_infrastructure.primary_exadata_infrastructure.id

  scans {
    hostname = "primary-prefix-nsubz-scan"
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
      db_server_id = data.oci_database_db_servers.primary_db_servers.db_servers.0.id
    }
    nodes {
      hostname = "myprefix2-cghdm2"
      ip       = "192.169.19.20"
      db_server_id = data.oci_database_db_servers.primary_db_servers.db_servers.1.id

    }
    vlan_id = "11"
  }
  vm_networks {
    domain_name  = "oracle.com"
    gateway      = "192.168.20.1"
    netmask      = "255.255.0.0"
    network_type = "CLIENT"
    nodes {
      hostname     = "primaryprefix1-r64zc1"
      ip           = "192.168.19.10"
      vip          = "192.168.19.11"
      vip_hostname = "myprefix1-r64zc1-vip"
      db_server_id = data.oci_database_db_servers.primary_db_servers.db_servers.0.id
    }
    nodes {
      hostname     = "primaryprefix1-r64zc2"
      ip           = "192.168.19.14"
      vip          = "192.168.19.15"
      vip_hostname = "primaryprefix1-r64zc2-vip"
      db_server_id = data.oci_database_db_servers.primary_db_servers.db_servers.1.id
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
}

resource "oci_database_vm_cluster_network" "standby_vm_cluster_network" {
  compartment_id = var.compartment_ocid
  display_name   = "standbyVmClusterRecommendedNetwork"
  dns = [
  "192.168.10.10"]
  ntp = [
  "192.168.10.20"]

  exadata_infrastructure_id = oci_database_exadata_infrastructure.standby_exadata_infrastructure.id

  scans {
    hostname = "standby-prefix-nsubz-scan"
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
      db_server_id = data.oci_database_db_servers.standby_db_servers.db_servers.0.id
    }
    nodes {
      hostname = "myprefix2-cghdm2"
      ip       = "192.169.19.20"
      db_server_id = data.oci_database_db_servers.standby_db_servers.db_servers.1.id
    }
    vlan_id = "11"
  }
  vm_networks {
    domain_name  = "oracle.com"
    gateway      = "192.168.20.1"
    netmask      = "255.255.0.0"
    network_type = "CLIENT"
    nodes {
      hostname     = "standbyprefix1-r64zc1"
      ip           = "192.168.19.10"
      vip          = "192.168.19.11"
      vip_hostname = "standbyprefix1-r64zc1-vip"
    db_server_id = data.oci_database_db_servers.standby_db_servers.db_servers.0.id

        }
    nodes {
      hostname     = "standbyprefix1-r64zc2"
      ip           = "192.168.19.14"
      vip          = "192.168.19.15"
      vip_hostname = "standbyprefix1-r64zc2-vip"
    db_server_id = data.oci_database_db_servers.standby_db_servers.db_servers.1.id

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
}

resource "oci_database_autonomous_vm_cluster" "primary_autonomous_vm_cluster" {
  #Required
  compartment_id                        = var.compartment_ocid
  display_name                          = "PrimaryVmCluster"
  exadata_infrastructure_id             = oci_database_exadata_infrastructure.primary_exadata_infrastructure.id
  vm_cluster_network_id                 = oci_database_vm_cluster_network.primary_vm_cluster_network.id
  cpu_core_count_per_node   = "20"
  autonomous_data_storage_size_in_tbs = "2.0"
  memory_per_oracle_compute_unit_in_gbs = "5"
  total_container_databases             = "2"
  #Optional
  is_local_backup_enabled = "false"
  license_model           = "LICENSE_INCLUDED"
  time_zone               = "US/Pacific"
  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "SampleTagValue"
  }

  freeform_tags = {
    "Department" = "Finance"
  }
}

resource "oci_database_autonomous_vm_cluster" "standby_autonomous_vm_cluster" {
  #Required
  compartment_id                        = var.compartment_ocid
  display_name                          = "StandbyVmCluster"
  exadata_infrastructure_id             = oci_database_exadata_infrastructure.standby_exadata_infrastructure.id
  vm_cluster_network_id                 = oci_database_vm_cluster_network.standby_vm_cluster_network.id
  cpu_core_count_per_node   = "20"
  autonomous_data_storage_size_in_tbs = "2.0"
  memory_per_oracle_compute_unit_in_gbs = "5"
  total_container_databases             = "2"
  #Optional
  is_local_backup_enabled = "false"
  license_model           = "LICENSE_INCLUDED"
  time_zone               = "US/Pacific"
  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "SampleTagValue"
  }

  freeform_tags = {
    "Department" = "Finance"
  }
}

data "oci_database_db_servers" "primary_db_servers" {
  #Required
  compartment_id            = var.compartment_ocid
  exadata_infrastructure_id = oci_database_exadata_infrastructure.primary_exadata_infrastructure.id
}

data "oci_database_db_servers" "standby_db_servers" {
  #Required
  compartment_id            = var.compartment_ocid
  exadata_infrastructure_id = oci_database_exadata_infrastructure.standby_exadata_infrastructure.id
}

resource "oci_database_autonomous_container_database" "test_autonomous_container_database_primary" {
  #Required
  autonomous_vm_cluster_id             = oci_database_autonomous_vm_cluster.primary_autonomous_vm_cluster.id
  display_name                         = "PrimaryACD"
  patch_model                          = "RELEASE_UPDATES"
  db_version                           = "19.28.0.1.0"
  db_name                              = "PRIMARY"

  #Optional
  backup_config {
      backup_destination_details {
        type = "LOCAL"
      }
      recovery_window_in_days = "7"
    }

  compartment_id               = var.compartment_ocid
  freeform_tags                = var.autonomous_database_freeform_tags
  service_level_agreement_type = "STANDARD"

  maintenance_window_details {
    preference = "CUSTOM_PREFERENCE"

    days_of_week {
      name = "MONDAY"
    }

    hours_of_day = ["4"]

    months {
      name = "JANUARY"
    }

    months {
      name = "APRIL"
    }

    months {
      name = "JULY"
    }

    months {
      name = "OCTOBER"
    }

    weeks_of_month = ["2"]
  }
  version_preference = "LATEST_RELEASE_UPDATE"

    lifecycle {
      ignore_changes = [
          peer_autonomous_container_database_display_name,
          peer_autonomous_exadata_infrastructure_id,
          peer_autonomous_vm_cluster_id,
          peer_cloud_autonomous_vm_cluster_id,
          peer_db_unique_name,
          service_level_agreement_type,
          protection_mode,
          peer_autonomous_container_database_backup_config,
      ]
    }

}

resource "oci_database_autonomous_container_database_add_standby" "test_autonomous_container_database_add_standby" {
  autonomous_container_database_id = oci_database_autonomous_container_database.test_autonomous_container_database_primary.id

#  add standby blocked auto fail over
#  fast_start_fail_over_lag_limit_in_seconds = "0"
#  is_automatic_failover_enabled = "true"

  peer_autonomous_container_database_backup_config {
    backup_destination_details {
      type = "LOCAL"
    }
    recovery_window_in_days = "7"
  }
  peer_autonomous_container_database_compartment_id = var.compartment_ocid
  peer_autonomous_container_database_display_name = "FirstStandby"
  peer_autonomous_vm_cluster_id = oci_database_autonomous_vm_cluster.standby_autonomous_vm_cluster.id
  protection_mode = "MAXIMUM_AVAILABILITY"
  standby_maintenance_buffer_in_days = "7"
}

resource "oci_database_autonomous_container_database_snapshot_standby" "test_autonomous_container_database_snapshot_standby" {
  autonomous_container_database_id = oci_database_autonomous_container_database_add_standby.test_autonomous_container_database_add_standby.dataguard_group_members.1.autonomous_container_database_id
  role = "SNAPSHOT_STANDBY"
  connection_strings_type="SNAPSHOT_SERVICES"
  depends_on = [oci_database_autonomous_container_database_add_standby.test_autonomous_container_database_add_standby]
}

resource "oci_database_autonomous_container_database_snapshot_standby" "test_autonomous_container_database_regular_standby" {
  autonomous_container_database_id = oci_database_autonomous_container_database_add_standby.test_autonomous_container_database_add_standby.dataguard_group_members.1.autonomous_container_database_id
  role = "STANDBY"
  depends_on = [oci_database_autonomous_container_database_snapshot_standby.test_autonomous_container_database_snapshot_standby]
}

#### End Resources ####
#######################





