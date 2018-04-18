# Gets a list of Availability Domains
data "oci_identity_availability_domains" "ADs" {
    compartment_id = "${var.tenancy_ocid}"
}

# Gets the custom image that will be created by this Terraform config
data "oci_core_images" "TFCustomImage" {
    compartment_id = "${var.tenancy_ocid}"

    filter {
        name = "id"
        values = ["${oci_core_image.TFCustomImage.id}"]
    }
}

# Gets a list of all Oracle Linux 7.4 images that support a given Instance shape
data "oci_core_images" "TFSupportedShapeImages" {
    compartment_id = "${var.tenancy_ocid}"
    shape = "${var.InstanceShape}"
    operating_system = "${var.ImageOS}"
    operating_system_version = "${var.ImageOSVersion}"
}
