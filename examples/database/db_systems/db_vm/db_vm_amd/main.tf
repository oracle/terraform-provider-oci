// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

#variable "kms_key_id" {
#}
#
#variable "kms_key_version_id" {
#}
#
#variable "vault_id" {
#}

variable "compartment_ocid" {
}

variable "ssh_public_key" {
}

variable "ssh_private_key" {
}

# DBSystem specific
variable "db_system_shape" {
  default = "VM.Standard.E4.Flex"
}

variable "cpu_core_count" {
  default = "2"
}

variable "db_system_storage_volume_performance_mode" {
  default = "BALANCED"
}

variable "db_edition" {
  default = "ENTERPRISE_EDITION"
}

variable "db_admin_password" {
  default = "BEstrO0ng_#12"
}

variable "db_version" {
  default = "19.0.0.0"
}

variable "db_disk_redundancy" {
  default = "NORMAL"
}

variable "sparse_diskgroup" {
  default = true
}

variable "hostname" {
  default = "myoracledb"
}

variable "host_user_name" {
  default = "opc"
}

variable "n_character_set" {
  default = "AL16UTF16"
}

variable "character_set" {
  default = "AL32UTF8"
}

variable "db_workload" {
  default = "OLTP"
}

variable "pdb_name" {
  default = "pdbName"
}

variable "data_storage_size_in_gb" {
  default = "256"
}

variable "license_model" {
  default = "LICENSE_INCLUDED"
}

variable "node_count" {
  default = "1"
}

variable "test_database_software_image_ocid" {

}

provider "oci" {
#  version = "4.70.0"
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

# Get DB node list
data "oci_database_db_nodes" "db_nodes" {
  compartment_id = var.compartment_ocid
  db_system_id   = oci_database_db_system.test_db_system.id
}

# Get DB node details
data "oci_database_db_node" "db_node_details" {
  db_node_id = data.oci_database_db_nodes.db_nodes.db_nodes[0]["id"]
}

# Gets the OCID of the first (default) vNIC
#data "oci_core_vnic" "db_node_vnic" {
#    vnic_id = data.oci_database_db_node.db_node_details.vnic_id
#}

data "oci_database_db_homes" "db_homes" {
  compartment_id = var.compartment_ocid
  db_system_id   = oci_database_db_system.test_db_system.id
}

data "oci_database_databases" "databases" {
  compartment_id = var.compartment_ocid
  db_home_id     = data.oci_database_db_homes.db_homes.db_homes[0].db_home_id
}

d



ata "oci_database_db_versions" "test_db_versions_by_db_system_id" {
  compartment_id = var.compartment_ocid
  db_system_id   = oci_database_db_system.test_db_system.id
}

resource "oci_database_backup" "test_backup" {
  database_id = "${data.oci_database_databases.databases.databases.0.id}"
  display_name = "Monthly Backup"
}

data "oci_database_db_system_shapes" "test_db_system_shapes" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid

  filter {
    name   = "shape"
    values = [var.db_system_shape]
  }
}

data "oci_database_db_systems" "db_systems" {
  compartment_id = var.compartment_ocid

  filter {
    name   = "id"
    values = [oci_database_db_system.test_db_system.id]
  }
}

resource "oci_core_vcn" "vcn" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "TFExampleVCNDBSystem"
  dns_label      = "tfexvcndbsys"
}

resource "oci_core_subnet" "subnet" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  cidr_block          = "10.1.20.0/24"
  display_name        = "TFExampleSubnetDBSystem"
  dns_label           = "tfexsubdbsys"
  security_list_ids   = [oci_core_security_list.ExampleSecurityList.id]
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.vcn.id
  route_table_id      = oci_core_route_table.route_table.id
  dhcp_options_id     = oci_core_vcn.vcn.default_dhcp_options_id
}

resource "oci_core_subnet" "subnet_backup" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  cidr_block          = "10.1.1.0/24"
  display_name        = "TFExampleSubnetDBSystemBackup"
  dns_label           = "tfexsubdbsysbp"
  security_list_ids   = [oci_core_security_list.ExampleSecurityList.id]
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.vcn.id
  route_table_id      = oci_core_route_table.route_table_backup.id
  dhcp_options_id     = oci_core_vcn.vcn.default_dhcp_options_id
}

resource "oci_core_internet_gateway" "internet_gateway" {
  compartment_id = var.compartment_ocid
  display_name   = "TFExampleIGDBSystem"
  vcn_id         = oci_core_vcn.vcn.id
}

