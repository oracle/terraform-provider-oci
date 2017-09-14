##Output a map of compartment names mapped to OCID

data "oci_identity_compartments" "compartmentList" {
  compartment_id = "${var.tenancy_ocid}"
}

resource "null_resource" "compartments" {
  count = "${length(data.oci_identity_compartments.compartmentList.compartments)}"
  triggers = {
    names = "${lookup(data.oci_identity_compartments.compartmentList.compartments[count.index],"name")}"
    ids = "${lookup(data.oci_identity_compartments.compartmentList.compartments[count.index],"id")}"
  }
}

output "compartments" {
  value = "${zipmap(null_resource.compartments.*.triggers.names,null_resource.compartments.*.triggers.ids)}"
}

##Output a map of image names mapped to OCID

data "oci_core_images" "imageList" {
  compartment_id = "${var.tenancy_ocid}"
}

resource "null_resource" "images" {
  count = "${length(data.oci_core_images.imageList.images)}"
  triggers = {
    names = "${lookup(data.oci_core_images.imageList.images[count.index],"display_name")}"
    ids = "${lookup(data.oci_core_images.imageList.images[count.index],"id")}"
  }
}

output "images" {
  value = "${zipmap(null_resource.images.*.triggers.names,null_resource.images.*.triggers.ids)}"
}



#Output a map of ADs
data "oci_identity_availability_domains" "ads" {
  compartment_id = "${var.tenancy_ocid}"
}

output "ads" {
  value = "${data.oci_identity_availability_domains.ads.availability_domains}"
}

