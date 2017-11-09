// vcn_dns_name is the outermost prefix in the <vcn_dns_name>.oraclevcn.com domain.
variable "vcn_dns_name" {
  default = "SQLAlwaysOn"
}

// vcn_cidr_block defines the IP address pool for the network.
// Check the locals configuration below for subnet declarations
variable "vcn_cidr_block" {
  default = "10.0.0.0/19"
}

// Compartment Name. This will only be used if the compartment_id variable in the terraform.tfvars file is not set.
variable "compartment_name" {
  default = "SQLAlwaysOn"
}

// Compartment Description. This will only be used if the compartment_id variable in the terraform.tfvars file is not set.
variable "compartment_description" {
  default = "SQL Always On compartment"
}

// How many availability domains are going to be used. Minimum 2, Maximum 3.
variable "ad_count" {
  default = "3"
}

variable "local_dns_server" {
  default = "1.1.1.1"
}

// In which AD deploy the Witness servers
variable "witness_deployment" {
  default = "2"
}

// Windows 2012
variable "image_id" {
  default = "ocid1.image.oc1.iad.aaaaaaaajlfsi5npxguvhad3v5d5lu7dc3zcylr2csfdexgd6kor3f6zeqeq"
}

// Size of the volumes in GB
variable "sql_db_size" {
  default = "2048"
}

variable "sql_backup_size" {
  default = "256"
}

variable "sql_log_size" {
  default = "512"
}

variable "witness_block_size" {
  default = "256"
}

// Hosts shapes definition.
variable "dmz_shape" {
  default = "VM.Standard1.2"
}

variable "sql_shape" {
  default = "VM.Standard1.8"
}

variable "admin_shape" {
  default = "VM.Standard1.4"
}

variable "witness_shape" {
  default = "VM.Standard1.4"
}

locals {
  // Subnet prefix declaration
  // If not set otherwise, subnets will be auto generated.
  // The first /21 block will be assigned to the DMZ-prefix.  
  // Second block to the Active Directory.  
  // Third for the SQL servers. Defaults are declared in the include.tf file.

  // CIDR /21 subnets assigned for the roles  
    // Example override:  
    // DMZ_prefix = "192.168.0.0/21"

  DMZ_prefix   = "${cidrsubnet("${var.vcn_cidr_block}", 2, 0)}"
  ADMIN_prefix = "${cidrsubnet("${var.vcn_cidr_block}", 2, 1)}"
  SQL_prefix   = "${cidrsubnet("${var.vcn_cidr_block}", 2, 2)}"

  // LISTS of /23 subnets assigned for roles inside the availability domain.
  // They can be overriden using list definition e.g
  // DMZ_subnets_cidrs = ["192.168.0.0/23","192.168.2.0/23","192.168.4.0/23"]

  DMZ_subnets_cidrs = [
    "${cidrsubnet("${local.DMZ_prefix}", 2, 0)}",
    "${cidrsubnet("${local.DMZ_prefix}", 2, 1)}",
    "${cidrsubnet("${local.DMZ_prefix}", 2, 2)}",
  ]
  ADMIN_subnets_cidrs = [
    "${cidrsubnet("${local.ADMIN_prefix}", 2, 0)}",
    "${cidrsubnet("${local.ADMIN_prefix}", 2, 1)}",
    "${cidrsubnet("${local.ADMIN_prefix}", 2, 2)}",
  ]
  SQL_subnets_cidrs = [
    "${cidrsubnet("${local.SQL_prefix}", 2, 0)}",
    "${cidrsubnet("${local.SQL_prefix}", 2, 1)}",
    "${cidrsubnet("${local.SQL_prefix}", 2, 2)}",
  ]
  Witness_subnets_cidrs = [
    "${cidrsubnet("${local.SQL_prefix}", 2, 3)}",
  ]
}
