Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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

variable "compartment_ocid" {
}

variable "ssh_public_key" {
}

variable "ssh_private_key" {
}

variable "source_db_system_id" {
}

variable "subnet_id" {
}

# DBSystem specific
variable "db_system_shape" {
  default = "VM.Standard2.1"
}

variable "db_admin_password" {
  default = "BEstrO0ng_#12"
}

variable "hostname" {
  default = "myoracledb"
}

variable "data_storage_size_in_gb" {
  default = "256"
}

variable "license_model" {
  default = "LICENSE_INCLUDED"
}

variable "node_count" {
  default = "2"
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
  ad_number      = 3
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

data "oci_database_db_versions" "test_db_versions_by_db_system_id" {
  compartment_id = var.compartment_ocid
  db_system_id   = oci_database_db_system.test_db_system.id
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

resource "oci_database_db_system" "test_db_system" {
  source              = "DB_SYSTEM"
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  source_db_system_id = var.source_db_system_id

  db_home {
    database {
      admin_password = var.db_admin_password

      db_backup_config {
        auto_backup_enabled = false
      }
    }
  }

  db_system_options {
    storage_management = "LVM"
  }

  shape                   = var.db_system_shape
  subnet_id               = var.subnet_id
  ssh_public_keys         = [var.ssh_public_key]
  display_name            = "MyTFDBSystemVM"
  hostname                = var.hostname
  data_storage_size_in_gb = var.data_storage_size_in_gb
  license_model           = var.license_model
  node_count              = data.oci_database_db_system_shapes.test_db_system_shapes.db_system_shapes[0]["minimum_node_count"]

  #To use defined_tags, set the values below to an existing tag namespace, refer to the identity example on how to create tag namespaces
  #  defined_tags = {
  #   "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"
  #  }
  freeform_tags = {
    "Department" = "Finance"
  }
}

