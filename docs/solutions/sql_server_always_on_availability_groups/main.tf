// VCN declaration
// Values are taken from the configuration.tf file

module "vcn" {
  source           = "modules/network/vcn"
  compartment_id   = "${module.compartment.id}"
  label_prefix     = "${var.label_prefix}"
  vcn_dns_name     = "${var.vcn_dns_name}"
  vcn_cidr_block   = "${var.vcn_cidr_block}"
  local_dns_server = "${var.local_dns_server}"
}

module "compartment" {
  // If not specified in tfvars it will be created
  source                  = "modules/iam/compartment"
  compartment_id          = "${var.compartment_id}"
  compartment_name        = "${var.compartment_name}"
  compartment_description = "${var.compartment_description}"
}

// Security lists configuration

module "securitylist" {
  source         = "modules/network/securitylist"
  compartment_id = "${module.compartment.id}"
  vcn_id         = "${module.vcn.vcn_id}"
  DMZ_prefix     = "${local.DMZ_prefix}"
  ADMIN_prefix   = "${local.ADMIN_prefix}"
  SQL_prefix     = "${local.SQL_prefix}"
}

// Public subnet containing bastion hosts.

module "dmz_subnets" {
  // Can be set to custom value
  dns_label  = "DMZ"
  cidr_block = "${local.DMZ_subnets_cidrs}"

  // Constants
  source           = "modules/network/subnet"
  compartment_id   = "${module.compartment.id}"
  vcn_id           = "${module.vcn.vcn_id}"
  ad_count         = "${var.ad_count}"
  route_table_id   = "${module.vcn.rt_id}"
  dhcp_options_id  = "${module.vcn.internet_dhcp_options_id}"
  security_list_id = ["${module.securitylist.dmz_id}"]
  label_prefix     = "${var.label_prefix}"
  tenancy_id       = "${var.tenancy_id}"
}

// Private subnet. Active Directory servers.

module "admin_subnets" {
  // Can be set to custom value
  dns_label  = "ADMIN"
  cidr_block = "${local.ADMIN_subnets_cidrs}"

  // Constants
  source           = "modules/network/subnet"
  compartment_id   = "${module.compartment.id}"
  vcn_id           = "${module.vcn.vcn_id}"
  ad_count         = "${var.ad_count}"
  route_table_id   = "${module.vcn.rt_id}"
  dhcp_options_id  = "${module.vcn.internet_dhcp_options_id}"
  security_list_id = ["${module.securitylist.admin_id}"]
  label_prefix     = "${var.label_prefix}"
  tenancy_id       = "${var.tenancy_id}"
  private          = "true"
}

// SQL servers subnet

module "sql_subnets" {
  // Can be set to custom value
  dns_label  = "SQL"
  cidr_block = "${local.SQL_subnets_cidrs}"

  // Constants
  source           = "modules/network/subnet"
  compartment_id   = "${module.compartment.id}"
  vcn_id           = "${module.vcn.vcn_id}"
  ad_count         = "${var.ad_count}"
  route_table_id   = "${module.vcn.rt_id}"
  dhcp_options_id  = "${module.vcn.internet_dhcp_options_id}"
  security_list_id = ["${module.securitylist.sql_id}"]
  label_prefix     = "${var.label_prefix}"
  tenancy_id       = "${var.tenancy_id}"
  private          = "true"
}

// Witness server subnet. Variable witness_deployment from the configuration.tf
// defines in which availability domain server will be installed

module "witness_subnets" {
  // Can be set to custom value
  dns_label  = "Witness"
  cidr_block = "${local.Witness_subnets_cidrs}"

  // Constants
  source           = "modules/network/subnet"
  compartment_id   = "${module.compartment.id}"
  vcn_id           = "${module.vcn.vcn_id}"
  ad_count         = "1"
  route_table_id   = "${module.vcn.rt_id}"
  dhcp_options_id  = "${module.vcn.internet_dhcp_options_id}"
  security_list_id = ["${module.securitylist.sql_id}"]
  label_prefix     = "${var.label_prefix}"
  tenancy_id       = "${var.tenancy_id}"
  private          = "true"
  ad_deployment    = "${var.witness_deployment}"
}

