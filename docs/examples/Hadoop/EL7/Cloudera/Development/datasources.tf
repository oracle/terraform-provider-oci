# Gets a list of Availability Domains
data "oci_identity_availability_domains" "ADs" {
  compartment_id = "${var.tenancy_ocid}"
}

# Get list of VNICS for Bastion - Master Nodes
data "oci_core_vnic_attachments" "bastion_vnics" {
  compartment_id      = "${var.compartment_ocid}"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  instance_id = "${oci_core_instance.Bastion.id}"
}

data "oci_core_vnic_attachments" "utility_node_vnics" {
  compartment_id      = "${var.compartment_ocid}"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
  instance_id = "${oci_core_instance.UtilityNode.id}"
}
# Get VNIC ID for first VNIC on Bastion - Master Node 
data "oci_core_vnic" "bastion_vnic" {
  vnic_id = "${lookup(data.oci_core_vnic_attachments.bastion_vnics.vnic_attachments[0],"vnic_id")}"
}
data "oci_core_vnic" "utility_node_vnic" {
  vnic_id = "${lookup(data.oci_core_vnic_attachments.utility_node_vnics.vnic_attachments[0],"vnic_id")}"
}
