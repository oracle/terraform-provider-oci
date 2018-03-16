output "ip2" {
  value = ["${oci_core_private_ip.ip2.*.ip_address}"]
}

output "ip3" {
  value = ["${oci_core_private_ip.ip3.*.ip_address}"]
}
