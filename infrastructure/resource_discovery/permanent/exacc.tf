// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_database_exadata_infrastructure" "exadata_infrastructure_rd" {
  #Required
  admin_network_cidr          = "192.168.0.0/16"
  cloud_control_plane_server1 = "192.168.19.1"
  cloud_control_plane_server2 = "192.168.19.2"
  compartment_id              = "${var.compartment_ocid}"
  display_name                = "ExaInfraRD"
  dns_server                  = ["192.168.10.10"]
  gateway                     = "192.168.20.1"
  infini_band_network_cidr    = "10.172.0.0/19"
  netmask                     = "255.255.0.0"
  ntp_server                  = ["192.168.10.20"]
  shape                       = "ExadataCC.Quarter3.100"
  time_zone                   = "US/Pacific"
  activation_file             = "activation.zip"

  #Optional
  corporate_proxy = "http://192.168.19.1:80"
  defined_tags    = "${map("${oci_identity_tag_namespace.tag_namespace_rd.name}.${oci_identity_tag.tag_rd.name}", "value1")}"

  freeform_tags = {
    "Department" = "Accounting"
  }
}

data "oci_database_gi_versions" "gi_version" {
  compartment_id = var.compartment_ocid
  shape = "ExadataCC.Quarter3.100"
}

resource "oci_database_vm_cluster_network" "vm_cluster_network_rd" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "VmClusterRecommendedNetworkRD"
  dns            = ["192.168.10.10"]
  ntp            = ["192.168.10.20"]

  exadata_infrastructure_id = "${oci_database_exadata_infrastructure.exadata_infrastructure_rd.id}"

  scans {
    hostname = "myprefix1-nsubz-scan"

    ips = [
      "192.168.19.7",
      "192.168.19.6",
      "192.168.19.8",
    ]

    port = 1521
  }

  vm_networks {
    domain_name  = "oracle.com"
    gateway      = "192.169.20.1"
    netmask      = "255.255.0.0"
    network_type = "BACKUP"

    nodes {
      hostname = "myprefix2-cghdm1"
      ip       = "192.169.19.18"
    }

    nodes {
      hostname = "myprefix2-cghdm2"
      ip       = "192.169.19.20"
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
    }

    nodes {
      hostname     = "myprefix1-r64zc2"
      ip           = "192.168.19.14"
      vip          = "192.168.19.15"
      vip_hostname = "myprefix1-r64zc2-vip"
    }

    vlan_id = "10"
  }

  #Optional
  defined_tags = "${map("${oci_identity_tag_namespace.tag_namespace_rd.name}.${oci_identity_tag.tag_rd.name}", "value1")}"

  freeform_tags = {
    "Department" = "Accounting"
  }

  validate_vm_cluster_network = true
}

resource "oci_database_vm_cluster" "test_vm_cluster" {
  #Required
  compartment_id            = "${var.compartment_ocid}"
  cpu_core_count            = "4"
  display_name              = "VmClusterRD"
  exadata_infrastructure_id = "${oci_database_exadata_infrastructure.exadata_infrastructure_rd.id}"
  gi_version                = data.oci_database_gi_versions.gi_version.gi_versions.0.version
  ssh_public_keys           = ["${var.ssh_public_key}"]
  vm_cluster_network_id     = "${oci_database_vm_cluster_network.vm_cluster_network_rd.id}"

  #Optional
  defined_tags = "${map("${oci_identity_tag_namespace.tag_namespace_rd.name}.${oci_identity_tag.tag_rd.name}", "value1")}"

  freeform_tags = {
    "Department" = "Accounting"
  }

  is_local_backup_enabled     = "false"
  is_sparse_diskgroup_enabled = "false"
  license_model               = "LICENSE_INCLUDED"
  data_storage_size_in_tbs    = "84"
  db_node_storage_size_in_gbs = "120"
  memory_size_in_gbs          = "60"
}

resource "oci_database_db_home" "test_db_home_vm_cluster" {
  vm_cluster_id = "${oci_database_vm_cluster.test_vm_cluster.id}"

  database {
    admin_password = "${var.db_admin_password}"
    db_name        = "VMClusRD"
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
        id   = "${oci_database_backup_destination.test_backup_destination_nfs.id}"
        type = "NFS"
      }
    }
  }

  source       = "VM_CLUSTER_NEW"
  db_version   = "12.1.0.2"
  display_name = "DbHomeRD"
}

resource "oci_database_backup_destination" "test_backup_destination_nfs" {
  #Required
  compartment_id = "${var.compartment_ocid}"
  display_name   = "testBackupDestinationRD"
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

resource "oci_database_database" "test_database_rd" {
  #Required
  database {
    admin_password = "${var.db_admin_password}"
    db_name        = "VMCluRD2"
    character_set  = "${var.character_set}"
    ncharacter_set = "${var.n_character_set}"
    db_workload    = "${var.db_workload}"

    db_backup_config {
      auto_backup_enabled = false
    }
  }

  db_home_id = "${oci_database_db_home.test_db_home_vm_cluster.id}"
  source     = "NONE"
}

resource "oci_database_autonomous_vm_cluster" "test_autonomous_vm_cluster" {
  #Required
  compartment_id            = "${var.compartment_ocid}"
  display_name              = "autonomousVmClusterRD"
  exadata_infrastructure_id = "${oci_database_exadata_infrastructure.exadata_infrastructure_rd.id}"
  vm_cluster_network_id     = "${oci_database_vm_cluster_network.vm_cluster_network_rd.id}"

  #Optional
  is_local_backup_enabled = "false"
  license_model           = "LICENSE_INCLUDED"
  time_zone               = "US/Pacific"
  defined_tags            = "${map("${oci_identity_tag_namespace.tag_namespace_rd.name}.${oci_identity_tag.tag_rd.name}", "value1")}"

  freeform_tags = {
    "Department" = "Finance"
  }
}
