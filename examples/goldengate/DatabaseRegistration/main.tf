variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "kms_vault_ocid" {}
variable "kms_key_ocid" {}


variable "database_registration_alias_name" {
  default = "aliasName"
}

variable "database_registration_connection_string" {
  default = "fqdndb.ggs.com:1521/orcl.us.oracle.com"
}

variable "database_registration_session_mode" {
  default = "DIRECT"
}

variable "database_registration_defined_tags_value" {
  default = "value"
}

variable "database_registration_description" {
  default = "description"
}

variable "database_registration_display_name" {
  default = "displayName"
}

variable "database_registration_fqdn" {
  default = "fqdndb.ggs.com"
}

variable "database_registration_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "database_registration_ip_address" {
  default = "10.0.0.10"
}



variable "database_registration_password" {
  default = "BEstrO0ng_#11"
}

variable "database_registration_state" {
  default = "ACTIVE"
}

variable "database_registration_username" {
  default = "username"
}

variable "database_registration_wallet" {
  default = "wallet"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_core_subnet" "test_subnet" {
	cidr_block = "10.0.0.0/24"
	compartment_id = var.compartment_ocid
	vcn_id = oci_core_vcn.test_vcn.id
}

resource "oci_core_vcn" "test_vcn" {
	cidr_block = "10.0.0.0/16"
	compartment_id = var.compartment_ocid
}

resource "oci_golden_gate_database_registration" "test_database_registration" {
  #Required
  alias_name     = var.database_registration_alias_name
  compartment_id = var.compartment_ocid
  display_name   = var.database_registration_display_name
  fqdn           = var.database_registration_fqdn
  password       = var.database_registration_password
  username       = var.database_registration_username

  #Optional
  connection_string     = var.database_registration_connection_string
  session_mode          = var.database_registration_session_mode
  database_id           = data.oci_database_databases.t.databases.0.id
  #defined_tags          = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.database_registration_defined_tags_value)
  description           = var.database_registration_description
  freeform_tags         = var.database_registration_freeform_tags
  ip_address            = var.database_registration_ip_address
  key_id                = var.kms_key_ocid
  secret_compartment_id = var.compartment_ocid
  subnet_id             = oci_core_subnet.test_subnet.id
  vault_id              = var.kms_vault_ocid
  wallet                = var.database_registration_wallet
}

data "oci_golden_gate_database_registrations" "test_database_registrations" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name = var.database_registration_display_name
  state        = var.database_registration_state
}

data "oci_database_db_systems" "t" {
	compartment_id = var.compartment_ocid			
}

data "oci_database_db_homes" "t" {
	compartment_id = var.compartment_ocid
	db_system_id = data.oci_database_db_systems.t.db_systems.0.id
}

data "oci_database_databases" "t" {
	compartment_id = var.compartment_ocid
	db_home_id = data.oci_database_db_homes.t.db_homes.0.id	
}

#data "oci_database_db_nodes" "t" {
#    #Required
#    compartment_id = var.compartment_ocid
#    #Optional
#    db_system_id = data.oci_database_db_systems.t.db_systems.0.id
#}

#data "oci_core_vnic" "t" {
    #Required
    #vnic_id = data.oci_database_db_nodes.t.vnic_id //believe this is null when using FAKEHOSTSERIAL header
#}
