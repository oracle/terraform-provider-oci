
output "BackendInstance1" {
    value = "${data.oci_core_vnic.Instance1Vnic.public_ip_address}"
}

output "AppFQDN" {
    value = "${var.ha_app_name}.${var.ha_app_domain}"
}

output "lb1_private_ip" {
  value = ["${oci_load_balancer.lb1.ip_addresses}"]
}

output "lb2_private_ip" {
  value = ["${oci_load_balancer.lb2.ip_addresses}"]
}

