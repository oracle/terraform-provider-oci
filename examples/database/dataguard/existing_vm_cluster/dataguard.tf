// Copyright (c) 2019, Oracle and/or its affiliates. All rights reserved.

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "compartment_id" {
}

variable "region" {
}

variable "ssh_public_key" {
  default = "ssh-rsa sample"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

resource "oci_core_security_list" "exadata_shapes_security_list" {
  compartment_id = var.compartment_id
  vcn_id         = oci_core_virtual_network.t.id
  display_name   = "TFExampleSecurityList"

  // allow outbound tcp traffic on all ports
  egress_security_rules {
    destination = "0.0.0.0/0"
    protocol    = "6"
  }

  ingress_security_rules {
    protocol = "6"
    source   = "0.0.0.0/0"
  }
}

#dataguard requires the port to be open on the subnet
resource "oci_core_virtual_network" "t" {
  compartment_id = var.compartment_id
  cidr_block     = "10.1.0.0/16"
  display_name   = "-tf-vcn"
  dns_label      = "tfvcn"
}

// An AD based subnet will supply an Availability Domain
resource "oci_core_subnet" "test_subnet" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  cidr_block          = "10.1.22.0/24"
  display_name        = "ExadataSubnet"
  compartment_id      = var.compartment_id
  vcn_id              = oci_core_virtual_network.t.id
  route_table_id      = oci_core_virtual_network.t.default_route_table_id
  dhcp_options_id     = oci_core_virtual_network.t.default_dhcp_options_id
  security_list_ids   = [oci_core_virtual_network.t.default_security_list_id, oci_core_security_list.exadata_shapes_security_list.id]
  dns_label           = "subnetexadata"
}

resource "oci_database_exadata_infrastructure" "test_exadata_infrastructure" {
  #Required
  admin_network_cidr          = "192.168.0.0/16"
  cloud_control_plane_server1 = "10.32.88.1"
  cloud_control_plane_server2 = "10.32.88.3"
  compartment_id              = var.compartment_id
  display_name                = "tstExaInfra"
  dns_server                  = ["10.231.225.65"]
  gateway                     = "10.32.88.5"
  infini_band_network_cidr    = "10.31.8.0/21"
  netmask                     = "255.255.0.0"
  ntp_server                  = ["10.231.225.76"]
  shape                       = "ExadataCC.Quarter3.100"
  time_zone                   = "US/Pacific"
  activation_file             = "activation.zip"

  #Optional
  corporate_proxy = "http://192.168.19.1:80"
}

data "oci_database_gi_versions" "gi_version" {
  compartment_id = var.compartment_id
  shape = "ExadataCC.Quarter3.100"
}

resource "oci_database_vm_cluster_network" "test_vm_cluster_network" {
  compartment_id = var.compartment_id

  display_name              = "testVmClusterNw"
  dns                       = ["192.168.10.12"]
  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id

  ntp = ["192.168.10.22"]

  scans {
    hostname = "myprefix2-ivmmj-scan"
    ips      = ["192.168.19.7", "192.168.19.8", "192.168.19.9"]
    port     = "1522"
  }

  validate_vm_cluster_network = "true"

  vm_networks {
    domain_name  = "oracle.com"
    gateway      = "192.169.20.2"
    netmask      = "255.255.192.0"
    network_type = "BACKUP"

    nodes {
      hostname = "myprefix2-xapb24"
      ip       = "192.169.19.19"
    }

    nodes {
      hostname = "myprefix2-xapb28"
      ip       = "192.169.19.21"
    }

    vlan_id = "100"
  }

  vm_networks {
    domain_name  = "oracle.com"
    gateway      = "192.168.20.2"
    netmask      = "255.255.192.0"
    network_type = "CLIENT"

    nodes {
      hostname     = "myprefix2-xapb22"
      ip           = "192.168.19.11"
      vip          = "192.168.19.13"
      vip_hostname = "myprefix2-xapb22-vip"
    }

    nodes {
      hostname     = "myprefix2-xapb26"
      ip           = "192.168.19.15"
      vip          = "192.168.19.17"
      vip_hostname = "myprefix2-xapb26-vip"
    }

    vlan_id = "101"
  }
}

resource "oci_database_vm_cluster" "test_exadata_vm_cluster_for_primary_db" {
  compartment_id            = var.compartment_id
  cpu_core_count            = "4"
  depends_on                = [oci_database_vm_cluster_network.test_vm_cluster_network]
  display_name              = "vmClusterForPrimaryDB"
  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
  gi_version                = data.oci_database_gi_versions.gi_version.gi_versions.0.version
  ssh_public_keys           = [var.ssh_public_key]
  vm_cluster_network_id     = oci_database_vm_cluster_network.test_vm_cluster_network.id
}

resource "oci_database_vm_cluster" "test_exadata_vm_cluster_for_standby_db" {
  compartment_id            = var.compartment_id
  cpu_core_count            = "4"
  depends_on                = [oci_database_vm_cluster_network.test_vm_cluster_network]
  display_name              = "vmClusterForStandbyDB"
  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
  gi_version                = data.oci_database_gi_versions.gi_version.gi_versions.0.version
  ssh_public_keys           = [var.ssh_public_key]
  vm_cluster_network_id     = oci_database_vm_cluster_network.test_vm_cluster_network.id
}

data "oci_database_databases" "exadb" {
  compartment_id = var.compartment_id
  db_home_id     = oci_database_db_home.test_db_home_vm_cluster.id
}

resource "oci_database_db_home" "test_db_home_vm_cluster" {
  vm_cluster_id = oci_database_vm_cluster.test_exadata_vm_cluster_for_primary_db.id

  database {
    admin_password = "BEstrO0ng_#11"
    db_name        = "dbVMClus"
    character_set  = "AL32UTF8"
    ncharacter_set = "AL16UTF16"
    db_workload    = "OLTP"
    pdb_name       = "pdbName"
  }

  source       = "VM_CLUSTER_NEW"
  db_version   = "12.1.0.2"
  display_name = "TFTestDbHome1"
}

resource "oci_database_data_guard_association" "test_exadata_data_guard_association" {
  creation_type                    = "ExistingVmCluster"
  database_admin_password          = "BEstrO0ng_#11"
  database_id                      = data.oci_database_databases.exadb.databases[0].id
  delete_standby_db_home_on_delete = "true"
  depends_on = [
    oci_database_vm_cluster.test_exadata_vm_cluster_for_primary_db,
    oci_database_vm_cluster.test_exadata_vm_cluster_for_standby_db,
  ]
  peer_vm_cluster_id = oci_database_vm_cluster.test_exadata_vm_cluster_for_standby_db.id
  protection_mode    = "MAXIMUM_PERFORMANCE"
  transport_type     = "ASYNC"
}

data "oci_database_data_guard_association" "test_exadata_data_guard_association_for_primary" {
  data_guard_association_id = oci_database_data_guard_association.test_exadata_data_guard_association.id
  database_id               = data.oci_database_databases.exadb.databases[0].id
}