// Definitions of block storage volumes
// Size is configurable from the configuration.tf

module "volumes" {
  // Constants
  source             = "modules/storage/volume"
  compartment_id     = "${module.compartment.id}"
  ad_count           = "${var.ad_count}"
  label_prefix       = "${var.label_prefix}"
  tenancy_id         = "${var.tenancy_id}"
  ad_deployment      = "${var.witness_deployment}"
  sql_db_size        = "${var.sql_db_size}"
  sql_log_size       = "${var.sql_log_size}"
  sql_backup_size    = "${var.sql_backup_size}"
  witness_block_size = "${var.witness_block_size}"
  ad_deployment      = "${var.witness_deployment}"
}

// Bastion Hosts declarations

module "dmz_hosts" {
  // Can be set to custom value
  dns_label = "Bastion"
  subnets   = "${module.dmz_subnets.subnet_id}"

  // Constants
  source         = "modules/instances/Bastion/"
  compartment_id = "${module.compartment.id}"
  vcn_id         = "${module.vcn.vcn_id}"
  ad_count       = "${var.ad_count}"
  label_prefix   = "${var.label_prefix}"
  tenancy_id     = "${var.tenancy_id}"
  image_id       = "${var.image_id}"
  shape          = "${var.dmz_shape}"
}

// Active Directory servers

module "admin_hosts" {
  // Can be set to custom value
  dns_label = "DC"
  subnets   = "${module.admin_subnets.subnet_id}"

  // Constants
  source         = "modules/instances/ActiveDirectory/"
  compartment_id = "${module.compartment.id}"
  vcn_id         = "${module.vcn.vcn_id}"
  ad_count       = "${var.ad_count}"
  label_prefix   = "${var.label_prefix}"
  tenancy_id     = "${var.tenancy_id}"
  image_id       = "${var.image_id}"
  shape          = "${var.admin_shape}"
}

// SQL servers

module "sql_hosts" {
  // Can be set to custom value
  ad_count  = "${var.ad_count}"
  dns_label = "SQL"
  subnets   = "${module.sql_subnets.subnet_id}"

  // Constants
  source         = "modules/instances/Sql/"
  compartment_id = "${module.compartment.id}"
  vcn_id         = "${module.vcn.vcn_id}"
  label_prefix   = "${var.label_prefix}"
  tenancy_id     = "${var.tenancy_id}"
  image_id       = "${var.image_id}"
  shape          = "${var.sql_shape}"
  db_volumes     = "${module.volumes.sql_db_id}"
  log_volumes    = "${module.volumes.sql_log_id}"
  backup_volumes = "${module.volumes.sql_backup_id}"
}

// Witness Host

module "witness_hosts" {
  // Can be set to custom value
  dns_label = "WITNESS"
  subnets   = "${module.witness_subnets.subnet_id}"

  // Constants
  source          = "modules/instances/witness/"
  compartment_id  = "${module.compartment.id}"
  vcn_id          = "${module.vcn.vcn_id}"
  ad_count        = "${var.ad_count}"
  label_prefix    = "${var.label_prefix}"
  tenancy_id      = "${var.tenancy_id}"
  image_id        = "${var.image_id}"
  shape           = "${var.witness_shape}"
  witness_volumes = "${module.volumes.witness_id}"
  ad_deployment   = "${var.witness_deployment}"
}

// Add additional IP addresses for SQL cluster.
// Retrieved vnic_ids from sql_hosts module are passed as a list
// to the module.

module "secondaryIPs" {
  // Can be set to custom value
  dns_label = "SQL"
  subnets   = "${module.sql_subnets.subnet_id}"

  // Constants
  source         = "modules/network/secondaryip/"
  compartment_id = "${module.compartment.id}"
  vcn_id         = "${module.vcn.vcn_id}"
  ad_count       = "${var.ad_count}"
  label_prefix   = "${var.label_prefix}"
  tenancy_id     = "${var.tenancy_id}"
  vnic_ids       = "${module.sql_hosts.vnic_ids}"
}
