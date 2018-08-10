locals {
  fss_client_instance_src_public_ip = "${oci_core_instance.fss_client_instance_src.public_ip}"
}

locals {
  fss_client_instance_dst_public_ip = "${oci_core_instance.fss_client_instance_dst.public_ip}"
}

locals {
  fss_client_instance_src_private_ip = "${oci_core_instance.fss_client_instance_src.private_ip}"
}

locals {
  fss_client_instance_dst_private_ip = "${oci_core_instance.fss_client_instance_dst.private_ip}"
}
