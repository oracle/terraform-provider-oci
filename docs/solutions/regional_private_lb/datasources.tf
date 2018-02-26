# Gets a list of Availability Domains
data "oci_identity_availability_domains" "ADs" {
    compartment_id = "${var.tenancy_ocid}"
}

# Gets a list of vNIC attachments of the instance
data "oci_core_vnic_attachments" "Instance1Vnics" {
    compartment_id = "${var.compartment_ocid}"
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD1 - 1],"name")}"
    instance_id = "${oci_core_instance.instance1.id}"
} 

# Gets the OCID of the first (default) vNIC
data "oci_core_vnic" "Instance1Vnic" {
    vnic_id = "${lookup(data.oci_core_vnic_attachments.Instance1Vnics.vnic_attachments[0],"vnic_id")}"
}
