/*
 * This example file shows how to create a compartment or reference an existing compartment as a resource.
 *
 * Note: the compartment resource internally resolves name collisions and returns a reference to the preexisting
 * compartment. Compartments can not be deleted, so removing a compartment resource from your .tf file will only
 * remove it from your statefile.
 */

resource "oci_identity_compartment" "compartment1" {
  name        = "tf-example-compartment"
  description = "compartment created by terraform"
}

data "oci_identity_compartments" "compartments1" {
  compartment_id = "${oci_identity_compartment.compartment1.compartment_id}"

  filter {
    name   = "name"
    values = ["tf-example-compartment"]
  }
}

output "compartments" {
  value = "${data.oci_identity_compartments.compartments1.compartments}"
}
