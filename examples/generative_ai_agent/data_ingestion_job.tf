resource "oci_generative_ai_agent_data_ingestion_job" "test_data_ingestion_job" {
  #Required
  compartment_id                 = var.compartment_ocid
  data_source_id                 = oci_generative_ai_agent_data_source.test_data_source.id

  #Optional
  display_name                  = var.test_data_ingestion_job_display_name
  description                   = var.test_data_ingestion_job_description
  #defined_tags not tested - cannot test in home region        
  freeform_tags                 = var.test_freeform_tags
}

data "oci_generative_ai_agent_data_ingestion_job" "test_data_ingestion_job" {
  #Required
  data_ingestion_job_id         = oci_generative_ai_agent_data_ingestion_job.test_data_ingestion_job.id
}

data "oci_generative_ai_agent_data_ingestion_jobs" "test_data_ingestion_jobs" {
  #Required
  compartment_id                = var.compartment_ocid
}
