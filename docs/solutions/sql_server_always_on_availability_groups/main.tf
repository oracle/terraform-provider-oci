module "vcn" {
  source           = "./modules/network/vcn"
  compartment_ocid = "${module.compartment.id}"
  label_prefix     = "${var.label_prefix}"
  vcn_dns_name     = "${var.vcn_dns_name}"
  vcn_cidr_block   = "${var.vcn_cidr_block}"
  local_dns_server = "${var.local_dns_server}"
}

module "compartment" {
  source                  = "./modules/iam/compartment"
  compartment_name        = "${var.compartment_name}"
  compartment_description = "${var.compartment_description}"
}

module "securitylist" {
  source         = "./modules/network/securitylist"
  compartment_ocid = "${module.compartment.id}"
  vcn_id         = "${module.vcn.vcn_id}"
  DMZ_prefix     = "${local.DMZ_prefix}"
  ADMIN_prefix   = "${local.ADMIN_prefix}"
  SQL_prefix     = "${local.SQL_prefix}"
}

module "dmz_subnets" {
  dns_label  = "DMZ"
  cidr_block = "${local.DMZ_subnets_cidrs}"
  source           = "./modules/network/subnet"
  compartment_ocid = "${module.compartment.id}"
  vcn_id           = "${module.vcn.vcn_id}"
  ad_count         = "${var.ad_count}"
  route_table_id   = "${module.vcn.rt_id}"
  dhcp_options_id  = "${module.vcn.internet_dhcp_options_id}"
  security_list_id = ["${module.securitylist.dmz_id}"]
  label_prefix     = "${var.label_prefix}"
  tenancy_ocid     = "${var.tenancy_ocid}"
}

module "admin_subnets" {
  dns_label  = "ADMIN"
  cidr_block = "${local.ADMIN_subnets_cidrs}"
  source           = "./modules/network/subnet"
  compartment_ocid = "${module.compartment.id}"
  vcn_id           = "${module.vcn.vcn_id}"
  ad_count         = "${var.ad_count}"
  route_table_id   = "${module.vcn.rt_id}"
  dhcp_options_id  = "${module.vcn.internet_dhcp_options_id}"
  security_list_id = ["${module.securitylist.admin_id}"]
  label_prefix     = "${var.label_prefix}"
  tenancy_ocid     = "${var.tenancy_ocid}"
  private          = "true"
}

module "sql_subnets" {
  dns_label  = "SQL"
  cidr_block = "${local.SQL_subnets_cidrs}"
  source           = "./modules/network/subnet"
  compartment_ocid = "${module.compartment.id}"
  vcn_id           = "${module.vcn.vcn_id}"
  ad_count         = "${var.ad_count}"
  route_table_id   = "${module.vcn.rt_id}"
  dhcp_options_id  = "${module.vcn.internet_dhcp_options_id}"
  security_list_id = ["${module.securitylist.sql_id}"]
  label_prefix     = "${var.label_prefix}"
  tenancy_ocid     = "${var.tenancy_ocid}"
  private          = "true"
}

module "witness_subnets" {
  dns_label  = "Witness"
  cidr_block = "${local.Witness_subnets_cidrs}"
  source           = "./modules/network/subnet"
  compartment_ocid = "${module.compartment.id}"
  vcn_id           = "${module.vcn.vcn_id}"
  ad_count         = "1"
  route_table_id   = "${module.vcn.rt_id}"
  dhcp_options_id  = "${module.vcn.internet_dhcp_options_id}"
  security_list_id = ["${module.securitylist.sql_id}"]
  label_prefix     = "${var.label_prefix}"
  tenancy_ocid     = "${var.tenancy_ocid}"
  private          = "true"
  ad_deployment    = "${var.witness_deployment}"
}

module "volumes" {
  source             = "./modules/storage/volume"
  compartment_ocid   = "${module.compartment.id}"
  ad_count           = "${var.ad_count}"
  label_prefix       = "${var.label_prefix}"
  tenancy_ocid       = "${var.tenancy_ocid}"
  ad_deployment      = "${var.witness_deployment}"
  sql_db_size        = "${var.sql_db_size}"
  sql_log_size       = "${var.sql_log_size}"
  sql_backup_size    = "${var.sql_backup_size}"
  witness_volume_size = "${var.witness_volume_size}"
  ad_deployment      = "${var.witness_deployment}"
}

module "dmz_hosts" {
  dns_label = "Bastion"
  subnets   = "${module.dmz_subnets.subnet_id}"
  source         = "./modules/instances/bastion/"
  compartment_ocid = "${module.compartment.id}"
  vcn_id         = "${module.vcn.vcn_id}"
  ad_count       = "${var.ad_count}"
  label_prefix   = "${var.label_prefix}"
  tenancy_ocid   = "${var.tenancy_ocid}"
  image_id       = "${var.image_id[var.region]}"
  shape          = "${var.dmz_shape}"
}

module "admin_hosts" {
  dns_label = "DC"
  subnets   = "${module.admin_subnets.subnet_id}"
  source         = "./modules/instances/active_directory/"
  compartment_ocid = "${module.compartment.id}"
  vcn_id         = "${module.vcn.vcn_id}"
  ad_count       = "${var.ad_count}"
  label_prefix   = "${var.label_prefix}"
  tenancy_ocid   = "${var.tenancy_ocid}"
  image_id       = "${var.image_id[var.region]}"
  shape          = "${var.admin_shape}"
}

module "sql_hosts" {
  dns_label = "SQL"
  ad_count  = "${var.ad_count}"
  subnets   = "${module.sql_subnets.subnet_id}"
  source         = "./modules/instances/sql/"
  compartment_ocid = "${module.compartment.id}"
  vcn_id         = "${module.vcn.vcn_id}"
  label_prefix   = "${var.label_prefix}"
  tenancy_ocid   = "${var.tenancy_ocid}"
  image_id       = "${var.image_id[var.region]}"
  shape          = "${var.sql_shape}"
  db_volumes     = "${module.volumes.sql_db_id}"
  log_volumes    = "${module.volumes.sql_log_id}"
  backup_volumes = "${module.volumes.sql_backup_id}"
}

module "witness_hosts" {
  dns_label = "WITNESS"
  subnets   = "${module.witness_subnets.subnet_id}"
  source          = "./modules/instances/witness/"
  compartment_ocid  = "${module.compartment.id}"
  vcn_id          = "${module.vcn.vcn_id}"
  ad_count        = "${var.ad_count}"
  label_prefix    = "${var.label_prefix}"
  tenancy_ocid    = "${var.tenancy_ocid}"
  image_id        = "${var.image_id[var.region]}"
  shape           = "${var.witness_shape}"
  witness_volumes = "${module.volumes.witness_id}"
  ad_deployment   = "${var.witness_deployment}"
}

module "secondaryIPs" {
  dns_label = "SQL"
  subnets   = "${module.sql_subnets.subnet_id}"
  source         = "./modules/network/secondaryip/"
  compartment_ocid = "${module.compartment.id}"
  vcn_id         = "${module.vcn.vcn_id}"
  ad_count       = "${var.ad_count}"
  label_prefix   = "${var.label_prefix}"
  tenancy_ocid   = "${var.tenancy_ocid}"
  vnic_ids       = "${module.sql_hosts.vnic_ids}"
}
