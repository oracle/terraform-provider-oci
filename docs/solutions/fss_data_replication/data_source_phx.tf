# Gets a list of Availability Domains
data "oci_identity_availability_domains" "ads_phx" {
  provider       = "oci.phx"
  compartment_id = "${var.tenancy_ocid}"
}

# Gets the list of file systems in the compartment
data "oci_file_storage_file_systems" "file_systems_phx" {
  provider            = "oci.phx"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ads_phx.availability_domains[0],"name")}"
  compartment_id      = "${var.compartment_id}"
  state               = "ACTIVE"
}

# Gets the list of mount targets in the compartment
data "oci_file_storage_mount_targets" "mount_targets_phx" {
  provider            = "oci.phx"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ads_phx.availability_domains[0],"name")}"
  compartment_id      = "${var.compartment_id}"
  state               = "ACTIVE"
}

# Gets the list of exports in the compartment
data "oci_file_storage_exports" "export_paths_phx" {
  provider       = "oci.phx"
  compartment_id = "${var.compartment_id}"
  state          = "ACTIVE"
}

data "oci_core_private_ips" "dst_mt_private_ip_phx_ad1" {
  provider  = "oci.phx"
  subnet_id = "${oci_core_subnet.subnet_phx_ad1.id}"

  filter {
    name   = "id"
    values = ["${oci_file_storage_mount_target.dst_mt_phx_ad1.private_ip_ids.0}"]
  }
}
