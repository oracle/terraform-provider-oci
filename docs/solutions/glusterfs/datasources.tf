data "oci_identity_availability_domains" "ADs" {
  compartment_id = "${var.compartment_ocid}"
}

# Gets the OCID of the image. This technique is for example purposes only. The results of oci_core_images may
# change over time for Oracle-provided images, so the only sure way to get the correct OCID is to supply it directly.
data "oci_core_images" "ServerImageList" {
  compartment_id = "${var.compartment_ocid}"
  display_name = "${var.ServerInstanceImage}"
}

data "oci_core_images" "ClientImageList" {
  compartment_id = "${var.compartment_ocid}"
  display_name = "${var.ClientInstanceImage}"
}