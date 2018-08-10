# Gets a list of Availability Domains
data "oci_identity_availability_domains" "ads_iad" {
  provider       = "oci.iad"
  compartment_id = "${var.tenancy_ocid}"
}

# Gets the list of file systems in the compartment
data "oci_file_storage_file_systems" "file_systems_iad" {
  provider            = "oci.iad"
  count               = "2"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ads_iad.availability_domains[count.index],"name")}"
  compartment_id      = "${var.compartment_id}"
  state               = "ACTIVE"
}

# Gets the list of mount targets in the compartment
data "oci_file_storage_mount_targets" "mount_targets_iad" {
  provider            = "oci.iad"
  count               = "2"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ads_iad.availability_domains[count.index],"name")}"
  compartment_id      = "${var.compartment_id}"
  state               = "ACTIVE"
}

# Gets the list of exports in the compartment
data "oci_file_storage_exports" "export_paths_iad" {
  provider       = "oci.iad"
  compartment_id = "${var.compartment_id}"
  state          = "ACTIVE"
}

data "oci_core_private_ips" "src_mt_private_ip_iad_ad1" {
  provider  = "oci.iad"
  subnet_id = "${oci_core_subnet.subnet_iad_ad1.id}"

  filter {
    name   = "id"
    values = ["${oci_file_storage_mount_target.src_mt_iad_ad1.private_ip_ids.0}"]
  }
}

data "oci_core_private_ips" "dst_mt_private_ip_iad_ad2" {
  provider  = "oci.iad"
  subnet_id = "${oci_core_subnet.subnet_iad_ad2.id}"

  filter {
    name   = "id"
    values = ["${oci_file_storage_mount_target.dst_mt_iad_ad2.private_ip_ids.0}"]
  }
}
