
output "DnsServer1" {
    value = ["${data.oci_core_vnic.DnsVMVnic.private_ip_address}"]
}

output "DnsServer2" {
    value = ["${data.oci_core_vnic.DnsVMVnic2.private_ip_address}"]
}

output "lb1_private_ip" {
  value = ["${oci_load_balancer.lb1.ip_addresses}"]
}

output "lb2_private_ip" {
  value = ["${oci_load_balancer.lb2.ip_addresses}"]
}

