data "oci_identity_availability_domains" "ADs" {
	compartment_id = "${var.tenancy_ocid}"
}

data "oci_identity_compartments" "compartment" {
	compartment_id = "${var.tenancy_ocid}"
	filter {
		name = "name"
		values = [ "${var.compartment_name}" ]
		regex = true
	}
}
