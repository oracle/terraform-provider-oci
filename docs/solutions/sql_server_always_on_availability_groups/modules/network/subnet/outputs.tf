output "subnet_id" {
  value = ["${oci_core_subnet.subnet.*.id}"]
}
