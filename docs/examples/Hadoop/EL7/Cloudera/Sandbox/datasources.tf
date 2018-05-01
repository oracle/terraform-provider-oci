# Gets a list of Availability Domains
data "oci_identity_availability_domains" "ADs" {
  compartment_id = "${var.tenancy_ocid}"
}

# Get list of VNICS for Bastion - Master Nodes
data "oci_core_vnic_attachments" "sandbox_vnics" {
  compartment_id      = "${var.compartment_ocid}"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  instance_id = "${oci_core_instance.Sandbox.id}"
}

# Get VNIC ID for first VNIC on Bastion - Master Node 
data "oci_core_vnic" "sandbox_vnic" {
  vnic_id = "${lookup(data.oci_core_vnic_attachments.sandbox_vnics.vnic_attachments[0],"vnic_id")}"
}
