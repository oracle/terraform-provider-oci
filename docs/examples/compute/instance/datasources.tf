# Gets a list of Availability Domains
data "oci_identity_availability_domains" "ADs" {
    compartment_id = "${var.tenancy_ocid}"
}

# Gets the OCID of the image. This technique is for example purposes only. The results of oci_core_images may
# change over time for Oracle-provided images, so the only sure way to get the correct OCID is to supply it directly.
data "oci_core_images" "OLImageOCID" {
    compartment_id = "${var.compartment_ocid}"
    display_name = "${var.InstanceImageDisplayName}"
}

# Gets a list of vNIC attachments on the instance
data "oci_core_vnic_attachments" "InstanceVnics" {
    compartment_id = "${var.compartment_ocid}"
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
    instance_id = "${oci_core_instance.TFInstance.id}"
} 

# Gets the OCID of the first (default) vNIC
data "oci_core_vnic" "InstanceVnic" {
    vnic_id = "${lookup(data.oci_core_vnic_attachments.InstanceVnics.vnic_attachments[0],"vnic_id")}"
}