resource "oci_core_route_table" "route_table" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.vcn.id
  display_name   = "TFExampleRouteTableDBSystem"

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = oci_core_internet_gateway.internet_gateway.id
  }
}

resource "oci_core_route_table" "route_table_backup" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.vcn.id
  display_name   = "TFExampleRouteTableDBSystemBackup"

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = oci_core_internet_gateway.internet_gateway.id
  }
}

resource "oci_core_security_list" "ExampleSecurityList" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.vcn.id
  display_name   = "TFExampleSecurityList"

  // allow outbound tcp traffic on all ports
  egress_security_rules {
    destination = "0.0.0.0/0"
    protocol    = "6"
  }

  // allow outbound udp traffic on a port range
  egress_security_rules {
    destination = "0.0.0.0/0"
    protocol    = "17" // udp
    stateless   = true
  }

  egress_security_rules {
    destination = "0.0.0.0/0"
    protocol    = "1"
    stateless   = true
  }

  // allow inbound ssh traffic from a specific port
  ingress_security_rules {
    protocol  = "6" // tcp
    source    = "0.0.0.0/0"
    stateless = false
  }

  // allow inbound icmp traffic of a specific type
  ingress_security_rules {
    protocol  = 1
    source    = "0.0.0.0/0"
    stateless = true
  }
}

resource "oci_core_network_security_group" "test_network_security_group" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.vcn.id
  display_name   = "displayName"
}

resource "oci_core_network_security_group" "test_network_security_group_backup" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.vcn.id
  display_name   = "displayName"
}

resource "oci_database_db_system" "test_db_system" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  database_edition    = var.db_edition

  db_home {
    database {
      admin_password = var.db_admin_password
#      kms_key_version_id    = var.kms_key_version_id
#      kms_key_id     = var.kms_key_id
#      vault_id       = var.vault_id
      db_name        = "aTFdbVm"
      character_set  = var.character_set
      ncharacter_set = var.n_character_set
      db_workload    = var.db_workload
      pdb_name       = var.pdb_name

      db_backup_config {
        auto_backup_enabled = false
      }
    }

    db_version   = "19.15.0.0"
    display_name = "MyTFDBHomeVm"
  }

  db_system_options {
    storage_management = "LVM"
  }

  disk_redundancy         = var.db_disk_redundancy
  shape                   = var.db_system_shape
  cpu_core_count          = var.cpu_core_count
  storage_volume_performance_mode = var.db_system_storage_volume_performance_mode
  subnet_id               = oci_core_subnet.subnet.id
  ssh_public_keys         = [var.ssh_public_key]
  display_name            = "MyTFDBSystemVM"
  hostname                = var.hostname
  data_storage_size_in_gb = var.data_storage_size_in_gb
  license_model           = var.license_model
  node_count              = data.oci_database_db_system_shapes.test_db_system_shapes.db_system_shapes[0]["minimum_node_count"]
  nsg_ids                 = [oci_core_network_security_group.test_network_security_group_backup.id, oci_core_network_security_group.test_network_security_group.id]

  #To use defined_tags, set the values below to an existing tag namespace, refer to the identity example on how to create tag namespaces
  #defined_tags  = {"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"}

  freeform_tags = {
    "Department" = "Finance"
  }
}

resource "oci_database_db_system" "db_system_bkup" {
  source = "DB_BACKUP"
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id = var.compartment_ocid
  subnet_id = oci_core_subnet.subnet.id
  database_edition = var.db_edition
  disk_redundancy = var.db_disk_redundancy
  shape = var.db_system_shape
  cpu_core_count= var.cpu_core_count
  storage_volume_performance_mode= var.db_system_storage_volume_performance_mode
  ssh_public_keys         = [var.ssh_public_key]
  hostname = var.hostname
  data_storage_size_in_gb = var.data_storage_size_in_gb
  license_model = var.license_model
  node_count = data.oci_database_db_system_shapes.test_db_system_shapes.db_system_shapes[0]["minimum_node_count"]
  display_name = "tfDbSystemFromBackupWithCustImg"

  db_home {
    db_version = "19.15.0.0"
#    database_software_image_id = var.test_database_software_image_ocid
    database {
      admin_password = "BEstrO0ng_#11"
      backup_tde_password = "BEstrO0ng_#11"
      backup_id = "${oci_database_backup.test_backup.id}"
      db_name = "dbback"
    }
  }
}
