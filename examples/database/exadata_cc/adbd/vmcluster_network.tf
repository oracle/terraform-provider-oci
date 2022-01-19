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
  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedvalue"
  }

  freeform_tags = {
    "Department" = "Accounting"
  }

  validate_vm_cluster_network = true
}

resource "oci_database_vm_cluster_network" "test_vm_cluster_network2" {
  compartment_id = var.compartment_ocid
  display_name   = "testVmClusterRecommendedNetwork"
  dns            = ["192.168.10.10"]
  ntp            = ["192.168.10.20"]

  exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id

  scans {
    hostname = "myprefix2-nsubz-scan"

    ips = [
      "192.168.19.37",
      "192.168.19.36",
      "192.168.19.38",
    ]

    port = 1521
    scan_listener_port_tcp = 1521
    scan_listener_port_tcp_ssl = 2484
  }

  vm_networks {
    domain_name  = "oracle.com"
    gateway      = "192.169.21.1"
    netmask      = "255.255.0.0"
    network_type = "BACKUP"

    nodes {
      hostname = "myprefix3-cghdm1"
      ip       = "192.169.19.48"
    }

    nodes {
      hostname = "myprefix3-cghdm2"
      ip       = "192.169.19.50"
    }

    vlan_id = "21"
  }

  vm_networks {
    domain_name  = "oracle.com"
    gateway      = "192.168.21.1"
    netmask      = "255.255.0.0"
    network_type = "CLIENT"

    nodes {
      hostname     = "myprefix4-r64zc1"
      ip           = "192.168.19.50"
      vip          = "192.168.19.51"
      vip_hostname = "myprefix4-r64zc1-vip"
    }

    nodes {
      hostname     = "myprefix4-r64zc2"
      ip           = "192.168.19.54"
      vip          = "192.168.19.55"
      vip_hostname = "myprefix4-r64zc2-vip"
    }

    vlan_id = "40"
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