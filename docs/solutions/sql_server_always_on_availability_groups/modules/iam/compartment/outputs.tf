output "id" {
  value = "${ join(" ", oci_identity_compartment.compartment.*.id) }" 
}
