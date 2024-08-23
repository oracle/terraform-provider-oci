// Copyright (c) 2019, Oracle and/or its affiliates. All rights reserved.

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "compartment_ocid" {
}

variable "region" {
}

variable "ssh_public_key" {
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

data "oci_core_services" "test_services" {
  filter {
    name   = "name"
    values = [".*Oracle.*Services.*Network"]
    regex  = true
  }
}

data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.tenancy_ocid
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "TFExampleVCN"
  dns_label      = "tfexamplevcn"
}

resource "oci_core_service_gateway" "test_service_gateway" {
  compartment_id = var.compartment_ocid
  display_name   = "test_service_gateway"
  vcn_id         = oci_core_vcn.test_vcn.id

  services {
    service_id = data.oci_core_services.test_services.services.0.id
  }

}

resource "oci_core_default_route_table" "test_vcn__default_route_table" {
  manage_default_resource_id = oci_core_vcn.test_vcn.default_route_table_id

  route_rules {
    network_entity_id = oci_core_service_gateway.test_service_gateway.id

    description 	  = "Internal traffic for OCI Services"
    destination       = data.oci_core_services.test_services.services[0].cidr_block
    destination_type  = "SERVICE_CIDR_BLOCK"
  }
}

resource "oci_core_route_table" "test_route_table" {
  compartment_id = var.compartment_ocid
  display_name   = "test_subnet_rt"
  vcn_id         = oci_core_vcn.test_vcn.id

  route_rules {
    network_entity_id = oci_core_service_gateway.test_service_gateway.id

    description 	  = "Internal traffic for OCI Services"
    destination       = data.oci_core_services.test_services.services[0].cidr_block
    destination_type  = "SERVICE_CIDR_BLOCK"
  }
}

resource "oci_core_security_list" "test_security_list" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
  display_name   = "test_security_list"

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

resource "oci_core_subnet" "test_subnet" {
  cidr_block          	   = "10.0.2.0/24"
  compartment_id      	   = var.compartment_ocid
  vcn_id              	   = oci_core_vcn.test_vcn.id
  display_name        	   = "test_subnet"
  security_list_ids   	   = [oci_core_security_list.test_security_list.id]
  route_table_id      	   = oci_core_route_table.test_route_table.id
  dhcp_options_id     	   = oci_core_vcn.test_vcn.default_dhcp_options_id
  dns_label           	   = "tftestsubnet"
  prohibit_public_ip_on_vnic = "true"
}


resource "oci_database_db_system" "test_db_system" {
  availability_domain     = data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name
  compartment_id          = var.compartment_ocid
  subnet_id               = oci_core_subnet.test_subnet.id
  database_edition        = "ENTERPRISE_EDITION"
  disk_redundancy         = "NORMAL"
  shape                   = "VM.Standard2.1"
  ssh_public_keys         = [var.ssh_public_key]
  domain                  = oci_core_subnet.test_subnet.subnet_domain_name
  hostname                = "myOracleDB"
  data_storage_size_in_gb = "256"
  license_model           = "LICENSE_INCLUDED"
  node_count              = "1"
  display_name            = "TFExampleDbSystem"

  db_home {
    db_version   = "12.1.0.2"
    display_name = "TFExampleDbHome"

    database {
      admin_password = "BEstrO0ng_#11"
      db_name        = "db1"
    }
  }
}

data "oci_database_db_homes" "t" {
  compartment_id = var.compartment_ocid
  db_system_id   = oci_database_db_system.test_db_system.id

  filter {
    name   = "display_name"
    values = ["TFExampleDbHome"]
  }
}

data "oci_database_databases" "db" {
  compartment_id = var.compartment_ocid
  db_home_id     = data.oci_database_db_homes.t.db_homes[0].db_home_id
}

resource "oci_database_data_guard_association" "test_data_guard_association" {
  #Required
  creation_type                    = "NewDbSystem"
  database_admin_password          = "BEstrO0ng_#11"
  database_id                      = data.oci_database_databases.db.databases[0].id
  protection_mode                  = "MAXIMUM_PERFORMANCE"
  transport_type                   = "ASYNC"
  delete_standby_db_home_on_delete = "true"

  #required for NewDbSystem creation_type
  display_name              = "TFExampleDataGuardAssociationVM"
  shape                     = "VM.Standard2.1"
  subnet_id                 = oci_core_subnet.test_subnet.id
  availability_domain       = data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name
  nsg_ids                   = [oci_core_network_security_group.test_network_security_group.id]
  hostname                  = "ocidb"
  db_system_defined_tags    = map("example-tag-namespace-all.example-tag", "dbSystemDefinedTags1")
  db_system_freeform_tags   = {"dbSystemFreeformTagsK" = "dbSystemFreeformTagsV"}
  database_defined_tags     = map("example-tag-namespace-all.example-tag", "databaseDefinedTags1")
  database_freeform_tags    = {"databaseFreeformTagsK" = "databaseFreeformTagsV"}
  fault_domains             = ["FAULT-DOMAIN-3"]
  license_model             = "LICENSE_INCLUDED"
  node_count                = "1"
  private_ip                = "10.0.2.223"
  time_zone                 = "US/Pacific"
  is_active_data_guard_enabled = "false"
}

resource "oci_core_network_security_group" "test_network_security_group" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
  display_name   = "tf-example-nsg"
}

