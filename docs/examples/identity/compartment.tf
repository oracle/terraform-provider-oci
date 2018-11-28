/*
 * This example file shows how to create a compartment or reference an existing compartment as a resource.
 *
 * Note: the compartment resource internally resolves name collisions and returns a reference to the preexisting
 * compartment by default. Use `enable_delete` to allow compartment deletion and prevent implicitly importing compartments.
 */

resource "oci_identity_compartment" "compartment1" {
  name           = "tf-example-compartment"
  description    = "compartment created by terraform"
  compartment_id = "${var.tenancy_ocid}"
  enable_delete  = false                              // true will cause this compartment to be deleted when running `terrafrom destroy`
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
