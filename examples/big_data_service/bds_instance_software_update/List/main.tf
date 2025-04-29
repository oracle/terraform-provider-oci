variable "bds_instance_id" { }

data "oci_bds_bds_instance_software_updates" "test_bds_instance_software_updates" {
bds_instance_id = var.bds_instance_id
}

output "list_software_keys" {
    value = data.oci_bds_bds_instance_software_updates.test_bds_instance_software_updates
}
