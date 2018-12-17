# Gets a list of Availability Domains
data "oci_identity_availability_domains" "ADs" {
  compartment_id = "${var.tenancy_ocid}"
}

# Gets a list of all Oracle Linux 7.5 images that support a given Instance shape
data "oci_core_images" "TFSupportedShapeImages" {
  compartment_id           = "${var.tenancy_ocid}"
  shape                    = "${var.instance_shape}"
  operating_system         = "${var.ImageOS}"
  operating_system_version = "${var.ImageOSVersion}"

  filter {
    name   = "launch_options.is_pv_encryption_in_transit_enabled"
    values = ["true"]
  }
}
