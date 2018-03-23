output "id" {
  value = "${oci_core_virtual_network.CoreVCN.id}"
}

output "subnet_ad1_id" {
  value = "${oci_core_subnet.MgmtSubnetAD1.id}"
}

output "subnet_ad2_id" {
  value = "${oci_core_subnet.MgmtSubnetAD2.id}"
}
output "subnet_ad3_id" {
  value = "${oci_core_subnet.MgmtSubnetAD3.id}"
}

output "subnets" {
  value = ["${oci_core_subnet.MgmtSubnetAD1.id}", "${oci_core_subnet.MgmtSubnetAD1.id}", "${oci_core_subnet.MgmtSubnetAD1.id}"]
}

