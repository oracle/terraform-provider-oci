variable "bds_instance_id" {}
variable "compartment_id" {}
variable "certificate_authority_id" {}
variable "display_name" {}
variable "certificate_type" {}
variable "cluster_admin_password" {} // Pass one of cluster_admin_password or secret_id
variable "secret_id" {}
variable "issue_certificate_trigger" {}
variable "renew_certificate_trigger" {}
variable "set_default_trigger" {}
variable "is_missing_nodes_only" {
  type    = bool
}

resource "oci_bds_bds_instance_bds_certificate_configuration" "oci_cert_config" {
  bds_instance_id          = var.bds_instance_id
  display_name             = var.display_name
  certificate_type         = var.certificate_type
  certificate_authority_id = var.certificate_authority_id
  compartment_id           = var.compartment_id

  cluster_admin_password   = var.cluster_admin_password != "" ? var.cluster_admin_password : null
  secret_id                = var.secret_id != "" ? var.secret_id : null

  issue_certificate_trigger = var.issue_certificate_trigger
  renew_certificate_trigger = var.renew_certificate_trigger
  set_default_trigger       = var.set_default_trigger
  is_missing_nodes_only     = var.is_missing_nodes_only
  timeouts {
    create = "30m"
    update = "45m"
    delete = "30m"
  }
}

data "oci_bds_bds_instance_bds_certificate_configuration" "oci_cert_config" {
  bds_instance_id                  = var.bds_instance_id
  bds_certificate_configuration_id = split("/", oci_bds_bds_instance_bds_certificate_configuration.oci_cert_config.id)[3]
}

output "bds_certificate_configuration_id" {
  value = split("/", oci_bds_bds_instance_bds_certificate_configuration.oci_cert_config.id)[3]
}

output "bds_certificate_configuration_composite_id" {
  value = oci_bds_bds_instance_bds_certificate_configuration.oci_cert_config.id
}

output "issue_certificate_trigger" {
  value = oci_bds_bds_instance_bds_certificate_configuration.oci_cert_config.issue_certificate_trigger
}

output "is_missing_nodes_only_input" {
  value = oci_bds_bds_instance_bds_certificate_configuration.oci_cert_config.is_missing_nodes_only
}

output "is_default_configuration" {
  value = data.oci_bds_bds_instance_bds_certificate_configuration.oci_cert_config.is_default_configuration
}

output "state" {
  value = data.oci_bds_bds_instance_bds_certificate_configuration.oci_cert_config.state
}

output "time_last_refreshed_or_generated" {
  value = data.oci_bds_bds_instance_bds_certificate_configuration.oci_cert_config.time_last_refreshed_or_generated
}