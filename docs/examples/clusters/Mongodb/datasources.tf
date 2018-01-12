# Gets a list of Availability Domains
data "oci_identity_availability_domains" "ADs" {
  compartment_id = "${var.tenancy_ocid}"
}

# Gets a list of vNIC attachments on the bastion host
data "oci_core_vnic_attachments" "BastionVnics" {
  compartment_id = "${var.compartment_ocid}"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
  instance_id = "${oci_core_instance.MongoDBBastion.id}"
} 

# Gets the OCID of the first (default) vNIC on the bastion host
data "oci_core_vnic" "BastionVnic" {
  vnic_id = "${lookup(data.oci_core_vnic_attachments.BastionVnics.vnic_attachments[0],"vnic_id")}"
}

# Gets a list of vNIC attachments on MongoDBAD1
data "oci_core_vnic_attachments" "MongoDBAD1Vnics" {
  compartment_id = "${var.compartment_ocid}"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
  instance_id = "${oci_core_instance.MongoDBAD1.id}"
}

# Gets the OCID of the first (default) vNIC on MongoDBAD1
data "oci_core_vnic" "MongoDBAD1Vnic" {
  vnic_id = "${lookup(data.oci_core_vnic_attachments.MongoDBAD1Vnics.vnic_attachments[0],"vnic_id")}"
}

# Gets a list of vNIC attachments on MongoDBAD2
data "oci_core_vnic_attachments" "MongoDBAD2Vnics" {
  compartment_id = "${var.compartment_ocid}"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[1],"name")}"
  instance_id = "${oci_core_instance.MongoDBAD2.id}"
}

# Gets the OCID of the first (default) vNIC on MongoDBAD2
data "oci_core_vnic" "MongoDBAD2Vnic" {
  vnic_id = "${lookup(data.oci_core_vnic_attachments.MongoDBAD2Vnics.vnic_attachments[0],"vnic_id")}"
}

