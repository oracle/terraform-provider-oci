# Gets a list of Availability Domains
data "oci_identity_availability_domains" "ADs" {
  compartment_id = "${var.tenancy_ocid}"
}

# Gets the custom image that will be created by this Terraform config
data "oci_core_images" "TFCustomImage" {
  compartment_id = "${var.compartment_ocid}"

  filter {
    name   = "id"
    values = ["${oci_core_image.TFCustomImage.id}"]
  }
}

# Gets a list of images within a tenancy
data "oci_core_images" "TFSupportedShapeImages" {
  compartment_id = "${var.tenancy_ocid}"

  # Uncomment below to filter images that support a specific instance shape 
  #shape                    = "VM.Standard2.1"

  # Uncomment below to filter images that are a specific OS
  #operating_system         = "Oracle Linux"

  # Uncomment below to filter images that are a specific OS version 
  #operating_system_version = "7.5"
}
