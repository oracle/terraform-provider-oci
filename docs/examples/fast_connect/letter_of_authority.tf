data "oci_core_letter_of_authority" "letter_of_authority" {
  #Required
  cross_connect_id = "${oci_core_cross_connect.cross_connect.id}"
}

output "letter_of_authority" {
  value = {
    id                     = "${data.oci_core_letter_of_authority.letter_of_authority.id}"
    authorized_entity_name = "${data.oci_core_letter_of_authority.letter_of_authority.authorized_entity_name}"
    circuit_type           = "${data.oci_core_letter_of_authority.letter_of_authority.circuit_type}"
    facility_location      = "${data.oci_core_letter_of_authority.letter_of_authority.facility_location}"
    port_name              = "${data.oci_core_letter_of_authority.letter_of_authority.port_name}"
    time_expires           = "${data.oci_core_letter_of_authority.letter_of_authority.time_expires}"
  }
}
