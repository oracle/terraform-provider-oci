
variable "tenancy_ocid" {}
variable "ssh_public_key" {}
variable "region" {}
variable "compartment_id" {}
variable "bds_instance_id" {}
variable "software_update_key" {}

resource "oci_bds_bds_instance_software_update_action" "test_bds_instance_software_update_action" {
    bds_instance_id = var.bds_instance_id
    software_update_key = var.software_update_key
}
