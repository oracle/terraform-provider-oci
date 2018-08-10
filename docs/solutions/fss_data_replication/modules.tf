//File System Data replication across AD1 and AD2
module "rsync_iad_ad1_to_ad2" {
  source = "./modules/rsync_filesystem_local"

  ssh_private_key_path = "${var.ssh_private_key_path}"
  ssh_public_key_path  = "${var.ssh_public_key_path}"

  availability_domain = "${lookup(data.oci_identity_availability_domains.ads_iad.availability_domains[0],"name")}"
  compartment_id      = "${var.compartment_id}"
  subnet_id           = "${oci_core_subnet.subnet_iad_ad1.id}"
  instance_hostname   = "fss-ad1-2"
  instance_shape      = "VM.Standard2.4"
  instance_image_id   = "${var.instance_image_id[var.region_iad]}"

  src_export_path             = "${var.src_export_path}"
  src_mount_target_private_ip = "${local.src_mt_private_ip_iad_ad1}"

  dst_export_path             = "${var.dst_export_path}"
  dst_mount_target_private_ip = "${local.dst_mt_private_ip_iad_ad2}"

  //data sync every 30'. 
  data_sync_frequency = "*/30 * * * *"
}

output "rsync_fs_instance_public_ip_iad_ad1" {
  value = "${module.rsync_iad_ad1_to_ad2.data_sync_public_ip}"
}

output "rsync_fs_instance_private_ip_iad_ad1" {
  value = "${module.rsync_iad_ad1_to_ad2.data_sync_private_ip}"
}

//data replication for snapshot across AD1 and AD2
module "rsync_snapshot_iad_ad1_to_ad2" {
  source = "./modules/rsync_snapshot_local"

  ssh_private_key_path        = "${var.ssh_private_key_path}"
  ssh_public_key_path         = "${var.ssh_public_key_path}"
  availability_domain         = "${lookup(data.oci_identity_availability_domains.ads_iad.availability_domains[0],"name")}"
  compartment_id              = "${var.compartment_id}"
  subnet_id                   = "${oci_core_subnet.subnet_iad_ad1.id}"
  instance_hostname           = "fss-snap-ad1-2"
  instance_shape              = "VM.Standard2.8"
  instance_image_id           = "${var.instance_image_id[var.region_iad]}"
  data_sync_frequency         = "*/5 * * * *"
  snapshot_frequency          = "*/10 * * * *"
  src_export_path             = "${var.src_export_path}"
  src_mount_target_private_ip = "${local.src_mt_private_ip_iad_ad1}"
  dst_export_path             = "${var.dst_export_path}"
  dst_mount_target_private_ip = "${local.dst_mt_private_ip_iad_ad2}"
}

output "rsync_snapshot_instance_public_ip_iad_ad1" {
  value = "${module.rsync_snapshot_iad_ad1_to_ad2.data_sync_public_ip}"
}

output "rsync_snapshot_instance_private_ip_iad_ad1" {
  value = "${module.rsync_snapshot_iad_ad1_to_ad2.data_sync_private_ip}"
}

//data replication for snapshot across regions
module "rsync_snapshot_across_ashburn_phoenix" {
  source = "./modules/rsync_snapshot_across_region"

  providers = {
    "oci.src" = "oci.iad"
    "oci.dst" = "oci.phx"
  }

  ssh_private_key_path = "${var.ssh_private_key_path}"
  ssh_public_key_path  = "${var.ssh_public_key_path}"

  compartment_id = "${var.compartment_id}"

  //Host Configuration
  src_availability_domain = "${lookup(data.oci_identity_availability_domains.ads_iad.availability_domains[0],"name")}"
  src_subnet_id           = "${oci_core_subnet.subnet_iad_ad1.id}"
  src_instance_hostname   = "rsync-fss-iad"
  src_instance_shape      = "VM.Standard2.16"
  src_instance_image_id   = "${var.instance_image_id[var.region_iad]}"

  dst_availability_domain = "${lookup(data.oci_identity_availability_domains.ads_phx.availability_domains[0],"name")}"
  dst_subnet_id           = "${oci_core_subnet.subnet_phx_ad1.id}"
  dst_instance_hostname   = "rsync-fss-phx"
  dst_instance_shape      = "VM.Standard2.16"
  dst_instance_image_id   = "${var.instance_image_id[var.region_phx]}"

  //FSS configuration
  src_export_path             = "${var.src_export_path}"
  src_mount_target_private_ip = "${local.src_mt_private_ip_iad_ad1}"

  dst_export_path             = "${var.dst_export_path}"
  dst_mount_target_private_ip = "${local.dst_mt_private_ip_phx_ad1}"

  data_sync_frequency = "*/30 * * * *"
  snapshot_frequency  = "@hourly"
}

output "rsync_snapshot_instance_src_public_ip" {
  value = "${module.rsync_snapshot_across_ashburn_phoenix.data_sync_src_public_ip}"
}

output "rsync_snapshot_instance_src_private_ip" {
  value = "${module.rsync_snapshot_across_ashburn_phoenix.data_sync_src_private_ip}"
}

output "rsync_snapshot_instance_dst_public_ip" {
  value = "${module.rsync_snapshot_across_ashburn_phoenix.data_sync_dst_public_ip}"
}

output "rsync_snapshot_instance_dst_private_ip" {
  value = "${module.rsync_snapshot_across_ashburn_phoenix.data_sync_dst_private_ip}"
}
