output "id" {
  value = "${
    var.compartment_id == "" ?
      join(" ", oci_identity_compartment.compartment.*.id)
    :
      var.compartment_id
  }"
}
