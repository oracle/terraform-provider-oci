output "private_ip" {
  value = "${oci_core_instance.instance.*.private_ip}"
}

output "public_ip" {
  value = "${oci_core_instance.instance.*.public_ip}"
}
