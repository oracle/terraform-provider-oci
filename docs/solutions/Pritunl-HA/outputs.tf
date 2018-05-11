# Output the private and public IPs of the instance

output "MongoPrimaryPrivateIP" {
value = ["${data.oci_core_vnic.MPVNIC.private_ip_address}"]
}

output "MongoPrimaryPublicIP" {
value = ["${data.oci_core_vnic.MPVNIC.public_ip_address}"]
}

output "MongoReplica1PrivateIP" {
value = ["${data.oci_core_vnic.MR1VNIC.private_ip_address}"]
}

output "MongoReplica1PublicIP" {
value = ["${data.oci_core_vnic.MR1VNIC.public_ip_address}"]
}

output "MongoReplica2PrivateIP" {
value = ["${data.oci_core_vnic.MR2VNIC.private_ip_address}"]
}

output "MongoReplica2PublicIP" {
value = ["${data.oci_core_vnic.MR2VNIC.public_ip_address}"]
}


output "Pritunl1PrivateIP" {
value = ["${data.oci_core_vnic.P1VNIC.private_ip_address}"]
}

output "Pritunl1PublicIP" {
value = ["${data.oci_core_vnic.P1VNIC.public_ip_address}"]
}

output "Pritunl2PrivateIP" {
value = ["${data.oci_core_vnic.P2VNIC.private_ip_address}"]
}

output "Pritunl2PublicIP" {
value = ["${data.oci_core_vnic.P2VNIC.public_ip_address}"]
}

output "PritunllinkPrivateIP" {
value = ["${data.oci_core_vnic.PTLVNIC.private_ip_address}"]
}

output "PritunllinkPublicIP" {
value = ["${data.oci_core_vnic.PTLVNIC.public_ip_address}"]
}

