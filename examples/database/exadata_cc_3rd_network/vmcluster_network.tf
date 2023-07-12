resource "oci_database_vm_cluster_network" "test_vm_cluster_network" {
  compartment_id = var.compartment_ocid
  display_name   = "testVmClusterRecommendedNetwork1"
  dns            = ["10.0.10.7"]
  ntp            = ["10.0.10.8"]

  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id

  scans {
    hostname = "clienthostname-igrik1-scan"

    ips = [
      "192.163.10.8",
      "192.163.10.9",
      "192.163.10.10",
    ]

    port = 1521
    scan_listener_port_tcp = 1521
    scan_listener_port_tcp_ssl = 2484
  }

  dr_scans {
    hostname = "drhostname-izt4c-drscan"

    ips = [
      "255.255.248.6",
      "255.255.248.7",
      "255.255.248.8",
    ]

    scan_listener_port_tcp = 1521
  }


  vm_networks {
    domain_name  = "backupdomainanme.com"
    gateway      = "192.161.13.1"
    netmask      = "255.255.248.0"
    network_type = "BACKUP"

    nodes {
      hostname = "backuphostname-q3rhi1"
      ip       = "192.161.13.2"
      db_server_id = data.oci_database_db_servers.test_db_servers.db_servers.0.id
    }

    nodes {
      hostname = "backuphostname-q3rhi2"
      ip       = "192.161.13.3"
      db_server_id = data.oci_database_db_servers.test_db_servers.db_servers.1.id
    }

    vlan_id = "232"
  }

  vm_networks {
    domain_name  = "oracle.com"
    gateway      = "192.163.10.3"
    netmask      = "255.255.224.0"
    network_type = "CLIENT"

    nodes {
      hostname     = "clienthostname-igrik1"
      ip           = "192.163.10.4"
      vip          = "192.163.10.5"
      vip_hostname = "clienthostname-igrik1-vip"
      db_server_id = data.oci_database_db_servers.test_db_servers.db_servers.0.id
    }

    nodes {
      hostname     = "clienthostname-igrik2"
      ip           = "192.163.10.6"
      vip          = "192.163.10.7"
      vip_hostname = "clienthostname-igrik2-vip"
      db_server_id = data.oci_database_db_servers.test_db_servers.db_servers.1.id
    }

    vlan_id = "324"
  }

  vm_networks {
    domain_name  = "drdomainname.com"
    gateway      = "255.255.248.1"
    netmask      = "255.255.248.0"
    network_type = "DISASTER_RECOVERY"

    nodes {
      hostname     = "drhostname-csuct1"
      ip           = "255.255.248.2"
      vip          = "255.255.248.3"
      vip_hostname = "drhostname-csuct1-vip"
      db_server_id = data.oci_database_db_servers.test_db_servers.db_servers.0.id
    }

    nodes {
      hostname     = "drhostname-csuct2"
      ip           = "255.255.248.4"
      vip          = "255.255.248.5"
      vip_hostname = "drhostname-csuct2-vip"
      db_server_id = data.oci_database_db_servers.test_db_servers.db_servers.1.id
    }

    vlan_id = "262"
  }

  #Optional
  /*defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedvalue"
  }

  freeform_tags = {
    "Department" = "Accounting"
  }*/

  validate_vm_cluster_network = true
  action = "ADD_DBSERVER_NETWORK"
  lifecycle {
    ignore_changes = [
      vm_networks,
    ]
  }
}

data "oci_database_db_servers" "test_db_servers" {
  #Required
  compartment_id            = var.compartment_ocid
  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
}
