resource "oci_generative_ai_routing_profile" "test_routing_profile" {
  compartment_id = var.compartment_ocid
  display_name   = "example-routing-profile"

  model_routing_policy {
    allowed_models = ["meta.llama-3-70b-instruct"]
  }

  region_routing_policy {
    allowed_regions = ["us-chicago-1"]
  }
}

data "oci_generative_ai_routing_profile" "test_routing_profile" {
  routing_profile_id = oci_generative_ai_routing_profile.test_routing_profile.id
}

data "oci_generative_ai_routing_profiles" "test_routing_profiles" {
  compartment_id = var.compartment_ocid
  display_name   = oci_generative_ai_routing_profile.test_routing_profile.display_name
  id             = oci_generative_ai_routing_profile.test_routing_profile.id
  state          = "ACTIVE"
}
