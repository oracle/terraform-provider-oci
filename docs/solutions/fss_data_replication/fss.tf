#The File Storage Service provides a durable, enterprise_grade network file system 
#that you can connect to from any bare metal, virtual machine, or container instance 
#in your virtual cloud network.
resource "oci_file_storage_file_system" "src_fs_iad_ad1" {
  provider            = "oci.iad"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ads_iad.availability_domains[0],"name")}"
  compartment_id      = "${var.compartment_id}"

  display_name = "${var.src_file_system}"
}

resource "oci_file_storage_file_system" "dst_fs_iad_ad2" {
  provider            = "oci.iad"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ads_iad.availability_domains[1],"name")}"
  compartment_id      = "${var.compartment_id}"

  display_name = "${var.dst_file_system}"
}

resource "oci_file_storage_file_system" "dst_fs_phx_ad1" {
  provider            = "oci.phx"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ads_phx.availability_domains[0],"name")}"
  compartment_id      = "${var.compartment_id}"

  display_name = "${var.dst_file_system}"
}

#An NFS endpoint that lives in a subnet of your choice and is highly available. 
#It provides the IP address or DNS name that is used in the mount command when 
#connecting NFS clients to File Storage Service. 
resource "oci_file_storage_mount_target" "src_mt_iad_ad1" {
  provider            = "oci.iad"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ads_iad.availability_domains[0],"name")}"
  compartment_id      = "${var.compartment_id}"
  subnet_id           = "${oci_core_subnet.subnet_iad_ad1.id}"

  display_name = "${var.src_mount_target}"
}

resource "oci_file_storage_mount_target" "dst_mt_iad_ad2" {
  provider            = "oci.iad"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ads_iad.availability_domains[1],"name")}"
  compartment_id      = "${var.compartment_id}"
  subnet_id           = "${oci_core_subnet.subnet_iad_ad2.id}"

  display_name = "${var.dst_mount_target}"
}

resource "oci_file_storage_mount_target" "dst_mt_phx_ad1" {
  provider            = "oci.phx"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ads_phx.availability_domains[0],"name")}"
  compartment_id      = "${var.compartment_id}"
  subnet_id           = "${oci_core_subnet.subnet_phx_ad1.id}"

  display_name = "${var.dst_mount_target}"
}

#A path that is specified when a file system is associated with a mount target. 
#It uniquely identifies the file system within the mount target, 
#letting you associate up to 100 file systems to a single mount target
resource "oci_file_storage_export" "src_export_path_iad_ad1" {
  provider       = "oci.iad"
  export_set_id  = "${oci_file_storage_mount_target.src_mt_iad_ad1.export_set_id}"
  file_system_id = "${oci_file_storage_file_system.src_fs_iad_ad1.id}"
  path           = "${var.src_export_path}"
}

resource "oci_file_storage_export" "dst_export_path_iad_ad2" {
  provider       = "oci.iad"
  export_set_id  = "${oci_file_storage_mount_target.dst_mt_iad_ad2.export_set_id}"
  file_system_id = "${oci_file_storage_file_system.dst_fs_iad_ad2.id}"
  path           = "${var.dst_export_path}"
}

resource "oci_file_storage_export" "dst_export_path_phx_ad1" {
  provider       = "oci.phx"
  export_set_id  = "${oci_file_storage_mount_target.dst_mt_phx_ad1.export_set_id}"
  file_system_id = "${oci_file_storage_file_system.dst_fs_phx_ad1.id}"
  path           = "${var.dst_export_path}"
}
