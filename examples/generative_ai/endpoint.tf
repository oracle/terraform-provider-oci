resource "oci_generative_ai_endpoint" "test_endpoint" {
  #Required
  compartment_id                 = var.compartment_ocid
  dedicated_ai_cluster_id        = data.oci_generative_ai_dedicated_ai_cluster.test_hosting_cluster.id
  model_id                       = local.servering_model_id

  #Optional
  display_name                  = var.test_endpoint_display_name
  description                   = var.test_endpoint_description
  #defined_tags not tested - cannot test in home region        
  freeform_tags                 = var.test_freeform_tags
}

data "oci_generative_ai_endpoint" "test_endpoint" {
  #Required
  endpoint_id                   = oci_generative_ai_endpoint.test_endpoint.id
}

data "oci_generative_ai_endpoints" "test_endpoints" {
  #Required
  compartment_id                = var.compartment_ocid
}



locals {
    filtered_serving_models = [
	for item in data.oci_generative_ai_models.serving_models.model_collection[0].items : item
	  if (
		(item.version == "14.2")
		&& length(item.capabilities) == 1
        && (item.display_name == "cohere.command-light")
	  )
	]

  servering_model_id = local.filtered_serving_models[0].id
}
	
data "oci_generative_ai_models" "serving_models" {
  compartment_id = var.compartment_ocid
  display_name = "cohere.command-light"
}