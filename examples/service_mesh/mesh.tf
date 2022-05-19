resource "oci_service_mesh_mesh" "test_mesh" {
  #Required
  certificate_authorities {
    #Required
    id = var.ca_auth_id
  }
  compartment_id = var.compartment_ocid
  display_name   = local.mesh_display_name

  #Optional
  description   = local.mesh_description
  freeform_tags = { "bar-key" = "value" }
  mtls {
    #Required
    minimum = local.mesh_mtls_minimum
  }
}

data "oci_service_mesh_mesh" "test_mesh" {
  #Required
  mesh_id = oci_service_mesh_mesh.test_mesh.id
}

locals {
  mesh_display_name = "test-mesh"
  mesh_description  = "test mesh description"
  mesh_mtls_minimum = "DISABLED" // allowed values are ("DISABLED", "STRICT", "PERMISSIVE")
  available_state   = "AVAILABLE"
}