# Gets a list of Availability Domains
data "oci_identity_availability_domains" "ADs" {
  compartment_id = "${var.tenancy_ocid}"
}

# Gets the OCID of the OS image to use
data "oci_core_images" "OLImageOCID" {
  compartment_id           = "${var.compartment_ocid}"
  operating_system         = "${var.InstanceOS}"
  operating_system_version = "${var.InstanceOSVersion}"
}

data "oci_core_vnic_attachments" "MongoPVNIC" {
  compartment_id      = "${var.compartment_ocid}"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  instance_id = "${oci_core_instance.MongoP.id}"
}

data "oci_core_vnic_attachments" "MongoR1VNIC" {
  compartment_id      = "${var.compartment_ocid}"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  instance_id = "${oci_core_instance.MongoR1.id}"
}

data "oci_core_vnic_attachments" "MongoR2VNIC" {
  compartment_id      = "${var.compartment_ocid}"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  instance_id = "${oci_core_instance.MongoR2.id}"
}

data "oci_core_vnic_attachments" "Pritunl1VNIC" {
  compartment_id      = "${var.compartment_ocid}"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  instance_id = "${oci_core_instance.Pritunl1.id}"
}

data "oci_core_vnic_attachments" "Pritunl2VNIC" {
  compartment_id      = "${var.compartment_ocid}"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  instance_id = "${oci_core_instance.Pritunl2.id}"
}

data "oci_core_vnic_attachments" "PritunllinkVNIC" {
  compartment_id      = "${var.compartment_ocid}"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  instance_id = "${oci_core_instance.Pritunllink.id}"
}

data "oci_core_vnic" "MPVNIC" {
  vnic_id = "${lookup(data.oci_core_vnic_attachments.MongoPVNIC.vnic_attachments[0],"vnic_id")}"
}

data "oci_core_vnic" "MR1VNIC" {
  vnic_id = "${lookup(data.oci_core_vnic_attachments.MongoR1VNIC.vnic_attachments[0],"vnic_id")}"
}

data "oci_core_vnic" "MR2VNIC" {
  vnic_id = "${lookup(data.oci_core_vnic_attachments.MongoR2VNIC.vnic_attachments[0],"vnic_id")}"
}

data "oci_core_vnic" "P1VNIC" {
  vnic_id = "${lookup(data.oci_core_vnic_attachments.Pritunl1VNIC.vnic_attachments[0],"vnic_id")}"
}

data "oci_core_vnic" "P2VNIC" {
  vnic_id = "${lookup(data.oci_core_vnic_attachments.Pritunl2VNIC.vnic_attachments[0],"vnic_id")}"
}

data "oci_core_vnic" "PTLVNIC" {
  vnic_id = "${lookup(data.oci_core_vnic_attachments.PritunllinkVNIC.vnic_attachments[0],"vnic_id")}"
}
