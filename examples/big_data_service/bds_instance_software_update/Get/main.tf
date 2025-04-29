variable "bds_instance_id" {}
variable "software_update_key" {}

data "oci_bds_bds_instance_software_update" "test_bds_instance_software_update" {
    bds_instance_id = var.bds_instance_id}
    software_update_key = var.software_update_key
}

output "get_software_keys" {
    value = data.oci_bds_bds_instance_software_update.test_bds_instance_software_update
}