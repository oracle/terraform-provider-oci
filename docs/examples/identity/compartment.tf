/*
 * This example file shows how to create a compartment or reference an existing compartment as a resource.
 * The compartment resource internally resolves name collisions and returns a reference to the preexisting 
 * compartment.   
 */

resource "oci_identity_compartment" "compartment1" {
  name = "-tf-compartment"
  description = "tf test compartment"
}

data "oci_identity_compartments" "compartments1" {
  compartment_id = "${oci_identity_compartment.compartment1.compartment_id}"
  filter {
    name = "name"
    values = ["-tf-compartment"]
  }
}

output "compartments" {
  value = "${data.oci_identity_compartments.compartments1.compartments}"
}
