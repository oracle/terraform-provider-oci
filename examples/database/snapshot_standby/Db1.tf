variable "tenancy_ocid" {
}

variable "compartment_id" {
}

variable "ssh_public_key" {
}

provider "oci" {
    version          = "0.0.0"
}

data "oci_identity_availability_domains" "ADs" {
    compartment_id = var.compartment_id

}
data "oci_identity_availability_domain" "ad" {
    compartment_id 		= var.compartment_id
    ad_number      		= 1

}
resource "oci_core_virtual_network" "t" {
    compartment_id = var.compartment_id
    cidr_block = "10.1.0.0/16"
    display_name = "-tf-vcn"
    dns_label = "tfvcn"

}
resource "oci_core_route_table" "t" {
    compartment_id = var.compartment_id
    vcn_id = oci_core_virtual_network.t.id
    route_rules {
        cidr_block = "0.0.0.0/0"
        network_entity_id = oci_core_internet_gateway.t.id
    }
}
resource "oci_core_internet_gateway" "t" {
    compartment_id = var.compartment_id
    vcn_id = oci_core_virtual_network.t.id
    display_name = "-tf-internet-gateway"
}

resource "oci_core_subnet" "t" {
    availability_domain = data.oci_identity_availability_domains.ADs.availability_domains.0.name
    cidr_block          = "10.1.20.0/24"
    display_name        = "TFSubnet1"
    compartment_id      = var.compartment_id
    vcn_id              = oci_core_virtual_network.t.id
    route_table_id      = oci_core_route_table.t.id
    dhcp_options_id     = oci_core_virtual_network.t.default_dhcp_options_id
    security_list_ids   = [oci_core_virtual_network.t.default_security_list_id]
    dns_label           = "tfsubnet"
}
resource "oci_core_subnet" "t2" {
    availability_domain = data.oci_identity_availability_domains.ADs.availability_domains.0.name
    cidr_block          = "10.1.21.0/24"
    display_name        = "TFSubnet2"
    compartment_id      = var.compartment_id
    vcn_id              = oci_core_virtual_network.t.id
    route_table_id      = oci_core_route_table.t.id
    dhcp_options_id     = oci_core_virtual_network.t.default_dhcp_options_id
    security_list_ids   = [oci_core_virtual_network.t.default_security_list_id]
    dns_label           = "tfsubnet2"
}
resource "oci_core_network_security_group" "test_network_security_group" {
     compartment_id  = var.compartment_id
     vcn_id            = oci_core_virtual_network.t.id
     display_name      =  "displayName"
}

resource "oci_core_network_security_group" "test_network_security_group_backup" {
    compartment_id = var.compartment_id
    vcn_id            = oci_core_virtual_network.t.id
}

resource "oci_core_subnet" "test_subnet1" {
    availability_domain = data.oci_identity_availability_domain.ad.name
    cidr_block          = "10.1.22.0/24"
    display_name        = "ExadataSubnet"
    compartment_id      = var.compartment_id
    vcn_id              = oci_core_virtual_network.t.id
    route_table_id      = oci_core_virtual_network.t.default_route_table_id
    dhcp_options_id     = oci_core_virtual_network.t.default_dhcp_options_id
    security_list_ids   = [oci_core_virtual_network.t.default_security_list_id, oci_core_security_list.exadata_shapes_security_list.id]
    dns_label           = "subnetexadata1"
}

resource "oci_core_subnet" "test_subnet_backup" {
    availability_domain = data.oci_identity_availability_domain.ad.name
    cidr_block          = "10.1.23.0/24"
    display_name        = "ExadataBackupSubnet"
    compartment_id      = var.compartment_id
    vcn_id              = oci_core_virtual_network.t.id
    route_table_id      = oci_core_virtual_network.t.default_route_table_id
    dhcp_options_id     = oci_core_virtual_network.t.default_dhcp_options_id
    security_list_ids   = [oci_core_virtual_network.t.default_security_list_id]
    dns_label           = "subnetexadata2"
}


resource "oci_core_security_list" "exadata_shapes_security_list" {
    compartment_id = var.compartment_id
    vcn_id         = oci_core_virtual_network.t.id
    display_name   = "ExadataSecurityList"
   ingress_security_rules {
       source    = "10.1.22.0/24"
       protocol  = "6"
   }

   ingress_security_rules {
       source    = "10.1.22.0/24"
       protocol  = "1"
   }

   egress_security_rules {
       destination = "10.1.22.0/24"
       protocol    = "6"
   }

   egress_security_rules {
       destination = "10.1.22.0/24"
       protocol    = "1"
   }
}

resource "oci_database_cloud_exadata_infrastructure" "test_cloud_exadata_infrastructure_primary" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_id
  display_name        = "TstExaInfra1"
  shape               = "Exadata.X8M"

  #Optional
  compute_count = 2
  storage_count = 3
}

resource "oci_database_cloud_vm_cluster" "test_cloud_vm_cluster" {
  #Required
  backup_subnet_id                = oci_core_subnet.test_subnet_backup.id
  cloud_exadata_infrastructure_id = oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure_primary.id
  compartment_id                  = var.compartment_id
  cpu_core_count                  = "22"
  display_name                    = "PrimaryCluster"
  gi_version                      = "19.0.0.0"
  hostname                        = "myOracleDB"
  ssh_public_keys                 = [var.ssh_public_key]
  subnet_id                       = oci_core_subnet.t.id

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

resource "oci_database_db_home" "test_db_home" {
  db_system_id = oci_database_cloud_vm_cluster.test_cloud_vm_cluster.id

  db_version   = "19.26.0.0"
  source = "NONE"
  display_name = "createdDbHomeNone"
}

resource "oci_database_database" "primary_database" {
  #Required
  database {
    admin_password = "BEstrO0ng_#11"
    db_name        = "PrimDb"
    character_set  = "AL32UTF8"
    ncharacter_set = "AL16UTF16"
    db_workload    = "OLTP"
  }

  db_home_id = oci_database_db_home.test_db_home.id
  source     = "NONE"
}
