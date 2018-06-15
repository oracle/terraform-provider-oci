### Do not alter this file ###
# datasources.tf - set datasources for the template

# Select the compartment based on the compartment name in variable.tf
data "oci_identity_compartments" "compartment" {
	compartment_id = "${var.tenancy_ocid}"
	filter {
		name = "name"
		values = [ "${var.vcn_state["compartment"]}" ]
		regex = true
	}
}

# Get a list of the ADs (proper AD names)
data "oci_identity_availability_domains" "ad" {
	compartment_id = "${var.tenancy_ocid}"
}

# Get a list of the existing VCNs
data "oci_core_virtual_networks" "vcn" {
	compartment_id = "${data.oci_identity_compartments.compartment.compartments.0.id}"
}

# Get a list of the exising subnets
data "oci_core_subnets" "subnets" {
	compartment_id = "${data.oci_identity_compartments.compartment.compartments.0.id}"
	vcn_id = "${oci_core_virtual_network.vcn.id}"
}
